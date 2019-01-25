#!/usr/bin/env bash

# Exit immediately if a command returns a non-zero status.
set -e

# Set git credentials
git config --global user.email "$GIT_EMAIL"
git config --global user.name "$GIT_USERNAME"

# let git know where to apply the changes
git checkout master

echo "Adding data files"
git add data/*

echo "Checking the number of files staged"
files=$(git diff --cached --numstat | wc -l | tr -d '[:space:]');
[[ $files -eq 0 ]] && echo "nothing to push, exiting..." && exit

echo "Commiting files"
git commit -m "Automated update repository metadata [skip-ci]"

echo "Pushing changes"
git push https://"$GIT_USER:$GITHUB_TOKEN"@github.com/veggiemonk/awesome-docker master >/dev/null 2>&1

echo "Done."
