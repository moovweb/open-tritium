package main

var base = `

@func read$$$$$$$$$$$$$$$$$$$$(Text %file) {
  // [native function]
}

@func export(Text %key, Text %value) {
  export(%key) {
    set(%value)
    yield()
  }
}

@func regexp(Text %exp) {
  regexp(%exp, "")
}

@func asset(Text %name) {
  concat($asset_host, %name) {
    yield()
  }
}

@func bm(Text %name) {
  log(concat(%name, ": ",
    time() {
      yield()
    }
  ))
}

@func match(Text %target, Text %comparitor) {
  match(%target) {
    with(%comparitor) {
      yield()
    }
  }
}

@func match(Text %target, Regexp %comparitor) {
  match(%target) {
    with(%comparitor) {
      yield()
    }
  }
}

@func match_not(Text %target, Text %comparitor) {
  match(%target) {
    not(%comparitor) {
      yield()
    }
  }
}

@func match_not(Text %target, Regexp %comparitor) {
  match(%target) {
    not(%comparitor) {
      yield()
    }
  }
}

@func Text.clear() {
  set("") {
    yield()
  }
}

@func Text.text() {
  this() {
    yield()
  }
}

@func Text.replace(Regexp %search, Text %with) {
  replace(%search) {
    set(%with)
    yield()
  }
}

@func Text.replace(Text %search, Text %with) {
  replace(%search) {
    set(%with)
  }
}

@func Text.length() {
  $input = this()
  length($input)
}

@func Text.fetch(Text %text) {
  deprecated("You can only use fetch() in a Node scope")
  ""
}

# DEPRECATED
@func match(Regexp %regexp, Text %against) {
  deprecated("Please use match(Text, Regexp). Reverse them!")
  match(%against, %regexp) {
    yield()
  }
}
@func Text.dump() {
  this()
}

@func encode64(Text %str) {
  base64_v1("encode", %str) {
    yield()
  }
}

@func decode64(Text %str) {
  base64_v1("decode", %str) {
    yield()
  }
}

`

var base_primitives = `
@func var(Text %name) Text Text

@func var(Text %name, Text %value) Text Text

@func time() Text

@func match(Text %match_target) Text

@func log(Text %log_message) Text Text

@func deprecated(Text %message) Text Text

@func not(Text %text) Text

@func not(Regexp %regexp) Text

@func with(Text %text) Text

@func with(Regexp %regexp) Text

@func Text.convert_encoding(Text %from, Text %to) Text

@func Text.guess_encoding() Text

@func length(Text %input) Text

@func else() Text

@func yield() Text

@func Text.this() Text

@func regexp(Text %expression, Text %options) Regexp Text

@func concat(Text %a, Text %b) Text Text

@func export(Text %key_name) Text Text

@func upcase(Text %input_string) Text

@func downcase(Text %input_string) Text

@func Text.set(Text %value) Text

@func Text.replace(Regexp %search) Text Text

@func Text.replace(Text %search) Text Text

@func Text.prepend(Text %text_to_prepend) Text

@func Text.append(Text %text_to_append) Text

@func Text.capture(Regexp %search) Text Text

@func base64_v1(Text %method, Text %str) Text Text

@func Base.env(Text %key) Text

`

