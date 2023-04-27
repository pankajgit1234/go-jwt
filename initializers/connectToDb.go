package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // creted this global variable to make db available in other parts of project
func ConnectToDb() {
	var err error
	//postgres://mlyewbpp:OYY1jeQ2LSpal87lC-rlXaiQPDMEYgaF@john.db.elephantsql.com/mlyewbpp
	dsn := os.Getenv("DB")
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
}
