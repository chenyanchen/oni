# oni

Oxygen Not Included tools.

# Pre-requirements

- Oxygen Not Included installed

# How to use oni?

```bash
❯ ./bin/oni -h
Oxygen Not Included tools

Usage:
  oni [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  element     element
  help        Help about any command

Flags:
  -h, --help   help for oni

Use "oni [command] --help" for more information about a command.
```

## element

```bash
❯ ./bin/oni element -h
element

Usage:
  oni element [flags]

Flags:
  -f, --format string   optional format: json, yaml, csv (default "json")
  -h, --help            help for element
  -o, --output string   output file
```

e.g.

```bash
❯ ./bin/oni element -f csv -o elements.csv
```
