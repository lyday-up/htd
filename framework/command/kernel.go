package command

import "github.com/htd/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(DemoCommand)

	root.AddCommand(initAppCommand())
	root.AddCommand(initEnvCommand())
	root.AddCommand(initConfigCommand())

}
