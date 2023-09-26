// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongdocx/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Body:
		tmp := GetInstanceDBFromInstance[models.Body, BodyDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Document:
		tmp := GetInstanceDBFromInstance[models.Document, DocumentDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Docx:
		tmp := GetInstanceDBFromInstance[models.Docx, DocxDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.File:
		tmp := GetInstanceDBFromInstance[models.File, FileDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			case "Files":
				if tmp != nil && tmp.Docx_FilesDBID.Int64 != 0 {
					id := uint(tmp.Docx_FilesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoDocx.Map_DocxDBID_DocxPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Node:
		tmp := GetInstanceDBFromInstance[models.Node, NodeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			case "Nodes":
				if tmp != nil && tmp.Node_NodesDBID.Int64 != 0 {
					id := uint(tmp.Node_NodesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Paragraph:
		tmp := GetInstanceDBFromInstance[models.Paragraph, ParagraphDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			case "Paragraphs":
				if tmp != nil && tmp.Body_ParagraphsDBID.Int64 != 0 {
					id := uint(tmp.Body_ParagraphsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoBody.Map_BodyDBID_BodyPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			case "Paragraphs":
				if tmp != nil && tmp.TableColumn_ParagraphsDBID.Int64 != 0 {
					id := uint(tmp.TableColumn_ParagraphsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTableColumn.Map_TableColumnDBID_TableColumnPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.ParagraphProperties:
		tmp := GetInstanceDBFromInstance[models.ParagraphProperties, ParagraphPropertiesDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.ParagraphStyle:
		tmp := GetInstanceDBFromInstance[models.ParagraphStyle, ParagraphStyleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Rune:
		tmp := GetInstanceDBFromInstance[models.Rune, RuneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			case "Runes":
				if tmp != nil && tmp.Paragraph_RunesDBID.Int64 != 0 {
					id := uint(tmp.Paragraph_RunesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoParagraph.Map_ParagraphDBID_ParagraphPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.RuneProperties:
		tmp := GetInstanceDBFromInstance[models.RuneProperties, RunePropertiesDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Table:
		tmp := GetInstanceDBFromInstance[models.Table, TableDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			case "Tables":
				if tmp != nil && tmp.Body_TablesDBID.Int64 != 0 {
					id := uint(tmp.Body_TablesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoBody.Map_BodyDBID_BodyPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.TableColumn:
		tmp := GetInstanceDBFromInstance[models.TableColumn, TableColumnDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			case "TableColumns":
				if tmp != nil && tmp.TableRow_TableColumnsDBID.Int64 != 0 {
					id := uint(tmp.TableRow_TableColumnsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTableRow.Map_TableRowDBID_TableRowPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.TableProperties:
		tmp := GetInstanceDBFromInstance[models.TableProperties, TablePropertiesDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.TableRow:
		tmp := GetInstanceDBFromInstance[models.TableRow, TableRowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			case "TableRows":
				if tmp != nil && tmp.Table_TableRowsDBID.Int64 != 0 {
					id := uint(tmp.Table_TableRowsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTable.Map_TableDBID_TablePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.TableStyle:
		tmp := GetInstanceDBFromInstance[models.TableStyle, TableStyleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Text:
		tmp := GetInstanceDBFromInstance[models.Text, TextDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Body:
		tmp := GetInstanceDBFromInstance[models.Body, BodyDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Document:
		tmp := GetInstanceDBFromInstance[models.Document, DocumentDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Docx:
		tmp := GetInstanceDBFromInstance[models.Docx, DocxDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.File:
		tmp := GetInstanceDBFromInstance[models.File, FileDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			case "Files":
				if tmp != nil && tmp.Docx_FilesDBID.Int64 != 0 {
					id := uint(tmp.Docx_FilesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoDocx.Map_DocxDBID_DocxPtr[id]
					res = reservePointerTarget
				}
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Node:
		tmp := GetInstanceDBFromInstance[models.Node, NodeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			case "Nodes":
				if tmp != nil && tmp.Node_NodesDBID.Int64 != 0 {
					id := uint(tmp.Node_NodesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[id]
					res = reservePointerTarget
				}
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Paragraph:
		tmp := GetInstanceDBFromInstance[models.Paragraph, ParagraphDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			case "Paragraphs":
				if tmp != nil && tmp.Body_ParagraphsDBID.Int64 != 0 {
					id := uint(tmp.Body_ParagraphsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoBody.Map_BodyDBID_BodyPtr[id]
					res = reservePointerTarget
				}
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			case "Paragraphs":
				if tmp != nil && tmp.TableColumn_ParagraphsDBID.Int64 != 0 {
					id := uint(tmp.TableColumn_ParagraphsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTableColumn.Map_TableColumnDBID_TableColumnPtr[id]
					res = reservePointerTarget
				}
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.ParagraphProperties:
		tmp := GetInstanceDBFromInstance[models.ParagraphProperties, ParagraphPropertiesDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.ParagraphStyle:
		tmp := GetInstanceDBFromInstance[models.ParagraphStyle, ParagraphStyleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Rune:
		tmp := GetInstanceDBFromInstance[models.Rune, RuneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			case "Runes":
				if tmp != nil && tmp.Paragraph_RunesDBID.Int64 != 0 {
					id := uint(tmp.Paragraph_RunesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoParagraph.Map_ParagraphDBID_ParagraphPtr[id]
					res = reservePointerTarget
				}
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.RuneProperties:
		tmp := GetInstanceDBFromInstance[models.RuneProperties, RunePropertiesDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Table:
		tmp := GetInstanceDBFromInstance[models.Table, TableDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			case "Tables":
				if tmp != nil && tmp.Body_TablesDBID.Int64 != 0 {
					id := uint(tmp.Body_TablesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoBody.Map_BodyDBID_BodyPtr[id]
					res = reservePointerTarget
				}
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.TableColumn:
		tmp := GetInstanceDBFromInstance[models.TableColumn, TableColumnDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			case "TableColumns":
				if tmp != nil && tmp.TableRow_TableColumnsDBID.Int64 != 0 {
					id := uint(tmp.TableRow_TableColumnsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTableRow.Map_TableRowDBID_TableRowPtr[id]
					res = reservePointerTarget
				}
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.TableProperties:
		tmp := GetInstanceDBFromInstance[models.TableProperties, TablePropertiesDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.TableRow:
		tmp := GetInstanceDBFromInstance[models.TableRow, TableRowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			case "TableRows":
				if tmp != nil && tmp.Table_TableRowsDBID.Int64 != 0 {
					id := uint(tmp.Table_TableRowsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTable.Map_TableDBID_TablePtr[id]
					res = reservePointerTarget
				}
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.TableStyle:
		tmp := GetInstanceDBFromInstance[models.TableStyle, TableStyleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.Text:
		tmp := GetInstanceDBFromInstance[models.Text, TextDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			}
		case "Document":
			switch reverseField.Fieldname {
			}
		case "Docx":
			switch reverseField.Fieldname {
			}
		case "File":
			switch reverseField.Fieldname {
			}
		case "Node":
			switch reverseField.Fieldname {
			}
		case "Paragraph":
			switch reverseField.Fieldname {
			}
		case "ParagraphProperties":
			switch reverseField.Fieldname {
			}
		case "ParagraphStyle":
			switch reverseField.Fieldname {
			}
		case "Rune":
			switch reverseField.Fieldname {
			}
		case "RuneProperties":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			}
		case "TableProperties":
			switch reverseField.Fieldname {
			}
		case "TableRow":
			switch reverseField.Fieldname {
			}
		case "TableStyle":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}
	
	default:
		_ = inst
	}
	return res
}
