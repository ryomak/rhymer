package util

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/BurntSushi/toml"
	"errors"
)

type Config struct {
	Port          string
	PublicKeyPath		string
	PrivateKeyPath	string
	DBConfig            DBConfig `toml:"database"`
}

var configure Config

func LoadConfig(filePath string) error {
	if _, err := toml.DecodeFile(filePath, &configure); err != nil {
		return errors.New("Failed to load config : " + err.Error())
	}

	return nil
}

func GetConfig() Config {
	return configure
}


type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Debug    bool
}

func InitDB(conf DBConfig) *gorm.DB {
	param := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" +
		conf.Database + "?parseTime=true&loc=UTC"

	db, err := gorm.Open("mysql", param)
	if err != nil {
		panic(err)
	}
	if conf.Debug{
		db = db.Debug()
	}
	return db
}

