package idea

import (
	"fmt"

	"github.com/JXSnack/flatmod/dependency"
	"github.com/beevik/etree"
)

func MakeLibrary(name string, nested []dependency.Dependency, jarSuffix string) *etree.Document {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// component
	component := doc.CreateElement("component")
	component.CreateAttr("name", "libraryTable")

	library := component.CreateElement("library")
	library.CreateAttr("name", name)

	// classes
	classes := library.CreateElement("CLASSES")

	appendRootClass(classes, fmt.Sprintf("jar://$USER_HOME$/.flatmod/cache/libraries/%s%s.jar!/", name, jarSuffix))
	for _, nest := range nested {
		appendRootClass(classes, fmt.Sprintf("jar://$USER_HOME$/.flatmod/cache/libraries/%s%s.jar!/", nest.Filename(), jarSuffix))
	}

	return doc
}

func appendRootClass(parent *etree.Element, url string) {
	root := parent.CreateElement("root")
	root.CreateAttr("url", url)
}
