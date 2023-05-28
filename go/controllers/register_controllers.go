package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongdocx/go")
	{ // insertion point for registrations
		v1.GET("/v1/documents", GetController().GetDocuments)
		v1.GET("/v1/documents/:id", GetController().GetDocument)
		v1.POST("/v1/documents", GetController().PostDocument)
		v1.PATCH("/v1/documents/:id", GetController().UpdateDocument)
		v1.PUT("/v1/documents/:id", GetController().UpdateDocument)
		v1.DELETE("/v1/documents/:id", GetController().DeleteDocument)

		v1.GET("/v1/docxs", GetController().GetDocxs)
		v1.GET("/v1/docxs/:id", GetController().GetDocx)
		v1.POST("/v1/docxs", GetController().PostDocx)
		v1.PATCH("/v1/docxs/:id", GetController().UpdateDocx)
		v1.PUT("/v1/docxs/:id", GetController().UpdateDocx)
		v1.DELETE("/v1/docxs/:id", GetController().DeleteDocx)

		v1.GET("/v1/files", GetController().GetFiles)
		v1.GET("/v1/files/:id", GetController().GetFile)
		v1.POST("/v1/files", GetController().PostFile)
		v1.PATCH("/v1/files/:id", GetController().UpdateFile)
		v1.PUT("/v1/files/:id", GetController().UpdateFile)
		v1.DELETE("/v1/files/:id", GetController().DeleteFile)

		v1.GET("/v1/nodes", GetController().GetNodes)
		v1.GET("/v1/nodes/:id", GetController().GetNode)
		v1.POST("/v1/nodes", GetController().PostNode)
		v1.PATCH("/v1/nodes/:id", GetController().UpdateNode)
		v1.PUT("/v1/nodes/:id", GetController().UpdateNode)
		v1.DELETE("/v1/nodes/:id", GetController().DeleteNode)

		v1.GET("/v1/paragraphs", GetController().GetParagraphs)
		v1.GET("/v1/paragraphs/:id", GetController().GetParagraph)
		v1.POST("/v1/paragraphs", GetController().PostParagraph)
		v1.PATCH("/v1/paragraphs/:id", GetController().UpdateParagraph)
		v1.PUT("/v1/paragraphs/:id", GetController().UpdateParagraph)
		v1.DELETE("/v1/paragraphs/:id", GetController().DeleteParagraph)

		v1.GET("/v1/paragraphpropertiess", GetController().GetParagraphPropertiess)
		v1.GET("/v1/paragraphpropertiess/:id", GetController().GetParagraphProperties)
		v1.POST("/v1/paragraphpropertiess", GetController().PostParagraphProperties)
		v1.PATCH("/v1/paragraphpropertiess/:id", GetController().UpdateParagraphProperties)
		v1.PUT("/v1/paragraphpropertiess/:id", GetController().UpdateParagraphProperties)
		v1.DELETE("/v1/paragraphpropertiess/:id", GetController().DeleteParagraphProperties)

		v1.GET("/v1/runes", GetController().GetRunes)
		v1.GET("/v1/runes/:id", GetController().GetRune)
		v1.POST("/v1/runes", GetController().PostRune)
		v1.PATCH("/v1/runes/:id", GetController().UpdateRune)
		v1.PUT("/v1/runes/:id", GetController().UpdateRune)
		v1.DELETE("/v1/runes/:id", GetController().DeleteRune)

		v1.GET("/v1/texts", GetController().GetTexts)
		v1.GET("/v1/texts/:id", GetController().GetText)
		v1.POST("/v1/texts", GetController().PostText)
		v1.PATCH("/v1/texts/:id", GetController().UpdateText)
		v1.PUT("/v1/texts/:id", GetController().UpdateText)
		v1.DELETE("/v1/texts/:id", GetController().DeleteText)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
