---
linkTitle: "Find Related Words"
title: "Related Word Search"
weight: 5
---
### The `--related-word` flag

To find words based on specific relationships like synonyms, antonyms, or homophones, you can use the `--related-word` flag and one of the related word search codes.

### Related Word Search Codes

The related word search codes allow you to specify the type of relationship you're interested in.

| Code  | Description                              | Example            |  
| :---: | ---------------------------------------- | ------------------ |  
| `syn` | Synonyms                                 | ocean → sea        |  
| `ant` | Antonyms                                 | late → early       |  
| `hom` | Homophones                               | course → coarse    |  
| `trg` | Trigger words (commonly associated)      | cow → milk         |  
| `jjb` | Adjectives used to modify the given noun | beach → sandy      |  
| `jja` | Nouns modified by the given adjective    | gradual → increase |  
| `spc` | Hypernyms ("kind of")                    | gondola → boat     |  
| `par` | Meronyms ("part of")                     | trunk → tree       |  

### Examples

#### Find antonyms for the word "joy"

This query returns **antonyms** of the word **"joy"** by using the [related-word
code](#related-word-search-codes) **`ant`**. We limit the number of
returned words to 3 using the **`--max`** [flag](/docs/flags/#optional-flags).

```bash
polyhymnia --max 3 --related-word ant joy
```

```text
sad

sorrow

sadden
```

#### Find homophones for the word "joy"

This query returns **homophones** of the word **"joy"** by using the [related-word
code](#related-word-search-codes) **`hom`**.

```bash
polyhymnia --related-word hom joy
```

```text
joye
```

#### Find synonyms for the word "joy"

This query returns **synonyms** of the word **"joy"** by using the [related-word
code](#related-word-search-codes) **`syn`**. We limit the number of
returned words to 3 using the **`--max`** [flag](/docs/flags/#optional-flags).

```bash
polyhymnia --max 3 --related-word syn joy
```

```text
delight

pleasure

rejoice
```
