package excel

import (
	"fmt"
	"import_river/config"
)

var sql = "insert into "

func GetExecuteSql(conf *config.Config) string {

	sql += conf.TableName + "(point) values (st_geomfromtext('Polygon(("
	var i int64 = 1
	info := leftShore[i]
	for info.check() {
		sql += fmt.Sprintf("%s %s,", info.Longitude, info.Latitude)
		i++
		info = leftShore[i]
	}
	i--
	info = rightShore[i]
	for info.check() {
		sql += fmt.Sprintf("%s %s,", info.Longitude, info.Latitude)
		i--
		info = rightShore[i]
	}
	sql += fmt.Sprintf("%s %s", leftShore[1].Longitude, leftShore[1].Latitude)
	sql += "))'));"
	return sql
}
