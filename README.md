# GoKid

GoKid is a CLI tool for managing changes, designed to remove the boilerplate work from creating branches, pull requests, and more.

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

## Usage

To create your first change (branch + PR):

```bash
gokid new "feat: my new feature"
```

or, using aliases:

```bash
gk n "feat: my new feature"
```

To get an overview of commands:

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
* p4: Go through the code and see if we use pointers where it makes sense
* p2: Go through Uber's Error style-guide and follow it
* p3: Go through the tdd-course and apply lessons?
* p3: Rename to warpgate?