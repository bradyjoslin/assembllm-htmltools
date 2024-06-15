build:
	tinygo build -target wasi -o assembllm-htmltools.wasm main.go

test:
	extism call assembllm-htmltools.wasm scraper --input='{"html": "<ul> <li>foo</li> <li>bar</li> </ul><p class='\''moon'\''>test text</p>", "selector": ".moon"}' --wasi
	extism call assembllm-htmltools.wasm htmlrewrite --input='{"html": "<html><body><h1>Title</h1><p>This is a paragraph.</p><div>Some <span>nested</span> text.</div></body></html>", "rules": [{"selector": "p", "html_content": "This is the new paragraph content."}, {"selector": "div", "html_content": "<b>New nested content</b>"}]}' --wasi