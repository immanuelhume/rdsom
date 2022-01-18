package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/immanuelhume/rdsom/internal"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{Use: "rdsom"}
	cmd.AddCommand(_initCmd)
	cmd.Execute()
}

var _initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize rdsom",
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one schema name")
		}
		return nil
	},
	RunE: func(_ *cobra.Command, args []string) error {
		rootdir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("initializing rdsom: %w", err)
		}
		if err := internal.InitRdsom(rootdir, args); err != nil {
			return fmt.Errorf("initializing rdsom: %w", err)
		}
		return nil
	},
}

var _genCmd = &cobra.Command{
	Use:   "gen",
	Short: "(Re)generate the ORM code",
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires path to schemas")
		}
		if _, err := os.Stat(args[0]); err != nil {
			return err
		}
		return nil
	},
	RunE: func(_ *cobra.Command, args []string) error {
		schemadir, err := filepath.Abs(args[0])
		if err != nil {
			return fmt.Errorf("running codegen: %w", err)
		}
		if err := internal.StartCodeGen(schemadir); err != nil {
			return fmt.Errorf("running codegen: %w", err)
		}
		return nil
	},
}
