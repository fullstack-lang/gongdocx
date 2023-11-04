// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdocx/go/models"
	"github.com/fullstack-lang/gongdocx/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__BodyFormCallback(
	body *models.Body,
	probe *Probe,
) (bodyFormCallback *BodyFormCallback) {
	bodyFormCallback = new(BodyFormCallback)
	bodyFormCallback.probe = probe
	bodyFormCallback.body = body

	bodyFormCallback.CreationMode = (body == nil)

	return
}

type BodyFormCallback struct {
	body *models.Body

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (bodyFormCallback *BodyFormCallback) OnSave() {

	log.Println("BodyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bodyFormCallback.probe.formStage.Checkout()

	if bodyFormCallback.body == nil {
		bodyFormCallback.body = new(models.Body).Stage(bodyFormCallback.probe.stageOfInterest)
	}
	body_ := bodyFormCallback.body
	_ = body_

	// get the formGroup
	formGroup := bodyFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(body_.Name), formDiv)
		case "LastParagraph":
			FormDivSelectFieldToField(&(body_.LastParagraph), bodyFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	bodyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Body](
		bodyFormCallback.probe,
	)
	bodyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if bodyFormCallback.CreationMode {
		bodyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__BodyFormCallback(
				nil,
				bodyFormCallback.probe,
			),
		}).Stage(bodyFormCallback.probe.formStage)
		body := new(models.Body)
		FillUpForm(body, newFormGroup, bodyFormCallback.probe)
		bodyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(bodyFormCallback.probe)
}
func __gong__New__DocumentFormCallback(
	document *models.Document,
	probe *Probe,
) (documentFormCallback *DocumentFormCallback) {
	documentFormCallback = new(DocumentFormCallback)
	documentFormCallback.probe = probe
	documentFormCallback.document = document

	documentFormCallback.CreationMode = (document == nil)

	return
}

type DocumentFormCallback struct {
	document *models.Document

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (documentFormCallback *DocumentFormCallback) OnSave() {

	log.Println("DocumentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentFormCallback.probe.formStage.Checkout()

	if documentFormCallback.document == nil {
		documentFormCallback.document = new(models.Document).Stage(documentFormCallback.probe.stageOfInterest)
	}
	document_ := documentFormCallback.document
	_ = document_

	// get the formGroup
	formGroup := documentFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(document_.Name), formDiv)
		case "File":
			FormDivSelectFieldToField(&(document_.File), documentFormCallback.probe.stageOfInterest, formDiv)
		case "Root":
			FormDivSelectFieldToField(&(document_.Root), documentFormCallback.probe.stageOfInterest, formDiv)
		case "Body":
			FormDivSelectFieldToField(&(document_.Body), documentFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	documentFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Document](
		documentFormCallback.probe,
	)
	documentFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if documentFormCallback.CreationMode {
		documentFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DocumentFormCallback(
				nil,
				documentFormCallback.probe,
			),
		}).Stage(documentFormCallback.probe.formStage)
		document := new(models.Document)
		FillUpForm(document, newFormGroup, documentFormCallback.probe)
		documentFormCallback.probe.formStage.Commit()
	}

	fillUpTree(documentFormCallback.probe)
}
func __gong__New__DocxFormCallback(
	docx *models.Docx,
	probe *Probe,
) (docxFormCallback *DocxFormCallback) {
	docxFormCallback = new(DocxFormCallback)
	docxFormCallback.probe = probe
	docxFormCallback.docx = docx

	docxFormCallback.CreationMode = (docx == nil)

	return
}

type DocxFormCallback struct {
	docx *models.Docx

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (docxFormCallback *DocxFormCallback) OnSave() {

	log.Println("DocxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	docxFormCallback.probe.formStage.Checkout()

	if docxFormCallback.docx == nil {
		docxFormCallback.docx = new(models.Docx).Stage(docxFormCallback.probe.stageOfInterest)
	}
	docx_ := docxFormCallback.docx
	_ = docx_

	// get the formGroup
	formGroup := docxFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(docx_.Name), formDiv)
		case "Document":
			FormDivSelectFieldToField(&(docx_.Document), docxFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	docxFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Docx](
		docxFormCallback.probe,
	)
	docxFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if docxFormCallback.CreationMode {
		docxFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DocxFormCallback(
				nil,
				docxFormCallback.probe,
			),
		}).Stage(docxFormCallback.probe.formStage)
		docx := new(models.Docx)
		FillUpForm(docx, newFormGroup, docxFormCallback.probe)
		docxFormCallback.probe.formStage.Commit()
	}

	fillUpTree(docxFormCallback.probe)
}
func __gong__New__FileFormCallback(
	file *models.File,
	probe *Probe,
) (fileFormCallback *FileFormCallback) {
	fileFormCallback = new(FileFormCallback)
	fileFormCallback.probe = probe
	fileFormCallback.file = file

	fileFormCallback.CreationMode = (file == nil)

	return
}

