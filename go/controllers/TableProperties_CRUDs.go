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
var __TableProperties__dummysDeclaration__ models.TableProperties
var __TableProperties_time__dummyDeclaration time.Duration

var mutexTableProperties sync.Mutex

// An TablePropertiesID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTableProperties updateTableProperties deleteTableProperties
type TablePropertiesID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TablePropertiesInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTableProperties updateTableProperties
type TablePropertiesInput struct {
	// The TableProperties to submit or modify
	// in: body
	TableProperties *orm.TablePropertiesAPI
}

// GetTablePropertiess
//
// swagger:route GET /tablepropertiess tablepropertiess getTablePropertiess
//
// # Get all tablepropertiess
//
// Responses:
// default: genericError
//
//	200: tablepropertiesDBResponse
func (controller *Controller) GetTablePropertiess(c *gin.Context) {

	// source slice
	var tablepropertiesDBs []orm.TablePropertiesDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTablePropertiess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableProperties.GetDB()

	_, err := db.Find(&tablepropertiesDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tablepropertiesAPIs := make([]orm.TablePropertiesAPI, 0)

	// for each tableproperties, update fields from the database nullable fields
	for idx := range tablepropertiesDBs {
		tablepropertiesDB := &tablepropertiesDBs[idx]
		_ = tablepropertiesDB
		var tablepropertiesAPI orm.TablePropertiesAPI

		// insertion point for updating fields
		tablepropertiesAPI.ID = tablepropertiesDB.ID
		tablepropertiesDB.CopyBasicFieldsToTableProperties_WOP(&tablepropertiesAPI.TableProperties_WOP)
		tablepropertiesAPI.TablePropertiesPointersEncoding = tablepropertiesDB.TablePropertiesPointersEncoding
		tablepropertiesAPIs = append(tablepropertiesAPIs, tablepropertiesAPI)
	}

	c.JSON(http.StatusOK, tablepropertiesAPIs)
}

// PostTableProperties
//
// swagger:route POST /tablepropertiess tablepropertiess postTableProperties
//
// Creates a tableproperties
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTableProperties(c *gin.Context) {

	mutexTableProperties.Lock()
	defer mutexTableProperties.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTablePropertiess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableProperties.GetDB()

	// Validate input
	var input orm.TablePropertiesAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tableproperties
	tablepropertiesDB := orm.TablePropertiesDB{}
	tablepropertiesDB.TablePropertiesPointersEncoding = input.TablePropertiesPointersEncoding
	tablepropertiesDB.CopyBasicFieldsFromTableProperties_WOP(&input.TableProperties_WOP)

	_, err = db.Create(&tablepropertiesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTableProperties.CheckoutPhaseOneInstance(&tablepropertiesDB)
	tableproperties := backRepo.BackRepoTableProperties.Map_TablePropertiesDBID_TablePropertiesPtr[tablepropertiesDB.ID]

	if tableproperties != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tableproperties)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tablepropertiesDB)
}

// GetTableProperties
//
// swagger:route GET /tablepropertiess/{ID} tablepropertiess getTableProperties
//
// Gets the details for a tableproperties.
//
// Responses:
// default: genericError
//
//	200: tablepropertiesDBResponse
func (controller *Controller) GetTableProperties(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTableProperties", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableProperties.GetDB()

	// Get tablepropertiesDB in DB
	var tablepropertiesDB orm.TablePropertiesDB
	if _, err := db.First(&tablepropertiesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tablepropertiesAPI orm.TablePropertiesAPI
	tablepropertiesAPI.ID = tablepropertiesDB.ID
	tablepropertiesAPI.TablePropertiesPointersEncoding = tablepropertiesDB.TablePropertiesPointersEncoding
	tablepropertiesDB.CopyBasicFieldsToTableProperties_WOP(&tablepropertiesAPI.TableProperties_WOP)

	c.JSON(http.StatusOK, tablepropertiesAPI)
}

// UpdateTableProperties
//
// swagger:route PATCH /tablepropertiess/{ID} tablepropertiess updateTableProperties
//
// # Update a tableproperties
//
// Responses:
// default: genericError
//
//	200: tablepropertiesDBResponse
func (controller *Controller) UpdateTableProperties(c *gin.Context) {

	mutexTableProperties.Lock()
	defer mutexTableProperties.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTableProperties", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableProperties.GetDB()

	// Validate input
	var input orm.TablePropertiesAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tablepropertiesDB orm.TablePropertiesDB

	// fetch the tableproperties
	_, err := db.First(&tablepropertiesDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tablepropertiesDB.CopyBasicFieldsFromTableProperties_WOP(&input.TableProperties_WOP)
	tablepropertiesDB.TablePropertiesPointersEncoding = input.TablePropertiesPointersEncoding

	db, _ = db.Model(&tablepropertiesDB)
	_, err = db.Updates(&tablepropertiesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tablepropertiesNew := new(models.TableProperties)
	tablepropertiesDB.CopyBasicFieldsToTableProperties(tablepropertiesNew)

	// redeem pointers
	tablepropertiesDB.DecodePointers(backRepo, tablepropertiesNew)

	// get stage instance from DB instance, and call callback function
	tablepropertiesOld := backRepo.BackRepoTableProperties.Map_TablePropertiesDBID_TablePropertiesPtr[tablepropertiesDB.ID]
	if tablepropertiesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tablepropertiesOld, tablepropertiesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tablepropertiesDB
	c.JSON(http.StatusOK, tablepropertiesDB)
}

// DeleteTableProperties
//
// swagger:route DELETE /tablepropertiess/{ID} tablepropertiess deleteTableProperties
//
// # Delete a tableproperties
//
// default: genericError
//
//	200: tablepropertiesDBResponse
func (controller *Controller) DeleteTableProperties(c *gin.Context) {

	mutexTableProperties.Lock()
	defer mutexTableProperties.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTableProperties", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableProperties.GetDB()

	// Get model if exist
	var tablepropertiesDB orm.TablePropertiesDB
	if _, err := db.First(&tablepropertiesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&tablepropertiesDB)

	// get an instance (not staged) from DB instance, and call callback function
	tablepropertiesDeleted := new(models.TableProperties)
	tablepropertiesDB.CopyBasicFieldsToTableProperties(tablepropertiesDeleted)

	// get stage instance from DB instance, and call callback function
	tablepropertiesStaged := backRepo.BackRepoTableProperties.Map_TablePropertiesDBID_TablePropertiesPtr[tablepropertiesDB.ID]
	if tablepropertiesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tablepropertiesStaged, tablepropertiesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
