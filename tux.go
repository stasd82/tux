package tux

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/uptrace/bunrouter"
)

type RouteFunc func(context.Context, http.ResponseWriter, *http.Request) error

type Tux struct {
	*bunrouter.CompatRouter
	shutdown chan os.Signal
}

func New(c chan os.Signal) *Tux {
	return &Tux{
		CompatRouter: bunrouter.New().Compat(),
		shutdown:     c,
	}
}

func (t *Tux) Finish() {
	t.shutdown <- syscall.SIGTERM
}

func (t *Tux) AddRoute(verb string, group string, path string, rf RouteFunc) {

	h := func(w http.ResponseWriter, r *http.Request) {

		ctx := context.TODO()
		if err := rf(ctx, w, r); err != nil {
			t.Finish()
			return
		}

	}

	team := t.NewGroup("/" + group)

	switch verb {
	case http.MethodGet:
		team.GET(path, h)
	case http.MethodPost:
		team.POST(path, h)
	case http.MethodPut:
		team.PUT(path, h)
	case http.MethodDelete:
		team.DELETE(path, h)
	}
}
