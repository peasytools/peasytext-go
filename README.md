# peasytext-go

[![Go Reference](https://pkg.go.dev/badge/github.com/peasytools/peasytext-go.svg)](https://pkg.go.dev/github.com/peasytools/peasytext-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/peasytools/peasytext-go)](https://goreportcard.com/report/github.com/peasytools/peasytext-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub stars](https://agentgif.com/badge/github/peasytools/peasytext-go/stars.svg)](https://github.com/peasytools/peasytext-go)

Go client for the [PeasyText](https://peasytext.com) API -- text case conversion, slug generation, word counting, and encoding utilities with tools for camelCase, snake_case, kebab-case, PascalCase, and more. Zero dependencies beyond the Go standard library.

Built from [PeasyText](https://peasytext.com), a comprehensive text processing toolkit offering free online tools for case conversion, slug generation, word counting, and text analysis. The glossary covers text concepts from character encodings to Unicode normalization, while guides explain text processing strategies and encoding best practices.

> **Try the interactive tools at [peasytext.com](https://peasytext.com)** -- [Text Counter](https://peasytext.com/text/text-counter/), [Case Converter](https://peasytext.com/text/text-case-converter/), [Slug Generator](https://peasytext.com/text/slug-generator/), and more.

<p align="center">
  <img src="demo.gif" alt="peasytext-go demo -- text case conversion, slug generation, and word count tools in Go terminal" width="800">
</p>

## Table of Contents

- [Install](#install)
- [Quick Start](#quick-start)
- [What You Can Do](#what-you-can-do)
  - [Text Processing Tools](#text-processing-tools)
  - [Browse Text Reference Content](#browse-text-reference-content)
  - [Search and Discovery](#search-and-discovery)
- [API Client](#api-client)
  - [Available Methods](#available-methods)
- [Learn More About Text Processing](#learn-more-about-text-processing)
- [Also Available](#also-available)
- [Peasy Developer Tools](#peasy-developer-tools)
- [License](#license)

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

## What You Can Do

### Text Processing Tools

Text transformation is fundamental to software development -- from generating URL-safe slugs for blog posts to converting variable names between camelCase and snake_case during code generation. Case conversion tools handle the mechanical work of transforming text between naming conventions used across programming languages (camelCase in JavaScript, snake_case in Python, kebab-case in CSS). Word counting and character analysis help content authors meet length requirements for SEO titles, meta descriptions, and social media posts.

| Tool | Description | Use Case |
|------|-------------|----------|
| Text Counter | Count words, characters, sentences, and paragraphs | Content length validation, SEO metadata |
| Case Converter | Convert between camelCase, snake_case, kebab-case, PascalCase | Code generation, API field mapping |
| Slug Generator | Create URL-safe slugs from any text input | Blog posts, CMS content, URL routing |

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

	// Fetch the case converter tool details
	tool, err := client.GetTool(ctx, "text-case-converter")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tool: %s\n", tool.Name)           // Case converter tool
	fmt.Printf("Category: %s\n", tool.Category)   // Text processing category

	// List supported text formats
	formats, err := client.ListFormats(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range formats.Results {
		fmt.Printf("%s (%s): %s\n", f.Name, f.Extension, f.MimeType)
	}

	// List available format conversions from plain text
	conversions, err := client.ListConversions(ctx, &peasytext.ListConversionsOptions{Source: str("text")})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d conversion paths from text\n", len(conversions.Results))
}

func str(s string) *string { return &s }
```

Learn more: [Text Counter](https://peasytext.com/text/text-counter/) · [Case Converter](https://peasytext.com/text/text-case-converter/) · [Text Encoding Guide](https://peasytext.com/guides/text-encoding-utf8-ascii/)

### Browse Text Reference Content

The text processing glossary explains key concepts from basic character encodings to advanced Unicode normalization forms. Understanding the difference between ASCII and UTF-8, how byte order marks (BOM) affect file parsing, and when to use URL encoding versus Base64 helps developers handle text data correctly across platforms and programming languages.

| Glossary Term | Description |
|---------------|-------------|
| Slug | URL-safe string derived from a title or phrase, using hyphens and lowercase |
| Case Conversion | Transforming text between naming conventions like camelCase and snake_case |
| UTF-8 | Variable-width character encoding capable of representing all Unicode code points |
| Whitespace | Characters representing horizontal or vertical space in text rendering |

```go
// Browse text processing glossary terms
glossary, err := client.ListGlossary(ctx, &peasytext.ListOptions{Search: str("case-conversion")})
if err != nil {
	log.Fatal(err)
}
for _, term := range glossary.Results {
	fmt.Printf("%s: %s\n", term.Term, term.Definition) // Text processing concept definition
}

// Read in-depth guides on text encoding and conversion
guides, err := client.ListGuides(ctx, &peasytext.ListGuidesOptions{Category: str("formatting")})
if err != nil {
	log.Fatal(err)
}
for _, g := range guides.Results {
	fmt.Printf("%s (%s)\n", g.Title, g.AudienceLevel) // Guide title and difficulty level
}
```

Learn more: [Slug Glossary](https://peasytext.com/glossary/slug/) · [ASCII Glossary](https://peasytext.com/glossary/ascii/) · [Regex Cheat Sheet](https://peasytext.com/guides/regex-cheat-sheet-essential-patterns/)

### Search and Discovery

The unified search endpoint queries across all text tools, glossary terms, guides, and supported formats simultaneously. This is useful for building text editor plugins, documentation search, or CLI tools that need to discover text processing capabilities.

```go
// Search across all text tools, glossary, and guides
results, err := client.Search(ctx, "slugify url", nil)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Found %d tools, %d glossary terms\n",
	len(results.Results.Tools),
	len(results.Results.Glossary)) // Cross-content text processing search results
```

Learn more: [BOM Glossary](https://peasytext.com/glossary/bom/) · [Slug Generator](https://peasytext.com/text/slug-generator/) · [All Text Guides](https://peasytext.com/guides/)

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

## Learn More About Text Processing

- **Tools**: [Text Counter](https://peasytext.com/text/text-counter/) · [Case Converter](https://peasytext.com/text/text-case-converter/) · [Slug Generator](https://peasytext.com/text/slug-generator/) · [All Tools](https://peasytext.com/)
- **Guides**: [Text Encoding Guide](https://peasytext.com/guides/text-encoding-utf8-ascii/) · [Regex Cheat Sheet](https://peasytext.com/guides/regex-cheat-sheet-essential-patterns/) · [All Guides](https://peasytext.com/guides/)
- **Glossary**: [Slug](https://peasytext.com/glossary/slug/) · [ASCII](https://peasytext.com/glossary/ascii/) · [BOM](https://peasytext.com/glossary/bom/) · [All Terms](https://peasytext.com/glossary/)
- **Formats**: [TXT](https://peasytext.com/formats/txt/) · [CSV](https://peasytext.com/formats/csv/) · [All Formats](https://peasytext.com/formats/)
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
| peasy-pdf | [PyPI](https://pypi.org/project/peasy-pdf/) | [npm](https://www.npmjs.com/package/peasy-pdf) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-pdf-go) | PDF merge, split, rotate, compress -- [peasypdf.com](https://peasypdf.com) |
| peasy-image | [PyPI](https://pypi.org/project/peasy-image/) | [npm](https://www.npmjs.com/package/peasy-image) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-image-go) | Image resize, crop, convert, compress -- [peasyimage.com](https://peasyimage.com) |
| peasy-audio | [PyPI](https://pypi.org/project/peasy-audio/) | [npm](https://www.npmjs.com/package/peasy-audio) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-audio-go) | Audio trim, merge, convert, normalize -- [peasyaudio.com](https://peasyaudio.com) |
| peasy-video | [PyPI](https://pypi.org/project/peasy-video/) | [npm](https://www.npmjs.com/package/peasy-video) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-video-go) | Video trim, resize, thumbnails, GIF -- [peasyvideo.com](https://peasyvideo.com) |
| peasy-css | [PyPI](https://pypi.org/project/peasy-css/) | [npm](https://www.npmjs.com/package/peasy-css) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-css-go) | CSS minify, format, analyze -- [peasycss.com](https://peasycss.com) |
| peasy-compress | [PyPI](https://pypi.org/project/peasy-compress/) | [npm](https://www.npmjs.com/package/peasy-compress) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-compress-go) | ZIP, TAR, gzip compression -- [peasytools.com](https://peasytools.com) |
| peasy-document | [PyPI](https://pypi.org/project/peasy-document/) | [npm](https://www.npmjs.com/package/peasy-document) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-document-go) | Markdown, HTML, CSV, JSON conversion -- [peasyformats.com](https://peasyformats.com) |
| **peasytext** | [PyPI](https://pypi.org/project/peasytext/) | [npm](https://www.npmjs.com/package/peasytext) | [Go](https://pkg.go.dev/github.com/peasytools/peasytext-go) | **Text case conversion, slugify, word count -- [peasytext.com](https://peasytext.com)** |

## License

MIT
