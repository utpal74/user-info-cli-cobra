package search

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/user-info-cli-tool/internal/controller/user"
	"github.com/user-info-cli-tool/internal/repository/memory"
)

var createCmd = &cobra.Command{
	Use: "create",
	Short: "create a user",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		repo := memory.New()
		ctrl := user.New(repo)
		if err := ctrl.CreateUser(ctx, args[0], args[1]); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("User created.")
	},
}

func init(){
	rootCmd.AddCommand(createCmd)
}