var libxml = `
@func XMLNode.$$(Text %css_selector) {
  $(css(%css_selector)) {
    yield()
  }
}

@func XMLNode.add_class(Text %class) {
  attribute("class") {
    value() {
      append(" ")
      append(%class)
    }
    yield()
  }
}

@func XMLNode.inner_wrap(Text %tag_name) {

  insert_top(%tag_name) {
    %wrapper = this()
    $("..") {
      move_children_to(%wrapper, position("bottom"))
    }
    yield()
  }
}

@func XMLNode.remove_text_nodes() {
  remove("./text()")
}

@func XMLNode.attribute(Text %name, Text %value) {
  attribute(%name) {
    value() {
      set(%value)
    }
    yield()
  }
}

@func Attribute.value(Text %value) {
  value() {
    set(%value)
    yield()
  }
}

@func Attribute.name(Text %name) {
  name() {
    set(%name)
    yield()
  }
}

@func sass(Text %filename) {
  asset("stylesheets/.css/" + %filename + ".css") {
    yield()
  }
}

@func XMLNode.set(Text %name, Text %value) {
  attribute(%name) {
    value(%value)
  }
}

@func XMLNode.attributes() {
  yield()
}

@func XMLNode.text() {
  inner_text() {
    yield()
  }
}

@func XMLNode.text(Text %value) {
  text() {
    set(%value)
    yield()
  }
}

@func XMLNode.absolutize(Text %xpath, Text %attribute) {

  # Absolutize IMG and SCRIPT SRCs
  var("slash_path") {
    # the 'slash_path' is the path of this page without anything following it's last slash
    set($path)
    replace(/[^\/]+$/, "")
    # turn empty string into a single slash because this is the only thing separating the host from the path relative path
    replace(/^$/, "/")
  }

  $(%xpath) {
    var("url", fetch(concat("./@", %attribute)))

    # skip URLs which: are empty, have a host (//www.example.com), or have a protocol (http:// or mailto:)
    match($url, /^(?![a-z]+\:)(?!\/\/)(?!$)/) {
      attribute(%attribute) {
        value() {
          match($url) {
            with(/^\//) {
              # host-relative URL: just add the host
              prepend($source_host)
              prepend("//")
            }
            else() {
              # TODO: I need a test case for this clause. I'm not sure what kind of path its trying to accomodate
              # path-relative URL: add the host and the path
              prepend($slash_path)
              prepend($source_host)
              prepend("//")
            }
          }
        }
      }
    }
    yield()
  }
}

@func XMLNode.absolutize(Text %xpath) {
  absolutize(%xpath, "src") {
    yield()
  }
}

@func XMLNode.absolutize() {
  absolutize(".//img|.//script") {
    yield()
  }
}

@func XMLNode.insert_javascript_at(Position %pos, Text %js) {
  insert_at(%pos, "script") {
    attribute("type", "text/javascript")
    cdata(concat("//<![CDATA[\n", %js, "\n//]]>"))
    yield()
  }
}

@func XMLNode.insert_javascript_at(Text %pos, Text %js) {
  insert_at(position(%pos), "script") {
    attribute("type", "text/javascript")
    cdata(concat("//<![CDATA[\n", %js, "\n//]]>"))
    yield()
  }
}

@func XMLNode.insert_javascript(Text %js) {
  insert_javascript_at(position(), %js) {
    yield()
  }
}

@func XMLNode.inner(Text %html) {
  inner() {
    set(%html)
    yield()
  }
}

@func XMLNode.wrap(Text %tag) {
  %parent_node = this()
  insert_at(position("before"), %tag) {
    move(%parent_node, this(), position("top"))
    yield()
  }
}

@func XMLNode.wrap(Text %tag) {
  %node = this()
  insert_at(position("before"), %tag) {
    move(%node, this(), position("top"))
    yield()
  }
}

@func Text.html(Text %from_enc, Text %to_enc) {
  html_doc(%from_enc, %to_enc) {
    yield()
  }
  export("Content-Type-Charset", %to_enc)
  match_not($__preservedSectionOfText__, "") {
    set($__preservedSectionOfText__)
  }
}

@func XMLNode.keep_only_this() {
  $__preservedSectionOfText__ = fetch(path())
}

@func Text.html(Text %enc) {
  $charset_determined = %enc
  html(%enc, %enc) {
    yield()
  }
}

@func Text.html() {
  match($charset_determined) {
    with("") {
      $encoding = guess_encoding()
    }
    else() {
      $encoding = $charset_determined
    }
  }
  html($encoding, $encoding) {
    yield()
  }
}

@func Text.html_fragment(Text %from_enc, Text %to_enc) {
  html_fragment_doc(%from_enc, %to_enc) {
    yield()
  }
  export("Content-Type-Charset", %to_enc)   # Right now we always output in utf-8, so set the response header appropriately
}

@func Text.html_fragment(Text %enc) {
  $charset_determined = %enc
  html_fragment(%enc, %enc) {
    yield()
  }
}

@func Text.html_fragment() {
  $encoding = guess_encoding()
  html_fragment($encoding, $encoding) {
    yield()
  }
}

@func XMLNode.insert_at(Position %pos, Text %tag, Text %inner) {
  insert_at(%pos, %tag) {
    inner(%inner)
    yield()
  }
}

@func XMLNode.insert_at(Text %pos, Text %tag, Text %content) {
  insert_at(position(%pos), %tag) {
    inner(%content)
    yield()
  }
}

@func XMLNode.insert(Text %tag, Text %inner) {
  insert_at(position("bottom"), %tag) {
    inner(%inner)
    yield()
  }
}

@func XMLNode.insert_bottom(Text %tag, Text %inner) {
  insert_at(position("bottom"), %tag) {
    inner(%inner)
    yield()
  }
}

@func XMLNode.insert_top(Text %tag, Text %inner) {
  insert_at(position("top"), %tag) {
    inner(%inner)
    yield()
  }
}

@func XMLNode.insert_after(Text %tag, Text %inner) {
  insert_at(position("after"), %tag) {
    inner(%inner)
    yield()
  }
}

@func XMLNode.insert_before(Text %tag, Text %inner) {
  insert_at(position("before"), %tag) {
    inner(%inner)
    yield()
  }
}

@func XMLNode.insert_javascript_bottom(Text %js) {
  insert_javascript_at(position("bottom"), %js) {
    yield()
  }
}

@func XMLNode.insert_javascript_top(Text %js) {
  insert_javascript_at(position("top"), %js) {
    yield()
  }
}

@func XMLNode.insert_javascript_after(Text %js) {
  insert_javascript_at(position("after"), %js) {
    yield()
  }
}

@func XMLNode.insert_javascript_before(Text %js) {
  insert_javascript_at(position("before"), %js) {
    yield()
  }
}

@func XMLNode.add_javascript_asset(Text %filename) {
  insert_bottom("script", data-keep: "true", type: "text/javascript", src: asset("javascript/" + %filename)) {
    yield()
  }
}

@func XMLNode.add_sass_asset(Text %filename) {
  insert_bottom("link", rel: "stylesheet", type: "text/css", data-mw-keep: "true", href: sass(%filename)) {
    yield()
  }
}

@func XMLNode.add_image_asset(Text %filename) {
  insert_bottom("img", src: asset("images/" + %filename)) {
    yield()
  }
}

@func asset(Text %name, Text %type) {
  deprecated("Please use asset('path/to/asset.jpg')")
  match(%type) {
    with("js") {
      $_deprecated_assets_tmp = asset(concat("javascript/", %name))
    }
    with("image") {
      $_deprecated_assets_tmp = asset(concat("images/", %name))
    }
    with("stylesheet") {
      $_deprecated_assets_tmp = asset(concat("stylesheets/.css/", %name))
    }
  }
  $_deprecated_assets_tmp
}

`

