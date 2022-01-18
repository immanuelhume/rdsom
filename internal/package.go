package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/tools/go/packages"
)

var schemaEnclosingPkg = regexp.MustCompile(".+/(.+)/rdsom/schemas$")

func pkgPath(fname string) (string, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName,
	}, fmt.Sprintf("file=%s", fname))
	if err != nil {
		return "", fmt.Errorf("internal.pkgPath: finding enclosing packages for %s: %w", fname, err)
	}
	if len(pkgs) == 0 {
		return "", fmt.Errorf("internal.pkgPath: %s is not in a valid Go package", fname)
	}
	return pkgs[0].PkgPath, nil
}

// enclosingPkgName returns the name of the "root" package enclosing
// an rdsom directory.
//
// E.g. given path/to/gopkg/rdsom/schemas, return "gopkg".
func enclosingPkgName(dirname string) (string, error) {
	fname, err := firstGoFile(dirname)
	if err != nil {
		return "", fmt.Errorf("internal.enclosingPkgName: %w", err)
	}
	pkg, err := pkgPath(fname)
	if err != nil {
		return "", fmt.Errorf("internal.enclosingPkgName: %w", err)
	}
	matches := schemaEnclosingPkg.FindStringSubmatch(pkg)
	if len(matches) < 2 {
		return "", fmt.Errorf("internal.enclosingPkgName: %s is not a valid schema pkgPath", pkg)
	}
	return matches[1], nil
}

// dirPkgPath attempts to find the Go module path of a given directory.
//
// E.g. given "/home/immanuelhume/rdsom/gen", it may return "github.com/immanuelhume/rdsom/gen".
func dirPkgPath(dirname string) (string, error) {
	fname, err := firstGoFile(dirname)
	if err != nil {
		return "", fmt.Errorf("internal.dirPkgPath: %w", err)
	}
	return pkgPath(fname)
}

// firstGoFile tries to find a .go file inside a directory.
// Returns error if no Go files are found.
func firstGoFile(dirname string) (string, error) {
	fls, err := os.ReadDir(dirname)
	if err != nil {
		return "", fmt.Errorf("internal.firstGoFile: %w", err)
	}
	dirname, err = filepath.Abs(dirname)
	if err != nil {
		return "", fmt.Errorf("internal.firstGoFile: %w", err)
	}
	for _, fl := range fls {
		if filepath.Ext(fl.Name()) == ".go" {
			return filepath.Join(dirname, fl.Name()), nil
		}
	}
	return "", fmt.Errorf("internal.firstGoFile: no go files in %s", dirname)
}
