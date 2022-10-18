package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryungmin/protocol-go/pkg/protocol"
)

var APPLICATION_NAME string = "protocol"

func init() {
	APPLICATION_NAME = filepath.Base(os.Args[0])
}

// Current version
var APPLICATION_VERSION string

func main() {
	if len(os.Args) == 2 {
		lowerArg1 := strings.ToLower(os.Args[1])

		if lowerArg1 == "-l" || lowerArg1 == "--list" {
			displayListOfSupportedProtocols(APPLICATION_NAME)
			os.Exit(0)
		}

		// protocol
		p, ok := protocol.Protocols[lowerArg1]

		specs := p.Specs

		if !ok {
			// specs
			specs = os.Args[1]
		}

		proto, err := protocol.NewProtocol(specs)

		if err != nil {
			os.Exit(1)
		}

		fmt.Println(proto)
	} else {
		display_help()
	}
}
