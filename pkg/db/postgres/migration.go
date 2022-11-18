package postgres

import (
	"fmt"
	gomigratev4 "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"strings"
)

type DBMigrate struct {
	Path string
	DSN  string
}

func NewMigrate(dbDsn string, migrationsDir string) *DBMigrate {
	return &DBMigrate{
		Path: migrationsDir,
		DSN:  dbDsn,
	}
}

func (m *DBMigrate) GetPath() (string, error) {
	mydir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	migrationsFilePath := "file://" + mydir + m.Path
	return migrationsFilePath, nil
}

func (m *DBMigrate) Up() error {
	mydir, err := m.GetPath()
	if err != nil {
		return err
	}
	mig, err := gomigratev4.New(mydir, m.DSN)
	if err != nil {
		return err
	}

	err = mig.Up()
	if err == nil {
		fmt.Println("Database migrations are successfull")
	} else if strings.Contains(err.Error(), "no change") {
		fmt.Println("No pending changes to apply on database migration")
		return nil
	}
	return err
}

func (m *DBMigrate) Down() error {
	mydir, err := m.GetPath()
	if err != nil {
		return err
	}

	mig, err := gomigratev4.New(mydir, m.DSN)
	mig.Steps(1)
	if err != nil {
		return err
	}
	return mig.Down()
}
