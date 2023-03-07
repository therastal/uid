# UID

I wanted to use [ULIDs](https://github.com/ulid/spec) as SQLite rowid aliases, but [the maximum size SQLite allows is 64 bits](https://www.sqlite.org/lang_createtable.html#rowid). This implementation allows me to benefit from ULID semantics while retaining the performance of SQLite rowid-based operations.

This is mostly a Python port of Mediagone's [Small UID](https://github.com/Mediagone/small-uid) library for PHP. The main functional difference is allowing a set of UIDs to be passed into the constructor in order to ensure true uniqueness across a given set of generated IDs, though that only becomes an issue when thousands of UIDs are being generated per millisecond.

## Usage

```python
from uid import UID

COUNT = 100000

with_duplicates = set()
unique_only = set()

for _ in range(COUNT):
    with_duplicates.add(UID())
    unique_only.add(UID(existing=unique_only))

assert len(with_duplicates) != COUNT
assert len(unique_only) == COUNT
```
