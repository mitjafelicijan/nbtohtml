# Newsboat converter

Converts unread [Newsboat](https://newsboat.org/) items to HTML for
reading outside of a terminal. You will need [DuckDB](https://duckdb.org/)
to query Newsboat's SQLite3 cache database.

You can use `nix-shell shell.nix` that will provide Golang and DuckDB.

**Example output**: Check how the example output looks like at
https://mitjafelicijan.github.io/nbtohtml/.

## Build and run

To build the software you do `go build .`. This will produce `nbtohtml`
binary.

The procedure to convert:

1. use Duckdb to get unread articles
2. pipe to `nbtohtml`
3. save to out HTML file

```sh
duckdb -json ~/.newsboat/cache.db "SELECT * FROM rss_item WHERE unread=1;" | nbtohtml > out.html
```

## Custom template file

`nbtohtml` allows you to provide your own template file with a flag
`-template`. Usage would then be `nbtohtml -template mytemplate.html`.

You can use the template below as a boilerplate. This template is included
with the application.

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Newsboat</title>

    <style>
      * {
        white-space: normal;
      }
    
      body {
        font-family: sans-serif;
        padding: 2em;
        white-space: normal;
      }
      
      .nb-article-list {
        display: flex;
        flex-direction: column;
        gap: 2em;
      }

      .nb-article-list * {
        font-size: inital!important;
      }
    </style>
  </head>
  <body>

    <section class="nb-article-list">
    {{ range .Items }}
      <article>
        <h1>
          <a href="{{ .Url }}" target="_blank">{{ .Title }}</a>
        </h1>
        <div style="text-decoration:underline; margin-bottom:1em;">
          {{ if ne .Author "" }}By {{ .Author }} on{{ end }}
          {{ .PubDate }}
        </div>
        <div style="margin-left:4em;">
          {{ .Content }}
        </div>
      </article>
    {{ end }}
    </section>
    
  </body>
</html>
```

## License

[nbtohtml](https://github.com/mitjafelicijan/nbtohtml) was written by
[Mitja Felicijan](https://mitjafelicijan.com) and is released under the
BSD two-clause license, see the LICENSE file for more information.
