---
name: triage-prs
description: Triage the open pull requests of this curated list — validate each one with the repo CLI, sort it into merge / changes-requested / hold / reject, then merge, review, label, or close it with a friendly comment. Use when asked to triage, sweep, or clean up the open PRs.
---

# Triage the PR queue

Almost every PR here adds one line to `README.md`. Triage is therefore two questions per PR — *is the entry mechanically correct?* (the CLI answers it) and *does the project belong on the list?* (you answer it from GitHub signals) — followed by exactly one outward action.

Ground rule for the second question, from `CLAUDE.md`: **the project has to be for Docker, not just using Docker.**

## 1. Collect the queue

```bash
gh pr list --limit 50 --json number,title,author,createdAt,mergeable \
  --jq '.[] | "\(.number)\t\(.author.login)\t\(.createdAt[0:10])\t\(.mergeable)\t\(.title)"'
gh pr diff <N> | grep -E '^\+\+\+|^\+[^+]'   # the proposed entry
```

Read the added line, not the title. Dependabot PRs skip to step 4.

## 2. Machine check — merge onto master, then lint

Never lint the PR branch alone: it was written against a stale base, so `gh pr diff` hunk headers name the wrong section, and ordering errors appear only after newer entries have landed. Merge each PR onto current master in a scratch worktree instead.

```bash
WT=<scratchpad>/wt
git worktree add -f "$WT" master && cp awesome-docker "$WT/"
for n in <PRs>; do git fetch -q origin "pull/$n/head:pr-$n" -f; done
for n in <PRs>; do
  git -C "$WT" reset -q --hard origin/master; git -C "$WT" clean -qfd
  if git -C "$WT" merge -q --no-edit "pr-$n" >/dev/null 2>&1; then
    echo "== PR $n $( (cd "$WT" && ./awesome-docker lint 2>&1 | head -3) )"
  else
    git -C "$WT" merge --abort; echo "== PR $n CONFLICT"
  fi
done
git worktree remove --force "$WT"; for n in <PRs>; do git branch -qD pr-$n; done
```

From the merged file, read the real target section (nearest `##`/`###` heading above the added line) and check the whole diff for changes beyond the one entry. Two problems hide there: author attribution added to a neighbour's description, and an entry that a past commit removed on purpose — `git log --oneline -S <name> -- README.md` exposes the second.

## 3. Judge the project

```bash
gh api repos/<owner>/<repo> --jq '[.stargazers_count,.created_at[0:10],.pushed_at[0:10],(.license.spdx_id//"none"),(.archived|tostring)]|join(" ")'
gh api repos/<owner>/<repo>/contributors --jq length
gh api repos/<owner>/<repo>/releases --jq length
gh api "repos/<owner>/<repo>/commits?per_page=1" -i | grep -i '^link:'   # last page number = commit count
```

Stars alone decide nothing — commit count, contributors, and releases separate a real young project from a weekend demo. Check non-GitHub links with `curl -s -o /dev/null -w '%{http_code}' -L`.

Sort into exactly one bucket:

| Bucket | Signals |
|---|---|
| **Merge** | for Docker, lint clean, maintained, some traction; `:yen:` present if and only if the service is commercial |
| **Changes requested** | good entry, mechanical fault: order error, stray edits, conflict with master |
| **Hold** | for Docker and real work, but days old, no users, or a large open-issue backlog |
| **Reject** | uses Docker rather than serving it; demonstration or template repository; a handful of commits with no license |

## 4. Act — one action per PR

**Merge.** Branch protection wants 1 approval and an up-to-date branch (`strict: true`), so `gh pr merge` alone fails with *"the base branch policy prohibits the merge"*.

```bash
gh pr comment <N> --body "Thank you for your contribution! <one line on why it fits>"
gh pr review <N> --approve --body "Correct section, alphabetical order kept, lint passes."
gh pr update-branch <N>          # when mergeStateStatus is not CLEAN
gh pr merge <N> --squash --delete-branch
```

Merge one at a time and confirm with `gh pr view <N> --json state`; each merge can push the next PR out of date. A PR that touches `.github/workflows/` fails on a token without `workflow` scope, for both update and merge — leave it approved and tell the human to merge it or to run `gh auth refresh -h github.com -s workflow`.

**Changes requested.** `gh pr review <N> --request-changes --body "..."` — quote the exact linter error and name the fix (`make lint-fix` for ordering, rebase for a conflict).

**Hold.** Comment with the reason and keep it open.

**Reject.** `gh pr close <N> --comment "..."` — name the rule it misses.

Then label every PR that stays open, so its state is visible at a glance:

| Label | For |
|---|---|
| `pending-evaluation` | hold |
| `pending-fix` | small fix requested |
| `pending-rebase` | conflicts with master |
| `pending-submitter-response` | waiting on the author for anything else |

## 5. Tone

Every comment, on a merge or a rejection, is written to a volunteer who did unpaid work:

- Thank them first, and name the thing you liked.
- Give the concrete reason, quoting the rule or the linter error.
- Close by inviting correction — "if we have judged this wrongly, reply and we will look again" — and encourage a future contribution.
- Reject the *fit*, never the project's quality; point rejected projects at a list where they belong.

## 6. Verify and report

```bash
gh pr list --limit 50 --json number,labels,title --jq '.[] | "\(.number)\t[\([.labels[].name]|join(", "))]\t\(.title)"'
git pull --ff-only && ./awesome-docker lint
```

Done when every open PR is merged, closed, or open with a state label; master lints clean; and the report names each PR's bucket plus anything that needs the human.
