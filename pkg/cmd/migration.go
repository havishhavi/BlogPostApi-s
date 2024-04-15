package cmd

import (
	"github.com/spf13/cobra"
	"www.blog.com/pkg/migration"
)

func init() {
	rootCmd.AddCommand(migrationCmd)
}

var migrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database migration",
	Long:  "Database table creation",
	Run: func(cmd *cobra.Command, args []string) {
		migration.Migrate()
	},
}
