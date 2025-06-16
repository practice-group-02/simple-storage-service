package config

import (
	"flag"
	"fmt"
	"strings"
)

var (
	Dir  string
	Port int
)

func ParseFlags() error {
	flag.StringVar(&Dir, "dir", "s3", "Path to directory")
	flag.IntVar(&Port, "port", 8080, "Port to serve on")
	flag.Usage = func() {
		printHelp()
	}
	flag.Parse()

	if Port < 1024 || Port > 65565 {
		return fmt.Errorf("Port must be between 1024 and 65565")
	}

	standardDirs := []string{"triple-s", "cmd", "config", "internal", "pkg"}
	for _, dir := range standardDirs {
		if strings.Contains(Dir, dir) {
			return fmt.Errorf("directory %s contains substring of standard package <%s>", Dir, dir)
		}
	}

	if Dir == ".." || Dir == "../" || Dir == "." {
		return fmt.Errorf("prohibited directory, please enter another one (s3)")
	}
	return nil
}

func printHelp() {
	fmt.Println(`Simple Storage Service.

**Usage:**
    triple-s [-port <N>] [-dir <S>]  
    triple-s --help

**Options:**
- --help     Show this screen.
- --port N   Port number
- --dir S    Path to the directory
`)
}
