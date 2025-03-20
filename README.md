# GoKid

GoKid is a CLI tool for managing changes, designed to remove the boilerplate work from creating branches and pull requests.

## Quick-start
1. Install the [GitHub CLI](https://cli.github.com/)

```bash
brew install gh && gh auth login
```

2. Install `gokid`:

```bash
brew update && brew install martinbernstorff/homebrew-tap/gk
```

or by cloning the repository, and:

```bash
cd gokid && go install .
```

To create your first change (branch + PR):

```bash
gokid new "my new PR"
```

or, using aliases:

```bash
gk n "my new PR"
```

## Settings
Write the default config to the current directory:

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
