package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdocx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram_1 models.StageStruct
var ___dummy__Time_NewDiagram_1 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram_1 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram_1 map[string]any = map[string]any{
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

	"ref_models.Rune.Text": (ref_models.Rune{}).Text,

	"ref_models.RuneProperties": &(ref_models.RuneProperties{}),

	"ref_models.RuneProperties.Content": (ref_models.RuneProperties{}).Content,

	"ref_models.RuneProperties.IsBold": (ref_models.RuneProperties{}).IsBold,

	"ref_models.RuneProperties.IsItalic": (ref_models.RuneProperties{}).IsItalic,

	"ref_models.RuneProperties.IsStrike": (ref_models.RuneProperties{}).IsStrike,

	"ref_models.RuneProperties.Name": (ref_models.RuneProperties{}).Name,

	"ref_models.RuneProperties.Node": (ref_models.RuneProperties{}).Node,

	"ref_models.Text": &(ref_models.Text{}),

	"ref_models.Text.Content": (ref_models.Text{}).Content,

	"ref_models.Text.Name": (ref_models.Text{}).Name,

	"ref_models.Text.Node": (ref_models.Text{}).Node,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram_1"] = NewDiagram_1Injection
// }

// NewDiagram_1Injection will stage objects of database "NewDiagram_1"
func NewDiagram_1Injection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram_1 := (&models.Classdiagram{Name: `NewDiagram_1`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Content := (&models.Field{Name: `Content`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_1_Document := (&models.GongStructShape{Name: `NewDiagram_1-Document`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_1_Node := (&models.GongStructShape{Name: `NewDiagram_1-Node`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_1_Paragraph := (&models.GongStructShape{Name: `NewDiagram_1-Paragraph`}).Stage(stage)
	__GongStructShape__000003_NewDiagram_1_Rune := (&models.GongStructShape{Name: `NewDiagram_1-Rune`}).Stage(stage)
	__GongStructShape__000004_NewDiagram_1_Text := (&models.GongStructShape{Name: `NewDiagram_1-Text`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Node := (&models.Link{Name: `Node`}).Stage(stage)
	__Link__000001_Node := (&models.Link{Name: `Node`}).Stage(stage)
	__Link__000002_Node := (&models.Link{Name: `Node`}).Stage(stage)
	__Link__000003_Nodes := (&models.Link{Name: `Nodes`}).Stage(stage)
	__Link__000004_Root := (&models.Link{Name: `Root`}).Stage(stage)
	__Link__000005_Runes := (&models.Link{Name: `Runes`}).Stage(stage)
	__Link__000006_Text := (&models.Link{Name: `Text`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_NodeOnText := (&models.NoteShape{Name: `NodeOnText`}).Stage(stage)
	__NoteShape__000001_NoteOnDocument := (&models.NoteShape{Name: `NoteOnDocument`}).Stage(stage)
	__NoteShape__000002_NoteOnParagraph := (&models.NoteShape{Name: `NoteOnParagraph`}).Stage(stage)
	__NoteShape__000003_NoteOnRune := (&models.NoteShape{Name: `NoteOnRune`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink
	__NoteShapeLink__000000_Paragraph := (&models.NoteShapeLink{Name: `Paragraph`}).Stage(stage)
	__NoteShapeLink__000001_Rune := (&models.NoteShapeLink{Name: `Rune`}).Stage(stage)
	__NoteShapeLink__000002_Text := (&models.NoteShapeLink{Name: `Text`}).Stage(stage)

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_1_Document := (&models.Position{Name: `Pos-NewDiagram_1-Document`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_1_Node := (&models.Position{Name: `Pos-NewDiagram_1-Node`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_1_Paragraph := (&models.Position{Name: `Pos-NewDiagram_1-Paragraph`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_1_Rune := (&models.Position{Name: `Pos-NewDiagram_1-Rune`}).Stage(stage)
	__Position__000004_Pos_NewDiagram_1_Text := (&models.Position{Name: `Pos-NewDiagram_1-Text`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Document and NewDiagram_1-Node`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Node_and_NewDiagram_1_Node := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Node and NewDiagram_1-Node`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Node := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Paragraph and NewDiagram_1-Node`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Rune := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Paragraph and NewDiagram_1-Rune`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Node := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Rune and NewDiagram_1-Node`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Text := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Rune and NewDiagram_1-Text`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Text_and_NewDiagram_1_Node := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Text and NewDiagram_1-Node`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram_1.Name = `NewDiagram_1`
	__Classdiagram__000000_NewDiagram_1.IsInDrawMode = true

	// Field values setup
	__Field__000000_Content.Name = `Content`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune.Content]
	__Field__000000_Content.Identifier = `ref_models.Rune.Content`
	__Field__000000_Content.FieldTypeAsString = ``
	__Field__000000_Content.Structname = `Rune`
	__Field__000000_Content.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document.Name]
	__Field__000001_Name.Identifier = `ref_models.Document.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Document`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text.Name]
	__Field__000002_Name.Identifier = `ref_models.Text.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Text`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph.Name]
	__Field__000003_Name.Identifier = `ref_models.Paragraph.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `Paragraph`
	__Field__000003_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_1_Document.Name = `NewDiagram_1-Document`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document]
	__GongStructShape__000000_NewDiagram_1_Document.Identifier = `ref_models.Document`
	__GongStructShape__000000_NewDiagram_1_Document.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_1_Document.NbInstances = 1
	__GongStructShape__000000_NewDiagram_1_Document.Width = 240.000000
	__GongStructShape__000000_NewDiagram_1_Document.Heigth = 78.000000
	__GongStructShape__000000_NewDiagram_1_Document.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_1_Node.Name = `NewDiagram_1-Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__GongStructShape__000001_NewDiagram_1_Node.Identifier = `ref_models.Node`
	__GongStructShape__000001_NewDiagram_1_Node.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_1_Node.NbInstances = 110
	__GongStructShape__000001_NewDiagram_1_Node.Width = 1199.000000
	__GongStructShape__000001_NewDiagram_1_Node.Heigth = 63.000000
	__GongStructShape__000001_NewDiagram_1_Node.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_1_Paragraph.Name = `NewDiagram_1-Paragraph`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph]
	__GongStructShape__000002_NewDiagram_1_Paragraph.Identifier = `ref_models.Paragraph`
	__GongStructShape__000002_NewDiagram_1_Paragraph.ShowNbInstances = true
	__GongStructShape__000002_NewDiagram_1_Paragraph.NbInstances = 12
	__GongStructShape__000002_NewDiagram_1_Paragraph.Width = 240.000000
	__GongStructShape__000002_NewDiagram_1_Paragraph.Heigth = 78.000000
	__GongStructShape__000002_NewDiagram_1_Paragraph.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_1_Rune.Name = `NewDiagram_1-Rune`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune]
	__GongStructShape__000003_NewDiagram_1_Rune.Identifier = `ref_models.Rune`
	__GongStructShape__000003_NewDiagram_1_Rune.ShowNbInstances = true
	__GongStructShape__000003_NewDiagram_1_Rune.NbInstances = 15
	__GongStructShape__000003_NewDiagram_1_Rune.Width = 240.000000
	__GongStructShape__000003_NewDiagram_1_Rune.Heigth = 78.000000
	__GongStructShape__000003_NewDiagram_1_Rune.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_1_Text.Name = `NewDiagram_1-Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__GongStructShape__000004_NewDiagram_1_Text.Identifier = `ref_models.Text`
	__GongStructShape__000004_NewDiagram_1_Text.ShowNbInstances = true
	__GongStructShape__000004_NewDiagram_1_Text.NbInstances = 15
	__GongStructShape__000004_NewDiagram_1_Text.Width = 240.000000
	__GongStructShape__000004_NewDiagram_1_Text.Heigth = 78.000000
	__GongStructShape__000004_NewDiagram_1_Text.IsSelected = false

	// Link values setup
	__Link__000000_Node.Name = `Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph.Node]
	__Link__000000_Node.Identifier = `ref_models.Paragraph.Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000000_Node.Fieldtypename = `ref_models.Node`
	__Link__000000_Node.FieldOffsetX = 13.000000
	__Link__000000_Node.FieldOffsetY = 24.000000
	__Link__000000_Node.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_Node.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_Node.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_Node.SourceMultiplicity = models.MANY
	__Link__000000_Node.SourceMultiplicityOffsetX = -26.000000
	__Link__000000_Node.SourceMultiplicityOffsetY = -10.000000
	__Link__000000_Node.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Node.StartRatio = 0.462500
	__Link__000000_Node.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Node.EndRatio = 0.080067
	__Link__000000_Node.CornerOffsetRatio = -1.346154

	// Link values setup
	__Link__000001_Node.Name = `Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text.Node]
	__Link__000001_Node.Identifier = `ref_models.Text.Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000001_Node.Fieldtypename = `ref_models.Node`
	__Link__000001_Node.FieldOffsetX = 15.000000
	__Link__000001_Node.FieldOffsetY = 21.000000
	__Link__000001_Node.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Node.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_Node.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_Node.SourceMultiplicity = models.MANY
	__Link__000001_Node.SourceMultiplicityOffsetX = -32.000000
	__Link__000001_Node.SourceMultiplicityOffsetY = -12.000000
	__Link__000001_Node.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_Node.StartRatio = 0.533333
	__Link__000001_Node.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_Node.EndRatio = 0.909925
	__Link__000001_Node.CornerOffsetRatio = -1.051282

	// Link values setup
	__Link__000002_Node.Name = `Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune.Node]
	__Link__000002_Node.Identifier = `ref_models.Rune.Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000002_Node.Fieldtypename = `ref_models.Node`
	__Link__000002_Node.FieldOffsetX = 24.000000
	__Link__000002_Node.FieldOffsetY = 23.000000
	__Link__000002_Node.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_Node.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_Node.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_Node.SourceMultiplicity = models.MANY
	__Link__000002_Node.SourceMultiplicityOffsetX = -31.000000
	__Link__000002_Node.SourceMultiplicityOffsetY = -12.000000
	__Link__000002_Node.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_Node.StartRatio = 0.650000
	__Link__000002_Node.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_Node.EndRatio = 0.513761
	__Link__000002_Node.CornerOffsetRatio = -0.198718

	// Link values setup
	__Link__000003_Nodes.Name = `Nodes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.Nodes]
	__Link__000003_Nodes.Identifier = `ref_models.Node.Nodes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000003_Nodes.Fieldtypename = `ref_models.Node`
	__Link__000003_Nodes.FieldOffsetX = -61.000000
	__Link__000003_Nodes.FieldOffsetY = -14.000000
	__Link__000003_Nodes.TargetMultiplicity = models.MANY
	__Link__000003_Nodes.TargetMultiplicityOffsetX = 11.000000
	__Link__000003_Nodes.TargetMultiplicityOffsetY = -9.000000
	__Link__000003_Nodes.SourceMultiplicity = models.ZERO_ONE
	__Link__000003_Nodes.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_Nodes.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_Nodes.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Nodes.StartRatio = 0.929108
	__Link__000003_Nodes.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Nodes.EndRatio = 0.842369
	__Link__000003_Nodes.CornerOffsetRatio = -0.706349

	// Link values setup
	__Link__000004_Root.Name = `Root`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document.Root]
	__Link__000004_Root.Identifier = `ref_models.Document.Root`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000004_Root.Fieldtypename = `ref_models.Node`
	__Link__000004_Root.FieldOffsetX = -50.000000
	__Link__000004_Root.FieldOffsetY = -16.000000
	__Link__000004_Root.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_Root.TargetMultiplicityOffsetX = 15.000000
	__Link__000004_Root.TargetMultiplicityOffsetY = -9.000000
	__Link__000004_Root.SourceMultiplicity = models.MANY
	__Link__000004_Root.SourceMultiplicityOffsetX = 17.000000
	__Link__000004_Root.SourceMultiplicityOffsetY = 25.000000
	__Link__000004_Root.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_Root.StartRatio = 0.541667
	__Link__000004_Root.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_Root.EndRatio = 0.111760
	__Link__000004_Root.CornerOffsetRatio = 2.012821

	// Link values setup
	__Link__000005_Runes.Name = `Runes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph.Runes]
	__Link__000005_Runes.Identifier = `ref_models.Paragraph.Runes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune]
	__Link__000005_Runes.Fieldtypename = `ref_models.Rune`
	__Link__000005_Runes.FieldOffsetX = -61.000000
	__Link__000005_Runes.FieldOffsetY = -18.000000
	__Link__000005_Runes.TargetMultiplicity = models.MANY
	__Link__000005_Runes.TargetMultiplicityOffsetX = -35.000000
	__Link__000005_Runes.TargetMultiplicityOffsetY = 33.000000
	__Link__000005_Runes.SourceMultiplicity = models.ZERO_ONE
	__Link__000005_Runes.SourceMultiplicityOffsetX = 17.000000
	__Link__000005_Runes.SourceMultiplicityOffsetY = 25.000000
	__Link__000005_Runes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_Runes.StartRatio = 0.500000
	__Link__000005_Runes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_Runes.EndRatio = 0.525641
	__Link__000005_Runes.CornerOffsetRatio = 1.375000

	// Link values setup
	__Link__000006_Text.Name = `Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune.Text]
	__Link__000006_Text.Identifier = `ref_models.Rune.Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__Link__000006_Text.Fieldtypename = `ref_models.Text`
	__Link__000006_Text.FieldOffsetX = -50.000000
	__Link__000006_Text.FieldOffsetY = -16.000000
	__Link__000006_Text.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_Text.TargetMultiplicityOffsetX = -41.000000
	__Link__000006_Text.TargetMultiplicityOffsetY = 27.000000
	__Link__000006_Text.SourceMultiplicity = models.MANY
	__Link__000006_Text.SourceMultiplicityOffsetX = 17.000000
	__Link__000006_Text.SourceMultiplicityOffsetY = 29.000000
	__Link__000006_Text.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_Text.StartRatio = 0.500000
	__Link__000006_Text.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_Text.EndRatio = 0.583333
	__Link__000006_Text.CornerOffsetRatio = 1.533333

	// NoteShape values setup
	__NoteShape__000000_NodeOnText.Name = `NodeOnText`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NodeOnText]
	__NoteShape__000000_NodeOnText.Identifier = `ref_models.NodeOnText`
	__NoteShape__000000_NodeOnText.Body = `:
[models.Text]
-
t: This stands for 'text'. This element is used to represent a run of text within
a paragraph. The t element will contain the actual string of text as its content.
-
When parsing these nodes, your code should handle each type of node appropriately
based on its name. For example, when you encounter a t node, you might simply
extract and store the text content.
`
	__NoteShape__000000_NodeOnText.BodyHTML = `<p>:
<a href="/models#Text">models.Text</a>
-
t: This stands for &apos;text&apos;. This element is used to represent a run of text within
a paragraph. The t element will contain the actual string of text as its content.
-
When parsing these nodes, your code should handle each type of node appropriately
based on its name. For example, when you encounter a t node, you might simply
extract and store the text content.
`
	__NoteShape__000000_NodeOnText.X = 701.000000
	__NoteShape__000000_NodeOnText.Y = 868.000000
	__NoteShape__000000_NodeOnText.Width = 670.000000
	__NoteShape__000000_NodeOnText.Heigth = 178.000000
	__NoteShape__000000_NodeOnText.Matched = false

	// NoteShape values setup
	__NoteShape__000001_NoteOnDocument.Name = `NoteOnDocument`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnDocument]
	__NoteShape__000001_NoteOnDocument.Identifier = `ref_models.NoteOnDocument`
	__NoteShape__000001_NoteOnDocument.Body = `
In the structure of a .docx file, the word/document.xml file is one of the most
crucial components. It contains the main content of the document, including the
text and its organization into paragraphs and other structures, as well as
references to other components of the document such as images, styles, formatting
instructions, etc.

 -

This XML file primarily houses the textual content and its associated XML tags
denote various properties of the text such as font size, style, alignment, and
more. All these pieces of information are represented using WordProcessingML, a
type of XML developed by Microsoft for Word documents.
-
The document.xml file refers to other files in the .docx structure to help render
the final document. For instance, it references styles from the styles.xml file,
numbered lists from the numbering.xml file, relationships from the
document.xml.rels file, and more. These references help assemble the complete,
rendered document that you see when you open a .docx file in a word processor.
-
There could be multiple document*.xml files in some situations. This is typically
seen when a Word document has been split into separate sections, perhaps for
editing or collaboration purposes. Each document*.xml file would contain a
different section of the overall document.
`
	__NoteShape__000001_NoteOnDocument.BodyHTML = `<p>In the structure of a .docx file, the word/document.xml file is one of the most
crucial components. It contains the main content of the document, including the
text and its organization into paragraphs and other structures, as well as
references to other components of the document such as images, styles, formatting
instructions, etc.
<pre>-
</pre>
<p>This XML file primarily houses the textual content and its associated XML tags
denote various properties of the text such as font size, style, alignment, and
more. All these pieces of information are represented using WordProcessingML, a
type of XML developed by Microsoft for Word documents.
-
The document.xml file refers to other files in the .docx structure to help render
the final document. For instance, it references styles from the styles.xml file,
numbered lists from the numbering.xml file, relationships from the
document.xml.rels file, and more. These references help assemble the complete,
rendered document that you see when you open a .docx file in a word processor.
-
There could be multiple document*.xml files in some situations. This is typically
seen when a Word document has been split into separate sections, perhaps for
editing or collaboration purposes. Each document*.xml file would contain a
different section of the overall document.
`
	__NoteShape__000001_NoteOnDocument.X = 356.000000
	__NoteShape__000001_NoteOnDocument.Y = 9.000000
	__NoteShape__000001_NoteOnDocument.Width = 669.000000
	__NoteShape__000001_NoteOnDocument.Heigth = 408.000000
	__NoteShape__000001_NoteOnDocument.Matched = false

	// NoteShape values setup
	__NoteShape__000002_NoteOnParagraph.Name = `NoteOnParagraph`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnParagraph]
	__NoteShape__000002_NoteOnParagraph.Identifier = `ref_models.NoteOnParagraph`
	__NoteShape__000002_NoteOnParagraph.Body = `
The "w:p" node represents a [models.Paragraph] in a Word document's XML structure.
-
It is one of the primary building blocks of a document's content within the
document.xml file.
-
Each "w:p" element contains a series of "w:r" (run) elements, which represent
sections of text within the paragraph that have consistent formatting. These
"w:r" elements, in turn, contain "w:t" elements that hold the actual text
content.
-
The "w:p" element may also contain various other child elements that provide
additional information about the paragraph. For example, "w:pPr" specifies
paragraph properties like alignment, indentation, and spacing.
-
As you parse a "w:p" node, you would typically create a new paragraph object in
your code, then parse the child nodes to fill in the content and properties of
that paragraph.
`
	__NoteShape__000002_NoteOnParagraph.BodyHTML = `<p>The &quot;w:p&quot; node represents a <a href="/models#Paragraph">models.Paragraph</a> in a Word document&apos;s XML structure.
-
It is one of the primary building blocks of a document&apos;s content within the
document.xml file.
-
Each &quot;w:p&quot; element contains a series of &quot;w:r&quot; (run) elements, which represent
sections of text within the paragraph that have consistent formatting. These
&quot;w:r&quot; elements, in turn, contain &quot;w:t&quot; elements that hold the actual text
content.
-
The &quot;w:p&quot; element may also contain various other child elements that provide
additional information about the paragraph. For example, &quot;w:pPr&quot; specifies
paragraph properties like alignment, indentation, and spacing.
-
As you parse a &quot;w:p&quot; node, you would typically create a new paragraph object in
your code, then parse the child nodes to fill in the content and properties of
that paragraph.
`
	__NoteShape__000002_NoteOnParagraph.X = 35.000000
	__NoteShape__000002_NoteOnParagraph.Y = 880.000000
	__NoteShape__000002_NoteOnParagraph.Width = 615.000000
	__NoteShape__000002_NoteOnParagraph.Heigth = 315.000000
	__NoteShape__000002_NoteOnParagraph.Matched = false

	// NoteShape values setup
	__NoteShape__000003_NoteOnRune.Name = `NoteOnRune`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnRune]
	__NoteShape__000003_NoteOnRune.Identifier = `ref_models.NoteOnRune`
	__NoteShape__000003_NoteOnRune.Body = `
for [models.Rune]
-
The "w:r" node, known as a run, represents a continuous run of text within a
paragraph ("w:p" node) in a Word document's XML structure. It is found within
the document.xml file.
-
Runs are segments of text within a paragraph that share the same formatting.
This can include properties like bolding, italics, underlining, color, size,
font, and more. The specific formatting is defined in a "w:rPr" (Run Properties)
element within the "w:r" node.
-
A "w:r" node contains one or more "w:t" nodes, which hold the actual text
content of the run. It can also contain other types of nodes like "w:br" for a
line break or "w:tab" for a tab character.
-
When parsing a "w:r" node, your code should handle the formatting information
provided in the "w:rPr" node (if present) and apply it to the text found within
the "w:t" nodes.
`
	__NoteShape__000003_NoteOnRune.BodyHTML = `<p>for <a href="/models#Rune">models.Rune</a>
-
The &quot;w:r&quot; node, known as a run, represents a continuous run of text within a
paragraph (&quot;w:p&quot; node) in a Word document&apos;s XML structure. It is found within
the document.xml file.
-
Runs are segments of text within a paragraph that share the same formatting.
This can include properties like bolding, italics, underlining, color, size,
font, and more. The specific formatting is defined in a &quot;w:rPr&quot; (Run Properties)
element within the &quot;w:r&quot; node.
-
A &quot;w:r&quot; node contains one or more &quot;w:t&quot; nodes, which hold the actual text
content of the run. It can also contain other types of nodes like &quot;w:br&quot; for a
line break or &quot;w:tab&quot; for a tab character.
-
When parsing a &quot;w:r&quot; node, your code should handle the formatting information
provided in the &quot;w:rPr&quot; node (if present) and apply it to the text found within
the &quot;w:t&quot; nodes.
`
	__NoteShape__000003_NoteOnRune.X = 685.000000
	__NoteShape__000003_NoteOnRune.Y = 1068.000000
	__NoteShape__000003_NoteOnRune.Width = 611.000000
	__NoteShape__000003_NoteOnRune.Heigth = 342.000000
	__NoteShape__000003_NoteOnRune.Matched = false

	// NoteShapeLink values setup
	__NoteShapeLink__000000_Paragraph.Name = `Paragraph`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph]
	__NoteShapeLink__000000_Paragraph.Identifier = `ref_models.Paragraph`
	__NoteShapeLink__000000_Paragraph.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// NoteShapeLink values setup
	__NoteShapeLink__000001_Rune.Name = `Rune`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune]
	__NoteShapeLink__000001_Rune.Identifier = `ref_models.Rune`
	__NoteShapeLink__000001_Rune.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// NoteShapeLink values setup
	__NoteShapeLink__000002_Text.Name = `Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__NoteShapeLink__000002_Text.Identifier = `ref_models.Text`
	__NoteShapeLink__000002_Text.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// Position values setup
	__Position__000000_Pos_NewDiagram_1_Document.X = 53.000000
	__Position__000000_Pos_NewDiagram_1_Document.Y = 18.000000
	__Position__000000_Pos_NewDiagram_1_Document.Name = `Pos-NewDiagram_1-Document`

	// Position values setup
	__Position__000001_Pos_NewDiagram_1_Node.X = 50.000000
	__Position__000001_Pos_NewDiagram_1_Node.Y = 483.000000
	__Position__000001_Pos_NewDiagram_1_Node.Name = `Pos-NewDiagram_1-Node`

	// Position values setup
	__Position__000002_Pos_NewDiagram_1_Paragraph.X = 41.000000
	__Position__000002_Pos_NewDiagram_1_Paragraph.Y = 714.000000
	__Position__000002_Pos_NewDiagram_1_Paragraph.Name = `Pos-NewDiagram_1-Paragraph`

	// Position values setup
	__Position__000003_Pos_NewDiagram_1_Rune.X = 510.000000
	__Position__000003_Pos_NewDiagram_1_Rune.Y = 714.000000
	__Position__000003_Pos_NewDiagram_1_Rune.Name = `Pos-NewDiagram_1-Rune`

	// Position values setup
	__Position__000004_Pos_NewDiagram_1_Text.X = 1017.000000
	__Position__000004_Pos_NewDiagram_1_Text.Y = 707.000000
	__Position__000004_Pos_NewDiagram_1_Text.Name = `Pos-NewDiagram_1-Text`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node.X = 429.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node.Y = 94.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Document and NewDiagram_1-Node`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Node_and_NewDiagram_1_Node.X = 1848.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Node_and_NewDiagram_1_Node.Y = 514.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Node_and_NewDiagram_1_Node.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Node and NewDiagram_1-Node`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Node.X = 638.500000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Node.Y = 613.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Node.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Paragraph and NewDiagram_1-Node`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Rune.X = 635.500000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Rune.Y = 753.000000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Rune.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Paragraph and NewDiagram_1-Rune`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Node.X = 640.000000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Node.Y = 630.000000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Node.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Rune and NewDiagram_1-Node`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Text.X = 1102.000000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Text.Y = 750.000000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Text.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Rune and NewDiagram_1-Text`

	// Vertice values setup
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Text_and_NewDiagram_1_Node.X = 405.000000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Text_and_NewDiagram_1_Node.Y = 413.000000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Text_and_NewDiagram_1_Node.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Text and NewDiagram_1-Node`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000000_NewDiagram_1_Document)
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000001_NewDiagram_1_Node)
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000004_NewDiagram_1_Text)
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000002_NewDiagram_1_Paragraph)
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000003_NewDiagram_1_Rune)
	__Classdiagram__000000_NewDiagram_1.NoteShapes = append(__Classdiagram__000000_NewDiagram_1.NoteShapes, __NoteShape__000001_NoteOnDocument)
	__Classdiagram__000000_NewDiagram_1.NoteShapes = append(__Classdiagram__000000_NewDiagram_1.NoteShapes, __NoteShape__000000_NodeOnText)
	__Classdiagram__000000_NewDiagram_1.NoteShapes = append(__Classdiagram__000000_NewDiagram_1.NoteShapes, __NoteShape__000002_NoteOnParagraph)
	__Classdiagram__000000_NewDiagram_1.NoteShapes = append(__Classdiagram__000000_NewDiagram_1.NoteShapes, __NoteShape__000003_NoteOnRune)
	__GongStructShape__000000_NewDiagram_1_Document.Position = __Position__000000_Pos_NewDiagram_1_Document
	__GongStructShape__000000_NewDiagram_1_Document.Fields = append(__GongStructShape__000000_NewDiagram_1_Document.Fields, __Field__000001_Name)
	__GongStructShape__000000_NewDiagram_1_Document.Links = append(__GongStructShape__000000_NewDiagram_1_Document.Links, __Link__000004_Root)
	__GongStructShape__000001_NewDiagram_1_Node.Position = __Position__000001_Pos_NewDiagram_1_Node
	__GongStructShape__000001_NewDiagram_1_Node.Links = append(__GongStructShape__000001_NewDiagram_1_Node.Links, __Link__000003_Nodes)
	__GongStructShape__000002_NewDiagram_1_Paragraph.Position = __Position__000002_Pos_NewDiagram_1_Paragraph
	__GongStructShape__000002_NewDiagram_1_Paragraph.Fields = append(__GongStructShape__000002_NewDiagram_1_Paragraph.Fields, __Field__000003_Name)
	__GongStructShape__000002_NewDiagram_1_Paragraph.Links = append(__GongStructShape__000002_NewDiagram_1_Paragraph.Links, __Link__000000_Node)
	__GongStructShape__000002_NewDiagram_1_Paragraph.Links = append(__GongStructShape__000002_NewDiagram_1_Paragraph.Links, __Link__000005_Runes)
	__GongStructShape__000003_NewDiagram_1_Rune.Position = __Position__000003_Pos_NewDiagram_1_Rune
	__GongStructShape__000003_NewDiagram_1_Rune.Fields = append(__GongStructShape__000003_NewDiagram_1_Rune.Fields, __Field__000000_Content)
	__GongStructShape__000003_NewDiagram_1_Rune.Links = append(__GongStructShape__000003_NewDiagram_1_Rune.Links, __Link__000002_Node)
	__GongStructShape__000003_NewDiagram_1_Rune.Links = append(__GongStructShape__000003_NewDiagram_1_Rune.Links, __Link__000006_Text)
	__GongStructShape__000004_NewDiagram_1_Text.Position = __Position__000004_Pos_NewDiagram_1_Text
	__GongStructShape__000004_NewDiagram_1_Text.Fields = append(__GongStructShape__000004_NewDiagram_1_Text.Fields, __Field__000002_Name)
	__GongStructShape__000004_NewDiagram_1_Text.Links = append(__GongStructShape__000004_NewDiagram_1_Text.Links, __Link__000001_Node)
	__Link__000000_Node.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Node
	__Link__000001_Node.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Text_and_NewDiagram_1_Node
	__Link__000002_Node.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Node
	__Link__000003_Nodes.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Node_and_NewDiagram_1_Node
	__Link__000004_Root.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node
	__Link__000005_Runes.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Paragraph_and_NewDiagram_1_Rune
	__Link__000006_Text.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Rune_and_NewDiagram_1_Text
	__NoteShape__000000_NodeOnText.NoteShapeLinks = append(__NoteShape__000000_NodeOnText.NoteShapeLinks, __NoteShapeLink__000002_Text)
	__NoteShape__000002_NoteOnParagraph.NoteShapeLinks = append(__NoteShape__000002_NoteOnParagraph.NoteShapeLinks, __NoteShapeLink__000000_Paragraph)
	__NoteShape__000003_NoteOnRune.NoteShapeLinks = append(__NoteShape__000003_NoteOnRune.NoteShapeLinks, __NoteShapeLink__000001_Rune)
}
