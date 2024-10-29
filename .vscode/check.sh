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

# Function to check status
check_status() {
    local result=$(gh api \
  -H "Accept: application/vnd.github+json" \
  /repos/$OWNER/$REPO/commits/$SHA/statuses \
  | jq '[.[] | select(.context=="signoff")] | sort_by(.created_at) | last | .state == "success"')
    
    if [ "$result" != "true" ]; then
        echo "Status check failed: signoff status is not success"
        exit 1
    fi
}

# Run the check
check_status

