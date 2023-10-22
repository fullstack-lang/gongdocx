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
				if _docx, ok := stage.Docx_Files_reverseMap[inst]; ok {
					res = _docx.Name
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
				if _node, ok := stage.Node_Nodes_reverseMap[inst]; ok {
					res = _node.Name
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
				if _body, ok := stage.Body_Paragraphs_reverseMap[inst]; ok {
					res = _body.Name
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
				if _tablecolumn, ok := stage.TableColumn_Paragraphs_reverseMap[inst]; ok {
					res = _tablecolumn.Name
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
				if _paragraph, ok := stage.Paragraph_Runes_reverseMap[inst]; ok {
					res = _paragraph.Name
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
				if _body, ok := stage.Body_Tables_reverseMap[inst]; ok {
					res = _body.Name
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
				if _tablerow, ok := stage.TableRow_TableColumns_reverseMap[inst]; ok {
					res = _tablerow.Name
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
				if _table, ok := stage.Table_TableRows_reverseMap[inst]; ok {
					res = _table.Name
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
				res = stage.Docx_Files_reverseMap[inst]
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
				res = stage.Node_Nodes_reverseMap[inst]
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
				res = stage.Body_Paragraphs_reverseMap[inst]
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
				res = stage.TableColumn_Paragraphs_reverseMap[inst]
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
				res = stage.Paragraph_Runes_reverseMap[inst]
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
				res = stage.Body_Tables_reverseMap[inst]
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
				res = stage.TableRow_TableColumns_reverseMap[inst]
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
				res = stage.Table_TableRows_reverseMap[inst]
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