type FileFormCallback struct {
	file *models.File

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (fileFormCallback *FileFormCallback) OnSave() {

	log.Println("FileFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fileFormCallback.probe.formStage.Checkout()

	if fileFormCallback.file == nil {
		fileFormCallback.file = new(models.File).Stage(fileFormCallback.probe.stageOfInterest)
	}
	file_ := fileFormCallback.file
	_ = file_

	// get the formGroup
	formGroup := fileFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(file_.Name), formDiv)
		case "Docx:Files":
			// we need to retrieve the field owner before the change
			var pastDocxOwner *models.Docx
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Docx"
			rf.Fieldname = "Files"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				fileFormCallback.probe.stageOfInterest,
				fileFormCallback.probe.backRepoOfInterest,
				file_,
				&rf)

			if reverseFieldOwner != nil {
				pastDocxOwner = reverseFieldOwner.(*models.Docx)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDocxOwner != nil {
					idx := slices.Index(pastDocxOwner.Files, file_)
					pastDocxOwner.Files = slices.Delete(pastDocxOwner.Files, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _docx := range *models.GetGongstructInstancesSet[models.Docx](fileFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _docx.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDocxOwner := _docx // we have a match
						if pastDocxOwner != nil {
							if newDocxOwner != pastDocxOwner {
								idx := slices.Index(pastDocxOwner.Files, file_)
								pastDocxOwner.Files = slices.Delete(pastDocxOwner.Files, idx, idx+1)
								newDocxOwner.Files = append(newDocxOwner.Files, file_)
							}
						} else {
							newDocxOwner.Files = append(newDocxOwner.Files, file_)
						}
					}
				}
			}
		}
	}

	fileFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.File](
		fileFormCallback.probe,
	)
	fileFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fileFormCallback.CreationMode {
		fileFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FileFormCallback(
				nil,
				fileFormCallback.probe,
			),
		}).Stage(fileFormCallback.probe.formStage)
		file := new(models.File)
		FillUpForm(file, newFormGroup, fileFormCallback.probe)
		fileFormCallback.probe.formStage.Commit()
	}

	fillUpTree(fileFormCallback.probe)
}
func __gong__New__NodeFormCallback(
	node *models.Node,
	probe *Probe,
) (nodeFormCallback *NodeFormCallback) {
	nodeFormCallback = new(NodeFormCallback)
	nodeFormCallback.probe = probe
	nodeFormCallback.node = node

	nodeFormCallback.CreationMode = (node == nil)

	return
}

type NodeFormCallback struct {
	node *models.Node

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (nodeFormCallback *NodeFormCallback) OnSave() {

	log.Println("NodeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	nodeFormCallback.probe.formStage.Checkout()

	if nodeFormCallback.node == nil {
		nodeFormCallback.node = new(models.Node).Stage(nodeFormCallback.probe.stageOfInterest)
	}
	node_ := nodeFormCallback.node
	_ = node_

	// get the formGroup
	formGroup := nodeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(node_.Name), formDiv)
		case "Node:Nodes":
			// we need to retrieve the field owner before the change
			var pastNodeOwner *models.Node
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Nodes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				nodeFormCallback.probe.stageOfInterest,
				nodeFormCallback.probe.backRepoOfInterest,
				node_,
				&rf)

			if reverseFieldOwner != nil {
				pastNodeOwner = reverseFieldOwner.(*models.Node)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNodeOwner != nil {
					idx := slices.Index(pastNodeOwner.Nodes, node_)
					pastNodeOwner.Nodes = slices.Delete(pastNodeOwner.Nodes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _node := range *models.GetGongstructInstancesSet[models.Node](nodeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _node.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNodeOwner := _node // we have a match
						if pastNodeOwner != nil {
							if newNodeOwner != pastNodeOwner {
								idx := slices.Index(pastNodeOwner.Nodes, node_)
								pastNodeOwner.Nodes = slices.Delete(pastNodeOwner.Nodes, idx, idx+1)
								newNodeOwner.Nodes = append(newNodeOwner.Nodes, node_)
							}
						} else {
							newNodeOwner.Nodes = append(newNodeOwner.Nodes, node_)
						}
					}
				}
			}
		}
	}

	nodeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Node](
		nodeFormCallback.probe,
	)
	nodeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if nodeFormCallback.CreationMode {
		nodeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__NodeFormCallback(
				nil,
				nodeFormCallback.probe,
			),
		}).Stage(nodeFormCallback.probe.formStage)
		node := new(models.Node)
		FillUpForm(node, newFormGroup, nodeFormCallback.probe)
		nodeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(nodeFormCallback.probe)
}
func __gong__New__ParagraphFormCallback(
	paragraph *models.Paragraph,
	probe *Probe,
) (paragraphFormCallback *ParagraphFormCallback) {
	paragraphFormCallback = new(ParagraphFormCallback)
	paragraphFormCallback.probe = probe
	paragraphFormCallback.paragraph = paragraph

	paragraphFormCallback.CreationMode = (paragraph == nil)

	return
}

