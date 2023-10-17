package cmd

import (
	"gint/cmd_init"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Quickly build project structure",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd_init.InitProjDir(args[0])
		},
	}
)
