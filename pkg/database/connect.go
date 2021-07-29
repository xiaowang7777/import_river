package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"import_river/config"
	"sync"
)

var (
	once      = &sync.Once{}
	dbConnect *sql.DB
)

func GetConnect(conf *config.Config) *sql.DB {
	once.Do(func() {
		if len(conf.Host) == 0 {
			logrus.Fatal("host must not be nil")
		}
		if len(conf.DatabaseName) == 0 {
			logrus.Fatal("database name must not be nil")
		}
		if len(conf.Username) == 0 {
			logrus.Fatal("database username must not be nil")
		}
		if len(conf.Password) == 0 {
			logrus.Fatal("database password must not be nil")
		}
		connectStr := fmt.Sprintf("user=%s password='%s' host=%s dbname=%s sslmode=disable port=%d", conf.Username, conf.Password, conf.Host, conf.DatabaseName, conf.Port)
		for true {
			db, err := sql.Open("postgres", connectStr)
			if err == nil {
				fmt.Println(fmt.Sprintf("database host:%s,port:%d,dbname:%s,username:%s", conf.Host, conf.Port, conf.DatabaseName, conf.Username))
				dbConnect = db
				break
			} else {
				logrus.Fatal(err)
			}
		}
	})
	return dbConnect
}
