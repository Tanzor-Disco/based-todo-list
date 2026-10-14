package migrations

import (
	"embed"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed *.sql
var migrations embed.FS

// Update updates to the lates available .sql migration
func Update(databaseURL string) error {
	source, err := iofs.New(migrations, ".")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance("iofs", source, databaseURL)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// Reset applies all the down migrations and then uses all up migrations
func Reset(databaseURL string) error {
	source, err := iofs.New(migrations, ".")
	m, err := migrate.NewWithSourceInstance("iofs", source, databaseURL)
	if err != nil {
		return err
	}

	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// Force forces a certain migration
func Force(databaseURL string, migrationNumber int) error {
	source, err := iofs.New(migrations, ".")
	m, err := migrate.NewWithSourceInstance("iofs", source, databaseURL)
	if err != nil {
		return err
	}
	err = m.Force(migrationNumber)
	return err
}
