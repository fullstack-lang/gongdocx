// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
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
		v1.GET("/v1/bodys", GetController().GetBodys)
		v1.GET("/v1/bodys/:id", GetController().GetBody)
		v1.POST("/v1/bodys", GetController().PostBody)
		v1.PATCH("/v1/bodys/:id", GetController().UpdateBody)
		v1.PUT("/v1/bodys/:id", GetController().UpdateBody)
		v1.DELETE("/v1/bodys/:id", GetController().DeleteBody)

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

		v1.GET("/v1/paragraphstyles", GetController().GetParagraphStyles)
		v1.GET("/v1/paragraphstyles/:id", GetController().GetParagraphStyle)
		v1.POST("/v1/paragraphstyles", GetController().PostParagraphStyle)
		v1.PATCH("/v1/paragraphstyles/:id", GetController().UpdateParagraphStyle)
		v1.PUT("/v1/paragraphstyles/:id", GetController().UpdateParagraphStyle)
		v1.DELETE("/v1/paragraphstyles/:id", GetController().DeleteParagraphStyle)

		v1.GET("/v1/runes", GetController().GetRunes)
		v1.GET("/v1/runes/:id", GetController().GetRune)
		v1.POST("/v1/runes", GetController().PostRune)
		v1.PATCH("/v1/runes/:id", GetController().UpdateRune)
		v1.PUT("/v1/runes/:id", GetController().UpdateRune)
		v1.DELETE("/v1/runes/:id", GetController().DeleteRune)

		v1.GET("/v1/runepropertiess", GetController().GetRunePropertiess)
		v1.GET("/v1/runepropertiess/:id", GetController().GetRuneProperties)
		v1.POST("/v1/runepropertiess", GetController().PostRuneProperties)
		v1.PATCH("/v1/runepropertiess/:id", GetController().UpdateRuneProperties)
		v1.PUT("/v1/runepropertiess/:id", GetController().UpdateRuneProperties)
		v1.DELETE("/v1/runepropertiess/:id", GetController().DeleteRuneProperties)

		v1.GET("/v1/tables", GetController().GetTables)
		v1.GET("/v1/tables/:id", GetController().GetTable)
		v1.POST("/v1/tables", GetController().PostTable)
		v1.PATCH("/v1/tables/:id", GetController().UpdateTable)
		v1.PUT("/v1/tables/:id", GetController().UpdateTable)
		v1.DELETE("/v1/tables/:id", GetController().DeleteTable)

		v1.GET("/v1/tablecolumns", GetController().GetTableColumns)
		v1.GET("/v1/tablecolumns/:id", GetController().GetTableColumn)
		v1.POST("/v1/tablecolumns", GetController().PostTableColumn)
		v1.PATCH("/v1/tablecolumns/:id", GetController().UpdateTableColumn)
		v1.PUT("/v1/tablecolumns/:id", GetController().UpdateTableColumn)
		v1.DELETE("/v1/tablecolumns/:id", GetController().DeleteTableColumn)

		v1.GET("/v1/tablepropertiess", GetController().GetTablePropertiess)
		v1.GET("/v1/tablepropertiess/:id", GetController().GetTableProperties)
		v1.POST("/v1/tablepropertiess", GetController().PostTableProperties)
		v1.PATCH("/v1/tablepropertiess/:id", GetController().UpdateTableProperties)
		v1.PUT("/v1/tablepropertiess/:id", GetController().UpdateTableProperties)
		v1.DELETE("/v1/tablepropertiess/:id", GetController().DeleteTableProperties)

		v1.GET("/v1/tablerows", GetController().GetTableRows)
		v1.GET("/v1/tablerows/:id", GetController().GetTableRow)
		v1.POST("/v1/tablerows", GetController().PostTableRow)
		v1.PATCH("/v1/tablerows/:id", GetController().UpdateTableRow)
		v1.PUT("/v1/tablerows/:id", GetController().UpdateTableRow)
		v1.DELETE("/v1/tablerows/:id", GetController().DeleteTableRow)

		v1.GET("/v1/tablestyles", GetController().GetTableStyles)
		v1.GET("/v1/tablestyles/:id", GetController().GetTableStyle)
		v1.POST("/v1/tablestyles", GetController().PostTableStyle)
		v1.PATCH("/v1/tablestyles/:id", GetController().UpdateTableStyle)
		v1.PUT("/v1/tablestyles/:id", GetController().UpdateTableStyle)
		v1.DELETE("/v1/tablestyles/:id", GetController().DeleteTableStyle)

		v1.GET("/v1/texts", GetController().GetTexts)
		v1.GET("/v1/texts/:id", GetController().GetText)
		v1.POST("/v1/texts", GetController().PostText)
		v1.PATCH("/v1/texts/:id", GetController().UpdateText)
		v1.PUT("/v1/texts/:id", GetController().UpdateText)
		v1.DELETE("/v1/texts/:id", GetController().DeleteText)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws", GetController().onWebSocketRequest)
	}
}

// onWebSocketRequest is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequest(c *gin.Context) {

	log.Println("onWebSocketRequest")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

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
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.Subscribe()

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {

		// Send elapsed time as a string over the WebSocket connection
		err = wsConnection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", nbCommitBackRepo)))
		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
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
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
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
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdocx/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
