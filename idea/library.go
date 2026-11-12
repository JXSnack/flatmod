package idea

import (
	"fmt"

	"github.com/beevik/etree"
)

func MakeLibrary(name string, nested []string) *etree.Document {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// component
	component := doc.CreateElement("component")
	component.CreateAttr("name", "libraryTable")

	library := component.CreateElement("library")
	library.CreateAttr("name", name)

	// classes
	classes := library.CreateElement("CLASSES")

	appendRootClass(classes, fmt.Sprintf("jar://$USER_HOME$/.flatmod/cache/libraries/%s.jar!/", name))
	for _, nest := range nested {
		appendRootClass(classes, fmt.Sprintf("jar://$USER_HOME$/.flatmod/cache/libraries/%s-nested/%s.jar!/", name, nest))
	}

	return doc
}

func appendRootClass(parent *etree.Element, url string) {
	root := parent.CreateElement("root")
	root.CreateAttr("url", url)
}
