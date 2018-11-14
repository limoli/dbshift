package lib

import "time"

const (
	Downgrade       MigrationType = 0
	Upgrade         MigrationType = 1
	DowngradeString string        = "down"
	UpgradeString   string        = "up"
)

type MigrationType int

func (t MigrationType) String() string {
	switch t {
	case Downgrade:
		return DowngradeString
	case Upgrade:
		return UpgradeString
	}
	return ""
}

type SystemMigration struct {
	Id          uint64        `yaml:"id"`
	Uuid        string        `yaml:"uuid"`
	Script      string        `yaml:"script"`
	Description string        `yaml:"description"`
	Type        MigrationType `yaml:"type"`
	CreatedAt   time.Time     `yaml:"createdAt"`
}

// Get migration id
func (m *SystemMigration) GetId() uint64           { return m.Id }

// Get migration uuid
func (m *SystemMigration) GetUuid() string         { return m.Uuid }

// Get migration script
func (m *SystemMigration) GetScript() string       { return m.Script }

// Get migration description
func (m *SystemMigration) GetDescription() string  { return m.Description }

// Get migration type
func (m *SystemMigration) GetType() MigrationType  { return m.Type }

// Get migration creation time
func (m *SystemMigration) GetCreatedAt() time.Time { return m.CreatedAt }