type ParagraphFormCallback struct {
	paragraph *models.Paragraph

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (paragraphFormCallback *ParagraphFormCallback) OnSave() {

	log.Println("ParagraphFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	paragraphFormCallback.probe.formStage.Checkout()

	if paragraphFormCallback.paragraph == nil {
		paragraphFormCallback.paragraph = new(models.Paragraph).Stage(paragraphFormCallback.probe.stageOfInterest)
	}
	paragraph_ := paragraphFormCallback.paragraph
	_ = paragraph_

	// get the formGroup
	formGroup := paragraphFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(paragraph_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(paragraph_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(paragraph_.Node), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "ParagraphProperties":
			FormDivSelectFieldToField(&(paragraph_.ParagraphProperties), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "Text":
			FormDivBasicFieldToField(&(paragraph_.Text), formDiv)
		case "Next":
			FormDivSelectFieldToField(&(paragraph_.Next), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "Previous":
			FormDivSelectFieldToField(&(paragraph_.Previous), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "EnclosingBody":
			FormDivSelectFieldToField(&(paragraph_.EnclosingBody), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "EnclosingTableColumn":
			FormDivSelectFieldToField(&(paragraph_.EnclosingTableColumn), paragraphFormCallback.probe.stageOfInterest, formDiv)
		case "Body:Paragraphs":
			// we need to retrieve the field owner before the change
			var pastBodyOwner *models.Body
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Body"
			rf.Fieldname = "Paragraphs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				paragraphFormCallback.probe.stageOfInterest,
				paragraphFormCallback.probe.backRepoOfInterest,
				paragraph_,
				&rf)

			if reverseFieldOwner != nil {
				pastBodyOwner = reverseFieldOwner.(*models.Body)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastBodyOwner != nil {
					idx := slices.Index(pastBodyOwner.Paragraphs, paragraph_)
					pastBodyOwner.Paragraphs = slices.Delete(pastBodyOwner.Paragraphs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _body := range *models.GetGongstructInstancesSet[models.Body](paragraphFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _body.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newBodyOwner := _body // we have a match
						if pastBodyOwner != nil {
							if newBodyOwner != pastBodyOwner {
								idx := slices.Index(pastBodyOwner.Paragraphs, paragraph_)
								pastBodyOwner.Paragraphs = slices.Delete(pastBodyOwner.Paragraphs, idx, idx+1)
								newBodyOwner.Paragraphs = append(newBodyOwner.Paragraphs, paragraph_)
							}
						} else {
							newBodyOwner.Paragraphs = append(newBodyOwner.Paragraphs, paragraph_)
						}
					}
				}
			}
		case "TableColumn:Paragraphs":
			// we need to retrieve the field owner before the change
			var pastTableColumnOwner *models.TableColumn
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TableColumn"
			rf.Fieldname = "Paragraphs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				paragraphFormCallback.probe.stageOfInterest,
				paragraphFormCallback.probe.backRepoOfInterest,
				paragraph_,
				&rf)

			if reverseFieldOwner != nil {
				pastTableColumnOwner = reverseFieldOwner.(*models.TableColumn)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTableColumnOwner != nil {
					idx := slices.Index(pastTableColumnOwner.Paragraphs, paragraph_)
					pastTableColumnOwner.Paragraphs = slices.Delete(pastTableColumnOwner.Paragraphs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _tablecolumn := range *models.GetGongstructInstancesSet[models.TableColumn](paragraphFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _tablecolumn.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTableColumnOwner := _tablecolumn // we have a match
						if pastTableColumnOwner != nil {
							if newTableColumnOwner != pastTableColumnOwner {
								idx := slices.Index(pastTableColumnOwner.Paragraphs, paragraph_)
								pastTableColumnOwner.Paragraphs = slices.Delete(pastTableColumnOwner.Paragraphs, idx, idx+1)
								newTableColumnOwner.Paragraphs = append(newTableColumnOwner.Paragraphs, paragraph_)
							}
						} else {
							newTableColumnOwner.Paragraphs = append(newTableColumnOwner.Paragraphs, paragraph_)
						}
					}
				}
			}
		}
	}

	paragraphFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Paragraph](
		paragraphFormCallback.probe,
	)
	paragraphFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if paragraphFormCallback.CreationMode {
		paragraphFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ParagraphFormCallback(
				nil,
				paragraphFormCallback.probe,
			),
		}).Stage(paragraphFormCallback.probe.formStage)
		paragraph := new(models.Paragraph)
		FillUpForm(paragraph, newFormGroup, paragraphFormCallback.probe)
		paragraphFormCallback.probe.formStage.Commit()
	}

	fillUpTree(paragraphFormCallback.probe)
}
func __gong__New__ParagraphPropertiesFormCallback(
	paragraphproperties *models.ParagraphProperties,
	probe *Probe,
) (paragraphpropertiesFormCallback *ParagraphPropertiesFormCallback) {
	paragraphpropertiesFormCallback = new(ParagraphPropertiesFormCallback)
	paragraphpropertiesFormCallback.probe = probe
	paragraphpropertiesFormCallback.paragraphproperties = paragraphproperties

	paragraphpropertiesFormCallback.CreationMode = (paragraphproperties == nil)

	return
}

type ParagraphPropertiesFormCallback struct {
	paragraphproperties *models.ParagraphProperties

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (paragraphpropertiesFormCallback *ParagraphPropertiesFormCallback) OnSave() {

	log.Println("ParagraphPropertiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	paragraphpropertiesFormCallback.probe.formStage.Checkout()

	if paragraphpropertiesFormCallback.paragraphproperties == nil {
		paragraphpropertiesFormCallback.paragraphproperties = new(models.ParagraphProperties).Stage(paragraphpropertiesFormCallback.probe.stageOfInterest)
	}
	paragraphproperties_ := paragraphpropertiesFormCallback.paragraphproperties
	_ = paragraphproperties_

	// get the formGroup
	formGroup := paragraphpropertiesFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(paragraphproperties_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(paragraphproperties_.Content), formDiv)
		case "ParagraphStyle":
			FormDivSelectFieldToField(&(paragraphproperties_.ParagraphStyle), paragraphpropertiesFormCallback.probe.stageOfInterest, formDiv)
		case "Node":
			FormDivSelectFieldToField(&(paragraphproperties_.Node), paragraphpropertiesFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	paragraphpropertiesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ParagraphProperties](
		paragraphpropertiesFormCallback.probe,
	)
	paragraphpropertiesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if paragraphpropertiesFormCallback.CreationMode {
		paragraphpropertiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ParagraphPropertiesFormCallback(
				nil,
				paragraphpropertiesFormCallback.probe,
			),
		}).Stage(paragraphpropertiesFormCallback.probe.formStage)
		paragraphproperties := new(models.ParagraphProperties)
		FillUpForm(paragraphproperties, newFormGroup, paragraphpropertiesFormCallback.probe)
		paragraphpropertiesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(paragraphpropertiesFormCallback.probe)
}
func __gong__New__ParagraphStyleFormCallback(
	paragraphstyle *models.ParagraphStyle,
	probe *Probe,
) (paragraphstyleFormCallback *ParagraphStyleFormCallback) {
	paragraphstyleFormCallback = new(ParagraphStyleFormCallback)
	paragraphstyleFormCallback.probe = probe
	paragraphstyleFormCallback.paragraphstyle = paragraphstyle

	paragraphstyleFormCallback.CreationMode = (paragraphstyle == nil)

	return
}

type ParagraphStyleFormCallback struct {
	paragraphstyle *models.ParagraphStyle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (paragraphstyleFormCallback *ParagraphStyleFormCallback) OnSave() {

	log.Println("ParagraphStyleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	paragraphstyleFormCallback.probe.formStage.Checkout()

	if paragraphstyleFormCallback.paragraphstyle == nil {
		paragraphstyleFormCallback.paragraphstyle = new(models.ParagraphStyle).Stage(paragraphstyleFormCallback.probe.stageOfInterest)
	}
	paragraphstyle_ := paragraphstyleFormCallback.paragraphstyle
	_ = paragraphstyle_

	// get the formGroup
	formGroup := paragraphstyleFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(paragraphstyle_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(paragraphstyle_.Node), paragraphstyleFormCallback.probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(paragraphstyle_.Content), formDiv)
		case "ValAttr":
			FormDivBasicFieldToField(&(paragraphstyle_.ValAttr), formDiv)
		}
	}

	paragraphstyleFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.ParagraphStyle](
		paragraphstyleFormCallback.probe,
	)
	paragraphstyleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if paragraphstyleFormCallback.CreationMode {
		paragraphstyleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ParagraphStyleFormCallback(
				nil,
				paragraphstyleFormCallback.probe,
			),
		}).Stage(paragraphstyleFormCallback.probe.formStage)
		paragraphstyle := new(models.ParagraphStyle)
		FillUpForm(paragraphstyle, newFormGroup, paragraphstyleFormCallback.probe)
		paragraphstyleFormCallback.probe.formStage.Commit()
	}

	fillUpTree(paragraphstyleFormCallback.probe)
}
func __gong__New__RuneFormCallback(
	rune *models.Rune,
	probe *Probe,
) (runeFormCallback *RuneFormCallback) {
	runeFormCallback = new(RuneFormCallback)
	runeFormCallback.probe = probe
	runeFormCallback.rune = rune

	runeFormCallback.CreationMode = (rune == nil)

	return
}

type RuneFormCallback struct {
	rune *models.Rune

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (runeFormCallback *RuneFormCallback) OnSave() {

	log.Println("RuneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	runeFormCallback.probe.formStage.Checkout()

	if runeFormCallback.rune == nil {
		runeFormCallback.rune = new(models.Rune).Stage(runeFormCallback.probe.stageOfInterest)
	}
	rune_ := runeFormCallback.rune
	_ = rune_

	// get the formGroup
	formGroup := runeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rune_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(rune_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(rune_.Node), runeFormCallback.probe.stageOfInterest, formDiv)
		case "Text":
			FormDivSelectFieldToField(&(rune_.Text), runeFormCallback.probe.stageOfInterest, formDiv)
		case "RuneProperties":
			FormDivSelectFieldToField(&(rune_.RuneProperties), runeFormCallback.probe.stageOfInterest, formDiv)
		case "EnclosingParagraph":
			FormDivSelectFieldToField(&(rune_.EnclosingParagraph), runeFormCallback.probe.stageOfInterest, formDiv)
		case "Paragraph:Runes":
			// we need to retrieve the field owner before the change
			var pastParagraphOwner *models.Paragraph
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Paragraph"
			rf.Fieldname = "Runes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				runeFormCallback.probe.stageOfInterest,
				runeFormCallback.probe.backRepoOfInterest,
				rune_,
				&rf)

			if reverseFieldOwner != nil {
				pastParagraphOwner = reverseFieldOwner.(*models.Paragraph)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastParagraphOwner != nil {
					idx := slices.Index(pastParagraphOwner.Runes, rune_)
					pastParagraphOwner.Runes = slices.Delete(pastParagraphOwner.Runes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _paragraph := range *models.GetGongstructInstancesSet[models.Paragraph](runeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _paragraph.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newParagraphOwner := _paragraph // we have a match
						if pastParagraphOwner != nil {
							if newParagraphOwner != pastParagraphOwner {
								idx := slices.Index(pastParagraphOwner.Runes, rune_)
								pastParagraphOwner.Runes = slices.Delete(pastParagraphOwner.Runes, idx, idx+1)
								newParagraphOwner.Runes = append(newParagraphOwner.Runes, rune_)
							}
						} else {
							newParagraphOwner.Runes = append(newParagraphOwner.Runes, rune_)
						}
					}
				}
			}
		}
	}

	runeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Rune](
		runeFormCallback.probe,
	)
	runeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if runeFormCallback.CreationMode {
		runeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RuneFormCallback(
				nil,
				runeFormCallback.probe,
			),
		}).Stage(runeFormCallback.probe.formStage)
		rune := new(models.Rune)
		FillUpForm(rune, newFormGroup, runeFormCallback.probe)
		runeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(runeFormCallback.probe)
}
func __gong__New__RunePropertiesFormCallback(
	runeproperties *models.RuneProperties,
	probe *Probe,
) (runepropertiesFormCallback *RunePropertiesFormCallback) {
	runepropertiesFormCallback = new(RunePropertiesFormCallback)
	runepropertiesFormCallback.probe = probe
	runepropertiesFormCallback.runeproperties = runeproperties

	runepropertiesFormCallback.CreationMode = (runeproperties == nil)

	return
}

