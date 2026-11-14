package remap

import (
	"fmt"
	"os"
	"os/exec"
)

// Jar remaps a single jar file at the input path. The remapped jar is saved at the
// output path.
func Jar(inputPath string, outputPath string, settings Settings) error {
	remapperPath, err := DownloadTinyRemapper()
	if err != nil {
		return fmt.Errorf("failed to download remapper: %w", err)
	}

	mappingPath, err := DownloadMappings(settings)
	if err != nil {
		return fmt.Errorf("failed to download mappings: %w", err)
	}

	args := []string{
		"-jar", remapperPath,
		inputPath,     // input
		outputPath,    // output
		mappingPath,   // .tiny file
		settings.From, // e.g. intermediary
		settings.To,   // e.g. official
	}
	cmd := exec.Command("java", args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
