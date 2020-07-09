package databases

import (
	"fmt"
	"myprojects/Shopping-lists-and-recipes/packages/settings"
)

// PostgreSQLGetConnString - получает строку соединения для PostgreSQL
// При начальной настройке строка возвращается без базы данных (она создаётся в процессе)
func PostgreSQLGetConnString(SQLsrv *settings.SQLServer, initialsetup bool) string {

	if initialsetup {
		return fmt.Sprintf("postgres://%v:%v@%v/", SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr)
	}

	return fmt.Sprintf("postgres://%v:%v@%v/%v", SQLsrv.Login, SQLsrv.Pass, SQLsrv.Addr, SQLsrv.DbName)

}
