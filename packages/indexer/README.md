# Tinysearch: Indexer

This is a subpackage in the tinysearch project. 
This package handles indexing (and re-indexing) content to be used in the search engine. 

## Design principles 

I want to make sure this indexer is both respectful to the sites on it's list, while still regularly checking and indexing posts.

- **The indexer should be respectful of traffic**: we should fetch site content every 6 hours. This will introduce some lead-time into the system, but will also mean that the indexer doesn't spam the sites we want to search through. 
  - It may eventually be worth introducing some heuristic (maybe resulte popularity?) to increase re-index frequency.    
- **The indexer should regularly update content**: we want searches to be relevant, and on relevant content. With that in mind, we should be continiously checking to see if the site *can* respectfully be indexed.
- **The indexer uses RSS as it's source of truth**: unlike a traditional search engine, all of our documents are pulled in via RSS feeds. We eventually want our search to be strictly semantic so we don't need to "crawl" pages via anchor tags like google would. This also makes sure our indexed document list stays manageable and indie. 

## Database Setup

for UUID's to work you have to run the following in the datbaase when it spins up. I should look at automating this somehow.

```SQL
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```