# tinysearch 🔍
a search engine for blogs and curated sites.

## Features
- Curated search pool of developer blogs, newsletters, and content.
- Focus on search/doc similarity > popularity. This search engine shies away from "PageRank" style of search in which more clicks to the page makes it more "relevant".
- Disjoint crawler and search engine allows for constant updates to the search index.

## Further reading 

- [Search Engine (computing)](https://en.wikipedia.org/wiki/Search_engine_(computing)?useskin=vector): good ole' wikipedia article that goes over the history and components of a search engine. 
- [A Search Engine in 80 lines of Python](https://www.alexmolas.com/2024/02/05/a-search-engine-in-80-lines.html): a fun exploration of creating a search engine in python with 80 lines of code.
- [Dr. Mark Smucker's "MSCI 541: Search Engines" Lectures](https://www.youtube.com/@msci541-searchengines3): a whole channel containing public lectures from Dr. Smucker of Waterloo Universty's graduate level Search Engine course.
- [Search Engine Design Interview Prep](https://www.youtube.com/watch?v=0LTXCcVRQi0): a high-level exploration of "scaleable" search engine architecture that you may see used in a "google" or "bing" like search engine. This is a bit overkill for this project but is an interesting watch none-the-less.  
- [FAISS Tutorial](https://www.pinecone.io/learn/series/faiss/faiss-tutorial/): similarity search tool that tinysearch will use to return documents. 
- [Demystifying LLM-Driven Search: Stop Comparing Embeddings or VectorDBs and Start Fine-Tuning](https://medium.com/thirdai-blog/demystifying-llm-driven-search-stop-comparing-embeddings-or-vectordbs-and-start-fine-tuning-d9b6791146fe): article that discusses the benefits of fine-tuning foundational models for search instead of using strictly embeddings to allow them to search. Will likely add-on some LLM embedding search later. 