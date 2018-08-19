package services

import (
	"company_info/models/request"
	"company_info/models/response"
	"company_info/repositories"
	"company_info/util/errors"
	"context"

	"log"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// HeaderServiceContract is the abstraction for service layer on roles resource
type HeaderServiceContract interface {
	CreateOne(*mrequest.HeaderCreate) (*mresponse.HeaderCreate, *mresponse.ErrorResponse)
	ReadOne(*mrequest.HeaderRead) (*mresponse.HeaderRead, *mresponse.ErrorResponse)
	UpdateOne(*mrequest.HeaderUpdate) (*mresponse.Header, *mresponse.ErrorResponse)
	DeleteOne(*mrequest.HeaderDelete) (*mresponse.Header, *mresponse.ErrorResponse)
	CreateMany(*[]*mrequest.HeaderCreate) (*[]*mresponse.HeaderCreate, *mresponse.ErrorResponse)
	List(request *mrequest.ListRequest) (*mresponse.HeaderList, *mresponse.ErrorResponse)
}

// HeaderService is the layer between http client and repository for Header resource
type HeaderService struct {
	HeaderRepository *repositories.HeaderRepository
}

// NewHeaderService is the constructor of HeaderService
func NewHeaderService(hr *repositories.HeaderRepository) *HeaderService {
	return &HeaderService{
		HeaderRepository: hr,
	}
}

// CreateOne saves provided model instance to database
func (this *HeaderService) CreateOne(request *mrequest.HeaderCreate) (*mresponse.HeaderCreate, *mresponse.ErrorResponse) {

	// validate request
	e := errors.ValidateRequest(request)
	if e != nil {
		return nil, e
	}

	res, err := this.HeaderRepository.CreateOne(request)

	if err != nil {
		errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
		return nil, errR
	}

	id := res.InsertedID.(objectid.ObjectID)
	p := mresponse.HeaderCreate{
		ID: id.Hex(),
	}

	return &p, nil
}

// ReadOne returns a HeaderRead read instance
func (this *HeaderService) ReadOne(request *mrequest.HeaderRead) (*mresponse.HeaderRead, *mresponse.ErrorResponse) {
	// validate request
	e := errors.ValidateRequest(request)
	if e != nil {
		return nil, e
	}

	res, err := this.HeaderRepository.ReadOne(request)

	if err != nil {
		errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
		return nil, errR
	}

	h := mresponse.HeaderRead{
		ID: res.IDdb.Hex(),
	}

	return &h, nil
}

// TODO: implement
func (this *HeaderService) UpdateOne(p *mrequest.HeaderUpdate) (*mresponse.Header, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (this *HeaderService) DeleteOne(p *mrequest.HeaderDelete) (*mresponse.Header, *mresponse.ErrorResponse) {
	return nil, nil
}

// CreateMany saves many Headers in one bulk operation
func (this *HeaderService) CreateMany(request *[]*mrequest.HeaderCreate) (*[]*mresponse.HeaderCreate, *mresponse.ErrorResponse) {

	res, err := this.HeaderRepository.InsertMany(request)

	if err != nil {
		mngBulkError := err.(mongo.BulkWriteError)
		writeErrors := mngBulkError.WriteErrors
		for _, err := range writeErrors {
			log.Println(err)
		}
	}

	result := make([]*mresponse.HeaderCreate, len(res.InsertedIDs))
	for i, insertedID := range res.InsertedIDs {
		id := insertedID.(objectid.ObjectID)
		result[i] = &mresponse.HeaderCreate{
			ID: id.Hex(),
		}
	}

	return &result, nil
}

// List returns a list of Headers with pagination and filtering options
func (this *HeaderService) List(request *mrequest.ListRequest) (*mresponse.HeaderList, *mresponse.ErrorResponse) {

	total, perPage, page, cursor, err := this.HeaderRepository.List(request)

	if err != nil {
		e := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
		return nil, e
	}

	docs := []*mresponse.HeaderRead{}

	for cursor.Next(context.Background()) {
		doc := mresponse.HeaderRead{}
		if err := cursor.Decode(&doc); err != nil {
			errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
			return nil, errR
		}

		doc.ID = doc.IDdb.Hex()

		docs = append(docs, &doc)
	}
	
	resp := mresponse.HeaderList{
		Total: total,
		PerPage: perPage,
		Page: page,
		Items: &docs,
	}
	return &resp, nil
}
