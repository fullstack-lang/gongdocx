package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdocx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_AboutTable models.StageStruct
var ___dummy__Time_AboutTable time.Time

// _ point for meta package dummy declaration
var ___dummy__ref_models_AboutTable ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
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

	"ref_models.NoteOnColumn": ref_models.NoteOnColumn,

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

	"ref_models.Table.TableRows": (ref_models.Table{}).TableRows,

	"ref_models.TableColumn": &(ref_models.TableColumn{}),

	"ref_models.TableColumn.Content": (ref_models.TableColumn{}).Content,

	"ref_models.TableColumn.Name": (ref_models.TableColumn{}).Name,

	"ref_models.TableColumn.Node": (ref_models.TableColumn{}).Node,

	"ref_models.TableColumn.Paragraphs": (ref_models.TableColumn{}).Paragraphs,

	"ref_models.TableProperties": &(ref_models.TableProperties{}),

	"ref_models.TableProperties.Content": (ref_models.TableProperties{}).Content,

	"ref_models.TableProperties.Name": (ref_models.TableProperties{}).Name,

	"ref_models.TableProperties.Node": (ref_models.TableProperties{}).Node,

	"ref_models.TableProperties.TableStyle": (ref_models.TableProperties{}).TableStyle,

	"ref_models.TableRow": &(ref_models.TableRow{}),

	"ref_models.TableRow.Content": (ref_models.TableRow{}).Content,

	"ref_models.TableRow.Name": (ref_models.TableRow{}).Name,

	"ref_models.TableRow.Node": (ref_models.TableRow{}).Node,

	"ref_models.TableRow.TableColumns": (ref_models.TableRow{}).TableColumns,

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
// but it has to properly reference the _ gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["AboutTable"] = _
// }

