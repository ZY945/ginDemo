package untils

import (
	"GinDemo/global"
)

func GenID() int64 {
	return global.Node.Generate().Int64()
}
