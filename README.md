# MediaSync

A simple tool written in Go to compare iTunes music libraries.

Depending on iTunes versions and whether or not the "Keep iTunes media folder organized" preference is enabled, the same track could have different filenames across different iTunes libraries. So this works by using some simple heuristics to match tracks by filename then falls back to matching by file hash.



### Usage

```
itunesSync --dir1 <iTunes library 1> --dir2 <iTunes library 2> --mode [import|sync]
```

- The *import* option will copy all files from _dir2_ not in _dir1_ to a separate directory. You can then add these files to iTunes by using the "Add to Library" option
- The *sync* option will copy all missing files from _dir1_ to _dir2_ and vice versa. (Semi useful for creating external backups)

Options:
- *--dryRun* will not copy or manipulate any files. Will just print out what it would do
- *--v* enable very verbose output