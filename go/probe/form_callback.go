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
	playground *Playground,
) (bodyFormCallback *BodyFormCallback) {
	bodyFormCallback = new(BodyFormCallback)
	bodyFormCallback.playground = playground
	bodyFormCallback.body = body

	bodyFormCallback.CreationMode = (body == nil)

	return
}

type BodyFormCallback struct {
	body *models.Body

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (bodyFormCallback *BodyFormCallback) OnSave() {

	log.Println("BodyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bodyFormCallback.playground.formStage.Checkout()

	if bodyFormCallback.body == nil {
		bodyFormCallback.body = new(models.Body).Stage(bodyFormCallback.playground.stageOfInterest)
	}
	body_ := bodyFormCallback.body
	_ = body_

	// get the formGroup
	formGroup := bodyFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(body_.Name), formDiv)
		case "LastParagraph":
			FormDivSelectFieldToField(&(body_.LastParagraph), bodyFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	bodyFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Body](
		bodyFormCallback.playground,
	)
	bodyFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if bodyFormCallback.CreationMode {
		bodyFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__BodyFormCallback(
				nil,
				bodyFormCallback.playground,
			),
		}).Stage(bodyFormCallback.playground.formStage)
		body := new(models.Body)
		FillUpForm(body, newFormGroup, bodyFormCallback.playground)
		bodyFormCallback.playground.formStage.Commit()
	}

	fillUpTree(bodyFormCallback.playground)
}
func __gong__New__DocumentFormCallback(
	document *models.Document,
	playground *Playground,
) (documentFormCallback *DocumentFormCallback) {
	documentFormCallback = new(DocumentFormCallback)
	documentFormCallback.playground = playground
	documentFormCallback.document = document

	documentFormCallback.CreationMode = (document == nil)

	return
}

type DocumentFormCallback struct {
	document *models.Document

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (documentFormCallback *DocumentFormCallback) OnSave() {

	log.Println("DocumentFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	documentFormCallback.playground.formStage.Checkout()

	if documentFormCallback.document == nil {
		documentFormCallback.document = new(models.Document).Stage(documentFormCallback.playground.stageOfInterest)
	}
	document_ := documentFormCallback.document
	_ = document_

	// get the formGroup
	formGroup := documentFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(document_.Name), formDiv)
		case "File":
			FormDivSelectFieldToField(&(document_.File), documentFormCallback.playground.stageOfInterest, formDiv)
		case "Root":
			FormDivSelectFieldToField(&(document_.Root), documentFormCallback.playground.stageOfInterest, formDiv)
		case "Body":
			FormDivSelectFieldToField(&(document_.Body), documentFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	documentFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Document](
		documentFormCallback.playground,
	)
	documentFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if documentFormCallback.CreationMode {
		documentFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DocumentFormCallback(
				nil,
				documentFormCallback.playground,
			),
		}).Stage(documentFormCallback.playground.formStage)
		document := new(models.Document)
		FillUpForm(document, newFormGroup, documentFormCallback.playground)
		documentFormCallback.playground.formStage.Commit()
	}

	fillUpTree(documentFormCallback.playground)
}
func __gong__New__DocxFormCallback(
	docx *models.Docx,
	playground *Playground,
) (docxFormCallback *DocxFormCallback) {
	docxFormCallback = new(DocxFormCallback)
	docxFormCallback.playground = playground
	docxFormCallback.docx = docx

	docxFormCallback.CreationMode = (docx == nil)

	return
}

