package repositories

import (
	"company_info/models/request"
	"company_info/models/response"
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/mongodb/mongo-go-driver/mongo/insertopt"
)

// HeaderRepository performs CRUD operations on users resource
type HeaderRepository struct {
	Headers MongoCollection
}

// NewHeaderRepository is the constructor for HeaderRepository
func NewHeaderRepository(db *DBCollections) *HeaderRepository {
	return &HeaderRepository{Headers: db.Header}
}

// CreateOne saves provided model instance to database
func (this *HeaderRepository) CreateOne(request *mrequest.HeaderCreate) (*mongo.InsertOneResult, error) {

	return this.Headers.InsertOne(context.Background(), request)
}

// ReadOne returns a Header based on TaxRegistrationNumber sent in request
// TODO: implement better query based on full request and not only the ProducCode
func (this *HeaderRepository) ReadOne(h *mrequest.HeaderRead) (*mresponse.HeaderRead, error) {
	result := this.Headers.FindOne(
		context.Background(),
		bson.NewDocument(bson.EC.String("TaxRegistrationNumber", string(h.TaxRegistrationNumber))),
	)

	res := mresponse.HeaderRead{}
	err := result.Decode(h)

	if err != nil {
		return nil, err
	}
	res.ID = res.IDdb.Hex()
	
	return &res, nil
}

// TODO: implement
func (this *HeaderRepository) UpdateOne(p *mrequest.HeaderUpdate) (*mresponse.Header, error) {
	return nil, nil
}

// TODO: implement
func (this *HeaderRepository) DeleteOne(p *mrequest.HeaderDelete) (*mresponse.Header, error) {
	return nil, nil
}

func (this *HeaderRepository) InsertMany(request *[]*mrequest.HeaderCreate) (*mongo.InsertManyResult, error) {
	// transform to []interface{} (https://golang.org/doc/faq#convert_slice_of_interface)
	s := make([]interface{}, len(*request))
	for i, v := range *request {
		s[i] = v
	}

	// { ordered: false } ordered is false in order to don't stop execution because an error ocurred on one of the inserts
	opt := insertopt.Ordered(false)
	return this.Headers.InsertMany(context.Background(), s, opt)
}

func (this *HeaderRepository) List(req *mrequest.ListRequest) (int64, int64, int64, mongo.Cursor, error) {

	args := []*bson.Element{}

	for i, v := range req.Filters {
		args = append(args, bson.EC.String(i, fmt.Sprintf("%v", v)))
	}

	total, e := this.Headers.Count(
		context.Background(),
		bson.NewDocument(args...),
	)

	sorting := map[string]int{}
	var sortingValue int
	if req.Order == "reverse" {
		sortingValue = -1
	} else {
		sortingValue = 1
	}
	sorting[req.Sort] = sortingValue

	perPage := int64(req.PerPage)
	page := int64(req.Page)
	cursor, e := this.Headers.Find(
		context.Background(),
		bson.NewDocument(args...),
		findopt.Sort(sorting),
		findopt.Skip(int64(req.PerPage*(req.Page-1))),
		findopt.Limit(perPage),
	)

	return total, perPage, page, cursor, e
}
