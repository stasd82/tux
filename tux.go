package tux

import (
	"os"
	"syscall"

	"github.com/uptrace/bunrouter"
)

type App struct {
	*bunrouter.Router
	shutdown chan os.Signal
}

func NewApp(shutdown chan os.Signal) *App {
	return &App{
		Router:   bunrouter.New(),
		shutdown: shutdown,
	}
}

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