type RunePropertiesFormCallback struct {
	runeproperties *models.RuneProperties

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (runepropertiesFormCallback *RunePropertiesFormCallback) OnSave() {

	log.Println("RunePropertiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	runepropertiesFormCallback.probe.formStage.Checkout()

	if runepropertiesFormCallback.runeproperties == nil {
		runepropertiesFormCallback.runeproperties = new(models.RuneProperties).Stage(runepropertiesFormCallback.probe.stageOfInterest)
	}
	runeproperties_ := runepropertiesFormCallback.runeproperties
	_ = runeproperties_

	// get the formGroup
	formGroup := runepropertiesFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(runeproperties_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(runeproperties_.Node), runepropertiesFormCallback.probe.stageOfInterest, formDiv)
		case "IsBold":
			FormDivBasicFieldToField(&(runeproperties_.IsBold), formDiv)
		case "IsStrike":
			FormDivBasicFieldToField(&(runeproperties_.IsStrike), formDiv)
		case "IsItalic":
			FormDivBasicFieldToField(&(runeproperties_.IsItalic), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(runeproperties_.Content), formDiv)
		}
	}

	runepropertiesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.RuneProperties](
		runepropertiesFormCallback.probe,
	)
	runepropertiesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if runepropertiesFormCallback.CreationMode {
		runepropertiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RunePropertiesFormCallback(
				nil,
				runepropertiesFormCallback.probe,
			),
		}).Stage(runepropertiesFormCallback.probe.formStage)
		runeproperties := new(models.RuneProperties)
		FillUpForm(runeproperties, newFormGroup, runepropertiesFormCallback.probe)
		runepropertiesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(runepropertiesFormCallback.probe)
}
func __gong__New__TableFormCallback(
	table *models.Table,
	probe *Probe,
) (tableFormCallback *TableFormCallback) {
	tableFormCallback = new(TableFormCallback)
	tableFormCallback.probe = probe
	tableFormCallback.table = table

	tableFormCallback.CreationMode = (table == nil)

	return
}

