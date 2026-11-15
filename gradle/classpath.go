package gradle

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ClassPath struct {
	Minecraft          string   `json:"minecraft"`
	MinecraftLibraries []string `json:"minecraftLibraries"`
	Dependencies       []string `json:"dependencies"`
}

// GetClasspath generates the libraries through Gradle
func (w *Workspace) GetClasspath() (ClassPath, error) {
	err := w.ExecGradlew([]string{"flatmodClasspath", "--quiet"})
	if err != nil {
		return ClassPath{}, err
	}

	data, err := os.ReadFile(filepath.Join(w.Dir(), "fmod-classpath.json"))
	if err != nil {
		return ClassPath{}, err
	}

	var classPath ClassPath
	err = json.Unmarshal(data, &classPath)
	if err != nil {
		return ClassPath{}, fmt.Errorf("cannot read classpath: %w", err)
	}

	return classPath, nil
}
