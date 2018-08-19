package mresponse

import (
	"company_info/models/saft/go_SaftT104"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type Header struct {
	ID                     objectid.ObjectID `json:"id,omitempty" bson:"_id"`
	go_SaftT104.TxsdHeader `bson:"inline"`
}

type HeaderCreate struct {
	ID                     string `json:"id,omitempty"`
	go_SaftT104.TxsdHeader `bson:"inline"`
}

type HeaderRead struct {
	ID                     string            `json:"id,omitempty"`
	IDdb                   objectid.ObjectID `json:"-" bson:"_id"`
	go_SaftT104.TxsdHeader `bson:"inline"`
}

type HeaderList struct {
	Total   int64          `json:"total"`
	PerPage int64          `json:"per_page"`
	Page    int64          `json:"page"`
	Items   *[]*HeaderRead `json:"items"`
}
