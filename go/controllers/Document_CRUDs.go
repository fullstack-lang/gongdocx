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
var __Document__dummysDeclaration__ models.Document
var __Document_time__dummyDeclaration time.Duration

var mutexDocument sync.Mutex

// An DocumentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDocument updateDocument deleteDocument
type DocumentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DocumentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDocument updateDocument
type DocumentInput struct {
	// The Document to submit or modify
	// in: body
	Document *orm.DocumentAPI
}

// GetDocuments
//
// swagger:route GET /documents documents getDocuments
//
// # Get all documents
//
// Responses:
// default: genericError
//
//	200: documentDBResponse
func (controller *Controller) GetDocuments(c *gin.Context) {

	// source slice
	var documentDBs []orm.DocumentDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDocuments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocument.GetDB()

	query := db.Find(&documentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	documentAPIs := make([]orm.DocumentAPI, 0)

	// for each document, update fields from the database nullable fields
	for idx := range documentDBs {
		documentDB := &documentDBs[idx]
		_ = documentDB
		var documentAPI orm.DocumentAPI

		// insertion point for updating fields
		documentAPI.ID = documentDB.ID
		documentDB.CopyBasicFieldsToDocument_WOP(&documentAPI.Document_WOP)
		documentAPI.DocumentPointersEncoding = documentDB.DocumentPointersEncoding
		documentAPIs = append(documentAPIs, documentAPI)
	}

	c.JSON(http.StatusOK, documentAPIs)
}

// PostDocument
//
// swagger:route POST /documents documents postDocument
//
// Creates a document
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDocument(c *gin.Context) {

	mutexDocument.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDocuments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocument.GetDB()

	// Validate input
	var input orm.DocumentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create document
	documentDB := orm.DocumentDB{}
	documentDB.DocumentPointersEncoding = input.DocumentPointersEncoding
	documentDB.CopyBasicFieldsFromDocument_WOP(&input.Document_WOP)

	query := db.Create(&documentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDocument.CheckoutPhaseOneInstance(&documentDB)
	document := backRepo.BackRepoDocument.Map_DocumentDBID_DocumentPtr[documentDB.ID]

	if document != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), document)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, documentDB)

	mutexDocument.Unlock()
}

// GetDocument
//
// swagger:route GET /documents/{ID} documents getDocument
//
// Gets the details for a document.
//
// Responses:
// default: genericError
//
//	200: documentDBResponse
func (controller *Controller) GetDocument(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDocument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocument.GetDB()

	// Get documentDB in DB
	var documentDB orm.DocumentDB
	if err := db.First(&documentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var documentAPI orm.DocumentAPI
	documentAPI.ID = documentDB.ID
	documentAPI.DocumentPointersEncoding = documentDB.DocumentPointersEncoding
	documentDB.CopyBasicFieldsToDocument_WOP(&documentAPI.Document_WOP)

	c.JSON(http.StatusOK, documentAPI)
}

// UpdateDocument
//
// swagger:route PATCH /documents/{ID} documents updateDocument
//
// # Update a document
//
// Responses:
// default: genericError
//
//	200: documentDBResponse
func (controller *Controller) UpdateDocument(c *gin.Context) {

	mutexDocument.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDocument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocument.GetDB()

	// Validate input
	var input orm.DocumentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var documentDB orm.DocumentDB

	// fetch the document
	query := db.First(&documentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	documentDB.CopyBasicFieldsFromDocument_WOP(&input.Document_WOP)
	documentDB.DocumentPointersEncoding = input.DocumentPointersEncoding

	query = db.Model(&documentDB).Updates(documentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	documentNew := new(models.Document)
	documentDB.CopyBasicFieldsToDocument(documentNew)

	// get stage instance from DB instance, and call callback function
	documentOld := backRepo.BackRepoDocument.Map_DocumentDBID_DocumentPtr[documentDB.ID]
	if documentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), documentOld, documentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the documentDB
	c.JSON(http.StatusOK, documentDB)

	mutexDocument.Unlock()
}

// DeleteDocument
//
// swagger:route DELETE /documents/{ID} documents deleteDocument
//
// # Delete a document
//
// default: genericError
//
//	200: documentDBResponse
func (controller *Controller) DeleteDocument(c *gin.Context) {

	mutexDocument.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDocument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDocument.GetDB()

	// Get model if exist
	var documentDB orm.DocumentDB
	if err := db.First(&documentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&documentDB)

	// get an instance (not staged) from DB instance, and call callback function
	documentDeleted := new(models.Document)
	documentDB.CopyBasicFieldsToDocument(documentDeleted)

	// get stage instance from DB instance, and call callback function
	documentStaged := backRepo.BackRepoDocument.Map_DocumentDBID_DocumentPtr[documentDB.ID]
	if documentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), documentStaged, documentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexDocument.Unlock()
}
