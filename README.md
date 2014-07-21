open-tritium
===========

Open source implementation of the Tritium language.

== Tritium ==

Tritium is a magical language that functions as a document modifying language.
It's JavaScript-like (except, without all that function() noise!) and simple to
learn.

It was designed by Hampton Catlin (@hcatlin), and has been heavily influenced
by Aaron Leung (@akhleung), the language's main engineer.

== Status ==

Tritium is proto-open source. We opened the source and changed the license.
It's still not very usable outside of a Moovweb build/run environment. If you
want to play around with the language, checkout http://tester.tritium.io

We are going to continue working hard to get this in shape for other uses outside
of Moovweb directly.

Currently, it focuses on modifying XML documents only (HTML and XHTML count too).

Check out the examples in /test/functional to get started... but basically.

 ```html
  <html>
    <body>
      <div id="awesome" />
    </body>
  </html>
  ```

If you pass that in as an input to the following script.

 ```xml
  $("/html/body/div[@id='awesome']") {
    attribute("class", "even_awesomer!")
    wrap("div", class: "wrapper")
  }
  ```

You'd get back...

 ```html
  <html>
    <body>
      <div class="wrapper">
        <div id="awesome" class="even_awesomer!"/>
      </div>
    </body>
  </html>
  ```


## Building Tritium Locally

This assumes the following:
* you have Go installed (64-bit) and have setup your $GOPATH
ex. `export GOPATH=/Users/Yourname/dev`
* This repo has been cloned into $GOPATH/src
* [Mercurial](http://mercurial.selenic.com/) is installed
* autoconf, automake, and libtool are installed. ex. `brew install autoconf; brew install automake; brew install libtool;`


When you have those requirements set up, you may run the build script in this folder:

`$GOPATH/src/open-tritium/build`

This will fetch and install all the required dependencies for tritium and set up some environment variables for you. Namely:
* code.google.com/p/goprotobuf/proto
* github.com/moovweb/gokogiri
* github.com/moovweb/rubex
* icu4c, libiconv, libxml2, libyaml, and oniguruma in $GOPATH/clibs
* Set $MOOV_HOME var to be your GOPATH


### Building the Driver from src (optional)

`cd open-tritium/driver`

`go build -ldflags -extldflags=-L$MOOV_HOME/clibs/lib`

Alternatively, inside /driver, you can also call ./build.sh which will build the driver package automatically.

### Run Tritium

You should now have a compiled tritium file in your /open-tritium/driver directory. Run by passing in the path to any tritium script and an html file to transform. Currently the driver emits the transformations to STDOUT, and can be piped into other functions.

`./tritium -f="driver/main.ts" -i="driver/input.html"`

The compiled binary can be used to transform HTML with any functions found in /mixers/tritium/lib. We have provided the primary functions at the core of the language. Why not try writing your own?


### Using Tritium in Your Code

You can also use tritium in your own go program to transform html or xml. Import "open-tritium/tr" in your program and call tritium.Transform(script, input). The Transform function takes a string of tritium code and a string of input HTML and returns a string of transformed HTML.

Before using this function, you will need to set the DYLD_LIBRARY_PATH to the clib library directory, usually $CLIBS_HOME/lib:

`export DYLD_LIBRARY_PATH=$CLIBS_HOME/lib`

Unfortunately, this environment variable may interfere with other command line interfaces such as git, so this may need to be unset after running your program with tritium:

`unset DYLD_LIBRARY_PATH`

or, set it to whatever it was before.



