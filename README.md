# GoKid

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

## Usage

```bash
gk --help
```

## Setup
Write the default config to the current directory, and you are ready to go! 🚀

```bash
gk init
```

Gokid looks for configuration files in the current directory, or any parent directory. It uses the values from the first config it encounters.

## Roadmap
* p1: Ensure shell run returns err if exit code is not 0

* p2: Setup golint with reasonable rules

* p2: Error if un-quoted "$(some_command_here)" in config

* p3: Support creating a PR from an existing branch. Prompt for confirmation.

* p4: Support deleting the current change

* p2: Plan-execute-rollback. 
    * Print out the plan before executing
    * Print the name of each step as it's executing
    * Print the name of a command when rolling back
