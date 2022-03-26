package internal

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

const templatesDir = "./web/templates/"
const templatesSuffix = ".gohtml"

func (app *application) sendInternalServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.logger.Error(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// func (app *application) sendNotFoundError(w http.ResponseWriter) {
// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// }

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.sendInternalServerError(w, fmt.Errorf("template %s does not exist", name))
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, td)
	if err != nil {
		app.sendInternalServerError(w, err)
	}

	buf.WriteTo(w)
}
