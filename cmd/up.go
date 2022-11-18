package cmd

import (
	"github.com/leapsquare/sheet-service/pkg/db/postgres"
	"github.com/leapsquare/sheet-service/pkg/utils"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := utils.LoadAndParseCfgFile()
		DSN := cfg.Postgres.DSN
		path := cfg.Postgres.MigrationPath
		m := postgres.NewMigrate(DSN, path)
		err := m.Up()
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	migrateCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