type TableFormCallback struct {
	table *models.Table

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (tableFormCallback *TableFormCallback) OnSave() {

	log.Println("TableFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tableFormCallback.probe.formStage.Checkout()

	if tableFormCallback.table == nil {
		tableFormCallback.table = new(models.Table).Stage(tableFormCallback.probe.stageOfInterest)
	}
	table_ := tableFormCallback.table
	_ = table_

	// get the formGroup
	formGroup := tableFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(table_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(table_.Node), tableFormCallback.probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(table_.Content), formDiv)
		case "TableProperties":
			FormDivSelectFieldToField(&(table_.TableProperties), tableFormCallback.probe.stageOfInterest, formDiv)
		case "Body:Tables":
			// we need to retrieve the field owner before the change
			var pastBodyOwner *models.Body
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Body"
			rf.Fieldname = "Tables"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				tableFormCallback.probe.stageOfInterest,
				tableFormCallback.probe.backRepoOfInterest,
				table_,
				&rf)

			if reverseFieldOwner != nil {
				pastBodyOwner = reverseFieldOwner.(*models.Body)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastBodyOwner != nil {
					idx := slices.Index(pastBodyOwner.Tables, table_)
					pastBodyOwner.Tables = slices.Delete(pastBodyOwner.Tables, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _body := range *models.GetGongstructInstancesSet[models.Body](tableFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _body.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newBodyOwner := _body // we have a match
						if pastBodyOwner != nil {
							if newBodyOwner != pastBodyOwner {
								idx := slices.Index(pastBodyOwner.Tables, table_)
								pastBodyOwner.Tables = slices.Delete(pastBodyOwner.Tables, idx, idx+1)
								newBodyOwner.Tables = append(newBodyOwner.Tables, table_)
							}
						} else {
							newBodyOwner.Tables = append(newBodyOwner.Tables, table_)
						}
					}
				}
			}
		}
	}

	tableFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Table](
		tableFormCallback.probe,
	)
	tableFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tableFormCallback.CreationMode {
		tableFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TableFormCallback(
				nil,
				tableFormCallback.probe,
			),
		}).Stage(tableFormCallback.probe.formStage)
		table := new(models.Table)
		FillUpForm(table, newFormGroup, tableFormCallback.probe)
		tableFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tableFormCallback.probe)
}
func __gong__New__TableColumnFormCallback(
	tablecolumn *models.TableColumn,
	probe *Probe,
) (tablecolumnFormCallback *TableColumnFormCallback) {
	tablecolumnFormCallback = new(TableColumnFormCallback)
	tablecolumnFormCallback.probe = probe
	tablecolumnFormCallback.tablecolumn = tablecolumn

	tablecolumnFormCallback.CreationMode = (tablecolumn == nil)

	return
}

