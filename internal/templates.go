package internal

import (
	"html/template"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type templateData struct {
}

type renderData struct {
	TemplateData *templateData
	I18n         *i18n.Bundle
	Config       *sharedConfig
}

func newTemplateCache(dir string, functions template.FuncMap) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "page.*"+templatesSuffix))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "layout.*"+templatesSuffix))
		if err != nil {
			return nil, err
		}

		// ts, err = ts.ParseGlob(filepath.Join(dir, "partial.*"+templatesSuffix))
		// if err != nil {
		// 	return nil, err
		// }

		cache[name] = ts
	}

	return cache, nil
}
