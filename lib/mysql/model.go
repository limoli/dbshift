package mysql

import (
	"github.com/limoli/dbshift/lib"
	"time"
)

type Migration struct {
	Id          uint64            `gorm:"primary_key" json:"id"`
	Uuid        string            `json:"uuid"`
	Script      string            `gorm:"index" json:"script"`
	Description string            `json:"description"`
	Type        lib.MigrationType `json:"type"`
	CreatedAt   time.Time         `gorm:"column:createdAt"`
}

// Get migration id
func (m *Migration) GetId() uint64              { return m.Id }

// Get migration uuid
func (m *Migration) GetUuid() string            { return m.Uuid }

// Get migration script
func (m *Migration) GetScript() string          { return m.Script }

// Get migration description
func (m *Migration) GetDescription() string     { return m.Description }

// Get migration type
func (m *Migration) GetType() lib.MigrationType { return m.Type }

// Get migration creation time
func (m *Migration) GetCreatedAt() time.Time    { return m.CreatedAt }

// Get migration table name
func (m *Migration) TableName() string {
	return "dbMigration"
}

type MigrationHistory struct {
	Id            uint64    `gorm:"primary_key" json:"id"`
	MigrationId   uint64    `gorm:"column:migrationId" json:"migrationId"`
	ExecutionTime float64   `gorm:"column:executionTime" json:"executionTime"`
	CreatedAt     time.Time `gorm:"column:createdAt"`
	Migration     Migration
}

func (m *MigrationHistory) TableName() string {
	return "dbHistory"
}