type DocxFormCallback struct {
	docx *models.Docx

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (docxFormCallback *DocxFormCallback) OnSave() {

	log.Println("DocxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	docxFormCallback.playground.formStage.Checkout()

	if docxFormCallback.docx == nil {
		docxFormCallback.docx = new(models.Docx).Stage(docxFormCallback.playground.stageOfInterest)
	}
	docx_ := docxFormCallback.docx
	_ = docx_

	// get the formGroup
	formGroup := docxFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(docx_.Name), formDiv)
		case "Document":
			FormDivSelectFieldToField(&(docx_.Document), docxFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	docxFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Docx](
		docxFormCallback.playground,
	)
	docxFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if docxFormCallback.CreationMode {
		docxFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DocxFormCallback(
				nil,
				docxFormCallback.playground,
			),
		}).Stage(docxFormCallback.playground.formStage)
		docx := new(models.Docx)
		FillUpForm(docx, newFormGroup, docxFormCallback.playground)
		docxFormCallback.playground.formStage.Commit()
	}

	fillUpTree(docxFormCallback.playground)
}
func __gong__New__FileFormCallback(
	file *models.File,
	playground *Playground,
) (fileFormCallback *FileFormCallback) {
	fileFormCallback = new(FileFormCallback)
	fileFormCallback.playground = playground
	fileFormCallback.file = file

	fileFormCallback.CreationMode = (file == nil)

	return
}

type FileFormCallback struct {
	file *models.File

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (fileFormCallback *FileFormCallback) OnSave() {

	log.Println("FileFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fileFormCallback.playground.formStage.Checkout()

	if fileFormCallback.file == nil {
		fileFormCallback.file = new(models.File).Stage(fileFormCallback.playground.stageOfInterest)
	}
	file_ := fileFormCallback.file
	_ = file_

	// get the formGroup
	formGroup := fileFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
				fileFormCallback.playground.stageOfInterest,
				fileFormCallback.playground.backRepoOfInterest,
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
				for _docx := range *models.GetGongstructInstancesSet[models.Docx](fileFormCallback.playground.stageOfInterest) {

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

	fileFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.File](
		fileFormCallback.playground,
	)
	fileFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if fileFormCallback.CreationMode {
		fileFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FileFormCallback(
				nil,
				fileFormCallback.playground,
			),
		}).Stage(fileFormCallback.playground.formStage)
		file := new(models.File)
		FillUpForm(file, newFormGroup, fileFormCallback.playground)
		fileFormCallback.playground.formStage.Commit()
	}

	fillUpTree(fileFormCallback.playground)
}
func __gong__New__NodeFormCallback(
	node *models.Node,
	playground *Playground,
) (nodeFormCallback *NodeFormCallback) {
	nodeFormCallback = new(NodeFormCallback)
	nodeFormCallback.playground = playground
	nodeFormCallback.node = node

	nodeFormCallback.CreationMode = (node == nil)

	return
}

type NodeFormCallback struct {
	node *models.Node

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (nodeFormCallback *NodeFormCallback) OnSave() {

	log.Println("NodeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	nodeFormCallback.playground.formStage.Checkout()

	if nodeFormCallback.node == nil {
		nodeFormCallback.node = new(models.Node).Stage(nodeFormCallback.playground.stageOfInterest)
	}
	node_ := nodeFormCallback.node
	_ = node_

	// get the formGroup
	formGroup := nodeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
				nodeFormCallback.playground.stageOfInterest,
				nodeFormCallback.playground.backRepoOfInterest,
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
				for _node := range *models.GetGongstructInstancesSet[models.Node](nodeFormCallback.playground.stageOfInterest) {

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

	nodeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Node](
		nodeFormCallback.playground,
	)
	nodeFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if nodeFormCallback.CreationMode {
		nodeFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__NodeFormCallback(
				nil,
				nodeFormCallback.playground,
			),
		}).Stage(nodeFormCallback.playground.formStage)
		node := new(models.Node)
		FillUpForm(node, newFormGroup, nodeFormCallback.playground)
		nodeFormCallback.playground.formStage.Commit()
	}

	fillUpTree(nodeFormCallback.playground)
}
func __gong__New__ParagraphFormCallback(
	paragraph *models.Paragraph,
	playground *Playground,
) (paragraphFormCallback *ParagraphFormCallback) {
	paragraphFormCallback = new(ParagraphFormCallback)
	paragraphFormCallback.playground = playground
	paragraphFormCallback.paragraph = paragraph

	paragraphFormCallback.CreationMode = (paragraph == nil)

	return
}

