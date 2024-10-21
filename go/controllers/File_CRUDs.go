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
var __File__dummysDeclaration__ models.File
var __File_time__dummyDeclaration time.Duration

var mutexFile sync.Mutex

// An FileID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFile updateFile deleteFile
type FileID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FileInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFile updateFile
type FileInput struct {
	// The File to submit or modify
	// in: body
	File *orm.FileAPI
}

// GetFiles
//
// swagger:route GET /files files getFiles
//
// # Get all files
//
// Responses:
// default: genericError
//
//	200: fileDBResponse
func (controller *Controller) GetFiles(c *gin.Context) {

	// source slice
	var fileDBs []orm.FileDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFiles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFile.GetDB()

	_, err := db.Find(&fileDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	fileAPIs := make([]orm.FileAPI, 0)

	// for each file, update fields from the database nullable fields
	for idx := range fileDBs {
		fileDB := &fileDBs[idx]
		_ = fileDB
		var fileAPI orm.FileAPI

		// insertion point for updating fields
		fileAPI.ID = fileDB.ID
		fileDB.CopyBasicFieldsToFile_WOP(&fileAPI.File_WOP)
		fileAPI.FilePointersEncoding = fileDB.FilePointersEncoding
		fileAPIs = append(fileAPIs, fileAPI)
	}

	c.JSON(http.StatusOK, fileAPIs)
}

// PostFile
//
// swagger:route POST /files files postFile
//
// Creates a file
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFile(c *gin.Context) {

	mutexFile.Lock()
	defer mutexFile.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFiles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFile.GetDB()

	// Validate input
	var input orm.FileAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create file
	fileDB := orm.FileDB{}
	fileDB.FilePointersEncoding = input.FilePointersEncoding
	fileDB.CopyBasicFieldsFromFile_WOP(&input.File_WOP)

	_, err = db.Create(&fileDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFile.CheckoutPhaseOneInstance(&fileDB)
	file := backRepo.BackRepoFile.Map_FileDBID_FilePtr[fileDB.ID]

	if file != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), file)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, fileDB)
}

// GetFile
//
// swagger:route GET /files/{ID} files getFile
//
// Gets the details for a file.
//
// Responses:
// default: genericError
//
//	200: fileDBResponse
func (controller *Controller) GetFile(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFile", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFile.GetDB()

	// Get fileDB in DB
	var fileDB orm.FileDB
	if _, err := db.First(&fileDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var fileAPI orm.FileAPI
	fileAPI.ID = fileDB.ID
	fileAPI.FilePointersEncoding = fileDB.FilePointersEncoding
	fileDB.CopyBasicFieldsToFile_WOP(&fileAPI.File_WOP)

	c.JSON(http.StatusOK, fileAPI)
}

// UpdateFile
//
// swagger:route PATCH /files/{ID} files updateFile
//
// # Update a file
//
// Responses:
// default: genericError
//
//	200: fileDBResponse
func (controller *Controller) UpdateFile(c *gin.Context) {

	mutexFile.Lock()
	defer mutexFile.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFile", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFile.GetDB()

	// Validate input
	var input orm.FileAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var fileDB orm.FileDB

	// fetch the file
	_, err := db.First(&fileDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	fileDB.CopyBasicFieldsFromFile_WOP(&input.File_WOP)
	fileDB.FilePointersEncoding = input.FilePointersEncoding

	db, _ = db.Model(&fileDB)
	_, err = db.Updates(fileDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	fileNew := new(models.File)
	fileDB.CopyBasicFieldsToFile(fileNew)

	// redeem pointers
	fileDB.DecodePointers(backRepo, fileNew)

	// get stage instance from DB instance, and call callback function
	fileOld := backRepo.BackRepoFile.Map_FileDBID_FilePtr[fileDB.ID]
	if fileOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), fileOld, fileNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the fileDB
	c.JSON(http.StatusOK, fileDB)
}

// DeleteFile
//
// swagger:route DELETE /files/{ID} files deleteFile
//
// # Delete a file
//
// default: genericError
//
//	200: fileDBResponse
func (controller *Controller) DeleteFile(c *gin.Context) {

	mutexFile.Lock()
	defer mutexFile.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFile", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFile.GetDB()

	// Get model if exist
	var fileDB orm.FileDB
	if _, err := db.First(&fileDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&fileDB)

	// get an instance (not staged) from DB instance, and call callback function
	fileDeleted := new(models.File)
	fileDB.CopyBasicFieldsToFile(fileDeleted)

	// get stage instance from DB instance, and call callback function
	fileStaged := backRepo.BackRepoFile.Map_FileDBID_FilePtr[fileDB.ID]
	if fileStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), fileStaged, fileDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
