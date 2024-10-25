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
var __TableRow__dummysDeclaration__ models.TableRow
var __TableRow_time__dummyDeclaration time.Duration

var mutexTableRow sync.Mutex

// An TableRowID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTableRow updateTableRow deleteTableRow
type TableRowID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TableRowInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTableRow updateTableRow
type TableRowInput struct {
	// The TableRow to submit or modify
	// in: body
	TableRow *orm.TableRowAPI
}

// GetTableRows
//
// swagger:route GET /tablerows tablerows getTableRows
//
// # Get all tablerows
//
// Responses:
// default: genericError
//
//	200: tablerowDBResponse
func (controller *Controller) GetTableRows(c *gin.Context) {

	// source slice
	var tablerowDBs []orm.TableRowDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTableRows", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableRow.GetDB()

	_, err := db.Find(&tablerowDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tablerowAPIs := make([]orm.TableRowAPI, 0)

	// for each tablerow, update fields from the database nullable fields
	for idx := range tablerowDBs {
		tablerowDB := &tablerowDBs[idx]
		_ = tablerowDB
		var tablerowAPI orm.TableRowAPI

		// insertion point for updating fields
		tablerowAPI.ID = tablerowDB.ID
		tablerowDB.CopyBasicFieldsToTableRow_WOP(&tablerowAPI.TableRow_WOP)
		tablerowAPI.TableRowPointersEncoding = tablerowDB.TableRowPointersEncoding
		tablerowAPIs = append(tablerowAPIs, tablerowAPI)
	}

	c.JSON(http.StatusOK, tablerowAPIs)
}

// PostTableRow
//
// swagger:route POST /tablerows tablerows postTableRow
//
// Creates a tablerow
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTableRow(c *gin.Context) {

	mutexTableRow.Lock()
	defer mutexTableRow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTableRows", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableRow.GetDB()

	// Validate input
	var input orm.TableRowAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tablerow
	tablerowDB := orm.TableRowDB{}
	tablerowDB.TableRowPointersEncoding = input.TableRowPointersEncoding
	tablerowDB.CopyBasicFieldsFromTableRow_WOP(&input.TableRow_WOP)

	_, err = db.Create(&tablerowDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTableRow.CheckoutPhaseOneInstance(&tablerowDB)
	tablerow := backRepo.BackRepoTableRow.Map_TableRowDBID_TableRowPtr[tablerowDB.ID]

	if tablerow != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tablerow)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tablerowDB)
}

// GetTableRow
//
// swagger:route GET /tablerows/{ID} tablerows getTableRow
//
// Gets the details for a tablerow.
//
// Responses:
// default: genericError
//
//	200: tablerowDBResponse
func (controller *Controller) GetTableRow(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTableRow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableRow.GetDB()

	// Get tablerowDB in DB
	var tablerowDB orm.TableRowDB
	if _, err := db.First(&tablerowDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tablerowAPI orm.TableRowAPI
	tablerowAPI.ID = tablerowDB.ID
	tablerowAPI.TableRowPointersEncoding = tablerowDB.TableRowPointersEncoding
	tablerowDB.CopyBasicFieldsToTableRow_WOP(&tablerowAPI.TableRow_WOP)

	c.JSON(http.StatusOK, tablerowAPI)
}

// UpdateTableRow
//
// swagger:route PATCH /tablerows/{ID} tablerows updateTableRow
//
// # Update a tablerow
//
// Responses:
// default: genericError
//
//	200: tablerowDBResponse
func (controller *Controller) UpdateTableRow(c *gin.Context) {

	mutexTableRow.Lock()
	defer mutexTableRow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTableRow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableRow.GetDB()

	// Validate input
	var input orm.TableRowAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tablerowDB orm.TableRowDB

	// fetch the tablerow
	_, err := db.First(&tablerowDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tablerowDB.CopyBasicFieldsFromTableRow_WOP(&input.TableRow_WOP)
	tablerowDB.TableRowPointersEncoding = input.TableRowPointersEncoding

	db, _ = db.Model(&tablerowDB)
	_, err = db.Updates(&tablerowDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tablerowNew := new(models.TableRow)
	tablerowDB.CopyBasicFieldsToTableRow(tablerowNew)

	// redeem pointers
	tablerowDB.DecodePointers(backRepo, tablerowNew)

	// get stage instance from DB instance, and call callback function
	tablerowOld := backRepo.BackRepoTableRow.Map_TableRowDBID_TableRowPtr[tablerowDB.ID]
	if tablerowOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tablerowOld, tablerowNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tablerowDB
	c.JSON(http.StatusOK, tablerowDB)
}

// DeleteTableRow
//
// swagger:route DELETE /tablerows/{ID} tablerows deleteTableRow
//
// # Delete a tablerow
//
// default: genericError
//
//	200: tablerowDBResponse
func (controller *Controller) DeleteTableRow(c *gin.Context) {

	mutexTableRow.Lock()
	defer mutexTableRow.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTableRow", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableRow.GetDB()

	// Get model if exist
	var tablerowDB orm.TableRowDB
	if _, err := db.First(&tablerowDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&tablerowDB)

	// get an instance (not staged) from DB instance, and call callback function
	tablerowDeleted := new(models.TableRow)
	tablerowDB.CopyBasicFieldsToTableRow(tablerowDeleted)

	// get stage instance from DB instance, and call callback function
	tablerowStaged := backRepo.BackRepoTableRow.Map_TableRowDBID_TableRowPtr[tablerowDB.ID]
	if tablerowStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tablerowStaged, tablerowDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
