package agent

import (
	"fmt"

	"github.com/aler9/gortsplib"
	"github.com/aler9/gortsplib/pkg/base"
)

func SourceQuery(path string) {
	url, _ := base.ParseURL(path)

	conn, _ := gortsplib.Dial(url.Scheme, url.Host)

	defer conn.Close()

	conn.Options(url)

	// Find tracks
	tracks, _, _, _ := conn.Describe(url)

	for _, track := range tracks {
		fmt.Print(track.Media.ConnectionInformation)
	}
}