// _ will stage objects of database "AboutTable"
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_AboutTable := (&models.Classdiagram{Name: `AboutTable`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000004_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000005_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000006_Val := (&models.Field{Name: `Val`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_AboutTable_Paragraph := (&models.GongStructShape{Name: `AboutTable-Paragraph`}).Stage(stage)
	__GongStructShape__000001_AboutTable_Table := (&models.GongStructShape{Name: `AboutTable-Table`}).Stage(stage)
	__GongStructShape__000002_AboutTable_TableColumn := (&models.GongStructShape{Name: `AboutTable-TableColumn`}).Stage(stage)
	__GongStructShape__000003_AboutTable_TableProperties := (&models.GongStructShape{Name: `AboutTable-TableProperties`}).Stage(stage)
	__GongStructShape__000004_AboutTable_TableRow := (&models.GongStructShape{Name: `AboutTable-TableRow`}).Stage(stage)
	__GongStructShape__000005_AboutTable_TableStyle := (&models.GongStructShape{Name: `AboutTable-TableStyle`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Paragraphs := (&models.Link{Name: `Paragraphs`}).Stage(stage)
	__Link__000001_TableColumns := (&models.Link{Name: `TableColumns`}).Stage(stage)
	__Link__000002_TableProperties := (&models.Link{Name: `TableProperties`}).Stage(stage)
	__Link__000003_TableRows := (&models.Link{Name: `TableRows`}).Stage(stage)
	__Link__000004_TableStyle := (&models.Link{Name: `TableStyle`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_NoteOnColumn := (&models.NoteShape{Name: `NoteOnColumn`}).Stage(stage)
	__NoteShape__000001_NoteOnTable := (&models.NoteShape{Name: `NoteOnTable`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_AboutTable_Paragraph := (&models.Position{Name: `Pos-AboutTable-Paragraph`}).Stage(stage)
	__Position__000001_Pos_AboutTable_Table := (&models.Position{Name: `Pos-AboutTable-Table`}).Stage(stage)
	__Position__000002_Pos_AboutTable_TableColumn := (&models.Position{Name: `Pos-AboutTable-TableColumn`}).Stage(stage)
	__Position__000003_Pos_AboutTable_TableProperties := (&models.Position{Name: `Pos-AboutTable-TableProperties`}).Stage(stage)
	__Position__000004_Pos_AboutTable_TableRow := (&models.Position{Name: `Pos-AboutTable-TableRow`}).Stage(stage)
	__Position__000005_Pos_AboutTable_TableStyle := (&models.Position{Name: `Pos-AboutTable-TableStyle`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableProperties := (&models.Vertice{Name: `Verticle in class diagram AboutTable in middle between AboutTable-Table and AboutTable-TableProperties`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableRow := (&models.Vertice{Name: `Verticle in class diagram AboutTable in middle between AboutTable-Table and AboutTable-TableRow`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableColumn_and_AboutTable_Paragraph := (&models.Vertice{Name: `Verticle in class diagram AboutTable in middle between AboutTable-TableColumn and AboutTable-Paragraph`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableProperties_and_AboutTable_TableStyle := (&models.Vertice{Name: `Verticle in class diagram AboutTable in middle between AboutTable-TableProperties and AboutTable-TableStyle`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableRow_and_AboutTable_TableColumn := (&models.Vertice{Name: `Verticle in class diagram AboutTable in middle between AboutTable-TableRow and AboutTable-TableColumn`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_AboutTable.Name = `AboutTable`
	__Classdiagram__000000_AboutTable.IsInDrawMode = true

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

	//gong:ident [ref_models.TableRow.Name]
	__Field__000002_Name.Identifier = `ref_models.TableRow.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `TableRow`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableProperties.Name]
	__Field__000003_Name.Identifier = `ref_models.TableProperties.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `TableProperties`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableColumn.Name]
	__Field__000004_Name.Identifier = `ref_models.TableColumn.Name`
	__Field__000004_Name.FieldTypeAsString = ``
	__Field__000004_Name.Structname = `TableColumn`
	__Field__000004_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph.Name]
	__Field__000005_Name.Identifier = `ref_models.Paragraph.Name`
	__Field__000005_Name.FieldTypeAsString = ``
	__Field__000005_Name.Structname = `Paragraph`
	__Field__000005_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000006_Val.Name = `Val`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableStyle.Val]
	__Field__000006_Val.Identifier = `ref_models.TableStyle.Val`
	__Field__000006_Val.FieldTypeAsString = ``
	__Field__000006_Val.Structname = `TableStyle`
	__Field__000006_Val.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_AboutTable_Paragraph.Name = `AboutTable-Paragraph`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph]
	__GongStructShape__000000_AboutTable_Paragraph.Identifier = `ref_models.Paragraph`
	__GongStructShape__000000_AboutTable_Paragraph.ShowNbInstances = false
	__GongStructShape__000000_AboutTable_Paragraph.NbInstances = 0
	__GongStructShape__000000_AboutTable_Paragraph.Width = 240.000000
	__GongStructShape__000000_AboutTable_Paragraph.Height = 78.000000
	__GongStructShape__000000_AboutTable_Paragraph.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_AboutTable_Table.Name = `AboutTable-Table`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table]
	__GongStructShape__000001_AboutTable_Table.Identifier = `ref_models.Table`
	__GongStructShape__000001_AboutTable_Table.ShowNbInstances = false
	__GongStructShape__000001_AboutTable_Table.NbInstances = 0
	__GongStructShape__000001_AboutTable_Table.Width = 240.000000
	__GongStructShape__000001_AboutTable_Table.Height = 78.000000
	__GongStructShape__000001_AboutTable_Table.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_AboutTable_TableColumn.Name = `AboutTable-TableColumn`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableColumn]
	__GongStructShape__000002_AboutTable_TableColumn.Identifier = `ref_models.TableColumn`
	__GongStructShape__000002_AboutTable_TableColumn.ShowNbInstances = false
	__GongStructShape__000002_AboutTable_TableColumn.NbInstances = 0
	__GongStructShape__000002_AboutTable_TableColumn.Width = 240.000000
	__GongStructShape__000002_AboutTable_TableColumn.Height = 78.000000
	__GongStructShape__000002_AboutTable_TableColumn.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_AboutTable_TableProperties.Name = `AboutTable-TableProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableProperties]
	__GongStructShape__000003_AboutTable_TableProperties.Identifier = `ref_models.TableProperties`
	__GongStructShape__000003_AboutTable_TableProperties.ShowNbInstances = false
	__GongStructShape__000003_AboutTable_TableProperties.NbInstances = 0
	__GongStructShape__000003_AboutTable_TableProperties.Width = 240.000000
	__GongStructShape__000003_AboutTable_TableProperties.Height = 78.000000
	__GongStructShape__000003_AboutTable_TableProperties.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_AboutTable_TableRow.Name = `AboutTable-TableRow`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableRow]
	__GongStructShape__000004_AboutTable_TableRow.Identifier = `ref_models.TableRow`
	__GongStructShape__000004_AboutTable_TableRow.ShowNbInstances = false
	__GongStructShape__000004_AboutTable_TableRow.NbInstances = 0
	__GongStructShape__000004_AboutTable_TableRow.Width = 240.000000
	__GongStructShape__000004_AboutTable_TableRow.Height = 78.000000
	__GongStructShape__000004_AboutTable_TableRow.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_AboutTable_TableStyle.Name = `AboutTable-TableStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableStyle]
	__GongStructShape__000005_AboutTable_TableStyle.Identifier = `ref_models.TableStyle`
	__GongStructShape__000005_AboutTable_TableStyle.ShowNbInstances = false
	__GongStructShape__000005_AboutTable_TableStyle.NbInstances = 0
	__GongStructShape__000005_AboutTable_TableStyle.Width = 240.000000
	__GongStructShape__000005_AboutTable_TableStyle.Height = 93.000000
	__GongStructShape__000005_AboutTable_TableStyle.IsSelected = false

	// Link values setup
	__Link__000000_Paragraphs.Name = `Paragraphs`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableColumn.Paragraphs]
	__Link__000000_Paragraphs.Identifier = `ref_models.TableColumn.Paragraphs`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph]
	__Link__000000_Paragraphs.Fieldtypename = `ref_models.Paragraph`
	__Link__000000_Paragraphs.FieldOffsetX = 10.000000
	__Link__000000_Paragraphs.FieldOffsetY = -19.000000
	__Link__000000_Paragraphs.TargetMultiplicity = models.MANY
	__Link__000000_Paragraphs.TargetMultiplicityOffsetX = -30.000000
	__Link__000000_Paragraphs.TargetMultiplicityOffsetY = -12.000000
	__Link__000000_Paragraphs.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_Paragraphs.SourceMultiplicityOffsetX = -39.000000
	__Link__000000_Paragraphs.SourceMultiplicityOffsetY = 25.000000
	__Link__000000_Paragraphs.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Paragraphs.StartRatio = 0.400000
	__Link__000000_Paragraphs.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Paragraphs.EndRatio = 0.433333
	__Link__000000_Paragraphs.CornerOffsetRatio = 1.205128

	// Link values setup
	__Link__000001_TableColumns.Name = `TableColumns`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableRow.TableColumns]
	__Link__000001_TableColumns.Identifier = `ref_models.TableRow.TableColumns`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableColumn]
	__Link__000001_TableColumns.Fieldtypename = `ref_models.TableColumn`
	__Link__000001_TableColumns.FieldOffsetX = 14.000000
	__Link__000001_TableColumns.FieldOffsetY = -17.000000
	__Link__000001_TableColumns.TargetMultiplicity = models.MANY
	__Link__000001_TableColumns.TargetMultiplicityOffsetX = -30.000000
	__Link__000001_TableColumns.TargetMultiplicityOffsetY = -14.000000
	__Link__000001_TableColumns.SourceMultiplicity = models.ZERO_ONE
	__Link__000001_TableColumns.SourceMultiplicityOffsetX = -35.000000
	__Link__000001_TableColumns.SourceMultiplicityOffsetY = 25.000000
	__Link__000001_TableColumns.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_TableColumns.StartRatio = 0.416667
	__Link__000001_TableColumns.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_TableColumns.EndRatio = 0.408333
	__Link__000001_TableColumns.CornerOffsetRatio = 1.320513

	// Link values setup
	__Link__000002_TableProperties.Name = `TableProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.TableProperties]
	__Link__000002_TableProperties.Identifier = `ref_models.Table.TableProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableProperties]
	__Link__000002_TableProperties.Fieldtypename = `ref_models.TableProperties`
	__Link__000002_TableProperties.FieldOffsetX = 8.000000
	__Link__000002_TableProperties.FieldOffsetY = -19.000000
	__Link__000002_TableProperties.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_TableProperties.TargetMultiplicityOffsetX = -43.000000
	__Link__000002_TableProperties.TargetMultiplicityOffsetY = -18.000000
	__Link__000002_TableProperties.SourceMultiplicity = models.MANY
	__Link__000002_TableProperties.SourceMultiplicityOffsetX = -30.000000
	__Link__000002_TableProperties.SourceMultiplicityOffsetY = 28.000000
	__Link__000002_TableProperties.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_TableProperties.StartRatio = 0.491699
	__Link__000002_TableProperties.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_TableProperties.EndRatio = 0.508366
	__Link__000002_TableProperties.CornerOffsetRatio = 1.230769

	// Link values setup
	__Link__000003_TableRows.Name = `TableRows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Table.TableRows]
	__Link__000003_TableRows.Identifier = `ref_models.Table.TableRows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableRow]
	__Link__000003_TableRows.Fieldtypename = `ref_models.TableRow`
	__Link__000003_TableRows.FieldOffsetX = -83.000000
	__Link__000003_TableRows.FieldOffsetY = -20.000000
	__Link__000003_TableRows.TargetMultiplicity = models.MANY
	__Link__000003_TableRows.TargetMultiplicityOffsetX = -38.000000
	__Link__000003_TableRows.TargetMultiplicityOffsetY = 24.000000
	__Link__000003_TableRows.SourceMultiplicity = models.ZERO_ONE
	__Link__000003_TableRows.SourceMultiplicityOffsetX = 18.000000
	__Link__000003_TableRows.SourceMultiplicityOffsetY = 22.000000
	__Link__000003_TableRows.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_TableRows.StartRatio = 0.500000
	__Link__000003_TableRows.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_TableRows.EndRatio = 0.551282
	__Link__000003_TableRows.CornerOffsetRatio = 1.358333

	// Link values setup
	__Link__000004_TableStyle.Name = `TableStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableProperties.TableStyle]
	__Link__000004_TableStyle.Identifier = `ref_models.TableProperties.TableStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.TableStyle]
	__Link__000004_TableStyle.Fieldtypename = `ref_models.TableStyle`
	__Link__000004_TableStyle.FieldOffsetX = 16.000000
	__Link__000004_TableStyle.FieldOffsetY = -17.000000
	__Link__000004_TableStyle.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_TableStyle.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_TableStyle.TargetMultiplicityOffsetY = -17.000000
	__Link__000004_TableStyle.SourceMultiplicity = models.MANY
	__Link__000004_TableStyle.SourceMultiplicityOffsetX = -28.000000
	__Link__000004_TableStyle.SourceMultiplicityOffsetY = 24.000000
	__Link__000004_TableStyle.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_TableStyle.StartRatio = 0.487533
	__Link__000004_TableStyle.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_TableStyle.EndRatio = 0.508366
	__Link__000004_TableStyle.CornerOffsetRatio = 1.192308

	// NoteShape values setup
	__NoteShape__000000_NoteOnColumn.Name = `NoteOnColumn`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnColumn]
	__NoteShape__000000_NoteOnColumn.Identifier = `ref_models.NoteOnColumn`
	__NoteShape__000000_NoteOnColumn.Body = `
The "w:tr" node represents a table row within a Word document's XML structure.
It is found as a child of the "w:tbl" (table) node in the document.xml file.
-
This node contains child elements that represent the individual cells ("w:tc")
in the row, as well as the row properties ("w:trPr"), which include attributes
like height, header status, and more.
-
The "w:tc" node, on the other hand, represents an individual table cell. It is
found as a child of the "w:tr" (table row) node.
-
This node contains the content of the cell, which can include text,
paragraphs, or even other tables. It also contains the cell properties
("w:tcPr"), which include attributes like cell width, vertical alignment,
borders, shading, and more.
-
When parsing "w:tr" and "w:tc" nodes, your code should correctly map the
structure of the table, row, and cell, and apply the appropriate properties
to each element.
`
	__NoteShape__000000_NoteOnColumn.BodyHTML = ``
	__NoteShape__000000_NoteOnColumn.X = 503.000000
	__NoteShape__000000_NoteOnColumn.Y = 652.000000
	__NoteShape__000000_NoteOnColumn.Width = 730.000000
	__NoteShape__000000_NoteOnColumn.Height = 366.000000
	__NoteShape__000000_NoteOnColumn.Matched = false

	// NoteShape values setup
	__NoteShape__000001_NoteOnTable.Name = `NoteOnTable`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnTable]
	__NoteShape__000001_NoteOnTable.Identifier = `ref_models.NoteOnTable`
	__NoteShape__000001_NoteOnTable.Body = `
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
	__NoteShape__000001_NoteOnTable.BodyHTML = `<p>The &quot;w:tbl&quot; node represents a table within a Word document&apos;s XML structure. It
is found in the document.xml file.
-
This node defines the structure and formatting of a table in the document. It
contains child elements that represent the table&apos;s properties (&quot;w:tblPr&quot;),
grid (&quot;w:tblGrid&quot;), and rows (&quot;w:tr&quot;).
-
The &quot;w:tblPr&quot; node defines the table&apos;s overall properties, such as its width,
alignment, borders, and shading.
-
The &quot;w:tblGrid&quot; node defines the table&apos;s grid structure - specifically, the
number and width of the columns.
-
The &quot;w:tr&quot; nodes represent table rows, and each row contains &quot;w:tc&quot; nodes that
represent the individual cells within that row.
-
When parsing a &quot;w:tbl&quot; node, your code should handle the table structure and
formatting information it provides to correctly represent the table in your
data structure or output format.
`
	__NoteShape__000001_NoteOnTable.X = 798.000000
	__NoteShape__000001_NoteOnTable.Y = 81.000000
	__NoteShape__000001_NoteOnTable.Width = 608.000000
	__NoteShape__000001_NoteOnTable.Height = 393.000000
	__NoteShape__000001_NoteOnTable.Matched = false

	// Position values setup
	__Position__000000_Pos_AboutTable_Paragraph.X = 513.000000
	__Position__000000_Pos_AboutTable_Paragraph.Y = 528.000000
	__Position__000000_Pos_AboutTable_Paragraph.Name = `Pos-AboutTable-Paragraph`

	// Position values setup
	__Position__000001_Pos_AboutTable_Table.X = 101.000000
	__Position__000001_Pos_AboutTable_Table.Y = 104.000000
	__Position__000001_Pos_AboutTable_Table.Name = `Pos-AboutTable-Table`

	// Position values setup
	__Position__000002_Pos_AboutTable_TableColumn.X = 521.000000
	__Position__000002_Pos_AboutTable_TableColumn.Y = 325.000000
	__Position__000002_Pos_AboutTable_TableColumn.Name = `Pos-AboutTable-TableColumn`

	// Position values setup
	__Position__000003_Pos_AboutTable_TableProperties.X = 99.000000
	__Position__000003_Pos_AboutTable_TableProperties.Y = 296.000000
	__Position__000003_Pos_AboutTable_TableProperties.Name = `Pos-AboutTable-TableProperties`

	// Position values setup
	__Position__000004_Pos_AboutTable_TableRow.X = 516.000000
	__Position__000004_Pos_AboutTable_TableRow.Y = 104.000000
	__Position__000004_Pos_AboutTable_TableRow.Name = `Pos-AboutTable-TableRow`

	// Position values setup
	__Position__000005_Pos_AboutTable_TableStyle.X = 92.000000
	__Position__000005_Pos_AboutTable_TableStyle.Y = 486.000000
	__Position__000005_Pos_AboutTable_TableStyle.Name = `Pos-AboutTable-TableStyle`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableProperties.X = 459.000000
	__Vertice__000000_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableProperties.Y = 231.000000
	__Vertice__000000_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableProperties.Name = `Verticle in class diagram AboutTable in middle between AboutTable-Table and AboutTable-TableProperties`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableRow.X = 668.500000
	__Vertice__000001_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableRow.Y = 143.000000
	__Vertice__000001_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableRow.Name = `Verticle in class diagram AboutTable in middle between AboutTable-Table and AboutTable-TableRow`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableColumn_and_AboutTable_Paragraph.X = 877.000000
	__Vertice__000002_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableColumn_and_AboutTable_Paragraph.Y = 465.500000
	__Vertice__000002_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableColumn_and_AboutTable_Paragraph.Name = `Verticle in class diagram AboutTable in middle between AboutTable-TableColumn and AboutTable-Paragraph`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableProperties_and_AboutTable_TableStyle.X = 418.500000
	__Vertice__000003_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableProperties_and_AboutTable_TableStyle.Y = 209.000000
	__Vertice__000003_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableProperties_and_AboutTable_TableStyle.Name = `Verticle in class diagram AboutTable in middle between AboutTable-TableProperties and AboutTable-TableStyle`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableRow_and_AboutTable_TableColumn.X = 878.500000
	__Vertice__000004_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableRow_and_AboutTable_TableColumn.Y = 253.500000
	__Vertice__000004_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableRow_and_AboutTable_TableColumn.Name = `Verticle in class diagram AboutTable in middle between AboutTable-TableRow and AboutTable-TableColumn`

	// Setup of pointers
	__Classdiagram__000000_AboutTable.GongStructShapes = append(__Classdiagram__000000_AboutTable.GongStructShapes, __GongStructShape__000001_AboutTable_Table)
	__Classdiagram__000000_AboutTable.GongStructShapes = append(__Classdiagram__000000_AboutTable.GongStructShapes, __GongStructShape__000003_AboutTable_TableProperties)
	__Classdiagram__000000_AboutTable.GongStructShapes = append(__Classdiagram__000000_AboutTable.GongStructShapes, __GongStructShape__000005_AboutTable_TableStyle)
	__Classdiagram__000000_AboutTable.GongStructShapes = append(__Classdiagram__000000_AboutTable.GongStructShapes, __GongStructShape__000004_AboutTable_TableRow)
	__Classdiagram__000000_AboutTable.GongStructShapes = append(__Classdiagram__000000_AboutTable.GongStructShapes, __GongStructShape__000002_AboutTable_TableColumn)
	__Classdiagram__000000_AboutTable.GongStructShapes = append(__Classdiagram__000000_AboutTable.GongStructShapes, __GongStructShape__000000_AboutTable_Paragraph)
	__Classdiagram__000000_AboutTable.NoteShapes = append(__Classdiagram__000000_AboutTable.NoteShapes, __NoteShape__000001_NoteOnTable)
	__Classdiagram__000000_AboutTable.NoteShapes = append(__Classdiagram__000000_AboutTable.NoteShapes, __NoteShape__000000_NoteOnColumn)
	__GongStructShape__000000_AboutTable_Paragraph.Position = __Position__000000_Pos_AboutTable_Paragraph
	__GongStructShape__000000_AboutTable_Paragraph.Fields = append(__GongStructShape__000000_AboutTable_Paragraph.Fields, __Field__000005_Name)
	__GongStructShape__000001_AboutTable_Table.Position = __Position__000001_Pos_AboutTable_Table
	__GongStructShape__000001_AboutTable_Table.Fields = append(__GongStructShape__000001_AboutTable_Table.Fields, __Field__000000_Name)
	__GongStructShape__000001_AboutTable_Table.Links = append(__GongStructShape__000001_AboutTable_Table.Links, __Link__000002_TableProperties)
	__GongStructShape__000001_AboutTable_Table.Links = append(__GongStructShape__000001_AboutTable_Table.Links, __Link__000003_TableRows)
	__GongStructShape__000002_AboutTable_TableColumn.Position = __Position__000002_Pos_AboutTable_TableColumn
	__GongStructShape__000002_AboutTable_TableColumn.Fields = append(__GongStructShape__000002_AboutTable_TableColumn.Fields, __Field__000004_Name)
	__GongStructShape__000002_AboutTable_TableColumn.Links = append(__GongStructShape__000002_AboutTable_TableColumn.Links, __Link__000000_Paragraphs)
	__GongStructShape__000003_AboutTable_TableProperties.Position = __Position__000003_Pos_AboutTable_TableProperties
	__GongStructShape__000003_AboutTable_TableProperties.Fields = append(__GongStructShape__000003_AboutTable_TableProperties.Fields, __Field__000003_Name)
	__GongStructShape__000003_AboutTable_TableProperties.Links = append(__GongStructShape__000003_AboutTable_TableProperties.Links, __Link__000004_TableStyle)
	__GongStructShape__000004_AboutTable_TableRow.Position = __Position__000004_Pos_AboutTable_TableRow
	__GongStructShape__000004_AboutTable_TableRow.Fields = append(__GongStructShape__000004_AboutTable_TableRow.Fields, __Field__000002_Name)
	__GongStructShape__000004_AboutTable_TableRow.Links = append(__GongStructShape__000004_AboutTable_TableRow.Links, __Link__000001_TableColumns)
	__GongStructShape__000005_AboutTable_TableStyle.Position = __Position__000005_Pos_AboutTable_TableStyle
	__GongStructShape__000005_AboutTable_TableStyle.Fields = append(__GongStructShape__000005_AboutTable_TableStyle.Fields, __Field__000001_Name)
	__GongStructShape__000005_AboutTable_TableStyle.Fields = append(__GongStructShape__000005_AboutTable_TableStyle.Fields, __Field__000006_Val)
	__Link__000000_Paragraphs.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableColumn_and_AboutTable_Paragraph
	__Link__000001_TableColumns.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableRow_and_AboutTable_TableColumn
	__Link__000002_TableProperties.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableProperties
	__Link__000003_TableRows.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_Table_and_AboutTable_TableRow
	__Link__000004_TableStyle.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_AboutTable_in_middle_between_AboutTable_TableProperties_and_AboutTable_TableStyle
}
