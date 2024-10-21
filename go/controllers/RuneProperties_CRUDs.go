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
var __RuneProperties__dummysDeclaration__ models.RuneProperties
var __RuneProperties_time__dummyDeclaration time.Duration

var mutexRuneProperties sync.Mutex

// An RunePropertiesID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRuneProperties updateRuneProperties deleteRuneProperties
type RunePropertiesID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RunePropertiesInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRuneProperties updateRuneProperties
type RunePropertiesInput struct {
	// The RuneProperties to submit or modify
	// in: body
	RuneProperties *orm.RunePropertiesAPI
}

// GetRunePropertiess
//
// swagger:route GET /runepropertiess runepropertiess getRunePropertiess
//
// # Get all runepropertiess
//
// Responses:
// default: genericError
//
//	200: runepropertiesDBResponse
func (controller *Controller) GetRunePropertiess(c *gin.Context) {

	// source slice
	var runepropertiesDBs []orm.RunePropertiesDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRunePropertiess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRuneProperties.GetDB()

	_, err := db.Find(&runepropertiesDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	runepropertiesAPIs := make([]orm.RunePropertiesAPI, 0)

	// for each runeproperties, update fields from the database nullable fields
	for idx := range runepropertiesDBs {
		runepropertiesDB := &runepropertiesDBs[idx]
		_ = runepropertiesDB
		var runepropertiesAPI orm.RunePropertiesAPI

		// insertion point for updating fields
		runepropertiesAPI.ID = runepropertiesDB.ID
		runepropertiesDB.CopyBasicFieldsToRuneProperties_WOP(&runepropertiesAPI.RuneProperties_WOP)
		runepropertiesAPI.RunePropertiesPointersEncoding = runepropertiesDB.RunePropertiesPointersEncoding
		runepropertiesAPIs = append(runepropertiesAPIs, runepropertiesAPI)
	}

	c.JSON(http.StatusOK, runepropertiesAPIs)
}

// PostRuneProperties
//
// swagger:route POST /runepropertiess runepropertiess postRuneProperties
//
// Creates a runeproperties
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRuneProperties(c *gin.Context) {

	mutexRuneProperties.Lock()
	defer mutexRuneProperties.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRunePropertiess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRuneProperties.GetDB()

	// Validate input
	var input orm.RunePropertiesAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create runeproperties
	runepropertiesDB := orm.RunePropertiesDB{}
	runepropertiesDB.RunePropertiesPointersEncoding = input.RunePropertiesPointersEncoding
	runepropertiesDB.CopyBasicFieldsFromRuneProperties_WOP(&input.RuneProperties_WOP)

	_, err = db.Create(&runepropertiesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRuneProperties.CheckoutPhaseOneInstance(&runepropertiesDB)
	runeproperties := backRepo.BackRepoRuneProperties.Map_RunePropertiesDBID_RunePropertiesPtr[runepropertiesDB.ID]

	if runeproperties != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), runeproperties)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, runepropertiesDB)
}

// GetRuneProperties
//
// swagger:route GET /runepropertiess/{ID} runepropertiess getRuneProperties
//
// Gets the details for a runeproperties.
//
// Responses:
// default: genericError
//
//	200: runepropertiesDBResponse
func (controller *Controller) GetRuneProperties(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRuneProperties", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRuneProperties.GetDB()

	// Get runepropertiesDB in DB
	var runepropertiesDB orm.RunePropertiesDB
	if _, err := db.First(&runepropertiesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var runepropertiesAPI orm.RunePropertiesAPI
	runepropertiesAPI.ID = runepropertiesDB.ID
	runepropertiesAPI.RunePropertiesPointersEncoding = runepropertiesDB.RunePropertiesPointersEncoding
	runepropertiesDB.CopyBasicFieldsToRuneProperties_WOP(&runepropertiesAPI.RuneProperties_WOP)

	c.JSON(http.StatusOK, runepropertiesAPI)
}

// UpdateRuneProperties
//
// swagger:route PATCH /runepropertiess/{ID} runepropertiess updateRuneProperties
//
// # Update a runeproperties
//
// Responses:
// default: genericError
//
//	200: runepropertiesDBResponse
func (controller *Controller) UpdateRuneProperties(c *gin.Context) {

	mutexRuneProperties.Lock()
	defer mutexRuneProperties.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRuneProperties", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRuneProperties.GetDB()

	// Validate input
	var input orm.RunePropertiesAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var runepropertiesDB orm.RunePropertiesDB

	// fetch the runeproperties
	_, err := db.First(&runepropertiesDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	runepropertiesDB.CopyBasicFieldsFromRuneProperties_WOP(&input.RuneProperties_WOP)
	runepropertiesDB.RunePropertiesPointersEncoding = input.RunePropertiesPointersEncoding

	db, _ = db.Model(&runepropertiesDB)
	_, err = db.Updates(runepropertiesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	runepropertiesNew := new(models.RuneProperties)
	runepropertiesDB.CopyBasicFieldsToRuneProperties(runepropertiesNew)

	// redeem pointers
	runepropertiesDB.DecodePointers(backRepo, runepropertiesNew)

	// get stage instance from DB instance, and call callback function
	runepropertiesOld := backRepo.BackRepoRuneProperties.Map_RunePropertiesDBID_RunePropertiesPtr[runepropertiesDB.ID]
	if runepropertiesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), runepropertiesOld, runepropertiesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the runepropertiesDB
	c.JSON(http.StatusOK, runepropertiesDB)
}

// DeleteRuneProperties
//
// swagger:route DELETE /runepropertiess/{ID} runepropertiess deleteRuneProperties
//
// # Delete a runeproperties
//
// default: genericError
//
//	200: runepropertiesDBResponse
func (controller *Controller) DeleteRuneProperties(c *gin.Context) {

	mutexRuneProperties.Lock()
	defer mutexRuneProperties.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRuneProperties", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRuneProperties.GetDB()

	// Get model if exist
	var runepropertiesDB orm.RunePropertiesDB
	if _, err := db.First(&runepropertiesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&runepropertiesDB)

	// get an instance (not staged) from DB instance, and call callback function
	runepropertiesDeleted := new(models.RuneProperties)
	runepropertiesDB.CopyBasicFieldsToRuneProperties(runepropertiesDeleted)

	// get stage instance from DB instance, and call callback function
	runepropertiesStaged := backRepo.BackRepoRuneProperties.Map_RunePropertiesDBID_RunePropertiesPtr[runepropertiesDB.ID]
	if runepropertiesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), runepropertiesStaged, runepropertiesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
