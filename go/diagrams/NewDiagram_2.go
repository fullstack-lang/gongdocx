package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdocx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram_2 models.StageStruct
var ___dummy__Time_NewDiagram_2 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram_2 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram_2 map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.BOLD": ref_models.BOLD,

	"ref_models.Document": &(ref_models.Document{}),

	"ref_models.Document.File": (ref_models.Document{}).File,

	"ref_models.Document.Name": (ref_models.Document{}).Name,

	"ref_models.Document.Root": (ref_models.Document{}).Root,

	"ref_models.Docx": &(ref_models.Docx{}),

	"ref_models.Docx.Files": (ref_models.Docx{}).Files,

	"ref_models.Docx.Name": (ref_models.Docx{}).Name,

	"ref_models.File": &(ref_models.File{}),

	"ref_models.File.Name": (ref_models.File{}).Name,

	"ref_models.ITALIC": ref_models.ITALIC,

	"ref_models.Node": &(ref_models.Node{}),

	"ref_models.Node.Name": (ref_models.Node{}).Name,

	"ref_models.Node.Nodes": (ref_models.Node{}).Nodes,

	"ref_models.NodeOnText": ref_models.NodeOnText,

	"ref_models.NoteOnDocument": ref_models.NoteOnDocument,

	"ref_models.NoteOnParagraph": ref_models.NoteOnParagraph,

	"ref_models.NoteOnParagraphProperties": ref_models.NoteOnParagraphProperties,

	"ref_models.NoteOnRunProperties": ref_models.NoteOnRunProperties,

	"ref_models.NoteOnRune": ref_models.NoteOnRune,

	"ref_models.NoteOnStyle": ref_models.NoteOnStyle,

	"ref_models.Paragraph": &(ref_models.Paragraph{}),

	"ref_models.Paragraph.Content": (ref_models.Paragraph{}).Content,

	"ref_models.Paragraph.Name": (ref_models.Paragraph{}).Name,

	"ref_models.Paragraph.Node": (ref_models.Paragraph{}).Node,

	"ref_models.ParagraphProperties": &(ref_models.ParagraphProperties{}),

	"ref_models.ParagraphProperties.Content": (ref_models.ParagraphProperties{}).Content,

	"ref_models.ParagraphProperties.Name": (ref_models.ParagraphProperties{}).Name,

	"ref_models.ParagraphProperties.Node": (ref_models.ParagraphProperties{}).Node,

	"ref_models.ParagraphStyle": &(ref_models.ParagraphStyle{}),

	"ref_models.ParagraphStyle.Content": (ref_models.ParagraphStyle{}).Content,

	"ref_models.ParagraphStyle.Name": (ref_models.ParagraphStyle{}).Name,

	"ref_models.ParagraphStyle.Node": (ref_models.ParagraphStyle{}).Node,

	"ref_models.ParagraphStyle.ValAttr": (ref_models.ParagraphStyle{}).ValAttr,

	"ref_models.Rune": &(ref_models.Rune{}),

	"ref_models.Rune.Content": (ref_models.Rune{}).Content,

	"ref_models.Rune.Name": (ref_models.Rune{}).Name,

	"ref_models.Rune.Node": (ref_models.Rune{}).Node,

	"ref_models.Rune.RuneStyle": (ref_models.Rune{}).RuneStyle,

	"ref_models.RuneProperties": &(ref_models.RuneProperties{}),

	"ref_models.RuneProperties.Content": (ref_models.RuneProperties{}).Content,

	"ref_models.RuneProperties.IsBold": (ref_models.RuneProperties{}).IsBold,

	"ref_models.RuneProperties.IsItalic": (ref_models.RuneProperties{}).IsItalic,

	"ref_models.RuneProperties.IsStrike": (ref_models.RuneProperties{}).IsStrike,

	"ref_models.RuneProperties.Name": (ref_models.RuneProperties{}).Name,

	"ref_models.RuneProperties.Node": (ref_models.RuneProperties{}).Node,

	"ref_models.RuneProperties.RuneStyle": (ref_models.RuneProperties{}).RuneStyle,

	"ref_models.RuneStyle": ref_models.RuneStyle(""),

	"ref_models.STRIKE": ref_models.STRIKE,

	"ref_models.Text": &(ref_models.Text{}),

	"ref_models.Text.Content": (ref_models.Text{}).Content,

	"ref_models.Text.Name": (ref_models.Text{}).Name,

	"ref_models.Text.Node": (ref_models.Text{}).Node,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram_2"] = NewDiagram_2Injection
// }

