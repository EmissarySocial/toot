package toot

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

/******************************************
 * Cross-Package Consistency
 *
 * Every Mastodon endpoint is declared four times: a handler field on API, a
 * path in route, a required scope in scope, and an input struct in txn. Nothing
 * in the compiler ties those four names together, so a typo in any one of them
 * produces a constant that simply never gets used -- which is exactly how
 * "PostAccont_Follow" and "PostAnnoucement_Dismis" survived. These tests parse
 * the source and check that all four agree.
 ******************************************/

// scopeVocabulary lists the constants in the scope package that name an OAuth
// scope itself, rather than the scope required by one endpoint.
var scopeVocabulary = map[string]bool{
	"Public": true, "Private": true, "Read": true, "Write": true,
	"Push": true, "AdminRead": true, "AdminWrite": true,
}

// txnHelpers lists the types in the txn package that are not endpoint inputs.
var txnHelpers = map[string]bool{"QueryPage": true}

// parseGoFiles parses every .go file in a directory, keyed by its path, the way
// parser.ParseDir did before it was deprecated
func parseGoFiles(t *testing.T, dir string) map[string]*ast.File {

	t.Helper()

	entries, err := os.ReadDir(dir)

	if err != nil {
		t.Fatalf("unable to read %s: %v", dir, err)
	}

	fileSet := token.NewFileSet()
	result := make(map[string]*ast.File)

	for _, entry := range entries {

		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		path := filepath.Join(dir, entry.Name())
		file, err := parser.ParseFile(fileSet, path, nil, 0)

		if err != nil {
			t.Fatalf("unable to parse %s: %v", path, err)
		}

		result[path] = file
	}

	return result
}

// parseDir returns the names declared in a package directory, keeping only those
// selected by the provided filter.
func parseDir(t *testing.T, dir string, collect func(ast.Decl, func(string))) map[string]bool {

	t.Helper()

	result := make(map[string]bool)

	for _, file := range parseGoFiles(t, dir) {
		for _, decl := range file.Decls {
			collect(decl, func(name string) { result[name] = true })
		}
	}

	return result
}

// routeNames returns every endpoint constant declared in the named directory.
func routeNames(t *testing.T, dir string) map[string]bool {

	return parseDir(t, dir, func(decl ast.Decl, add func(string)) {

		genDecl, ok := decl.(*ast.GenDecl)

		if !ok || genDecl.Tok != token.CONST {
			return
		}

		for _, spec := range genDecl.Specs {
			if valueSpec, ok := spec.(*ast.ValueSpec); ok {
				for _, name := range valueSpec.Names {
					add(name.Name)
				}
			}
		}
	})
}

// txnNames returns every input struct declared in the txn package.
func txnNames(t *testing.T) map[string]bool {

	return parseDir(t, "txn", func(decl ast.Decl, add func(string)) {

		genDecl, ok := decl.(*ast.GenDecl)

		if !ok || genDecl.Tok != token.TYPE {
			return
		}

		for _, spec := range genDecl.Specs {
			if typeSpec, ok := spec.(*ast.TypeSpec); ok {
				if _, isStruct := typeSpec.Type.(*ast.StructType); isStruct {
					if !txnHelpers[typeSpec.Name.Name] {
						add(typeSpec.Name.Name)
					}
				}
			}
		}
	})
}

// apiNames returns every handler field declared on the API struct.
func apiNames(t *testing.T) map[string]bool {

	t.Helper()

	file, err := parser.ParseFile(token.NewFileSet(), "api.go", nil, 0)

	if err != nil {
		t.Fatalf("unable to parse api.go: %v", err)
	}

	result := make(map[string]bool)

	ast.Inspect(file, func(node ast.Node) bool {

		typeSpec, ok := node.(*ast.TypeSpec)

		if !ok || typeSpec.Name.Name != "API" {
			return true
		}

		structType, ok := typeSpec.Type.(*ast.StructType)

		if !ok {
			return true
		}

		for _, field := range structType.Fields.List {
			if _, isFunc := field.Type.(*ast.FuncType); isFunc {
				for _, name := range field.Names {
					if name.Name != "Authorize" {
						result[name.Name] = true
					}
				}
			}
		}

		return false
	})

	return result
}

