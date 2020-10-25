package result

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mkhstar/inmemdb/platform"
)

// Echo returns a string format of the ResultInfo
func (resultInfo *Info) Echo() string {
	if resultInfo.Error != nil {
		return fmt.Sprintf("(error) %s", resultInfo.Error)
	}
	switch v := resultInfo.Result.(type) {
	case nil:
		return "nil"
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case [0]string:
		return "(empty array)"
	case []string:
		var result strings.Builder
		for i, val := range v {
			fmt.Fprintf(&result, "%d) %s", i+1, val)
			if i != (len(v) - 1) {
				fmt.Fprintf(&result, "%s", platform.LineBreak)
			}
		}
		return result.String()
	default:
		panic("Unknown type")
	}
}
