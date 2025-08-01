package miniblog

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewMiniBlogCommand create a *cobra.Command object. after, we can use Command object's Execute func to startup the service.
func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "miniblog",
		Short: "miniblog",
		Long: `miniblog service, used to create user with basic information.
			Find more miniblog information at: https://github.com/DanteSu/miniblog`,
		SilenceUsage: true,
		// when cmd.Execute(), call this run fuc
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		// when command runs, do not need to indicate command line parameter
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	return cmd
}

// real entrypoint for service
func run() error {
	fmt.Println("hello miniblog")
	return nil
}
