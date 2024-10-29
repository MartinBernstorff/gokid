#!/usr/bin/env bash

# Abort sign off on any error
set -e

# Start the benchmark timer
SECONDS=0

# Repository introspection
OWNER=$(gh repo view --json owner --jq .owner.login)
REPO=$(gh repo view --json name --jq .name)
SHA=$(git rev-parse HEAD)
USER=$(git config user.name)

GREEN=32; RED=31;
announce() { echo -e "\033[0;$2m$1\033[0m"; }
if [[ -n $(git status --porcelain) ]]; then
  announce "Can't sign off on a dirty repository!" $RED
  git status
  exit 1
else
  announce "Attempting to sign off on $SHA in $OWNER/$REPO as $USER" $GREEN
  go test ./...
fi

# Report successful sign off to GitHub
gh api \
  --method POST --silent \
  -H "Accept: application/vnd.github+json" -H "X-GitHub-Api-Version: 2022-11-28" \
  /repos/$OWNER/$REPO/statuses/$SHA \
  -f "context=signoff" -f "state=success" -f "description=Signed off by $USER ($SECONDS seconds)"

announce "Signed off on $SHA in $SECONDS seconds" $GREEN