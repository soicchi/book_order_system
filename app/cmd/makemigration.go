package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const content = `package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var {{.MigrationName}} = &gormigrate.Migration{
	ID: "{{.MigrationID}}",
	Migrate: func(tx *gorm.DB) error {
		// Write your migration code here
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		// Write your rollback code here
		return nil
	},
}
`
const targetDir = "internal/infrastructure/postgres/database/migrations"

var migrationName string

type MigrationInfo struct {
	MigrationID   string
	MigrationName string
}

// Migration用のファイルを作成するコマンドを定義
var makemigrationCmd = &cobra.Command{
	Use:   "makemigration",
	Short: "A brief description of your command",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("makeMigration called")

		// generate file name based on current time and migration name
		now := time.Now()
		migrationID := now.Format("20060102150405")
		fileName := fmt.Sprintf("%s_%s.go", migrationID, migrationName)

		// create file
		file, err := os.Create(fmt.Sprintf("%s/%s", targetDir, fileName))
		if err != nil {
			return err
		}

		defer file.Close()

		migrationInfo := MigrationInfo{
			MigrationID:   migrationID,
			MigrationName: toCamelCase(migrationName),
		}

		// create template
		t := template.Must(template.New("migration").Parse(content))

		// map the migrationInfo to the template and write it to file
		if err := t.Execute(file, migrationInfo); err != nil {
			return err
		}

		return nil
	},
}

func toCamelCase(s string) string {
	words := strings.Split(s, "_")
	for i, word := range words {
		words[i] = cases.Title(language.English).String(word)
	}

	return strings.Join(words, "")
}

func init() {
	rootCmd.AddCommand(makemigrationCmd)
	makemigrationCmd.Flags().StringVarP(&migrationName, "name", "n", "", "migration name")
}
