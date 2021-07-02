package version

import (
	"log"
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
)

func ShowBuildVersion() {
	ts, _ := strconv.ParseInt(BuildTS, 10, 32)

	log.Printf("Starting Pravaah\n")
	log.Printf("	Build machine     : %s", BuildMachine)
	log.Printf("	Build user        : %s", BuildUser)
	log.Printf("	Build go version  : %s", GOVersion)
	log.Printf("	Build host arch   : %s", BuildArch)
	log.Printf("	Build target arch : %s", TargetArch)
	log.Printf("	Build timestamp   : %s", time.Unix(ts, 0))
	log.Printf("	Build git info    : %s", GITInfo)
	log.Printf("\n")
}
