// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongdocx/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Body:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Document:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Docx:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.File:
		switch reverseField.GongstructName {
		// insertion point
		case "Docx":
			switch reverseField.Fieldname {
			case "Files":
				if _docx, ok := stage.Docx_Files_reverseMap[inst]; ok {
					res = _docx.Name
				}
			}
		}

	case *models.Node:
		switch reverseField.GongstructName {
		// insertion point
		case "Node":
			switch reverseField.Fieldname {
			case "Nodes":
				if _node, ok := stage.Node_Nodes_reverseMap[inst]; ok {
					res = _node.Name
				}
			}
		}

	case *models.Paragraph:
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			case "Paragraphs":
				if _body, ok := stage.Body_Paragraphs_reverseMap[inst]; ok {
					res = _body.Name
				}
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			case "Paragraphs":
				if _tablecolumn, ok := stage.TableColumn_Paragraphs_reverseMap[inst]; ok {
					res = _tablecolumn.Name
				}
			}
		}

	case *models.ParagraphProperties:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ParagraphStyle:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Rune:
		switch reverseField.GongstructName {
		// insertion point
		case "Paragraph":
			switch reverseField.Fieldname {
			case "Runes":
				if _paragraph, ok := stage.Paragraph_Runes_reverseMap[inst]; ok {
					res = _paragraph.Name
				}
			}
		}

	case *models.RuneProperties:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Table:
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			case "Tables":
				if _body, ok := stage.Body_Tables_reverseMap[inst]; ok {
					res = _body.Name
				}
			}
		}

	case *models.TableColumn:
		switch reverseField.GongstructName {
		// insertion point
		case "TableRow":
			switch reverseField.Fieldname {
			case "TableColumns":
				if _tablerow, ok := stage.TableRow_TableColumns_reverseMap[inst]; ok {
					res = _tablerow.Name
				}
			}
		}

	case *models.TableProperties:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TableRow:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "TableRows":
				if _table, ok := stage.Table_TableRows_reverseMap[inst]; ok {
					res = _table.Name
				}
			}
		}

	case *models.TableStyle:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Text:
		switch reverseField.GongstructName {
		// insertion point
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
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Document:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Docx:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.File:
		switch reverseField.GongstructName {
		// insertion point
		case "Docx":
			switch reverseField.Fieldname {
			case "Files":
				res = stage.Docx_Files_reverseMap[inst]
			}
		}

	case *models.Node:
		switch reverseField.GongstructName {
		// insertion point
		case "Node":
			switch reverseField.Fieldname {
			case "Nodes":
				res = stage.Node_Nodes_reverseMap[inst]
			}
		}

	case *models.Paragraph:
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			case "Paragraphs":
				res = stage.Body_Paragraphs_reverseMap[inst]
			}
		case "TableColumn":
			switch reverseField.Fieldname {
			case "Paragraphs":
				res = stage.TableColumn_Paragraphs_reverseMap[inst]
			}
		}

	case *models.ParagraphProperties:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.ParagraphStyle:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Rune:
		switch reverseField.GongstructName {
		// insertion point
		case "Paragraph":
			switch reverseField.Fieldname {
			case "Runes":
				res = stage.Paragraph_Runes_reverseMap[inst]
			}
		}

	case *models.RuneProperties:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Table:
		switch reverseField.GongstructName {
		// insertion point
		case "Body":
			switch reverseField.Fieldname {
			case "Tables":
				res = stage.Body_Tables_reverseMap[inst]
			}
		}

	case *models.TableColumn:
		switch reverseField.GongstructName {
		// insertion point
		case "TableRow":
			switch reverseField.Fieldname {
			case "TableColumns":
				res = stage.TableRow_TableColumns_reverseMap[inst]
			}
		}

	case *models.TableProperties:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.TableRow:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "TableRows":
				res = stage.Table_TableRows_reverseMap[inst]
			}
		}

	case *models.TableStyle:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Text:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
