#!/usr/bin/env bash

# Abort sign off on any error
set -e

# Start the benchmark timer
SECONDS=0

# Progress reporting
RED=31;
announce() { echo -e "\033[0;$2m$1\033[0m"; }

# Sign off requires a clean repository
if [[ -n $(git status --porcelain) ]]; then
    announce "Can't merge a dirty repository!" $RED
    git status
    exit 1
else
    announce "Checking signoff status" $BLUE

    OWNER=$(gh repo view --json owner --jq .owner.login)
    REPO=$(gh repo view --json name --jq .name)
    SHA=$(git rev-parse HEAD)
    USER=$(git config user.name)

    result=$(gh api \
        -H "Accept: application/vnd.github+json" \
        /repos/$OWNER/$REPO/commits/$SHA/statuses \
        | jq '[.[] | select(.context=="signoff")] | sort_by(.created_at) | last | .state == "success"')
    
    if [ "$result" != "true" ]; then
        echo "Status check failed: signoff status is not success"
        exit 1
    fi
fi