package tritium

import (
  "time"
)

import(
  "open-tritium/whale"
  tp "open-tritium/proto"
  "open-tritium/packager"
  "open-tritium/dependencies/golog"
  "open-tritium/dependencies/steno/dummy"
  "open-tritium/linker"
)

var pkg *tp.Package

func Transform(tscript string, input string, importpath string) string {
  logger := golog.NewLogger("tritium")
  logger.AddProcessor("info", golog.NewConsoleProcessor(golog.LOG_INFO, true))

  pkgr := packager.New_OSS(importpath, logger, func(name, version string) (mxr *tp.Mixer, err error) {return nil, nil})
  pkgr.Build_OSS(lib, types)
  pkg = pkgr.Mixer.Package
  script, _ := linker.RunWithPackage_OSS(tscript, importpath, pkg, make([]string, 0))

  // input := readFile(inputfile)

  debugger := &dummy.DummyDebugger{}
  eng := whale.NewEngine(debugger)
  d, _ := time.ParseDuration("10m")
  exh := eng.Run(script, nil, input, make(map[string]string, 0), time.Now().Add(d), "test", "test", "test", make([]string, 0), false)
  // os.Stderr = os.Stdout
  // fmt.Fprintf(os.Stderr, "%s", exh.Output)
  return exh.Output
}