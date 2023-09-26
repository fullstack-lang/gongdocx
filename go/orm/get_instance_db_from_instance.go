// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongdocx/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | BodyDB | DocumentDB | DocxDB | FileDB | NodeDB | ParagraphDB | ParagraphPropertiesDB | ParagraphStyleDB | RuneDB | RunePropertiesDB | TableDB | TableColumnDB | TablePropertiesDB | TableRowDB | TableStyleDB | TextDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Body:
		bodyInstance := any(concreteInstance).(*models.Body)
		ret2 := backRepo.BackRepoBody.GetBodyDBFromBodyPtr(bodyInstance)
		ret = any(ret2).(*T2)
	case *models.Document:
		documentInstance := any(concreteInstance).(*models.Document)
		ret2 := backRepo.BackRepoDocument.GetDocumentDBFromDocumentPtr(documentInstance)
		ret = any(ret2).(*T2)
	case *models.Docx:
		docxInstance := any(concreteInstance).(*models.Docx)
		ret2 := backRepo.BackRepoDocx.GetDocxDBFromDocxPtr(docxInstance)
		ret = any(ret2).(*T2)
	case *models.File:
		fileInstance := any(concreteInstance).(*models.File)
		ret2 := backRepo.BackRepoFile.GetFileDBFromFilePtr(fileInstance)
		ret = any(ret2).(*T2)
	case *models.Node:
		nodeInstance := any(concreteInstance).(*models.Node)
		ret2 := backRepo.BackRepoNode.GetNodeDBFromNodePtr(nodeInstance)
		ret = any(ret2).(*T2)
	case *models.Paragraph:
		paragraphInstance := any(concreteInstance).(*models.Paragraph)
		ret2 := backRepo.BackRepoParagraph.GetParagraphDBFromParagraphPtr(paragraphInstance)
		ret = any(ret2).(*T2)
	case *models.ParagraphProperties:
		paragraphpropertiesInstance := any(concreteInstance).(*models.ParagraphProperties)
		ret2 := backRepo.BackRepoParagraphProperties.GetParagraphPropertiesDBFromParagraphPropertiesPtr(paragraphpropertiesInstance)
		ret = any(ret2).(*T2)
	case *models.ParagraphStyle:
		paragraphstyleInstance := any(concreteInstance).(*models.ParagraphStyle)
		ret2 := backRepo.BackRepoParagraphStyle.GetParagraphStyleDBFromParagraphStylePtr(paragraphstyleInstance)
		ret = any(ret2).(*T2)
	case *models.Rune:
		runeInstance := any(concreteInstance).(*models.Rune)
		ret2 := backRepo.BackRepoRune.GetRuneDBFromRunePtr(runeInstance)
		ret = any(ret2).(*T2)
	case *models.RuneProperties:
		runepropertiesInstance := any(concreteInstance).(*models.RuneProperties)
		ret2 := backRepo.BackRepoRuneProperties.GetRunePropertiesDBFromRunePropertiesPtr(runepropertiesInstance)
		ret = any(ret2).(*T2)
	case *models.Table:
		tableInstance := any(concreteInstance).(*models.Table)
		ret2 := backRepo.BackRepoTable.GetTableDBFromTablePtr(tableInstance)
		ret = any(ret2).(*T2)
	case *models.TableColumn:
		tablecolumnInstance := any(concreteInstance).(*models.TableColumn)
		ret2 := backRepo.BackRepoTableColumn.GetTableColumnDBFromTableColumnPtr(tablecolumnInstance)
		ret = any(ret2).(*T2)
	case *models.TableProperties:
		tablepropertiesInstance := any(concreteInstance).(*models.TableProperties)
		ret2 := backRepo.BackRepoTableProperties.GetTablePropertiesDBFromTablePropertiesPtr(tablepropertiesInstance)
		ret = any(ret2).(*T2)
	case *models.TableRow:
		tablerowInstance := any(concreteInstance).(*models.TableRow)
		ret2 := backRepo.BackRepoTableRow.GetTableRowDBFromTableRowPtr(tablerowInstance)
		ret = any(ret2).(*T2)
	case *models.TableStyle:
		tablestyleInstance := any(concreteInstance).(*models.TableStyle)
		ret2 := backRepo.BackRepoTableStyle.GetTableStyleDBFromTableStylePtr(tablestyleInstance)
		ret = any(ret2).(*T2)
	case *models.Text:
		textInstance := any(concreteInstance).(*models.Text)
		ret2 := backRepo.BackRepoText.GetTextDBFromTextPtr(textInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Body:
		tmp := GetInstanceDBFromInstance[models.Body, BodyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Document:
		tmp := GetInstanceDBFromInstance[models.Document, DocumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Docx:
		tmp := GetInstanceDBFromInstance[models.Docx, DocxDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.File:
		tmp := GetInstanceDBFromInstance[models.File, FileDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Node:
		tmp := GetInstanceDBFromInstance[models.Node, NodeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Paragraph:
		tmp := GetInstanceDBFromInstance[models.Paragraph, ParagraphDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ParagraphProperties:
		tmp := GetInstanceDBFromInstance[models.ParagraphProperties, ParagraphPropertiesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ParagraphStyle:
		tmp := GetInstanceDBFromInstance[models.ParagraphStyle, ParagraphStyleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Rune:
		tmp := GetInstanceDBFromInstance[models.Rune, RuneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RuneProperties:
		tmp := GetInstanceDBFromInstance[models.RuneProperties, RunePropertiesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Table:
		tmp := GetInstanceDBFromInstance[models.Table, TableDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TableColumn:
		tmp := GetInstanceDBFromInstance[models.TableColumn, TableColumnDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TableProperties:
		tmp := GetInstanceDBFromInstance[models.TableProperties, TablePropertiesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TableRow:
		tmp := GetInstanceDBFromInstance[models.TableRow, TableRowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TableStyle:
		tmp := GetInstanceDBFromInstance[models.TableStyle, TableStyleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Text:
		tmp := GetInstanceDBFromInstance[models.Text, TextDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Body:
		tmp := GetInstanceDBFromInstance[models.Body, BodyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Document:
		tmp := GetInstanceDBFromInstance[models.Document, DocumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Docx:
		tmp := GetInstanceDBFromInstance[models.Docx, DocxDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.File:
		tmp := GetInstanceDBFromInstance[models.File, FileDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Node:
		tmp := GetInstanceDBFromInstance[models.Node, NodeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Paragraph:
		tmp := GetInstanceDBFromInstance[models.Paragraph, ParagraphDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ParagraphProperties:
		tmp := GetInstanceDBFromInstance[models.ParagraphProperties, ParagraphPropertiesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ParagraphStyle:
		tmp := GetInstanceDBFromInstance[models.ParagraphStyle, ParagraphStyleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Rune:
		tmp := GetInstanceDBFromInstance[models.Rune, RuneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RuneProperties:
		tmp := GetInstanceDBFromInstance[models.RuneProperties, RunePropertiesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Table:
		tmp := GetInstanceDBFromInstance[models.Table, TableDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TableColumn:
		tmp := GetInstanceDBFromInstance[models.TableColumn, TableColumnDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TableProperties:
		tmp := GetInstanceDBFromInstance[models.TableProperties, TablePropertiesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TableRow:
		tmp := GetInstanceDBFromInstance[models.TableRow, TableRowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.TableStyle:
		tmp := GetInstanceDBFromInstance[models.TableStyle, TableStyleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Text:
		tmp := GetInstanceDBFromInstance[models.Text, TextDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
