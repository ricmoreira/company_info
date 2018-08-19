package mrequest

import (
	"company_info/models/saft/go_SaftT104"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type HeaderCreate struct {
	go_SaftT104.TxsdHeader `bson:"inline"`
}

type HeaderRead struct {
	ID                     objectid.ObjectID `json:"id,omitempty" bson:"_id"`
	go_SaftT104.TxsdHeader `bson:"inline"`
}

type HeaderUpdate struct {
	go_SaftT104.TxsdHeader `bson:"inline"`
}

type HeaderDelete struct {
	ID                     objectid.ObjectID `bson:"_id" json:"id,omitempty" valid:"required~Cannot be empty" bson:"_id"`
	go_SaftT104.TxsdHeader `bson:"inline"`
}
