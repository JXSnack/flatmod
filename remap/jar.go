package remap

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/helper"
)

// Jar remaps a single jar file at the input path. The remapped jar is saved at the
// output path. It returns the content of the remapped jar on success.
func Jar(inputPath string, outputPath string, settings Settings) ([]byte, error) {
	helper.Vprintf("remap: %s\n", filepath.Base(inputPath))

	remapperPath, err := DownloadTinyRemapper()
	if err != nil {
		return []byte{}, fmt.Errorf("failed to download remapper: %w", err)
	}

	mappingPath, err := DownloadMappings(settings)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to download mappings: %w", err)
	}

	// run
	args := []string{
		"-jar", remapperPath,
		inputPath,     // input
		outputPath,    // output
		mappingPath,   // .tiny file
		settings.From, // e.g. intermediary
		settings.To,   // e.g. official
	}
	cmd := exec.Command("java", args...)

	if helper.Verbose {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return []byte{}, err
	}

	// read output
	data, err := os.ReadFile(outputPath)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// JarCached runs Jar if the outputPath isn't cached. Therefore, the outputPath will have
// the cache file directory prepended.
func JarCached(inputPath string, outputPath string, settings Settings) error {
	_, err := cache.Read(
		outputPath,
		helper.ZeroSum(),
		func() ([]byte, error) {
			return Jar(inputPath, filepath.Join(cache.FileDir, outputPath), settings)
		},
	)
	return err
}