type TableColumnFormCallback struct {
	tablecolumn *models.TableColumn

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (tablecolumnFormCallback *TableColumnFormCallback) OnSave() {

	log.Println("TableColumnFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablecolumnFormCallback.probe.formStage.Checkout()

	if tablecolumnFormCallback.tablecolumn == nil {
		tablecolumnFormCallback.tablecolumn = new(models.TableColumn).Stage(tablecolumnFormCallback.probe.stageOfInterest)
	}
	tablecolumn_ := tablecolumnFormCallback.tablecolumn
	_ = tablecolumn_

	// get the formGroup
	formGroup := tablecolumnFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tablecolumn_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tablecolumn_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tablecolumn_.Node), tablecolumnFormCallback.probe.stageOfInterest, formDiv)
		case "TableRow:TableColumns":
			// we need to retrieve the field owner before the change
			var pastTableRowOwner *models.TableRow
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TableRow"
			rf.Fieldname = "TableColumns"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				tablecolumnFormCallback.probe.stageOfInterest,
				tablecolumnFormCallback.probe.backRepoOfInterest,
				tablecolumn_,
				&rf)

			if reverseFieldOwner != nil {
				pastTableRowOwner = reverseFieldOwner.(*models.TableRow)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTableRowOwner != nil {
					idx := slices.Index(pastTableRowOwner.TableColumns, tablecolumn_)
					pastTableRowOwner.TableColumns = slices.Delete(pastTableRowOwner.TableColumns, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _tablerow := range *models.GetGongstructInstancesSet[models.TableRow](tablecolumnFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _tablerow.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTableRowOwner := _tablerow // we have a match
						if pastTableRowOwner != nil {
							if newTableRowOwner != pastTableRowOwner {
								idx := slices.Index(pastTableRowOwner.TableColumns, tablecolumn_)
								pastTableRowOwner.TableColumns = slices.Delete(pastTableRowOwner.TableColumns, idx, idx+1)
								newTableRowOwner.TableColumns = append(newTableRowOwner.TableColumns, tablecolumn_)
							}
						} else {
							newTableRowOwner.TableColumns = append(newTableRowOwner.TableColumns, tablecolumn_)
						}
					}
				}
			}
		}
	}

	tablecolumnFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TableColumn](
		tablecolumnFormCallback.probe,
	)
	tablecolumnFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tablecolumnFormCallback.CreationMode {
		tablecolumnFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TableColumnFormCallback(
				nil,
				tablecolumnFormCallback.probe,
			),
		}).Stage(tablecolumnFormCallback.probe.formStage)
		tablecolumn := new(models.TableColumn)
		FillUpForm(tablecolumn, newFormGroup, tablecolumnFormCallback.probe)
		tablecolumnFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tablecolumnFormCallback.probe)
}
func __gong__New__TablePropertiesFormCallback(
	tableproperties *models.TableProperties,
	probe *Probe,
) (tablepropertiesFormCallback *TablePropertiesFormCallback) {
	tablepropertiesFormCallback = new(TablePropertiesFormCallback)
	tablepropertiesFormCallback.probe = probe
	tablepropertiesFormCallback.tableproperties = tableproperties

	tablepropertiesFormCallback.CreationMode = (tableproperties == nil)

	return
}

