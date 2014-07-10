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
1. Setting up local vars:

* This assumes you have setup your $GOPATH
ex. `export GOPATH=/Users/Yourname/dev`

Setup $MOOV_HOME var to be your GOPATH:
`export MOOV_HOME=$GOPATH`

2. Fetch the dependencies

Ensure the dependent clibs are inside a /clibs folder within GOPATH

Fetching dependencies, for now check out the oss branch of each public repository:

`cd $GOPATH`
`go get github.com/moovweb/gokogiri`
`cd github.com/moovweb/gokogiri ; git checkout oss`

`cd $GOPATH`
`go get github.com/moovweb/rubex`
`cd github.com/moovweb/rubex ; git checkout oss`

3. Installing Tritium

Clone the tritium_oss repo into your GOPATH/src folder:

`cd $GOPATH`
`git clone git@github.com:moovweb/tritium_oss`

4. Building the Driver from src

`cd tritium_oss/trit`
`go build -ldflags -extldflags=-L$MOOV_HOME/clibs/lib

5. Run Tritium

You should now have a trit file in your /tritium_oss/trit directory.

Run by passing in the path to any tritium script and an html file to transform.

`./trit main.ts input.html`



