package mysql

import (
	"github.com/limoli/dbshift/lib"
	"time"
)

var (
	errBadMigrationDescription = lib.NewError("bad migration description")
	errTransactionAssertion    = lib.NewError("transaction assertion")
)

func (db *Db) HasMigrationStore() bool {
	migrationEntity := new(Migration)
	return db.tx.HasTable(migrationEntity)
}

func (db *Db) HasHistoryStore() bool {
	historyEntity := new(MigrationHistory)
	return db.tx.HasTable(historyEntity)
}

func (db *Db) CreateMigrationStore() error {
	migrationEntity := new(Migration)
	resultMigration := db.tx.
		CreateTable(migrationEntity).
		AddUniqueIndex("script_uidx", "script")
	if resultMigration.Error != nil {
		return resultMigration.Error
	}
	return nil
}

func (db *Db) CreateHistoryStore() error {
	historyEntity := new(MigrationHistory)
	migrationEntity := new(Migration)
	resultHistory := db.tx.
		CreateTable(historyEntity).
		AddForeignKey("migrationId", migrationEntity.TableName()+"(id)", "RESTRICT", "RESTRICT")
	if resultHistory.Error != nil {
		return resultHistory.Error
	}
	return nil
}

func (db *Db) ImportMigrations(migrations []lib.IMigration) (uint, error) {
	var counter uint

	for _, v := range migrations {
		m := &Migration{
			Id:          v.GetId(),
			Uuid:        v.GetUuid(),
			Type:        v.GetType(),
			Description: v.GetDescription(),
			Script:      v.GetScript(),
			CreatedAt:   v.GetCreatedAt(),
		}
		result := db.tx.Where("script=?", v.GetScript()).First(m)
		if result.Error != nil {
			result = db.tx.Create(m)
			if result.Error != nil {
				return counter, result.Error
			}
			counter++
		}
	}

	return counter, nil
}

func (db *Db) GetStatus() (int, int) {
	var counterMigrationUpgrade int
	var counterMigrationDowngrade int

	lastDone := db.getLastExecutedMigration()

	db.tx.Model(new(Migration)).
		Where("type = ? AND id > ?", lib.Upgrade, lastDone.MigrationId).
		Count(&counterMigrationUpgrade)

	db.tx.Model(new(Migration)).
		Where("type = ? AND id < ?", lib.Downgrade, lastDone.MigrationId).
		Count(&counterMigrationDowngrade)

	return counterMigrationUpgrade, counterMigrationDowngrade
}

func (db *Db) CreateMigration(transaction lib.IDbTransaction, sm *lib.SystemMigration) error {
	txWr, ok := transaction.(*TxWrapper)
	if !ok {
		return errTransactionAssertion
	}

	m := &Migration{
		Uuid:        sm.GetUuid(),
		Type:        sm.GetType(),
		Description: sm.GetDescription(),
		Script:      sm.GetScript(),
		CreatedAt:   sm.GetCreatedAt(),
	}

	result := txWr.tx.Create(m)
	err := result.Error
	if err != nil {
		return err
	}

	sm.Id = m.Id
	return nil
}

func (db *Db) GetMigrationsByType(migrationType lib.MigrationType, migrationId *uint64) ([]lib.IMigration, error) {
	var migrations []Migration

	lastDone := db.getLastExecutedMigration()
	query := db.tx.Model(new(Migration))

	switch migrationType {
	case lib.Upgrade:
		query = query.Where("type = ? AND id > ?", migrationType, lastDone.MigrationId).Order("id ASC")
	case lib.Downgrade:
		query = query.Where("type = ? AND id < ?", migrationType, lastDone.MigrationId).Order("id DESC")
	}

	result := query.Find(&migrations)
	if result.Error != nil {
		return nil, result.Error
	}

	var iMigrations []lib.IMigration
	for k := range migrations {
		iMigrations = append(iMigrations, &migrations[k])
	}

	return iMigrations, nil
}

func (db *Db) ExecuteMigration(transaction lib.IDbTransaction, script []byte) (*time.Duration, error) {
	txWr, ok := transaction.(*TxWrapper)
	if !ok {
		return nil, errTransactionAssertion
	}

	timeStart := time.Now()

	result := txWr.tx.Exec(string(script))
	if result.Error != nil {
		return nil, result.Error
	}

	executionTime := time.Since(timeStart)
	return &executionTime, nil
}

func (db *Db) UpdateMigrationHistory(migrationId uint64, executionTime time.Duration) error {
	updatedHistory := MigrationHistory{
		MigrationId:   migrationId,
		ExecutionTime: executionTime.Seconds(),
		CreatedAt:     time.Now(),
	}
	result := db.tx.Create(&updatedHistory)
	return result.Error
}

func (db *Db) getLastExecutedMigration() *MigrationHistory {
	lastDone := new(MigrationHistory)
	result := db.tx.Last(lastDone)
	if result.Error != nil {
		lastDone.MigrationId = 0
	}
	return lastDone
}
