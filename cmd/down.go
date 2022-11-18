package cmd

import (
	"fmt"
	"github.com/leapsquare/sheet-service/pkg/dbmigrate"
	"github.com/leapsquare/sheet-service/pkg/utils"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := utils.LoadAndParseCfgFile()
		DSN := cfg.Postgres.DSN
		path := cfg.Postgres.MigrationPath
		fmt.Println("down called")
		m := dbmigrate.NewMigrate(DSN, path)
		err := m.Down()
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {

	migrateCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
