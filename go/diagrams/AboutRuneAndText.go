package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongdocx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_AboutRuneAndText models.StageStruct
var ___dummy__Time_AboutRuneAndText time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_AboutRuneAndText ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_AboutRuneAndText map[string]any = map[string]any{
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
// 	InjectionGateway["AboutRuneAndText"] = AboutRuneAndTextInjection
// }

// AboutRuneAndTextInjection will stage objects of database "AboutRuneAndText"
func AboutRuneAndTextInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_AboutRuneAndText := (&models.Classdiagram{Name: `AboutRuneAndText`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Content := (&models.Field{Name: `Content`}).Stage(stage)
	__Field__000001_Content := (&models.Field{Name: `Content`}).Stage(stage)
	__Field__000002_IsBold := (&models.Field{Name: `IsBold`}).Stage(stage)
	__Field__000003_IsItalic := (&models.Field{Name: `IsItalic`}).Stage(stage)
	__Field__000004_IsStrike := (&models.Field{Name: `IsStrike`}).Stage(stage)
	__Field__000005_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000006_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_AboutRuneAndText_Rune := (&models.GongStructShape{Name: `AboutRuneAndText-Rune`}).Stage(stage)
	__GongStructShape__000001_AboutRuneAndText_RuneProperties := (&models.GongStructShape{Name: `AboutRuneAndText-RuneProperties`}).Stage(stage)
	__GongStructShape__000002_AboutRuneAndText_Text := (&models.GongStructShape{Name: `AboutRuneAndText-Text`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Text := (&models.Link{Name: `Text`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_NoteOnRunProperties := (&models.NoteShape{Name: `NoteOnRunProperties`}).Stage(stage)
	__NoteShape__000001_NoteOnRune := (&models.NoteShape{Name: `NoteOnRune`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_AboutRuneAndText_Rune := (&models.Position{Name: `Pos-AboutRuneAndText-Rune`}).Stage(stage)
	__Position__000001_Pos_AboutRuneAndText_RuneProperties := (&models.Position{Name: `Pos-AboutRuneAndText-RuneProperties`}).Stage(stage)
	__Position__000002_Pos_AboutRuneAndText_Text := (&models.Position{Name: `Pos-AboutRuneAndText-Text`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_AboutRuneAndText_in_middle_between_AboutRuneAndText_Rune_and_AboutRuneAndText_Text := (&models.Vertice{Name: `Verticle in class diagram AboutRuneAndText in middle between AboutRuneAndText-Rune and AboutRuneAndText-Text`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_AboutRuneAndText.Name = `AboutRuneAndText`
	__Classdiagram__000000_AboutRuneAndText.IsInDrawMode = true

	// Field values setup
	__Field__000000_Content.Name = `Content`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RuneProperties.Content]
	__Field__000000_Content.Identifier = `ref_models.RuneProperties.Content`
	__Field__000000_Content.FieldTypeAsString = ``
	__Field__000000_Content.Structname = `RuneProperties`
	__Field__000000_Content.Fieldtypename = `string`

	// Field values setup
	__Field__000001_Content.Name = `Content`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune.Content]
	__Field__000001_Content.Identifier = `ref_models.Rune.Content`
	__Field__000001_Content.FieldTypeAsString = ``
	__Field__000001_Content.Structname = `Rune`
	__Field__000001_Content.Fieldtypename = `string`

	// Field values setup
	__Field__000002_IsBold.Name = `IsBold`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RuneProperties.IsBold]
	__Field__000002_IsBold.Identifier = `ref_models.RuneProperties.IsBold`
	__Field__000002_IsBold.FieldTypeAsString = ``
	__Field__000002_IsBold.Structname = `RuneProperties`
	__Field__000002_IsBold.Fieldtypename = `bool`

	// Field values setup
	__Field__000003_IsItalic.Name = `IsItalic`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RuneProperties.IsItalic]
	__Field__000003_IsItalic.Identifier = `ref_models.RuneProperties.IsItalic`
	__Field__000003_IsItalic.FieldTypeAsString = ``
	__Field__000003_IsItalic.Structname = `RuneProperties`
	__Field__000003_IsItalic.Fieldtypename = `bool`

	// Field values setup
	__Field__000004_IsStrike.Name = `IsStrike`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RuneProperties.IsStrike]
	__Field__000004_IsStrike.Identifier = `ref_models.RuneProperties.IsStrike`
	__Field__000004_IsStrike.FieldTypeAsString = ``
	__Field__000004_IsStrike.Structname = `RuneProperties`
	__Field__000004_IsStrike.Fieldtypename = `bool`

	// Field values setup
	__Field__000005_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RuneProperties.Name]
	__Field__000005_Name.Identifier = `ref_models.RuneProperties.Name`
	__Field__000005_Name.FieldTypeAsString = ``
	__Field__000005_Name.Structname = `RuneProperties`
	__Field__000005_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000006_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune.Name]
	__Field__000006_Name.Identifier = `ref_models.Rune.Name`
	__Field__000006_Name.FieldTypeAsString = ``
	__Field__000006_Name.Structname = `Rune`
	__Field__000006_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_AboutRuneAndText_Rune.Name = `AboutRuneAndText-Rune`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune]
	__GongStructShape__000000_AboutRuneAndText_Rune.Identifier = `ref_models.Rune`
	__GongStructShape__000000_AboutRuneAndText_Rune.ShowNbInstances = true
	__GongStructShape__000000_AboutRuneAndText_Rune.NbInstances = 25
	__GongStructShape__000000_AboutRuneAndText_Rune.Width = 240.000000
	__GongStructShape__000000_AboutRuneAndText_Rune.Heigth = 93.000000
	__GongStructShape__000000_AboutRuneAndText_Rune.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Name = `AboutRuneAndText-RuneProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RuneProperties]
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Identifier = `ref_models.RuneProperties`
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.ShowNbInstances = true
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.NbInstances = 25
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Width = 240.000000
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Heigth = 138.000000
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_AboutRuneAndText_Text.Name = `AboutRuneAndText-Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__GongStructShape__000002_AboutRuneAndText_Text.Identifier = `ref_models.Text`
	__GongStructShape__000002_AboutRuneAndText_Text.ShowNbInstances = false
	__GongStructShape__000002_AboutRuneAndText_Text.NbInstances = 0
	__GongStructShape__000002_AboutRuneAndText_Text.Width = 240.000000
	__GongStructShape__000002_AboutRuneAndText_Text.Heigth = 63.000000
	__GongStructShape__000002_AboutRuneAndText_Text.IsSelected = false

	// Link values setup
	__Link__000000_Text.Name = `Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rune.Text]
	__Link__000000_Text.Identifier = `ref_models.Rune.Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__Link__000000_Text.Fieldtypename = `ref_models.Text`
	__Link__000000_Text.FieldOffsetX = -50.000000
	__Link__000000_Text.FieldOffsetY = -16.000000
	__Link__000000_Text.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_Text.TargetMultiplicityOffsetX = 21.000000
	__Link__000000_Text.TargetMultiplicityOffsetY = -19.000000
	__Link__000000_Text.SourceMultiplicity = models.MANY
	__Link__000000_Text.SourceMultiplicityOffsetX = 21.000000
	__Link__000000_Text.SourceMultiplicityOffsetY = 26.000000
	__Link__000000_Text.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Text.StartRatio = 0.758333
	__Link__000000_Text.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Text.EndRatio = 0.766667
	__Link__000000_Text.CornerOffsetRatio = 1.161290

	// NoteShape values setup
	__NoteShape__000000_NoteOnRunProperties.Name = `NoteOnRunProperties`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnRunProperties]
	__NoteShape__000000_NoteOnRunProperties.Identifier = `ref_models.NoteOnRunProperties`
	__NoteShape__000000_NoteOnRunProperties.Body = `
The "w:rPr" node represents run properties in a Word document's XML structure.
It is found within a "w:r" (run) node in the document.xml file.
-
This node defines the formatting for a specific run of text within a paragraph.
It can include properties like font size, font type, color, highlighting,
bolding, italics, underlining, and more.
-
For example, a "w:rPr" node might contain a "w:sz" element for font size, a
"w:color" element for text color, or "w:b" for bold formatting. The presence of
elements like "w:b" (bold), "w:i" (italic), and "w:u" (underline) indicate that
the formatting is applied, as they are toggled by their presence alone.
-
When parsing a "w:rPr" node, your code should use the information it provides to
apply the appropriate formatting to the text in the run ("w:r") that contains
this "w:rPr" node.
`
	__NoteShape__000000_NoteOnRunProperties.BodyHTML = `<p>The &quot;w:rPr&quot; node represents run properties in a Word document&apos;s XML structure.
It is found within a &quot;w:r&quot; (run) node in the document.xml file.
-
This node defines the formatting for a specific run of text within a paragraph.
It can include properties like font size, font type, color, highlighting,
bolding, italics, underlining, and more.
-
For example, a &quot;w:rPr&quot; node might contain a &quot;w:sz&quot; element for font size, a
&quot;w:color&quot; element for text color, or &quot;w:b&quot; for bold formatting. The presence of
elements like &quot;w:b&quot; (bold), &quot;w:i&quot; (italic), and &quot;w:u&quot; (underline) indicate that
the formatting is applied, as they are toggled by their presence alone.
-
When parsing a &quot;w:rPr&quot; node, your code should use the information it provides to
apply the appropriate formatting to the text in the run (&quot;w:r&quot;) that contains
this &quot;w:rPr&quot; node.
`
	__NoteShape__000000_NoteOnRunProperties.X = 766.000000
	__NoteShape__000000_NoteOnRunProperties.Y = 347.000000
	__NoteShape__000000_NoteOnRunProperties.Width = 595.000000
	__NoteShape__000000_NoteOnRunProperties.Heigth = 300.000000
	__NoteShape__000000_NoteOnRunProperties.Matched = false

	// NoteShape values setup
	__NoteShape__000001_NoteOnRune.Name = `NoteOnRune`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnRune]
	__NoteShape__000001_NoteOnRune.Identifier = `ref_models.NoteOnRune`
	__NoteShape__000001_NoteOnRune.Body = `
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
	__NoteShape__000001_NoteOnRune.BodyHTML = `<p>for <a href="/models#Rune">models.Rune</a>
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
	__NoteShape__000001_NoteOnRune.X = 44.000000
	__NoteShape__000001_NoteOnRune.Y = 324.000000
	__NoteShape__000001_NoteOnRune.Width = 631.000000
	__NoteShape__000001_NoteOnRune.Heigth = 352.000000
	__NoteShape__000001_NoteOnRune.Matched = false

	// Position values setup
	__Position__000000_Pos_AboutRuneAndText_Rune.X = 47.000000
	__Position__000000_Pos_AboutRuneAndText_Rune.Y = 73.000000
	__Position__000000_Pos_AboutRuneAndText_Rune.Name = `Pos-AboutRuneAndText-Rune`

	// Position values setup
	__Position__000001_Pos_AboutRuneAndText_RuneProperties.X = 610.000000
	__Position__000001_Pos_AboutRuneAndText_RuneProperties.Y = 73.000000
	__Position__000001_Pos_AboutRuneAndText_RuneProperties.Name = `Pos-AboutRuneAndText-RuneProperties`

	// Position values setup
	__Position__000002_Pos_AboutRuneAndText_Text.X = 45.000000
	__Position__000002_Pos_AboutRuneAndText_Text.Y = 238.000000
	__Position__000002_Pos_AboutRuneAndText_Text.Name = `Pos-AboutRuneAndText-Text`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_AboutRuneAndText_in_middle_between_AboutRuneAndText_Rune_and_AboutRuneAndText_Text.X = 406.000000
	__Vertice__000000_Verticle_in_class_diagram_AboutRuneAndText_in_middle_between_AboutRuneAndText_Rune_and_AboutRuneAndText_Text.Y = 202.000000
	__Vertice__000000_Verticle_in_class_diagram_AboutRuneAndText_in_middle_between_AboutRuneAndText_Rune_and_AboutRuneAndText_Text.Name = `Verticle in class diagram AboutRuneAndText in middle between AboutRuneAndText-Rune and AboutRuneAndText-Text`

	// Setup of pointers
	__Classdiagram__000000_AboutRuneAndText.GongStructShapes = append(__Classdiagram__000000_AboutRuneAndText.GongStructShapes, __GongStructShape__000000_AboutRuneAndText_Rune)
	__Classdiagram__000000_AboutRuneAndText.GongStructShapes = append(__Classdiagram__000000_AboutRuneAndText.GongStructShapes, __GongStructShape__000001_AboutRuneAndText_RuneProperties)
	__Classdiagram__000000_AboutRuneAndText.GongStructShapes = append(__Classdiagram__000000_AboutRuneAndText.GongStructShapes, __GongStructShape__000002_AboutRuneAndText_Text)
	__Classdiagram__000000_AboutRuneAndText.NoteShapes = append(__Classdiagram__000000_AboutRuneAndText.NoteShapes, __NoteShape__000001_NoteOnRune)
	__Classdiagram__000000_AboutRuneAndText.NoteShapes = append(__Classdiagram__000000_AboutRuneAndText.NoteShapes, __NoteShape__000000_NoteOnRunProperties)
	__GongStructShape__000000_AboutRuneAndText_Rune.Position = __Position__000000_Pos_AboutRuneAndText_Rune
	__GongStructShape__000000_AboutRuneAndText_Rune.Fields = append(__GongStructShape__000000_AboutRuneAndText_Rune.Fields, __Field__000006_Name)
	__GongStructShape__000000_AboutRuneAndText_Rune.Fields = append(__GongStructShape__000000_AboutRuneAndText_Rune.Fields, __Field__000001_Content)
	__GongStructShape__000000_AboutRuneAndText_Rune.Links = append(__GongStructShape__000000_AboutRuneAndText_Rune.Links, __Link__000000_Text)
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Position = __Position__000001_Pos_AboutRuneAndText_RuneProperties
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields = append(__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields, __Field__000005_Name)
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields = append(__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields, __Field__000002_IsBold)
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields = append(__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields, __Field__000004_IsStrike)
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields = append(__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields, __Field__000003_IsItalic)
	__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields = append(__GongStructShape__000001_AboutRuneAndText_RuneProperties.Fields, __Field__000000_Content)
	__GongStructShape__000002_AboutRuneAndText_Text.Position = __Position__000002_Pos_AboutRuneAndText_Text
	__Link__000000_Text.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_AboutRuneAndText_in_middle_between_AboutRuneAndText_Rune_and_AboutRuneAndText_Text
}
