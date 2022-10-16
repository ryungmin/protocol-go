package main

import (
	"fmt"
	"os"
	"path/filepath"
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
		proto, err := protocol.NewProtocol(os.Args[1])

		if err != nil {
			os.Exit(1)
		}

		fmt.Println(proto)
	} else {
		display_help()
	}
	// proto, err := protocol.NewProtocol(protocol.Udp)

	// spec := "PClient version:4, System ID:8, User ID:50, User Password:50, NIC:1, t_NICInfo 0:12, t_NICInfo 1:12, t_NICInfo 2:12, t_NICInfo 3:12, t_NICInfo 4:12, Computer name:255, Workgroup name:235, OTP code:20, Local IP:4, Login type:1"
	// spec := "PClient version:4"
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
}

// @return a string containing application usage information
func get_usage(applicationName string) string {
	return fmt.Sprintf("Usage: %s \"{<spec>}\"", applicationName)
}

// Prints usage information to standard output
func display_usage() {
	fmt.Println(get_usage(APPLICATION_NAME))
}
