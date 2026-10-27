package txn

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// TestFormTagsHaveMatchingJSONTags requires every form-tagged field to carry the same name as a json tag.
// Clients such as Ice Cubes post JSON bodies, which the JSON binder reads only through json tags.
func TestFormTagsHaveMatchingJSONTags(t *testing.T) {

	files := parseGoFiles(t, ".", false)

	checked := 0

	for _, source := range files {
		ast.Inspect(source, func(node ast.Node) bool {

			field, ok := node.(*ast.Field)
			if !ok || field.Tag == nil {
				return true
			}

			raw, err := strconv.Unquote(field.Tag.Value)
			if err != nil {
				t.Errorf("unreadable tag %s", field.Tag.Value)
				return true
			}

			tag := reflect.StructTag(raw)
			form, hasForm := tag.Lookup("form")
			if !hasForm {
				return true
			}

			// A form array is named "x[]"; its JSON name is plain "x"
			want := strings.TrimSuffix(form, "[]")
			name := fieldName(field)
			checked++

			if got, hasJSON := tag.Lookup("json"); !hasJSON {
				t.Errorf("field %s has form:%q but no json tag", name, form)
			} else if got != want {
				t.Errorf("field %s has form:%q but json:%q (want %q)", name, form, got, want)
			}

			return true
		})
	}

	if checked == 0 {
		t.Fatal("no form-tagged fields found; the parser filter is wrong")
	}
}

// fieldName returns a field's name for error messages, or "(embedded)" when it has none.
func fieldName(field *ast.Field) string {
	if len(field.Names) == 0 {
		return "(embedded)"
	}
	return field.Names[0].Name
}

// TestPathParamsHaveParamTags requires each ":name" in a route path to have a field tagged
// param:"name" on the request type of the same name
func TestPathParamsHaveParamTags(t *testing.T) {

	// Echo binds path values only through param tags, so any other tag (such as uri:"id") leaves the field empty

	routes := routePaths(t)
	checked := 0

	files := parseGoFiles(t, ".", false)

	for _, source := range files {
		ast.Inspect(source, func(node ast.Node) bool {

			spec, ok := node.(*ast.TypeSpec)
			if !ok {
				return true
			}

			path, found := routes[spec.Name.Name]
			structType, isStruct := spec.Type.(*ast.StructType)

			if !found || !isStruct {
				return true
			}

			params := map[string]bool{}

			for _, field := range structType.Fields.List {
				if field.Tag == nil {
					continue
				}
				raw, _ := strconv.Unquote(field.Tag.Value)
				if name, ok := reflect.StructTag(raw).Lookup("param"); ok {
					params[name] = true
				}
			}

			for _, segment := range strings.Split(path, "/") {
				if name, isParam := strings.CutPrefix(segment, ":"); isParam {
					checked++
					if !params[name] {
						t.Errorf("%s: route %s has :%s but no field tagged param:%q", spec.Name.Name, path, name, name)
					}
				}
			}

			return true
		})
	}

	if checked == 0 {
		t.Fatal("no path parameters checked; the route lookup is wrong")
	}
}

// routePaths returns every route constant in ../route, keyed by its name
func routePaths(t *testing.T) map[string]string {

	files := parseGoFiles(t, "../route", true)

	result := map[string]string{}

	for _, source := range files {
		ast.Inspect(source, func(node ast.Node) bool {

			spec, ok := node.(*ast.ValueSpec)
			if !ok || len(spec.Names) != 1 || len(spec.Values) != 1 {
				return true
			}

			if literal, ok := spec.Values[0].(*ast.BasicLit); ok && literal.Kind == token.STRING {
				if value, err := strconv.Unquote(literal.Value); err == nil {
					result[spec.Names[0].Name] = value
				}
			}

			return true
		})
	}

	return result
}

// parseGoFiles parses the .go files in a directory, leaving out _test.go files unless withTests is set
func parseGoFiles(t *testing.T, dir string, withTests bool) []*ast.File {

	t.Helper()

	entries, err := os.ReadDir(dir)

	if err != nil {
		t.Fatalf("reading %s: %v", dir, err)
	}

	fileSet := token.NewFileSet()
	result := make([]*ast.File, 0, len(entries))

	for _, entry := range entries {

		name := entry.Name()

		if entry.IsDir() || !strings.HasSuffix(name, ".go") || (!withTests && strings.HasSuffix(name, "_test.go")) {
			continue
		}

		file, err := parser.ParseFile(fileSet, filepath.Join(dir, name), nil, 0)

		if err != nil {
			t.Fatalf("parsing %s: %v", name, err)
		}

		result = append(result, file)
	}

	return result
}
