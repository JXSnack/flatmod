package idea

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JXSnack/flatmod/helper"
	"github.com/beevik/etree"
)

// GetJDK tries to find the JDK based on the IntelliJ configuration. If something goes
// wrong, it either defaults back to the home JDK (and prints a warning) or returns an
// error in more severe cases.
func GetJDK() (string, error) {
	project := etree.NewDocument()
	err := project.ReadFromFile(filepath.Join(".idea", "misc.xml"))
	if err != nil {
		fmt.Println("flatmod: using home jdk, because IntelliJ configuration wasn't found")
		return "", nil
	}

	name := project.FindElement("//component[@name='ProjectRootManager']")
	if name == nil {
		return "", fmt.Errorf("no project sdk configured")
	}

	jdk := name.SelectAttrValue("project-jdk-name", "")
	userConfig, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	files, _ := filepath.Glob(
		filepath.Join(userConfig, "JetBrains", "*", "options", "jdk.table.xml"),
	)

	for _, file := range files {
		doc := etree.NewDocument()
		if doc.ReadFromFile(file) != nil {
			continue
		}

		for _, jdkTag := range doc.FindElements("//jdk") {
			if jdkTag.SelectElement("name").SelectAttrValue("value", "") == jdk {
				helper.Vprintf("idea: using project sdk %q\n", jdk)
				return jdkTag.SelectElement("homePath").SelectAttrValue("value", ""), nil
			}
		}
	}

	fmt.Printf("flatmod: using home jdk instead of %q\n", jdk)
	return "", nil
}
