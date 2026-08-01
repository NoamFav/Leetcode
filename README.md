# Leetcode

Personal collection of [LeetCode](https://leetcode.com/) solutions in Go, managed with
[leetcode.nvim](https://github.com/kawre/leetcode.nvim).

![Update README](https://github.com/NoamFav/Leetcode/actions/workflows/update-readme.yml/badge.svg)

## Solutions

The table below is generated automatically — see [Automation](#automation).

<!-- STATS:START -->
**9 / 10 solved**
<!-- STATS:END -->

<!-- TABLE:START -->
| # | Problem | Solution | Status |
| --- | --- | --- | --- |
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | [Go](1.two-sum.go) | ✅ Solved |
| 2 | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | [Go](2.add-two-numbers.go) | ✅ Solved |
| 4 | [Median Of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | [Go](4.median-of-two-sorted-arrays.go) | ✅ Solved |
| 8 | [String To Integer Atoi](https://leetcode.com/problems/string-to-integer-atoi/) | [Go](8.string-to-integer-atoi.go) | ✅ Solved |
| 9 | [Palindrome Number](https://leetcode.com/problems/palindrome-number/) | [Go](9.palindrome-number.go) | ✅ Solved |
| 12 | [Integer To Roman](https://leetcode.com/problems/integer-to-roman/) | [Go](12.integer-to-roman.go) | ✅ Solved |
| 13 | [Roman To Integer](https://leetcode.com/problems/roman-to-integer/) | [Go](13.roman-to-integer.go) | ✅ Solved |
| 234 | [Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/) | [Go](234.palindrome-linked-list.go) | ✅ Solved |
| 1034 | [Coloring A Border](https://leetcode.com/problems/coloring-a-border/) | [Go](1034.coloring-a-border.go) | ✅ Solved |
| 3518 | [Smallest Palindromic Rearrangement II](https://leetcode.com/problems/smallest-palindromic-rearrangement-ii/) | [Go](3518.smallest-palindromic-rearrangement-ii.go) | 🚧 In progress |
<!-- TABLE:END -->

## Automation

Every solution lives at the repo root as `<problem-id>.<slug>.go`, the naming scheme
leetcode.nvim uses when you open or submit a problem. There's nothing to maintain by hand:

- [`scripts/generate_readme.py`](scripts/generate_readme.py) scans the repo for solution
  files, figures out each problem's title and LeetCode URL from its filename, and checks
  whether the solution body is still an empty stub (🚧 In progress) or has real code
  (✅ Solved).
- The [Update README workflow](.github/workflows/update-readme.yml) runs that script on
  every push to `main` that touches a `.go` file, and commits the refreshed README back
  to the repo — so this file is always in sync with what's actually solved.

To regenerate it locally:

```sh
python3 scripts/generate_readme.py
```