// NewDiagram_2Injection will stage objects of database "NewDiagram_2"
func NewDiagram_2Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram_2 := (&models.Classdiagram{Name: `NewDiagram_2`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Content := (&models.Field{Name: `Content`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000004_ValAttr := (&models.Field{Name: `ValAttr`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_2_Paragraph := (&models.GongStructShape{Name: `NewDiagram_2-Paragraph`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties := (&models.GongStructShape{Name: `NewDiagram_2-ParagraphProperties`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle := (&models.GongStructShape{Name: `NewDiagram_2-ParagraphStyle`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_NoteOnParagraphProperties := (&models.NoteShape{Name: `NoteOnParagraphProperties`}).Stage(stage)
	__NoteShape__000001_NoteOnStyle := (&models.NoteShape{Name: `NoteOnStyle`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_2_Paragraph := (&models.Position{Name: `Pos-NewDiagram_2-Paragraph`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_2_ParagraphProperties := (&models.Position{Name: `Pos-NewDiagram_2-ParagraphProperties`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_2_ParagraphStyle := (&models.Position{Name: `Pos-NewDiagram_2-ParagraphStyle`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram_2.Name = `NewDiagram_2`
	__Classdiagram__000000_NewDiagram_2.IsInDrawMode = true

	// Field values setup
	__Field__000000_Content.Name = `Content`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphStyle.Content]
	__Field__000000_Content.Identifier = `ref_models.ParagraphStyle.Content`
	__Field__000000_Content.FieldTypeAsString = ``
	__Field__000000_Content.Structname = `ParagraphStyle`
	__Field__000000_Content.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphProperties.Name]
	__Field__000001_Name.Identifier = `ref_models.ParagraphProperties.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `ParagraphProperties`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph.Name]
	__Field__000002_Name.Identifier = `ref_models.Paragraph.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Paragraph`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphStyle.Name]
	__Field__000003_Name.Identifier = `ref_models.ParagraphStyle.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `ParagraphStyle`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_ValAttr.Name = `ValAttr`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphStyle.ValAttr]
	__Field__000004_ValAttr.Identifier = `ref_models.ParagraphStyle.ValAttr`
	__Field__000004_ValAttr.FieldTypeAsString = ``
	__Field__000004_ValAttr.Structname = `ParagraphStyle`
	__Field__000004_ValAttr.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_2_Paragraph.Name = `NewDiagram_2-Paragraph`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph]
	__GongStructShape__000000_NewDiagram_2_Paragraph.Identifier = `ref_models.Paragraph`
	__GongStructShape__000000_NewDiagram_2_Paragraph.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_2_Paragraph.NbInstances = 9
	__GongStructShape__000000_NewDiagram_2_Paragraph.Width = 240.000000
	__GongStructShape__000000_NewDiagram_2_Paragraph.Heigth = 78.000000
	__GongStructShape__000000_NewDiagram_2_Paragraph.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Name = `NewDiagram_2-ParagraphProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphProperties]
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Identifier = `ref_models.ParagraphProperties`
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.NbInstances = 9
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Width = 240.000000
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Heigth = 78.000000
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Name = `NewDiagram_2-ParagraphStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphStyle]
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Identifier = `ref_models.ParagraphStyle`
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.NbInstances = 0
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Width = 240.000000
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Heigth = 108.000000
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.IsSelected = false

	// NoteShape values setup
	__NoteShape__000000_NoteOnParagraphProperties.Name = `NoteOnParagraphProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnParagraphProperties]
	__NoteShape__000000_NoteOnParagraphProperties.Identifier = `ref_models.NoteOnParagraphProperties`
	__NoteShape__000000_NoteOnParagraphProperties.Body = `
The "w:pPr" node represents paragraph properties within a Word document's XML
structure, and is found within a "w:p" (paragraph) node in the document.xml
file.
-
It contains information about the formatting and layout of a paragraph. This can
include properties like text alignment (left, right, centered, or justified),
indentation, spacing before or after the paragraph, line spacing, and more.
-
The "w:pPr" node may also contain a "w:numPr" element for numbered or bulleted
lists, and a "w:sectPr" element for section properties (if this paragraph marks
the end of a section).
-
When parsing a "w:pPr" node, your code should use the information it provides to
format the paragraph appropriately in your data structure or output format.
`
	__NoteShape__000000_NoteOnParagraphProperties.BodyHTML = `<p>The &quot;w:pPr&quot; node represents paragraph properties within a Word document&apos;s XML
structure, and is found within a &quot;w:p&quot; (paragraph) node in the document.xml
file.
-
It contains information about the formatting and layout of a paragraph. This can
include properties like text alignment (left, right, centered, or justified),
indentation, spacing before or after the paragraph, line spacing, and more.
-
The &quot;w:pPr&quot; node may also contain a &quot;w:numPr&quot; element for numbered or bulleted
lists, and a &quot;w:sectPr&quot; element for section properties (if this paragraph marks
the end of a section).
-
When parsing a &quot;w:pPr&quot; node, your code should use the information it provides to
format the paragraph appropriately in your data structure or output format.
`
	__NoteShape__000000_NoteOnParagraphProperties.X = 726.000000
	__NoteShape__000000_NoteOnParagraphProperties.Y = 67.000000
	__NoteShape__000000_NoteOnParagraphProperties.Width = 666.000000
	__NoteShape__000000_NoteOnParagraphProperties.Heigth = 302.000000
	__NoteShape__000000_NoteOnParagraphProperties.Matched = false

	// NoteShape values setup
	__NoteShape__000001_NoteOnStyle.Name = `NoteOnStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnStyle]
	__NoteShape__000001_NoteOnStyle.Identifier = `ref_models.NoteOnStyle`
	__NoteShape__000001_NoteOnStyle.Body = `
The "w:pStyle" element, found within the paragraph properties ("w:pPr") node in
a Word document's XML structure, defines the paragraph style for a given
paragraph ("w:p").
-
The "w:pStyle" element includes an attribute "w:val" that references the ID of
the style being applied to the paragraph. This style ID correlates with the
styles defined in the styles.xml part of the .docx package.
-
A style in Word includes a predefined set of formatting instructions. It can
control multiple aspects of the paragraph's appearance, including alignment,
spacing, font, size, color, and more.
-
By using styles, a document can maintain a consistent look and feel, and
changing the style in one place will automatically update all paragraphs that
reference that style.
-
When parsing a "w:pStyle" node, your code should map the style ID to the
corresponding style in the styles.xml file and apply the associated formatting
to the paragraph.
`
	__NoteShape__000001_NoteOnStyle.BodyHTML = ``
	__NoteShape__000001_NoteOnStyle.X = 585.000000
	__NoteShape__000001_NoteOnStyle.Y = 409.000000
	__NoteShape__000001_NoteOnStyle.Width = 668.000000
	__NoteShape__000001_NoteOnStyle.Heigth = 368.000000
	__NoteShape__000001_NoteOnStyle.Matched = false

	// Position values setup
	__Position__000000_Pos_NewDiagram_2_Paragraph.X = 64.000000
	__Position__000000_Pos_NewDiagram_2_Paragraph.Y = 71.000000
	__Position__000000_Pos_NewDiagram_2_Paragraph.Name = `Pos-NewDiagram_2-Paragraph`

	// Position values setup
	__Position__000001_Pos_NewDiagram_2_ParagraphProperties.X = 434.000000
	__Position__000001_Pos_NewDiagram_2_ParagraphProperties.Y = 69.000000
	__Position__000001_Pos_NewDiagram_2_ParagraphProperties.Name = `Pos-NewDiagram_2-ParagraphProperties`

	// Position values setup
	__Position__000002_Pos_NewDiagram_2_ParagraphStyle.X = 62.000000
	__Position__000002_Pos_NewDiagram_2_ParagraphStyle.Y = 477.000000
	__Position__000002_Pos_NewDiagram_2_ParagraphStyle.Name = `Pos-NewDiagram_2-ParagraphStyle`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram_2.GongStructShapes = append(__Classdiagram__000000_NewDiagram_2.GongStructShapes, __GongStructShape__000000_NewDiagram_2_Paragraph)
	__Classdiagram__000000_NewDiagram_2.GongStructShapes = append(__Classdiagram__000000_NewDiagram_2.GongStructShapes, __GongStructShape__000001_NewDiagram_2_ParagraphProperties)
	__Classdiagram__000000_NewDiagram_2.GongStructShapes = append(__Classdiagram__000000_NewDiagram_2.GongStructShapes, __GongStructShape__000002_NewDiagram_2_ParagraphStyle)
	__Classdiagram__000000_NewDiagram_2.NoteShapes = append(__Classdiagram__000000_NewDiagram_2.NoteShapes, __NoteShape__000000_NoteOnParagraphProperties)
	__Classdiagram__000000_NewDiagram_2.NoteShapes = append(__Classdiagram__000000_NewDiagram_2.NoteShapes, __NoteShape__000001_NoteOnStyle)
	__GongStructShape__000000_NewDiagram_2_Paragraph.Position = __Position__000000_Pos_NewDiagram_2_Paragraph
	__GongStructShape__000000_NewDiagram_2_Paragraph.Fields = append(__GongStructShape__000000_NewDiagram_2_Paragraph.Fields, __Field__000002_Name)
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Position = __Position__000001_Pos_NewDiagram_2_ParagraphProperties
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Fields = append(__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Fields, __Field__000001_Name)
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Position = __Position__000002_Pos_NewDiagram_2_ParagraphStyle
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Fields = append(__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Fields, __Field__000003_Name)
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Fields = append(__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Fields, __Field__000000_Content)
	__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Fields = append(__GongStructShape__000002_NewDiagram_2_ParagraphStyle.Fields, __Field__000004_ValAttr)
}


