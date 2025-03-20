# GoKid

GoKid is a CLI tool for managing changes, designed to remove the boilerplate work from creating branches and pull requests.

## Installation
Requires the [GitHub CLI](https://cli.github.com/) to be installed under `gh`, e.g.

```bash
brew install gh && gh auth login
```

Can then be installed by:

```bash
brew update && brew install martinbernstorff/homebrew-tap/gk
```

or by cloning the repository, and:

```bash
cd gokid && go install .
```

## Quick-start

To create your first change (branch + PR):

```bash
gokid new "my new PR"
```

or, using aliases:

```bash
gk n "my new PR"
```

## Usage
Write the default config to the current directory, and you are ready to go! 🚀

```bash
gk init
```

Gokid looks for configuration files (by priority): 
1. In the current directory
2. Any parent directory
3. `~/.config/gokid/.gokid` 

It uses the values from the first config it encounters.

To get an overview of commands:

```bash
gk --help
```

## Roadmap
* p3: Go through the tdd-course and apply lessons?
* p3: Rename to warpgate?