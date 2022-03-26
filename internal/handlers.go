package internal

import (
	"net/http"
)

func (app *application) handleView(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "page.home.gohtml", &templateData{})
}
