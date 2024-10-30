# GoKid - Trunk + branch 1!

GoKid is a CLI tool for managing changes, designed to remove the boilerplate work from creating branches, pull requests, and more.

## Installation
Requires the [GitHub CLI](https://cli.github.com/) to be installed under `gh`, e.g.

```bash
brew install gh
```

Can then be installed by:

```bash
brew update && brew install martinbernstorff/homebrew-tap/gk
```

## Usage - Trunk!

```bash
gk --help
```

## Setup
Write the default config to the current directory, and you are ready to go! 🚀

```bash
gk init
```

Gokid looks for configuration files in the current directory, or any parent directory. It uses the values from the first config it encounters.
