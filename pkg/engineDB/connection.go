package engineDB

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	// registering database driver
	_ "github.com/lib/pq"
)

func getConnection() (*gorm.DB, error) {
	uri := os.Getenv("DATABASE_URI")
	return gorm.Open(postgres.Open(uri), &gorm.Config{})
}
