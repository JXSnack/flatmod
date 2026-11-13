package idea

import (
	"os"
	"strings"
)

// ResolveVars resolves IntelliJ path variables in strings, like $PROJECT_DIR$ or $USER_HOME$.
func ResolveVars(raw string) string {
	wd, err := os.Getwd()
	if err != nil {
		wd = ""
	}

	home, err := os.UserHomeDir()
	if err != nil {
		home = ""
	}

	replacements := map[string]string{
		"$PROJECT_DIR$": wd,
		"$USER_HOME$":   home,
		"$USER_DIR$":    home,
	}

	for shortcut, value := range replacements {
		raw = strings.ReplaceAll(raw, shortcut, value)
	}

	return raw
}
