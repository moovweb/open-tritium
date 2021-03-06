package linker

import (
	// "errors"
	"log"

	parser "open-tritium/parser"
	tp "open-tritium/proto"
	. "open-tritium/util"
)

func RunStringWithPackage(src, projectPath, scriptPath, fileName string, pkg *tp.Package, activeLayers []string, ranges ...Range) (*tp.Transform, error) {
	objs := parser.Parse(src, projectPath, scriptPath, fileName, false, activeLayers)
	return runWithObjs(objs, pkg, projectPath, scriptPath, ranges...)
}

func RunWithPackage(projectPath, scriptPath, fileName string, pkg *tp.Package, activeLayers []string) (*tp.Transform, error) {
	objs := parser.ParseFileSet(projectPath, scriptPath, fileName, false, activeLayers)
	return runWithObjs(objs, pkg, projectPath, scriptPath)
}

func RunWithPackage_OSS(src string, scriptPath string, pkg *tp.Package, activeLayers []string) (*tp.Transform, error) {
	objs := parser.Parse(src, "", scriptPath, "", false, activeLayers)
	return runWithObjs(objs, pkg, "", scriptPath)
}

func runWithObjs(objs []*tp.ScriptObject, pkg *tp.Package, projectPath, scriptPath string, ranges ...Range) (*tp.Transform, error) {
	ctx := NewObjectLinkingContext(pkg, objs, projectPath, scriptPath, ranges...)
	ctx.Link()
	if ctx.HasErrors() {
		message := ""
		for _, msg := range ctx.Errors {
			message = message + "\n" + msg
		}
		log.Fatal(message)
		// return nil, errors.New(message)
	}

	return ctx.Transform, nil
}
