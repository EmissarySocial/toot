package txn

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// TestFormTagsHaveMatchingJSONTags requires every form-tagged field to carry the same name as a json tag.
// Clients such as Ice Cubes post JSON bodies, which the JSON binder reads only through json tags.
func TestFormTagsHaveMatchingJSONTags(t *testing.T) {

	packages, err := parser.ParseDir(token.NewFileSet(), ".", func(info fs.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, 0)

	if err != nil {
		t.Fatalf("parsing txn: %v", err)
	}

	checked := 0

	for _, pkg := range packages {
		ast.Inspect(pkg, func(node ast.Node) bool {

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
