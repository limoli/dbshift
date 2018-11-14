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

func (m *SystemMigration) GetId() uint64           { return m.Id }
func (m *SystemMigration) GetUuid() string         { return m.Uuid }
func (m *SystemMigration) GetScript() string       { return m.Script }
func (m *SystemMigration) GetDescription() string  { return m.Description }
func (m *SystemMigration) GetType() MigrationType  { return m.Type }
func (m *SystemMigration) GetCreatedAt() time.Time { return m.CreatedAt }
