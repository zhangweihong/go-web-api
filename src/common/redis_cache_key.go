//redis cache key
package common

import "gin-framework/basic/src/tool"

const AppPrefix string = "APPPREFIX_"

func TokenKey(id string) string {
	return tool.Splicing(AppPrefix, id)
}
