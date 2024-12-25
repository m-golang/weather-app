package main

import (
	"html/template"
	"path/filepath"

	"github.com/m-golang/weather-app/internal/models"
)

// templateData holds data for rendering templates
type templateData struct {
	CurrentYear int
	Weather     *models.Weather
}

// newTemplateCache loads and caches all templates for efficient reuse
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

		// Fetch all page templates
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

		// Parse templates and store them in the cache
	for _, page := range pages {
		name := filepath.Base(page)

				// Load base template and page-specific templates

		files := []string{
			"./ui/html/base.html",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil
}
