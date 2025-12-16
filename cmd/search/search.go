package search

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/utpal74/user-info-cli-cobra/internal/controller/user"
	"github.com/utpal74/user-info-cli-cobra/internal/repository/memory"
)

var caseSensitive bool

var searchCmd = &cobra.Command{
	Use: "search <name> [--case-sensitive]",
	Short: "find user by name",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		repo := memory.New()
		ctrl := user.New(repo)
		searchKey := name
		if !caseSensitive {
			searchKey = strings.ToLower(name)
		}
		u, err := ctrl.GetUser(context.Background(), searchKey)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("User found: %+v\n", u)
	},
}

func init(){
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&caseSensitive, "case-sensitive", "c", false, "Perform case-sensitive search")
}
