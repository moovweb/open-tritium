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

