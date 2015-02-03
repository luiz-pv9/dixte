package properties

import (
	"github.com/luiz-pv9/dixte-analytics/databasemanager"
)

func DeleteAll() error {
	db := databasemanager.Db.Conn
	_, err := db.Exec("DELETE FROM properties")
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM property_values")
	return err
}
