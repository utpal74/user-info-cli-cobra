package search

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "user-info",
	Short: "user crud operation",
	Long: `store user information.
	can create a user
	can display a user by name if queried.
	`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil{
		fmt.Printf("trouble executing comamnd: %v\n", err)
		os.Exit(1)
	}
}
