# tinysearch 🔍
an LLM-enabled search engine for blogs and curated sites.

## Features
- Curated search pool of developer blogs, newsletters, and content.
- Focus on search/doc similarity > popularity. This search engine shies away from "PageRank" style of search in which more clicks to the page makes it more "relevant".
- Custom fine-tuned LLM for search that is constantly tuned to content in crawler.
- Disjoint crawler and search engine allows for constant updates to the search index.

## Further reading 

- [Search Engine (computing)](https://en.wikipedia.org/wiki/Search_engine_(computing)?useskin=vector): good ole' wikipedia article that goes over the history and components of a search engine. 
- [Demystifying LLM-Driven Search: Stop Comparing Embeddings or VectorDBs and Start Fine-Tuning](https://medium.com/thirdai-blog/demystifying-llm-driven-search-stop-comparing-embeddings-or-vectordbs-and-start-fine-tuning-d9b6791146fe): article that discusses the benefits of fine-tuning foundational models for search instead of using strictly embeddings to allow them to search 
- [A Search Engine in 80 lines of Python](https://www.alexmolas.com/2024/02/05/a-search-engine-in-80-lines.html): a fun exploration of creating a search engine in python with 80 lines of code.
