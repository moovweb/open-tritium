# Open Tritium's GitHub Page

Found at [open-tritium.io](http://open-tritium.io).

## Getting started 

Install [Jekyll](http://jekyllrb.com/) 

	gem install jekyll

This will install Jekyll 2.x, which converts Sass, but since GitHub Pages is still on 1.5.x, we need to manually update the CSS:

	sass --watch stylesheets/scss:stylesheets --style compact

In a separate terminal process, run the server (this will also listen for changes)

	jekyll serve --watch --baseurl=''

Browse it at [localhost:4000](http://localhost:4000).

That's it!

###TODO

Update to Jekyll 2.2.0

