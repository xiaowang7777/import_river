package client

import (
	"github.com/sirupsen/logrus"
	"import_river/config"
	"import_river/pkg/database"
	"import_river/pkg/excel"
)

func Run(conf *config.Config) {
	err := excel.ReadExcel(conf.FilePath)
	if err != nil {
		logrus.Fatal(err)
	}
	database.Execute(conf)
}
