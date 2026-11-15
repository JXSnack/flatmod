package idea

import (
	"fmt"

	"github.com/JXSnack/flatmod/dependency"
	"github.com/beevik/etree"
)

// makeBaseLib creates a basic library XML and returns the base document and CLASSES.
func makeBaseLib(name string) (*etree.Document, *etree.Element) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// component
	component := doc.CreateElement("component")
	component.CreateAttr("name", "libraryTable")

	library := component.CreateElement("library")
	library.CreateAttr("name", name)

	// classes
	classes := library.CreateElement("CLASSES")
	return doc, classes
}

func MakeLibrary(name string, nested []dependency.Dependency, facet string, jarSuffix string) *etree.Document {
	doc, classes := makeBaseLib(name)

	appendRootClass(classes, fmt.Sprintf("jar://$USER_HOME$/.flatmod/cache/libraries/%s%s.jar!/", name, jarSuffix))
	for _, nest := range nested {
		appendRootClass(classes, fmt.Sprintf("jar://$USER_HOME$/.flatmod/cache/libraries/%s%s.jar!/", nest.Filename(facet), jarSuffix))
	}

	return doc
}

func MakeMinecraft(version string, libs []string, jarSuffix string) *etree.Document {
	doc, classes := makeBaseLib("minecraft")

	appendRootClass(classes, fmt.Sprintf("jar://$USER_HOME$/.flatmod/cache/minecraft/%s/client%s.jar!/", version, jarSuffix))
	appendRootClass(classes, fmt.Sprintf("jar://$USER_HOME$/.flatmod/cache/minecraft/%s/server%s.jar!/", version, jarSuffix))
	for _, lib := range libs {
		appendRootClass(classes, fmt.Sprintf("jar://$USER_HOME$/.flatmod/cache/libraries/%s!/", lib))
	}

	return doc
}

func appendRootClass(parent *etree.Element, url string) {
	root := parent.CreateElement("root")
	root.CreateAttr("url", url)
}
