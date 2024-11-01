// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Body_Identifiers := make(map[*Body]string)
	_ = map_Body_Identifiers

	bodyOrdered := []*Body{}
	for body := range stage.Bodys {
		bodyOrdered = append(bodyOrdered, body)
	}
	sort.Slice(bodyOrdered[:], func(i, j int) bool {
		return bodyOrdered[i].Name < bodyOrdered[j].Name
	})
	if len(bodyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, body := range bodyOrdered {

		id = generatesIdentifier("Body", idx, body.Name)
		map_Body_Identifiers[body] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Body")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", body.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(body.Name))
		initializerStatements += setValueField

	}

	map_Document_Identifiers := make(map[*Document]string)
	_ = map_Document_Identifiers

	documentOrdered := []*Document{}
	for document := range stage.Documents {
		documentOrdered = append(documentOrdered, document)
	}
	sort.Slice(documentOrdered[:], func(i, j int) bool {
		return documentOrdered[i].Name < documentOrdered[j].Name
	})
	if len(documentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, document := range documentOrdered {

		id = generatesIdentifier("Document", idx, document.Name)
		map_Document_Identifiers[document] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Document")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", document.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(document.Name))
		initializerStatements += setValueField

	}

	map_Docx_Identifiers := make(map[*Docx]string)
	_ = map_Docx_Identifiers

	docxOrdered := []*Docx{}
	for docx := range stage.Docxs {
		docxOrdered = append(docxOrdered, docx)
	}
	sort.Slice(docxOrdered[:], func(i, j int) bool {
		return docxOrdered[i].Name < docxOrdered[j].Name
	})
	if len(docxOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, docx := range docxOrdered {

		id = generatesIdentifier("Docx", idx, docx.Name)
		map_Docx_Identifiers[docx] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Docx")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", docx.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(docx.Name))
		initializerStatements += setValueField

	}

	map_File_Identifiers := make(map[*File]string)
	_ = map_File_Identifiers

	fileOrdered := []*File{}
	for file := range stage.Files {
		fileOrdered = append(fileOrdered, file)
	}
	sort.Slice(fileOrdered[:], func(i, j int) bool {
		return fileOrdered[i].Name < fileOrdered[j].Name
	})
	if len(fileOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, file := range fileOrdered {

		id = generatesIdentifier("File", idx, file.Name)
		map_File_Identifiers[file] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "File")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", file.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(file.Name))
		initializerStatements += setValueField

	}

	map_Node_Identifiers := make(map[*Node]string)
	_ = map_Node_Identifiers

	nodeOrdered := []*Node{}
	for node := range stage.Nodes {
		nodeOrdered = append(nodeOrdered, node)
	}
	sort.Slice(nodeOrdered[:], func(i, j int) bool {
		return nodeOrdered[i].Name < nodeOrdered[j].Name
	})
	if len(nodeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, node := range nodeOrdered {

		id = generatesIdentifier("Node", idx, node.Name)
		map_Node_Identifiers[node] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Node")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", node.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(node.Name))
		initializerStatements += setValueField

	}

	map_Paragraph_Identifiers := make(map[*Paragraph]string)
	_ = map_Paragraph_Identifiers

	paragraphOrdered := []*Paragraph{}
	for paragraph := range stage.Paragraphs {
		paragraphOrdered = append(paragraphOrdered, paragraph)
	}
	sort.Slice(paragraphOrdered[:], func(i, j int) bool {
		return paragraphOrdered[i].Name < paragraphOrdered[j].Name
	})
	if len(paragraphOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, paragraph := range paragraphOrdered {

		id = generatesIdentifier("Paragraph", idx, paragraph.Name)
		map_Paragraph_Identifiers[paragraph] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Paragraph")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", paragraph.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Content))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraph.Text))
		initializerStatements += setValueField

	}

	map_ParagraphProperties_Identifiers := make(map[*ParagraphProperties]string)
	_ = map_ParagraphProperties_Identifiers

	paragraphpropertiesOrdered := []*ParagraphProperties{}
	for paragraphproperties := range stage.ParagraphPropertiess {
		paragraphpropertiesOrdered = append(paragraphpropertiesOrdered, paragraphproperties)
	}
	sort.Slice(paragraphpropertiesOrdered[:], func(i, j int) bool {
		return paragraphpropertiesOrdered[i].Name < paragraphpropertiesOrdered[j].Name
	})
	if len(paragraphpropertiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, paragraphproperties := range paragraphpropertiesOrdered {

		id = generatesIdentifier("ParagraphProperties", idx, paragraphproperties.Name)
		map_ParagraphProperties_Identifiers[paragraphproperties] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParagraphProperties")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", paragraphproperties.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphproperties.Content))
		initializerStatements += setValueField

	}

	map_ParagraphStyle_Identifiers := make(map[*ParagraphStyle]string)
	_ = map_ParagraphStyle_Identifiers

	paragraphstyleOrdered := []*ParagraphStyle{}
	for paragraphstyle := range stage.ParagraphStyles {
		paragraphstyleOrdered = append(paragraphstyleOrdered, paragraphstyle)
	}
	sort.Slice(paragraphstyleOrdered[:], func(i, j int) bool {
		return paragraphstyleOrdered[i].Name < paragraphstyleOrdered[j].Name
	})
	if len(paragraphstyleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, paragraphstyle := range paragraphstyleOrdered {

		id = generatesIdentifier("ParagraphStyle", idx, paragraphstyle.Name)
		map_ParagraphStyle_Identifiers[paragraphstyle] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParagraphStyle")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", paragraphstyle.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.Content))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ValAttr")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(paragraphstyle.ValAttr))
		initializerStatements += setValueField

	}

	map_Rune_Identifiers := make(map[*Rune]string)
	_ = map_Rune_Identifiers

	runeOrdered := []*Rune{}
	for rune := range stage.Runes {
		runeOrdered = append(runeOrdered, rune)
	}
	sort.Slice(runeOrdered[:], func(i, j int) bool {
		return runeOrdered[i].Name < runeOrdered[j].Name
	})
	if len(runeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, rune := range runeOrdered {

		id = generatesIdentifier("Rune", idx, rune.Name)
		map_Rune_Identifiers[rune] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rune")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rune.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rune.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(rune.Content))
		initializerStatements += setValueField

	}

	map_RuneProperties_Identifiers := make(map[*RuneProperties]string)
	_ = map_RuneProperties_Identifiers

	runepropertiesOrdered := []*RuneProperties{}
	for runeproperties := range stage.RunePropertiess {
		runepropertiesOrdered = append(runepropertiesOrdered, runeproperties)
	}
	sort.Slice(runepropertiesOrdered[:], func(i, j int) bool {
		return runepropertiesOrdered[i].Name < runepropertiesOrdered[j].Name
	})
	if len(runepropertiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, runeproperties := range runepropertiesOrdered {

		id = generatesIdentifier("RuneProperties", idx, runeproperties.Name)
		map_RuneProperties_Identifiers[runeproperties] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RuneProperties")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", runeproperties.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(runeproperties.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsBold")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsBold))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsStrike")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsStrike))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsItalic")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", runeproperties.IsItalic))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(runeproperties.Content))
		initializerStatements += setValueField

	}

	map_Table_Identifiers := make(map[*Table]string)
	_ = map_Table_Identifiers

	tableOrdered := []*Table{}
	for table := range stage.Tables {
		tableOrdered = append(tableOrdered, table)
	}
	sort.Slice(tableOrdered[:], func(i, j int) bool {
		return tableOrdered[i].Name < tableOrdered[j].Name
	})
	if len(tableOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, table := range tableOrdered {

		id = generatesIdentifier("Table", idx, table.Name)
		map_Table_Identifiers[table] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Table")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", table.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.Content))
		initializerStatements += setValueField

	}

	map_TableColumn_Identifiers := make(map[*TableColumn]string)
	_ = map_TableColumn_Identifiers

	tablecolumnOrdered := []*TableColumn{}
	for tablecolumn := range stage.TableColumns {
		tablecolumnOrdered = append(tablecolumnOrdered, tablecolumn)
	}
	sort.Slice(tablecolumnOrdered[:], func(i, j int) bool {
		return tablecolumnOrdered[i].Name < tablecolumnOrdered[j].Name
	})
	if len(tablecolumnOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tablecolumn := range tablecolumnOrdered {

		id = generatesIdentifier("TableColumn", idx, tablecolumn.Name)
		map_TableColumn_Identifiers[tablecolumn] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableColumn")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tablecolumn.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablecolumn.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablecolumn.Content))
		initializerStatements += setValueField

	}

	map_TableProperties_Identifiers := make(map[*TableProperties]string)
	_ = map_TableProperties_Identifiers

	tablepropertiesOrdered := []*TableProperties{}
	for tableproperties := range stage.TablePropertiess {
		tablepropertiesOrdered = append(tablepropertiesOrdered, tableproperties)
	}
	sort.Slice(tablepropertiesOrdered[:], func(i, j int) bool {
		return tablepropertiesOrdered[i].Name < tablepropertiesOrdered[j].Name
	})
	if len(tablepropertiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tableproperties := range tablepropertiesOrdered {

		id = generatesIdentifier("TableProperties", idx, tableproperties.Name)
		map_TableProperties_Identifiers[tableproperties] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableProperties")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tableproperties.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tableproperties.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tableproperties.Content))
		initializerStatements += setValueField

	}

	map_TableRow_Identifiers := make(map[*TableRow]string)
	_ = map_TableRow_Identifiers

	tablerowOrdered := []*TableRow{}
	for tablerow := range stage.TableRows {
		tablerowOrdered = append(tablerowOrdered, tablerow)
	}
	sort.Slice(tablerowOrdered[:], func(i, j int) bool {
		return tablerowOrdered[i].Name < tablerowOrdered[j].Name
	})
	if len(tablerowOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tablerow := range tablerowOrdered {

		id = generatesIdentifier("TableRow", idx, tablerow.Name)
		map_TableRow_Identifiers[tablerow] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableRow")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tablerow.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablerow.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablerow.Content))
		initializerStatements += setValueField

	}

	map_TableStyle_Identifiers := make(map[*TableStyle]string)
	_ = map_TableStyle_Identifiers

	tablestyleOrdered := []*TableStyle{}
	for tablestyle := range stage.TableStyles {
		tablestyleOrdered = append(tablestyleOrdered, tablestyle)
	}
	sort.Slice(tablestyleOrdered[:], func(i, j int) bool {
		return tablestyleOrdered[i].Name < tablestyleOrdered[j].Name
	})
	if len(tablestyleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, tablestyle := range tablestyleOrdered {

		id = generatesIdentifier("TableStyle", idx, tablestyle.Name)
		map_TableStyle_Identifiers[tablestyle] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableStyle")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tablestyle.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Content))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Val")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(tablestyle.Val))
		initializerStatements += setValueField

	}

	map_Text_Identifiers := make(map[*Text]string)
	_ = map_Text_Identifiers

	textOrdered := []*Text{}
	for text := range stage.Texts {
		textOrdered = append(textOrdered, text)
	}
	sort.Slice(textOrdered[:], func(i, j int) bool {
		return textOrdered[i].Name < textOrdered[j].Name
	})
	if len(textOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, text := range textOrdered {

		id = generatesIdentifier("Text", idx, text.Name)
		map_Text_Identifiers[text] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Text")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", text.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Content")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(text.Content))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PreserveWhiteSpace")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", text.PreserveWhiteSpace))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, body := range bodyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Body", idx, body.Name)
		map_Body_Identifiers[body] = id

		// Initialisation of values
		for _, _paragraph := range body.Paragraphs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paragraphs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[_paragraph])
			pointersInitializesStatements += setPointerField
		}

		for _, _table := range body.Tables {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Tables")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Table_Identifiers[_table])
			pointersInitializesStatements += setPointerField
		}

		if body.LastParagraph != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LastParagraph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[body.LastParagraph])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, document := range documentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Document", idx, document.Name)
		map_Document_Identifiers[document] = id

		// Initialisation of values
		if document.File != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "File")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_File_Identifiers[document.File])
			pointersInitializesStatements += setPointerField
		}

		if document.Root != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Root")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[document.Root])
			pointersInitializesStatements += setPointerField
		}

		if document.Body != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Body")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Body_Identifiers[document.Body])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, docx := range docxOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Docx", idx, docx.Name)
		map_Docx_Identifiers[docx] = id

		// Initialisation of values
		for _, _file := range docx.Files {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Files")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_File_Identifiers[_file])
			pointersInitializesStatements += setPointerField
		}

		if docx.Document != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Document")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Document_Identifiers[docx.Document])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, file := range fileOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("File", idx, file.Name)
		map_File_Identifiers[file] = id

		// Initialisation of values
	}

	for idx, node := range nodeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Node", idx, node.Name)
		map_Node_Identifiers[node] = id

		// Initialisation of values
		for _, _node := range node.Nodes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Nodes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[_node])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, paragraph := range paragraphOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Paragraph", idx, paragraph.Name)
		map_Paragraph_Identifiers[paragraph] = id

		// Initialisation of values
		if paragraph.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[paragraph.Node])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.ParagraphProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParagraphProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ParagraphProperties_Identifiers[paragraph.ParagraphProperties])
			pointersInitializesStatements += setPointerField
		}

		for _, _rune := range paragraph.Runes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Runes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rune_Identifiers[_rune])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.Next != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Next")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[paragraph.Next])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.Previous != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Previous")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[paragraph.Previous])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.EnclosingBody != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingBody")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Body_Identifiers[paragraph.EnclosingBody])
			pointersInitializesStatements += setPointerField
		}

		if paragraph.EnclosingTableColumn != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingTableColumn")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableColumn_Identifiers[paragraph.EnclosingTableColumn])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, paragraphproperties := range paragraphpropertiesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ParagraphProperties", idx, paragraphproperties.Name)
		map_ParagraphProperties_Identifiers[paragraphproperties] = id

		// Initialisation of values
		if paragraphproperties.ParagraphStyle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ParagraphStyle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ParagraphStyle_Identifiers[paragraphproperties.ParagraphStyle])
			pointersInitializesStatements += setPointerField
		}

		if paragraphproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[paragraphproperties.Node])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, paragraphstyle := range paragraphstyleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ParagraphStyle", idx, paragraphstyle.Name)
		map_ParagraphStyle_Identifiers[paragraphstyle] = id

		// Initialisation of values
		if paragraphstyle.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[paragraphstyle.Node])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, rune := range runeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Rune", idx, rune.Name)
		map_Rune_Identifiers[rune] = id

		// Initialisation of values
		if rune.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[rune.Node])
			pointersInitializesStatements += setPointerField
		}

		if rune.Text != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Text")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Text_Identifiers[rune.Text])
			pointersInitializesStatements += setPointerField
		}

		if rune.RuneProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RuneProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_RuneProperties_Identifiers[rune.RuneProperties])
			pointersInitializesStatements += setPointerField
		}

		if rune.EnclosingParagraph != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingParagraph")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[rune.EnclosingParagraph])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, runeproperties := range runepropertiesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("RuneProperties", idx, runeproperties.Name)
		map_RuneProperties_Identifiers[runeproperties] = id

		// Initialisation of values
		if runeproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[runeproperties.Node])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, table := range tableOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Table", idx, table.Name)
		map_Table_Identifiers[table] = id

		// Initialisation of values
		if table.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[table.Node])
			pointersInitializesStatements += setPointerField
		}

		if table.TableProperties != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableProperties")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableProperties_Identifiers[table.TableProperties])
			pointersInitializesStatements += setPointerField
		}

		for _, _tablerow := range table.TableRows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableRows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableRow_Identifiers[_tablerow])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, tablecolumn := range tablecolumnOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TableColumn", idx, tablecolumn.Name)
		map_TableColumn_Identifiers[tablecolumn] = id

		// Initialisation of values
		if tablecolumn.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[tablecolumn.Node])
			pointersInitializesStatements += setPointerField
		}

		for _, _paragraph := range tablecolumn.Paragraphs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Paragraphs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Paragraph_Identifiers[_paragraph])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, tableproperties := range tablepropertiesOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TableProperties", idx, tableproperties.Name)
		map_TableProperties_Identifiers[tableproperties] = id

		// Initialisation of values
		if tableproperties.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[tableproperties.Node])
			pointersInitializesStatements += setPointerField
		}

		if tableproperties.TableStyle != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableStyle")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableStyle_Identifiers[tableproperties.TableStyle])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, tablerow := range tablerowOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TableRow", idx, tablerow.Name)
		map_TableRow_Identifiers[tablerow] = id

		// Initialisation of values
		if tablerow.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[tablerow.Node])
			pointersInitializesStatements += setPointerField
		}

		for _, _tablecolumn := range tablerow.TableColumns {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TableColumns")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TableColumn_Identifiers[_tablecolumn])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, tablestyle := range tablestyleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TableStyle", idx, tablestyle.Name)
		map_TableStyle_Identifiers[tablestyle] = id

		// Initialisation of values
		if tablestyle.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[tablestyle.Node])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, text := range textOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Text", idx, text.Name)
		map_Text_Identifiers[text] = id

		// Initialisation of values
		if text.Node != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Node")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Node_Identifiers[text.Node])
			pointersInitializesStatements += setPointerField
		}

		if text.EnclosingRune != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingRune")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Rune_Identifiers[text.EnclosingRune])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