type ParagraphFormCallback struct {
	paragraph *models.Paragraph

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (paragraphFormCallback *ParagraphFormCallback) OnSave() {

	log.Println("ParagraphFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	paragraphFormCallback.playground.formStage.Checkout()

	if paragraphFormCallback.paragraph == nil {
		paragraphFormCallback.paragraph = new(models.Paragraph).Stage(paragraphFormCallback.playground.stageOfInterest)
	}
	paragraph_ := paragraphFormCallback.paragraph
	_ = paragraph_

	// get the formGroup
	formGroup := paragraphFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(paragraph_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(paragraph_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(paragraph_.Node), paragraphFormCallback.playground.stageOfInterest, formDiv)
		case "ParagraphProperties":
			FormDivSelectFieldToField(&(paragraph_.ParagraphProperties), paragraphFormCallback.playground.stageOfInterest, formDiv)
		case "Text":
			FormDivBasicFieldToField(&(paragraph_.Text), formDiv)
		case "Next":
			FormDivSelectFieldToField(&(paragraph_.Next), paragraphFormCallback.playground.stageOfInterest, formDiv)
		case "Previous":
			FormDivSelectFieldToField(&(paragraph_.Previous), paragraphFormCallback.playground.stageOfInterest, formDiv)
		case "EnclosingBody":
			FormDivSelectFieldToField(&(paragraph_.EnclosingBody), paragraphFormCallback.playground.stageOfInterest, formDiv)
		case "EnclosingTableColumn":
			FormDivSelectFieldToField(&(paragraph_.EnclosingTableColumn), paragraphFormCallback.playground.stageOfInterest, formDiv)
		case "Body:Paragraphs":
			// we need to retrieve the field owner before the change
			var pastBodyOwner *models.Body
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Body"
			rf.Fieldname = "Paragraphs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				paragraphFormCallback.playground.stageOfInterest,
				paragraphFormCallback.playground.backRepoOfInterest,
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
				for _body := range *models.GetGongstructInstancesSet[models.Body](paragraphFormCallback.playground.stageOfInterest) {

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
				paragraphFormCallback.playground.stageOfInterest,
				paragraphFormCallback.playground.backRepoOfInterest,
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
				for _tablecolumn := range *models.GetGongstructInstancesSet[models.TableColumn](paragraphFormCallback.playground.stageOfInterest) {

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

	paragraphFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Paragraph](
		paragraphFormCallback.playground,
	)
	paragraphFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if paragraphFormCallback.CreationMode {
		paragraphFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ParagraphFormCallback(
				nil,
				paragraphFormCallback.playground,
			),
		}).Stage(paragraphFormCallback.playground.formStage)
		paragraph := new(models.Paragraph)
		FillUpForm(paragraph, newFormGroup, paragraphFormCallback.playground)
		paragraphFormCallback.playground.formStage.Commit()
	}

	fillUpTree(paragraphFormCallback.playground)
}
func __gong__New__ParagraphPropertiesFormCallback(
	paragraphproperties *models.ParagraphProperties,
	playground *Playground,
) (paragraphpropertiesFormCallback *ParagraphPropertiesFormCallback) {
	paragraphpropertiesFormCallback = new(ParagraphPropertiesFormCallback)
	paragraphpropertiesFormCallback.playground = playground
	paragraphpropertiesFormCallback.paragraphproperties = paragraphproperties

	paragraphpropertiesFormCallback.CreationMode = (paragraphproperties == nil)

	return
}

type ParagraphPropertiesFormCallback struct {
	paragraphproperties *models.ParagraphProperties

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (paragraphpropertiesFormCallback *ParagraphPropertiesFormCallback) OnSave() {

	log.Println("ParagraphPropertiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	paragraphpropertiesFormCallback.playground.formStage.Checkout()

	if paragraphpropertiesFormCallback.paragraphproperties == nil {
		paragraphpropertiesFormCallback.paragraphproperties = new(models.ParagraphProperties).Stage(paragraphpropertiesFormCallback.playground.stageOfInterest)
	}
	paragraphproperties_ := paragraphpropertiesFormCallback.paragraphproperties
	_ = paragraphproperties_

	// get the formGroup
	formGroup := paragraphpropertiesFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(paragraphproperties_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(paragraphproperties_.Content), formDiv)
		case "ParagraphStyle":
			FormDivSelectFieldToField(&(paragraphproperties_.ParagraphStyle), paragraphpropertiesFormCallback.playground.stageOfInterest, formDiv)
		case "Node":
			FormDivSelectFieldToField(&(paragraphproperties_.Node), paragraphpropertiesFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	paragraphpropertiesFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.ParagraphProperties](
		paragraphpropertiesFormCallback.playground,
	)
	paragraphpropertiesFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if paragraphpropertiesFormCallback.CreationMode {
		paragraphpropertiesFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ParagraphPropertiesFormCallback(
				nil,
				paragraphpropertiesFormCallback.playground,
			),
		}).Stage(paragraphpropertiesFormCallback.playground.formStage)
		paragraphproperties := new(models.ParagraphProperties)
		FillUpForm(paragraphproperties, newFormGroup, paragraphpropertiesFormCallback.playground)
		paragraphpropertiesFormCallback.playground.formStage.Commit()
	}

	fillUpTree(paragraphpropertiesFormCallback.playground)
}
func __gong__New__ParagraphStyleFormCallback(
	paragraphstyle *models.ParagraphStyle,
	playground *Playground,
) (paragraphstyleFormCallback *ParagraphStyleFormCallback) {
	paragraphstyleFormCallback = new(ParagraphStyleFormCallback)
	paragraphstyleFormCallback.playground = playground
	paragraphstyleFormCallback.paragraphstyle = paragraphstyle

	paragraphstyleFormCallback.CreationMode = (paragraphstyle == nil)

	return
}

type ParagraphStyleFormCallback struct {
	paragraphstyle *models.ParagraphStyle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (paragraphstyleFormCallback *ParagraphStyleFormCallback) OnSave() {

	log.Println("ParagraphStyleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	paragraphstyleFormCallback.playground.formStage.Checkout()

	if paragraphstyleFormCallback.paragraphstyle == nil {
		paragraphstyleFormCallback.paragraphstyle = new(models.ParagraphStyle).Stage(paragraphstyleFormCallback.playground.stageOfInterest)
	}
	paragraphstyle_ := paragraphstyleFormCallback.paragraphstyle
	_ = paragraphstyle_

	// get the formGroup
	formGroup := paragraphstyleFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(paragraphstyle_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(paragraphstyle_.Node), paragraphstyleFormCallback.playground.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(paragraphstyle_.Content), formDiv)
		case "ValAttr":
			FormDivBasicFieldToField(&(paragraphstyle_.ValAttr), formDiv)
		}
	}

	paragraphstyleFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.ParagraphStyle](
		paragraphstyleFormCallback.playground,
	)
	paragraphstyleFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if paragraphstyleFormCallback.CreationMode {
		paragraphstyleFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ParagraphStyleFormCallback(
				nil,
				paragraphstyleFormCallback.playground,
			),
		}).Stage(paragraphstyleFormCallback.playground.formStage)
		paragraphstyle := new(models.ParagraphStyle)
		FillUpForm(paragraphstyle, newFormGroup, paragraphstyleFormCallback.playground)
		paragraphstyleFormCallback.playground.formStage.Commit()
	}

	fillUpTree(paragraphstyleFormCallback.playground)
}
func __gong__New__RuneFormCallback(
	rune *models.Rune,
	playground *Playground,
) (runeFormCallback *RuneFormCallback) {
	runeFormCallback = new(RuneFormCallback)
	runeFormCallback.playground = playground
	runeFormCallback.rune = rune

	runeFormCallback.CreationMode = (rune == nil)

	return
}

