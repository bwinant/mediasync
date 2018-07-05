package pkg

import (
	"encoding/hex"
	"crypto/md5"
	"io/ioutil"
	"time"
)

// Compares two sets of files and returns all files in files1 that do not have an exact name match in files2
func ExactMatch(files1, files2 []*File) []*File {
	mapper := func(file *File) string {
		return file.Name
	}

	return Match(files1, files2, mapper);
}

// For each file in files1, looks for a corresponding matching file in files2
// mapper is a function used to generate map keys from a property of a file, in order to build a mapping from key -> file
func Match(files1 []*File, files2 []*File, mapper func(*File) string) []*File {
	fileMap := make(map[string]*File)
	for _, f := range files2 {
		key := mapper(f)
		fileMap[key] = f
	}

	diff := make([]*File, 0)

	for _, f := range files1 {
		key := mapper(f)

		match := fileMap[key]
		if match != nil {
			// Have similar files - are they really the same file?
			if f.Size != match.Size {
				diff = append(diff, f)
			} else {
				debugf("Matched: %s to %s", f.AbsolutePath, match.AbsolutePath)
			}
		} else {
			diff = append(diff, f)
		}
	}

	return diff
}

// Compares the results of two directory scans.
// Returns list of files are in d1 but not in d2
func Compare(d1, d2 *DirScan) []*File {
	infof("Comparing %s to %s", d1, d2)

	// 0) Create map with name and normalized name as keys
	fileMap := make(map[string]*File)
	for _, f := range d2.Files {
		fileMap[f.Name] = f
		fileMap[f.AbsolutePath] = f
		fileMap[Normalize(f.Name)] = f
	}

	// 1) Find filename matches
	count := 0
	for _, f := range d1.Files {
		var match *File

		// 1a) Did file already match?
		existing := f.Metadata["match"]
		if existing != "" {
			match = fileMap[existing]
		}

		// 1b) Look for exact name match
		if match == nil {
			match = fileMap[f.Name]
			if match != nil && f.Size == match.Size {
				debugf("Matched: %s to %s", f.AbsolutePath, match.AbsolutePath)
			} else {
				match = nil
			}
		}

		// 1c) Failing that, look for close enough match
		if match == nil {
			match = fileMap[Normalize(f.Name)]
			if match != nil && f.Size == match.Size {
				debugf("Matched: %s to %s", f.AbsolutePath, match.AbsolutePath)
			} else {
				match = nil
			}
		}

		// 1d) Track matches
		if match != nil {
			f.Metadata["match"] = match.AbsolutePath
			match.Metadata["match"] = f.AbsolutePath
		} else {
			count++
		}
	}

	// 2) If there are files without name matches, calculate hashes for all files that didn't match to anything
	if count > 0 {
		infof("Calculating file hashes ...")
		getHashes(d1, nil)

		cb := func(f *File, hash string) {
			fileMap[hash] = f
		}
		getHashes(d2, cb)
	}

	// 3) Match files based on hash
	diff := make([]*File, 0)
	for _, f := range d1.Files {
		if f.Metadata["match"] == "" {
			hash := f.Metadata["hash"]
			if hash != "" {
				match := fileMap[hash]
				if match != nil {
					debugf("Matched: %s to %s", f.AbsolutePath, match.AbsolutePath)
					f.Metadata["match"] = match.AbsolutePath
					match.Metadata["match"] = f.AbsolutePath
				} else {
					diff = append(diff, f)
				}
			} else {
				diff = append(diff, f)
			}
		}
	}

	if len(diff) > 0 {
		infof("%d files in %s are not in %s", len(diff), d1.Name, d2.Name)
	} else {
		infof("All files in %s found in %s", d1.Name, d2.Name)
	}

	return diff
}

func getHashes(d *DirScan, cb func(f *File, hash string)) {
	start := time.Now()
	count := 0

	for _, f := range d.Files {
		if f.Metadata["match"] == "" && f.Metadata["hash"] == "" {
			hash, err := fileHash(f.AbsolutePath)
			if err == nil {
				debugf("Hash: %s = %s", f.AbsolutePath, hash)

				f.Metadata["hash"] = hash
				count++

				if cb != nil {
					cb(f, hash)
				}
			} else {
				errorf("Unable to generate hash for file %s", f.AbsolutePath)
			}
		}
	}

	elapsed := time.Since(start)
	infof("Obtained hashes for %d files in %s in %s", count, d, elapsed)
}

// Don't care about security, use MD5 for performance reasons
// Reuse it to try to save some GC pressure
var hashFn = md5.New()
func fileHash(filename string) (string, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	hashFn.Reset()
	hashFn.Write(buf)
	hash := hashFn.Sum(nil)
	return hex.EncodeToString(hash), nil
}

func stopwatch(start time.Time, label string) {
	elapsed := time.Since(start)
	debugf("%s took %s", label, elapsed)
}