package dao

import (
	"fmt"
	"govue/helpers"
	"govue/models"

	"github.com/eaciit/dbox"
)

type MasterUserDao struct {
	Conn dbox.IConnection
}

func (m *MasterUserDao) FindAll() []models.MasterUserModel {
	var ret []models.MasterUserModel
	e := m.Conn.NewQuery().From(helpers.CollectionMasterUser)
	q, err := e.Cursor(nil)
	if err != nil {
		fmt.Println("Error FINDALL Collection MasterUser", err.Error())
		return nil
	}
	ret = make([]models.MasterUserModel, 0)
	q.Fetch(&ret, 0, false)
	defer q.Close()
	return ret
}