type RuneFormCallback struct {
	rune *models.Rune

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (runeFormCallback *RuneFormCallback) OnSave() {

	log.Println("RuneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	runeFormCallback.playground.formStage.Checkout()

	if runeFormCallback.rune == nil {
		runeFormCallback.rune = new(models.Rune).Stage(runeFormCallback.playground.stageOfInterest)
	}
	rune_ := runeFormCallback.rune
	_ = rune_

	// get the formGroup
	formGroup := runeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rune_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(rune_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(rune_.Node), runeFormCallback.playground.stageOfInterest, formDiv)
		case "Text":
			FormDivSelectFieldToField(&(rune_.Text), runeFormCallback.playground.stageOfInterest, formDiv)
		case "RuneProperties":
			FormDivSelectFieldToField(&(rune_.RuneProperties), runeFormCallback.playground.stageOfInterest, formDiv)
		case "EnclosingParagraph":
			FormDivSelectFieldToField(&(rune_.EnclosingParagraph), runeFormCallback.playground.stageOfInterest, formDiv)
		case "Paragraph:Runes":
			// we need to retrieve the field owner before the change
			var pastParagraphOwner *models.Paragraph
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Paragraph"
			rf.Fieldname = "Runes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				runeFormCallback.playground.stageOfInterest,
				runeFormCallback.playground.backRepoOfInterest,
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
				for _paragraph := range *models.GetGongstructInstancesSet[models.Paragraph](runeFormCallback.playground.stageOfInterest) {

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

	runeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Rune](
		runeFormCallback.playground,
	)
	runeFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if runeFormCallback.CreationMode {
		runeFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RuneFormCallback(
				nil,
				runeFormCallback.playground,
			),
		}).Stage(runeFormCallback.playground.formStage)
		rune := new(models.Rune)
		FillUpForm(rune, newFormGroup, runeFormCallback.playground)
		runeFormCallback.playground.formStage.Commit()
	}

	fillUpTree(runeFormCallback.playground)
}
func __gong__New__RunePropertiesFormCallback(
	runeproperties *models.RuneProperties,
	playground *Playground,
) (runepropertiesFormCallback *RunePropertiesFormCallback) {
	runepropertiesFormCallback = new(RunePropertiesFormCallback)
	runepropertiesFormCallback.playground = playground
	runepropertiesFormCallback.runeproperties = runeproperties

	runepropertiesFormCallback.CreationMode = (runeproperties == nil)

	return
}

