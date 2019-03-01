package cli

import (
	gpflag "github.com/octago/sflags/gen/gpflag"
	cobra "github.com/spf13/cobra"
)

// flag definitions here
// https://github.com/octago/sflags#flags-based-on-structures------
type MainConfig struct {
}

func Main() *cobra.Command {
	var flags MainConfig
	var cmd = &cobra.Command{
		Use: "flex",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	if err := gpflag.ParseTo(&flags, cmd.PersistentFlags()); err != nil {
		panic(err)
	}
	return cmd
}
