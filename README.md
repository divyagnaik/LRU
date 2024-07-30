### Problem Statement
- Build a simple cache system in Go with fixed capacity and Least Recently Used (LRU) eviction policy

### Approach
- Use a combination of a doubly linked list and a hash map. 
- The doubly linked list will help keep track of the order of access, and the hash map will provide efficient access to the cache items.
