---
linkTitle: "Find Words by Spelling"
title: "Find Words by Spelling"
weight: 4
---
### The `--spelled-like` flag

To search for words that match a specific spelling pattern, use the
`--spelled-like` flag. The spelling can include wildcards such as (`*` and `?`).

#### Example

This example searches for words that begin with `all` and end in the letter `e`.
The results include matching words like "alleviate" and "allude".

```bash
polyhymnia --max 5 --spelled-like "all*e"
```

```text
alleviate
   Score: 2240

allude
   Score: 1968

allure
   Score: 1569

alliance
   Score: 1526

allege
   Score: 1458
```
