# yup-xargs

```
NAME:
   xargs - build and execute command lines from standard input

USAGE:
   xargs [OPTIONS] [COMMAND [INITIAL-ARGS]]

      Execute COMMAND with arguments from standard input.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --max-args value, -n value           use at most MAX-ARGS arguments per command line (default: 0)
   --max-lines value, -L value          use at most MAX-LINES non-blank input lines per command line (default: 0)
   --max-chars value, -s value          limit length of command line to MAX-CHARS (default: 0)
   --max-procs value, -P value          run up to MAX-PROCS processes at a time (default: 1)
   --delimiter value, -d value          items in input stream are separated by CHARACTER
   --replace value, -I value, -i value  replace REPLACE-STR in INITIAL-ARGS with names read from standard input
   --null, -0                           items are separated by a null, not whitespace (default: false)
   --print, -p                          print commands before running them (default: false)
   --interactive, -t                    be verbose; print commands before executing (default: false)
   --no-run-if-empty, -r                do not run COMMAND if standard input is empty (default: false)
   --verbose                            print commands before executing them (default: false)
   --help, -h                           show help
```
