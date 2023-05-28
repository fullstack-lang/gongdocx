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

	"ref_models.Docx": &(ref_models.Docx{}),

	"ref_models.Docx.Files": (ref_models.Docx{}).Files,

	"ref_models.Docx.Name": (ref_models.Docx{}).Name,

	"ref_models.File": &(ref_models.File{}),

	"ref_models.File.Name": (ref_models.File{}).Name,

	"ref_models.NoteOnDocument": ref_models.NoteOnDocument,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Document := (&models.GongStructShape{Name: `NewDiagram-Document`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Docx := (&models.GongStructShape{Name: `NewDiagram-Docx`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_File := (&models.GongStructShape{Name: `NewDiagram-File`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_File := (&models.Link{Name: `File`}).Stage(stage)
	__Link__000001_Files := (&models.Link{Name: `Files`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Document := (&models.Position{Name: `Pos-NewDiagram-Document`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_Docx := (&models.Position{Name: `Pos-NewDiagram-Docx`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_File := (&models.Position{Name: `Pos-NewDiagram-File`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Document and NewDiagram-File`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Docx and NewDiagram-File`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
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

	//gong:ident [ref_models.Document.Name]
	__Field__000001_Name.Identifier = `ref_models.Document.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Document`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Docx.Name]
	__Field__000002_Name.Identifier = `ref_models.Docx.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Docx`
	__Field__000002_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Document.Name = `NewDiagram-Document`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document]
	__GongStructShape__000000_NewDiagram_Document.Identifier = `ref_models.Document`
	__GongStructShape__000000_NewDiagram_Document.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Document.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Document.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Document.Heigth = 78.000000
	__GongStructShape__000000_NewDiagram_Document.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Docx.Name = `NewDiagram-Docx`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Docx]
	__GongStructShape__000001_NewDiagram_Docx.Identifier = `ref_models.Docx`
	__GongStructShape__000001_NewDiagram_Docx.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_Docx.NbInstances = 1
	__GongStructShape__000001_NewDiagram_Docx.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Docx.Heigth = 78.000000
	__GongStructShape__000001_NewDiagram_Docx.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_File.Name = `NewDiagram-File`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.File]
	__GongStructShape__000002_NewDiagram_File.Identifier = `ref_models.File`
	__GongStructShape__000002_NewDiagram_File.ShowNbInstances = true
	__GongStructShape__000002_NewDiagram_File.NbInstances = 11
	__GongStructShape__000002_NewDiagram_File.Width = 240.000000
	__GongStructShape__000002_NewDiagram_File.Heigth = 78.000000
	__GongStructShape__000002_NewDiagram_File.IsSelected = false

	// Link values setup
	__Link__000000_File.Name = `File`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document.File]
	__Link__000000_File.Identifier = `ref_models.Document.File`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.File]
	__Link__000000_File.Fieldtypename = `ref_models.File`
	__Link__000000_File.FieldOffsetX = -48.000000
	__Link__000000_File.FieldOffsetY = 20.000000
	__Link__000000_File.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_File.TargetMultiplicityOffsetX = 12.000000
	__Link__000000_File.TargetMultiplicityOffsetY = 20.000000
	__Link__000000_File.SourceMultiplicity = models.MANY
	__Link__000000_File.SourceMultiplicityOffsetX = 13.000000
	__Link__000000_File.SourceMultiplicityOffsetY = -6.000000
	__Link__000000_File.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_File.StartRatio = 0.612500
	__Link__000000_File.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_File.EndRatio = 0.629167
	__Link__000000_File.CornerOffsetRatio = -0.141026

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
	__Link__000001_Files.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_Files.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_Files.SourceMultiplicity = models.ZERO_ONE
	__Link__000001_Files.SourceMultiplicityOffsetX = 17.000000
	__Link__000001_Files.SourceMultiplicityOffsetY = 25.000000
	__Link__000001_Files.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Files.StartRatio = 0.500000
	__Link__000001_Files.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Files.EndRatio = 0.500000
	__Link__000001_Files.CornerOffsetRatio = 1.433333

	// Position values setup
	__Position__000000_Pos_NewDiagram_Document.X = 599.000000
	__Position__000000_Pos_NewDiagram_Document.Y = 287.000000
	__Position__000000_Pos_NewDiagram_Document.Name = `Pos-NewDiagram-Document`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Docx.X = 105.000000
	__Position__000001_Pos_NewDiagram_Docx.Y = 73.000000
	__Position__000001_Pos_NewDiagram_Docx.Name = `Pos-NewDiagram-Docx`

	// Position values setup
	__Position__000002_Pos_NewDiagram_File.X = 595.000000
	__Position__000002_Pos_NewDiagram_File.Y = 72.000000
	__Position__000002_Pos_NewDiagram_File.Name = `Pos-NewDiagram-File`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File.X = 957.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File.Y = 211.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Document and NewDiagram-File`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File.X = 465.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File.Y = 109.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Docx and NewDiagram-File`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Docx)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_File)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Document)
	__GongStructShape__000000_NewDiagram_Document.Position = __Position__000000_Pos_NewDiagram_Document
	__GongStructShape__000000_NewDiagram_Document.Fields = append(__GongStructShape__000000_NewDiagram_Document.Fields, __Field__000001_Name)
	__GongStructShape__000000_NewDiagram_Document.Links = append(__GongStructShape__000000_NewDiagram_Document.Links, __Link__000000_File)
	__GongStructShape__000001_NewDiagram_Docx.Position = __Position__000001_Pos_NewDiagram_Docx
	__GongStructShape__000001_NewDiagram_Docx.Fields = append(__GongStructShape__000001_NewDiagram_Docx.Fields, __Field__000002_Name)
	__GongStructShape__000001_NewDiagram_Docx.Links = append(__GongStructShape__000001_NewDiagram_Docx.Links, __Link__000001_Files)
	__GongStructShape__000002_NewDiagram_File.Position = __Position__000002_Pos_NewDiagram_File
	__GongStructShape__000002_NewDiagram_File.Fields = append(__GongStructShape__000002_NewDiagram_File.Fields, __Field__000000_Name)
	__Link__000000_File.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Document_and_NewDiagram_File
	__Link__000001_Files.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File
}


