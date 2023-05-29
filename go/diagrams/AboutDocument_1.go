package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdocx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_AboutDocument_1 models.StageStruct
var ___dummy__Time_AboutDocument_1 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_AboutDocument_1 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_AboutDocument_1 map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Document": &(ref_models.Document{}),

	"ref_models.Document.File": (ref_models.Document{}).File,

	"ref_models.Document.Name": (ref_models.Document{}).Name,

	"ref_models.Document.Root": (ref_models.Document{}).Root,

	"ref_models.Docx": &(ref_models.Docx{}),

	"ref_models.Docx.Files": (ref_models.Docx{}).Files,

	"ref_models.Docx.Name": (ref_models.Docx{}).Name,

	"ref_models.File": &(ref_models.File{}),

	"ref_models.File.Name": (ref_models.File{}).Name,

	"ref_models.Node": &(ref_models.Node{}),

	"ref_models.Node.Name": (ref_models.Node{}).Name,

	"ref_models.Node.Nodes": (ref_models.Node{}).Nodes,

	"ref_models.NodeOnText": ref_models.NodeOnText,

	"ref_models.NoteOnDocument": ref_models.NoteOnDocument,

	"ref_models.NoteOnParagraph": ref_models.NoteOnParagraph,

	"ref_models.NoteOnParagraphProperties": ref_models.NoteOnParagraphProperties,

	"ref_models.NoteOnParagraphStyle": ref_models.NoteOnParagraphStyle,

	"ref_models.NoteOnRunProperties": ref_models.NoteOnRunProperties,

	"ref_models.NoteOnRune": ref_models.NoteOnRune,

	"ref_models.NoteOnTable": ref_models.NoteOnTable,

	"ref_models.Paragraph": &(ref_models.Paragraph{}),

	"ref_models.Paragraph.Content": (ref_models.Paragraph{}).Content,

	"ref_models.Paragraph.Name": (ref_models.Paragraph{}).Name,

	"ref_models.Paragraph.Node": (ref_models.Paragraph{}).Node,

	"ref_models.Paragraph.ParagraphProperties": (ref_models.Paragraph{}).ParagraphProperties,

	"ref_models.Paragraph.Runes": (ref_models.Paragraph{}).Runes,

	"ref_models.ParagraphProperties": &(ref_models.ParagraphProperties{}),

	"ref_models.ParagraphProperties.Content": (ref_models.ParagraphProperties{}).Content,

	"ref_models.ParagraphProperties.Name": (ref_models.ParagraphProperties{}).Name,

	"ref_models.ParagraphProperties.Node": (ref_models.ParagraphProperties{}).Node,

	"ref_models.ParagraphProperties.ParagraphStyle": (ref_models.ParagraphProperties{}).ParagraphStyle,

	"ref_models.ParagraphStyle": &(ref_models.ParagraphStyle{}),

	"ref_models.ParagraphStyle.Content": (ref_models.ParagraphStyle{}).Content,

	"ref_models.ParagraphStyle.Name": (ref_models.ParagraphStyle{}).Name,

	"ref_models.ParagraphStyle.Node": (ref_models.ParagraphStyle{}).Node,

	"ref_models.ParagraphStyle.ValAttr": (ref_models.ParagraphStyle{}).ValAttr,

	"ref_models.Rune": &(ref_models.Rune{}),

	"ref_models.Rune.Content": (ref_models.Rune{}).Content,

	"ref_models.Rune.Name": (ref_models.Rune{}).Name,

	"ref_models.Rune.Node": (ref_models.Rune{}).Node,

	"ref_models.Rune.RuneProperties": (ref_models.Rune{}).RuneProperties,

	"ref_models.Rune.Text": (ref_models.Rune{}).Text,

	"ref_models.RuneProperties": &(ref_models.RuneProperties{}),

	"ref_models.RuneProperties.Content": (ref_models.RuneProperties{}).Content,

	"ref_models.RuneProperties.IsBold": (ref_models.RuneProperties{}).IsBold,

	"ref_models.RuneProperties.IsItalic": (ref_models.RuneProperties{}).IsItalic,

	"ref_models.RuneProperties.IsStrike": (ref_models.RuneProperties{}).IsStrike,

	"ref_models.RuneProperties.Name": (ref_models.RuneProperties{}).Name,

	"ref_models.RuneProperties.Node": (ref_models.RuneProperties{}).Node,

	"ref_models.Table": &(ref_models.Table{}),

	"ref_models.Table.Content": (ref_models.Table{}).Content,

	"ref_models.Table.Name": (ref_models.Table{}).Name,

	"ref_models.Table.Node": (ref_models.Table{}).Node,

	"ref_models.Table.TableProperties": (ref_models.Table{}).TableProperties,

	"ref_models.TableProperties": &(ref_models.TableProperties{}),

	"ref_models.TableProperties.Content": (ref_models.TableProperties{}).Content,

	"ref_models.TableProperties.Name": (ref_models.TableProperties{}).Name,

	"ref_models.TableProperties.Node": (ref_models.TableProperties{}).Node,

	"ref_models.TableProperties.TableStyle": (ref_models.TableProperties{}).TableStyle,

	"ref_models.TableStyle": &(ref_models.TableStyle{}),

	"ref_models.TableStyle.Content": (ref_models.TableStyle{}).Content,

	"ref_models.TableStyle.Name": (ref_models.TableStyle{}).Name,

	"ref_models.TableStyle.Node": (ref_models.TableStyle{}).Node,

	"ref_models.TableStyle.Val": (ref_models.TableStyle{}).Val,

	"ref_models.Text": &(ref_models.Text{}),

	"ref_models.Text.Content": (ref_models.Text{}).Content,

	"ref_models.Text.Name": (ref_models.Text{}).Name,

	"ref_models.Text.Node": (ref_models.Text{}).Node,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["AboutTable"] = AboutDocument_1Injection
