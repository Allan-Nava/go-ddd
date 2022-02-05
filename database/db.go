package db
/*
* Copyright © 2022 Allan Nava <>
* Created 05/02/2022
* Updated 05/02/2022
*
 */
/*
import (
	"github.com/Allan-Nava/go-ddd/config"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	//
	config := config.CONFIGURATION
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.dbUsername, config.dbPassword, config.dbHost, config.dbPort, config.dbName )
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("error connection on %s, err: %s", dsn, err.Error())
	}
	db, _ := conn.DB()
	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(max)
	//
	return conn
}
*/