var libxml_primitives = `
@func css(Text %selector) Text Text

@func Text.xml() Text XMLNode

@func Text.html_doc(Text %input_encoding, Text %output_encoding) Text XMLNode

@func Text.html_fragment_doc(Text %input_encoding, Text %output_encoding) Text XMLNode

@func XMLNode.cdata(Text %contents) Text

@func XMLNode.remove(Text %xpath_selector) Text

@func XMLNode.inner() Text Text

@func XMLNode.inner_text() Text Text

@func XMLNode.attribute(Text %name) Text Attribute

@func equal(XMLNode %a, XMLNode %b) Text

@func equal(Node %a, Node %b) Text

@func XMLNode.wrap_text_children(Text %tag_name) Text XMLNode

@func XMLNode.move_children_to(Node %tag_name, Position %pos) Text

@func Attribute.remove() Text

@func Attribute.value() Text Text

@func Attribute.name() Text Text

`

var node = `
@func Node.$(Text %xpath) {
  select(%xpath) {
    yield()
  }
}

@func position() {
  position("bottom")
}

@func Node.node() {
  this() {
    yield()
  }
}

@func Node.index() {
  index(this()) {
    yield()
  }
}

@func Node.name(Text %value) {
  name() {
    set(%value)
    yield()
  }
}

@func Node.copy_here(Text %xpath, Position %pos) {
  %calling_node = this()
  $(%xpath) {
    dup() {
      move(this(), %calling_node, %pos)
      yield()
    }
  }
}

@func Node.copy_here(Text %xpath, Text %pos) {
  copy_here(%xpath, position(%pos)) {
    yield()
  }
}

@func Node.copy_here(Text %xpath) {
  copy_here(%xpath, position()) {
    yield()
  }
}


@func Node.copy_to(Text %xpath, Position %pos) {
  %source = this()
  $(%xpath) {
    %destination = this()
    %source {
      dup() {
        %copied_node = this()
        move(%copied_node, %destination, %pos)
        yield()
      }
    }
  }
}

@func Node.copy_to(Text %xpath, Text %pos) {
  copy_to(%xpath, position(%pos)) {
    yield()
  }
}

@func Node.copy_to(Text %xpath) {
  copy_to(%xpath, position()) {
    yield()
  }
}

@func Node.inject(Text %html) {
  inject_at(position("bottom"), %html) {
    yield()
  }
}

@func Node.move_to(Text %xpath, Position %pos) {
  %parent_node = this()
  $(%xpath) {
    move(%parent_node, this(), %pos)
    yield()
  }
}

@func Node.move_to(Text %xpath, Text %pos) {
  move_to(%xpath, position(%pos)) {
    yield()
  }
}

@func Node.inject_at(Position %pos, Text %html) {
  inject_at_v1(%pos, %html) {
    yield()
  }
}

@func Node.inject_at(Text %pos, Text %html) {
  inject_at(position(%pos), %html) {
    yield()
  }
}

@func Node.move_to(Text %xpath) {
  move_to(%xpath, position()) {
    yield()
  }
}

@func Node.move_here(Text %where, Position %pos) {
  %parent = this()
  select(%where) {
    move(this(), %parent, %pos)
    yield()
  }
}

@func Node.move_here(Text %where, Text %pos) {
  move_here(%where, position(%pos)) {
    yield()
  }
}

@func Node.move_here(Text %where) {
  move_here(%where, position("bottom")) {
    yield()
  }
}

@func Node.insert(Text %tag) {
  insert_at(position(), %tag) {
    yield()
  }
}

@func Node.insert_bottom(Text %tag) {
  insert_at(position(), %tag) {
    yield()
  }
}

@func Node.insert_top(Text %tag) {
  insert_at(position("top"), %tag) {
    yield()
  }
}

@func Node.insert_after(Text %tag) {
  insert_at(position("after"), %tag) {
    yield()
  }
}

@func Node.insert_before(Text %tag) {
  insert_at(position("before"), %tag) {
    yield()
  }
}

@func Node.inject(Text %html) {
  inject_at(position("bottom"), %html) {
    yield()
  }
}

@func Node.inject_bottom(Text %html) {
  inject_at(position("bottom"), %html) {
    yield()
  }
}

@func Node.inject_top(Text %html) {
  inject_at(position("top"), %html) {
    yield()
  }
}

@func Node.inject_after(Text %html) {
  inject_at(position("after"), %html) {
    yield()
  }
}

@func Node.inject_before(Text %html) {
  inject_at(position("before"), %html) {
    yield()
  }
}

`

