# Tinysearch: Indexer

This is a subpackage in the tinysearch project. 
This package handles indexing (and re-indexing) content to be used in the search engine. 

## Database Setup

for UUID's to work you have to run the following in the datbaase when it spins up. I should look at automating this somehow.

```SQL
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```