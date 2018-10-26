package model

import "github.com/jinzhu/gorm"

const(
	DBKEY="DB"
)

func Migrate(db *gorm.DB)error{
	db.AutoMigrate(Word{})
	db.AutoMigrate(RhymeWord{})
	return nil
}