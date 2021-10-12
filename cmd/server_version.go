package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type serverVersionCmd struct {
	out io.Writer
}

// NewServerVersionCommand creates the command for rendering the Kubernetes server version.
func NewServerVersionCommand(streams genericclioptions.IOStreams) *cobra.Command {
	helloWorldCmd := &serverVersionCmd{
		out: streams.Out,
	}

	cmd := &cobra.Command{
		Use:          "server-version",
		Short:        "Prints Kubernetes server version",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) != 0 {
				return errors.New("this command does not accept arguments")
			}
			return helloWorldCmd.run()
		},
	}

	cmd.AddCommand(newVersionCmd(streams.Out))
	return cmd
}

func (sv *serverVersionCmd) run() error {
	serverVersion, err := getServerVersion()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sv.out, "Hello from Kubernetes server with version %s!\n", serverVersion)
	if err != nil {
		return err
	}
	return nil
}
