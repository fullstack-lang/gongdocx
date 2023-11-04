// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdocx/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.probe = probe
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	probe *Probe
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.Body:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Document:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Docx:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.File:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Node:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Paragraph:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ParagraphProperties:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ParagraphStyle:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Rune:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.RuneProperties:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Table:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TableColumn:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TableProperties:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TableRow:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.TableStyle:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Text:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.probe)
	fillUpTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}

