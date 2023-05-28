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

	"ref_models.NoteOnRune": ref_models.NoteOnRune,

	"ref_models.Paragraph": &(ref_models.Paragraph{}),

	"ref_models.Paragraph.Name": (ref_models.Paragraph{}).Name,

	"ref_models.Paragraph.Node": (ref_models.Paragraph{}).Node,

	"ref_models.ParagraphProperties": &(ref_models.ParagraphProperties{}),

	"ref_models.ParagraphProperties.Name": (ref_models.ParagraphProperties{}).Name,

	"ref_models.ParagraphProperties.Node": (ref_models.ParagraphProperties{}).Node,

	"ref_models.Rune": &(ref_models.Rune{}),

	"ref_models.Rune.Name": (ref_models.Rune{}).Name,

	"ref_models.Rune.Node": (ref_models.Rune{}).Node,

	"ref_models.Rune.RuneStyle": (ref_models.Rune{}).RuneStyle,

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
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_2_Paragraph := (&models.GongStructShape{Name: `NewDiagram_2-Paragraph`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties := (&models.GongStructShape{Name: `NewDiagram_2-ParagraphProperties`}).Stage(stage)

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_NoteOnParagraphProperties := (&models.NoteShape{Name: `NoteOnParagraphProperties`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_2_Paragraph := (&models.Position{Name: `Pos-NewDiagram_2-Paragraph`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_2_ParagraphProperties := (&models.Position{Name: `Pos-NewDiagram_2-ParagraphProperties`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram_2.Name = `NewDiagram_2`
	__Classdiagram__000000_NewDiagram_2.IsInDrawMode = true

	// Field values setup
	__Field__000000_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphProperties.Name]
	__Field__000000_Name.Identifier = `ref_models.ParagraphProperties.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `ParagraphProperties`
	__Field__000000_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph.Name]
	__Field__000001_Name.Identifier = `ref_models.Paragraph.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Paragraph`
	__Field__000001_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_2_Paragraph.Name = `NewDiagram_2-Paragraph`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph]
	__GongStructShape__000000_NewDiagram_2_Paragraph.Identifier = `ref_models.Paragraph`
	__GongStructShape__000000_NewDiagram_2_Paragraph.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_2_Paragraph.NbInstances = 0
	__GongStructShape__000000_NewDiagram_2_Paragraph.Width = 240.000000
	__GongStructShape__000000_NewDiagram_2_Paragraph.Heigth = 78.000000
	__GongStructShape__000000_NewDiagram_2_Paragraph.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Name = `NewDiagram_2-ParagraphProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphProperties]
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Identifier = `ref_models.ParagraphProperties`
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.NbInstances = 0
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Width = 240.000000
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Heigth = 78.000000
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.IsSelected = false

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
	__NoteShape__000000_NoteOnParagraphProperties.BodyHTML = ``
	__NoteShape__000000_NoteOnParagraphProperties.X = 726.000000
	__NoteShape__000000_NoteOnParagraphProperties.Y = 67.000000
	__NoteShape__000000_NoteOnParagraphProperties.Width = 666.000000
	__NoteShape__000000_NoteOnParagraphProperties.Heigth = 302.000000
	__NoteShape__000000_NoteOnParagraphProperties.Matched = false

	// Position values setup
	__Position__000000_Pos_NewDiagram_2_Paragraph.X = 64.000000
	__Position__000000_Pos_NewDiagram_2_Paragraph.Y = 71.000000
	__Position__000000_Pos_NewDiagram_2_Paragraph.Name = `Pos-NewDiagram_2-Paragraph`

	// Position values setup
	__Position__000001_Pos_NewDiagram_2_ParagraphProperties.X = 434.000000
	__Position__000001_Pos_NewDiagram_2_ParagraphProperties.Y = 69.000000
	__Position__000001_Pos_NewDiagram_2_ParagraphProperties.Name = `Pos-NewDiagram_2-ParagraphProperties`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram_2.GongStructShapes = append(__Classdiagram__000000_NewDiagram_2.GongStructShapes, __GongStructShape__000000_NewDiagram_2_Paragraph)
	__Classdiagram__000000_NewDiagram_2.GongStructShapes = append(__Classdiagram__000000_NewDiagram_2.GongStructShapes, __GongStructShape__000001_NewDiagram_2_ParagraphProperties)
	__Classdiagram__000000_NewDiagram_2.NoteShapes = append(__Classdiagram__000000_NewDiagram_2.NoteShapes, __NoteShape__000000_NoteOnParagraphProperties)
	__GongStructShape__000000_NewDiagram_2_Paragraph.Position = __Position__000000_Pos_NewDiagram_2_Paragraph
	__GongStructShape__000000_NewDiagram_2_Paragraph.Fields = append(__GongStructShape__000000_NewDiagram_2_Paragraph.Fields, __Field__000001_Name)
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Position = __Position__000001_Pos_NewDiagram_2_ParagraphProperties
	__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Fields = append(__GongStructShape__000001_NewDiagram_2_ParagraphProperties.Fields, __Field__000000_Name)
}


