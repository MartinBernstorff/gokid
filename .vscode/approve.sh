#!/usr/bin/env bash

# Abort sign off on any error
set -euo pipefail

# Start the benchmark timer
SECONDS=0

# Progress reporting
GREEN=32; RED=31; BLUE=34
announce() { echo -e "\033[0;$2m$1\033[0m"; }

# Sign off requires a clean repository
if [[ -n $(git status --porcelain) ]]; then
  announce "Can't sign off on a dirty repository!" $RED
  git status
  exit 1
fi

run() {
  local SPLIT=$SECONDS
  announce "\nRun $1" $BLUE
  eval "$1"
  local INTERVAL=$((SECONDS-SPLIT))
  announce "Completed $1 in $INTERVAL seconds" $GREEN
}

# Run the command
run "go test ./..."

# Repository introspection
OWNER=$(gh repo view --json owner --jq .owner.login)
REPO=$(gh repo view --json name --jq .name)
SHA=$(git rev-parse HEAD)
USER=$(git config user.name)

announce "Attempting to sign off on $SHA in $OWNER/$REPO as $USER" $GREEN

# Report successful sign off to GitHub
gh api \
  --method POST --silent \
  -H "Accept: application/vnd.github+json" -H "X-GitHub-Api-Version: 2022-11-28" \
  /repos/$OWNER/$REPO/statuses/$SHA \
  -f "context=signoff" -f "state=success" -f "description=Signed off by $USER ($SECONDS seconds)"

announce "Signed off on $SHA in $SECONDS seconds" $GREEN