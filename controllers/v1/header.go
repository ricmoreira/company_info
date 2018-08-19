package controllers

import (
	"company_info/models/request"
	"company_info/services"
	"company_info/util/errors"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type (
	// HeaderController represents the controller for operating on the Headers resource
	HeaderController struct {
		HeaderService services.HeaderServiceContract
	}
)

// NewHeaderController is the constructor of HeaderController
func NewHeaderController(ps *services.HeaderService) *HeaderController {
	return &HeaderController{
		HeaderService: ps,
	}
}

// CreateAction creates a new Header
func (pc HeaderController) CreateAction(c *gin.Context) {
	iReq := mrequest.HeaderCreate{}
	json.NewDecoder(c.Request.Body).Decode(&iReq)

	e := errors.ValidateRequest(&iReq)
	if e != nil {
		c.JSON(e.HttpCode, e)
		return
	}

	iRes, err := pc.HeaderService.CreateOne(&iReq)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	c.JSON(200, iRes)
}

// ListAction list Headers
func (pc HeaderController) ListAction(c *gin.Context) {
	validSorts := map[string]string{}
	validSorts["_id"] = "_id"

	validFilters := map[string]string{}
	validFilters["TaxRegistrationNumber"] = "TaxRegistrationNumber"
	validFilters["_id"] = "_id"

	qValues := c.Request.URL.Query()
	req := mrequest.NewListRequest(qValues, validSorts, validFilters)

	res, err := pc.HeaderService.List(req)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	c.JSON(200, res)
}
