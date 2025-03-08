package app

func (app *App) Shutdown() {
	if app.server != nil {
		_ = app.server.Shutdown()
	}

	if app.db != nil {
		_ = app.db.Close()
	}

	if app.rpc != nil {
		_ = app.rpc.Close()
	}
}
