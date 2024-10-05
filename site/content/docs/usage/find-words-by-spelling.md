---
linkTitle: "Find Words by Spelling"
title: "Find Words by Spelling"
weight: 4
---
### The `--spelled-like` flag

To search for words that match a specific spelling pattern, use the
`--spelled-like` flag. The spelling can include wildcards such as (`*` and `?`).

#### Example

This example searches for words that begin with `all` and end in the letter `e`. It limits the results to 5 words.
The results include matching words like "alleviate" and "allude".

```bash
polyhymnia --max 5 --spelled-like "all*e"
```

```text
alleviate

allude

allure

alliance

allege
```
