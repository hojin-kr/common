package common

import (
	"github.com/hojin-kr/datastore"
)

type Common struct {
	UUID         string `json:"uuid"`
	Lattitude    string `json:"lattitude"`
	Longitude    string `json:"longitude"`
	PushToken    string `json:"pushToken"`
	Nickname     string `json:"nickname"`
	ProfileImage string `json:"profileImage"`
}

var ds datastore.GcpDatastore

// get common data
func (common *Common) GetCommon() {
	ds.Init()
	ds.GetEntity(common.UUID, common)
}

// set common data
func (common *Common) SetCommon() {
	ds.Init()
	ds.PutEntity(common.UUID, common)
}
