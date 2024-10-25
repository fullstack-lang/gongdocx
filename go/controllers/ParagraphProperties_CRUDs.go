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
var __ParagraphProperties__dummysDeclaration__ models.ParagraphProperties
var __ParagraphProperties_time__dummyDeclaration time.Duration

var mutexParagraphProperties sync.Mutex

// An ParagraphPropertiesID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getParagraphProperties updateParagraphProperties deleteParagraphProperties
type ParagraphPropertiesID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ParagraphPropertiesInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postParagraphProperties updateParagraphProperties
type ParagraphPropertiesInput struct {
	// The ParagraphProperties to submit or modify
	// in: body
	ParagraphProperties *orm.ParagraphPropertiesAPI
}

// GetParagraphPropertiess
//
// swagger:route GET /paragraphpropertiess paragraphpropertiess getParagraphPropertiess
//
// # Get all paragraphpropertiess
//
// Responses:
// default: genericError
//
//	200: paragraphpropertiesDBResponse
func (controller *Controller) GetParagraphPropertiess(c *gin.Context) {

	// source slice
	var paragraphpropertiesDBs []orm.ParagraphPropertiesDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetParagraphPropertiess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphProperties.GetDB()

	_, err := db.Find(&paragraphpropertiesDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	paragraphpropertiesAPIs := make([]orm.ParagraphPropertiesAPI, 0)

	// for each paragraphproperties, update fields from the database nullable fields
	for idx := range paragraphpropertiesDBs {
		paragraphpropertiesDB := &paragraphpropertiesDBs[idx]
		_ = paragraphpropertiesDB
		var paragraphpropertiesAPI orm.ParagraphPropertiesAPI

		// insertion point for updating fields
		paragraphpropertiesAPI.ID = paragraphpropertiesDB.ID
		paragraphpropertiesDB.CopyBasicFieldsToParagraphProperties_WOP(&paragraphpropertiesAPI.ParagraphProperties_WOP)
		paragraphpropertiesAPI.ParagraphPropertiesPointersEncoding = paragraphpropertiesDB.ParagraphPropertiesPointersEncoding
		paragraphpropertiesAPIs = append(paragraphpropertiesAPIs, paragraphpropertiesAPI)
	}

	c.JSON(http.StatusOK, paragraphpropertiesAPIs)
}

// PostParagraphProperties
//
// swagger:route POST /paragraphpropertiess paragraphpropertiess postParagraphProperties
//
// Creates a paragraphproperties
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostParagraphProperties(c *gin.Context) {

	mutexParagraphProperties.Lock()
	defer mutexParagraphProperties.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostParagraphPropertiess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphProperties.GetDB()

	// Validate input
	var input orm.ParagraphPropertiesAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create paragraphproperties
	paragraphpropertiesDB := orm.ParagraphPropertiesDB{}
	paragraphpropertiesDB.ParagraphPropertiesPointersEncoding = input.ParagraphPropertiesPointersEncoding
	paragraphpropertiesDB.CopyBasicFieldsFromParagraphProperties_WOP(&input.ParagraphProperties_WOP)

	_, err = db.Create(&paragraphpropertiesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoParagraphProperties.CheckoutPhaseOneInstance(&paragraphpropertiesDB)
	paragraphproperties := backRepo.BackRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr[paragraphpropertiesDB.ID]

	if paragraphproperties != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), paragraphproperties)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, paragraphpropertiesDB)
}

// GetParagraphProperties
//
// swagger:route GET /paragraphpropertiess/{ID} paragraphpropertiess getParagraphProperties
//
// Gets the details for a paragraphproperties.
//
// Responses:
// default: genericError
//
//	200: paragraphpropertiesDBResponse
func (controller *Controller) GetParagraphProperties(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetParagraphProperties", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphProperties.GetDB()

	// Get paragraphpropertiesDB in DB
	var paragraphpropertiesDB orm.ParagraphPropertiesDB
	if _, err := db.First(&paragraphpropertiesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var paragraphpropertiesAPI orm.ParagraphPropertiesAPI
	paragraphpropertiesAPI.ID = paragraphpropertiesDB.ID
	paragraphpropertiesAPI.ParagraphPropertiesPointersEncoding = paragraphpropertiesDB.ParagraphPropertiesPointersEncoding
	paragraphpropertiesDB.CopyBasicFieldsToParagraphProperties_WOP(&paragraphpropertiesAPI.ParagraphProperties_WOP)

	c.JSON(http.StatusOK, paragraphpropertiesAPI)
}

// UpdateParagraphProperties
//
// swagger:route PATCH /paragraphpropertiess/{ID} paragraphpropertiess updateParagraphProperties
//
// # Update a paragraphproperties
//
// Responses:
// default: genericError
//
//	200: paragraphpropertiesDBResponse
func (controller *Controller) UpdateParagraphProperties(c *gin.Context) {

	mutexParagraphProperties.Lock()
	defer mutexParagraphProperties.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateParagraphProperties", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphProperties.GetDB()

	// Validate input
	var input orm.ParagraphPropertiesAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var paragraphpropertiesDB orm.ParagraphPropertiesDB

	// fetch the paragraphproperties
	_, err := db.First(&paragraphpropertiesDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	paragraphpropertiesDB.CopyBasicFieldsFromParagraphProperties_WOP(&input.ParagraphProperties_WOP)
	paragraphpropertiesDB.ParagraphPropertiesPointersEncoding = input.ParagraphPropertiesPointersEncoding

	db, _ = db.Model(&paragraphpropertiesDB)
	_, err = db.Updates(&paragraphpropertiesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	paragraphpropertiesNew := new(models.ParagraphProperties)
	paragraphpropertiesDB.CopyBasicFieldsToParagraphProperties(paragraphpropertiesNew)

	// redeem pointers
	paragraphpropertiesDB.DecodePointers(backRepo, paragraphpropertiesNew)

	// get stage instance from DB instance, and call callback function
	paragraphpropertiesOld := backRepo.BackRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr[paragraphpropertiesDB.ID]
	if paragraphpropertiesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), paragraphpropertiesOld, paragraphpropertiesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the paragraphpropertiesDB
	c.JSON(http.StatusOK, paragraphpropertiesDB)
}

// DeleteParagraphProperties
//
// swagger:route DELETE /paragraphpropertiess/{ID} paragraphpropertiess deleteParagraphProperties
//
// # Delete a paragraphproperties
//
// default: genericError
//
//	200: paragraphpropertiesDBResponse
func (controller *Controller) DeleteParagraphProperties(c *gin.Context) {

	mutexParagraphProperties.Lock()
	defer mutexParagraphProperties.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteParagraphProperties", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphProperties.GetDB()

	// Get model if exist
	var paragraphpropertiesDB orm.ParagraphPropertiesDB
	if _, err := db.First(&paragraphpropertiesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&paragraphpropertiesDB)

	// get an instance (not staged) from DB instance, and call callback function
	paragraphpropertiesDeleted := new(models.ParagraphProperties)
	paragraphpropertiesDB.CopyBasicFieldsToParagraphProperties(paragraphpropertiesDeleted)

	// get stage instance from DB instance, and call callback function
	paragraphpropertiesStaged := backRepo.BackRepoParagraphProperties.Map_ParagraphPropertiesDBID_ParagraphPropertiesPtr[paragraphpropertiesDB.ID]
	if paragraphpropertiesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), paragraphpropertiesStaged, paragraphpropertiesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
