tritium_oss
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

  <html>
    <body>
      <div id="awesome" />
    </body>
  </html>

If you pass that in as an input to the following script.

  $("/html/body/div[@id='awesome']") {
    attribute("class", "even_awesomer!")
    wrap("div", class: "wrapper")
  }

You'd get back...

  <html>
    <body>
      <div class="wrapper">
        <div id="awesome" class="even_awesomer!"/>
      </div>
    </body>
  </html>


## Building Tritium Locally

* This assumes you have Go installed (64-bit) and have setup your $GOPATH
ex. `export GOPATH=/Users/Yourname/dev`

### Setup MOOV_HOME

Set $MOOV_HOME var to be your GOPATH:
`export MOOV_HOME=$GOPATH`

### Fetch the dependencies

`go get code.google.com/p/goprotobuf/proto`

Note: goprotobuf requires [Mercurial](http://mercurial.selenic.com/) which is required to fetch the latest version.

Fetch the moovweb repositories for gokogiri.
For current compatability, switch to the 'oss' branch.

`cd $GOPATH`

`go get github.com/moovweb/gokogiri`

`cd src/github.com/moovweb/gokogiri ; git checkout oss`

`go install`

### Build and install our required clibs.

Versions have been locked for our development processes. Clibs can be installed anywhere but the following steps assume they are placed inside a GOPATH/clibs/ folder. Source can be cloned into GOPATH/clibs/src and built. EX:

`mkdir $GOPATH/clibs; mkdir $GOPATH/clibs/src; cd $GOPATH/clibs/src`

You will need to have autoconf, automake, and libtool installed. If you dont:

`brew install autoconf`
`brew install automake`
`brew install libtool`

Note: you may wish to create an environment variable to this clibs directory, $CLIBS_HOME

`git clone git@github.com:moovweb/icu4c`

`git clone git@github.com:moovweb/libiconv`

`git clone git@github.com:moovweb/libxml2`

`git clone git@github.com:moovweb/libsass`

`git clone git@github.com:moovweb/libyaml`

`git clone git@github.com:moovweb/oniguruma`


For each of the 6 libraries, enter the folder and build them with

`./build.sh`

$GOPATH/clibs/ should now contain the built /bin /includes and /lib folders for use within the Tritium application

Fetch the moovweb repositories for rubex.
For current compatability, switch to the 'oss' branch.

`cd $GOPATH`

`go get -d github.com/moovweb/rubex`

`cd src/github.com/moovweb/rubex ; git checkout oss`

`go install`


### Clone this Repo

Clone this repo into your $GOPATH/src folder:

`cd $GOPATH/src`

`git clone git@github.com:moovweb/tritium_oss`


### Building the Driver from src (optional)

`cd tritium_oss/driver`

`go build -ldflags -extldflags=-L$MOOV_HOME/clibs/lib`

Alternatively, inside /driver, you can also call ./build.sh which will build the driver package automatically.

### Run Tritium

You should now have a compiled tritium file in your /tritium_oss/driver directory. Run by passing in the path to any tritium script and an html file to transform. Currently the driver emits the transformations to STDOUT, and can be piped into other functions.

`./tritium -f="driver/main.ts" -i="driver/input.html"`

The compiled binary can be used to transform HTML with any functions found in /mixers/tritium/lib. We have provided the primary functions at the core of the language. Why not try writing your own?





