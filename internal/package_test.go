package internal

import (
	"testing"
)

func TestEnclosingPkgName(t *testing.T) {
	pkgname, err := enclosingPkgName("testdata/pkga/rdsom/schemas/")
	if err != nil {
		t.Fatal(err)
	}
	if pkgname != "pkga" {
		t.Errorf("got %q, want %q", pkgname, "pkga")
	}
}

func TestDirPkgPath(t *testing.T) {
	pkgpath, err := dirPkgPath(".")
	if err != nil {
		t.Fatal(err)
	}
	if pkgpath != "github.com/immanuelhume/rdsom/internal" {
		t.Errorf("got %s, want %s", pkgpath, "github.com/immanuelhume/rdsom/internal")
	}
}
