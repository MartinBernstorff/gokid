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
gokid new "feat: my new feature"
```

or, using aliases:

```bash
gk n "feat: my new feature"
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
* p4: Go through the code and see if we use pointers where it makes sense
* p1: Go through Uber's Error style-guide and follow it
* p3: Go through the tdd-course and apply lessons?
* p2: Make the output more quiet
* p2: Output the URL to the PR on creation
* p3: Rename to warpgate?