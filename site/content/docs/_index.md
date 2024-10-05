---
title: 'Polyhymnia Documentation'
---
Welcome to **Polyhymnia**—your gateway to unlocking the mysteries of the
English language, all from the comfort of your command line! Whether
you're a writer in search of the perfect word, a poet crafting rhymes, a
language lover, or even a cryptographer diving into word patterns,
Polyhymnia is here to elevate your word searches. Powered by the mighty
**Datamuse API**, it’s built to help you explore words based on their
meaning, sound, spelling, and much more.

## What is Polyhymnia?

Polyhymnia is a command-line tool that leverages the Datamuse API to
enable users to search for words based on [meaning](/docs/usage/find-words-by-meaning/), [sound](/docs/usage/find-words-by-sound/), [spelling](/docs/usage/find-words-by-spelling/), and [relationships](/docs/usage/find-related-words/). It's designed for writers, poets, language enthusiasts, and anyone who needs to explore the nuances of the English and Spanish languages.

## 🎯 Use Cases

Polyhymnia isn’t just for language geeks—it has a wide range of practical applications!

1. **Creative Writing & Poetry**: Perfect for finding synonyms, rhymes, or words that fit your poem’s rhythm or themes.
2. **Songwriting**: Quickly grab rhymes or similar-sounding words to smooth out your lyrics.
3. **Solving Word Puzzles**: Find words that match specific clues based on patterns, meanings, or letter combinations.
4. **Marketing & Branding**: Generate catchy names, taglines, or slogans by searching for words that match a sound or idea.
5. **Language Learning**: Boost your vocabulary with related words and synonyms, filtered by context.
6. **Game Development**: Create word puzzles, crosswords, or word association games with dynamic word lists.
7. **Linguistic Research**: Query the API for insights into phonetics, semantics, or lexical relationships.

### 🚀 Word Search Modes

Polyhymnia puts the power of words at your fingertips with four primary search modes:

* **Meaning**: Discover words with similar meanings using the `--means-like` flag. Perfect for finding that elusive synonym!
* **Sound**: Want words that sound alike? Use the `--sounds-like` flag and find perfect rhymes or assonance matches.
* **Spelling**: Search for words that match a specific pattern with the `--spelled-like` flag, even using wildcards!
* **Related Words**: Explore synonyms, antonyms, homophones, and more using the `--related-word` flag with custom relationship codes.

{{< cards cols="2" >}}
  {{< card link="/docs/usage/"
        title="Usage"
        subtitle="Learn the ins and outs of how to use Polyhymnia."
        icon="document" >}}
{{< /cards >}}

### Customization Options

#### Required Flags

Start every search with one of the essential flags—`--means-like`, `--sounds-like`, `--spelled-like`, or `--related-word`—and Polyhymnia will do the rest!

#### Optional Flags

Fine-tune your results with powerful optional flags like `--max`, `--left-context`, `--right-context`, `--topics`, and `--vocabulary`. Whether you're narrowing down topics or adjusting context, Polyhymnia has you covered.

#### Metadata Flags

Need more details? Fetch definitions, frequency, part of speech, and more with [metadata flags](/docs/flags/) like `--def`, `--freq`, and `--pos`. Or, use the convenient `--metadata` flag to grab multiple at once.

{{< cards cols="2" >}}
  {{< card link="/docs/flags/"
        title="Flags and Options"
        subtitle="Explore all the flags and options Polyhymnia offers."
        icon="document" >}}
{{< /cards >}}

### Advanced Features

#### Related Word Codes

Looking for specific word relationships? Use [advanced codes](/docs/usage/find-related-words/) like `syn` (synonyms), `ant` (antonyms), `hom` (homophones), `trg` (trigger words), and more to explore deeper word connections.

{{< cards cols="2" >}}
  {{< card link="/docs/usage/find-related-words/"
        title="Find Related Words"
        subtitle="Dig deeper into related words with detailed options."
        icon="document" >}}
{{< /cards >}}

#### Vocabulary

Customize your searches by specifying vocabularies. From the default English language dictionary to specialized sets like `enwiki` (English Wikipedia) or even `es` (Spanish), Polyhymnia’s `--vocabulary` flag gives you the flexibility you need.

{{< cards cols="2" >}}
  {{< card link="/docs/flags/advanced/#vocabulary"
        title="Advanced Vocabulary Options"
        subtitle="Unlock advanced vocab options for tailored searches."
        icon="document" >}}
{{< /cards >}}

### ⚙️ Installation

Getting started is easy! Polyhymnia requires Go and can be installed via `go install` or by cloning the repository and building it manually.

{{< cards cols="2" >}}
  {{< card link="/docs/installation"
        title="Installation"
        subtitle="Follow the steps to install Polyhymnia."
        icon="document" >}}
{{< /cards >}}

### 📚 Help & Support

Need assistance? Polyhymnia’s got your back:

* Use the `--help` flag for a quick overview of commands and options.
* Check out the man page for detailed information.
* Found a bug or have a feature request? Open an issue on the project repository!

{{< cards cols="2" >}}
  {{< card link="/docs/help"
        title="Help"
        subtitle="Get all the help you need here."
        icon="document" >}}
{{< /cards >}}

### Key Takeaways

* **Polyhymnia** is a versatile, user-friendly tool for exploring the vast world of words and their relationships.
* With intuitive commands and robust options, it’s designed to cater to a wide range of users and use cases.
* Comprehensive documentation and active support make your journey smooth and enjoyable.
