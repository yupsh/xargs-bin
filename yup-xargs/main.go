package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/xargs"
)

const (
	flagMaxArgs     = "max-args"
	flagMaxLines    = "max-lines"
	flagMaxChars    = "max-chars"
	flagMaxProcs    = "max-procs"
	flagDelimiter   = "delimiter"
	flagReplaceStr  = "replace"
	flagNullDelim   = "null"
	flagPrint       = "print"
	flagInteractive = "interactive"
	flagNoRunEmpty  = "no-run-if-empty"
	flagVerbose     = "verbose"
)

func main() {
	app := &cli.App{
		Name:  "xargs",
		Usage: "build and execute command lines from standard input",
		UsageText: `xargs [OPTIONS] [COMMAND [INITIAL-ARGS]]

   Execute COMMAND with arguments from standard input.`,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagMaxArgs,
				Aliases: []string{"n"},
				Usage:   "use at most MAX-ARGS arguments per command line",
			},
			&cli.IntFlag{
				Name:    flagMaxLines,
				Aliases: []string{"L"},
				Usage:   "use at most MAX-LINES non-blank input lines per command line",
			},
			&cli.IntFlag{
				Name:    flagMaxChars,
				Aliases: []string{"s"},
				Usage:   "limit length of command line to MAX-CHARS",
			},
			&cli.IntFlag{
				Name:    flagMaxProcs,
				Aliases: []string{"P"},
				Usage:   "run up to MAX-PROCS processes at a time",
				Value:   1,
			},
			&cli.StringFlag{
				Name:    flagDelimiter,
				Aliases: []string{"d"},
				Usage:   "items in input stream are separated by CHARACTER",
			},
			&cli.StringFlag{
				Name:    flagReplaceStr,
				Aliases: []string{"I", "i"},
				Usage:   "replace REPLACE-STR in INITIAL-ARGS with names read from standard input",
			},
			&cli.BoolFlag{
				Name:    flagNullDelim,
				Aliases: []string{"0"},
				Usage:   "items are separated by a null, not whitespace",
			},
			&cli.BoolFlag{
				Name:    flagPrint,
				Aliases: []string{"p"},
				Usage:   "print commands before running them",
			},
			&cli.BoolFlag{
				Name:    flagInteractive,
				Aliases: []string{"t"},
				Usage:   "be verbose; print commands before executing",
			},
			&cli.BoolFlag{
				Name:    flagNoRunEmpty,
				Aliases: []string{"r"},
				Usage:   "do not run COMMAND if standard input is empty",
			},
			&cli.BoolFlag{
				Name:  flagVerbose,
				Usage: "print commands before executing them",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "xargs: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments as command/args
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.IsSet(flagMaxArgs) {
		params = append(params, MaxArgs(c.Int(flagMaxArgs)))
	}
	if c.IsSet(flagMaxLines) {
		params = append(params, MaxLines(c.Int(flagMaxLines)))
	}
	if c.IsSet(flagMaxChars) {
		params = append(params, MaxChars(c.Int(flagMaxChars)))
	}
	if c.IsSet(flagMaxProcs) {
		params = append(params, MaxProcs(c.Int(flagMaxProcs)))
	}
	if c.IsSet(flagDelimiter) {
		params = append(params, Delimiter(c.String(flagDelimiter)))
	}
	if c.IsSet(flagReplaceStr) {
		params = append(params, ReplaceStr(c.String(flagReplaceStr)))
	}
	if c.Bool(flagNullDelim) {
		params = append(params, NullDelim)
	}
	if c.Bool(flagPrint) {
		params = append(params, Print)
	}
	if c.Bool(flagInteractive) {
		params = append(params, Interactive)
	}
	if c.Bool(flagNoRunEmpty) {
		params = append(params, NoRunEmpty)
	}
	if c.Bool(flagVerbose) {
		params = append(params, Verbose)
	}

	// Create and execute the xargs command
	cmd := Xargs(params...)
	return yup.Run(cmd)
}
