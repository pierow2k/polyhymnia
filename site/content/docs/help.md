---
linkTitle: Help
title: 'Getting Help'
weight: 5
---
### The Help Flag

Polyhymnia's `--help` flag provides an overview of the application's usage
and the various [flags](/docs/flags) that can be used to customize queries
and output.

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
      --max int                    Maximum number of results to return (1-1000) (default 100)
  -l, --means-like                 Words with meaning similar to this string
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

### Manual

Polyhymnia includes a man page in both troff (man page) and in PDF formats.

- **Man Page in troff Format** [troff format](/files/polyhymnia.1)
- **man Page in PDF Format** [PDF](/files/polyhymnia.1.pdf)

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
