package internal

import (
	"bytes"
	"fmt"
	"go/ast"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"golang.org/x/tools/go/packages"
)

func StartCodeGen(schemadir string) error {
	// First, check if the given schemadir is inside a directory named "rdsom".
	rdsomDir := filepath.Dir(schemadir)
	if !strings.HasSuffix(rdsomDir, "rdsom") {
		return fmt.Errorf("internal.RunCodeGen: schema directory %s is not inside an rdsom package", schemadir)
	}
	// Get the Go import path for the schema package.
	// It should look something like "github.com/foo/bar/schemas".
	sPkgPath, err := dirPkgPath(schemadir)
	if err != nil {
		return fmt.Errorf("internal.RunCodeGen: %w", err)
	}
	// This is an estimate of the project's name. It's used for some keys in redis.
	_, err = enclosingPkgName(schemadir)
	if err != nil {
		return fmt.Errorf("internal.RunCodeGen: %w", err)
	}

	// Create the main go file, run it, then delete it.
	timestamp := time.Now().Unix()
	mainFile := filepath.Join(rdsomDir, fmt.Sprintf("main-%d.go", timestamp))
	f, err := os.Create(mainFile)
	if err != nil {
		return fmt.Errorf("internal.RunCodeGen: %w", err)
	}
	defer os.Remove(mainFile)

	snames, err := schemas(sPkgPath)
	if err != nil {
		return fmt.Errorf("internal.RunCodeGen: %w", err)
	}

	if err := NewTemplater(f, "main.gotmpl").Do(
		struct {
			SchemaPkgPath string
			SchemaNames   []string
		}{
			SchemaPkgPath: sPkgPath,
			SchemaNames:   snames,
		},
	); err != nil {
		return fmt.Errorf("internal.RunCodeGen: %w", err)
	}

	_, err = run(mainFile)
	if err != nil {
		return fmt.Errorf("internal.RunCodeGen: %w", err)
	}
	return nil
}

// run 'go run' command and return its output.
func run(target string) (string, error) {
	cmd := exec.Command("go", "run", target)
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("running %s: %s", target, stderr)
	}
	return stdout.String(), nil
}

// schemas returns the names of schemas declared within a package.
func schemas(pkgPath string) ([]string, error) {
	var names []string
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo,
	}, pkgPath)
	if err != nil {
		return nil, fmt.Errorf("internal.schemas: loading schema names: %w", err)
	}
	if npkgs := len(pkgs); npkgs != 1 {
		return nil, fmt.Errorf("internal.schemas: expected 1 package, got %d", npkgs)
	}
	pkg := pkgs[0]
	if len(pkg.Errors) != 0 {
		return nil, fmt.Errorf("internal.schemas: %w", pkg.Errors[0])
	}

	// From here, it's just a clusterf*ck of type casting.
	// The goal is to check that it's a struct with rdsom.Schema embedded.
	for k := range pkg.TypesInfo.Defs {
		if !k.IsExported() || k.Obj == nil {
			continue
		}
		spec, ok := k.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return nil, fmt.Errorf("internal.schemas: invalid declaration %T for %s", k.Obj.Decl, k.Name)
		}
		st, ok := spec.Type.(*ast.StructType)
		if !ok {
			return nil, fmt.Errorf("internal.schemas: invalid spec type %T for %s", spec.Type, k.Name)
		}
		if nfields := len(st.Fields.List); nfields != 1 {
			return nil, fmt.Errorf("internal.schemas: expected 1 field, got %d for %s", nfields, k.Name)
		}
		se, ok := st.Fields.List[0].Type.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		x, ok := se.X.(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf("internal.schemas: invalid declaration %T for %s", se.Sel.Obj.Decl, k.Name)
		}
		if x.Name != "rdsom" || se.Sel.Name != "Schema" {
			continue
		}
		names = append(names, k.Name)
	}

	sort.Strings(names)
	return names, nil
}
