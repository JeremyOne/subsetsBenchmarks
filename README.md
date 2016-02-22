A utility for benchmarking various subset algorithms.

Utility will load text files, convert them to arrays, then check if they are subsets of each other using methods provided in the subsets package

Also provided are small set of books for sample comparisons

Example set:
```
go run subsetsBenchmarks bible.txt shakespeare.txt
```

This benchmark depends on these supporting libraries:

[JeremyOne/runeTree] (https://github.com/JeremyOne/runeTree)
[JeremyOne/subsets] (https://github.com/JeremyOne/subsets)
