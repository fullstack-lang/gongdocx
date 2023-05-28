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

	"ref_models.Docx": &(ref_models.Docx{}),

	"ref_models.Docx.Files": (ref_models.Docx{}).Files,

	"ref_models.Docx.Name": (ref_models.Docx{}).Name,

	"ref_models.File": &(ref_models.File{}),

	"ref_models.File.Name": (ref_models.File{}).Name,
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

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Docx := (&models.GongStructShape{Name: `NewDiagram-Docx`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_File := (&models.GongStructShape{Name: `NewDiagram-File`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Files := (&models.Link{Name: `Files`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Docx := (&models.Position{Name: `Pos-NewDiagram-Docx`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_File := (&models.Position{Name: `Pos-NewDiagram-File`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Docx and NewDiagram-File`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Docx.Name]
	__Field__000000_Name.Identifier = `ref_models.Docx.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Docx`
	__Field__000000_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.File.Name]
	__Field__000001_Name.Identifier = `ref_models.File.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `File`
	__Field__000001_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Docx.Name = `NewDiagram-Docx`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Docx]
	__GongStructShape__000000_NewDiagram_Docx.Identifier = `ref_models.Docx`
	__GongStructShape__000000_NewDiagram_Docx.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Docx.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Docx.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Docx.Heigth = 78.000000
	__GongStructShape__000000_NewDiagram_Docx.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_File.Name = `NewDiagram-File`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.File]
	__GongStructShape__000001_NewDiagram_File.Identifier = `ref_models.File`
	__GongStructShape__000001_NewDiagram_File.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_File.NbInstances = 0
	__GongStructShape__000001_NewDiagram_File.Width = 240.000000
	__GongStructShape__000001_NewDiagram_File.Heigth = 78.000000
	__GongStructShape__000001_NewDiagram_File.IsSelected = false

	// Link values setup
	__Link__000000_Files.Name = `Files`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Docx.Files]
	__Link__000000_Files.Identifier = `ref_models.Docx.Files`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.File]
	__Link__000000_Files.Fieldtypename = `ref_models.File`
	__Link__000000_Files.FieldOffsetX = -50.000000
	__Link__000000_Files.FieldOffsetY = -16.000000
	__Link__000000_Files.TargetMultiplicity = models.MANY
	__Link__000000_Files.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_Files.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_Files.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_Files.SourceMultiplicityOffsetX = 17.000000
	__Link__000000_Files.SourceMultiplicityOffsetY = 25.000000
	__Link__000000_Files.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Files.StartRatio = 0.500000
	__Link__000000_Files.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Files.EndRatio = 0.500000
	__Link__000000_Files.CornerOffsetRatio = 1.433333

	// Position values setup
	__Position__000000_Pos_NewDiagram_Docx.X = 105.000000
	__Position__000000_Pos_NewDiagram_Docx.Y = 73.000000
	__Position__000000_Pos_NewDiagram_Docx.Name = `Pos-NewDiagram-Docx`

	// Position values setup
	__Position__000001_Pos_NewDiagram_File.X = 595.000000
	__Position__000001_Pos_NewDiagram_File.Y = 72.000000
	__Position__000001_Pos_NewDiagram_File.Name = `Pos-NewDiagram-File`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File.X = 465.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File.Y = 109.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Docx and NewDiagram-File`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Docx)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_File)
	__GongStructShape__000000_NewDiagram_Docx.Position = __Position__000000_Pos_NewDiagram_Docx
	__GongStructShape__000000_NewDiagram_Docx.Fields = append(__GongStructShape__000000_NewDiagram_Docx.Fields, __Field__000000_Name)
	__GongStructShape__000000_NewDiagram_Docx.Links = append(__GongStructShape__000000_NewDiagram_Docx.Links, __Link__000000_Files)
	__GongStructShape__000001_NewDiagram_File.Position = __Position__000001_Pos_NewDiagram_File
	__GongStructShape__000001_NewDiagram_File.Fields = append(__GongStructShape__000001_NewDiagram_File.Fields, __Field__000001_Name)
	__Link__000000_Files.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Docx_and_NewDiagram_File
}


