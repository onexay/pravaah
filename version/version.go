package version

import (
	"fmt"
	"strconv"
	"time"
)

var (
	BuildMachine string
	BuildUser    string
	GOVersion    string
	BuildArch    string
	TargetArch   string
	GITInfo      string
	BuildTS      string
	DevBuild     string
)

var IsDevBuild bool

func ShowBuildVersion() {
	ts, _ := strconv.ParseInt(BuildTS, 10, 32)
	IsDevBuild, _ = strconv.ParseBool(DevBuild)

	fmt.Printf("Pravaah RTP streaming framework.\n")
	fmt.Printf("Akshay Ranjan <akshay.ranjan@gmail.com>\n")
	fmt.Printf("\n")
	fmt.Printf("	Build machine     : %s\n", BuildMachine)
	fmt.Printf("	Build user        : %s\n", BuildUser)
	fmt.Printf("	Build go version  : %s\n", GOVersion)
	fmt.Printf("	Build host arch   : %s\n", BuildArch)
	fmt.Printf("	Build target arch : %s\n", TargetArch)
	fmt.Printf("	Build timestamp   : %s\n", time.Unix(ts, 0))
	fmt.Printf("	Build git info    : %s\n", GITInfo)
	if IsDevBuild {
		fmt.Printf("	Build type        : development\n")
	} else {
		fmt.Printf("	Build type        : release\n")
	}
	fmt.Printf("\n")
}
