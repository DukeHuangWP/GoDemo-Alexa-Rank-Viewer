package flags

import (
	"flag"
)

var Args []string

func init() {
	flag.Parse()
	Args = flag.Args()
}
