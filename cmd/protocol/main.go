package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

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
			specs = lowerArg1
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Displays command-line usage help to standard output.
func display_help() {
	year, _, _ := time.Now().Date()

	fmt.Println("")
	fmt.Printf("%s v%s\n", APPLICATION_NAME, APPLICATION_VERSION)
	fmt.Printf("Copyright (C) %d, %s (%s).\n", max(2022, year), APPLICATION_AUTHOR, APPLICATION_AUTHOR_EMAIL)
	fmt.Println("This software comes with ABSOLUTELY NO WARRANTY.")
	fmt.Println("")
	display_usage()
	fmt.Println("PARAMETERS:")
	fmt.Println(" <spec>              : Field by field specification of non-existing protocol")
	fmt.Println(" <protocol>          : Name of an existing protocol")
	fmt.Println(" -l, --list          : List of supported protocols")
}

// @return a string containing application usage information
func get_usage(applicationName string) string {
	return fmt.Sprintf("Usage: %s \"{<protocol> or <spec>}\"", applicationName)
}

// Prints usage information to standard output
func display_usage() {
	fmt.Println(get_usage(APPLICATION_NAME))
}

// Displays command-line usage list of supported protocols to standard output.
func displayListOfSupportedProtocols(applicationName string) {
	rows := [][]string{}

	for proto, v := range protocol.Protocols {
		rows = append(rows, []string{v.Desc, proto})
	}
	sort.Slice(rows, func(a, b int) bool {
		// sorted by first column
		return strings.Compare(rows[a][0], rows[b][0]) < 0
	})

	fmt.Printf("Usage: %s \"<protocol>\"\n\n", applicationName)
	printTable([]string{"Protocol", "<protocol>"}, rows)
}

// from https://gosamples.dev/string-padding/
func printTable(header []string, rows [][]string) {
	table := [][]string{}
	table = append(table, header)
	table = append(table, rows...)

	// get number of columns from the first table row
	columnLengths := make([]int, len(table[0]))
	for _, line := range table {
		for i, val := range line {
			if len(val) > columnLengths[i] {
				columnLengths[i] = len(val)
			}
		}
	}

	var lineLength int
	for _, c := range columnLengths {
		lineLength += c + 3 // +3 for 3 additional characters before and after each field: "| %s "
	}
	lineLength += 1 // +1 for the last "|" in the line

	for i, line := range table {
		if i == 0 { // table header
			fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2)) // lineLength-2 because of "+" as first and last character
		}
		for j, val := range line {
			fmt.Printf("| %-*s ", columnLengths[j], val)
			if j == len(line)-1 {
				fmt.Printf("|\n")
			}
		}
		if i == 0 || i == len(table)-1 { // table header or last line
			fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2)) // lineLength-2 because of "+" as first and last character
		}
	}
}
