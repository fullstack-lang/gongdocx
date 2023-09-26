// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongdocx/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
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
				playground,
			),
		}).Stage(formStage)
		body := new(models.Body)
		FillUpForm(body, formGroup, playground)
	case "Document":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Document Form",
			OnSave: __gong__New__DocumentFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		document := new(models.Document)
		FillUpForm(document, formGroup, playground)
	case "Docx":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Docx Form",
			OnSave: __gong__New__DocxFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		docx := new(models.Docx)
		FillUpForm(docx, formGroup, playground)
	case "File":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " File Form",
			OnSave: __gong__New__FileFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		file := new(models.File)
		FillUpForm(file, formGroup, playground)
	case "Node":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Node Form",
			OnSave: __gong__New__NodeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		node := new(models.Node)
		FillUpForm(node, formGroup, playground)
	case "Paragraph":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Paragraph Form",
			OnSave: __gong__New__ParagraphFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		paragraph := new(models.Paragraph)
		FillUpForm(paragraph, formGroup, playground)
	case "ParagraphProperties":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " ParagraphProperties Form",
			OnSave: __gong__New__ParagraphPropertiesFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		paragraphproperties := new(models.ParagraphProperties)
		FillUpForm(paragraphproperties, formGroup, playground)
	case "ParagraphStyle":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " ParagraphStyle Form",
			OnSave: __gong__New__ParagraphStyleFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		paragraphstyle := new(models.ParagraphStyle)
		FillUpForm(paragraphstyle, formGroup, playground)
	case "Rune":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Rune Form",
			OnSave: __gong__New__RuneFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		rune := new(models.Rune)
		FillUpForm(rune, formGroup, playground)
	case "RuneProperties":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " RuneProperties Form",
			OnSave: __gong__New__RunePropertiesFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		runeproperties := new(models.RuneProperties)
		FillUpForm(runeproperties, formGroup, playground)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Table Form",
			OnSave: __gong__New__TableFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		table := new(models.Table)
		FillUpForm(table, formGroup, playground)
	case "TableColumn":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " TableColumn Form",
			OnSave: __gong__New__TableColumnFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		tablecolumn := new(models.TableColumn)
		FillUpForm(tablecolumn, formGroup, playground)
	case "TableProperties":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " TableProperties Form",
			OnSave: __gong__New__TablePropertiesFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		tableproperties := new(models.TableProperties)
		FillUpForm(tableproperties, formGroup, playground)
	case "TableRow":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " TableRow Form",
			OnSave: __gong__New__TableRowFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		tablerow := new(models.TableRow)
		FillUpForm(tablerow, formGroup, playground)
	case "TableStyle":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " TableStyle Form",
			OnSave: __gong__New__TableStyleFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		tablestyle := new(models.TableStyle)
		FillUpForm(tablestyle, formGroup, playground)
	case "Text":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Text Form",
			OnSave: __gong__New__TextFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		text := new(models.Text)
		FillUpForm(text, formGroup, playground)
	}
	formStage.Commit()
}