// }

// AboutDocument_1Injection will stage objects of database "AboutTable"
func AboutDocument_1Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_AboutDocument_1 := (&models.Classdiagram{Name: `AboutTable`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Val := (&models.Field{Name: `Val`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_AboutDocument_1_Table := (&models.GongStructShape{Name: `AboutTable-Table`}).Stage(stage)
	__GongStructShape__000001_AboutDocument_1_TableProperties := (&models.GongStructShape{Name: `AboutTable-TableProperties`}).Stage(stage)
	__GongStructShape__000002_AboutDocument_1_TableStyle := (&models.GongStructShape{Name: `AboutTable-TableStyle`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_TableProperties := (&models.Link{Name: `TableProperties`}).Stage(stage)
	__Link__000001_TableStyle := (&models.Link{Name: `TableStyle`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_NoteOnTable := (&models.NoteShape{Name: `NoteOnTable`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_AboutDocument_1_Table := (&models.Position{Name: `Pos-AboutTable-Table`}).Stage(stage)
	__Position__000001_Pos_AboutDocument_1_TableProperties := (&models.Position{Name: `Pos-AboutTable-TableProperties`}).Stage(stage)
	__Position__000002_Pos_AboutDocument_1_TableStyle := (&models.Position{Name: `Pos-AboutTable-TableStyle`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_Table_and_AboutDocument_1_TableProperties := (&models.Vertice{Name: `Verticle in class diagram AboutTable in middle between AboutTable-Table and AboutTable-TableProperties`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_TableProperties_and_AboutDocument_1_TableStyle := (&models.Vertice{Name: `Verticle in class diagram AboutTable in middle between AboutTable-TableProperties and AboutTable-TableStyle`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_AboutDocument_1.Name = `AboutTable`
	__Classdiagram__000000_AboutDocument_1.IsInDrawMode = true

	// Field values setup
	__Field__000000_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.Name]
	__Field__000000_Name.Identifier = `ref_models.Table.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Table`
	__Field__000000_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableStyle.Name]
	__Field__000001_Name.Identifier = `ref_models.TableStyle.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `TableStyle`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableProperties.Name]
	__Field__000002_Name.Identifier = `ref_models.TableProperties.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `TableProperties`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Val.Name = `Val`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableStyle.Val]
	__Field__000003_Val.Identifier = `ref_models.TableStyle.Val`
	__Field__000003_Val.FieldTypeAsString = ``
	__Field__000003_Val.Structname = `TableStyle`
	__Field__000003_Val.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_AboutDocument_1_Table.Name = `AboutTable-Table`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table]
	__GongStructShape__000000_AboutDocument_1_Table.Identifier = `ref_models.Table`
	__GongStructShape__000000_AboutDocument_1_Table.ShowNbInstances = false
	__GongStructShape__000000_AboutDocument_1_Table.NbInstances = 0
	__GongStructShape__000000_AboutDocument_1_Table.Width = 240.000000
	__GongStructShape__000000_AboutDocument_1_Table.Heigth = 78.000000
	__GongStructShape__000000_AboutDocument_1_Table.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_AboutDocument_1_TableProperties.Name = `AboutTable-TableProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableProperties]
	__GongStructShape__000001_AboutDocument_1_TableProperties.Identifier = `ref_models.TableProperties`
	__GongStructShape__000001_AboutDocument_1_TableProperties.ShowNbInstances = false
	__GongStructShape__000001_AboutDocument_1_TableProperties.NbInstances = 0
	__GongStructShape__000001_AboutDocument_1_TableProperties.Width = 240.000000
	__GongStructShape__000001_AboutDocument_1_TableProperties.Heigth = 78.000000
	__GongStructShape__000001_AboutDocument_1_TableProperties.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_AboutDocument_1_TableStyle.Name = `AboutTable-TableStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableStyle]
	__GongStructShape__000002_AboutDocument_1_TableStyle.Identifier = `ref_models.TableStyle`
	__GongStructShape__000002_AboutDocument_1_TableStyle.ShowNbInstances = false
	__GongStructShape__000002_AboutDocument_1_TableStyle.NbInstances = 0
	__GongStructShape__000002_AboutDocument_1_TableStyle.Width = 240.000000
	__GongStructShape__000002_AboutDocument_1_TableStyle.Heigth = 93.000000
	__GongStructShape__000002_AboutDocument_1_TableStyle.IsSelected = false

	// Link values setup
	__Link__000000_TableProperties.Name = `TableProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.TableProperties]
	__Link__000000_TableProperties.Identifier = `ref_models.Table.TableProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableProperties]
	__Link__000000_TableProperties.Fieldtypename = `ref_models.TableProperties`
	__Link__000000_TableProperties.FieldOffsetX = 8.000000
	__Link__000000_TableProperties.FieldOffsetY = -19.000000
	__Link__000000_TableProperties.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_TableProperties.TargetMultiplicityOffsetX = -43.000000
	__Link__000000_TableProperties.TargetMultiplicityOffsetY = -18.000000
	__Link__000000_TableProperties.SourceMultiplicity = models.MANY
	__Link__000000_TableProperties.SourceMultiplicityOffsetX = -30.000000
	__Link__000000_TableProperties.SourceMultiplicityOffsetY = 28.000000
	__Link__000000_TableProperties.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_TableProperties.StartRatio = 0.491699
	__Link__000000_TableProperties.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_TableProperties.EndRatio = 0.508366
	__Link__000000_TableProperties.CornerOffsetRatio = 1.230769

	// Link values setup
	__Link__000001_TableStyle.Name = `TableStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableProperties.TableStyle]
	__Link__000001_TableStyle.Identifier = `ref_models.TableProperties.TableStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableStyle]
	__Link__000001_TableStyle.Fieldtypename = `ref_models.TableStyle`
	__Link__000001_TableStyle.FieldOffsetX = 16.000000
	__Link__000001_TableStyle.FieldOffsetY = -17.000000
	__Link__000001_TableStyle.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_TableStyle.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_TableStyle.TargetMultiplicityOffsetY = -17.000000
	__Link__000001_TableStyle.SourceMultiplicity = models.MANY
	__Link__000001_TableStyle.SourceMultiplicityOffsetX = -28.000000
	__Link__000001_TableStyle.SourceMultiplicityOffsetY = 24.000000
	__Link__000001_TableStyle.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_TableStyle.StartRatio = 0.487533
	__Link__000001_TableStyle.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_TableStyle.EndRatio = 0.508366
	__Link__000001_TableStyle.CornerOffsetRatio = 1.192308

	// NoteShape values setup
	__NoteShape__000000_NoteOnTable.Name = `NoteOnTable`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnTable]
	__NoteShape__000000_NoteOnTable.Identifier = `ref_models.NoteOnTable`
	__NoteShape__000000_NoteOnTable.Body = `
The "w:tbl" node represents a table within a Word document's XML structure. It
is found in the document.xml file.
-
This node defines the structure and formatting of a table in the document. It
contains child elements that represent the table's properties ("w:tblPr"),
grid ("w:tblGrid"), and rows ("w:tr").
-
The "w:tblPr" node defines the table's overall properties, such as its width,
alignment, borders, and shading.
-
The "w:tblGrid" node defines the table's grid structure - specifically, the
number and width of the columns.
-
The "w:tr" nodes represent table rows, and each row contains "w:tc" nodes that
represent the individual cells within that row.
-
When parsing a "w:tbl" node, your code should handle the table structure and
formatting information it provides to correctly represent the table in your
data structure or output format.
`
	__NoteShape__000000_NoteOnTable.BodyHTML = ``
	__NoteShape__000000_NoteOnTable.X = 471.000000
	__NoteShape__000000_NoteOnTable.Y = 66.000000
	__NoteShape__000000_NoteOnTable.Width = 608.000000
	__NoteShape__000000_NoteOnTable.Heigth = 393.000000
	__NoteShape__000000_NoteOnTable.Matched = false

	// Position values setup
	__Position__000000_Pos_AboutDocument_1_Table.X = 101.000000
	__Position__000000_Pos_AboutDocument_1_Table.Y = 104.000000
	__Position__000000_Pos_AboutDocument_1_Table.Name = `Pos-AboutTable-Table`

	// Position values setup
	__Position__000001_Pos_AboutDocument_1_TableProperties.X = 99.000000
	__Position__000001_Pos_AboutDocument_1_TableProperties.Y = 296.000000
	__Position__000001_Pos_AboutDocument_1_TableProperties.Name = `Pos-AboutTable-TableProperties`

	// Position values setup
	__Position__000002_Pos_AboutDocument_1_TableStyle.X = 92.000000
	__Position__000002_Pos_AboutDocument_1_TableStyle.Y = 486.000000
	__Position__000002_Pos_AboutDocument_1_TableStyle.Name = `Pos-AboutTable-TableStyle`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_Table_and_AboutDocument_1_TableProperties.X = 459.000000
	__Vertice__000000_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_Table_and_AboutDocument_1_TableProperties.Y = 231.000000
	__Vertice__000000_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_Table_and_AboutDocument_1_TableProperties.Name = `Verticle in class diagram AboutTable in middle between AboutTable-Table and AboutTable-TableProperties`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_TableProperties_and_AboutDocument_1_TableStyle.X = 418.500000
	__Vertice__000001_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_TableProperties_and_AboutDocument_1_TableStyle.Y = 209.000000
	__Vertice__000001_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_TableProperties_and_AboutDocument_1_TableStyle.Name = `Verticle in class diagram AboutTable in middle between AboutTable-TableProperties and AboutTable-TableStyle`

	// Setup of pointers
	__Classdiagram__000000_AboutDocument_1.GongStructShapes = append(__Classdiagram__000000_AboutDocument_1.GongStructShapes, __GongStructShape__000000_AboutDocument_1_Table)
	__Classdiagram__000000_AboutDocument_1.GongStructShapes = append(__Classdiagram__000000_AboutDocument_1.GongStructShapes, __GongStructShape__000001_AboutDocument_1_TableProperties)
	__Classdiagram__000000_AboutDocument_1.GongStructShapes = append(__Classdiagram__000000_AboutDocument_1.GongStructShapes, __GongStructShape__000002_AboutDocument_1_TableStyle)
	__Classdiagram__000000_AboutDocument_1.NoteShapes = append(__Classdiagram__000000_AboutDocument_1.NoteShapes, __NoteShape__000000_NoteOnTable)
	__GongStructShape__000000_AboutDocument_1_Table.Position = __Position__000000_Pos_AboutDocument_1_Table
	__GongStructShape__000000_AboutDocument_1_Table.Fields = append(__GongStructShape__000000_AboutDocument_1_Table.Fields, __Field__000000_Name)
	__GongStructShape__000000_AboutDocument_1_Table.Links = append(__GongStructShape__000000_AboutDocument_1_Table.Links, __Link__000000_TableProperties)
	__GongStructShape__000001_AboutDocument_1_TableProperties.Position = __Position__000001_Pos_AboutDocument_1_TableProperties
	__GongStructShape__000001_AboutDocument_1_TableProperties.Fields = append(__GongStructShape__000001_AboutDocument_1_TableProperties.Fields, __Field__000002_Name)
	__GongStructShape__000001_AboutDocument_1_TableProperties.Links = append(__GongStructShape__000001_AboutDocument_1_TableProperties.Links, __Link__000001_TableStyle)
	__GongStructShape__000002_AboutDocument_1_TableStyle.Position = __Position__000002_Pos_AboutDocument_1_TableStyle
	__GongStructShape__000002_AboutDocument_1_TableStyle.Fields = append(__GongStructShape__000002_AboutDocument_1_TableStyle.Fields, __Field__000001_Name)
	__GongStructShape__000002_AboutDocument_1_TableStyle.Fields = append(__GongStructShape__000002_AboutDocument_1_TableStyle.Fields, __Field__000003_Val)
	__Link__000000_TableProperties.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_Table_and_AboutDocument_1_TableProperties
	__Link__000001_TableStyle.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_AboutDocument_1_in_middle_between_AboutDocument_1_TableProperties_and_AboutDocument_1_TableStyle
}
