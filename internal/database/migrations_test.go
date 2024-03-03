package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrationConfigForDefaultPath(t *testing.T) {
	config := MigrationConfig{
		Driver: "some_driver",
		URL:    "some_url",
	}

	assert.Equal(t, defaultMigrationsPath, config.MigrationPath())
}

func TestCustomMigrationPath(t *testing.T) {
	config := MigrationConfig{
		Driver: "some_driver",
		URL:    "some_url",
		Path:   "some_path",
	}

	assert.Equal(t, "some_path", config.MigrationPath())
}
