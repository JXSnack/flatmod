package idea

import (
	"fmt"

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

func MakeDependencies(name string, dependencies []string) *etree.Document {
	doc, classes := makeBaseLib(name + "-dependencies")

	for _, dep := range dependencies {
		appendRootClass(classes, fmt.Sprintf("jar://%s!/", dep))
	}

	return doc
}

func MakeMinecraft(name string, minecraft string, libraries []string) *etree.Document {
	doc, classes := makeBaseLib(name + "-minecraft")

	appendRootClass(classes, fmt.Sprintf("jar://%s!/", minecraft))
	for _, lib := range libraries {
		appendRootClass(classes, fmt.Sprintf("jar://%s!/", lib))
	}

	return doc
}

func appendRootClass(parent *etree.Element, url string) {
	root := parent.CreateElement("root")
	root.CreateAttr("url", url)
}
