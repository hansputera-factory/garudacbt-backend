package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"hanifu.id/hansputera-factory/garudacbt-backend/config"
)

type Database interface {
	GetDb() *Queries
}

type garudaDatabase struct {
	Db *Queries
}

var (
	once       sync.Once
	dbInstance *garudaDatabase
)

func NewGarudaDatabase(conf *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.DBName,
		)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}

		dbInstance = &garudaDatabase{
			Db: New(db),
		}
	})

	return dbInstance
}

func (g *garudaDatabase) GetDb() *Queries {
	return g.Db
}
