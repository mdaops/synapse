package cmd

import (
	"github.com/spf13/cobra"
)

type InstallCMD struct {
	CMD *cobra.Command
}

func NewInstallCMD() *InstallCMD {
	cmd := &cobra.Command{
		Use:     "install",
		Short:   "Install Control Plane",
		GroupID: "control_plane",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return &InstallCMD{
		CMD: cmd,
	}
}
