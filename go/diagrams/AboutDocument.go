package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdocx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram map[string]any = map[string]any{
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
// 	InjectionGateway["AboutDocument"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "AboutDocument"
func NewDiagramInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `AboutDocument`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Document := (&models.GongStructShape{Name: `AboutDocument-Document`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Docx := (&models.GongStructShape{Name: `AboutDocument-Docx`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_File := (&models.GongStructShape{Name: `AboutDocument-File`}).Stage(stage)
	__GongStructShape__000003_NewDiagram_Node := (&models.GongStructShape{Name: `AboutDocument-Node`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_File := (&models.Link{Name: `File`}).Stage(stage)
	__Link__000001_Files := (&models.Link{Name: `Files`}).Stage(stage)
	__Link__000002_Nodes := (&models.Link{Name: `Nodes`}).Stage(stage)
	__Link__000003_Root := (&models.Link{Name: `Root`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Document := (&models.Position{Name: `Pos-AboutDocument-Document`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_Docx := (&models.Position{Name: `Pos-AboutDocument-Docx`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_File := (&models.Position{Name: `Pos-AboutDocument-File`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_Node := (&models.Position{Name: `Pos-AboutDocument-Node`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File := (&models.Vertice{Name: `Verticle in class diagram AboutDocument in middle between AboutDocument-Document and AboutDocument-File`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_Node := (&models.Vertice{Name: `Verticle in class diagram AboutDocument in middle between AboutDocument-Document and AboutDocument-Node`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File := (&models.Vertice{Name: `Verticle in class diagram AboutDocument in middle between AboutDocument-Docx and AboutDocument-File`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Node_and_NewDiagram_Node := (&models.Vertice{Name: `Verticle in class diagram AboutDocument in middle between AboutDocument-Node and AboutDocument-Node`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `AboutDocument`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.File.Name]
	__Field__000000_Name.Identifier = `ref_models.File.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `File`
	__Field__000000_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.Name]
	__Field__000001_Name.Identifier = `ref_models.Node.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Node`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document.Name]
	__Field__000002_Name.Identifier = `ref_models.Document.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Document`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Docx.Name]
	__Field__000003_Name.Identifier = `ref_models.Docx.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `Docx`
	__Field__000003_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Document.Name = `AboutDocument-Document`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document]
	__GongStructShape__000000_NewDiagram_Document.Identifier = `ref_models.Document`
	__GongStructShape__000000_NewDiagram_Document.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_Document.NbInstances = 1
	__GongStructShape__000000_NewDiagram_Document.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Document.Heigth = 78.000000
	__GongStructShape__000000_NewDiagram_Document.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Docx.Name = `AboutDocument-Docx`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Docx]
	__GongStructShape__000001_NewDiagram_Docx.Identifier = `ref_models.Docx`
	__GongStructShape__000001_NewDiagram_Docx.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_Docx.NbInstances = 1
	__GongStructShape__000001_NewDiagram_Docx.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Docx.Heigth = 78.000000
	__GongStructShape__000001_NewDiagram_Docx.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_File.Name = `AboutDocument-File`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.File]
	__GongStructShape__000002_NewDiagram_File.Identifier = `ref_models.File`
	__GongStructShape__000002_NewDiagram_File.ShowNbInstances = true
	__GongStructShape__000002_NewDiagram_File.NbInstances = 11
	__GongStructShape__000002_NewDiagram_File.Width = 244.000000
	__GongStructShape__000002_NewDiagram_File.Heigth = 78.000000
	__GongStructShape__000002_NewDiagram_File.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_Node.Name = `AboutDocument-Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__GongStructShape__000003_NewDiagram_Node.Identifier = `ref_models.Node`
	__GongStructShape__000003_NewDiagram_Node.ShowNbInstances = true
	__GongStructShape__000003_NewDiagram_Node.NbInstances = 110
	__GongStructShape__000003_NewDiagram_Node.Width = 240.000000
	__GongStructShape__000003_NewDiagram_Node.Heigth = 131.000000
	__GongStructShape__000003_NewDiagram_Node.IsSelected = false

	// Link values setup
	__Link__000000_File.Name = `File`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document.File]
	__Link__000000_File.Identifier = `ref_models.Document.File`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.File]
	__Link__000000_File.Fieldtypename = `ref_models.File`
	__Link__000000_File.FieldOffsetX = -42.000000
	__Link__000000_File.FieldOffsetY = 27.000000
	__Link__000000_File.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_File.TargetMultiplicityOffsetX = 13.000000
	__Link__000000_File.TargetMultiplicityOffsetY = 29.000000
	__Link__000000_File.SourceMultiplicity = models.MANY
	__Link__000000_File.SourceMultiplicityOffsetX = 13.000000
	__Link__000000_File.SourceMultiplicityOffsetY = -6.000000
	__Link__000000_File.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_File.StartRatio = 0.516667
	__Link__000000_File.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_File.EndRatio = 0.504098
	__Link__000000_File.CornerOffsetRatio = -0.538462

	// Link values setup
	__Link__000001_Files.Name = `Files`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Docx.Files]
	__Link__000001_Files.Identifier = `ref_models.Docx.Files`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.File]
	__Link__000001_Files.Fieldtypename = `ref_models.File`
	__Link__000001_Files.FieldOffsetX = -50.000000
	__Link__000001_Files.FieldOffsetY = -16.000000
	__Link__000001_Files.TargetMultiplicity = models.MANY
	__Link__000001_Files.TargetMultiplicityOffsetX = 20.000000
	__Link__000001_Files.TargetMultiplicityOffsetY = -10.000000
	__Link__000001_Files.SourceMultiplicity = models.ZERO_ONE
	__Link__000001_Files.SourceMultiplicityOffsetX = 17.000000
	__Link__000001_Files.SourceMultiplicityOffsetY = 25.000000
	__Link__000001_Files.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_Files.StartRatio = 0.512500
	__Link__000001_Files.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_Files.EndRatio = 0.520492
	__Link__000001_Files.CornerOffsetRatio = 1.410256

	// Link values setup
	__Link__000002_Nodes.Name = `Nodes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.Nodes]
	__Link__000002_Nodes.Identifier = `ref_models.Node.Nodes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000002_Nodes.Fieldtypename = `ref_models.Node`
	__Link__000002_Nodes.FieldOffsetX = -78.000000
	__Link__000002_Nodes.FieldOffsetY = -16.000000
	__Link__000002_Nodes.TargetMultiplicity = models.MANY
	__Link__000002_Nodes.TargetMultiplicityOffsetX = -26.000000
	__Link__000002_Nodes.TargetMultiplicityOffsetY = 22.000000
	__Link__000002_Nodes.SourceMultiplicity = models.ZERO_ONE
	__Link__000002_Nodes.SourceMultiplicityOffsetX = -42.000000
	__Link__000002_Nodes.SourceMultiplicityOffsetY = 23.000000
	__Link__000002_Nodes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Nodes.StartRatio = 0.692308
	__Link__000002_Nodes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Nodes.EndRatio = 0.217949
	__Link__000002_Nodes.CornerOffsetRatio = -0.545833

	// Link values setup
	__Link__000003_Root.Name = `Root`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document.Root]
	__Link__000003_Root.Identifier = `ref_models.Document.Root`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000003_Root.Fieldtypename = `ref_models.Node`
	__Link__000003_Root.FieldOffsetX = -50.000000
	__Link__000003_Root.FieldOffsetY = -16.000000
	__Link__000003_Root.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_Root.TargetMultiplicityOffsetX = 13.000000
	__Link__000003_Root.TargetMultiplicityOffsetY = -12.000000
	__Link__000003_Root.SourceMultiplicity = models.MANY
	__Link__000003_Root.SourceMultiplicityOffsetX = 16.000000
	__Link__000003_Root.SourceMultiplicityOffsetY = 19.000000
	__Link__000003_Root.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Root.StartRatio = 0.520833
	__Link__000003_Root.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Root.EndRatio = 0.516667
	__Link__000003_Root.CornerOffsetRatio = 1.653846

	// Position values setup
	__Position__000000_Pos_NewDiagram_Document.X = 192.000000
	__Position__000000_Pos_NewDiagram_Document.Y = 391.000000
	__Position__000000_Pos_NewDiagram_Document.Name = `Pos-AboutDocument-Document`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Docx.X = 196.000000
	__Position__000001_Pos_NewDiagram_Docx.Y = 77.000000
	__Position__000001_Pos_NewDiagram_Docx.Name = `Pos-AboutDocument-Docx`

	// Position values setup
	__Position__000002_Pos_NewDiagram_File.X = 192.000000
	__Position__000002_Pos_NewDiagram_File.Y = 239.000000
	__Position__000002_Pos_NewDiagram_File.Name = `Pos-AboutDocument-File`

	// Position values setup
	__Position__000003_Pos_NewDiagram_Node.X = 194.000000
	__Position__000003_Pos_NewDiagram_Node.Y = 554.000000
	__Position__000003_Pos_NewDiagram_Node.Name = `Pos-AboutDocument-Node`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File.X = 957.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File.Y = 211.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File.Name = `Verticle in class diagram AboutDocument in middle between AboutDocument-Document and AboutDocument-File`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_Node.X = 961.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_Node.Y = 419.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_Node.Name = `Verticle in class diagram AboutDocument in middle between AboutDocument-Document and AboutDocument-Node`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File.X = 465.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File.Y = 109.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File.Name = `Verticle in class diagram AboutDocument in middle between AboutDocument-Docx and AboutDocument-File`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Node_and_NewDiagram_Node.X = 941.000000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Node_and_NewDiagram_Node.Y = 503.500000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Node_and_NewDiagram_Node.Name = `Verticle in class diagram AboutDocument in middle between AboutDocument-Node and AboutDocument-Node`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Docx)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_File)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Document)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_Node)
	__GongStructShape__000000_NewDiagram_Document.Position = __Position__000000_Pos_NewDiagram_Document
	__GongStructShape__000000_NewDiagram_Document.Fields = append(__GongStructShape__000000_NewDiagram_Document.Fields, __Field__000002_Name)
	__GongStructShape__000000_NewDiagram_Document.Links = append(__GongStructShape__000000_NewDiagram_Document.Links, __Link__000000_File)
	__GongStructShape__000000_NewDiagram_Document.Links = append(__GongStructShape__000000_NewDiagram_Document.Links, __Link__000003_Root)
	__GongStructShape__000001_NewDiagram_Docx.Position = __Position__000001_Pos_NewDiagram_Docx
	__GongStructShape__000001_NewDiagram_Docx.Fields = append(__GongStructShape__000001_NewDiagram_Docx.Fields, __Field__000003_Name)
	__GongStructShape__000001_NewDiagram_Docx.Links = append(__GongStructShape__000001_NewDiagram_Docx.Links, __Link__000001_Files)
	__GongStructShape__000002_NewDiagram_File.Position = __Position__000002_Pos_NewDiagram_File
	__GongStructShape__000002_NewDiagram_File.Fields = append(__GongStructShape__000002_NewDiagram_File.Fields, __Field__000000_Name)
	__GongStructShape__000003_NewDiagram_Node.Position = __Position__000003_Pos_NewDiagram_Node
	__GongStructShape__000003_NewDiagram_Node.Fields = append(__GongStructShape__000003_NewDiagram_Node.Fields, __Field__000001_Name)
	__GongStructShape__000003_NewDiagram_Node.Links = append(__GongStructShape__000003_NewDiagram_Node.Links, __Link__000002_Nodes)
	__Link__000000_File.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File
	__Link__000001_Files.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File
	__Link__000002_Nodes.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Node_and_NewDiagram_Node
	__Link__000003_Root.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_Node
}
