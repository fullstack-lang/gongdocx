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
var __Rune__dummysDeclaration__ models.Rune
var __Rune_time__dummyDeclaration time.Duration

var mutexRune sync.Mutex

// An RuneID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRune updateRune deleteRune
type RuneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RuneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRune updateRune
type RuneInput struct {
	// The Rune to submit or modify
	// in: body
	Rune *orm.RuneAPI
}

// GetRunes
//
// swagger:route GET /runes runes getRunes
//
// # Get all runes
//
// Responses:
// default: genericError
//
//	200: runeDBResponse
func (controller *Controller) GetRunes(c *gin.Context) {

	// source slice
	var runeDBs []orm.RuneDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRunes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRune.GetDB()

	query := db.Find(&runeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	runeAPIs := make([]orm.RuneAPI, 0)

	// for each rune, update fields from the database nullable fields
	for idx := range runeDBs {
		runeDB := &runeDBs[idx]
		_ = runeDB
		var runeAPI orm.RuneAPI

		// insertion point for updating fields
		runeAPI.ID = runeDB.ID
		runeDB.CopyBasicFieldsToRune_WOP(&runeAPI.Rune_WOP)
		runeAPI.RunePointersEncoding = runeDB.RunePointersEncoding
		runeAPIs = append(runeAPIs, runeAPI)
	}

	c.JSON(http.StatusOK, runeAPIs)
}

// PostRune
//
// swagger:route POST /runes runes postRune
//
// Creates a rune
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRune(c *gin.Context) {

	mutexRune.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRunes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRune.GetDB()

	// Validate input
	var input orm.RuneAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rune
	runeDB := orm.RuneDB{}
	runeDB.RunePointersEncoding = input.RunePointersEncoding
	runeDB.CopyBasicFieldsFromRune_WOP(&input.Rune_WOP)

	query := db.Create(&runeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRune.CheckoutPhaseOneInstance(&runeDB)
	rune := backRepo.BackRepoRune.Map_RuneDBID_RunePtr[runeDB.ID]

	if rune != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rune)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, runeDB)

	mutexRune.Unlock()
}

// GetRune
//
// swagger:route GET /runes/{ID} runes getRune
//
// Gets the details for a rune.
//
// Responses:
// default: genericError
//
//	200: runeDBResponse
func (controller *Controller) GetRune(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRune", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRune.GetDB()

	// Get runeDB in DB
	var runeDB orm.RuneDB
	if err := db.First(&runeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var runeAPI orm.RuneAPI
	runeAPI.ID = runeDB.ID
	runeAPI.RunePointersEncoding = runeDB.RunePointersEncoding
	runeDB.CopyBasicFieldsToRune_WOP(&runeAPI.Rune_WOP)

	c.JSON(http.StatusOK, runeAPI)
}

// UpdateRune
//
// swagger:route PATCH /runes/{ID} runes updateRune
//
// # Update a rune
//
// Responses:
// default: genericError
//
//	200: runeDBResponse
func (controller *Controller) UpdateRune(c *gin.Context) {

	mutexRune.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRune", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRune.GetDB()

	// Validate input
	var input orm.RuneAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var runeDB orm.RuneDB

	// fetch the rune
	query := db.First(&runeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	runeDB.CopyBasicFieldsFromRune_WOP(&input.Rune_WOP)
	runeDB.RunePointersEncoding = input.RunePointersEncoding

	query = db.Model(&runeDB).Updates(runeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	runeNew := new(models.Rune)
	runeDB.CopyBasicFieldsToRune(runeNew)

	// redeem pointers
	runeDB.DecodePointers(backRepo, runeNew)

	// get stage instance from DB instance, and call callback function
	runeOld := backRepo.BackRepoRune.Map_RuneDBID_RunePtr[runeDB.ID]
	if runeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), runeOld, runeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the runeDB
	c.JSON(http.StatusOK, runeDB)

	mutexRune.Unlock()
}

// DeleteRune
//
// swagger:route DELETE /runes/{ID} runes deleteRune
//
// # Delete a rune
//
// default: genericError
//
//	200: runeDBResponse
func (controller *Controller) DeleteRune(c *gin.Context) {

	mutexRune.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRune", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRune.GetDB()

	// Get model if exist
	var runeDB orm.RuneDB
	if err := db.First(&runeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&runeDB)

	// get an instance (not staged) from DB instance, and call callback function
	runeDeleted := new(models.Rune)
	runeDB.CopyBasicFieldsToRune(runeDeleted)

	// get stage instance from DB instance, and call callback function
	runeStaged := backRepo.BackRepoRune.Map_RuneDBID_RunePtr[runeDB.ID]
	if runeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), runeStaged, runeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexRune.Unlock()
}