var node_primitives = `
@func index(Node %node) Text

@func Node.fetch(Text %selector) Text Text

@func Node.this() Node Node

@func position(Text %position) Position

@func Node.select(Text %xpath_selector) Text Node

@func Node.text() Text Text

@func Node.move(Node %what, Node %where, Position %pos) Text Node

@func Node.dup() Node Node

@func Node.name() Text Text

@func Node.remove() Text Node

@func Node.path() Text Text

@func Node.insert_at(Position %pos, Text %tag_name) Text

@func Node.inject_at(Position %pos, Text %html) Text

@func Node.inject_at_v1(Position %pos, Text %html) Text

@func Node.set(Text %value) Text

`

var stdlib_impl = `
@func XMLNode.wrap_together(Text %selector, Text %tag) {
  $("(" + %selector + ")[1]") {
    %first = this()
    $("./..") {
      $tmp = $("(" + %selector + ")[position() > 1]") {
        move(this(), %first, position("after"))
      }
    }

    wrap(%tag) {
      $("./following-sibling::*[position() <= " + $tmp + "]") {
        move(this(), %first, position("after"))
      }
      yield()
    }
  }
}

@func add_cookie(Text %raw) {
  match_not(%raw, "") {
    $__cookie__ = concat($__cookie__, "\r\n")
    $__cookie__ = concat($__cookie__, "Set-Cookie: ")
    $__cookie__ = concat($__cookie__, %raw)
  }
}

@func add_cookie(Text %name, Text %value) {
  match_not(%name, "") {
    $raw = concat(%name, "=")
    $raw = concat($raw, %value)
    add_cookie($raw)
  }
}

@func add_cookie(Text %name, Text %value, Text %domain) {
  match_not(%name, "") {
    $raw = concat(%name, "=")
    $raw = concat($raw, %value)
    $raw = concat($raw, "; domain=")
    $raw = concat($raw, %domain)
    add_cookie($raw)
  }
}

@func add_cookie(Text %name, Text %value, Text %domain, Text %path) {
  match_not(%name, "") {
    $raw = concat(%name, "=")
    $raw = concat($raw, %value)
    $raw = concat($raw, "; domain=")
    $raw = concat($raw, %domain)
    $raw = concat($raw, "; path=")
    $raw = concat($raw, %path)
    add_cookie($raw)
  }
}

@func add_cookie(Text %name, Text %value, Text %domain, Text %path, Text %expire, Text %secure, Text %httponly) {
  match_not(%name, "") {
    $raw = concat(%name, "=")
    $raw = concat($raw, %value)
    $raw = concat($raw, "; domain=")
    $raw = concat($raw, %domain)
    $raw = concat($raw, "; path=")
    $raw = concat($raw, %path)
    $raw = concat($raw, "; expires=")
    $raw = concat($raw, %expire)
    match(%secure, "true") {
      $raw = concat($raw, "; secure")
    }
    match(%httponly, "true") {
      $raw = concat($raw, "; httponly")
    }
    add_cookie($raw)
  }
}

@func XMLNode.yank(Text %xpath) {
  $innards = ""
  $(%xpath) {
    $innards = inner()
  }
  $innards
}

@func XMLNode.unwrap() {
  move_children_to(this(), position("before"))
  remove()
}

@func XMLNode.attr(Text %name) {
  attribute(%name) {
    yield()
  }
}

@func XMLNode.attr(Text %name, Text %value) {
  attribute(%name, %value) {
    yield()
  }
}

@func normalize(Text %input) {
  %input {
    replace(/\s\s+/, " ")
    replace(/^\s+|\s+$/, "")
  }
  %input
}

@func XMLNode.remove_class(Text %delete_me) {
  attribute("class") {
    value() {
      replace(regexp("\\b" + %delete_me + "\\b"), "")
      set(normalize(this()))
    }
  }
}

@func XMLNode.add_class(Text %add_me) {
  attribute("class") {
    value() {
      append(" ")
      append(%add_me)
      // Exactly like the native function,
      // except that we normalize.
      set(normalize(this()))
    }
    yield()
  }
}

@func XMLNode.insert_at(Text %pos, Text %tag) {
  insert_at(position(%pos), %tag) {
    inner("")
    yield()
  }
}

@func redirect_temporary(Text %url) {
  $__redirect_permanent__ = "false"
  export("Location", %url)
}

@func redirect_permanent(Text %url) {
  $__redirect_permanent__ = "true"
  export("Location", %url)
}

@func XMLNode.move(Text %what, Text %where, Text %pos) {
  %current_node = this()
  $(%where) {
    %where_node = this()
    %current_node {
      $(%what) {
        move(this(), %where_node, position(%pos)) {
          yield()
        }
      }
    }
  }
}

@func XMLNode.move(Text %what, Text %where) {
  move(%what, %where, "bottom") {
    yield()
  }
}

@func Text.trim() {
  replace(/^\s+|\s+$/, "")
}

@func XMLNode.remove_attributes() {
  remove('./@*')
}

@func XMLNode.flatten() {
  inner(text())
}

@func XMLNode.filter(Text %xpath_selector) {
  $(%xpath_selector) {
    attribute('data-blessed', 'save')
  }
  remove('.//*[not(descendant-or-self::*[@data-blessed])]')
  $(%xpath_selector) {
    attribute('data-blessed','')
  }
}

@func Text.reverse() {
  %rev = ""
  capture(/(.)/){
    %rev {
      prepend($1)
    }
  }
  set(%rev)
  yield()
  %rev
}

@func reverse(Text %str) {
  %str {
    reverse()
    yield()
  }
}

`

