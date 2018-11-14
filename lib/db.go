package lib

import (
	"time"
)

const (
	DatabaseTypeMySQL DatabaseType = "mysql"
)

var (
	AvailableDatabases = []DatabaseType{DatabaseTypeMySQL}
)

type DatabaseType string

type IDb interface {
	IDbAction
	IDbCheck
	IDbConnection
	IDbImport
	IDbScript
}

type IDbAction interface {
	CreateMigration(transaction IDbTransaction, m *SystemMigration) error
	ExecuteMigration(transaction IDbTransaction, script []byte) (*time.Duration, error)
	GetMigrationsByType(migrationType MigrationType, migrationId *uint64) ([]IMigration, error)
	GetStatus() (int, int)
	GetTransaction() IDbTransaction
	UpdateMigrationHistory(migrationId uint64, executionTime time.Duration) error
}

type IDbCheck interface {
	HasMigrationStore() bool
	HasHistoryStore() bool
	CreateMigrationStore() error
	CreateHistoryStore() error
}

type IDbConnection interface {
	Connect(IDbConfig) error
}

type IDbImport interface {
	ImportMigrations([]IMigration) (uint, error)
}

type IDbScript interface {
	GetExtension() string
}

type IDbTransaction interface {
	Rollback() error
	Commit() error
}

type IMigration interface {
	GetId() uint64
	GetUuid() string
	GetScript() string
	GetDescription() string
	GetType() MigrationType
	GetCreatedAt() time.Time
}
