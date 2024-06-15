# assembllm HTML Tools

Extism web assembly plugin written in Go that performs web scraping and HTML rewriting. The plugin exposes two main functionalities: scraper and htmlrewrite.

Built to aid in generating prompts and related results using [assembllm](https://github.com/bradyjoslin/assembllm/), but can be used by any Extism host.

## Features

- The `scraper` function allows extracting text content from HTML elements that match a given CSS selector.
- The `htmlrewrite` function modifies HTML content based on a set of rewrite rules, where each rule specifies a CSS selector and new HTML content to replace the matched elements.

## Usage

### Scraper

**Input**

The `scraper` function expects a JSON input with the following structure:

- `html`: The HTML content as a string.
- `selector`: A CSS selector to identify the elements to extract text from.

**Output**

The function outputs the text content of the matched elements.

Example:

```sh
extism call \
    assembllm-htmltools.wasm scraper \
    --input='{"html": "<ul> <li>foo</li> <li>bar</li> </ul><p class='\''moon'\''>test text</p>", "selector": ".moon"}' \
    --wasi

# ==> test text
```

### HTML Rewriter

The `htmlrewrite` function expects a JSON input with the following structure:

```json
{
  "html": "<html-content>",
  "rules": [
    {
      "selector": "<css-selector>",
      "html_content": "<new-html-content>"
    },
    ...
  ]
}
```

- `html`: The HTML content as a string.
- `selector`: A CSS selector to identify the elements to extract text from.

**Output**

The function outputs the modified HTML content as a string.

Example:

```sh
extism call assembllm-htmltools.wasm htmlrewrite \
    --input='{"html": "<html><body><h1>Title</h1><p>This is a paragraph.</p><div>Some <span>nested</span> text.</div></body></html>", "rules": [{"selector": "p", "html_content": "This is the new paragraph content."}, {"selector": "div", "html_content": "<b>New nested content</b>"}]}' \
    --wasi

# ==> <html><head></head><body><h1>Title</h1><p>This is the new paragraph content.</p><div><b>New nested content</b></div></body></html>
```

## Build

```sh
make build
```

## Test

```sh
make test
```

