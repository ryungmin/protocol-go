package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/ryungmin/protocol-go/pkg/protocol"
)

var APPLICATION_NAME string = "protocol"

const APPLICATION_DESC string = "Protocol is a command-line tool to display ASCII RFC-like protocol header diagrams."

func init() {
	APPLICATION_NAME = filepath.Base(os.Args[0])
}

// Current version
var APPLICATION_VERSION string

func main() {
	parse()

	lowerArg1 := strings.ToLower(os.Args[1])

	// protocol
	p, ok := protocol.Protocols[lowerArg1]

	specs := p.Specs

	if !ok {
		// specs
		specs = os.Args[1]
	}

	proto, err := protocol.NewProtocol(specs)

	if err != nil {
		display_help()

		os.Exit(1)
	}

	fmt.Println(proto)
}

func parse() protocol.ProtocolDesc {
	parser := argparse.NewParser(APPLICATION_NAME, APPLICATION_DESC)

	// --bits <n>       : Number of bits per line. By default it's 32, same as
	//                  IETF. This is useful for protocols that don't align
	//                  perfectly to 32-bit words, like Ethernet.
	// --numbers <Y/N>  : Instructs protocol to print or avoid printing the bit
	//                  counts on top of the header.
	// --evenchar <c>   : Instructs protocol to use the supplied character, instead
	//                  of the default "-" as the character in even positions of
	//                  the horizontal lines.
	// --oddchar <c>    : Same as evenchar but for characters in odd positions of
	//                  the horizontal lines. By default it takes the value '+'
	// --startchar <c>  : Instructs protocol to use the supplied character instead
	//                  of the default "+" for the first position of an horizontal
	//                  line.
	// --endchar <c>    : Same as startchar but for the character in the last
	//                  position of the horizonal lines.
	// --sepchar <c>    : Instructs protocol to use the supplied character instead
	//                  of the default "|" for the field separator character.

	// --list <tcp,>

	bits := parser.Int("", "bits",
		&argparse.Options{
			Required: false,
			Help:     "Number of bits per line. By default it's 32, same as IETF. This is useful for protocols that don't align perfectly to 32-bit words, like Ethernet.",
			Default:  32,
		})
	numbers := parser.Selector(
		"",
		"numbers",
		[]string{"Y", "N"},
		&argparse.Options{
			Required: false,
			Help:     "Instructs protocol to print or avoid printing the bit counts on top of the header.",
			Default:  "Y",
		})
	evenchar := parser.String(
		"",
		"evenchar",
		&argparse.Options{
			Required: false,
			Help:     "Instructs protocol to use the supplied character, instead of the default '-' as the character in even positions of the horizontal lines.",
			Default:  "-",
		})

	oddchar := parser.String(
		"",
		"oddchar",
		&argparse.Options{
			Required: false,
			Help:     "Same as evenchar but for characters in odd positions of the horizontal lines. By default it takes the value '+'",
			Default:  "|",
		})

	startchar := parser.String(
		"",
		"startchar",
		&argparse.Options{
			Required: false,
			Help:     "Instructs protocol to use the supplied character instead of the default '+' for the first position of an horizontal line.dd positions of the horizontal Instructs protocol to use the supplied character instead of the default '+' for the first position of an horizontal line.",
			Default:  "+",
		})

	endchar := parser.String(
		"",
		"endchar",
		&argparse.Options{
			Required: false,
			Help:     "Same as startchar but for the character in the last position of the horizonal lines.",
			Default:  "+",
		})
	sepchar := parser.String(
		"",
		"sepchar",
		&argparse.Options{
			Required: false,
			Help:     "Instructs protocol to use the supplied character instead of the default '|' for the field separator character.",
			Default:  "|",
		})

	list := parser.Flag(
		"l",
		"list",
		&argparse.Options{
			Required: false,
			Default:  false,
		})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(0)
	}

	desc := protocol.ProtocolDesc{}

	if bits != nil {
	}
	if numbers != nil {
	}
	if evenchar != nil {
	}
	if oddchar != nil {
	}
	if startchar != nil {
	}
	if endchar != nil {
	}
	if sepchar != nil {
	}
	if list != nil && *list {
		displayListOfSupportedProtocols(APPLICATION_NAME)
		os.Exit(0)
	}

	return desc
}
