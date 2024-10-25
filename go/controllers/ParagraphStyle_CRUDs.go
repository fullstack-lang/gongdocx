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
var __ParagraphStyle__dummysDeclaration__ models.ParagraphStyle
var __ParagraphStyle_time__dummyDeclaration time.Duration

var mutexParagraphStyle sync.Mutex

// An ParagraphStyleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getParagraphStyle updateParagraphStyle deleteParagraphStyle
type ParagraphStyleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ParagraphStyleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postParagraphStyle updateParagraphStyle
type ParagraphStyleInput struct {
	// The ParagraphStyle to submit or modify
	// in: body
	ParagraphStyle *orm.ParagraphStyleAPI
}

// GetParagraphStyles
//
// swagger:route GET /paragraphstyles paragraphstyles getParagraphStyles
//
// # Get all paragraphstyles
//
// Responses:
// default: genericError
//
//	200: paragraphstyleDBResponse
func (controller *Controller) GetParagraphStyles(c *gin.Context) {

	// source slice
	var paragraphstyleDBs []orm.ParagraphStyleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetParagraphStyles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphStyle.GetDB()

	_, err := db.Find(&paragraphstyleDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	paragraphstyleAPIs := make([]orm.ParagraphStyleAPI, 0)

	// for each paragraphstyle, update fields from the database nullable fields
	for idx := range paragraphstyleDBs {
		paragraphstyleDB := &paragraphstyleDBs[idx]
		_ = paragraphstyleDB
		var paragraphstyleAPI orm.ParagraphStyleAPI

		// insertion point for updating fields
		paragraphstyleAPI.ID = paragraphstyleDB.ID
		paragraphstyleDB.CopyBasicFieldsToParagraphStyle_WOP(&paragraphstyleAPI.ParagraphStyle_WOP)
		paragraphstyleAPI.ParagraphStylePointersEncoding = paragraphstyleDB.ParagraphStylePointersEncoding
		paragraphstyleAPIs = append(paragraphstyleAPIs, paragraphstyleAPI)
	}

	c.JSON(http.StatusOK, paragraphstyleAPIs)
}

// PostParagraphStyle
//
// swagger:route POST /paragraphstyles paragraphstyles postParagraphStyle
//
// Creates a paragraphstyle
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostParagraphStyle(c *gin.Context) {

	mutexParagraphStyle.Lock()
	defer mutexParagraphStyle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostParagraphStyles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphStyle.GetDB()

	// Validate input
	var input orm.ParagraphStyleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create paragraphstyle
	paragraphstyleDB := orm.ParagraphStyleDB{}
	paragraphstyleDB.ParagraphStylePointersEncoding = input.ParagraphStylePointersEncoding
	paragraphstyleDB.CopyBasicFieldsFromParagraphStyle_WOP(&input.ParagraphStyle_WOP)

	_, err = db.Create(&paragraphstyleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoParagraphStyle.CheckoutPhaseOneInstance(&paragraphstyleDB)
	paragraphstyle := backRepo.BackRepoParagraphStyle.Map_ParagraphStyleDBID_ParagraphStylePtr[paragraphstyleDB.ID]

	if paragraphstyle != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), paragraphstyle)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, paragraphstyleDB)
}

// GetParagraphStyle
//
// swagger:route GET /paragraphstyles/{ID} paragraphstyles getParagraphStyle
//
// Gets the details for a paragraphstyle.
//
// Responses:
// default: genericError
//
//	200: paragraphstyleDBResponse
func (controller *Controller) GetParagraphStyle(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetParagraphStyle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphStyle.GetDB()

	// Get paragraphstyleDB in DB
	var paragraphstyleDB orm.ParagraphStyleDB
	if _, err := db.First(&paragraphstyleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var paragraphstyleAPI orm.ParagraphStyleAPI
	paragraphstyleAPI.ID = paragraphstyleDB.ID
	paragraphstyleAPI.ParagraphStylePointersEncoding = paragraphstyleDB.ParagraphStylePointersEncoding
	paragraphstyleDB.CopyBasicFieldsToParagraphStyle_WOP(&paragraphstyleAPI.ParagraphStyle_WOP)

	c.JSON(http.StatusOK, paragraphstyleAPI)
}

// UpdateParagraphStyle
//
// swagger:route PATCH /paragraphstyles/{ID} paragraphstyles updateParagraphStyle
//
// # Update a paragraphstyle
//
// Responses:
// default: genericError
//
//	200: paragraphstyleDBResponse
func (controller *Controller) UpdateParagraphStyle(c *gin.Context) {

	mutexParagraphStyle.Lock()
	defer mutexParagraphStyle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateParagraphStyle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphStyle.GetDB()

	// Validate input
	var input orm.ParagraphStyleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var paragraphstyleDB orm.ParagraphStyleDB

	// fetch the paragraphstyle
	_, err := db.First(&paragraphstyleDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	paragraphstyleDB.CopyBasicFieldsFromParagraphStyle_WOP(&input.ParagraphStyle_WOP)
	paragraphstyleDB.ParagraphStylePointersEncoding = input.ParagraphStylePointersEncoding

	db, _ = db.Model(&paragraphstyleDB)
	_, err = db.Updates(&paragraphstyleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	paragraphstyleNew := new(models.ParagraphStyle)
	paragraphstyleDB.CopyBasicFieldsToParagraphStyle(paragraphstyleNew)

	// redeem pointers
	paragraphstyleDB.DecodePointers(backRepo, paragraphstyleNew)

	// get stage instance from DB instance, and call callback function
	paragraphstyleOld := backRepo.BackRepoParagraphStyle.Map_ParagraphStyleDBID_ParagraphStylePtr[paragraphstyleDB.ID]
	if paragraphstyleOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), paragraphstyleOld, paragraphstyleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the paragraphstyleDB
	c.JSON(http.StatusOK, paragraphstyleDB)
}

// DeleteParagraphStyle
//
// swagger:route DELETE /paragraphstyles/{ID} paragraphstyles deleteParagraphStyle
//
// # Delete a paragraphstyle
//
// default: genericError
//
//	200: paragraphstyleDBResponse
func (controller *Controller) DeleteParagraphStyle(c *gin.Context) {

	mutexParagraphStyle.Lock()
	defer mutexParagraphStyle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteParagraphStyle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraphStyle.GetDB()

	// Get model if exist
	var paragraphstyleDB orm.ParagraphStyleDB
	if _, err := db.First(&paragraphstyleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&paragraphstyleDB)

	// get an instance (not staged) from DB instance, and call callback function
	paragraphstyleDeleted := new(models.ParagraphStyle)
	paragraphstyleDB.CopyBasicFieldsToParagraphStyle(paragraphstyleDeleted)

	// get stage instance from DB instance, and call callback function
	paragraphstyleStaged := backRepo.BackRepoParagraphStyle.Map_ParagraphStyleDBID_ParagraphStylePtr[paragraphstyleDB.ID]
	if paragraphstyleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), paragraphstyleStaged, paragraphstyleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
