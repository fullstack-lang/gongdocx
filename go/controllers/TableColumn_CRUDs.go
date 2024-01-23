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
var __TableColumn__dummysDeclaration__ models.TableColumn
var __TableColumn_time__dummyDeclaration time.Duration

var mutexTableColumn sync.Mutex

// An TableColumnID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTableColumn updateTableColumn deleteTableColumn
type TableColumnID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TableColumnInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTableColumn updateTableColumn
type TableColumnInput struct {
	// The TableColumn to submit or modify
	// in: body
	TableColumn *orm.TableColumnAPI
}

// GetTableColumns
//
// swagger:route GET /tablecolumns tablecolumns getTableColumns
//
// # Get all tablecolumns
//
// Responses:
// default: genericError
//
//	200: tablecolumnDBResponse
func (controller *Controller) GetTableColumns(c *gin.Context) {

	// source slice
	var tablecolumnDBs []orm.TableColumnDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTableColumns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableColumn.GetDB()

	query := db.Find(&tablecolumnDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tablecolumnAPIs := make([]orm.TableColumnAPI, 0)

	// for each tablecolumn, update fields from the database nullable fields
	for idx := range tablecolumnDBs {
		tablecolumnDB := &tablecolumnDBs[idx]
		_ = tablecolumnDB
		var tablecolumnAPI orm.TableColumnAPI

		// insertion point for updating fields
		tablecolumnAPI.ID = tablecolumnDB.ID
		tablecolumnDB.CopyBasicFieldsToTableColumn_WOP(&tablecolumnAPI.TableColumn_WOP)
		tablecolumnAPI.TableColumnPointersEncoding = tablecolumnDB.TableColumnPointersEncoding
		tablecolumnAPIs = append(tablecolumnAPIs, tablecolumnAPI)
	}

	c.JSON(http.StatusOK, tablecolumnAPIs)
}

// PostTableColumn
//
// swagger:route POST /tablecolumns tablecolumns postTableColumn
//
// Creates a tablecolumn
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTableColumn(c *gin.Context) {

	mutexTableColumn.Lock()
	defer mutexTableColumn.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTableColumns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableColumn.GetDB()

	// Validate input
	var input orm.TableColumnAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tablecolumn
	tablecolumnDB := orm.TableColumnDB{}
	tablecolumnDB.TableColumnPointersEncoding = input.TableColumnPointersEncoding
	tablecolumnDB.CopyBasicFieldsFromTableColumn_WOP(&input.TableColumn_WOP)

	query := db.Create(&tablecolumnDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTableColumn.CheckoutPhaseOneInstance(&tablecolumnDB)
	tablecolumn := backRepo.BackRepoTableColumn.Map_TableColumnDBID_TableColumnPtr[tablecolumnDB.ID]

	if tablecolumn != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tablecolumn)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tablecolumnDB)
}

// GetTableColumn
//
// swagger:route GET /tablecolumns/{ID} tablecolumns getTableColumn
//
// Gets the details for a tablecolumn.
//
// Responses:
// default: genericError
//
//	200: tablecolumnDBResponse
func (controller *Controller) GetTableColumn(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTableColumn", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableColumn.GetDB()

	// Get tablecolumnDB in DB
	var tablecolumnDB orm.TableColumnDB
	if err := db.First(&tablecolumnDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tablecolumnAPI orm.TableColumnAPI
	tablecolumnAPI.ID = tablecolumnDB.ID
	tablecolumnAPI.TableColumnPointersEncoding = tablecolumnDB.TableColumnPointersEncoding
	tablecolumnDB.CopyBasicFieldsToTableColumn_WOP(&tablecolumnAPI.TableColumn_WOP)

	c.JSON(http.StatusOK, tablecolumnAPI)
}

// UpdateTableColumn
//
// swagger:route PATCH /tablecolumns/{ID} tablecolumns updateTableColumn
//
// # Update a tablecolumn
//
// Responses:
// default: genericError
//
//	200: tablecolumnDBResponse
func (controller *Controller) UpdateTableColumn(c *gin.Context) {

	mutexTableColumn.Lock()
	defer mutexTableColumn.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTableColumn", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableColumn.GetDB()

	// Validate input
	var input orm.TableColumnAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tablecolumnDB orm.TableColumnDB

	// fetch the tablecolumn
	query := db.First(&tablecolumnDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tablecolumnDB.CopyBasicFieldsFromTableColumn_WOP(&input.TableColumn_WOP)
	tablecolumnDB.TableColumnPointersEncoding = input.TableColumnPointersEncoding

	query = db.Model(&tablecolumnDB).Updates(tablecolumnDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tablecolumnNew := new(models.TableColumn)
	tablecolumnDB.CopyBasicFieldsToTableColumn(tablecolumnNew)

	// redeem pointers
	tablecolumnDB.DecodePointers(backRepo, tablecolumnNew)

	// get stage instance from DB instance, and call callback function
	tablecolumnOld := backRepo.BackRepoTableColumn.Map_TableColumnDBID_TableColumnPtr[tablecolumnDB.ID]
	if tablecolumnOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tablecolumnOld, tablecolumnNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tablecolumnDB
	c.JSON(http.StatusOK, tablecolumnDB)
}

// DeleteTableColumn
//
// swagger:route DELETE /tablecolumns/{ID} tablecolumns deleteTableColumn
//
// # Delete a tablecolumn
//
// default: genericError
//
//	200: tablecolumnDBResponse
func (controller *Controller) DeleteTableColumn(c *gin.Context) {

	mutexTableColumn.Lock()
	defer mutexTableColumn.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTableColumn", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTableColumn.GetDB()

	// Get model if exist
	var tablecolumnDB orm.TableColumnDB
	if err := db.First(&tablecolumnDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tablecolumnDB)

	// get an instance (not staged) from DB instance, and call callback function
	tablecolumnDeleted := new(models.TableColumn)
	tablecolumnDB.CopyBasicFieldsToTableColumn(tablecolumnDeleted)

	// get stage instance from DB instance, and call callback function
	tablecolumnStaged := backRepo.BackRepoTableColumn.Map_TableColumnDBID_TableColumnPtr[tablecolumnDB.ID]
	if tablecolumnStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tablecolumnStaged, tablecolumnDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
