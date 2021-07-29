package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"import_river/config"
	"import_river/pkg/excel"
)

func Execute(conf *config.Config) {
	connect := GetConnect(conf)
	tx, err := connect.Begin()
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println("connect successful!")
	sql := excel.GetExecuteSql(conf)
	fmt.Println("execute sql ", sql)
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Exec(sql)
	if err != nil {
		panic(err)
	}
}
