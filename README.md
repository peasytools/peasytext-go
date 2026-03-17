# peasytext-go

[![Go Reference](https://pkg.go.dev/badge/github.com/peasytools/peasytext-go.svg)](https://pkg.go.dev/github.com/peasytools/peasytext-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/peasytools/peasytext-go)](https://goreportcard.com/report/github.com/peasytools/peasytext-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [PeasyText](https://peasytext.com) API — text case conversion, slugify, and word count. Zero dependencies beyond the Go standard library.

Built from [PeasyText](https://peasytext.com), a comprehensive text processing toolkit offering free online tools for case conversion, slug generation, word counting, and text analysis with detailed guides and glossary.

> **Try the interactive tools at [peasytext.com](https://peasytext.com)** — [Text Tools](https://peasytext.com/), [Text Glossary](https://peasytext.com/glossary/), [Text Guides](https://peasytext.com/guides/)

## Install

```bash
go get github.com/peasytools/peasytext-go
```

Requires Go 1.21+.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	peasytext "github.com/peasytools/peasytext-go"
)

func main() {
	client := peasytext.New()
	ctx := context.Background()

	// List available text tools
	tools, err := client.ListTools(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tools.Results {
		fmt.Printf("%s: %s\n", t.Name, t.Description)
	}
}
```

## API Client

The client wraps the [PeasyText REST API](https://peasytext.com/developers/) with typed Go structs and zero external dependencies.

```go
client := peasytext.New()
// Or with a custom base URL:
// client := peasytext.New(peasytext.WithBaseURL("https://custom.example.com"))
ctx := context.Background()

// List tools with pagination
tools, _ := client.ListTools(ctx, &peasytext.ListOptions{Page: 1, Limit: 10})

// Get a specific tool by slug
tool, _ := client.GetTool(ctx, "text-case")
fmt.Println(tool.Name, tool.Description)

// Search across all content
results, _ := client.Search(ctx, "slugify url", nil)
fmt.Printf("Found %d tools\n", len(results.Results.Tools))

// Browse the glossary
glossary, _ := client.ListGlossary(ctx, &peasytext.ListOptions{Search: str("case-conversion")})
for _, term := range glossary.Results {
	fmt.Printf("%s: %s\n", term.Term, term.Definition)
}

// Discover guides
guides, _ := client.ListGuides(ctx, &peasytext.ListGuidesOptions{Category: str("formatting")})
for _, g := range guides.Results {
	fmt.Printf("%s (%s)\n", g.Title, g.AudienceLevel)
}

// List file format conversions
conversions, _ := client.ListConversions(ctx, &peasytext.ListConversionsOptions{Source: str("text")})

// Get format details
format, _ := client.GetFormat(ctx, "slug")
fmt.Printf("%s (%s): %s\n", format.Name, format.Extension, format.MimeType)
```

Helper for optional string parameters:

```go
func str(s string) *string { return &s }
```

### Available Methods

| Method | Description |
|--------|-------------|
| `ListTools(ctx, opts)` | List tools (paginated, filterable) |
| `GetTool(ctx, slug)` | Get tool by slug |
| `ListCategories(ctx, opts)` | List tool categories |
| `ListFormats(ctx, opts)` | List file formats |
| `GetFormat(ctx, slug)` | Get format by slug |
| `ListConversions(ctx, opts)` | List format conversions |
| `ListGlossary(ctx, opts)` | List glossary terms |
| `GetGlossaryTerm(ctx, slug)` | Get glossary term |
| `ListGuides(ctx, opts)` | List guides |
| `GetGuide(ctx, slug)` | Get guide by slug |
| `ListUseCases(ctx, opts)` | List use cases |
| `Search(ctx, query, limit)` | Search across all content |
| `ListSites(ctx)` | List Peasy sites |
| `OpenAPISpec(ctx)` | Get OpenAPI specification |

Full API documentation at [peasytext.com/developers/](https://peasytext.com/developers/).
OpenAPI 3.1.0 spec: [peasytext.com/api/openapi.json](https://peasytext.com/api/openapi.json).

## Learn More

- **Tools**: [Text Case](https://peasytext.com/tools/text-case/) · [Text Slugify](https://peasytext.com/tools/text-slugify/) · [Text Word Count](https://peasytext.com/tools/text-word-count/) · [All Tools](https://peasytext.com/)
- **Guides**: [Case Conversion Guide](https://peasytext.com/guides/case-conversion/) · [All Guides](https://peasytext.com/guides/)
- **Glossary**: [Slug](https://peasytext.com/glossary/slug/) · [Case Conversion](https://peasytext.com/glossary/case-conversion/) · [All Terms](https://peasytext.com/glossary/)
- **Formats**: [Plain Text](https://peasytext.com/formats/plain-text/) · [Slug](https://peasytext.com/formats/slug/) · [All Formats](https://peasytext.com/formats/)
- **API**: [REST API Docs](https://peasytext.com/developers/) · [OpenAPI Spec](https://peasytext.com/api/openapi.json)

## Also Available

| Language | Package | Install |
|----------|---------|---------|
| **Python** | [peasytext](https://pypi.org/project/peasytext/) | `pip install "peasytext[all]"` |
| **TypeScript** | [peasytext](https://www.npmjs.com/package/peasytext) | `npm install peasytext` |
| **Rust** | [peasytext](https://crates.io/crates/peasytext) | `cargo add peasytext` |
| **Ruby** | [peasytext](https://rubygems.org/gems/peasytext) | `gem install peasytext` |

## Peasy Developer Tools

Part of the [Peasy Tools](https://peasytools.com) open-source developer ecosystem.

| Package | PyPI | npm | Go | Description |
|---------|------|-----|----|-------------|
| peasy-pdf | [PyPI](https://pypi.org/project/peasy-pdf/) | [npm](https://www.npmjs.com/package/peasy-pdf) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-pdf-go) | PDF merge, split, rotate, compress — [peasypdf.com](https://peasypdf.com) |
| peasy-image | [PyPI](https://pypi.org/project/peasy-image/) | [npm](https://www.npmjs.com/package/peasy-image) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-image-go) | Image resize, crop, convert, compress — [peasyimage.com](https://peasyimage.com) |
| peasy-audio | [PyPI](https://pypi.org/project/peasy-audio/) | [npm](https://www.npmjs.com/package/peasy-audio) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-audio-go) | Audio trim, merge, convert, normalize — [peasyaudio.com](https://peasyaudio.com) |
| peasy-video | [PyPI](https://pypi.org/project/peasy-video/) | [npm](https://www.npmjs.com/package/peasy-video) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-video-go) | Video trim, resize, thumbnails, GIF — [peasyvideo.com](https://peasyvideo.com) |
| peasy-css | [PyPI](https://pypi.org/project/peasy-css/) | [npm](https://www.npmjs.com/package/peasy-css) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-css-go) | CSS minify, format, analyze — [peasycss.com](https://peasycss.com) |
| peasy-compress | [PyPI](https://pypi.org/project/peasy-compress/) | [npm](https://www.npmjs.com/package/peasy-compress) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-compress-go) | ZIP, TAR, gzip compression — [peasytools.com](https://peasytools.com) |
| peasy-document | [PyPI](https://pypi.org/project/peasy-document/) | [npm](https://www.npmjs.com/package/peasy-document) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-document-go) | Markdown, HTML, CSV, JSON conversion — [peasyformats.com](https://peasyformats.com) |
| **peasytext** | [PyPI](https://pypi.org/project/peasytext/) | [npm](https://www.npmjs.com/package/peasytext) | [Go](https://pkg.go.dev/github.com/peasytools/peasytext-go) | **Text case conversion, slugify, word count — [peasytext.com](https://peasytext.com)** |

## License

MIT