type RunePropertiesFormCallback struct {
	runeproperties *models.RuneProperties

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (runepropertiesFormCallback *RunePropertiesFormCallback) OnSave() {

	log.Println("RunePropertiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	runepropertiesFormCallback.playground.formStage.Checkout()

	if runepropertiesFormCallback.runeproperties == nil {
		runepropertiesFormCallback.runeproperties = new(models.RuneProperties).Stage(runepropertiesFormCallback.playground.stageOfInterest)
	}
	runeproperties_ := runepropertiesFormCallback.runeproperties
	_ = runeproperties_

	// get the formGroup
	formGroup := runepropertiesFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(runeproperties_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(runeproperties_.Node), runepropertiesFormCallback.playground.stageOfInterest, formDiv)
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

	runepropertiesFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.RuneProperties](
		runepropertiesFormCallback.playground,
	)
	runepropertiesFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if runepropertiesFormCallback.CreationMode {
		runepropertiesFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RunePropertiesFormCallback(
				nil,
				runepropertiesFormCallback.playground,
			),
		}).Stage(runepropertiesFormCallback.playground.formStage)
		runeproperties := new(models.RuneProperties)
		FillUpForm(runeproperties, newFormGroup, runepropertiesFormCallback.playground)
		runepropertiesFormCallback.playground.formStage.Commit()
	}

	fillUpTree(runepropertiesFormCallback.playground)
}
func __gong__New__TableFormCallback(
	table *models.Table,
	playground *Playground,
) (tableFormCallback *TableFormCallback) {
	tableFormCallback = new(TableFormCallback)
	tableFormCallback.playground = playground
	tableFormCallback.table = table

	tableFormCallback.CreationMode = (table == nil)

	return
}

