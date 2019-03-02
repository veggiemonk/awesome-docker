# workflow "New workflow" {
#   on = "push"
#   resolves = ["push changes"]
# }

# action "skip-commit" {
#   uses = "veggiemonk/skip-commit@449e94fa83e7918c4079f37322205e17b868f993"
#   env = {
#     COMMIT_FILTER = "skip-ci"
#   }
# }

# action "npm install" {
#   uses = "actions/npm@de7a3705a9510ee12702e124482fad6af249991b"
#   needs = ["skip-commit"]
#   args = "install"
# }

# action "npm run build" {
#   uses = "actions/npm@de7a3705a9510ee12702e124482fad6af249991b"
#   needs = ["npm install"]
#   args = "run build"
# }

# action "Build metadata" {
#   needs = ["npm run build"]
#   uses = "actions/npm@master"
#   runs = "sh -l -c"
#   args = ["node buildMetadata.js"]
#   secrets = ["GITHUB_TOKEN"]
#   env = {
#     GIT_EMAIL = "alex.blaine@layder.io"
#     GIT_USERNAME = "veggiemonk"
#   }
# }

# action "push changes" {
#   uses = "veggiemonk/bin/git@master"
#   needs = ["Build metadata"]
#   runs = "sh -c $@"
#   args = "push.sh"
#   secrets = ["GITHUB_TOKEN"]
#   env = {
#     GIT_EMAIL = "alex.blaine@layder.io"
#     GIT_USERNAME = "veggiemonk-bot"
#     GIT_USER = "veggiemonk"
#   }
# }
