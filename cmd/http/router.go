package routerpkg

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Router struct {
	render func(http.ResponseWriter, string, any)
}

func New(render func(http.ResponseWriter, string, any)) http.Handler {
	r := &Router{render: render}
	mux := chi.NewRouter()
	mux.Get("/", r.Home)
	return mux
}

func (r *Router) Home(writer http.ResponseWriter, _ *http.Request) {
	data := map[string]any{
		"Title": "Homepage",
		"User":  "foo",
		"Year":  time.Now().Year(),
	}
	r.render(writer, "home", data)
}
