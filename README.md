![Polyhymnia](./doc/polyhymnia.jpg)

[![Website](https://img.shields.io/website?url=https%3A%2F%2Fpolyhymnia.daspyro.de&link=https%3A%2F%2Fpolyhymnia.daspyro.de)](https://polyhymnia.daspyro.de) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pierow2k/polyhymnia) ![License](https://img.shields.io/github/license/pierow2k/polyhymnia) ![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/pierow2k/polyhymnia) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/47d449d11aab41258192067c99c0ebde)](https://app.codacy.com/gh/pierow2k/polyhymnia/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

# Polyhymnia — Unleashing the magic of words, one query at a time!

**Polyhymnia** - A powerful command-line tool that lets you dive deep
into the English language. Whether you're a writer, poet, or language
enthusiast, Polyhymnia helps you discover words based on their meaning,
sound, spelling, and much more—all from the comfort of your terminal.
Polyhymnia is powered by the the incredible [Datamuse API](https://www.datamuse.com/api/).

**You can find more detailed information on the Polyhymnia website at [https://polyhymnia.daspyro.de](https://polyhymnia.daspyro.de)**

<!-- TABLE OF CONTENTS -->
<details closed="closed">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
	<li><a href="#what-is-polyhymnia">What is Polyhymnia?</a></li>
  <li><a href="#-features">Features</a></li>
  <li><a href="#-use-cases">Use Cases</a></li>
  <li><a href="#-example-queries">Example Queries</a></li>
  <li><a href="#-installation">Installation</a></li>
  <li><a href="#-usage">Usage</a></li>
  <li><a href="#command-line-options">Command-line Options</a></li>
  <li><a href="#metadata-options">Metadata Options</a></li>
  <li><a href="#related-word">Related Word</a></li>
  <li><a href="#vocabulary">Vocabulary</a></li>
  <li><a href="#output">Output</a></li>
  <li><a href="#parts-of-speech">Parts of Speech</a></li>
  <li><a href="#-getting-help">Getting Help</a></li>
  <li><a href="#man-page">Man Page</a></li>
  <li><a href="#open-an-issue">Open an Issue</a></li>
  <li><a href="#-license">License</a></li>
  <li><a href="#-contributing">Contributing</a></li>
</details>

## What is Polyhymnia?

Polyhymnia is a command-line tool that leverages the Datamuse API to enable users to search for words based on [meaning](https://polyhymnia.daspyro.de/docs/usage/find-words-by-meaning/), [sound](https://polyhymnia.daspyro.de/docs/usage/find-words-by-sound/), [spelling](https://polyhymnia.daspyro.de/docs/usage/find-words-by-spelling/), and [relationships](https://polyhymnia.daspyro.de/docs/usage/find-related-words/). It's designed for writers, poets, language enthusiasts, and anyone who needs to explore the nuances of the English and Spanish languages.

## 🚀 Features

Polyhymnia puts the power of words at your fingertips with four primary
search modes:

* **Meaning**: Discover [words with similar meanings](https://polyhymnia.daspyro.de/docs/usage/find-words-by-meaning/) using the `--means-like` flag. Perfect for finding that elusive synonym!
* **Sound**: Want [words that sound alike](https://polyhymnia.daspyro.de/docs/usage/find-words-by-sound/)? Use the `--sounds-like` flag and find perfect rhymes or assonance matches.
* **Spelling**: Search for [words that match a specific pattern](https://polyhymnia.daspyro.de/docs/usage/find-words-by-spelling/) with the `--spelled-like` flag, even using wildcards!
* **Related Words**: Explore [synonyms, antonyms, homophones, and more](https://polyhymnia.daspyro.de/docs/usage/find-related-words/) using the `--related-word` flag with custom relationship codes.

## 🎯 Use Cases

Polyhymnia isn’t just for language geeks—[it has a wide range of practical
applications!](https://polyhymnia.daspyro.de/docs/#-use-cases)

1. **Creative Writing & Poetry**: Perfect for finding synonyms, rhymes, or words that fit your poem’s rhythm or themes.
2. **Songwriting**: Quickly grab rhymes or similar-sounding words to smooth out your lyrics.
3. **Solving Word Puzzles**: Find words that match specific clues based on patterns, meanings, or letter combinations.
4. **Marketing & Branding**: Generate catchy names, taglines, or slogans by searching for words that match a sound or idea.
5. **Language Learning**: Boost your vocabulary with related words and synonyms, filtered by context.
6. **Game Development**: Create word puzzles, crosswords, or word association games with dynamic word lists.
7. **Linguistic Research**: Query the API for insights into phonetics, semantics, or lexical relationships.

## 🛠️ Example Queries

**[More examples are available on the Polyhymnia website.](https://polyhymnia.daspyro.de/docs/examples/)**

### Find Words by Meaning

Search for words with a meaning similar to **"joy"** and limit the results to 3 words using the **`--max`** [flag](https://polyhymnia.daspyro.de/docs/flags/#optional-flags).

```bash
polyhymnia --max 3 --means-like "joy"
```

```text
gladden

delight

rejoice
```

### Find Words by Sound

This example finds words that sound like **"kevin"** and limits the results to
5 words.

```bash
polyhymnia --max 5 --sounds-like "kevin"
```

```text
kevin

kevyn

kevan

cave in

cave-in
```

### Find Words by Spelling

This example searches for words that **begin with `all`** and **end in the
letter `e`**. It limits the results to 5 words.
The results include matching words like "alleviate" and
"allude".

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

### Find Related Words

#### The `--related-word` flag

To find words based on specific relationships like synonyms, antonyms, or
homophones, you can use the `--related-word` flag and one of the related
word search codes.

#### Related Word Search Codes

The related word search codes allow you to specify the type of
relationship you're interested in.

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

#### Examples

##### Find antonyms for the word "joy"

This query returns **antonyms** of the word **"joy"** by using the [related-word
code](https://polyhymnia.daspyro.de/docs/usage/find-related-words/#related-word-search-codes) **`ant`**. We limit the number of returned words to 3 using the **`--max`**
[flag](https://polyhymnia.daspyro.de/docs/flags/#optional-flags).

```bash
polyhymnia --max 3 --related-word ant joy
```

```text
sad

sorrow

sadden
```

##### Find homophones for the word "joy"

This query returns **homophones** of the word **"joy"** by using the [related-word
code](https://polyhymnia.daspyro.de/docs/usage/find-related-words/#related-word-search-codes) **`hom`**.

```bash
polyhymnia --related-word hom joy
```

```text
joye
```

##### Find synonyms for the word "joy"

This query returns **synonyms** of the word **"joy"** by using the [related-word
code](https://polyhymnia.daspyro.de/docs/usage/find-related-words/#related-word-search-codes) **`syn`**. We limit the number of returned words to 3 using the **`--max`**
[flag](https://polyhymnia.daspyro.de/docs/flags/#optional-flags).

```bash
polyhymnia --max 3 --related-word syn joy
```

```text
delight

pleasure

rejoice
```


## ⚙️ Installation

First, make sure you have [Go](https://golang.org/dl/) installed on your system.

### Install via `go install`

```bash
go install github.com/pierow2k/polyhymnia@latest
```

### Clone and Build

Alternatively, you can clone the repo and build the application manually:

```bash
git clone https://github.com/pierow2k/polyhymnia.git
cd polyhymnia
sudo make install
```

This will create an executable called `polyhymnia` and install the `man`
page to `/usr/local/share/man/man1`.

## 📖 Usage

```bash
./polyhymnia [options]
```

### Command-line Options

#### Required Flags

One of these flags is **required** for every query:

| Flag              | Description                                                       |  
| ----------------- | ----------------------------------------------------------------- |  
| `--means-like`    | Find words with similar meaning ([Example](#find-words-by-meaning)) |  
| `--sounds-like`   | Find words that sound similar ([Example](#find-words-by-sound))|  
| `--spelled-like`  | Find words with a similar spelling (supports wildcards) ([Example](#find-words-by-spelling)) |  
| `--related-word` | Find words based on [specific relationships](#find-related-words) like synonyms, antonyms, or usage. |

#### Optional Flags

| Flag              | Description                                                       |  
| ----------------- | ----------------------------------------------------------------- |  
| `--left-context`  | Specify words that should appear before the results               |  
| `--right-context` | Specify words that should appear after the results                |  
| `--max`           | Limit the number of results (e.g., `--max 5`)                     |  
| `--metadata`      | Return additional [metadata information](#metadata-options) such as definitions or frequency |  
| `--related-word`  | Refer to [Related Word](#related-word) below (Multiple values allowed)  |  
| `--topics`        | Filter results by topic (comma-separated values) ([Example](#find-related-words-with-context)) |  
| `--vocabulary`    | Specify a vocabulary to search (e.g., "enwiki")                   |  
| `--help`          | Show the [help message](https://polyhymnia.daspyro.de/docs/help/#the-help-flag)                                             |  
| `--version`       | Show version information                                          |

#### Metadata Flags

Use these flags to include specific types of metadata in your query results:

| Flag            | Description                         |  
| --------------- | :---------------------------------- |  
| `-d`, `--def`   | Include definitions in results.     |  
| `-f`, `--freq`  | Include word frequency in results.  |  
| `-p`, `--pos`   | Include parts of speech in results. |  
| `-r`, `--pro`   | Include pronunciation in results.   |  
| `-s`, `--score` | Include score/ranking in results.   |  
| `-y`, `--syl`   | Include syllable count in results.  |  

### Metadata Options

The `--metadata` flag allows you to request additional information about
each word returned by Polyhymnia. Use a single-letter shorthand to quickly
request common metadata:

| Letter | Meaning        |  
| :----: | -------------- |  
|   d    | Definitions    |  
|   f    | Word frequency |  
|   p    | Parts of speech|  
|   r    | Pronunciation  |  
|   s    | Syllables      |  

You can combine metadata types to return multiple pieces of information.
For example, `--metadata df` is equivalent to using `--def --freq`.

### Related Word

Polyhymnia's "Related Word" feature offers a deeper dive into lexical
relationships, allowing you to discover words that are connected through
various semantic, phonetic, and usage-based associations. This feature is
especially useful for anyone looking to explore how words are related
beyond mere definitions—whether you're hunting for synonyms, antonyms, or
words commonly paired together. By leveraging advanced search codes, you
can fine-tune your queries to surface words with specific kinds of
relationships, all powered by rich data sources like WordNet and Google
Books Ngrams.

Explore the various relationship types below to uncover new word
connections that can inspire creativity, solve puzzles, or enrich your
linguistic research.

| code | Description | Example |
|:---:|-----|-----|
| jja | Popular nouns modified by the given adjective, per Google Books Ngrams | gradual → increase |
| jjb | Popular adjectives used to modify the given noun, per Google Books Ngrams | beach → sandy |
| syn | Synonyms (words contained within the same WordNet synset) | ocean → sea |
| trg | "Triggers" (words that are statistically associated with the query word in the same piece of text.) | cow → milking |
| ant | Antonyms (per WordNet) | late → early |
| spc | "Kind of" (direct hypernyms, per WordNet) | gondola → boat |
| gen | "More general than" (direct hyponyms, per WordNet) | boat → gondola |
| com | "Comprises" (direct holonyms, per WordNet) | car → accelerator |
| par | "Part of" (direct meronyms, per WordNet) | trunk → tree |
| bga | Frequent followers (w′ such that P(w′|w) ≥ 0.001, per Google Books Ngrams) | wreak → havoc |
| bgb | Frequent predecessors (w′ such that P(w|w′) ≥ 0.001, per Google Books Ngrams) | havoc → wreak |
| hom | Homophones (sound-alike words) | course → coarse |
| cns | Consonant match | sample → simple |

### Vocabulary

The `--vocabulary` flag specifies which vocabulary to use.

* If no vocabulary is specified, a 550,000-term vocabulary of English words and multiword expressions is used.
* `es`  - a 500,000-term vocabulary of words from Spanish-language books.
* `enwiki` - approximately 6 million-term vocabulary of article titles from the English-language Wikipedia, updated monthly.

## Output

Polyhymnia provides the following results:

* **Word**: The vocabulary word.
* **Score**: Relevance score.
* **Syllables**: Number of syllables.
* **Pronunciation**: Phonetic representation (if available).
* **Frequency**: How common the word is.
* **Parts of speech**: Noun, verb, etc. (See [Parts of Speech](#parts-of-speech) below)
* **Definition**: The meaning of the word (if available).

Use the [`--metadata`](#metadata-options) flag to customize the data returned.

### Parts of Speech

One or more part-of-speech codes will be returned with the search results. 

* `adj` means adjective
* `adv` means adverb
* `n` means noun
* `v` means verb
* `u` means that the part of speech is none of these or cannot be determined.

*Polyhymnia adds multiple entries when the word's part of speech is ambiguous,
listing the most popular part of speech first.*

## 📚 Getting Help

Need assistance? Polyhymnia’s got your back:

* The [Polyhymia website](https://polyhymnia.daspyro.de) includes detailed documentation.
* Use the `--help` flag for a quick overview of commands and options.
* Check out the man page for detailed information.
* Found a bug or have a feature request? Open an issue on the project repository!

### Help Command

Polyhymnia's `--help` flag provides an overview of the application's usage
and the various [flags](https://polyhymnia.daspyro.de/docs/flags/) that
can be used to customize queries and output.

```bash
polyhymnia --help
```

```text
Polyhymnia leverages the Datamuse API to enable users to search for words
based on meaning, sound, spelling, and relationships.

Usage:
  polyhymnia [search term] [flags]

Flags:
  -c, --count                      Show number of words returned by query
  -d, --def                        Include definitions in results
  -f, --freq                       Include frequency in results
  -h, --help                       help for polyhymnia
      --left-context string        Left context
      --max int                    Maximum number of results to return (
                                   1-1000) (default 100)
  -l, --means-like                 Words with meaning similar to this
                                   string
      --metadata string            Metadata flags (dfprs)
  -p, --pos                        Include parts of speech in results
  -r, --pro                        Include pronunciation in results
      --related-word stringArray   Related word constraints
      --right-context string       Right context
  -s, --score                      Include score in results
  -q, --show-query                 Show the URL used for the query
  -n, --sounds-like                Words that sound like this string
  -t, --spelled-like               Words spelled like this string
  -y, --syl                        Include syllables in results
      --topics stringArray         Topics (comma-separated)
  -v, --version                    version for polyhymnia
      --vocabulary string          Vocabulary identifier
```

### Man Page

Check out the `man` page in [troff](./doc/polyhymnia.1) or [PDF](./doc/polyhymnia.1.pdf) format for detailed documentation.

```text
POLYHYMNIA(1)               General Commands Manual              POLYHYMNIA(1)

NAME
       polyhymnia - query the Datamuse API from the command line

SYNOPSIS
       polyhymnia options

DESCRIPTION
       polyhymnia  provides  a  means  to query the Datamuse API from the com-
       mand-line

   Key Features
       - Find words by meaning, sound, or spelling: Query the Datamuse API for
       words similar to, sounding like, or spelled like a given word.
       - Retrieve related words: Get words related to a specific concept using
       the API's rel codes.
       - Optional search filters: Limit results by specifying vocabulary, con-
       text, maximum number of results, and more.
       - Word metadata: Receive additional information such as word  frequency
       and pronunciation.

OPTIONS
   Required Flags
       -means-like
              Find words with a meaning similar to this string

       -sounds-like
              Find words that sound similar to this string

       -spelled-like
              Find words spelled similarly to this string

   Optional Flags
…
```

### Open an Issue

Have an idea for a new feature or noticed something that isn’t working
quite right? [Open an issue](https://github.com/pierow2k/polyhymnia/issues)
to let us know. Your feedback helps us keep Polyhymnia reliable and
feature-rich.

## 📝 License

Polyhymnia is released under the MIT License. See the [LICENSE](LICENSE)
file for more details.

## 🤝 Contributing

We love contributions! Feel free to submit a pull request or open an issue
to discuss new ideas or improvements.
