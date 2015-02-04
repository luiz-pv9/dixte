package profiles

import (
	"database/sql"
	"encoding/json"
	"github.com/luiz-pv9/dixte/databasemanager"
)

func LoadFromSqlRow(row *sql.Row) (*Profile, error) {
	profile := &Profile{}
	var properties json.RawMessage
	err := row.Scan(&profile.ProfileId, &profile.ExternalId,
		&profile.CreatedAt, &profile.UpdatedAt, &properties, &profile.AppToken)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(properties, &profile.Properties)
	return profile, err
}

func FindInAppByExternalId(appToken, externalId string) (*Profile, error) {
	db := databasemanager.Db.Conn
	row := db.QueryRow(`SELECT profiles.* FROM profiles 
		WHERE profiles.app_token = $1 AND profiles.external_id = $2`,
		appToken, externalId)
	return LoadFromSqlRow(row)
}
