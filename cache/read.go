package cache

import (
	"fmt"
	"os"
)

// Read compares the cached checksum to the requested checksum. If they are
// the same, it returns the cached value. If the two checksums are not equal,
// the get func is used to get the information. The returned value from get
// is automatically saved to the cache.
//
// If the cache isn't initialized, Read always returns get and prints a warning.
// The function does not return any errors by itself, it only escalates errors
// from get. If get returns an error, its value is not cached.
//
// Cached values are reused between flatmod projects to avoid unnecessary
// information duplication.
func Read(path string, checksum [16]byte, get func() ([]byte, error)) ([]byte, error) {
	// skip if cache isn't initialized
	if InfoDir == "" || FileDir == "" {
		fmt.Fprintf(os.Stderr, "cache isn't initialized, skipping read for %q", path)
		return get()
	}

	// check if cached
	if _, err := os.Stat(path); err == nil {
		// check checksum
		cachedSum, err := GetChecksum(path)
		if err == nil {
			if cachedSum == checksum {
				// read from cache
				data, err := os.ReadFile(path)
				if err == nil {
					return data, nil
				}
			}
		}
	}

	// escalate errors
	data, err := get()
	if err != nil {
		return data, err
	}

	// write
	err = Write(path, checksum, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}

	return data, nil
}
