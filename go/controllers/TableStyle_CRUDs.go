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
var __TableStyle__dummysDeclaration__ models.TableStyle
var __TableStyle_time__dummyDeclaration time.Duration

var mutexTableStyle sync.Mutex

// An TableStyleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTableStyle updateTableStyle deleteTableStyle
type TableStyleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TableStyleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTableStyle updateTableStyle
type TableStyleInput struct {
	// The TableStyle to submit or modify
	// in: body
	TableStyle *orm.TableStyleAPI
}

// GetTableStyles
//
// swagger:route GET /tablestyles tablestyles getTableStyles
//
// # Get all tablestyles
//
// Responses:
// default: genericError
//
//	200: tablestyleDBResponse
func (controller *Controller) GetTableStyles(c *gin.Context) {

	// source slice
	var tablestyleDBs []orm.TableStyleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTableStyles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableStyle.GetDB()

	_, err := db.Find(&tablestyleDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tablestyleAPIs := make([]orm.TableStyleAPI, 0)

	// for each tablestyle, update fields from the database nullable fields
	for idx := range tablestyleDBs {
		tablestyleDB := &tablestyleDBs[idx]
		_ = tablestyleDB
		var tablestyleAPI orm.TableStyleAPI

		// insertion point for updating fields
		tablestyleAPI.ID = tablestyleDB.ID
		tablestyleDB.CopyBasicFieldsToTableStyle_WOP(&tablestyleAPI.TableStyle_WOP)
		tablestyleAPI.TableStylePointersEncoding = tablestyleDB.TableStylePointersEncoding
		tablestyleAPIs = append(tablestyleAPIs, tablestyleAPI)
	}

	c.JSON(http.StatusOK, tablestyleAPIs)
}

// PostTableStyle
//
// swagger:route POST /tablestyles tablestyles postTableStyle
//
// Creates a tablestyle
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTableStyle(c *gin.Context) {

	mutexTableStyle.Lock()
	defer mutexTableStyle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTableStyles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableStyle.GetDB()

	// Validate input
	var input orm.TableStyleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tablestyle
	tablestyleDB := orm.TableStyleDB{}
	tablestyleDB.TableStylePointersEncoding = input.TableStylePointersEncoding
	tablestyleDB.CopyBasicFieldsFromTableStyle_WOP(&input.TableStyle_WOP)

	_, err = db.Create(&tablestyleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTableStyle.CheckoutPhaseOneInstance(&tablestyleDB)
	tablestyle := backRepo.BackRepoTableStyle.Map_TableStyleDBID_TableStylePtr[tablestyleDB.ID]

	if tablestyle != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tablestyle)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tablestyleDB)
}

// GetTableStyle
//
// swagger:route GET /tablestyles/{ID} tablestyles getTableStyle
//
// Gets the details for a tablestyle.
//
// Responses:
// default: genericError
//
//	200: tablestyleDBResponse
func (controller *Controller) GetTableStyle(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTableStyle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableStyle.GetDB()

	// Get tablestyleDB in DB
	var tablestyleDB orm.TableStyleDB
	if _, err := db.First(&tablestyleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tablestyleAPI orm.TableStyleAPI
	tablestyleAPI.ID = tablestyleDB.ID
	tablestyleAPI.TableStylePointersEncoding = tablestyleDB.TableStylePointersEncoding
	tablestyleDB.CopyBasicFieldsToTableStyle_WOP(&tablestyleAPI.TableStyle_WOP)

	c.JSON(http.StatusOK, tablestyleAPI)
}

// UpdateTableStyle
//
// swagger:route PATCH /tablestyles/{ID} tablestyles updateTableStyle
//
// # Update a tablestyle
//
// Responses:
// default: genericError
//
//	200: tablestyleDBResponse
func (controller *Controller) UpdateTableStyle(c *gin.Context) {

	mutexTableStyle.Lock()
	defer mutexTableStyle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTableStyle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableStyle.GetDB()

	// Validate input
	var input orm.TableStyleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tablestyleDB orm.TableStyleDB

	// fetch the tablestyle
	_, err := db.First(&tablestyleDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tablestyleDB.CopyBasicFieldsFromTableStyle_WOP(&input.TableStyle_WOP)
	tablestyleDB.TableStylePointersEncoding = input.TableStylePointersEncoding

	db, _ = db.Model(&tablestyleDB)
	_, err = db.Updates(&tablestyleDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tablestyleNew := new(models.TableStyle)
	tablestyleDB.CopyBasicFieldsToTableStyle(tablestyleNew)

	// redeem pointers
	tablestyleDB.DecodePointers(backRepo, tablestyleNew)

	// get stage instance from DB instance, and call callback function
	tablestyleOld := backRepo.BackRepoTableStyle.Map_TableStyleDBID_TableStylePtr[tablestyleDB.ID]
	if tablestyleOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tablestyleOld, tablestyleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tablestyleDB
	c.JSON(http.StatusOK, tablestyleDB)
}

// DeleteTableStyle
//
// swagger:route DELETE /tablestyles/{ID} tablestyles deleteTableStyle
//
// # Delete a tablestyle
//
// default: genericError
//
//	200: tablestyleDBResponse
func (controller *Controller) DeleteTableStyle(c *gin.Context) {

	mutexTableStyle.Lock()
	defer mutexTableStyle.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTableStyle", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableStyle.GetDB()

	// Get model if exist
	var tablestyleDB orm.TableStyleDB
	if _, err := db.First(&tablestyleDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&tablestyleDB)

	// get an instance (not staged) from DB instance, and call callback function
	tablestyleDeleted := new(models.TableStyle)
	tablestyleDB.CopyBasicFieldsToTableStyle(tablestyleDeleted)

	// get stage instance from DB instance, and call callback function
	tablestyleStaged := backRepo.BackRepoTableStyle.Map_TableStyleDBID_TableStylePtr[tablestyleDB.ID]
	if tablestyleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tablestyleStaged, tablestyleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
