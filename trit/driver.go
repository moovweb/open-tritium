package main

import "path/filepath"
import "tritium_oss/whale"
import "runtime"
import "fmt"
import tp "tritium_oss/proto"
import "tritium_oss/packager"
import "tritium_oss/dependencies/golog"
import "time"
import "tritium_oss/dependencies/steno/dummy"
import "tritium_oss/linker"
import "os"
import "io/ioutil"

func readFile(filename string) string {
  f, err := ioutil.ReadFile(filename)
  if err!= nil {
    panic(err)
  }
  return string(f)
}

func writeFile(input string) {
  byteinput := []byte(input)
  err := ioutil.WriteFile("output.html", byteinput, 0755)
  if err!= nil {
    panic(err)
  }
}

func relativeDirectory(directoryFromRoot string) (directory string, ok bool) {
  _, file, _, ok := runtime.Caller(0)

  if !ok {
    return
  }

  directory = filepath.Join(file, "../../", directoryFromRoot)

  return
}

var pkg *tp.Package

func transform(tscript string, inputfile string) {
  logger := golog.NewLogger("tritium")
  logger.AddProcessor("info", golog.NewConsoleProcessor(golog.LOG_INFO, true))

  pkgr := packager.New("../mixers/tritium", "lib", false, logger, func(name, version string) (mxr *tp.Mixer, err error) {return nil, nil})
  pkgr.Build()
  pkg = pkgr.Mixer.Package
  script, _ := linker.RunWithPackage(".", ".", tscript, pkg, make([]string, 0))

  input := readFile(inputfile)

  debugger := &dummy.DummyDebugger{}
  eng := whale.NewEngine(debugger)
  d, _ := time.ParseDuration("10m")
  exh := eng.Run(script, nil, input, make(map[string]string, 0), time.Now().Add(d), "test", "test", "test", make([]string, 0), false)
  os.Stderr = os.Stdout
  fmt.Fprintf(os.Stderr, "%s", exh.Output)
}

func main() {

  transform(os.Args[1], os.Args[2])

}