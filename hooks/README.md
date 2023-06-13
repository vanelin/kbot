# How to stops leaks

## Source:
- [Gitleaks](https://github.com/gitleaks/gitleaks)

## Requirements

- [Install the `curl`](https://everything.curl.dev/get)

- Install the `wget`

- [Using a central config Gitleaks](https://raw.githubusercontent.com/gitleaks/gitleaks/master/config/gitleaks.toml)

- Tested on `maos` *arm64/x64* and `linux` *arm64/x64*

## Installation notes

Clone the repository and go to `hooks` folder.

`make install` will install `gitleaks`. The install will:

- install `gitleaks`
- add a global `pre-commit` hook to `$HOME/.git-support/hooks/pre-commit`
- add the configuration with central config patterns to `$HOME/.git-support/gitleaks.toml`

You now have the gitleaks pre-commit hook enabled globally.

## Deletion notes
`make clean` will remove `gitleaks` and clean `git config`.

## Usage notes
- `make detect` - detect secrets in code
- `make version` - displaying the installed version of gitleaks

[![gitleaks](https://asciinema.org/a/591296.svg)](https://asciinema.org/a/591296)
