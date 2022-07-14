package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// TODO: Create our own logging wrapper for zerolog
//       https://github.com/rs/zerolog

var (
	debug bool = true // FIXME: Set to false and load from .env instead, to make development easier?
)

func main() {
	// TODO: Use a different library for dealing with arguments, with full help and all that
	// TODO: Use a separate argument for pretty console output, something that makes more sense?
	flag.BoolVar(&debug, "debug", false, "sets log level to debug")
	flag.Parse()

	if debug {
		// TODO: Only use this when necessary, as this directly impacts performance!
		// Use the console writer for pretty output
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		// TODO: Only use this in production? Since unix time isn't human readable
		// UNIX Time is faster and smaller than most timestamps
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}

	// Test logging to the console
	log.Warn().
		Bool("debug", debug).
		Str("user", "N/A").
		Msg("WELCOME TO B(.)(.)BS - BBS.. bbs?")

	log.Debug().
		Bool("debug", debug).
		Str("scale", "833 cents").
		Float64("interval", 833.09).
		Msg("I see you.")

	log.Info().
		Bool("debug", debug).
		Str("firstName", "Tom").
		Str("lastName", "Hanks").
		Send()
}
