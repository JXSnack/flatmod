package cache

import (
	"fmt"
	"os"
	"path/filepath"
)

// cachedChecksums stores checksums that were accessed previously and
// updates when the checksum changes. It does not contain every single checksum,
// which is why you should instead use GetChecksum to get the checksum
// of a path.
var cachedChecksums = map[string][16]byte{}

func GetChecksum(path string) ([16]byte, error) {
	diskPath := filepath.Join(InfoDir, path+".sum")

	// skip if cache isn't initialized
	if InfoDir == "" || FileDir == "" {
		return [16]byte{}, fmt.Errorf("cache isn't initialized, skipping GetChecksum for %q", path)
	}

	// check cache
	if checksum, ok := cachedChecksums[path]; ok {
		return checksum, nil
	}

	// check disk
	data, err := os.ReadFile(diskPath)
	if err == nil {
		if len(data) < 16 {
			os.Remove(diskPath)
			return [16]byte{}, fmt.Errorf("invalid checksum length: got %d, expected 16", len(data))
		}

		var checksum [16]byte
		copy(checksum[:], data[:16])

		// update in-memory caches
		cachedChecksums[path] = checksum
		return checksum, nil
	}

	return [16]byte{}, fmt.Errorf("no checksum for %q stored", path)
}
