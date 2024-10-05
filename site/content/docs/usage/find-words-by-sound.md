---
linkTitle: "Find Words by Sound"
title: "Find Words by Sound"
weight: 3
---
### The `--sounds-like` flag

To search for words that sound like a specific word, use the `--sounds-like` flag.

#### Example

This example finds words that sound like "kevin" and limits the results to 5 words.

```bash
polyhymnia --max 5 --sounds-like "kevin"
```

```text
kevin
   Score: 100
   Num Syllables: 2

kevyn
   Score: 100
   Num Syllables: 2

kevan
   Score: 97
   Num Syllables: 2

cave in
   Score: 95
   Num Syllables: 2

cave-in
   Score: 95
   Num Syllables: 2

```
