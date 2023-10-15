// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongdocx/go/models"
	"github.com/fullstack-lang/gongdocx/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Body__dummysDeclaration__ models.Body
var __Body_time__dummyDeclaration time.Duration

var mutexBody sync.Mutex

// An BodyID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBody updateBody deleteBody
type BodyID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BodyInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBody updateBody
type BodyInput struct {
	// The Body to submit or modify
	// in: body
	Body *orm.BodyAPI
}

// GetBodys
//
// swagger:route GET /bodys bodys getBodys
//
// # Get all bodys
//
// Responses:
// default: genericError
//
//	200: bodyDBResponse
func (controller *Controller) GetBodys(c *gin.Context) {

	// source slice
	var bodyDBs []orm.BodyDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBodys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBody.GetDB()

	query := db.Find(&bodyDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bodyAPIs := make([]orm.BodyAPI, 0)

	// for each body, update fields from the database nullable fields
	for idx := range bodyDBs {
		bodyDB := &bodyDBs[idx]
		_ = bodyDB
		var bodyAPI orm.BodyAPI

		// insertion point for updating fields
		bodyAPI.ID = bodyDB.ID
		bodyDB.CopyBasicFieldsToBody_WOP(&bodyAPI.Body_WOP)
		bodyAPI.BodyPointersEncoding = bodyDB.BodyPointersEncoding
		bodyAPIs = append(bodyAPIs, bodyAPI)
	}

	c.JSON(http.StatusOK, bodyAPIs)
}

// PostBody
//
// swagger:route POST /bodys bodys postBody
//
// Creates a body
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBody(c *gin.Context) {

	mutexBody.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBodys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBody.GetDB()

	// Validate input
	var input orm.BodyAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create body
	bodyDB := orm.BodyDB{}
	bodyDB.BodyPointersEncoding = input.BodyPointersEncoding
	bodyDB.CopyBasicFieldsFromBody_WOP(&input.Body_WOP)

	query := db.Create(&bodyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBody.CheckoutPhaseOneInstance(&bodyDB)
	body := backRepo.BackRepoBody.Map_BodyDBID_BodyPtr[bodyDB.ID]

	if body != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), body)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bodyDB)

	mutexBody.Unlock()
}

// GetBody
//
// swagger:route GET /bodys/{ID} bodys getBody
//
// Gets the details for a body.
//
// Responses:
// default: genericError
//
//	200: bodyDBResponse
func (controller *Controller) GetBody(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBody", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBody.GetDB()

	// Get bodyDB in DB
	var bodyDB orm.BodyDB
	if err := db.First(&bodyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bodyAPI orm.BodyAPI
	bodyAPI.ID = bodyDB.ID
	bodyAPI.BodyPointersEncoding = bodyDB.BodyPointersEncoding
	bodyDB.CopyBasicFieldsToBody_WOP(&bodyAPI.Body_WOP)

	c.JSON(http.StatusOK, bodyAPI)
}

// UpdateBody
//
// swagger:route PATCH /bodys/{ID} bodys updateBody
//
// # Update a body
//
// Responses:
// default: genericError
//
//	200: bodyDBResponse
func (controller *Controller) UpdateBody(c *gin.Context) {

	mutexBody.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBody", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBody.GetDB()

	// Validate input
	var input orm.BodyAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bodyDB orm.BodyDB

	// fetch the body
	query := db.First(&bodyDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bodyDB.CopyBasicFieldsFromBody_WOP(&input.Body_WOP)
	bodyDB.BodyPointersEncoding = input.BodyPointersEncoding

	query = db.Model(&bodyDB).Updates(bodyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bodyNew := new(models.Body)
	bodyDB.CopyBasicFieldsToBody(bodyNew)

	// get stage instance from DB instance, and call callback function
	bodyOld := backRepo.BackRepoBody.Map_BodyDBID_BodyPtr[bodyDB.ID]
	if bodyOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bodyOld, bodyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bodyDB
	c.JSON(http.StatusOK, bodyDB)

	mutexBody.Unlock()
}

// DeleteBody
//
// swagger:route DELETE /bodys/{ID} bodys deleteBody
//
// # Delete a body
//
// default: genericError
//
//	200: bodyDBResponse
func (controller *Controller) DeleteBody(c *gin.Context) {

	mutexBody.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBody", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBody.GetDB()

	// Get model if exist
	var bodyDB orm.BodyDB
	if err := db.First(&bodyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bodyDB)

	// get an instance (not staged) from DB instance, and call callback function
	bodyDeleted := new(models.Body)
	bodyDB.CopyBasicFieldsToBody(bodyDeleted)

	// get stage instance from DB instance, and call callback function
	bodyStaged := backRepo.BackRepoBody.Map_BodyDBID_BodyPtr[bodyDB.ID]
	if bodyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bodyStaged, bodyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexBody.Unlock()
}
