# GoKid

GoKid is a CLI tool for managing changes, designed to remove the boilerplate work from creating branches, pull requests, and more.

## Installation

```bash
brew tap martinbernstorff/homebrew-tap
brew update && brew install martinbernstorff/homebrew-tap/gokid
```

## Usage

```bash
gokid --help
```

Available Commands:
- `new`: Create a new change
- `ready`: Mark a change as ready for review

Use `gokid [command] --help` for more information about a command.

## Commands

### New Change

Create a new change with the following command:

```bash
gokid new
```

This command will:
1. Prompt you for a change title (must contain a colon)
2. Create a new branch
3. Create a pull request with the given title

### Mark Ready

Mark a change as ready for review:

```bash
gokid ready
```

This command will:
1. Mark the current pull request for auto-merge
2. Set the merge method to squash