type TablePropertiesFormCallback struct {
	tableproperties *models.TableProperties

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (tablepropertiesFormCallback *TablePropertiesFormCallback) OnSave() {

	log.Println("TablePropertiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablepropertiesFormCallback.probe.formStage.Checkout()

	if tablepropertiesFormCallback.tableproperties == nil {
		tablepropertiesFormCallback.tableproperties = new(models.TableProperties).Stage(tablepropertiesFormCallback.probe.stageOfInterest)
	}
	tableproperties_ := tablepropertiesFormCallback.tableproperties
	_ = tableproperties_

	// get the formGroup
	formGroup := tablepropertiesFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tableproperties_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tableproperties_.Node), tablepropertiesFormCallback.probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tableproperties_.Content), formDiv)
		case "TableStyle":
			FormDivSelectFieldToField(&(tableproperties_.TableStyle), tablepropertiesFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	tablepropertiesFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TableProperties](
		tablepropertiesFormCallback.probe,
	)
	tablepropertiesFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tablepropertiesFormCallback.CreationMode {
		tablepropertiesFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TablePropertiesFormCallback(
				nil,
				tablepropertiesFormCallback.probe,
			),
		}).Stage(tablepropertiesFormCallback.probe.formStage)
		tableproperties := new(models.TableProperties)
		FillUpForm(tableproperties, newFormGroup, tablepropertiesFormCallback.probe)
		tablepropertiesFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tablepropertiesFormCallback.probe)
}
func __gong__New__TableRowFormCallback(
	tablerow *models.TableRow,
	probe *Probe,
) (tablerowFormCallback *TableRowFormCallback) {
	tablerowFormCallback = new(TableRowFormCallback)
	tablerowFormCallback.probe = probe
	tablerowFormCallback.tablerow = tablerow

	tablerowFormCallback.CreationMode = (tablerow == nil)

	return
}

