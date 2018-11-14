package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/limoli/dbshift/lib"
)

type Db struct {
	tx *gorm.DB
}

// Connect to MySQL database
func (db *Db) Connect(config lib.IDbConfig) error {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		config.GetUser(),
		config.GetPassword(),
		config.GetHost(),
		config.GetPort(),
		config.GetName())

	db.tx, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	db.tx.LogMode(false)
	return nil
}

// Get MySQL extension
func (db *Db) GetExtension() string {
	return "sql"
}

// Get MySQL transaction
func (db *Db) GetTransaction() lib.IDbTransaction {
	return &TxWrapper{db.tx.Begin()}
}

type TxWrapper struct {
	tx *gorm.DB
}

// Commit a MySQL transaction
func (tx *TxWrapper) Commit() error {
	result := tx.tx.Commit()
	return result.Error
}

// Rollback a MySQL transaction
func (tx *TxWrapper) Rollback() error {
	result := tx.tx.Rollback()
	return result.Error
}
