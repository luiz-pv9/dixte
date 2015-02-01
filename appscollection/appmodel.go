package appscollection

// The App struct represents a register from the apps table
type App struct {
	Id    int
	Name  string
	Token string
}

func (app *App) Persist() error {
	return nil
}

func (app *App) GenerateToken() {
}

func NewApp(name string) *App {
	return &App{Name: name}
}
