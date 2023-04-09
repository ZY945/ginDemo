package snowflake

import (
	"GinAndSqlx/global"
)

func GenID() int64 {
	return global.Node.Generate().Int64()
}
