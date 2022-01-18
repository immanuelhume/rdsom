package internal_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/immanuelhume/rdsom/internal"
)

func TestInitRdsom(t *testing.T) {
	cases := []struct {
		desc      string
		snames    []string
		filenames []string // files to be created
	}{
		{
			desc:      "works with lower case args",
			snames:    []string{"pet", "owner"},
			filenames: []string{"pet.go", "owner.go"},
		},
		{
			desc:      "works with upper case args",
			snames:    []string{"Shiba", "Inu"},
			filenames: []string{"shiba.go", "inu.go"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			dir := t.TempDir()
			if err := internal.InitRdsom(dir, tt.snames); err != nil {
				t.Fatal(err)
			}
			for _, f := range tt.filenames {
				if _, err := os.Stat(filepath.Join(dir, "rdsom", "schemas", f)); err != nil {
					t.Error(err)
				}
			}
		})
	}
}
