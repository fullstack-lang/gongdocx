// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongdocx/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	playground *Playground
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	playground *Playground,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.playground = playground
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Body" {
		fillUpTable[models.Body](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Document" {
		fillUpTable[models.Document](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Docx" {
		fillUpTable[models.Docx](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "File" {
		fillUpTable[models.File](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Node" {
		fillUpTable[models.Node](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Paragraph" {
		fillUpTable[models.Paragraph](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ParagraphProperties" {
		fillUpTable[models.ParagraphProperties](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ParagraphStyle" {
		fillUpTable[models.ParagraphStyle](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rune" {
		fillUpTable[models.Rune](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RuneProperties" {
		fillUpTable[models.RuneProperties](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Table" {
		fillUpTable[models.Table](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TableColumn" {
		fillUpTable[models.TableColumn](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TableProperties" {
		fillUpTable[models.TableProperties](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TableRow" {
		fillUpTable[models.TableRow](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TableStyle" {
		fillUpTable[models.TableStyle](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Text" {
		fillUpTable[models.Text](nodeImplGongstruct.playground)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.playground.tableStage.Commit()
}
