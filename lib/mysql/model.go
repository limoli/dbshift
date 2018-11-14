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

func (m *Migration) GetId() uint64              { return m.Id }
func (m *Migration) GetUuid() string            { return m.Uuid }
func (m *Migration) GetScript() string          { return m.Script }
func (m *Migration) GetDescription() string     { return m.Description }
func (m *Migration) GetType() lib.MigrationType { return m.Type }
func (m *Migration) GetCreatedAt() time.Time    { return m.CreatedAt }

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
