---
linkTitle: "Examples"
title: "Example Queries"
weight: 4
---
Here are some practical examples of how to use Polyhymnia.

## Find words with a similar meaning to "happy" and include definitions:

```bash
polyhymnia --max 3 --means-like "happy" --metadata d
```

```text
   pleased
    Score: 40004395
    Part of Speech: syn
    Part of Speech: adj
    Part of Speech: v
    Part of Speech: results_type:primary_rel
    Definition: adj happy, content

blissful
    Score: 40004156
    Part of Speech: syn
    Part of Speech: adj
    Definition: adj Extremely happy; full of joy; experiencing, indicating, causing, or characterized by bliss.
    Definition: adj (obsolete) Blessed; glorified.

content
    Score: 40004024
    Part of Speech: syn
    Part of Speech: adj
    Definition: adj Satisfied, pleased, contented.
    Definition: adj (obsolete) Contained.
```

## Find words that sound like "seven" and return a maximum of 3 results

```bash
polyhymnia --max 3 --sounds-like "seven"
```

## Search for words that are spelled like "cat\*og"

```bash
polyhymnia --spelled-like "cat*og"
```

```text
catalog
    Score: 1024

cat-and-dog
    Score: 15

cattle dog
    Score: 15

cat and dog
    Score: 12

catch dog
    Score: 1
```

## Find synonyms for "ocean":

```bash
polyhymnia --related-word syn "ocean"
```

```text
sea
    Score: 4730
```

## Find words that start with the letter "w" where the word drink appears before the word.

```bash
polyhymnia --max 5 --left-context "drink" --spelled-like "w*"
```

```text
with
    Score: 17269

water
    Score: 9976

wine
    Score: 8740

was
    Score: 7382

when
    Score: 4684
```
