package load

import (
	"log"

	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type DiagramPackageCallbacksSingloton struct {
	DiagramPackageCallback DiagramPackageCallback
}

func (diagramPackageCallbacksSingloton *DiagramPackageCallbacksSingloton) OnAfterUpdate(
	gongdocStage *gongdoc_models.StageStruct,
	stagedDiagramPackage, frontDiagramPackage *gongdoc_models.DiagramPackage) {

	if stagedDiagramPackage.IsReloaded != frontDiagramPackage.IsReloaded {

		// reset the IsReloaded to false
		stagedDiagramPackage.Checkout(gongdocStage)
		stagedDiagramPackage.IsReloaded = false
		stagedDiagramPackage.Commit(gongdocStage)

		log.Println("Reload requested")
		if stagedDiagramPackage.IsEditable {
			Reload(gongdocStage, stagedDiagramPackage)
		}
	}
}

type DiagramPackageCallback interface {
	HasSelected(gongstructName string)
}
