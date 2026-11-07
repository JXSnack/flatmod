package cache

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
)

func Write(path string, data []byte) error {
	// skip if cache isn't initialized
	if InfoDir == "" || FileDir == "" {
		return fmt.Errorf("cache isn't initialized, skipping write for %q", path)
	}

	// create directories
	diskPath := filepath.Join(FileDir, path)
	sumPath := filepath.Join(InfoDir, path+".sum")

	if err := os.MkdirAll(filepath.Dir(diskPath), 0700); err != nil {
		return fmt.Errorf("failed cache write: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(sumPath), 0700); err != nil {
		return fmt.Errorf("failed cache write: %w", err)
	}

	// write file
	if err := os.WriteFile(diskPath, data, 0644); err != nil {
		return fmt.Errorf("failed cache write: %w", err)
	}

	// update checksum
	checksum := md5.Sum(data)
	if err := os.WriteFile(sumPath, checksum[:], 0644); err != nil {
		os.Remove(diskPath) // remove cached file
		return fmt.Errorf("failed checksum write: %w", err)
	}

	cachedChecksums[path] = checksum
	return nil
}
