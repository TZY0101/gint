package cmd

import (
	"gint/cmd_gen"
	"github.com/spf13/cobra"
)

var (
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate api module related fileds",
		RunE: cmd_gen.GenProjByDesc,
	}
)

func init() {
	genCmd.Flags().StringVar(&cmd_gen.VarStringDesc, "desc", "", "The desc file")
}