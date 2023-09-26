// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdocx/go/models"
	"github.com/fullstack-lang/gongdocx/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Body:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Paragraphs", instanceWithInferedType, &instanceWithInferedType.Paragraphs, formGroup, playground)
		AssociationSliceToForm("Tables", instanceWithInferedType, &instanceWithInferedType.Tables, formGroup, playground)
		AssociationFieldToForm("LastParagraph", instanceWithInferedType.LastParagraph, formGroup, playground)

	case *models.Document:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("File", instanceWithInferedType.File, formGroup, playground)
		AssociationFieldToForm("Root", instanceWithInferedType.Root, formGroup, playground)
		AssociationFieldToForm("Body", instanceWithInferedType.Body, formGroup, playground)

	case *models.Docx:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Files", instanceWithInferedType, &instanceWithInferedType.Files, formGroup, playground)
		AssociationFieldToForm("Document", instanceWithInferedType.Document, formGroup, playground)

	case *models.File:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Docx"
			rf.Fieldname = "Files"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Docx),
					"Files",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Docx, *models.File](
					nil,
					"Files",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Node:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Nodes", instanceWithInferedType, &instanceWithInferedType.Nodes, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Nodes"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Node),
					"Nodes",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Node, *models.Node](
					nil,
					"Nodes",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Paragraph:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		AssociationFieldToForm("ParagraphProperties", instanceWithInferedType.ParagraphProperties, formGroup, playground)
		AssociationSliceToForm("Runes", instanceWithInferedType, &instanceWithInferedType.Runes, formGroup, playground)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Next", instanceWithInferedType.Next, formGroup, playground)
		AssociationFieldToForm("Previous", instanceWithInferedType.Previous, formGroup, playground)
		AssociationFieldToForm("EnclosingBody", instanceWithInferedType.EnclosingBody, formGroup, playground)
		AssociationFieldToForm("EnclosingTableColumn", instanceWithInferedType.EnclosingTableColumn, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Body"
			rf.Fieldname = "Paragraphs"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Body),
					"Paragraphs",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Body, *models.Paragraph](
					nil,
					"Paragraphs",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TableColumn"
			rf.Fieldname = "Paragraphs"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.TableColumn),
					"Paragraphs",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.TableColumn, *models.Paragraph](
					nil,
					"Paragraphs",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.ParagraphProperties:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("ParagraphStyle", instanceWithInferedType.ParagraphStyle, formGroup, playground)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)

	case *models.ParagraphStyle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("ValAttr", instanceWithInferedType.ValAttr, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Rune:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		AssociationFieldToForm("Text", instanceWithInferedType.Text, formGroup, playground)
		AssociationFieldToForm("RuneProperties", instanceWithInferedType.RuneProperties, formGroup, playground)
		AssociationFieldToForm("EnclosingParagraph", instanceWithInferedType.EnclosingParagraph, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Paragraph"
			rf.Fieldname = "Runes"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Paragraph),
					"Runes",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Paragraph, *models.Rune](
					nil,
					"Runes",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.RuneProperties:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		BasicFieldtoForm("IsBold", instanceWithInferedType.IsBold, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsStrike", instanceWithInferedType.IsStrike, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsItalic", instanceWithInferedType.IsItalic, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Table:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("TableProperties", instanceWithInferedType.TableProperties, formGroup, playground)
		AssociationSliceToForm("TableRows", instanceWithInferedType, &instanceWithInferedType.TableRows, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Body"
			rf.Fieldname = "Tables"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Body),
					"Tables",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Body, *models.Table](
					nil,
					"Tables",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.TableColumn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		AssociationSliceToForm("Paragraphs", instanceWithInferedType, &instanceWithInferedType.Paragraphs, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TableRow"
			rf.Fieldname = "TableColumns"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.TableRow),
					"TableColumns",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.TableRow, *models.TableColumn](
					nil,
					"TableColumns",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.TableProperties:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("TableStyle", instanceWithInferedType.TableStyle, formGroup, playground)

	case *models.TableRow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		AssociationSliceToForm("TableColumns", instanceWithInferedType, &instanceWithInferedType.TableColumns, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "TableRows"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Table),
					"TableRows",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Table, *models.TableRow](
					nil,
					"TableRows",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.TableStyle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Val", instanceWithInferedType.Val, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, playground)
		BasicFieldtoForm("PreserveWhiteSpace", instanceWithInferedType.PreserveWhiteSpace, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("EnclosingRune", instanceWithInferedType.EnclosingRune, formGroup, playground)

	default:
		_ = instanceWithInferedType
	}
}
