package apps

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/luiz-pv9/dixte-analytics/databasemanager"
	"github.com/luiz-pv9/dixte-analytics/environment"
	"log"
)

// Type definition used by the finder and register
type App struct {
	AppId int64
	Name  string
	Token string
}

func NewApp(name string) *App {
	return &App{Name: name}
}

func (app *App) Persist() error {
	if app.AppId == int64(0) {
		return app.Create()
	}
	return app.Update()
}

func (app *App) Create() error {
	db := databasemanager.Db.Conn
	err := db.QueryRow(`INSERT INTO apps (name, token) 
		VALUES ($1, $2) RETURNING app_id`, app.Name, app.Token).Scan(&app.AppId)
	return err
}

func (app *App) Update() error {
	db := databasemanager.Db.Conn
	_, err := db.Exec(`UPDATE apps SET name = $1, token = $2 
		WHERE app_id = $3`, app.Name, app.Token, app.AppId)
	return err
}

func (app *App) GenerateToken(dc *environment.Config) error {
	searching := true
	for searching {
		token := generateRandomToken(int(dc.App.Token_Size))
		appFound, err := FindByToken(token)
		if err != nil {
			log.Printf("appfinder.FindByToken\n%v\n", err)
			return err
		}
		if appFound == nil {
			app.Token = token
			searching = false
		}
	}
	return nil
}

func generateRandomToken(size int) string {
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		log.Println(err)
		return ""
	}
	// For some reason the result from encoding generates a string with length
	// greater than the length of the array fof bytes
	return base64.URLEncoding.EncodeToString(rb)[0:size]
}

func DeleteAll() (int64, error) {
	db := databasemanager.Db.Conn
	result, err := db.Exec("DELETE FROM apps")
	if err != nil {
		return int64(0), err
	}
	val, err := result.RowsAffected()
	if err != nil {
		return int64(0), err
	}
	return val, err
}
