package apps

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/luiz-pv9/dixte-analytics/environment"
	"log"
)

// Type definition used by the finder and register
type App struct {
	Id    int
	Name  string
	Token string
}

func generateRandomToken(size int) string {
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(rb)
}

func (app *App) GenerateToken(dc *environment.Config) error {
	searching := false
	for searching {
		// token := generateRandomToken(int(dc.App.Token_Size))
		// appFound, err := appfinder.ByToken(token)
		// if appFound == nil {
		// 	app.Token = token
		// 	searching = false
		// }
	}
	return nil
}
