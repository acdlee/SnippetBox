package main

import (
	"html/template"
	"path/filepath"

	"github.com/acdlee/SnippetBox/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

// Returns an in-memory map to cache parsed templates.
// Cache is indexed by the name of the page (e.g. 'home.tmpl.html).
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	// Parse and add to cache the following templates:
	// current 'page' template; base page template; all partial templates
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
