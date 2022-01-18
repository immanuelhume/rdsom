package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// InitRdsom creates bare schema files. If the rdsom/ folder is not created, create that too.
func InitRdsom(dir string, ss []string) error {
	rdsomroot := filepath.Join(dir, "rdsom")
	sroot := filepath.Join(rdsomroot, "schemas")
	if err := os.MkdirAll(sroot, 0777); err != nil {
		return fmt.Errorf("internal.InitRdsom: initialising rdsom directory: %w", err)
	}
	for _, s := range ss {
		if err := initSchema(s, sroot); err != nil {
			return fmt.Errorf("internal.Rdsom: creating schema file for %q: %w", s, err)
		}
	}
	return nil
}

func initSchema(sname string, dir string) error {
	fname := filepath.Join(dir, strings.ToLower(sname)+".go")
	if _, err := os.Stat(fname); err == nil {
		return fmt.Errorf("internal.initSchema: %w", err)
	}

	f, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("internal.initSchema: %w", err)
	}
	defer f.Close()

	if err := NewTemplater(f, "newschema.gotmpl").Do(
		struct {
			SchemaName string
		}{
			SchemaName: sname,
		},
	); err != nil {
		return fmt.Errorf("internal.initSchema: %w", err)
	}

	return nil
}