type TableRowFormCallback struct {
	tablerow *models.TableRow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (tablerowFormCallback *TableRowFormCallback) OnSave() {

	log.Println("TableRowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablerowFormCallback.probe.formStage.Checkout()

	if tablerowFormCallback.tablerow == nil {
		tablerowFormCallback.tablerow = new(models.TableRow).Stage(tablerowFormCallback.probe.stageOfInterest)
	}
	tablerow_ := tablerowFormCallback.tablerow
	_ = tablerow_

	// get the formGroup
	formGroup := tablerowFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tablerow_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tablerow_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tablerow_.Node), tablerowFormCallback.probe.stageOfInterest, formDiv)
		case "Table:TableRows":
			// we need to retrieve the field owner before the change
			var pastTableOwner *models.Table
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "TableRows"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				tablerowFormCallback.probe.stageOfInterest,
				tablerowFormCallback.probe.backRepoOfInterest,
				tablerow_,
				&rf)

			if reverseFieldOwner != nil {
				pastTableOwner = reverseFieldOwner.(*models.Table)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTableOwner != nil {
					idx := slices.Index(pastTableOwner.TableRows, tablerow_)
					pastTableOwner.TableRows = slices.Delete(pastTableOwner.TableRows, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _table := range *models.GetGongstructInstancesSet[models.Table](tablerowFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _table.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTableOwner := _table // we have a match
						if pastTableOwner != nil {
							if newTableOwner != pastTableOwner {
								idx := slices.Index(pastTableOwner.TableRows, tablerow_)
								pastTableOwner.TableRows = slices.Delete(pastTableOwner.TableRows, idx, idx+1)
								newTableOwner.TableRows = append(newTableOwner.TableRows, tablerow_)
							}
						} else {
							newTableOwner.TableRows = append(newTableOwner.TableRows, tablerow_)
						}
					}
				}
			}
		}
	}

	tablerowFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TableRow](
		tablerowFormCallback.probe,
	)
	tablerowFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tablerowFormCallback.CreationMode {
		tablerowFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TableRowFormCallback(
				nil,
				tablerowFormCallback.probe,
			),
		}).Stage(tablerowFormCallback.probe.formStage)
		tablerow := new(models.TableRow)
		FillUpForm(tablerow, newFormGroup, tablerowFormCallback.probe)
		tablerowFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tablerowFormCallback.probe)
}
func __gong__New__TableStyleFormCallback(
	tablestyle *models.TableStyle,
	probe *Probe,
) (tablestyleFormCallback *TableStyleFormCallback) {
	tablestyleFormCallback = new(TableStyleFormCallback)
	tablestyleFormCallback.probe = probe
	tablestyleFormCallback.tablestyle = tablestyle

	tablestyleFormCallback.CreationMode = (tablestyle == nil)

	return
}

type TableStyleFormCallback struct {
	tablestyle *models.TableStyle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (tablestyleFormCallback *TableStyleFormCallback) OnSave() {

	log.Println("TableStyleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablestyleFormCallback.probe.formStage.Checkout()

	if tablestyleFormCallback.tablestyle == nil {
		tablestyleFormCallback.tablestyle = new(models.TableStyle).Stage(tablestyleFormCallback.probe.stageOfInterest)
	}
	tablestyle_ := tablestyleFormCallback.tablestyle
	_ = tablestyle_

	// get the formGroup
	formGroup := tablestyleFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tablestyle_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tablestyle_.Node), tablestyleFormCallback.probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tablestyle_.Content), formDiv)
		case "Val":
			FormDivBasicFieldToField(&(tablestyle_.Val), formDiv)
		}
	}

	tablestyleFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.TableStyle](
		tablestyleFormCallback.probe,
	)
	tablestyleFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if tablestyleFormCallback.CreationMode {
		tablestyleFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TableStyleFormCallback(
				nil,
				tablestyleFormCallback.probe,
			),
		}).Stage(tablestyleFormCallback.probe.formStage)
		tablestyle := new(models.TableStyle)
		FillUpForm(tablestyle, newFormGroup, tablestyleFormCallback.probe)
		tablestyleFormCallback.probe.formStage.Commit()
	}

	fillUpTree(tablestyleFormCallback.probe)
}
func __gong__New__TextFormCallback(
	text *models.Text,
	probe *Probe,
) (textFormCallback *TextFormCallback) {
	textFormCallback = new(TextFormCallback)
	textFormCallback.probe = probe
	textFormCallback.text = text

	textFormCallback.CreationMode = (text == nil)

	return
}

type TextFormCallback struct {
	text *models.Text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (textFormCallback *TextFormCallback) OnSave() {

	log.Println("TextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	textFormCallback.probe.formStage.Checkout()

	if textFormCallback.text == nil {
		textFormCallback.text = new(models.Text).Stage(textFormCallback.probe.stageOfInterest)
	}
	text_ := textFormCallback.text
	_ = text_

	// get the formGroup
	formGroup := textFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(text_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(text_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(text_.Node), textFormCallback.probe.stageOfInterest, formDiv)
		case "PreserveWhiteSpace":
			FormDivBasicFieldToField(&(text_.PreserveWhiteSpace), formDiv)
		case "EnclosingRune":
			FormDivSelectFieldToField(&(text_.EnclosingRune), textFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	textFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Text](
		textFormCallback.probe,
	)
	textFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if textFormCallback.CreationMode {
		textFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TextFormCallback(
				nil,
				textFormCallback.probe,
			),
		}).Stage(textFormCallback.probe.formStage)
		text := new(models.Text)
		FillUpForm(text, newFormGroup, textFormCallback.probe)
		textFormCallback.probe.formStage.Commit()
	}

	fillUpTree(textFormCallback.probe)
}
