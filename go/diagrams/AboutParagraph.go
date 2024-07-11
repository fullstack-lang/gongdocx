package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdocx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_AboutParagraph models.StageStruct
var ___dummy__Time_AboutParagraph time.Time

// _ point for meta package dummy declaration
var ___dummy__ref_models_AboutParagraph ref_models.StageStruct

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
// but it has to properly reference the _ gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["AboutParagraph"] = _
// }

// _ will stage objects of database "AboutParagraph"
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_AboutParagraph := (&models.Classdiagram{Name: `AboutParagraph`}).Stage(stage)

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
	__GongStructShape__000000_AboutParagraph_Paragraph := (&models.GongStructShape{Name: `AboutParagraph-Paragraph`}).Stage(stage)
	__GongStructShape__000001_AboutParagraph_ParagraphProperties := (&models.GongStructShape{Name: `AboutParagraph-ParagraphProperties`}).Stage(stage)
	__GongStructShape__000002_AboutParagraph_ParagraphStyle := (&models.GongStructShape{Name: `AboutParagraph-ParagraphStyle`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_ParagraphProperties := (&models.Link{Name: `ParagraphProperties`}).Stage(stage)
	__Link__000001_ParagraphStyle := (&models.Link{Name: `ParagraphStyle`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_NoteOnParagraphProperties := (&models.NoteShape{Name: `NoteOnParagraphProperties`}).Stage(stage)
	__NoteShape__000001_NoteOnStyle := (&models.NoteShape{Name: `NoteOnStyle`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink
	__NoteShapeLink__000000_ParagraphProperties := (&models.NoteShapeLink{Name: `ParagraphProperties`}).Stage(stage)
	__NoteShapeLink__000001_ParagraphStyle := (&models.NoteShapeLink{Name: `ParagraphStyle`}).Stage(stage)

	// Declarations of staged instances of Position
	__Position__000000_Pos_AboutParagraph_Paragraph := (&models.Position{Name: `Pos-AboutParagraph-Paragraph`}).Stage(stage)
	__Position__000001_Pos_AboutParagraph_ParagraphProperties := (&models.Position{Name: `Pos-AboutParagraph-ParagraphProperties`}).Stage(stage)
	__Position__000002_Pos_AboutParagraph_ParagraphStyle := (&models.Position{Name: `Pos-AboutParagraph-ParagraphStyle`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_Paragraph_and_AboutParagraph_ParagraphProperties := (&models.Vertice{Name: `Verticle in class diagram AboutParagraph in middle between AboutParagraph-Paragraph and AboutParagraph-ParagraphProperties`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_ParagraphProperties_and_AboutParagraph_ParagraphStyle := (&models.Vertice{Name: `Verticle in class diagram AboutParagraph in middle between AboutParagraph-ParagraphProperties and AboutParagraph-ParagraphStyle`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_AboutParagraph.Name = `AboutParagraph`
	__Classdiagram__000000_AboutParagraph.IsInDrawMode = true

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

	//gong:ident [ref_models.ParagraphStyle.Name]
	__Field__000001_Name.Identifier = `ref_models.ParagraphStyle.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `ParagraphStyle`
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

	//gong:ident [ref_models.ParagraphProperties.Name]
	__Field__000003_Name.Identifier = `ref_models.ParagraphProperties.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `ParagraphProperties`
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
	__GongStructShape__000000_AboutParagraph_Paragraph.Name = `AboutParagraph-Paragraph`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph]
	__GongStructShape__000000_AboutParagraph_Paragraph.Identifier = `ref_models.Paragraph`
	__GongStructShape__000000_AboutParagraph_Paragraph.ShowNbInstances = true
	__GongStructShape__000000_AboutParagraph_Paragraph.NbInstances = 9
	__GongStructShape__000000_AboutParagraph_Paragraph.Width = 353.000000
	__GongStructShape__000000_AboutParagraph_Paragraph.Height = 78.000000
	__GongStructShape__000000_AboutParagraph_Paragraph.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.Name = `AboutParagraph-ParagraphProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphProperties]
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.Identifier = `ref_models.ParagraphProperties`
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.ShowNbInstances = true
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.NbInstances = 9
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.Width = 353.000000
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.Height = 78.000000
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.Name = `AboutParagraph-ParagraphStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphStyle]
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.Identifier = `ref_models.ParagraphStyle`
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.ShowNbInstances = true
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.NbInstances = 2
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.Width = 351.000000
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.Height = 108.000000
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.IsSelected = false

	// Link values setup
	__Link__000000_ParagraphProperties.Name = `ParagraphProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Paragraph.ParagraphProperties]
	__Link__000000_ParagraphProperties.Identifier = `ref_models.Paragraph.ParagraphProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphProperties]
	__Link__000000_ParagraphProperties.Fieldtypename = `ref_models.ParagraphProperties`
	__Link__000000_ParagraphProperties.FieldOffsetX = 10.000000
	__Link__000000_ParagraphProperties.FieldOffsetY = -19.000000
	__Link__000000_ParagraphProperties.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_ParagraphProperties.TargetMultiplicityOffsetX = -41.000000
	__Link__000000_ParagraphProperties.TargetMultiplicityOffsetY = -18.000000
	__Link__000000_ParagraphProperties.SourceMultiplicity = models.MANY
	__Link__000000_ParagraphProperties.SourceMultiplicityOffsetX = -21.000000
	__Link__000000_ParagraphProperties.SourceMultiplicityOffsetY = 26.000000
	__Link__000000_ParagraphProperties.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_ParagraphProperties.StartRatio = 0.430595
	__Link__000000_ParagraphProperties.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_ParagraphProperties.EndRatio = 0.437500
	__Link__000000_ParagraphProperties.CornerOffsetRatio = 1.615385

	// Link values setup
	__Link__000001_ParagraphStyle.Name = `ParagraphStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphProperties.ParagraphStyle]
	__Link__000001_ParagraphStyle.Identifier = `ref_models.ParagraphProperties.ParagraphStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphStyle]
	__Link__000001_ParagraphStyle.Fieldtypename = `ref_models.ParagraphStyle`
	__Link__000001_ParagraphStyle.FieldOffsetX = 24.000000
	__Link__000001_ParagraphStyle.FieldOffsetY = -18.000000
	__Link__000001_ParagraphStyle.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_ParagraphStyle.TargetMultiplicityOffsetX = -42.000000
	__Link__000001_ParagraphStyle.TargetMultiplicityOffsetY = -20.000000
	__Link__000001_ParagraphStyle.SourceMultiplicity = models.MANY
	__Link__000001_ParagraphStyle.SourceMultiplicityOffsetX = -22.000000
	__Link__000001_ParagraphStyle.SourceMultiplicityOffsetY = 27.000000
	__Link__000001_ParagraphStyle.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_ParagraphStyle.StartRatio = 0.422096
	__Link__000001_ParagraphStyle.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_ParagraphStyle.EndRatio = 0.433333
	__Link__000001_ParagraphStyle.CornerOffsetRatio = 1.782051

	// NoteShape values setup
	__NoteShape__000000_NoteOnParagraphProperties.Name = `NoteOnParagraphProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnParagraphProperties]
	__NoteShape__000000_NoteOnParagraphProperties.Identifier = `ref_models.NoteOnParagraphProperties`
	__NoteShape__000000_NoteOnParagraphProperties.Body = `[models.ParagraphProperties]
-
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
	__NoteShape__000000_NoteOnParagraphProperties.BodyHTML = `<p><a href="/models#ParagraphProperties">models.ParagraphProperties</a>
-
The &quot;w:pPr&quot; node represents paragraph properties within a Word document&apos;s XML
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
	__NoteShape__000000_NoteOnParagraphProperties.X = 585.000000
	__NoteShape__000000_NoteOnParagraphProperties.Y = 66.000000
	__NoteShape__000000_NoteOnParagraphProperties.Width = 666.000000
	__NoteShape__000000_NoteOnParagraphProperties.Height = 302.000000
	__NoteShape__000000_NoteOnParagraphProperties.Matched = false

	// NoteShape values setup
	__NoteShape__000001_NoteOnStyle.Name = `NoteOnStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnParagraphStyle]
	__NoteShape__000001_NoteOnStyle.Identifier = `ref_models.NoteOnParagraphStyle`
	__NoteShape__000001_NoteOnStyle.Body = `[models.ParagraphStyle]
-
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
	__NoteShape__000001_NoteOnStyle.BodyHTML = `<p><a href="/models#ParagraphStyle">models.ParagraphStyle</a>
-
The &quot;w:pStyle&quot; element, found within the paragraph properties (&quot;w:pPr&quot;) node in
a Word document&apos;s XML structure, defines the paragraph style for a given
paragraph (&quot;w:p&quot;).
-
The &quot;w:pStyle&quot; element includes an attribute &quot;w:val&quot; that references the ID of
the style being applied to the paragraph. This style ID correlates with the
styles defined in the styles.xml part of the .docx package.
-
A style in Word includes a predefined set of formatting instructions. It can
control multiple aspects of the paragraph&apos;s appearance, including alignment,
spacing, font, size, color, and more.
-
By using styles, a document can maintain a consistent look and feel, and
changing the style in one place will automatically update all paragraphs that
reference that style.
-
When parsing a &quot;w:pStyle&quot; node, your code should map the style ID to the
corresponding style in the styles.xml file and apply the associated formatting
to the paragraph.
`
	__NoteShape__000001_NoteOnStyle.X = 585.000000
	__NoteShape__000001_NoteOnStyle.Y = 409.000000
	__NoteShape__000001_NoteOnStyle.Width = 668.000000
	__NoteShape__000001_NoteOnStyle.Height = 368.000000
	__NoteShape__000001_NoteOnStyle.Matched = false

	// NoteShapeLink values setup
	__NoteShapeLink__000000_ParagraphProperties.Name = `ParagraphProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphProperties]
	__NoteShapeLink__000000_ParagraphProperties.Identifier = `ref_models.ParagraphProperties`
	__NoteShapeLink__000000_ParagraphProperties.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// NoteShapeLink values setup
	__NoteShapeLink__000001_ParagraphStyle.Name = `ParagraphStyle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ParagraphStyle]
	__NoteShapeLink__000001_ParagraphStyle.Identifier = `ref_models.ParagraphStyle`
	__NoteShapeLink__000001_ParagraphStyle.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// Position values setup
	__Position__000000_Pos_AboutParagraph_Paragraph.X = 65.000000
	__Position__000000_Pos_AboutParagraph_Paragraph.Y = 71.000000
	__Position__000000_Pos_AboutParagraph_Paragraph.Name = `Pos-AboutParagraph-Paragraph`

	// Position values setup
	__Position__000001_Pos_AboutParagraph_ParagraphProperties.X = 67.000000
	__Position__000001_Pos_AboutParagraph_ParagraphProperties.Y = 257.000000
	__Position__000001_Pos_AboutParagraph_ParagraphProperties.Name = `Pos-AboutParagraph-ParagraphProperties`

	// Position values setup
	__Position__000002_Pos_AboutParagraph_ParagraphStyle.X = 67.000000
	__Position__000002_Pos_AboutParagraph_ParagraphStyle.Y = 467.000000
	__Position__000002_Pos_AboutParagraph_ParagraphStyle.Name = `Pos-AboutParagraph-ParagraphStyle`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_Paragraph_and_AboutParagraph_ParagraphProperties.X = 609.000000
	__Vertice__000000_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_Paragraph_and_AboutParagraph_ParagraphProperties.Y = 109.000000
	__Vertice__000000_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_Paragraph_and_AboutParagraph_ParagraphProperties.Name = `Verticle in class diagram AboutParagraph in middle between AboutParagraph-Paragraph and AboutParagraph-ParagraphProperties`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_ParagraphProperties_and_AboutParagraph_ParagraphStyle.X = 427.000000
	__Vertice__000001_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_ParagraphProperties_and_AboutParagraph_ParagraphStyle.Y = 401.000000
	__Vertice__000001_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_ParagraphProperties_and_AboutParagraph_ParagraphStyle.Name = `Verticle in class diagram AboutParagraph in middle between AboutParagraph-ParagraphProperties and AboutParagraph-ParagraphStyle`

	// Setup of pointers
	__Classdiagram__000000_AboutParagraph.GongStructShapes = append(__Classdiagram__000000_AboutParagraph.GongStructShapes, __GongStructShape__000000_AboutParagraph_Paragraph)
	__Classdiagram__000000_AboutParagraph.GongStructShapes = append(__Classdiagram__000000_AboutParagraph.GongStructShapes, __GongStructShape__000001_AboutParagraph_ParagraphProperties)
	__Classdiagram__000000_AboutParagraph.GongStructShapes = append(__Classdiagram__000000_AboutParagraph.GongStructShapes, __GongStructShape__000002_AboutParagraph_ParagraphStyle)
	__Classdiagram__000000_AboutParagraph.NoteShapes = append(__Classdiagram__000000_AboutParagraph.NoteShapes, __NoteShape__000000_NoteOnParagraphProperties)
	__Classdiagram__000000_AboutParagraph.NoteShapes = append(__Classdiagram__000000_AboutParagraph.NoteShapes, __NoteShape__000001_NoteOnStyle)
	__GongStructShape__000000_AboutParagraph_Paragraph.Position = __Position__000000_Pos_AboutParagraph_Paragraph
	__GongStructShape__000000_AboutParagraph_Paragraph.Fields = append(__GongStructShape__000000_AboutParagraph_Paragraph.Fields, __Field__000002_Name)
	__GongStructShape__000000_AboutParagraph_Paragraph.Links = append(__GongStructShape__000000_AboutParagraph_Paragraph.Links, __Link__000000_ParagraphProperties)
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.Position = __Position__000001_Pos_AboutParagraph_ParagraphProperties
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.Fields = append(__GongStructShape__000001_AboutParagraph_ParagraphProperties.Fields, __Field__000003_Name)
	__GongStructShape__000001_AboutParagraph_ParagraphProperties.Links = append(__GongStructShape__000001_AboutParagraph_ParagraphProperties.Links, __Link__000001_ParagraphStyle)
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.Position = __Position__000002_Pos_AboutParagraph_ParagraphStyle
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.Fields = append(__GongStructShape__000002_AboutParagraph_ParagraphStyle.Fields, __Field__000001_Name)
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.Fields = append(__GongStructShape__000002_AboutParagraph_ParagraphStyle.Fields, __Field__000000_Content)
	__GongStructShape__000002_AboutParagraph_ParagraphStyle.Fields = append(__GongStructShape__000002_AboutParagraph_ParagraphStyle.Fields, __Field__000004_ValAttr)
	__Link__000000_ParagraphProperties.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_Paragraph_and_AboutParagraph_ParagraphProperties
	__Link__000001_ParagraphStyle.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_AboutParagraph_in_middle_between_AboutParagraph_ParagraphProperties_and_AboutParagraph_ParagraphStyle
	__NoteShape__000000_NoteOnParagraphProperties.NoteShapeLinks = append(__NoteShape__000000_NoteOnParagraphProperties.NoteShapeLinks, __NoteShapeLink__000000_ParagraphProperties)
	__NoteShape__000001_NoteOnStyle.NoteShapeLinks = append(__NoteShape__000001_NoteOnStyle.NoteShapeLinks, __NoteShapeLink__000001_ParagraphStyle)
}
