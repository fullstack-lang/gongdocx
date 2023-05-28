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

	"ref_models.NoteOnDocument": ref_models.NoteOnDocument,
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
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_1_Document := (&models.GongStructShape{Name: `NewDiagram_1-Document`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_1_Node := (&models.GongStructShape{Name: `NewDiagram_1-Node`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Root := (&models.Link{Name: `Root`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_NoteOnDocument := (&models.NoteShape{Name: `NoteOnDocument`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_1_Document := (&models.Position{Name: `Pos-NewDiagram_1-Document`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_1_Node := (&models.Position{Name: `Pos-NewDiagram_1-Node`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node := (&models.Vertice{Name: `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Document and NewDiagram_1-Node`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram_1.Name = `NewDiagram_1`
	__Classdiagram__000000_NewDiagram_1.IsInDrawMode = true

	// Field values setup
	__Field__000000_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document.Name]
	__Field__000000_Name.Identifier = `ref_models.Document.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Document`
	__Field__000000_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_1_Document.Name = `NewDiagram_1-Document`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document]
	__GongStructShape__000000_NewDiagram_1_Document.Identifier = `ref_models.Document`
	__GongStructShape__000000_NewDiagram_1_Document.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_1_Document.NbInstances = 0
	__GongStructShape__000000_NewDiagram_1_Document.Width = 240.000000
	__GongStructShape__000000_NewDiagram_1_Document.Heigth = 78.000000
	__GongStructShape__000000_NewDiagram_1_Document.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_1_Node.Name = `NewDiagram_1-Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__GongStructShape__000001_NewDiagram_1_Node.Identifier = `ref_models.Node`
	__GongStructShape__000001_NewDiagram_1_Node.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_1_Node.NbInstances = 0
	__GongStructShape__000001_NewDiagram_1_Node.Width = 240.000000
	__GongStructShape__000001_NewDiagram_1_Node.Heigth = 63.000000
	__GongStructShape__000001_NewDiagram_1_Node.IsSelected = false

	// Link values setup
	__Link__000000_Root.Name = `Root`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Document.Root]
	__Link__000000_Root.Identifier = `ref_models.Document.Root`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000000_Root.Fieldtypename = `ref_models.Node`
	__Link__000000_Root.FieldOffsetX = -50.000000
	__Link__000000_Root.FieldOffsetY = -16.000000
	__Link__000000_Root.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_Root.TargetMultiplicityOffsetX = 15.000000
	__Link__000000_Root.TargetMultiplicityOffsetY = -9.000000
	__Link__000000_Root.SourceMultiplicity = models.MANY
	__Link__000000_Root.SourceMultiplicityOffsetX = 17.000000
	__Link__000000_Root.SourceMultiplicityOffsetY = 25.000000
	__Link__000000_Root.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Root.StartRatio = 0.541667
	__Link__000000_Root.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Root.EndRatio = 0.591667
	__Link__000000_Root.CornerOffsetRatio = 2.012821

	// NoteShape values setup
	__NoteShape__000000_NoteOnDocument.Name = `NoteOnDocument`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnDocument]
	__NoteShape__000000_NoteOnDocument.Identifier = `ref_models.NoteOnDocument`
	__NoteShape__000000_NoteOnDocument.Body = `
In the structure of a .docx file, the word/document.xml file is one of the most
crucial components. It contains the main content of the document, including the
text and its organization into paragraphs and other structures, as well as
references to other components of the document such as images, styles, formatting
instructions, etc.

This XML file primarily houses the textual content and its associated XML tags
denote various properties of the text such as font size, style, alignment, and
more. All these pieces of information are represented using WordProcessingML, a
type of XML developed by Microsoft for Word documents.

The document.xml file refers to other files in the .docx structure to help render
the final document. For instance, it references styles from the styles.xml file,
numbered lists from the numbering.xml file, relationships from the
document.xml.rels file, and more. These references help assemble the complete,
rendered document that you see when you open a .docx file in a word processor.

There could be multiple document*.xml files in some situations. This is typically
seen when a Word document has been split into separate sections, perhaps for
editing or collaboration purposes. Each document*.xml file would contain a
different section of the overall document.
`
	__NoteShape__000000_NoteOnDocument.BodyHTML = ``
	__NoteShape__000000_NoteOnDocument.X = 356.000000
	__NoteShape__000000_NoteOnDocument.Y = 9.000000
	__NoteShape__000000_NoteOnDocument.Width = 669.000000
	__NoteShape__000000_NoteOnDocument.Heigth = 353.000000
	__NoteShape__000000_NoteOnDocument.Matched = false

	// Position values setup
	__Position__000000_Pos_NewDiagram_1_Document.X = 53.000000
	__Position__000000_Pos_NewDiagram_1_Document.Y = 18.000000
	__Position__000000_Pos_NewDiagram_1_Document.Name = `Pos-NewDiagram_1-Document`

	// Position values setup
	__Position__000001_Pos_NewDiagram_1_Node.X = 42.000000
	__Position__000001_Pos_NewDiagram_1_Node.Y = 280.000000
	__Position__000001_Pos_NewDiagram_1_Node.Name = `Pos-NewDiagram_1-Node`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node.X = 429.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node.Y = 94.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node.Name = `Verticle in class diagram NewDiagram_1 in middle between NewDiagram_1-Document and NewDiagram_1-Node`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000000_NewDiagram_1_Document)
	__Classdiagram__000000_NewDiagram_1.GongStructShapes = append(__Classdiagram__000000_NewDiagram_1.GongStructShapes, __GongStructShape__000001_NewDiagram_1_Node)
	__Classdiagram__000000_NewDiagram_1.NoteShapes = append(__Classdiagram__000000_NewDiagram_1.NoteShapes, __NoteShape__000000_NoteOnDocument)
	__GongStructShape__000000_NewDiagram_1_Document.Position = __Position__000000_Pos_NewDiagram_1_Document
	__GongStructShape__000000_NewDiagram_1_Document.Fields = append(__GongStructShape__000000_NewDiagram_1_Document.Fields, __Field__000000_Name)
	__GongStructShape__000000_NewDiagram_1_Document.Links = append(__GongStructShape__000000_NewDiagram_1_Document.Links, __Link__000000_Root)
	__GongStructShape__000001_NewDiagram_1_Node.Position = __Position__000001_Pos_NewDiagram_1_Node
	__Link__000000_Root.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_1_in_middle_between_NewDiagram_1_Document_and_NewDiagram_1_Node
}


