package main

import (
	"os"
	"testing"

	"github.com/romanyx/polluter"
	"gorm.io/gorm"
	"stockexchange.com/config"
	"stockexchange.com/setup"
)

func exampleSuite(t *testing.T) {
	var db *gorm.DB = config.SetupDatabaseConnection()
	setup.Init(db)
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("failed to convert because: %s", err)
	}

	seed, e := os.Open("dataseed.yaml")

	if e != nil {
		t.Fatalf("failed to open seed file: %v", err)
	}
	defer seed.Close()
	defer config.CloseDatabaseConnection(db)

	p := polluter.New(polluter.MySQLEngine(sqlDB))

	if err := p.Pollute(seed); err != nil {
		t.Fatalf("failed to pollute: %s", err)
	}

}

func TestExample(t *testing.T) {
	t.Parallel()
	defer exampleSuite(t)
}
