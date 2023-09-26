// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdocx/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.playground = playground
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.Body:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Document:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Docx:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.File:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Node:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Paragraph:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.ParagraphProperties:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.ParagraphStyle:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Rune:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.RuneProperties:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Table:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.TableColumn:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.TableProperties:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.TableRow:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.TableStyle:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Text:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.playground.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.playground)
	fillUpTree(cellDeleteIconImpl.playground)
	cellDeleteIconImpl.playground.tableStage.Commit()
}

