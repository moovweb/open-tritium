
# Simply start by parsing the document as 'html'
html() {
  # Select the body
  $("/html/body") {
    # Append a class to the body tag
    add_class("moov")

    # Sometimes empty blocks break stuff
    $(".") {
    }
  }
}