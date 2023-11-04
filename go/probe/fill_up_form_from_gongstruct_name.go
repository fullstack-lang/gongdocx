// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdocx/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Body":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Body Form",
			OnSave: __gong__New__BodyFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		body := new(models.Body)
		FillUpForm(body, formGroup, probe)
	case "Document":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Document Form",
			OnSave: __gong__New__DocumentFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		document := new(models.Document)
		FillUpForm(document, formGroup, probe)
	case "Docx":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Docx Form",
			OnSave: __gong__New__DocxFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		docx := new(models.Docx)
		FillUpForm(docx, formGroup, probe)
	case "File":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " File Form",
			OnSave: __gong__New__FileFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		file := new(models.File)
		FillUpForm(file, formGroup, probe)
	case "Node":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Node Form",
			OnSave: __gong__New__NodeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		node := new(models.Node)
		FillUpForm(node, formGroup, probe)
	case "Paragraph":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Paragraph Form",
			OnSave: __gong__New__ParagraphFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		paragraph := new(models.Paragraph)
		FillUpForm(paragraph, formGroup, probe)
	case "ParagraphProperties":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " ParagraphProperties Form",
			OnSave: __gong__New__ParagraphPropertiesFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		paragraphproperties := new(models.ParagraphProperties)
		FillUpForm(paragraphproperties, formGroup, probe)
	case "ParagraphStyle":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " ParagraphStyle Form",
			OnSave: __gong__New__ParagraphStyleFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		paragraphstyle := new(models.ParagraphStyle)
		FillUpForm(paragraphstyle, formGroup, probe)
	case "Rune":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Rune Form",
			OnSave: __gong__New__RuneFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		rune := new(models.Rune)
		FillUpForm(rune, formGroup, probe)
	case "RuneProperties":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " RuneProperties Form",
			OnSave: __gong__New__RunePropertiesFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		runeproperties := new(models.RuneProperties)
		FillUpForm(runeproperties, formGroup, probe)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Table Form",
			OnSave: __gong__New__TableFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		table := new(models.Table)
		FillUpForm(table, formGroup, probe)
	case "TableColumn":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " TableColumn Form",
			OnSave: __gong__New__TableColumnFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		tablecolumn := new(models.TableColumn)
		FillUpForm(tablecolumn, formGroup, probe)
	case "TableProperties":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " TableProperties Form",
			OnSave: __gong__New__TablePropertiesFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		tableproperties := new(models.TableProperties)
		FillUpForm(tableproperties, formGroup, probe)
	case "TableRow":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " TableRow Form",
			OnSave: __gong__New__TableRowFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		tablerow := new(models.TableRow)
		FillUpForm(tablerow, formGroup, probe)
	case "TableStyle":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " TableStyle Form",
			OnSave: __gong__New__TableStyleFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		tablestyle := new(models.TableStyle)
		FillUpForm(tablestyle, formGroup, probe)
	case "Text":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Text Form",
			OnSave: __gong__New__TextFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		text := new(models.Text)
		FillUpForm(text, formGroup, probe)
	}
	formStage.Commit()
}
