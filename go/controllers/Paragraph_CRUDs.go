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
var __Paragraph__dummysDeclaration__ models.Paragraph
var __Paragraph_time__dummyDeclaration time.Duration

var mutexParagraph sync.Mutex

// An ParagraphID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getParagraph updateParagraph deleteParagraph
type ParagraphID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ParagraphInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postParagraph updateParagraph
type ParagraphInput struct {
	// The Paragraph to submit or modify
	// in: body
	Paragraph *orm.ParagraphAPI
}

// GetParagraphs
//
// swagger:route GET /paragraphs paragraphs getParagraphs
//
// # Get all paragraphs
//
// Responses:
// default: genericError
//
//	200: paragraphDBResponse
func (controller *Controller) GetParagraphs(c *gin.Context) {

	// source slice
	var paragraphDBs []orm.ParagraphDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetParagraphs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraph.GetDB()

	query := db.Find(&paragraphDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	paragraphAPIs := make([]orm.ParagraphAPI, 0)

	// for each paragraph, update fields from the database nullable fields
	for idx := range paragraphDBs {
		paragraphDB := &paragraphDBs[idx]
		_ = paragraphDB
		var paragraphAPI orm.ParagraphAPI

		// insertion point for updating fields
		paragraphAPI.ID = paragraphDB.ID
		paragraphDB.CopyBasicFieldsToParagraph_WOP(&paragraphAPI.Paragraph_WOP)
		paragraphAPI.ParagraphPointersEncoding = paragraphDB.ParagraphPointersEncoding
		paragraphAPIs = append(paragraphAPIs, paragraphAPI)
	}

	c.JSON(http.StatusOK, paragraphAPIs)
}

// PostParagraph
//
// swagger:route POST /paragraphs paragraphs postParagraph
//
// Creates a paragraph
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostParagraph(c *gin.Context) {

	mutexParagraph.Lock()
	defer mutexParagraph.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostParagraphs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraph.GetDB()

	// Validate input
	var input orm.ParagraphAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create paragraph
	paragraphDB := orm.ParagraphDB{}
	paragraphDB.ParagraphPointersEncoding = input.ParagraphPointersEncoding
	paragraphDB.CopyBasicFieldsFromParagraph_WOP(&input.Paragraph_WOP)

	query := db.Create(&paragraphDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoParagraph.CheckoutPhaseOneInstance(&paragraphDB)
	paragraph := backRepo.BackRepoParagraph.Map_ParagraphDBID_ParagraphPtr[paragraphDB.ID]

	if paragraph != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), paragraph)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, paragraphDB)
}

// GetParagraph
//
// swagger:route GET /paragraphs/{ID} paragraphs getParagraph
//
// Gets the details for a paragraph.
//
// Responses:
// default: genericError
//
//	200: paragraphDBResponse
func (controller *Controller) GetParagraph(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetParagraph", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraph.GetDB()

	// Get paragraphDB in DB
	var paragraphDB orm.ParagraphDB
	if err := db.First(&paragraphDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var paragraphAPI orm.ParagraphAPI
	paragraphAPI.ID = paragraphDB.ID
	paragraphAPI.ParagraphPointersEncoding = paragraphDB.ParagraphPointersEncoding
	paragraphDB.CopyBasicFieldsToParagraph_WOP(&paragraphAPI.Paragraph_WOP)

	c.JSON(http.StatusOK, paragraphAPI)
}

// UpdateParagraph
//
// swagger:route PATCH /paragraphs/{ID} paragraphs updateParagraph
//
// # Update a paragraph
//
// Responses:
// default: genericError
//
//	200: paragraphDBResponse
func (controller *Controller) UpdateParagraph(c *gin.Context) {

	mutexParagraph.Lock()
	defer mutexParagraph.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateParagraph", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraph.GetDB()

	// Validate input
	var input orm.ParagraphAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var paragraphDB orm.ParagraphDB

	// fetch the paragraph
	query := db.First(&paragraphDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	paragraphDB.CopyBasicFieldsFromParagraph_WOP(&input.Paragraph_WOP)
	paragraphDB.ParagraphPointersEncoding = input.ParagraphPointersEncoding

	query = db.Model(&paragraphDB).Updates(paragraphDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	paragraphNew := new(models.Paragraph)
	paragraphDB.CopyBasicFieldsToParagraph(paragraphNew)

	// redeem pointers
	paragraphDB.DecodePointers(backRepo, paragraphNew)

	// get stage instance from DB instance, and call callback function
	paragraphOld := backRepo.BackRepoParagraph.Map_ParagraphDBID_ParagraphPtr[paragraphDB.ID]
	if paragraphOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), paragraphOld, paragraphNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the paragraphDB
	c.JSON(http.StatusOK, paragraphDB)
}

// DeleteParagraph
//
// swagger:route DELETE /paragraphs/{ID} paragraphs deleteParagraph
//
// # Delete a paragraph
//
// default: genericError
//
//	200: paragraphDBResponse
func (controller *Controller) DeleteParagraph(c *gin.Context) {

	mutexParagraph.Lock()
	defer mutexParagraph.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteParagraph", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoParagraph.GetDB()

	// Get model if exist
	var paragraphDB orm.ParagraphDB
	if err := db.First(&paragraphDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&paragraphDB)

	// get an instance (not staged) from DB instance, and call callback function
	paragraphDeleted := new(models.Paragraph)
	paragraphDB.CopyBasicFieldsToParagraph(paragraphDeleted)

	// get stage instance from DB instance, and call callback function
	paragraphStaged := backRepo.BackRepoParagraph.Map_ParagraphDBID_ParagraphPtr[paragraphDB.ID]
	if paragraphStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), paragraphStaged, paragraphDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
