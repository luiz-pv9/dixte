package profiles

import (
	"errors"
	"github.com/luiz-pv9/dixte/databasemanager"
)

type Profile struct {
	ProfileId  int64
	ExternalId string
	CreatedAt  int64
	UpdatedAt  int64
	AppToken   string
	Properties map[string]interface{}
}

func DeleteAll() (int64, error) {
	db := databasemanager.Db.Conn
	result, err := db.Exec("DELETE FROM profiles")
	if err != nil {
		return int64(0), err
	}
	val, err := result.RowsAffected()
	if err != nil {
		return int64(0), err
	}
	return val, nil
}

func (p *Profile) CheckFields() error {
	if p.ExternalId == "" {
		return errors.New("Field ExternalId nt set in profile")
	}

	if p.AppToken == "" {
		return errors.New("Field AppToken not set in profile")
	}
	return nil
}

func (p *Profile) Track() error {
	previousProfile, err := FindInAppByExternalId(p.AppToken, p.ExternalId)
	if err != nil {
		return err
	}
	if previousProfile == nil {
		return p.create()
	}
	return p.update()
}

func (p *Profile) create() error {
	return nil
}

func (p *Profile) Delete() error {
}
