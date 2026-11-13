package idea

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/dependency"
	"github.com/beevik/etree"
)

func GetIML() string {
	iml := ""
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if strings.HasPrefix(path, ".") && !strings.HasPrefix(path, ".idea") {
			return nil
		}

		if strings.HasSuffix(path, "flatmod-project.iml") {
			iml = path
		}
		return nil
	})

	return iml
}

func MakeIML(platform string, group string, dependencies []dependency.Dependency) *etree.Document {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// module
	module := doc.CreateElement("module")
	module.CreateAttr("type", "JAVA_MODULE")
	module.CreateAttr("version", "4")

	// facet manager & mcp
	appendFacetManager(module, platform)
	appendMcp(module)

	// module root, directory highlighting & libraries
	nmrm := appendNMRM(module)
	appendModuleContent(nmrm, group)
	appendLibraries(nmrm, dependencies)

	return doc
}

func appendFacetManager(module *etree.Element, platform string) {
	component := module.CreateElement("component")
	component.CreateAttr("name", "FacetManager")

	facet := component.CreateElement("facet")
	facet.CreateAttr("type", "minecraft")
	facet.CreateAttr("name", "Minecraft")

	conf := facet.CreateElement("configuration")
	autoDetectTypes := conf.CreateElement("autoDetectTypes")
	appendElement(autoDetectTypes, "platformType", "MCP")
	appendElement(autoDetectTypes, "platformType", platform)
	appendElement(autoDetectTypes, "platformType", "MIXIN")
}

func appendElement(parent *etree.Element, tag string, inner string) *etree.Element {
	e := parent.CreateElement(tag)
	e.SetText(inner)
	return e
}

func appendMcp(module *etree.Element) {
	component := module.CreateElement("component")
	component.CreateAttr("name", "McpModuleSettings")

	option := component.CreateElement("option")
	option.CreateAttr("name", "srgType")
	option.CreateAttr("value", "SRG")
}

func appendNMRM(module *etree.Element) *etree.Element {
	nmrm := module.CreateElement("component")
	nmrm.CreateAttr("name", "NewModuleRootManager")
	nmrm.CreateAttr("LANGUAGE_LEVEL", "JDK_21")

	inheritedJdk := nmrm.CreateElement("orderEntry")
	inheritedJdk.CreateAttr("type", "inheritedJdk")

	sourceFolder := nmrm.CreateElement("orderEntry")
	sourceFolder.CreateAttr("type", "sourceFolder")
	sourceFolder.CreateAttr("forTests", "false")

	return nmrm
}

func appendModuleContent(nmrm *etree.Element, group string) {
	content := nmrm.CreateElement("content")
	content.CreateAttr("url", "file://$MODULE_DIR$")

	pkg := content.CreateElement("sourceFolder")
	pkg.CreateAttr("url", "file://$MODULE_DIR$/")
	pkg.CreateAttr("isTestSource", "false")
	pkg.CreateAttr("packagePrefix", group)

	resources := content.CreateElement("sourceFolder")
	resources.CreateAttr("url", "file://$MODULE_DIR$/resources")
	resources.CreateAttr("type", "java-resource")

	dotIdea := content.CreateElement("excludeFolder")
	dotIdea.CreateAttr("url", "file://$MODULE_DIR$/.idea")

	dotFlatmod := content.CreateElement("excludeFolder")
	dotFlatmod.CreateAttr("url", "file://$MODULE_DIR$/.flatmod")
}

func appendLibraries(nmrm *etree.Element, libraries []dependency.Dependency) {
	for _, lib := range libraries {
		e := nmrm.CreateElement("orderEntry")
		e.CreateAttr("type", "library")
		e.CreateAttr("name", lib.Filename())
		e.CreateAttr("level", "project")
	}
}
