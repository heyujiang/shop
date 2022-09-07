package setting

import "fmt"

const AppName = "shop"

var (
	GitInfo    string
	AppVersion string
)

func init() {
	GitInfo = fmt.Sprintf("")
	AppVersion = fmt.Sprintf("")
}
