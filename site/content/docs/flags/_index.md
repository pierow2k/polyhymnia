---
linkTitle: "Flags and Options"
title: "Flags and Options"
next: advanced
weight: 3
---
Polyhymnia provides several flags to customize and fine-tune your queries. These flags fall into three categories, [required flags](#required-flags), [optional flags](#optional-flags) and [metadata flags](#metadata-flags).

## Required Flags

At least one of the following flags **must be included** in every Polyhymnia query. These flags define the type of word search to perform:

| Flag             | Description                        | Example |
| ---------------- | ---------------------------------- | ------- |
| `--means-like`   | Find words with a similar meaning. | `--means-like "joy"` |
| `--sounds-like`  | Find words that sound similar to the specified word. | `--sounds-like "kevin"` |
| `--spelled-like` | Find words that match a specific spelling pattern. | `--spelled-like "all*e"`  |
| `--related-word` | Find words based on specific relationships like synonyms, antonyms, or usage. | `--related-word ant joy` (returns antonyms for "joy") |

## Optional Flags

These flags allow you to customize your query, providing context, limiting results, or retrieving additional metadata:

| Flag              | Description                                | Example |
| ----------------- | ------------------------------------------ | ------- |
| `-c`, `--count`   | Display the total number of words returned. | `--count` |
| `--help`          | Display a help message explaining all commands and options. | `polyhymnia --help` |  
| `--left-context`  | Specify words that should appear before the results. | `--left-context "happy"` |  
| `--max`           | Limit the number of results. (Default: 100) | `--max 5` to return only 5 words. |  
| `--metadata`      | Return additional [metadata](#metadata-options), such as word frequency, pronunciation, or syllables. | `--metadata fp` for frequency and pronunciation. |  
| `-q`, `--show-query` | Display the Datamuse API URL used for the query. | `--show-query` |
| `--right-context` | Specify words that should appear after the results. | `--right-context "day"`  |  
| `--topics`        | Filter results by topic (comma-separated values). | `--topics "flight,aviation"` |
| `--version`       | Show the current version of Polyhymnia. | `polyhymnia --version` |
| `--vocabulary`    | Specify a vocabulary to search within. | `--vocabulary enwiki` to use the English Wikipedia vocabulary. |

## Metadata Flags

Use these flags to include specific types of metadata in your query results:

| Flag            | Description                         |  
| --------------- | :---------------------------------- |  
| `-d`, `--def`   | Include definitions in results.     |  
| `-f`, `--freq`  | Include word frequency in results.  |  
| `-p`, `--pos`   | Include parts of speech in results. |  
| `-r`, `--pro`   | Include pronunciation in results.   |  
| `-s`, `--score` | Include score/ranking in results.   |  
| `-y`, `--syl`   | Include syllable count in results.  |  

### The `--metadata` flag

The `--metadata` flag allows you to request additional information about each word returned by Polyhymnia. Use a single-letter shorthand to quickly request common metadata:

| Letter | Meaning        |  
| :----: | -------------- |  
|   d    | Definitions    |  
|   f    | Word frequency |  
|   p    | Parts of speech|  
|   r    | Pronunciation  |  
|   s    | Syllables      |  

You can combine metadata types to return multiple pieces of information. For example, `--metadata df` is equivalent to using `--def --freq`.

**Example:**

This query finds words similar to "happiness," includes definitions and word frequency data, and limits the results to the top two:

```bash
polyhymnia --max 2 --means-like "happiness" --metadata df
```

```text
felicity
   Frequency: 1.319658
   Definition: n  (uncountable) Happiness; (countable) an instance of this.
   Definition: n  (uncountable) An apt and pleasing style in speech, writing, etc.; (countable) an apt and pleasing choice of words.

pleasure
   Frequency: 34.032901
   Definition: n  (uncountable) A state of being pleased or contented; gratification.
   Definition: n  (countable) A person, thing, or action that causes enjoyment.
```

For further customization and advanced features, explore the [Usage Guide](/docs/usage), [Advanced Flags](advanced), and [Examples](/docs/examples).