type TableFormCallback struct {
	table *models.Table

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (tableFormCallback *TableFormCallback) OnSave() {

	log.Println("TableFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tableFormCallback.playground.formStage.Checkout()

	if tableFormCallback.table == nil {
		tableFormCallback.table = new(models.Table).Stage(tableFormCallback.playground.stageOfInterest)
	}
	table_ := tableFormCallback.table
	_ = table_

	// get the formGroup
	formGroup := tableFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(table_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(table_.Node), tableFormCallback.playground.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(table_.Content), formDiv)
		case "TableProperties":
			FormDivSelectFieldToField(&(table_.TableProperties), tableFormCallback.playground.stageOfInterest, formDiv)
		case "Body:Tables":
			// we need to retrieve the field owner before the change
			var pastBodyOwner *models.Body
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Body"
			rf.Fieldname = "Tables"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				tableFormCallback.playground.stageOfInterest,
				tableFormCallback.playground.backRepoOfInterest,
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
				for _body := range *models.GetGongstructInstancesSet[models.Body](tableFormCallback.playground.stageOfInterest) {

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

	tableFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Table](
		tableFormCallback.playground,
	)
	tableFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if tableFormCallback.CreationMode {
		tableFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TableFormCallback(
				nil,
				tableFormCallback.playground,
			),
		}).Stage(tableFormCallback.playground.formStage)
		table := new(models.Table)
		FillUpForm(table, newFormGroup, tableFormCallback.playground)
		tableFormCallback.playground.formStage.Commit()
	}

	fillUpTree(tableFormCallback.playground)
}
func __gong__New__TableColumnFormCallback(
	tablecolumn *models.TableColumn,
	playground *Playground,
) (tablecolumnFormCallback *TableColumnFormCallback) {
	tablecolumnFormCallback = new(TableColumnFormCallback)
	tablecolumnFormCallback.playground = playground
	tablecolumnFormCallback.tablecolumn = tablecolumn

	tablecolumnFormCallback.CreationMode = (tablecolumn == nil)

	return
}

type TableColumnFormCallback struct {
	tablecolumn *models.TableColumn

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (tablecolumnFormCallback *TableColumnFormCallback) OnSave() {

	log.Println("TableColumnFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablecolumnFormCallback.playground.formStage.Checkout()

	if tablecolumnFormCallback.tablecolumn == nil {
		tablecolumnFormCallback.tablecolumn = new(models.TableColumn).Stage(tablecolumnFormCallback.playground.stageOfInterest)
	}
	tablecolumn_ := tablecolumnFormCallback.tablecolumn
	_ = tablecolumn_

	// get the formGroup
	formGroup := tablecolumnFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tablecolumn_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tablecolumn_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tablecolumn_.Node), tablecolumnFormCallback.playground.stageOfInterest, formDiv)
		case "TableRow:TableColumns":
			// we need to retrieve the field owner before the change
			var pastTableRowOwner *models.TableRow
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TableRow"
			rf.Fieldname = "TableColumns"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				tablecolumnFormCallback.playground.stageOfInterest,
				tablecolumnFormCallback.playground.backRepoOfInterest,
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
				for _tablerow := range *models.GetGongstructInstancesSet[models.TableRow](tablecolumnFormCallback.playground.stageOfInterest) {

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

	tablecolumnFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.TableColumn](
		tablecolumnFormCallback.playground,
	)
	tablecolumnFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if tablecolumnFormCallback.CreationMode {
		tablecolumnFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TableColumnFormCallback(
				nil,
				tablecolumnFormCallback.playground,
			),
		}).Stage(tablecolumnFormCallback.playground.formStage)
		tablecolumn := new(models.TableColumn)
		FillUpForm(tablecolumn, newFormGroup, tablecolumnFormCallback.playground)
		tablecolumnFormCallback.playground.formStage.Commit()
	}

	fillUpTree(tablecolumnFormCallback.playground)
}
func __gong__New__TablePropertiesFormCallback(
	tableproperties *models.TableProperties,
	playground *Playground,
) (tablepropertiesFormCallback *TablePropertiesFormCallback) {
	tablepropertiesFormCallback = new(TablePropertiesFormCallback)
	tablepropertiesFormCallback.playground = playground
	tablepropertiesFormCallback.tableproperties = tableproperties

	tablepropertiesFormCallback.CreationMode = (tableproperties == nil)

	return
}

