package basic

import (
	"fmt"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints 'Hello World'",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}

var rootCmd = &cobra.Command{
	Use:     "bf-cli",
	Short:   "bf-cli is a command line tool for breeze-functions",
	Long:    "Breeze Functions CLI is a command line tool for breeze-functions.",
	Example: "bf-cli subcommand [OPTIONS]",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
