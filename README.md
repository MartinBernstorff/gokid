# GoKid

GoKid is a CLI tool for managing changes, designed to remove the boilerplate work from creating branches, pull requests, and more.

## Installation
Requires the github CLI to be installed under `gh`.

```bash
brew update && brew install martinbernstorff/homebrew-tap/gk
```

## Usage

```bash
gk --help
```

## Configuration
Write a test config to the current directory:

```bash
gk config --write
```

Gokid looks for configuration files in the current directory, or any parent directory.
