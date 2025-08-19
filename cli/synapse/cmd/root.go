package cmd

import (
	"github.com/spf13/cobra"
)

const Version = "0.1.0"

type RootCMD struct {
	Name    string
	Version string
	CMD     *cobra.Command
}

func NewRootCMD() *RootCMD {
	cmd := &cobra.Command{
		Use:   "synapse",
		Short: "Synapse CLI",
		Long:  "Synapse CLI",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddGroup(&cobra.Group{
		ID:    "control_plane",
		Title: "Control Plane",
	})

	cmd.AddCommand(NewInstallCMD().CMD)

	return &RootCMD{
		Name:    "synapse",
		Version: Version,
		CMD:     cmd,
	}
}