func sorted(names map[string]bool) []string {

	result := make([]string, 0, len(names))

	for name := range names {
		result = append(result, name)
	}

	sort.Strings(result)
	return result
}

func TestConsistency_EveryEndpointIsDeclaredFourTimes(t *testing.T) {

	api := apiNames(t)
	route := routeNames(t, "route")
	txn := txnNames(t)

	scope := make(map[string]bool)
	for name := range routeNames(t, "scope") {
		if !scopeVocabulary[name] && !strings.HasPrefix(name, "Read") &&
			!strings.HasPrefix(name, "Write") && !strings.HasPrefix(name, "Admin") {
			scope[name] = true
		}
	}

	if len(api) == 0 || len(route) == 0 || len(scope) == 0 || len(txn) == 0 {
		t.Fatalf("parsed an empty package: api=%d route=%d scope=%d txn=%d",
			len(api), len(route), len(scope), len(txn))
	}

	all := make(map[string]bool)
	for _, set := range []map[string]bool{api, route, scope, txn} {
		for name := range set {
			all[name] = true
		}
	}

	for _, name := range sorted(all) {

		missing := make([]string, 0, 4)

		for _, each := range []struct {
			label string
			names map[string]bool
		}{
			{"api.go", api}, {"route", route}, {"scope", scope}, {"txn", txn},
		} {
			if !each.names[name] {
				missing = append(missing, each.label)
			}
		}

		if len(missing) > 0 {
			t.Errorf("%s is not declared in: %s", name, strings.Join(missing, ", "))
		}
	}
}

// A route constant must hold a path, and a scope constant must not.
func TestConsistency_RouteValuesArePaths(t *testing.T) {

	files := parseGoFiles(t, "route")

	for path, f := range files {
		for _, decl := range f.Decls {

			genDecl, ok := decl.(*ast.GenDecl)

			if !ok || genDecl.Tok != token.CONST {
				continue
			}

			for _, spec := range genDecl.Specs {

				valueSpec, ok := spec.(*ast.ValueSpec)

				if !ok || len(valueSpec.Values) == 0 {
					continue
				}

				literal, ok := valueSpec.Values[0].(*ast.BasicLit)

				if !ok {
					t.Errorf("%s: %s is not a string literal", path, valueSpec.Names[0].Name)
					continue
				}

				if !strings.HasPrefix(literal.Value, `"/api/`) &&
					!strings.HasPrefix(literal.Value, `"/oauth/`) &&
					!strings.HasPrefix(literal.Value, `"/.well-known/`) {
					t.Errorf("%s: %s = %s does not look like an API path",
						path, valueSpec.Names[0].Name, literal.Value)
				}
			}
		}
	}
}

// Every endpoint's scope must be one of the constants the scope package defines,
// never a bare string literal that could drift from the vocabulary.
func TestConsistency_ScopeValuesAreNamedConstants(t *testing.T) {

	files := parseGoFiles(t, "scope")

	for path, file := range files {

		// scope.go declares the vocabulary itself, as string literals.
		if strings.HasSuffix(path, "scope.go") {
			continue
		}

		for _, decl := range file.Decls {

			genDecl, ok := decl.(*ast.GenDecl)

			if !ok || genDecl.Tok != token.CONST {
				continue
			}

			for _, spec := range genDecl.Specs {

				valueSpec, ok := spec.(*ast.ValueSpec)

				if !ok || len(valueSpec.Values) == 0 {
					continue
				}

				if _, isLiteral := valueSpec.Values[0].(*ast.BasicLit); isLiteral {
					t.Errorf("%s: %s is a bare string literal; use a scope constant instead",
						path, valueSpec.Names[0].Name)
				}
			}
		}
	}
}