var util = `
@func XMLNode.tag_count(Text %tag) {
  %num = ""
  $("//"+%tag) {
    %num {
      append("1")
    }
  }
  %num {
    set(length(%num))
  }
  %num
}

@func XMLNode.proxied_tag_count(Text %tag) {
  %num = ""
  $("//"+%tag) {
    match(fetch("@src"), /^\/[^\/]/) {
      match_not(fetch("@src"), /moovweb\_local\_asset/) {
        %num {
          append("1")
        }
      }
    }
  }
  %num {
    set(length(%num))
  }
  %num
}

@func XMLNode.print_asset_count() {
  %imgs_count = tag_count("img")
  %scripts_count = tag_count("script[@src]")

  %proxied_imgs_count = proxied_tag_count("img")
  %proxied_scripts_count = proxied_tag_count("script[@src]")

  match_not(%proxied_imgs_count, "0") {
    %proxied_imgs_count {
      append(". Use absolute URLs so images are not proxied.")
    }
  }

  log("### Asset Count ###")
  log("Total Number of Images: "+%imgs_count)
  log("Total Number of Scripts: "+%scripts_count)
  log("Number of Proxied Images: "+%proxied_imgs_count)
  log("Number of Proxied Scripts: "+%proxied_scripts_count)
}
`

var lib =  base_primitives +
           base +
           node_primitives +
           node +
           libxml_primitives +
           libxml +
           stdlib_impl +
           util

var types = `
- Base
- Text
- Regexp
- Header
- Node
- Position
- XMLNode < Node
- Attribute
`
