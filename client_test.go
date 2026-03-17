package peasytext

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func paginatedJSON(t *testing.T, results any, count int) string {
	t.Helper()
	b, err := json.Marshal(results)
	if err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf(`{"count":%d,"next":null,"previous":null,"results":%s}`, count, string(b))
}

func TestNew(t *testing.T) {
	c := New()
	if c.baseURL != DefaultBaseURL {
		t.Errorf("expected %s, got %s", DefaultBaseURL, c.baseURL)
	}
	c2 := New(WithBaseURL("https://example.com"))
	if c2.baseURL != "https://example.com" {
		t.Errorf("expected https://example.com, got %s", c2.baseURL)
	}
}

func TestWithBaseURLTrailingSlash(t *testing.T) {
	c := New(WithBaseURL("https://example.com/"))
	if c.baseURL != "https://example.com" {
		t.Errorf("expected trailing slash stripped, got %s", c.baseURL)
	}
}

func TestListTools(t *testing.T) {
	ctx := context.Background()
	tools := []Tool{{Slug: "image-resize", Name: "Image Resize", Description: "Resize images", Category: "image"}}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/tools/" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		fmt.Fprint(w, paginatedJSON(t, tools, 1))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	resp, err := c.ListTools(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Count != 1 {
		t.Errorf("expected count 1, got %d", resp.Count)
	}
	if resp.Results[0].Slug != "image-resize" {
		t.Errorf("expected image-resize, got %s", resp.Results[0].Slug)
	}
}

func TestListToolsWithSearch(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if q := r.URL.Query().Get("search"); q != "resize" {
			t.Errorf("expected search=resize, got %s", q)
		}
		fmt.Fprint(w, paginatedJSON(t, []Tool{}, 0))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	_, err := c.ListTools(ctx, ListOptions{Search: "resize"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTool(t *testing.T) {
	ctx := context.Background()
	tool := Tool{Slug: "image-resize", Name: "Image Resize", Description: "Resize images", Category: "image"}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/tools/image-resize/" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(tool)
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	got, err := c.GetTool(ctx, "image-resize")
	if err != nil {
		t.Fatal(err)
	}
	if got.Slug != "image-resize" {
		t.Errorf("expected image-resize, got %s", got.Slug)
	}
}

func TestListGlossary(t *testing.T) {
	ctx := context.Background()
	terms := []GlossaryTerm{{Slug: "webp", Term: "WebP", Definition: "Modern image format", Category: "formats"}}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, paginatedJSON(t, terms, 1))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	resp, err := c.ListGlossary(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Results[0].Term != "WebP" {
		t.Errorf("expected WebP, got %s", resp.Results[0].Term)
	}
}

func TestGetGlossaryTerm(t *testing.T) {
	ctx := context.Background()
	term := GlossaryTerm{Slug: "webp", Term: "WebP", Definition: "Modern image format", Category: "formats"}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(term)
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	got, err := c.GetGlossaryTerm(ctx, "webp")
	if err != nil {
		t.Fatal(err)
	}
	if got.Definition != "Modern image format" {
		t.Errorf("expected Modern image format, got %s", got.Definition)
	}
}

func TestListGuides(t *testing.T) {
	ctx := context.Background()
	guides := []Guide{{Slug: "resize-images", Title: "How to Resize Images", Category: "optimization", AudienceLevel: "beginner", WordCount: 1200}}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, paginatedJSON(t, guides, 1))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	resp, err := c.ListGuides(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Results[0].Title != "How to Resize Images" {
		t.Errorf("expected How to Resize Images, got %s", resp.Results[0].Title)
	}
}

func TestSearch(t *testing.T) {
	ctx := context.Background()
	result := SearchResult{
		Query: "resize",
		Results: SearchCategories{
			Tools:    []Tool{{Slug: "image-resize", Name: "Image Resize"}},
			Formats:  []Format{},
			Glossary: []GlossaryTerm{},
		},
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if q := r.URL.Query().Get("q"); q != "resize" {
			t.Errorf("expected q=resize, got %s", q)
		}
		json.NewEncoder(w).Encode(result)
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	got, err := c.Search(ctx, "resize")
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Results.Tools) != 1 {
		t.Errorf("expected 1 tool result, got %d", len(got.Results.Tools))
	}
}

func TestNotFoundError(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"detail":"Not found."}`)
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	_, err := c.GetTool(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error")
	}
	var nfe *NotFoundError
	if !errors.As(err, &nfe) {
		t.Errorf("expected NotFoundError, got %T", err)
	}
}

func TestServerError(t *testing.T) {
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	_, err := c.ListTools(ctx)
	if err == nil {
		t.Fatal("expected error")
	}
	var pe *PeasyError
	if !errors.As(err, &pe) {
		t.Errorf("expected PeasyError, got %T", err)
	}
	if pe.StatusCode != 500 {
		t.Errorf("expected 500, got %d", pe.StatusCode)
	}
}
