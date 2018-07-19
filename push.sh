#!/usr/bin/env bash

# Exit immediately if a command returns a non-zero status.
set -e

# Set git credentials
git config --global user.email "info@veggiemonk.ovh"
git config --global user.name "veggiemonk-bot"

# let git know where to apply the changes
git checkout master

echo 'Adding data files'
git add data/*

echo 'Commiting files'
git commit -m 'Automated update repository metadata'

echo 'Pushing changes'
git push https://$GITHUB_USER:$GITHUB_TOKEN@github.com/veggiemonk/awesome-docker master >/dev/null 2>&1

echo 'Done.'