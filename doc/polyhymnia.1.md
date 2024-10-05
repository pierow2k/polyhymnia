% POLYHYMNIA(1) Version {{version}} | General Commands Manual
%
% {{date}}

NAME
====

**polyhymnia** - query the Datamuse API from the command line

SYNOPSIS
========

| **polyhymnia** [options]

DESCRIPTION
===========

**polyhymnia** provides a means to query the Datamuse API from the command-line


Key Features
------------

\- **Find words by meaning, sound, or spelling**: Query the Datamuse API for words similar to, sounding like, or spelled like a given word.  
\- **Retrieve related words**: Get words related to a specific concept using the API's `rel` codes.  
\- **Optional search filters**: Limit results by specifying vocabulary, context, maximum number of results, and more.  
\- **Word metadata**: Receive additional information such as word frequency and pronunciation.  


OPTIONS
=======

Required Flags
--------------

**--means-like**
:    Find words with a meaning similar to this string  

**--sounds-like**
:    Find words that sound similar to this string  

**--spelled-like**
:    Find words spelled similarly to this string  

Optional Flags
--------------

**-c, --count**
:    Show number of words returned by query

**-h, \-\-help**  
:    print the polyhymnia command syntax usage message, and exit

**--left-context**  
:    Provide left context for the search (i.e., words that appear immediately before)

**--max**  
:    Limit the number of results returned (e.g., --max 5)

**--meta-data**  
:    Specify metadata flags (e.g., dpsrf for definitions, parts of speech, syllables) (Refer to Metadata below)

**--related-word**
:    Retrieve words related to a specific concept (multiple values allowed)

**--right-context**  
:    Provide right context for the search (i.e., words that appear immediately after)

**-q, --show-query**
:    Display the Datamuse API URL used for the query.

**--topics**  
:    List topics to filter results (comma-separated values)

**-v, \-\-version**
:    print version information, and exit.

**--vocabulary**  
:    Specify a vocabulary to search in (e.g., "enwiki")

Metadata Display Flags
----------------------

Metadata display flags control the metadata that is displayed with search results.

**-d, --def**
:    Include definitions in results.  
**-f, --freq**
:    Include word frequency in results.  
**-p, --pos**
:    Include parts of speech in results.  
**-r, --pro**
:    Include pronunciation in results.  
**-s, --score**
:    Include score/ranking in results.  
**-y, --syl**
:    Include syllable count in results.  



Metadata
--------

The **--metadata** flag allows you to request additional information about each word returned by Polyhymnia. Use a single-letter shorthand to quickly request common metadata:

**none**
:    Return only the word and score. (Default)  
**d**
:    Definitions  
**f**
:    Word frequency  
**p**
:    Parts of speech  
**r**
:    Pronunciation  
**s**
:    Syllable count  

You can combine metadata types to return multiple pieces of information. For example, **--metadata df** is equivalent to using **--def --freq**.

Related Word
------------

The **--related-word** option allows you to find words that share specific lexical relationships with the input term, using a set of predefined codes. These codes represent various semantic, phonetic, and corpus-statistics-based relations. You can use multiple `--related-word` flags to retrieve words with different relations.

- **jja**: Popular nouns modified by the given adjective, per Google Books Ngrams.  
  Example: *gradual* `->` *increase*
  
- **jjb**: Popular adjectives used to modify the given noun, per Google Books Ngrams.  
  Example: *beach* `->` *sandy*

- **syn**: Synonyms (words contained within the same WordNet synset).  
  Example: *ocean* `->` *sea*

- **trg**: "Triggers" (words statistically associated with the query word in the same text).  
  Example: *cow* `->` *milking*

- **ant**: Antonyms (per WordNet).  
  Example: *late* `->` *early*

- **spc**: "Kind of" (direct hypernyms, per WordNet).  
  Example: *gondola* `->` *boat*

- **gen**: "More general than" (direct hyponyms, per WordNet).  
  Example: *boat* `->` *gondola*

- **com**: "Comprises" (direct holonyms, per WordNet).  
  Example: *car* `->` *accelerator*

- **par**: "Part of" (direct meronyms, per WordNet).  
  Example: *trunk* `->` *tree*

- **bga**: Frequent followers (per Google Books Ngrams).  
  Example: *wreak* `->` *havoc*

- **bgb**: Frequent predecessors (per Google Books Ngrams).  
  Example: *havoc* `->` *wreak*

- **hom**: Homophones (words that sound alike).  
  Example: *course* `->` *coarse*

- **cns**: Consonant match.  
  Example: *sample* `->` *simple*

Vocabulary
----------

The `--vocabulary` flag specifies which vocabulary to use.

\- If no vocabulary is specified, a 550,000-term vocabulary of English words and multiword expressions is used.  
\- **`es`**  - a 500,000-term vocabulary of words from Spanish-language books.  
\- **`enwiki`** - approximately 6 million-term vocabulary of article titles from the English-language Wikipedia, updated monthly.  

OUTPUT
======

Results include:

**Word**
:    The vocabulary word.  

**Score**
:    Relevance score.  

**Syllables**
:    Number of syllables.  

**Pronunciation**
:    Phonetic representation (if available).  

**Frequency**
:    How common the word is.  

**Parts of speech**
:    Noun, verb, etc.  

**Definition**
:    The meaning of the word (if available).  

Parts of Speech
---------------

One or more part-of-speech codes will be returned with the search results.

* **adj** - Adjective  
* **adv** - Adverb  
* **n** - Noun  
* **v** - Verb  
* **u** - The part of speech is none of these or cannot be determined.  

Multiple entries will be added when the word's part of speech is ambiguous, with the most popular part of speech listed first. This field is derived from an analysis of Google Books Ngrams data.

EXIT STATUS
===========

**polyhymnia** returns the following exit status codes:  

* **0** - Successful  
* **1** - An error occurred  

BUGS
====

Report bugs to the polyhymnia GitHub repository.  
[https://github.com/pierow2k/polyhymnia](https://github.com/pierow2k/polyhymnia)

AUTHOR
======

Written by Pierow2K.

COPYRIGHT
=========

Copyright (C) {{copyright_date}} Pierow2K.

LICENSE
=======

polyhymnia is licensed under the MIT License.  
