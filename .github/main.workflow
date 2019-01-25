workflow "New workflow" {
  on = "push"
  resolves = ["push changes"]
}

action "skip-commit" {
  uses = "veggiemonk/skip-commit@master"
  env = {
    COMMIT_FILTER = "skip-ci"
  }
}

action "npm install" {
  uses = "actions/npm@de7a3705a9510ee12702e124482fad6af249991b"
  needs = ["skip-commit"]
  args = "install"
}

action "npm run build" {
  uses = "actions/npm@de7a3705a9510ee12702e124482fad6af249991b"
  needs = ["npm install"]
  args = "run build"
}

action "Build metadata" {
  uses = "veggiemonk/bin/git@master"
  runs = "sh -c \"$@\""
  args = "node buildMetadata.js"
  needs = ["npm run build"]
  secrets = ["GITHUB_TOKEN"]
  env = {
    GIT_EMAIL = "alex.blaine@layder.io"
    GIT_USERNAME = "veggiemonk"
  }
}

action "push changes" {
  uses = "actions/bin/shell@master"
  needs = ["Build metadata"]
  runs = "sh -c $@"
  args = "push.sh"
  secrets = ["GITHUB_TOKEN"]
  env = {
    GIT_EMAIL = "alex.blaine@layder.io"
    GIT_USERNAME = "veggiemonk-bot"
    GIT_USER = "veggiemonk"
  }
}
