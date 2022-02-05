package db

/*
import (
	"github.com/Allan-Nava/go-ddd/config"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	config := config.CONFIGURATION
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("error connection on %s, err: %s", dsn, err.Error())
	}
	db, _ := conn.DB()
	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(max)

	return conn
}
*/
