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
import "io"
import "io/ioutil"
import "flag"
import "log"

var pkg *tp.Package

func readFile(filename string) string {
  f, err := ioutil.ReadFile(filename)
  if err!= nil {
    log.Fatal(err.Error())
  }
  return string(f)
}

func readFromIO(in io.Reader) string {
  bytes, err := ioutil.ReadAll(in)
  if err != nil && err != io.EOF {
    log.Fatal(err.Error())
  }
  return string(bytes)
}

func writeFile(input string) {
  byteinput := []byte(input)
  err := ioutil.WriteFile("output.html", byteinput, 0755)
  if err!= nil {
    log.Fatal(err.Error())
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

func Transform(tscript string, input string) string {
  logger := golog.NewLogger("tritium")
  logger.AddProcessor("info", golog.NewConsoleProcessor(golog.LOG_INFO, true))

  pkgr := packager.New_OSS(logger, func(name, version string) (mxr *tp.Mixer, err error) {return nil, nil})
  pkgr.Build_OSS(lib, types)
  pkg = pkgr.Mixer.Package
  script, _ := linker.RunWithPackage_OSS(tscript, pkg, make([]string, 0))

  // input := readFile(inputfile)

  debugger := &dummy.DummyDebugger{}
  eng := whale.NewEngine(debugger)
  d, _ := time.ParseDuration("10m")
  exh := eng.Run(script, nil, input, make(map[string]string, 0), time.Now().Add(d), "test", "test", "test", make([]string, 0), false)
  // os.Stderr = os.Stdout
  // fmt.Fprintf(os.Stderr, "%s", exh.Output)
  return exh.Output
}

func show_usage() {
  fmt.Println("General purpose Tritium command line interface. Language for html transformation.")
  fmt.Println("\tUsage: \n\t\t 1) tritium [-e|-f]=tritium_script \n\t\t\t Input assumed to be provided as Stdin")
  fmt.Println("\t\t 2) tritium [-e|-f]=\"tritium_script -i=\"input_file\" \n\t\t\t Input provided as a filepath")
  fmt.Println("\tOutput is streamed to Stdout")
  fmt.Println("\tFlags:")
  fmt.Println("\t\t -e=\"\": a one-line tritium program passed as a string")
  fmt.Println("\t\t -f=\"\": a filepath to tritium script")
  fmt.Println("\t\t -i=\"\" (optional): a filepath to input html")
}

func main() {

  var input, tscript string

  var e, f, i string
  flag.StringVar(&e, "e", "", "executable tritium expression")
  flag.StringVar(&f, "f", "", "filepath to tritium script")
  flag.StringVar(&i, "i", "", "filepath to input html")

  flag.Parse()

  if len(os.Args) == 1 {
    show_usage()
    return
  }

  if len(e) > 0 {
    tscript = e
  } else if len(f) > 0 {
    tscript = readFile(f)
  } else {
    fmt.Println("please provide tritium script")
    return
  }

  if len(i) > 0 {
    input = readFile(i)
  } else {

    // read from stdin
    input = readFromIO(os.Stdin)
  }
  // fmt.Println(tscript)
  // fmt.Println(input)
  output := Transform(tscript, input)

  // os.Stderr = os.Stdout
  fmt.Fprintf(os.Stdout, "%s", output)
  // println(output)

}