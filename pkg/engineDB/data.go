package engineDB

import (
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

// Data manages the connection to the database.
type Data struct {
	DB *gorm.DB
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	data = &Data{DB: db}
}

func Factory() *Data {
	once.Do(initDB)
	return data
}
