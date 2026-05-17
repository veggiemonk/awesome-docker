#!/usr/bin/env python3
"""Normalize :yen: placement in README.md entries.

Transformations (only on list entries — leaves legend lines and section prose alone):
  - [Title :yen:](url) - desc   ->  - [Title](url) - :yen: desc
  - [Title](url) :yen:- desc    ->  - [Title](url) - :yen: desc
  - [Title](url) :yen: desc     ->  - [Title](url) - :yen: desc
  - [Title](url) :yen:          ->  - [Title](url) - :yen:

Idempotent: already-correct lines pass through untouched.
"""
from __future__ import annotations

import re
import sys
from pathlib import Path

# Entry where :yen: is inside the link title.
RE_IN_TITLE = re.compile(r'^(\s*[-*]\s+\[)([^\]]*?)\s+:yen:(\s*\]\([^)]+\))(\s*-?\s*)(.*)$')

# Entry where :yen: sits right after the URL, followed by "- desc", " desc", or nothing.
RE_AFTER_URL = re.compile(r'^(\s*[-*]\s+\[[^\]]+\]\([^)]+\))\s+:yen:\s*(-\s*)?(.*)$')

# Already-correct entries: "(url) - :yen: ..."
RE_ALREADY_OK = re.compile(r'\]\([^)]+\)\s+-\s+:yen:(\s|$)')


def transform(line: str) -> str:
    if ':yen:' not in line:
        return line
    if RE_ALREADY_OK.search(line):
        return line

    m = RE_IN_TITLE.match(line)
    if m:
        prefix, title, link_close, _sep, desc = m.groups()
        desc = desc.lstrip()
        if desc:
            return f"{prefix}{title.rstrip()}{link_close} - :yen: {desc}"
        return f"{prefix}{title.rstrip()}{link_close} - :yen:"

    m = RE_AFTER_URL.match(line)
    if m:
        head, _dash, desc = m.groups()
        desc = desc.lstrip()
        if desc:
            return f"{head} - :yen: {desc}"
        return f"{head} - :yen:"

    return line


def main(path: str) -> int:
    p = Path(path)
    original = p.read_text()
    out_lines = []
    changed = 0
    for line in original.splitlines():
        new = transform(line)
        if new != line:
            changed += 1
        out_lines.append(new)
    # Preserve trailing newline.
    new_text = "\n".join(out_lines)
    if original.endswith("\n"):
        new_text += "\n"
    if new_text != original:
        p.write_text(new_text)
    print(f"{changed} line(s) changed in {path}")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv[1] if len(sys.argv) > 1 else "README.md"))