type TablePropertiesFormCallback struct {
	tableproperties *models.TableProperties

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (tablepropertiesFormCallback *TablePropertiesFormCallback) OnSave() {

	log.Println("TablePropertiesFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablepropertiesFormCallback.playground.formStage.Checkout()

	if tablepropertiesFormCallback.tableproperties == nil {
		tablepropertiesFormCallback.tableproperties = new(models.TableProperties).Stage(tablepropertiesFormCallback.playground.stageOfInterest)
	}
	tableproperties_ := tablepropertiesFormCallback.tableproperties
	_ = tableproperties_

	// get the formGroup
	formGroup := tablepropertiesFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tableproperties_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tableproperties_.Node), tablepropertiesFormCallback.playground.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tableproperties_.Content), formDiv)
		case "TableStyle":
			FormDivSelectFieldToField(&(tableproperties_.TableStyle), tablepropertiesFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	tablepropertiesFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.TableProperties](
		tablepropertiesFormCallback.playground,
	)
	tablepropertiesFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if tablepropertiesFormCallback.CreationMode {
		tablepropertiesFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TablePropertiesFormCallback(
				nil,
				tablepropertiesFormCallback.playground,
			),
		}).Stage(tablepropertiesFormCallback.playground.formStage)
		tableproperties := new(models.TableProperties)
		FillUpForm(tableproperties, newFormGroup, tablepropertiesFormCallback.playground)
		tablepropertiesFormCallback.playground.formStage.Commit()
	}

	fillUpTree(tablepropertiesFormCallback.playground)
}
func __gong__New__TableRowFormCallback(
	tablerow *models.TableRow,
	playground *Playground,
) (tablerowFormCallback *TableRowFormCallback) {
	tablerowFormCallback = new(TableRowFormCallback)
	tablerowFormCallback.playground = playground
	tablerowFormCallback.tablerow = tablerow

	tablerowFormCallback.CreationMode = (tablerow == nil)

	return
}

type TableRowFormCallback struct {
	tablerow *models.TableRow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (tablerowFormCallback *TableRowFormCallback) OnSave() {

	log.Println("TableRowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablerowFormCallback.playground.formStage.Checkout()

	if tablerowFormCallback.tablerow == nil {
		tablerowFormCallback.tablerow = new(models.TableRow).Stage(tablerowFormCallback.playground.stageOfInterest)
	}
	tablerow_ := tablerowFormCallback.tablerow
	_ = tablerow_

	// get the formGroup
	formGroup := tablerowFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tablerow_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tablerow_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tablerow_.Node), tablerowFormCallback.playground.stageOfInterest, formDiv)
		case "Table:TableRows":
			// we need to retrieve the field owner before the change
			var pastTableOwner *models.Table
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "TableRows"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				tablerowFormCallback.playground.stageOfInterest,
				tablerowFormCallback.playground.backRepoOfInterest,
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
				for _table := range *models.GetGongstructInstancesSet[models.Table](tablerowFormCallback.playground.stageOfInterest) {

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

	tablerowFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.TableRow](
		tablerowFormCallback.playground,
	)
	tablerowFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if tablerowFormCallback.CreationMode {
		tablerowFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TableRowFormCallback(
				nil,
				tablerowFormCallback.playground,
			),
		}).Stage(tablerowFormCallback.playground.formStage)
		tablerow := new(models.TableRow)
		FillUpForm(tablerow, newFormGroup, tablerowFormCallback.playground)
		tablerowFormCallback.playground.formStage.Commit()
	}

	fillUpTree(tablerowFormCallback.playground)
}
func __gong__New__TableStyleFormCallback(
	tablestyle *models.TableStyle,
	playground *Playground,
) (tablestyleFormCallback *TableStyleFormCallback) {
	tablestyleFormCallback = new(TableStyleFormCallback)
	tablestyleFormCallback.playground = playground
	tablestyleFormCallback.tablestyle = tablestyle

	tablestyleFormCallback.CreationMode = (tablestyle == nil)

	return
}

