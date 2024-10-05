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

This query returns antonyms of the word "joy" by using the related-word code `ant`. We limit the number of returned words to 3 using the `--max` flag.

```bash
polyhymnia --max 3 --related-word ant joy
```

```text
sad
   Score: 3401

sorrow
   Score: 1141

sadden
   Score: 311
```

#### Find homophones for the word "joy"

```bash
polyhymnia --max 3 --related-word hom joy
```

```text
joye
    Score: 44
    Num Syllables: 1
```

#### Find synonyms for the word "joy"

```bash
polyhymnia --max 3 --related-word syn joy
```

```text
delight
    Score: 1746

pleasure
    Score: 1601

rejoice
    Score: 925
```
