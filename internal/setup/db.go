package setup

import (
	"gin-admin/global"
	"gin-admin/pkg/db"
)

func setupDB() error {
	var err error
	global.DB, err = db.NewDB(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
