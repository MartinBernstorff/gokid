# GoKid

GoKid is a CLI tool for unifying the management of branches, pull requests etc. as one 'change'.

## Installation
Requires the github CLI to be installed under `gh`.

```bash
brew update && brew install martinbernstorff/homebrew-tap/gk
```

## Usage

```bash
gk --help
```

Available Commands:
- `new`: Create a new change
- `ready`: Mark a change as ready for review

Use `gk [command] --help` for more information about a command.

## Commands

### New Change

Create a new change with the following command:

```bash
gk new
```

This command will:
1. Prompt you for a change title (must contain a colon)
2. Create a new branch
3. Create a pull request with the given title

### Mark Ready

Mark a change as ready for review:

```bash
gk ready
```

This command will:
1. Mark the current pull request for auto-merge
2. Set the merge method to squash
