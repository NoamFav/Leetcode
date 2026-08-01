#!/usr/bin/env python3
"""Regenerate the solutions table in README.md from the .go solution files.

Run manually with `python3 scripts/generate_readme.py`, or let the
"Update README" GitHub Action run it on every push.
"""

import re
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
README = ROOT / "README.md"

FILENAME_RE = re.compile(r"^(\d+)\.([a-z0-9-]+)\.go$")
BLOCK_COMMENT_RE = re.compile(r"/\*.*?\*/", re.DOTALL)
LINE_COMMENT_RE = re.compile(r"//[^\n]*")
WHITESPACE_RE = re.compile(r"\s+")
SIGNATURE_RE = re.compile(r"func\w*\([^)]*\)[^{]*\{")
LEET_BLOCK_RE = re.compile(r"@leet start(.*)@leet end", re.DOTALL)

ROMAN_NUMERALS = {"i", "ii", "iii", "iv", "v", "vi", "vii", "viii", "ix", "x"}

DEFAULT_TEMPLATE = """# Leetcode

Personal collection of [LeetCode](https://leetcode.com/) solutions in Go, \
managed with [leetcode.nvim](https://github.com/kawre/leetcode.nvim).

## Solutions

<!-- STATS:START -->
<!-- STATS:END -->

<!-- TABLE:START -->
<!-- TABLE:END -->
"""


def title_from_slug(slug: str) -> str:
    words = slug.split("-")
    return " ".join(w.upper() if w in ROMAN_NUMERALS else w.capitalize() for w in words)


def is_solved(go_source: str) -> bool:
    match = LEET_BLOCK_RE.search(go_source)
    body = match.group(1) if match else go_source
    body = BLOCK_COMMENT_RE.sub("", body)
    body = LINE_COMMENT_RE.sub("", body)
    body = WHITESPACE_RE.sub("", body)
    body = SIGNATURE_RE.sub("", body)
    body = body.replace("{", "").replace("}", "")
    return len(body) > 0


def collect_problems():
    problems = []
    for path in sorted(ROOT.glob("*.go")):
        match = FILENAME_RE.match(path.name)
        if not match:
            continue
        number, slug = match.groups()
        problems.append(
            {
                "number": int(number),
                "slug": slug,
                "title": title_from_slug(slug),
                "filename": path.name,
                "solved": is_solved(path.read_text()),
            }
        )
    return sorted(problems, key=lambda p: p["number"])


def render_table(problems) -> str:
    if not problems:
        return "_No solutions yet._"
    lines = [
        "| # | Problem | Solution | Status |",
        "| --- | --- | --- | --- |",
    ]
    for p in problems:
        url = f"https://leetcode.com/problems/{p['slug']}/"
        status = "✅ Solved" if p["solved"] else "🚧 In progress"
        lines.append(
            f"| {p['number']} | [{p['title']}]({url}) | [Go]({p['filename']}) | {status} |"
        )
    return "\n".join(lines)


def render_stats(problems) -> str:
    solved = sum(1 for p in problems if p["solved"])
    return f"**{solved} / {len(problems)} solved**"


def replace_block(text: str, marker: str, new_content: str) -> str:
    pattern = re.compile(
        rf"(<!-- {marker}:START -->)(.*?)(<!-- {marker}:END -->)", re.DOTALL
    )
    replacement = f"\\1\n{new_content}\n\\3"
    if not pattern.search(text):
        raise SystemExit(
            f"README.md is missing the <!-- {marker}:START/END --> markers"
        )
    return pattern.sub(replacement, text)


def main() -> None:
    problems = collect_problems()
    text = README.read_text() if README.exists() else DEFAULT_TEMPLATE
    text = replace_block(text, "STATS", render_stats(problems))
    text = replace_block(text, "TABLE", render_table(problems))
    README.write_text(text)


if __name__ == "__main__":
    main()
