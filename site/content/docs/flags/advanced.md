---
linkTitle: "Advanced"
title: "Advanced Flags and Options"
weight: 2
---
These options provide more fine-tuned control over the results, especially when looking for specific metadata or relationships between words.

### Related Word Search

The `--related-word` flag allows you to explore words related to the query word based on various lexical relationships. Use the following codes to specify the type of relationship you're interested in. Examples are provided in the [Related Word Search](/docs/usage/find-related-words/) section of the [Usage Guide](/docs/usage).

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

### Vocabulary

The `--vocabulary` flag specifies which vocabulary source to use in your search.

| Vocabulary | Description                                         |  
| ---------- | --------------------------------------------------- |  
| `default`  | A 550,000-term vocabulary of English words and multiword expressions. |  
| `enwiki`   | A 6 million-term vocabulary from English Wikipedia. |  
| `es`       | A 500,000-term vocabulary of words from Spanish-language books. |

#### Wikipedia Vocabulary Example

This query searches for words with a meaning similar to "joy" using the English Wikipedia vocabulary.

```bash
polyhymnia --max 3 --means-like "joy" --vocabulary enwiki
```

```text
delight
   Score: 39999926
   Num Syllables: 2
   Pronunciation: D IH0 L AY1 T
   Frequency: 11.006901
   Part of Speech: syn
   Part of Speech: n
   Part of Speech: results_type:primary_rel
   Definition: n  Joy; pleasure.
   Definition: n  Something that gives great joy or pleasure.

rejoice
   Score: 39999920
   Num Syllables: 2
   Pronunciation: R IH0 JH OY1 S
   Frequency: 2.931221
   Part of Speech: syn
   Part of Speech: v
   Definition: v  (intransitive) To be very happy, be delighted, exult; to feel joy.
   Definition: v  (obsolete, transitive) To have (someone) as a lover or spouse; to enjoy sexually.
   Definition: v  (transitive) To make happy, exhilarate.
   Definition: v  (obsolete) To enjoy.
```

#### Spanish Vocabulary Example

This query searches for words that are spelled "ba" followed by any single character and ending with the letter "o" using the Spanish vocabulary.

```bash
polyhymnia --max 3 --spelled-like "ba?o" --vocabulary es
```

```text
bajo
   Score: 216190
   Num Syllables: 2
   Pronunciation: B AA1 XH OW0
   Frequency: 403.005755
   Part of Speech: n
   Definition: n  estribación sumergida ( o parcialmente sumergida ) en un río o a lo largo de una orilla costera

baño
   Score: 9137
   Num Syllables: 2
   Pronunciation: B AA1 N Y OW0
   Frequency: 0.242895
   Part of Speech: n
   Definition: n  trabajo de aplicar algo

bajó
   Score: 5325
   Num Syllables: 2
   Pronunciation: B AA0 XH OW1
   Frequency: 9.927057
   Part of Speech: v
   Definition: v  recortar hacer una reducción
   Definition: v  descender en forma vertical
   Definition: v  deteriorarse
   Definition: v  caer de las nubes
   Definition: v  desabrochar , como una vela , de un mástil o un soporte
```
