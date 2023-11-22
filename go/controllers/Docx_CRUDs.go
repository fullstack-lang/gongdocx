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
var __Docx__dummysDeclaration__ models.Docx
var __Docx_time__dummyDeclaration time.Duration

var mutexDocx sync.Mutex

// An DocxID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDocx updateDocx deleteDocx
type DocxID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DocxInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDocx updateDocx
type DocxInput struct {
	// The Docx to submit or modify
	// in: body
	Docx *orm.DocxAPI
}

// GetDocxs
//
// swagger:route GET /docxs docxs getDocxs
//
// # Get all docxs
//
// Responses:
// default: genericError
//
//	200: docxDBResponse
func (controller *Controller) GetDocxs(c *gin.Context) {

	// source slice
	var docxDBs []orm.DocxDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDocxs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocx.GetDB()

	query := db.Find(&docxDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	docxAPIs := make([]orm.DocxAPI, 0)

	// for each docx, update fields from the database nullable fields
	for idx := range docxDBs {
		docxDB := &docxDBs[idx]
		_ = docxDB
		var docxAPI orm.DocxAPI

		// insertion point for updating fields
		docxAPI.ID = docxDB.ID
		docxDB.CopyBasicFieldsToDocx_WOP(&docxAPI.Docx_WOP)
		docxAPI.DocxPointersEncoding = docxDB.DocxPointersEncoding
		docxAPIs = append(docxAPIs, docxAPI)
	}

	c.JSON(http.StatusOK, docxAPIs)
}

// PostDocx
//
// swagger:route POST /docxs docxs postDocx
//
// Creates a docx
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDocx(c *gin.Context) {

	mutexDocx.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDocxs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocx.GetDB()

	// Validate input
	var input orm.DocxAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create docx
	docxDB := orm.DocxDB{}
	docxDB.DocxPointersEncoding = input.DocxPointersEncoding
	docxDB.CopyBasicFieldsFromDocx_WOP(&input.Docx_WOP)

	query := db.Create(&docxDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDocx.CheckoutPhaseOneInstance(&docxDB)
	docx := backRepo.BackRepoDocx.Map_DocxDBID_DocxPtr[docxDB.ID]

	if docx != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), docx)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, docxDB)

	mutexDocx.Unlock()
}

// GetDocx
//
// swagger:route GET /docxs/{ID} docxs getDocx
//
// Gets the details for a docx.
//
// Responses:
// default: genericError
//
//	200: docxDBResponse
func (controller *Controller) GetDocx(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDocx", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocx.GetDB()

	// Get docxDB in DB
	var docxDB orm.DocxDB
	if err := db.First(&docxDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var docxAPI orm.DocxAPI
	docxAPI.ID = docxDB.ID
	docxAPI.DocxPointersEncoding = docxDB.DocxPointersEncoding
	docxDB.CopyBasicFieldsToDocx_WOP(&docxAPI.Docx_WOP)

	c.JSON(http.StatusOK, docxAPI)
}

// UpdateDocx
//
// swagger:route PATCH /docxs/{ID} docxs updateDocx
//
// # Update a docx
//
// Responses:
// default: genericError
//
//	200: docxDBResponse
func (controller *Controller) UpdateDocx(c *gin.Context) {

	mutexDocx.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDocx", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocx.GetDB()

	// Validate input
	var input orm.DocxAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var docxDB orm.DocxDB

	// fetch the docx
	query := db.First(&docxDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	docxDB.CopyBasicFieldsFromDocx_WOP(&input.Docx_WOP)
	docxDB.DocxPointersEncoding = input.DocxPointersEncoding

	query = db.Model(&docxDB).Updates(docxDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	docxNew := new(models.Docx)
	docxDB.CopyBasicFieldsToDocx(docxNew)

	// redeem pointers
	docxDB.DecodePointers(backRepo, docxNew)

	// get stage instance from DB instance, and call callback function
	docxOld := backRepo.BackRepoDocx.Map_DocxDBID_DocxPtr[docxDB.ID]
	if docxOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), docxOld, docxNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the docxDB
	c.JSON(http.StatusOK, docxDB)

	mutexDocx.Unlock()
}

// DeleteDocx
//
// swagger:route DELETE /docxs/{ID} docxs deleteDocx
//
// # Delete a docx
//
// default: genericError
//
//	200: docxDBResponse
func (controller *Controller) DeleteDocx(c *gin.Context) {

	mutexDocx.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDocx", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocx.GetDB()

	// Get model if exist
	var docxDB orm.DocxDB
	if err := db.First(&docxDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&docxDB)

	// get an instance (not staged) from DB instance, and call callback function
	docxDeleted := new(models.Docx)
	docxDB.CopyBasicFieldsToDocx(docxDeleted)

	// get stage instance from DB instance, and call callback function
	docxStaged := backRepo.BackRepoDocx.Map_DocxDBID_DocxPtr[docxDB.ID]
	if docxStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), docxStaged, docxDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexDocx.Unlock()
}
