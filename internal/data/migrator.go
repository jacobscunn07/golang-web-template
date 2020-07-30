package data

import (
  "github.com/golang-migrate/migrate/v4"
  "log"
)

func Migrate(databaseURL string) {
  m, err := migrate.New("file:///app/internal/data/migrations",databaseURL)
  if err != nil {
    log.Fatal(err)
  }

  if err := m.Up(); err != nil && err != migrate.ErrNoChange {
    log.Fatal(err)
  }
}