type TableStyleFormCallback struct {
	tablestyle *models.TableStyle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (tablestyleFormCallback *TableStyleFormCallback) OnSave() {

	log.Println("TableStyleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	tablestyleFormCallback.playground.formStage.Checkout()

	if tablestyleFormCallback.tablestyle == nil {
		tablestyleFormCallback.tablestyle = new(models.TableStyle).Stage(tablestyleFormCallback.playground.stageOfInterest)
	}
	tablestyle_ := tablestyleFormCallback.tablestyle
	_ = tablestyle_

	// get the formGroup
	formGroup := tablestyleFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tablestyle_.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(tablestyle_.Node), tablestyleFormCallback.playground.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(tablestyle_.Content), formDiv)
		case "Val":
			FormDivBasicFieldToField(&(tablestyle_.Val), formDiv)
		}
	}

	tablestyleFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.TableStyle](
		tablestyleFormCallback.playground,
	)
	tablestyleFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if tablestyleFormCallback.CreationMode {
		tablestyleFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TableStyleFormCallback(
				nil,
				tablestyleFormCallback.playground,
			),
		}).Stage(tablestyleFormCallback.playground.formStage)
		tablestyle := new(models.TableStyle)
		FillUpForm(tablestyle, newFormGroup, tablestyleFormCallback.playground)
		tablestyleFormCallback.playground.formStage.Commit()
	}

	fillUpTree(tablestyleFormCallback.playground)
}
func __gong__New__TextFormCallback(
	text *models.Text,
	playground *Playground,
) (textFormCallback *TextFormCallback) {
	textFormCallback = new(TextFormCallback)
	textFormCallback.playground = playground
	textFormCallback.text = text

	textFormCallback.CreationMode = (text == nil)

	return
}

type TextFormCallback struct {
	text *models.Text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (textFormCallback *TextFormCallback) OnSave() {

	log.Println("TextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	textFormCallback.playground.formStage.Checkout()

	if textFormCallback.text == nil {
		textFormCallback.text = new(models.Text).Stage(textFormCallback.playground.stageOfInterest)
	}
	text_ := textFormCallback.text
	_ = text_

	// get the formGroup
	formGroup := textFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(text_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(text_.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(text_.Node), textFormCallback.playground.stageOfInterest, formDiv)
		case "PreserveWhiteSpace":
			FormDivBasicFieldToField(&(text_.PreserveWhiteSpace), formDiv)
		case "EnclosingRune":
			FormDivSelectFieldToField(&(text_.EnclosingRune), textFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	textFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Text](
		textFormCallback.playground,
	)
	textFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if textFormCallback.CreationMode {
		textFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TextFormCallback(
				nil,
				textFormCallback.playground,
			),
		}).Stage(textFormCallback.playground.formStage)
		text := new(models.Text)
		FillUpForm(text, newFormGroup, textFormCallback.playground)
		textFormCallback.playground.formStage.Commit()
	}

	fillUpTree(textFormCallback.playground)
}
