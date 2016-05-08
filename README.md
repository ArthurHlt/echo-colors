# echo-colors

Echo in CLI with colors easily by using [colorstring](https://github.com/mitchellh/colorstring).

## Installation

### On *nix system

You can install this via the command-line with either `curl` or `wget`.

#### via curl

```bash
$ sh -c "$(curl -fsSL https://raw.github.com/ArthurHlt/echo/master/bin/install.sh)"
```

#### via wget

```bash
$ sh -c "$(wget https://raw.github.com/ArthurHlt/echo-colors/master/bin/install.sh -O -)"
```

### On windows

You can install it by downloading the `.exe` corresponding to your cpu from releases page: https://github.com/ArthurHlt/echo-colors/releases .
Alternatively, if you have terminal interpreting shell you can also use command line script above, it will download file in your current working dir.

### From go command line

Simply run in terminal:

```bash
$ go get github.com/ArthurHlt/echo-colors
```

## Usage

```
NAME:
   echoc - Echo in colors some text easily

USAGE:
   echoc "[red] this is red [yellow]and yellow [_red_]with red background [bold] and bold [reset] and no colors"

VERSION:
   1.0.0

COMMANDS:
GLOBAL OPTIONS:
   -n			Optional. Do not print the trailing newline character
   --help, -h		show help
   --version, -v	print the version
```