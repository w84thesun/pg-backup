package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"pg-backup/internal/cli"
)

const (
	dumpUsage = `dump	Dump database`
)

func printUsageAndExit() {
	flag.Usage()

	// If a command is not found we exit with a status 2 to match the behavior
	// of flag.Parse() with flag.ExitOnError when parsing an invalid flag.
	os.Exit(2)
}

func main() {
	helpPtr := flag.Bool("help", false, "")
	databasePtr := flag.String("database", "", "")
	outputPtr := flag.String("output", "", "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			`Usage: backup OPTIONS COMMAND [arg...]
       backup [ -version | -help ]
Options:
  -database          Source for dump (postgres://url)
  -output            Location of dump file
  -help            Print usage
Commands:
  %s
`, dumpUsage)
	}

	flag.Parse()

	// show help
	if *helpPtr {
		flag.Usage()
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		printUsageAndExit()
	}

	switch flag.Arg(0) {
	case "dump":
		if databasePtr == nil || outputPtr == nil {
			printUsageAndExit()
		}

		dir, err := os.Stat(*outputPtr)
		if err != nil {
			log.Fatalf("failed to open directory, error: %w", err)
		}
		if !dir.IsDir() {
			log.Fatalf("%q is not a directory", dir.Name())
		}

		if err := cli.DumpCmd(*databasePtr, *outputPtr); err != nil {
			log.Fatal(err)
		}
	}
}
