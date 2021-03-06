package cmd

import (
	"github.com/spf13/cobra"
)

var idmCmd = &cobra.Command{
	Use:   "idm",
	Short: "Identity Management commands",
	Long:  `
Commands to manage users, groups and roles. 

See the help of respective sub-commands for further details.
`,
	Run: func(cm *cobra.Command, args []string) {
		cm.Usage()
	},
}

func init() {
	RootCmd.AddCommand(idmCmd)
}
