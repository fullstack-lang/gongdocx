// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Bodys           map[*Body]any
	Bodys_mapString map[string]*Body

	// insertion point for slice of pointers maps
	Body_Paragraphs_reverseMap map[*Paragraph]*Body

	Body_Tables_reverseMap map[*Table]*Body

	OnAfterBodyCreateCallback OnAfterCreateInterface[Body]
	OnAfterBodyUpdateCallback OnAfterUpdateInterface[Body]
	OnAfterBodyDeleteCallback OnAfterDeleteInterface[Body]
	OnAfterBodyReadCallback   OnAfterReadInterface[Body]

	Documents           map[*Document]any
	Documents_mapString map[string]*Document

	// insertion point for slice of pointers maps
	OnAfterDocumentCreateCallback OnAfterCreateInterface[Document]
	OnAfterDocumentUpdateCallback OnAfterUpdateInterface[Document]
	OnAfterDocumentDeleteCallback OnAfterDeleteInterface[Document]
	OnAfterDocumentReadCallback   OnAfterReadInterface[Document]

	Docxs           map[*Docx]any
	Docxs_mapString map[string]*Docx

	// insertion point for slice of pointers maps
	Docx_Files_reverseMap map[*File]*Docx

	OnAfterDocxCreateCallback OnAfterCreateInterface[Docx]
	OnAfterDocxUpdateCallback OnAfterUpdateInterface[Docx]
	OnAfterDocxDeleteCallback OnAfterDeleteInterface[Docx]
	OnAfterDocxReadCallback   OnAfterReadInterface[Docx]

	Files           map[*File]any
	Files_mapString map[string]*File

	// insertion point for slice of pointers maps
	OnAfterFileCreateCallback OnAfterCreateInterface[File]
	OnAfterFileUpdateCallback OnAfterUpdateInterface[File]
	OnAfterFileDeleteCallback OnAfterDeleteInterface[File]
	OnAfterFileReadCallback   OnAfterReadInterface[File]

	Nodes           map[*Node]any
	Nodes_mapString map[string]*Node

	// insertion point for slice of pointers maps
	Node_Nodes_reverseMap map[*Node]*Node

	OnAfterNodeCreateCallback OnAfterCreateInterface[Node]
	OnAfterNodeUpdateCallback OnAfterUpdateInterface[Node]
	OnAfterNodeDeleteCallback OnAfterDeleteInterface[Node]
	OnAfterNodeReadCallback   OnAfterReadInterface[Node]

	Paragraphs           map[*Paragraph]any
	Paragraphs_mapString map[string]*Paragraph

	// insertion point for slice of pointers maps
	Paragraph_Runes_reverseMap map[*Rune]*Paragraph

	OnAfterParagraphCreateCallback OnAfterCreateInterface[Paragraph]
	OnAfterParagraphUpdateCallback OnAfterUpdateInterface[Paragraph]
	OnAfterParagraphDeleteCallback OnAfterDeleteInterface[Paragraph]
	OnAfterParagraphReadCallback   OnAfterReadInterface[Paragraph]

	ParagraphPropertiess           map[*ParagraphProperties]any
	ParagraphPropertiess_mapString map[string]*ParagraphProperties

	// insertion point for slice of pointers maps
	OnAfterParagraphPropertiesCreateCallback OnAfterCreateInterface[ParagraphProperties]
	OnAfterParagraphPropertiesUpdateCallback OnAfterUpdateInterface[ParagraphProperties]
	OnAfterParagraphPropertiesDeleteCallback OnAfterDeleteInterface[ParagraphProperties]
	OnAfterParagraphPropertiesReadCallback   OnAfterReadInterface[ParagraphProperties]

	ParagraphStyles           map[*ParagraphStyle]any
	ParagraphStyles_mapString map[string]*ParagraphStyle

	// insertion point for slice of pointers maps
	OnAfterParagraphStyleCreateCallback OnAfterCreateInterface[ParagraphStyle]
	OnAfterParagraphStyleUpdateCallback OnAfterUpdateInterface[ParagraphStyle]
	OnAfterParagraphStyleDeleteCallback OnAfterDeleteInterface[ParagraphStyle]
	OnAfterParagraphStyleReadCallback   OnAfterReadInterface[ParagraphStyle]

	Runes           map[*Rune]any
	Runes_mapString map[string]*Rune

	// insertion point for slice of pointers maps
	OnAfterRuneCreateCallback OnAfterCreateInterface[Rune]
	OnAfterRuneUpdateCallback OnAfterUpdateInterface[Rune]
	OnAfterRuneDeleteCallback OnAfterDeleteInterface[Rune]
	OnAfterRuneReadCallback   OnAfterReadInterface[Rune]

	RunePropertiess           map[*RuneProperties]any
	RunePropertiess_mapString map[string]*RuneProperties

	// insertion point for slice of pointers maps
	OnAfterRunePropertiesCreateCallback OnAfterCreateInterface[RuneProperties]
	OnAfterRunePropertiesUpdateCallback OnAfterUpdateInterface[RuneProperties]
	OnAfterRunePropertiesDeleteCallback OnAfterDeleteInterface[RuneProperties]
	OnAfterRunePropertiesReadCallback   OnAfterReadInterface[RuneProperties]

	Tables           map[*Table]any
	Tables_mapString map[string]*Table

	// insertion point for slice of pointers maps
	Table_TableRows_reverseMap map[*TableRow]*Table

	OnAfterTableCreateCallback OnAfterCreateInterface[Table]
	OnAfterTableUpdateCallback OnAfterUpdateInterface[Table]
	OnAfterTableDeleteCallback OnAfterDeleteInterface[Table]
	OnAfterTableReadCallback   OnAfterReadInterface[Table]

	TableColumns           map[*TableColumn]any
	TableColumns_mapString map[string]*TableColumn

	// insertion point for slice of pointers maps
	TableColumn_Paragraphs_reverseMap map[*Paragraph]*TableColumn

	OnAfterTableColumnCreateCallback OnAfterCreateInterface[TableColumn]
	OnAfterTableColumnUpdateCallback OnAfterUpdateInterface[TableColumn]
	OnAfterTableColumnDeleteCallback OnAfterDeleteInterface[TableColumn]
	OnAfterTableColumnReadCallback   OnAfterReadInterface[TableColumn]

	TablePropertiess           map[*TableProperties]any
	TablePropertiess_mapString map[string]*TableProperties

	// insertion point for slice of pointers maps
	OnAfterTablePropertiesCreateCallback OnAfterCreateInterface[TableProperties]
	OnAfterTablePropertiesUpdateCallback OnAfterUpdateInterface[TableProperties]
	OnAfterTablePropertiesDeleteCallback OnAfterDeleteInterface[TableProperties]
	OnAfterTablePropertiesReadCallback   OnAfterReadInterface[TableProperties]

	TableRows           map[*TableRow]any
	TableRows_mapString map[string]*TableRow

	// insertion point for slice of pointers maps
	TableRow_TableColumns_reverseMap map[*TableColumn]*TableRow

	OnAfterTableRowCreateCallback OnAfterCreateInterface[TableRow]
	OnAfterTableRowUpdateCallback OnAfterUpdateInterface[TableRow]
	OnAfterTableRowDeleteCallback OnAfterDeleteInterface[TableRow]
	OnAfterTableRowReadCallback   OnAfterReadInterface[TableRow]

	TableStyles           map[*TableStyle]any
	TableStyles_mapString map[string]*TableStyle

	// insertion point for slice of pointers maps
	OnAfterTableStyleCreateCallback OnAfterCreateInterface[TableStyle]
	OnAfterTableStyleUpdateCallback OnAfterUpdateInterface[TableStyle]
	OnAfterTableStyleDeleteCallback OnAfterDeleteInterface[TableStyle]
	OnAfterTableStyleReadCallback   OnAfterReadInterface[TableStyle]

	Texts           map[*Text]any
	Texts_mapString map[string]*Text

	// insertion point for slice of pointers maps
	OnAfterTextCreateCallback OnAfterCreateInterface[Text]
	OnAfterTextUpdateCallback OnAfterUpdateInterface[Text]
	OnAfterTextDeleteCallback OnAfterDeleteInterface[Text]
	OnAfterTextReadCallback   OnAfterReadInterface[Text]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongdocx/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitBody(body *Body)
	CheckoutBody(body *Body)
	CommitDocument(document *Document)
	CheckoutDocument(document *Document)
	CommitDocx(docx *Docx)
	CheckoutDocx(docx *Docx)
	CommitFile(file *File)
	CheckoutFile(file *File)
	CommitNode(node *Node)
	CheckoutNode(node *Node)
	CommitParagraph(paragraph *Paragraph)
	CheckoutParagraph(paragraph *Paragraph)
	CommitParagraphProperties(paragraphproperties *ParagraphProperties)
	CheckoutParagraphProperties(paragraphproperties *ParagraphProperties)
	CommitParagraphStyle(paragraphstyle *ParagraphStyle)
	CheckoutParagraphStyle(paragraphstyle *ParagraphStyle)
	CommitRune(rune *Rune)
	CheckoutRune(rune *Rune)
	CommitRuneProperties(runeproperties *RuneProperties)
	CheckoutRuneProperties(runeproperties *RuneProperties)
	CommitTable(table *Table)
	CheckoutTable(table *Table)
	CommitTableColumn(tablecolumn *TableColumn)
	CheckoutTableColumn(tablecolumn *TableColumn)
	CommitTableProperties(tableproperties *TableProperties)
	CheckoutTableProperties(tableproperties *TableProperties)
	CommitTableRow(tablerow *TableRow)
	CheckoutTableRow(tablerow *TableRow)
	CommitTableStyle(tablestyle *TableStyle)
	CheckoutTableStyle(tablestyle *TableStyle)
	CommitText(text *Text)
	CheckoutText(text *Text)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Bodys:           make(map[*Body]any),
		Bodys_mapString: make(map[string]*Body),

		Documents:           make(map[*Document]any),
		Documents_mapString: make(map[string]*Document),

		Docxs:           make(map[*Docx]any),
		Docxs_mapString: make(map[string]*Docx),

		Files:           make(map[*File]any),
		Files_mapString: make(map[string]*File),

		Nodes:           make(map[*Node]any),
		Nodes_mapString: make(map[string]*Node),

		Paragraphs:           make(map[*Paragraph]any),
		Paragraphs_mapString: make(map[string]*Paragraph),

		ParagraphPropertiess:           make(map[*ParagraphProperties]any),
		ParagraphPropertiess_mapString: make(map[string]*ParagraphProperties),

		ParagraphStyles:           make(map[*ParagraphStyle]any),
		ParagraphStyles_mapString: make(map[string]*ParagraphStyle),

		Runes:           make(map[*Rune]any),
		Runes_mapString: make(map[string]*Rune),

		RunePropertiess:           make(map[*RuneProperties]any),
		RunePropertiess_mapString: make(map[string]*RuneProperties),

		Tables:           make(map[*Table]any),
		Tables_mapString: make(map[string]*Table),

		TableColumns:           make(map[*TableColumn]any),
		TableColumns_mapString: make(map[string]*TableColumn),

		TablePropertiess:           make(map[*TableProperties]any),
		TablePropertiess_mapString: make(map[string]*TableProperties),

		TableRows:           make(map[*TableRow]any),
		TableRows_mapString: make(map[string]*TableRow),

		TableStyles:           make(map[*TableStyle]any),
		TableStyles_mapString: make(map[string]*TableStyle),

		Texts:           make(map[*Text]any),
		Texts_mapString: make(map[string]*Text),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Body"] = len(stage.Bodys)
	stage.Map_GongStructName_InstancesNb["Document"] = len(stage.Documents)
	stage.Map_GongStructName_InstancesNb["Docx"] = len(stage.Docxs)
	stage.Map_GongStructName_InstancesNb["File"] = len(stage.Files)
	stage.Map_GongStructName_InstancesNb["Node"] = len(stage.Nodes)
	stage.Map_GongStructName_InstancesNb["Paragraph"] = len(stage.Paragraphs)
	stage.Map_GongStructName_InstancesNb["ParagraphProperties"] = len(stage.ParagraphPropertiess)
	stage.Map_GongStructName_InstancesNb["ParagraphStyle"] = len(stage.ParagraphStyles)
	stage.Map_GongStructName_InstancesNb["Rune"] = len(stage.Runes)
	stage.Map_GongStructName_InstancesNb["RuneProperties"] = len(stage.RunePropertiess)
	stage.Map_GongStructName_InstancesNb["Table"] = len(stage.Tables)
	stage.Map_GongStructName_InstancesNb["TableColumn"] = len(stage.TableColumns)
	stage.Map_GongStructName_InstancesNb["TableProperties"] = len(stage.TablePropertiess)
	stage.Map_GongStructName_InstancesNb["TableRow"] = len(stage.TableRows)
	stage.Map_GongStructName_InstancesNb["TableStyle"] = len(stage.TableStyles)
	stage.Map_GongStructName_InstancesNb["Text"] = len(stage.Texts)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Body"] = len(stage.Bodys)
	stage.Map_GongStructName_InstancesNb["Document"] = len(stage.Documents)
	stage.Map_GongStructName_InstancesNb["Docx"] = len(stage.Docxs)
	stage.Map_GongStructName_InstancesNb["File"] = len(stage.Files)
	stage.Map_GongStructName_InstancesNb["Node"] = len(stage.Nodes)
	stage.Map_GongStructName_InstancesNb["Paragraph"] = len(stage.Paragraphs)
	stage.Map_GongStructName_InstancesNb["ParagraphProperties"] = len(stage.ParagraphPropertiess)
	stage.Map_GongStructName_InstancesNb["ParagraphStyle"] = len(stage.ParagraphStyles)
	stage.Map_GongStructName_InstancesNb["Rune"] = len(stage.Runes)
	stage.Map_GongStructName_InstancesNb["RuneProperties"] = len(stage.RunePropertiess)
	stage.Map_GongStructName_InstancesNb["Table"] = len(stage.Tables)
	stage.Map_GongStructName_InstancesNb["TableColumn"] = len(stage.TableColumns)
	stage.Map_GongStructName_InstancesNb["TableProperties"] = len(stage.TablePropertiess)
	stage.Map_GongStructName_InstancesNb["TableRow"] = len(stage.TableRows)
	stage.Map_GongStructName_InstancesNb["TableStyle"] = len(stage.TableStyles)
	stage.Map_GongStructName_InstancesNb["Text"] = len(stage.Texts)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts body to the model stage
func (body *Body) Stage(stage *StageStruct) *Body {
	stage.Bodys[body] = __member
	stage.Bodys_mapString[body.Name] = body

	return body
}

// Unstage removes body off the model stage
func (body *Body) Unstage(stage *StageStruct) *Body {
	delete(stage.Bodys, body)
	delete(stage.Bodys_mapString, body.Name)
	return body
}

// UnstageVoid removes body off the model stage
func (body *Body) UnstageVoid(stage *StageStruct) {
	delete(stage.Bodys, body)
	delete(stage.Bodys_mapString, body.Name)
}

// commit body to the back repo (if it is already staged)
func (body *Body) Commit(stage *StageStruct) *Body {
	if _, ok := stage.Bodys[body]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBody(body)
		}
	}
	return body
}

func (body *Body) CommitVoid(stage *StageStruct) {
	body.Commit(stage)
}

// Checkout body to the back repo (if it is already staged)
func (body *Body) Checkout(stage *StageStruct) *Body {
	if _, ok := stage.Bodys[body]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBody(body)
		}
	}
	return body
}

// for satisfaction of GongStruct interface
func (body *Body) GetName() (res string) {
	return body.Name
}

// Stage puts document to the model stage
func (document *Document) Stage(stage *StageStruct) *Document {
	stage.Documents[document] = __member
	stage.Documents_mapString[document.Name] = document

	return document
}

// Unstage removes document off the model stage
func (document *Document) Unstage(stage *StageStruct) *Document {
	delete(stage.Documents, document)
	delete(stage.Documents_mapString, document.Name)
	return document
}

// UnstageVoid removes document off the model stage
func (document *Document) UnstageVoid(stage *StageStruct) {
	delete(stage.Documents, document)
	delete(stage.Documents_mapString, document.Name)
}

// commit document to the back repo (if it is already staged)
func (document *Document) Commit(stage *StageStruct) *Document {
	if _, ok := stage.Documents[document]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocument(document)
		}
	}
	return document
}

func (document *Document) CommitVoid(stage *StageStruct) {
	document.Commit(stage)
}

// Checkout document to the back repo (if it is already staged)
func (document *Document) Checkout(stage *StageStruct) *Document {
	if _, ok := stage.Documents[document]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocument(document)
		}
	}
	return document
}

// for satisfaction of GongStruct interface
func (document *Document) GetName() (res string) {
	return document.Name
}

// Stage puts docx to the model stage
func (docx *Docx) Stage(stage *StageStruct) *Docx {
	stage.Docxs[docx] = __member
	stage.Docxs_mapString[docx.Name] = docx

	return docx
}

// Unstage removes docx off the model stage
func (docx *Docx) Unstage(stage *StageStruct) *Docx {
	delete(stage.Docxs, docx)
	delete(stage.Docxs_mapString, docx.Name)
	return docx
}

// UnstageVoid removes docx off the model stage
func (docx *Docx) UnstageVoid(stage *StageStruct) {
	delete(stage.Docxs, docx)
	delete(stage.Docxs_mapString, docx.Name)
}

// commit docx to the back repo (if it is already staged)
func (docx *Docx) Commit(stage *StageStruct) *Docx {
	if _, ok := stage.Docxs[docx]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocx(docx)
		}
	}
	return docx
}

func (docx *Docx) CommitVoid(stage *StageStruct) {
	docx.Commit(stage)
}

// Checkout docx to the back repo (if it is already staged)
func (docx *Docx) Checkout(stage *StageStruct) *Docx {
	if _, ok := stage.Docxs[docx]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocx(docx)
		}
	}
	return docx
}

// for satisfaction of GongStruct interface
func (docx *Docx) GetName() (res string) {
	return docx.Name
}

// Stage puts file to the model stage
func (file *File) Stage(stage *StageStruct) *File {
	stage.Files[file] = __member
	stage.Files_mapString[file.Name] = file

	return file
}

// Unstage removes file off the model stage
func (file *File) Unstage(stage *StageStruct) *File {
	delete(stage.Files, file)
	delete(stage.Files_mapString, file.Name)
	return file
}

// UnstageVoid removes file off the model stage
func (file *File) UnstageVoid(stage *StageStruct) {
	delete(stage.Files, file)
	delete(stage.Files_mapString, file.Name)
}

// commit file to the back repo (if it is already staged)
func (file *File) Commit(stage *StageStruct) *File {
	if _, ok := stage.Files[file]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFile(file)
		}
	}
	return file
}

func (file *File) CommitVoid(stage *StageStruct) {
	file.Commit(stage)
}

// Checkout file to the back repo (if it is already staged)
func (file *File) Checkout(stage *StageStruct) *File {
	if _, ok := stage.Files[file]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFile(file)
		}
	}
	return file
}

// for satisfaction of GongStruct interface
func (file *File) GetName() (res string) {
	return file.Name
}

// Stage puts node to the model stage
func (node *Node) Stage(stage *StageStruct) *Node {
	stage.Nodes[node] = __member
	stage.Nodes_mapString[node.Name] = node

	return node
}

// Unstage removes node off the model stage
func (node *Node) Unstage(stage *StageStruct) *Node {
	delete(stage.Nodes, node)
	delete(stage.Nodes_mapString, node.Name)
	return node
}

// UnstageVoid removes node off the model stage
func (node *Node) UnstageVoid(stage *StageStruct) {
	delete(stage.Nodes, node)
	delete(stage.Nodes_mapString, node.Name)
}

// commit node to the back repo (if it is already staged)
func (node *Node) Commit(stage *StageStruct) *Node {
	if _, ok := stage.Nodes[node]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNode(node)
		}
	}
	return node
}

func (node *Node) CommitVoid(stage *StageStruct) {
	node.Commit(stage)
}

// Checkout node to the back repo (if it is already staged)
func (node *Node) Checkout(stage *StageStruct) *Node {
	if _, ok := stage.Nodes[node]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNode(node)
		}
	}
	return node
}

// for satisfaction of GongStruct interface
func (node *Node) GetName() (res string) {
	return node.Name
}

// Stage puts paragraph to the model stage
func (paragraph *Paragraph) Stage(stage *StageStruct) *Paragraph {
	stage.Paragraphs[paragraph] = __member
	stage.Paragraphs_mapString[paragraph.Name] = paragraph

	return paragraph
}

// Unstage removes paragraph off the model stage
func (paragraph *Paragraph) Unstage(stage *StageStruct) *Paragraph {
	delete(stage.Paragraphs, paragraph)
	delete(stage.Paragraphs_mapString, paragraph.Name)
	return paragraph
}

// UnstageVoid removes paragraph off the model stage
func (paragraph *Paragraph) UnstageVoid(stage *StageStruct) {
	delete(stage.Paragraphs, paragraph)
	delete(stage.Paragraphs_mapString, paragraph.Name)
}

// commit paragraph to the back repo (if it is already staged)
func (paragraph *Paragraph) Commit(stage *StageStruct) *Paragraph {
	if _, ok := stage.Paragraphs[paragraph]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParagraph(paragraph)
		}
	}
	return paragraph
}

func (paragraph *Paragraph) CommitVoid(stage *StageStruct) {
	paragraph.Commit(stage)
}

// Checkout paragraph to the back repo (if it is already staged)
func (paragraph *Paragraph) Checkout(stage *StageStruct) *Paragraph {
	if _, ok := stage.Paragraphs[paragraph]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParagraph(paragraph)
		}
	}
	return paragraph
}

// for satisfaction of GongStruct interface
func (paragraph *Paragraph) GetName() (res string) {
	return paragraph.Name
}

// Stage puts paragraphproperties to the model stage
func (paragraphproperties *ParagraphProperties) Stage(stage *StageStruct) *ParagraphProperties {
	stage.ParagraphPropertiess[paragraphproperties] = __member
	stage.ParagraphPropertiess_mapString[paragraphproperties.Name] = paragraphproperties

	return paragraphproperties
}

// Unstage removes paragraphproperties off the model stage
func (paragraphproperties *ParagraphProperties) Unstage(stage *StageStruct) *ParagraphProperties {
	delete(stage.ParagraphPropertiess, paragraphproperties)
	delete(stage.ParagraphPropertiess_mapString, paragraphproperties.Name)
	return paragraphproperties
}

// UnstageVoid removes paragraphproperties off the model stage
func (paragraphproperties *ParagraphProperties) UnstageVoid(stage *StageStruct) {
	delete(stage.ParagraphPropertiess, paragraphproperties)
	delete(stage.ParagraphPropertiess_mapString, paragraphproperties.Name)
}

// commit paragraphproperties to the back repo (if it is already staged)
func (paragraphproperties *ParagraphProperties) Commit(stage *StageStruct) *ParagraphProperties {
	if _, ok := stage.ParagraphPropertiess[paragraphproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParagraphProperties(paragraphproperties)
		}
	}
	return paragraphproperties
}

func (paragraphproperties *ParagraphProperties) CommitVoid(stage *StageStruct) {
	paragraphproperties.Commit(stage)
}

// Checkout paragraphproperties to the back repo (if it is already staged)
func (paragraphproperties *ParagraphProperties) Checkout(stage *StageStruct) *ParagraphProperties {
	if _, ok := stage.ParagraphPropertiess[paragraphproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParagraphProperties(paragraphproperties)
		}
	}
	return paragraphproperties
}

// for satisfaction of GongStruct interface
func (paragraphproperties *ParagraphProperties) GetName() (res string) {
	return paragraphproperties.Name
}

// Stage puts paragraphstyle to the model stage
func (paragraphstyle *ParagraphStyle) Stage(stage *StageStruct) *ParagraphStyle {
	stage.ParagraphStyles[paragraphstyle] = __member
	stage.ParagraphStyles_mapString[paragraphstyle.Name] = paragraphstyle

	return paragraphstyle
}

// Unstage removes paragraphstyle off the model stage
func (paragraphstyle *ParagraphStyle) Unstage(stage *StageStruct) *ParagraphStyle {
	delete(stage.ParagraphStyles, paragraphstyle)
	delete(stage.ParagraphStyles_mapString, paragraphstyle.Name)
	return paragraphstyle
}

// UnstageVoid removes paragraphstyle off the model stage
func (paragraphstyle *ParagraphStyle) UnstageVoid(stage *StageStruct) {
	delete(stage.ParagraphStyles, paragraphstyle)
	delete(stage.ParagraphStyles_mapString, paragraphstyle.Name)
}

// commit paragraphstyle to the back repo (if it is already staged)
func (paragraphstyle *ParagraphStyle) Commit(stage *StageStruct) *ParagraphStyle {
	if _, ok := stage.ParagraphStyles[paragraphstyle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParagraphStyle(paragraphstyle)
		}
	}
	return paragraphstyle
}

func (paragraphstyle *ParagraphStyle) CommitVoid(stage *StageStruct) {
	paragraphstyle.Commit(stage)
}

// Checkout paragraphstyle to the back repo (if it is already staged)
func (paragraphstyle *ParagraphStyle) Checkout(stage *StageStruct) *ParagraphStyle {
	if _, ok := stage.ParagraphStyles[paragraphstyle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParagraphStyle(paragraphstyle)
		}
	}
	return paragraphstyle
}

// for satisfaction of GongStruct interface
func (paragraphstyle *ParagraphStyle) GetName() (res string) {
	return paragraphstyle.Name
}

// Stage puts rune to the model stage
func (rune *Rune) Stage(stage *StageStruct) *Rune {
	stage.Runes[rune] = __member
	stage.Runes_mapString[rune.Name] = rune

	return rune
}

// Unstage removes rune off the model stage
func (rune *Rune) Unstage(stage *StageStruct) *Rune {
	delete(stage.Runes, rune)
	delete(stage.Runes_mapString, rune.Name)
	return rune
}

// UnstageVoid removes rune off the model stage
func (rune *Rune) UnstageVoid(stage *StageStruct) {
	delete(stage.Runes, rune)
	delete(stage.Runes_mapString, rune.Name)
}

// commit rune to the back repo (if it is already staged)
func (rune *Rune) Commit(stage *StageStruct) *Rune {
	if _, ok := stage.Runes[rune]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRune(rune)
		}
	}
	return rune
}

func (rune *Rune) CommitVoid(stage *StageStruct) {
	rune.Commit(stage)
}

// Checkout rune to the back repo (if it is already staged)
func (rune *Rune) Checkout(stage *StageStruct) *Rune {
	if _, ok := stage.Runes[rune]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRune(rune)
		}
	}
	return rune
}

// for satisfaction of GongStruct interface
func (rune *Rune) GetName() (res string) {
	return rune.Name
}

// Stage puts runeproperties to the model stage
func (runeproperties *RuneProperties) Stage(stage *StageStruct) *RuneProperties {
	stage.RunePropertiess[runeproperties] = __member
	stage.RunePropertiess_mapString[runeproperties.Name] = runeproperties

	return runeproperties
}

// Unstage removes runeproperties off the model stage
func (runeproperties *RuneProperties) Unstage(stage *StageStruct) *RuneProperties {
	delete(stage.RunePropertiess, runeproperties)
	delete(stage.RunePropertiess_mapString, runeproperties.Name)
	return runeproperties
}

// UnstageVoid removes runeproperties off the model stage
func (runeproperties *RuneProperties) UnstageVoid(stage *StageStruct) {
	delete(stage.RunePropertiess, runeproperties)
	delete(stage.RunePropertiess_mapString, runeproperties.Name)
}

// commit runeproperties to the back repo (if it is already staged)
func (runeproperties *RuneProperties) Commit(stage *StageStruct) *RuneProperties {
	if _, ok := stage.RunePropertiess[runeproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRuneProperties(runeproperties)
		}
	}
	return runeproperties
}

func (runeproperties *RuneProperties) CommitVoid(stage *StageStruct) {
	runeproperties.Commit(stage)
}

// Checkout runeproperties to the back repo (if it is already staged)
func (runeproperties *RuneProperties) Checkout(stage *StageStruct) *RuneProperties {
	if _, ok := stage.RunePropertiess[runeproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRuneProperties(runeproperties)
		}
	}
	return runeproperties
}

// for satisfaction of GongStruct interface
func (runeproperties *RuneProperties) GetName() (res string) {
	return runeproperties.Name
}

// Stage puts table to the model stage
func (table *Table) Stage(stage *StageStruct) *Table {
	stage.Tables[table] = __member
	stage.Tables_mapString[table.Name] = table

	return table
}

// Unstage removes table off the model stage
func (table *Table) Unstage(stage *StageStruct) *Table {
	delete(stage.Tables, table)
	delete(stage.Tables_mapString, table.Name)
	return table
}

// UnstageVoid removes table off the model stage
func (table *Table) UnstageVoid(stage *StageStruct) {
	delete(stage.Tables, table)
	delete(stage.Tables_mapString, table.Name)
}

// commit table to the back repo (if it is already staged)
func (table *Table) Commit(stage *StageStruct) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTable(table)
		}
	}
	return table
}

func (table *Table) CommitVoid(stage *StageStruct) {
	table.Commit(stage)
}

// Checkout table to the back repo (if it is already staged)
func (table *Table) Checkout(stage *StageStruct) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTable(table)
		}
	}
	return table
}

// for satisfaction of GongStruct interface
func (table *Table) GetName() (res string) {
	return table.Name
}

// Stage puts tablecolumn to the model stage
func (tablecolumn *TableColumn) Stage(stage *StageStruct) *TableColumn {
	stage.TableColumns[tablecolumn] = __member
	stage.TableColumns_mapString[tablecolumn.Name] = tablecolumn

	return tablecolumn
}

// Unstage removes tablecolumn off the model stage
func (tablecolumn *TableColumn) Unstage(stage *StageStruct) *TableColumn {
	delete(stage.TableColumns, tablecolumn)
	delete(stage.TableColumns_mapString, tablecolumn.Name)
	return tablecolumn
}

// UnstageVoid removes tablecolumn off the model stage
func (tablecolumn *TableColumn) UnstageVoid(stage *StageStruct) {
	delete(stage.TableColumns, tablecolumn)
	delete(stage.TableColumns_mapString, tablecolumn.Name)
}

// commit tablecolumn to the back repo (if it is already staged)
func (tablecolumn *TableColumn) Commit(stage *StageStruct) *TableColumn {
	if _, ok := stage.TableColumns[tablecolumn]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTableColumn(tablecolumn)
		}
	}
	return tablecolumn
}

func (tablecolumn *TableColumn) CommitVoid(stage *StageStruct) {
	tablecolumn.Commit(stage)
}

// Checkout tablecolumn to the back repo (if it is already staged)
func (tablecolumn *TableColumn) Checkout(stage *StageStruct) *TableColumn {
	if _, ok := stage.TableColumns[tablecolumn]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTableColumn(tablecolumn)
		}
	}
	return tablecolumn
}

// for satisfaction of GongStruct interface
func (tablecolumn *TableColumn) GetName() (res string) {
	return tablecolumn.Name
}

// Stage puts tableproperties to the model stage
func (tableproperties *TableProperties) Stage(stage *StageStruct) *TableProperties {
	stage.TablePropertiess[tableproperties] = __member
	stage.TablePropertiess_mapString[tableproperties.Name] = tableproperties

	return tableproperties
}

// Unstage removes tableproperties off the model stage
func (tableproperties *TableProperties) Unstage(stage *StageStruct) *TableProperties {
	delete(stage.TablePropertiess, tableproperties)
	delete(stage.TablePropertiess_mapString, tableproperties.Name)
	return tableproperties
}

// UnstageVoid removes tableproperties off the model stage
func (tableproperties *TableProperties) UnstageVoid(stage *StageStruct) {
	delete(stage.TablePropertiess, tableproperties)
	delete(stage.TablePropertiess_mapString, tableproperties.Name)
}

// commit tableproperties to the back repo (if it is already staged)
func (tableproperties *TableProperties) Commit(stage *StageStruct) *TableProperties {
	if _, ok := stage.TablePropertiess[tableproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTableProperties(tableproperties)
		}
	}
	return tableproperties
}

func (tableproperties *TableProperties) CommitVoid(stage *StageStruct) {
	tableproperties.Commit(stage)
}

// Checkout tableproperties to the back repo (if it is already staged)
func (tableproperties *TableProperties) Checkout(stage *StageStruct) *TableProperties {
	if _, ok := stage.TablePropertiess[tableproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTableProperties(tableproperties)
		}
	}
	return tableproperties
}

// for satisfaction of GongStruct interface
func (tableproperties *TableProperties) GetName() (res string) {
	return tableproperties.Name
}

// Stage puts tablerow to the model stage
func (tablerow *TableRow) Stage(stage *StageStruct) *TableRow {
	stage.TableRows[tablerow] = __member
	stage.TableRows_mapString[tablerow.Name] = tablerow

	return tablerow
}

// Unstage removes tablerow off the model stage
func (tablerow *TableRow) Unstage(stage *StageStruct) *TableRow {
	delete(stage.TableRows, tablerow)
	delete(stage.TableRows_mapString, tablerow.Name)
	return tablerow
}

// UnstageVoid removes tablerow off the model stage
func (tablerow *TableRow) UnstageVoid(stage *StageStruct) {
	delete(stage.TableRows, tablerow)
	delete(stage.TableRows_mapString, tablerow.Name)
}

// commit tablerow to the back repo (if it is already staged)
func (tablerow *TableRow) Commit(stage *StageStruct) *TableRow {
	if _, ok := stage.TableRows[tablerow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTableRow(tablerow)
		}
	}
	return tablerow
}

func (tablerow *TableRow) CommitVoid(stage *StageStruct) {
	tablerow.Commit(stage)
}

// Checkout tablerow to the back repo (if it is already staged)
func (tablerow *TableRow) Checkout(stage *StageStruct) *TableRow {
	if _, ok := stage.TableRows[tablerow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTableRow(tablerow)
		}
	}
	return tablerow
}

// for satisfaction of GongStruct interface
func (tablerow *TableRow) GetName() (res string) {
	return tablerow.Name
}

// Stage puts tablestyle to the model stage
func (tablestyle *TableStyle) Stage(stage *StageStruct) *TableStyle {
	stage.TableStyles[tablestyle] = __member
	stage.TableStyles_mapString[tablestyle.Name] = tablestyle

	return tablestyle
}

// Unstage removes tablestyle off the model stage
func (tablestyle *TableStyle) Unstage(stage *StageStruct) *TableStyle {
	delete(stage.TableStyles, tablestyle)
	delete(stage.TableStyles_mapString, tablestyle.Name)
	return tablestyle
}

// UnstageVoid removes tablestyle off the model stage
func (tablestyle *TableStyle) UnstageVoid(stage *StageStruct) {
	delete(stage.TableStyles, tablestyle)
	delete(stage.TableStyles_mapString, tablestyle.Name)
}

// commit tablestyle to the back repo (if it is already staged)
func (tablestyle *TableStyle) Commit(stage *StageStruct) *TableStyle {
	if _, ok := stage.TableStyles[tablestyle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTableStyle(tablestyle)
		}
	}
	return tablestyle
}

func (tablestyle *TableStyle) CommitVoid(stage *StageStruct) {
	tablestyle.Commit(stage)
}

// Checkout tablestyle to the back repo (if it is already staged)
func (tablestyle *TableStyle) Checkout(stage *StageStruct) *TableStyle {
	if _, ok := stage.TableStyles[tablestyle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTableStyle(tablestyle)
		}
	}
	return tablestyle
}

// for satisfaction of GongStruct interface
func (tablestyle *TableStyle) GetName() (res string) {
	return tablestyle.Name
}

// Stage puts text to the model stage
func (text *Text) Stage(stage *StageStruct) *Text {
	stage.Texts[text] = __member
	stage.Texts_mapString[text.Name] = text

	return text
}

// Unstage removes text off the model stage
func (text *Text) Unstage(stage *StageStruct) *Text {
	delete(stage.Texts, text)
	delete(stage.Texts_mapString, text.Name)
	return text
}

// UnstageVoid removes text off the model stage
func (text *Text) UnstageVoid(stage *StageStruct) {
	delete(stage.Texts, text)
	delete(stage.Texts_mapString, text.Name)
}

// commit text to the back repo (if it is already staged)
func (text *Text) Commit(stage *StageStruct) *Text {
	if _, ok := stage.Texts[text]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitText(text)
		}
	}
	return text
}

func (text *Text) CommitVoid(stage *StageStruct) {
	text.Commit(stage)
}

// Checkout text to the back repo (if it is already staged)
func (text *Text) Checkout(stage *StageStruct) *Text {
	if _, ok := stage.Texts[text]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutText(text)
		}
	}
	return text
}

// for satisfaction of GongStruct interface
func (text *Text) GetName() (res string) {
	return text.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMBody(Body *Body)
	CreateORMDocument(Document *Document)
	CreateORMDocx(Docx *Docx)
	CreateORMFile(File *File)
	CreateORMNode(Node *Node)
	CreateORMParagraph(Paragraph *Paragraph)
	CreateORMParagraphProperties(ParagraphProperties *ParagraphProperties)
	CreateORMParagraphStyle(ParagraphStyle *ParagraphStyle)
	CreateORMRune(Rune *Rune)
	CreateORMRuneProperties(RuneProperties *RuneProperties)
	CreateORMTable(Table *Table)
	CreateORMTableColumn(TableColumn *TableColumn)
	CreateORMTableProperties(TableProperties *TableProperties)
	CreateORMTableRow(TableRow *TableRow)
	CreateORMTableStyle(TableStyle *TableStyle)
	CreateORMText(Text *Text)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMBody(Body *Body)
	DeleteORMDocument(Document *Document)
	DeleteORMDocx(Docx *Docx)
	DeleteORMFile(File *File)
	DeleteORMNode(Node *Node)
	DeleteORMParagraph(Paragraph *Paragraph)
	DeleteORMParagraphProperties(ParagraphProperties *ParagraphProperties)
	DeleteORMParagraphStyle(ParagraphStyle *ParagraphStyle)
	DeleteORMRune(Rune *Rune)
	DeleteORMRuneProperties(RuneProperties *RuneProperties)
	DeleteORMTable(Table *Table)
	DeleteORMTableColumn(TableColumn *TableColumn)
	DeleteORMTableProperties(TableProperties *TableProperties)
	DeleteORMTableRow(TableRow *TableRow)
	DeleteORMTableStyle(TableStyle *TableStyle)
	DeleteORMText(Text *Text)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Bodys = make(map[*Body]any)
	stage.Bodys_mapString = make(map[string]*Body)

	stage.Documents = make(map[*Document]any)
	stage.Documents_mapString = make(map[string]*Document)

	stage.Docxs = make(map[*Docx]any)
	stage.Docxs_mapString = make(map[string]*Docx)

	stage.Files = make(map[*File]any)
	stage.Files_mapString = make(map[string]*File)

	stage.Nodes = make(map[*Node]any)
	stage.Nodes_mapString = make(map[string]*Node)

	stage.Paragraphs = make(map[*Paragraph]any)
	stage.Paragraphs_mapString = make(map[string]*Paragraph)

	stage.ParagraphPropertiess = make(map[*ParagraphProperties]any)
	stage.ParagraphPropertiess_mapString = make(map[string]*ParagraphProperties)

	stage.ParagraphStyles = make(map[*ParagraphStyle]any)
	stage.ParagraphStyles_mapString = make(map[string]*ParagraphStyle)

	stage.Runes = make(map[*Rune]any)
	stage.Runes_mapString = make(map[string]*Rune)

	stage.RunePropertiess = make(map[*RuneProperties]any)
	stage.RunePropertiess_mapString = make(map[string]*RuneProperties)

	stage.Tables = make(map[*Table]any)
	stage.Tables_mapString = make(map[string]*Table)

	stage.TableColumns = make(map[*TableColumn]any)
	stage.TableColumns_mapString = make(map[string]*TableColumn)

	stage.TablePropertiess = make(map[*TableProperties]any)
	stage.TablePropertiess_mapString = make(map[string]*TableProperties)

	stage.TableRows = make(map[*TableRow]any)
	stage.TableRows_mapString = make(map[string]*TableRow)

	stage.TableStyles = make(map[*TableStyle]any)
	stage.TableStyles_mapString = make(map[string]*TableStyle)

	stage.Texts = make(map[*Text]any)
	stage.Texts_mapString = make(map[string]*Text)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Bodys = nil
	stage.Bodys_mapString = nil

	stage.Documents = nil
	stage.Documents_mapString = nil

	stage.Docxs = nil
	stage.Docxs_mapString = nil

	stage.Files = nil
	stage.Files_mapString = nil

	stage.Nodes = nil
	stage.Nodes_mapString = nil

	stage.Paragraphs = nil
	stage.Paragraphs_mapString = nil

	stage.ParagraphPropertiess = nil
	stage.ParagraphPropertiess_mapString = nil

	stage.ParagraphStyles = nil
	stage.ParagraphStyles_mapString = nil

	stage.Runes = nil
	stage.Runes_mapString = nil

	stage.RunePropertiess = nil
	stage.RunePropertiess_mapString = nil

	stage.Tables = nil
	stage.Tables_mapString = nil

	stage.TableColumns = nil
	stage.TableColumns_mapString = nil

	stage.TablePropertiess = nil
	stage.TablePropertiess_mapString = nil

	stage.TableRows = nil
	stage.TableRows_mapString = nil

	stage.TableStyles = nil
	stage.TableStyles_mapString = nil

	stage.Texts = nil
	stage.Texts_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for body := range stage.Bodys {
		body.Unstage(stage)
	}

	for document := range stage.Documents {
		document.Unstage(stage)
	}

	for docx := range stage.Docxs {
		docx.Unstage(stage)
	}

	for file := range stage.Files {
		file.Unstage(stage)
	}

	for node := range stage.Nodes {
		node.Unstage(stage)
	}

	for paragraph := range stage.Paragraphs {
		paragraph.Unstage(stage)
	}

	for paragraphproperties := range stage.ParagraphPropertiess {
		paragraphproperties.Unstage(stage)
	}

	for paragraphstyle := range stage.ParagraphStyles {
		paragraphstyle.Unstage(stage)
	}

	for rune := range stage.Runes {
		rune.Unstage(stage)
	}

	for runeproperties := range stage.RunePropertiess {
		runeproperties.Unstage(stage)
	}

	for table := range stage.Tables {
		table.Unstage(stage)
	}

	for tablecolumn := range stage.TableColumns {
		tablecolumn.Unstage(stage)
	}

	for tableproperties := range stage.TablePropertiess {
		tableproperties.Unstage(stage)
	}

	for tablerow := range stage.TableRows {
		tablerow.Unstage(stage)
	}

	for tablestyle := range stage.TableStyles {
		tablestyle.Unstage(stage)
	}

	for text := range stage.Texts {
		text.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Body]any:
		return any(&stage.Bodys).(*Type)
	case map[*Document]any:
		return any(&stage.Documents).(*Type)
	case map[*Docx]any:
		return any(&stage.Docxs).(*Type)
	case map[*File]any:
		return any(&stage.Files).(*Type)
	case map[*Node]any:
		return any(&stage.Nodes).(*Type)
	case map[*Paragraph]any:
		return any(&stage.Paragraphs).(*Type)
	case map[*ParagraphProperties]any:
		return any(&stage.ParagraphPropertiess).(*Type)
	case map[*ParagraphStyle]any:
		return any(&stage.ParagraphStyles).(*Type)
	case map[*Rune]any:
		return any(&stage.Runes).(*Type)
	case map[*RuneProperties]any:
		return any(&stage.RunePropertiess).(*Type)
	case map[*Table]any:
		return any(&stage.Tables).(*Type)
	case map[*TableColumn]any:
		return any(&stage.TableColumns).(*Type)
	case map[*TableProperties]any:
		return any(&stage.TablePropertiess).(*Type)
	case map[*TableRow]any:
		return any(&stage.TableRows).(*Type)
	case map[*TableStyle]any:
		return any(&stage.TableStyles).(*Type)
	case map[*Text]any:
		return any(&stage.Texts).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Body:
		return any(&stage.Bodys_mapString).(*Type)
	case map[string]*Document:
		return any(&stage.Documents_mapString).(*Type)
	case map[string]*Docx:
		return any(&stage.Docxs_mapString).(*Type)
	case map[string]*File:
		return any(&stage.Files_mapString).(*Type)
	case map[string]*Node:
		return any(&stage.Nodes_mapString).(*Type)
	case map[string]*Paragraph:
		return any(&stage.Paragraphs_mapString).(*Type)
	case map[string]*ParagraphProperties:
		return any(&stage.ParagraphPropertiess_mapString).(*Type)
	case map[string]*ParagraphStyle:
		return any(&stage.ParagraphStyles_mapString).(*Type)
	case map[string]*Rune:
		return any(&stage.Runes_mapString).(*Type)
	case map[string]*RuneProperties:
		return any(&stage.RunePropertiess_mapString).(*Type)
	case map[string]*Table:
		return any(&stage.Tables_mapString).(*Type)
	case map[string]*TableColumn:
		return any(&stage.TableColumns_mapString).(*Type)
	case map[string]*TableProperties:
		return any(&stage.TablePropertiess_mapString).(*Type)
	case map[string]*TableRow:
		return any(&stage.TableRows_mapString).(*Type)
	case map[string]*TableStyle:
		return any(&stage.TableStyles_mapString).(*Type)
	case map[string]*Text:
		return any(&stage.Texts_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Body:
		return any(&stage.Bodys).(*map[*Type]any)
	case Document:
		return any(&stage.Documents).(*map[*Type]any)
	case Docx:
		return any(&stage.Docxs).(*map[*Type]any)
	case File:
		return any(&stage.Files).(*map[*Type]any)
	case Node:
		return any(&stage.Nodes).(*map[*Type]any)
	case Paragraph:
		return any(&stage.Paragraphs).(*map[*Type]any)
	case ParagraphProperties:
		return any(&stage.ParagraphPropertiess).(*map[*Type]any)
	case ParagraphStyle:
		return any(&stage.ParagraphStyles).(*map[*Type]any)
	case Rune:
		return any(&stage.Runes).(*map[*Type]any)
	case RuneProperties:
		return any(&stage.RunePropertiess).(*map[*Type]any)
	case Table:
		return any(&stage.Tables).(*map[*Type]any)
	case TableColumn:
		return any(&stage.TableColumns).(*map[*Type]any)
	case TableProperties:
		return any(&stage.TablePropertiess).(*map[*Type]any)
	case TableRow:
		return any(&stage.TableRows).(*map[*Type]any)
	case TableStyle:
		return any(&stage.TableStyles).(*map[*Type]any)
	case Text:
		return any(&stage.Texts).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Body:
		return any(&stage.Bodys).(*map[Type]any)
	case *Document:
		return any(&stage.Documents).(*map[Type]any)
	case *Docx:
		return any(&stage.Docxs).(*map[Type]any)
	case *File:
		return any(&stage.Files).(*map[Type]any)
	case *Node:
		return any(&stage.Nodes).(*map[Type]any)
	case *Paragraph:
		return any(&stage.Paragraphs).(*map[Type]any)
	case *ParagraphProperties:
		return any(&stage.ParagraphPropertiess).(*map[Type]any)
	case *ParagraphStyle:
		return any(&stage.ParagraphStyles).(*map[Type]any)
	case *Rune:
		return any(&stage.Runes).(*map[Type]any)
	case *RuneProperties:
		return any(&stage.RunePropertiess).(*map[Type]any)
	case *Table:
		return any(&stage.Tables).(*map[Type]any)
	case *TableColumn:
		return any(&stage.TableColumns).(*map[Type]any)
	case *TableProperties:
		return any(&stage.TablePropertiess).(*map[Type]any)
	case *TableRow:
		return any(&stage.TableRows).(*map[Type]any)
	case *TableStyle:
		return any(&stage.TableStyles).(*map[Type]any)
	case *Text:
		return any(&stage.Texts).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Body:
		return any(&stage.Bodys_mapString).(*map[string]*Type)
	case Document:
		return any(&stage.Documents_mapString).(*map[string]*Type)
	case Docx:
		return any(&stage.Docxs_mapString).(*map[string]*Type)
	case File:
		return any(&stage.Files_mapString).(*map[string]*Type)
	case Node:
		return any(&stage.Nodes_mapString).(*map[string]*Type)
	case Paragraph:
		return any(&stage.Paragraphs_mapString).(*map[string]*Type)
	case ParagraphProperties:
		return any(&stage.ParagraphPropertiess_mapString).(*map[string]*Type)
	case ParagraphStyle:
		return any(&stage.ParagraphStyles_mapString).(*map[string]*Type)
	case Rune:
		return any(&stage.Runes_mapString).(*map[string]*Type)
	case RuneProperties:
		return any(&stage.RunePropertiess_mapString).(*map[string]*Type)
	case Table:
		return any(&stage.Tables_mapString).(*map[string]*Type)
	case TableColumn:
		return any(&stage.TableColumns_mapString).(*map[string]*Type)
	case TableProperties:
		return any(&stage.TablePropertiess_mapString).(*map[string]*Type)
	case TableRow:
		return any(&stage.TableRows_mapString).(*map[string]*Type)
	case TableStyle:
		return any(&stage.TableStyles_mapString).(*map[string]*Type)
	case Text:
		return any(&stage.Texts_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Body:
		return any(&Body{
			// Initialisation of associations
			// field is initialized with an instance of Paragraph with the name of the field
			Paragraphs: []*Paragraph{{Name: "Paragraphs"}},
			// field is initialized with an instance of Table with the name of the field
			Tables: []*Table{{Name: "Tables"}},
			// field is initialized with an instance of Paragraph with the name of the field
			LastParagraph: &Paragraph{Name: "LastParagraph"},
		}).(*Type)
	case Document:
		return any(&Document{
			// Initialisation of associations
			// field is initialized with an instance of File with the name of the field
			File: &File{Name: "File"},
			// field is initialized with an instance of Node with the name of the field
			Root: &Node{Name: "Root"},
			// field is initialized with an instance of Body with the name of the field
			Body: &Body{Name: "Body"},
		}).(*Type)
	case Docx:
		return any(&Docx{
			// Initialisation of associations
			// field is initialized with an instance of File with the name of the field
			Files: []*File{{Name: "Files"}},
			// field is initialized with an instance of Document with the name of the field
			Document: &Document{Name: "Document"},
		}).(*Type)
	case File:
		return any(&File{
			// Initialisation of associations
		}).(*Type)
	case Node:
		return any(&Node{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Nodes: []*Node{{Name: "Nodes"}},
		}).(*Type)
	case Paragraph:
		return any(&Paragraph{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of ParagraphProperties with the name of the field
			ParagraphProperties: &ParagraphProperties{Name: "ParagraphProperties"},
			// field is initialized with an instance of Rune with the name of the field
			Runes: []*Rune{{Name: "Runes"}},
			// field is initialized with an instance of Paragraph with the name of the field
			Next: &Paragraph{Name: "Next"},
			// field is initialized with an instance of Paragraph with the name of the field
			Previous: &Paragraph{Name: "Previous"},
			// field is initialized with an instance of Body with the name of the field
			EnclosingBody: &Body{Name: "EnclosingBody"},
			// field is initialized with an instance of TableColumn with the name of the field
			EnclosingTableColumn: &TableColumn{Name: "EnclosingTableColumn"},
		}).(*Type)
	case ParagraphProperties:
		return any(&ParagraphProperties{
			// Initialisation of associations
			// field is initialized with an instance of ParagraphStyle with the name of the field
			ParagraphStyle: &ParagraphStyle{Name: "ParagraphStyle"},
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
		}).(*Type)
	case ParagraphStyle:
		return any(&ParagraphStyle{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
		}).(*Type)
	case Rune:
		return any(&Rune{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of Text with the name of the field
			Text: &Text{Name: "Text"},
			// field is initialized with an instance of RuneProperties with the name of the field
			RuneProperties: &RuneProperties{Name: "RuneProperties"},
			// field is initialized with an instance of Paragraph with the name of the field
			EnclosingParagraph: &Paragraph{Name: "EnclosingParagraph"},
		}).(*Type)
	case RuneProperties:
		return any(&RuneProperties{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
		}).(*Type)
	case Table:
		return any(&Table{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of TableProperties with the name of the field
			TableProperties: &TableProperties{Name: "TableProperties"},
			// field is initialized with an instance of TableRow with the name of the field
			TableRows: []*TableRow{{Name: "TableRows"}},
		}).(*Type)
	case TableColumn:
		return any(&TableColumn{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of Paragraph with the name of the field
			Paragraphs: []*Paragraph{{Name: "Paragraphs"}},
		}).(*Type)
	case TableProperties:
		return any(&TableProperties{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of TableStyle with the name of the field
			TableStyle: &TableStyle{Name: "TableStyle"},
		}).(*Type)
	case TableRow:
		return any(&TableRow{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of TableColumn with the name of the field
			TableColumns: []*TableColumn{{Name: "TableColumns"}},
		}).(*Type)
	case TableStyle:
		return any(&TableStyle{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
		}).(*Type)
	case Text:
		return any(&Text{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of Rune with the name of the field
			EnclosingRune: &Rune{Name: "EnclosingRune"},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Body
	case Body:
		switch fieldname {
		// insertion point for per direct association field
		case "LastParagraph":
			res := make(map[*Paragraph][]*Body)
			for body := range stage.Bodys {
				if body.LastParagraph != nil {
					paragraph_ := body.LastParagraph
					var bodys []*Body
					_, ok := res[paragraph_]
					if ok {
						bodys = res[paragraph_]
					} else {
						bodys = make([]*Body, 0)
					}
					bodys = append(bodys, body)
					res[paragraph_] = bodys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Document
	case Document:
		switch fieldname {
		// insertion point for per direct association field
		case "File":
			res := make(map[*File][]*Document)
			for document := range stage.Documents {
				if document.File != nil {
					file_ := document.File
					var documents []*Document
					_, ok := res[file_]
					if ok {
						documents = res[file_]
					} else {
						documents = make([]*Document, 0)
					}
					documents = append(documents, document)
					res[file_] = documents
				}
			}
			return any(res).(map[*End][]*Start)
		case "Root":
			res := make(map[*Node][]*Document)
			for document := range stage.Documents {
				if document.Root != nil {
					node_ := document.Root
					var documents []*Document
					_, ok := res[node_]
					if ok {
						documents = res[node_]
					} else {
						documents = make([]*Document, 0)
					}
					documents = append(documents, document)
					res[node_] = documents
				}
			}
			return any(res).(map[*End][]*Start)
		case "Body":
			res := make(map[*Body][]*Document)
			for document := range stage.Documents {
				if document.Body != nil {
					body_ := document.Body
					var documents []*Document
					_, ok := res[body_]
					if ok {
						documents = res[body_]
					} else {
						documents = make([]*Document, 0)
					}
					documents = append(documents, document)
					res[body_] = documents
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Docx
	case Docx:
		switch fieldname {
		// insertion point for per direct association field
		case "Document":
			res := make(map[*Document][]*Docx)
			for docx := range stage.Docxs {
				if docx.Document != nil {
					document_ := docx.Document
					var docxs []*Docx
					_, ok := res[document_]
					if ok {
						docxs = res[document_]
					} else {
						docxs = make([]*Docx, 0)
					}
					docxs = append(docxs, docx)
					res[document_] = docxs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of File
	case File:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Paragraph
	case Paragraph:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.Node != nil {
					node_ := paragraph.Node
					var paragraphs []*Paragraph
					_, ok := res[node_]
					if ok {
						paragraphs = res[node_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[node_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "ParagraphProperties":
			res := make(map[*ParagraphProperties][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.ParagraphProperties != nil {
					paragraphproperties_ := paragraph.ParagraphProperties
					var paragraphs []*Paragraph
					_, ok := res[paragraphproperties_]
					if ok {
						paragraphs = res[paragraphproperties_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[paragraphproperties_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Next":
			res := make(map[*Paragraph][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.Next != nil {
					paragraph_ := paragraph.Next
					var paragraphs []*Paragraph
					_, ok := res[paragraph_]
					if ok {
						paragraphs = res[paragraph_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[paragraph_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Previous":
			res := make(map[*Paragraph][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.Previous != nil {
					paragraph_ := paragraph.Previous
					var paragraphs []*Paragraph
					_, ok := res[paragraph_]
					if ok {
						paragraphs = res[paragraph_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[paragraph_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "EnclosingBody":
			res := make(map[*Body][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.EnclosingBody != nil {
					body_ := paragraph.EnclosingBody
					var paragraphs []*Paragraph
					_, ok := res[body_]
					if ok {
						paragraphs = res[body_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[body_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "EnclosingTableColumn":
			res := make(map[*TableColumn][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.EnclosingTableColumn != nil {
					tablecolumn_ := paragraph.EnclosingTableColumn
					var paragraphs []*Paragraph
					_, ok := res[tablecolumn_]
					if ok {
						paragraphs = res[tablecolumn_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[tablecolumn_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParagraphProperties
	case ParagraphProperties:
		switch fieldname {
		// insertion point for per direct association field
		case "ParagraphStyle":
			res := make(map[*ParagraphStyle][]*ParagraphProperties)
			for paragraphproperties := range stage.ParagraphPropertiess {
				if paragraphproperties.ParagraphStyle != nil {
					paragraphstyle_ := paragraphproperties.ParagraphStyle
					var paragraphpropertiess []*ParagraphProperties
					_, ok := res[paragraphstyle_]
					if ok {
						paragraphpropertiess = res[paragraphstyle_]
					} else {
						paragraphpropertiess = make([]*ParagraphProperties, 0)
					}
					paragraphpropertiess = append(paragraphpropertiess, paragraphproperties)
					res[paragraphstyle_] = paragraphpropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		case "Node":
			res := make(map[*Node][]*ParagraphProperties)
			for paragraphproperties := range stage.ParagraphPropertiess {
				if paragraphproperties.Node != nil {
					node_ := paragraphproperties.Node
					var paragraphpropertiess []*ParagraphProperties
					_, ok := res[node_]
					if ok {
						paragraphpropertiess = res[node_]
					} else {
						paragraphpropertiess = make([]*ParagraphProperties, 0)
					}
					paragraphpropertiess = append(paragraphpropertiess, paragraphproperties)
					res[node_] = paragraphpropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParagraphStyle
	case ParagraphStyle:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*ParagraphStyle)
			for paragraphstyle := range stage.ParagraphStyles {
				if paragraphstyle.Node != nil {
					node_ := paragraphstyle.Node
					var paragraphstyles []*ParagraphStyle
					_, ok := res[node_]
					if ok {
						paragraphstyles = res[node_]
					} else {
						paragraphstyles = make([]*ParagraphStyle, 0)
					}
					paragraphstyles = append(paragraphstyles, paragraphstyle)
					res[node_] = paragraphstyles
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Rune
	case Rune:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*Rune)
			for rune := range stage.Runes {
				if rune.Node != nil {
					node_ := rune.Node
					var runes []*Rune
					_, ok := res[node_]
					if ok {
						runes = res[node_]
					} else {
						runes = make([]*Rune, 0)
					}
					runes = append(runes, rune)
					res[node_] = runes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Text":
			res := make(map[*Text][]*Rune)
			for rune := range stage.Runes {
				if rune.Text != nil {
					text_ := rune.Text
					var runes []*Rune
					_, ok := res[text_]
					if ok {
						runes = res[text_]
					} else {
						runes = make([]*Rune, 0)
					}
					runes = append(runes, rune)
					res[text_] = runes
				}
			}
			return any(res).(map[*End][]*Start)
		case "RuneProperties":
			res := make(map[*RuneProperties][]*Rune)
			for rune := range stage.Runes {
				if rune.RuneProperties != nil {
					runeproperties_ := rune.RuneProperties
					var runes []*Rune
					_, ok := res[runeproperties_]
					if ok {
						runes = res[runeproperties_]
					} else {
						runes = make([]*Rune, 0)
					}
					runes = append(runes, rune)
					res[runeproperties_] = runes
				}
			}
			return any(res).(map[*End][]*Start)
		case "EnclosingParagraph":
			res := make(map[*Paragraph][]*Rune)
			for rune := range stage.Runes {
				if rune.EnclosingParagraph != nil {
					paragraph_ := rune.EnclosingParagraph
					var runes []*Rune
					_, ok := res[paragraph_]
					if ok {
						runes = res[paragraph_]
					} else {
						runes = make([]*Rune, 0)
					}
					runes = append(runes, rune)
					res[paragraph_] = runes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RuneProperties
	case RuneProperties:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*RuneProperties)
			for runeproperties := range stage.RunePropertiess {
				if runeproperties.Node != nil {
					node_ := runeproperties.Node
					var runepropertiess []*RuneProperties
					_, ok := res[node_]
					if ok {
						runepropertiess = res[node_]
					} else {
						runepropertiess = make([]*RuneProperties, 0)
					}
					runepropertiess = append(runepropertiess, runeproperties)
					res[node_] = runepropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*Table)
			for table := range stage.Tables {
				if table.Node != nil {
					node_ := table.Node
					var tables []*Table
					_, ok := res[node_]
					if ok {
						tables = res[node_]
					} else {
						tables = make([]*Table, 0)
					}
					tables = append(tables, table)
					res[node_] = tables
				}
			}
			return any(res).(map[*End][]*Start)
		case "TableProperties":
			res := make(map[*TableProperties][]*Table)
			for table := range stage.Tables {
				if table.TableProperties != nil {
					tableproperties_ := table.TableProperties
					var tables []*Table
					_, ok := res[tableproperties_]
					if ok {
						tables = res[tableproperties_]
					} else {
						tables = make([]*Table, 0)
					}
					tables = append(tables, table)
					res[tableproperties_] = tables
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableColumn
	case TableColumn:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*TableColumn)
			for tablecolumn := range stage.TableColumns {
				if tablecolumn.Node != nil {
					node_ := tablecolumn.Node
					var tablecolumns []*TableColumn
					_, ok := res[node_]
					if ok {
						tablecolumns = res[node_]
					} else {
						tablecolumns = make([]*TableColumn, 0)
					}
					tablecolumns = append(tablecolumns, tablecolumn)
					res[node_] = tablecolumns
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableProperties
	case TableProperties:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*TableProperties)
			for tableproperties := range stage.TablePropertiess {
				if tableproperties.Node != nil {
					node_ := tableproperties.Node
					var tablepropertiess []*TableProperties
					_, ok := res[node_]
					if ok {
						tablepropertiess = res[node_]
					} else {
						tablepropertiess = make([]*TableProperties, 0)
					}
					tablepropertiess = append(tablepropertiess, tableproperties)
					res[node_] = tablepropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		case "TableStyle":
			res := make(map[*TableStyle][]*TableProperties)
			for tableproperties := range stage.TablePropertiess {
				if tableproperties.TableStyle != nil {
					tablestyle_ := tableproperties.TableStyle
					var tablepropertiess []*TableProperties
					_, ok := res[tablestyle_]
					if ok {
						tablepropertiess = res[tablestyle_]
					} else {
						tablepropertiess = make([]*TableProperties, 0)
					}
					tablepropertiess = append(tablepropertiess, tableproperties)
					res[tablestyle_] = tablepropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableRow
	case TableRow:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*TableRow)
			for tablerow := range stage.TableRows {
				if tablerow.Node != nil {
					node_ := tablerow.Node
					var tablerows []*TableRow
					_, ok := res[node_]
					if ok {
						tablerows = res[node_]
					} else {
						tablerows = make([]*TableRow, 0)
					}
					tablerows = append(tablerows, tablerow)
					res[node_] = tablerows
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableStyle
	case TableStyle:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*TableStyle)
			for tablestyle := range stage.TableStyles {
				if tablestyle.Node != nil {
					node_ := tablestyle.Node
					var tablestyles []*TableStyle
					_, ok := res[node_]
					if ok {
						tablestyles = res[node_]
					} else {
						tablestyles = make([]*TableStyle, 0)
					}
					tablestyles = append(tablestyles, tablestyle)
					res[node_] = tablestyles
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Text
	case Text:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*Text)
			for text := range stage.Texts {
				if text.Node != nil {
					node_ := text.Node
					var texts []*Text
					_, ok := res[node_]
					if ok {
						texts = res[node_]
					} else {
						texts = make([]*Text, 0)
					}
					texts = append(texts, text)
					res[node_] = texts
				}
			}
			return any(res).(map[*End][]*Start)
		case "EnclosingRune":
			res := make(map[*Rune][]*Text)
			for text := range stage.Texts {
				if text.EnclosingRune != nil {
					rune_ := text.EnclosingRune
					var texts []*Text
					_, ok := res[rune_]
					if ok {
						texts = res[rune_]
					} else {
						texts = make([]*Text, 0)
					}
					texts = append(texts, text)
					res[rune_] = texts
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Body
	case Body:
		switch fieldname {
		// insertion point for per direct association field
		case "Paragraphs":
			res := make(map[*Paragraph]*Body)
			for body := range stage.Bodys {
				for _, paragraph_ := range body.Paragraphs {
					res[paragraph_] = body
				}
			}
			return any(res).(map[*End]*Start)
		case "Tables":
			res := make(map[*Table]*Body)
			for body := range stage.Bodys {
				for _, table_ := range body.Tables {
					res[table_] = body
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Document
	case Document:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Docx
	case Docx:
		switch fieldname {
		// insertion point for per direct association field
		case "Files":
			res := make(map[*File]*Docx)
			for docx := range stage.Docxs {
				for _, file_ := range docx.Files {
					res[file_] = docx
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of File
	case File:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		case "Nodes":
			res := make(map[*Node]*Node)
			for node := range stage.Nodes {
				for _, node_ := range node.Nodes {
					res[node_] = node
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Paragraph
	case Paragraph:
		switch fieldname {
		// insertion point for per direct association field
		case "Runes":
			res := make(map[*Rune]*Paragraph)
			for paragraph := range stage.Paragraphs {
				for _, rune_ := range paragraph.Runes {
					res[rune_] = paragraph
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ParagraphProperties
	case ParagraphProperties:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ParagraphStyle
	case ParagraphStyle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Rune
	case Rune:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RuneProperties
	case RuneProperties:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		case "TableRows":
			res := make(map[*TableRow]*Table)
			for table := range stage.Tables {
				for _, tablerow_ := range table.TableRows {
					res[tablerow_] = table
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of TableColumn
	case TableColumn:
		switch fieldname {
		// insertion point for per direct association field
		case "Paragraphs":
			res := make(map[*Paragraph]*TableColumn)
			for tablecolumn := range stage.TableColumns {
				for _, paragraph_ := range tablecolumn.Paragraphs {
					res[paragraph_] = tablecolumn
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of TableProperties
	case TableProperties:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TableRow
	case TableRow:
		switch fieldname {
		// insertion point for per direct association field
		case "TableColumns":
			res := make(map[*TableColumn]*TableRow)
			for tablerow := range stage.TableRows {
				for _, tablecolumn_ := range tablerow.TableColumns {
					res[tablecolumn_] = tablerow
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of TableStyle
	case TableStyle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Text
	case Text:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Body:
		res = "Body"
	case Document:
		res = "Document"
	case Docx:
		res = "Docx"
	case File:
		res = "File"
	case Node:
		res = "Node"
	case Paragraph:
		res = "Paragraph"
	case ParagraphProperties:
		res = "ParagraphProperties"
	case ParagraphStyle:
		res = "ParagraphStyle"
	case Rune:
		res = "Rune"
	case RuneProperties:
		res = "RuneProperties"
	case Table:
		res = "Table"
	case TableColumn:
		res = "TableColumn"
	case TableProperties:
		res = "TableProperties"
	case TableRow:
		res = "TableRow"
	case TableStyle:
		res = "TableStyle"
	case Text:
		res = "Text"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Body:
		res = "Body"
	case *Document:
		res = "Document"
	case *Docx:
		res = "Docx"
	case *File:
		res = "File"
	case *Node:
		res = "Node"
	case *Paragraph:
		res = "Paragraph"
	case *ParagraphProperties:
		res = "ParagraphProperties"
	case *ParagraphStyle:
		res = "ParagraphStyle"
	case *Rune:
		res = "Rune"
	case *RuneProperties:
		res = "RuneProperties"
	case *Table:
		res = "Table"
	case *TableColumn:
		res = "TableColumn"
	case *TableProperties:
		res = "TableProperties"
	case *TableRow:
		res = "TableRow"
	case *TableStyle:
		res = "TableStyle"
	case *Text:
		res = "Text"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Body:
		res = []string{"Name", "Paragraphs", "Tables", "LastParagraph"}
	case Document:
		res = []string{"Name", "File", "Root", "Body"}
	case Docx:
		res = []string{"Name", "Files", "Document"}
	case File:
		res = []string{"Name"}
	case Node:
		res = []string{"Name", "Nodes"}
	case Paragraph:
		res = []string{"Name", "Content", "Node", "ParagraphProperties", "Runes", "Text", "Next", "Previous", "EnclosingBody", "EnclosingTableColumn"}
	case ParagraphProperties:
		res = []string{"Name", "Content", "ParagraphStyle", "Node"}
	case ParagraphStyle:
		res = []string{"Name", "Node", "Content", "ValAttr"}
	case Rune:
		res = []string{"Name", "Content", "Node", "Text", "RuneProperties", "EnclosingParagraph"}
	case RuneProperties:
		res = []string{"Name", "Node", "IsBold", "IsStrike", "IsItalic", "Content"}
	case Table:
		res = []string{"Name", "Node", "Content", "TableProperties", "TableRows"}
	case TableColumn:
		res = []string{"Name", "Content", "Node", "Paragraphs"}
	case TableProperties:
		res = []string{"Name", "Node", "Content", "TableStyle"}
	case TableRow:
		res = []string{"Name", "Content", "Node", "TableColumns"}
	case TableStyle:
		res = []string{"Name", "Node", "Content", "Val"}
	case Text:
		res = []string{"Name", "Content", "Node", "PreserveWhiteSpace", "EnclosingRune"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Body:
		var rf ReverseField
		_ = rf
	case Document:
		var rf ReverseField
		_ = rf
	case Docx:
		var rf ReverseField
		_ = rf
	case File:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Docx"
		rf.Fieldname = "Files"
		res = append(res, rf)
	case Node:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Node"
		rf.Fieldname = "Nodes"
		res = append(res, rf)
	case Paragraph:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Body"
		rf.Fieldname = "Paragraphs"
		res = append(res, rf)
		rf.GongstructName = "TableColumn"
		rf.Fieldname = "Paragraphs"
		res = append(res, rf)
	case ParagraphProperties:
		var rf ReverseField
		_ = rf
	case ParagraphStyle:
		var rf ReverseField
		_ = rf
	case Rune:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Paragraph"
		rf.Fieldname = "Runes"
		res = append(res, rf)
	case RuneProperties:
		var rf ReverseField
		_ = rf
	case Table:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Body"
		rf.Fieldname = "Tables"
		res = append(res, rf)
	case TableColumn:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TableRow"
		rf.Fieldname = "TableColumns"
		res = append(res, rf)
	case TableProperties:
		var rf ReverseField
		_ = rf
	case TableRow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Table"
		rf.Fieldname = "TableRows"
		res = append(res, rf)
	case TableStyle:
		var rf ReverseField
		_ = rf
	case Text:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Body:
		res = []string{"Name", "Paragraphs", "Tables", "LastParagraph"}
	case *Document:
		res = []string{"Name", "File", "Root", "Body"}
	case *Docx:
		res = []string{"Name", "Files", "Document"}
	case *File:
		res = []string{"Name"}
	case *Node:
		res = []string{"Name", "Nodes"}
	case *Paragraph:
		res = []string{"Name", "Content", "Node", "ParagraphProperties", "Runes", "Text", "Next", "Previous", "EnclosingBody", "EnclosingTableColumn"}
	case *ParagraphProperties:
		res = []string{"Name", "Content", "ParagraphStyle", "Node"}
	case *ParagraphStyle:
		res = []string{"Name", "Node", "Content", "ValAttr"}
	case *Rune:
		res = []string{"Name", "Content", "Node", "Text", "RuneProperties", "EnclosingParagraph"}
	case *RuneProperties:
		res = []string{"Name", "Node", "IsBold", "IsStrike", "IsItalic", "Content"}
	case *Table:
		res = []string{"Name", "Node", "Content", "TableProperties", "TableRows"}
	case *TableColumn:
		res = []string{"Name", "Content", "Node", "Paragraphs"}
	case *TableProperties:
		res = []string{"Name", "Node", "Content", "TableStyle"}
	case *TableRow:
		res = []string{"Name", "Content", "Node", "TableColumns"}
	case *TableStyle:
		res = []string{"Name", "Node", "Content", "Val"}
	case *Text:
		res = []string{"Name", "Content", "Node", "PreserveWhiteSpace", "EnclosingRune"}
	}
	return
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt     GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat   GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool    GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeOthers  GongFieldValueType = "GongFieldValueTypeOthers"
)

type GongFieldValue struct {
	valueString string
	GongFieldValueType
	valueInt   int
	valueFloat float64
	valueBool  bool
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}
	
func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}
	
func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

func GetFieldStringValueFromPointer(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Body:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Paragraphs":
			for idx, __instance__ := range inferedInstance.Paragraphs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Tables":
			for idx, __instance__ := range inferedInstance.Tables {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "LastParagraph":
			if inferedInstance.LastParagraph != nil {
				res.valueString = inferedInstance.LastParagraph.Name
			}
		}
	case *Document:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "File":
			if inferedInstance.File != nil {
				res.valueString = inferedInstance.File.Name
			}
		case "Root":
			if inferedInstance.Root != nil {
				res.valueString = inferedInstance.Root.Name
			}
		case "Body":
			if inferedInstance.Body != nil {
				res.valueString = inferedInstance.Body.Name
			}
		}
	case *Docx:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Files":
			for idx, __instance__ := range inferedInstance.Files {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Document":
			if inferedInstance.Document != nil {
				res.valueString = inferedInstance.Document.Name
			}
		}
	case *File:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Node:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Nodes":
			for idx, __instance__ := range inferedInstance.Nodes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Paragraph:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "ParagraphProperties":
			if inferedInstance.ParagraphProperties != nil {
				res.valueString = inferedInstance.ParagraphProperties.Name
			}
		case "Runes":
			for idx, __instance__ := range inferedInstance.Runes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Text":
			res.valueString = inferedInstance.Text
		case "Next":
			if inferedInstance.Next != nil {
				res.valueString = inferedInstance.Next.Name
			}
		case "Previous":
			if inferedInstance.Previous != nil {
				res.valueString = inferedInstance.Previous.Name
			}
		case "EnclosingBody":
			if inferedInstance.EnclosingBody != nil {
				res.valueString = inferedInstance.EnclosingBody.Name
			}
		case "EnclosingTableColumn":
			if inferedInstance.EnclosingTableColumn != nil {
				res.valueString = inferedInstance.EnclosingTableColumn.Name
			}
		}
	case *ParagraphProperties:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "ParagraphStyle":
			if inferedInstance.ParagraphStyle != nil {
				res.valueString = inferedInstance.ParagraphStyle.Name
			}
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		}
	case *ParagraphStyle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Content":
			res.valueString = inferedInstance.Content
		case "ValAttr":
			res.valueString = inferedInstance.ValAttr
		}
	case *Rune:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Text":
			if inferedInstance.Text != nil {
				res.valueString = inferedInstance.Text.Name
			}
		case "RuneProperties":
			if inferedInstance.RuneProperties != nil {
				res.valueString = inferedInstance.RuneProperties.Name
			}
		case "EnclosingParagraph":
			if inferedInstance.EnclosingParagraph != nil {
				res.valueString = inferedInstance.EnclosingParagraph.Name
			}
		}
	case *RuneProperties:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "IsBold":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsBold)
			res.valueBool = inferedInstance.IsBold
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsStrike":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsStrike)
			res.valueBool = inferedInstance.IsStrike
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsItalic":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsItalic)
			res.valueBool = inferedInstance.IsItalic
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Content":
			res.valueString = inferedInstance.Content
		}
	case *Table:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Content":
			res.valueString = inferedInstance.Content
		case "TableProperties":
			if inferedInstance.TableProperties != nil {
				res.valueString = inferedInstance.TableProperties.Name
			}
		case "TableRows":
			for idx, __instance__ := range inferedInstance.TableRows {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *TableColumn:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Paragraphs":
			for idx, __instance__ := range inferedInstance.Paragraphs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *TableProperties:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Content":
			res.valueString = inferedInstance.Content
		case "TableStyle":
			if inferedInstance.TableStyle != nil {
				res.valueString = inferedInstance.TableStyle.Name
			}
		}
	case *TableRow:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "TableColumns":
			for idx, __instance__ := range inferedInstance.TableColumns {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *TableStyle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Content":
			res.valueString = inferedInstance.Content
		case "Val":
			res.valueString = inferedInstance.Val
		}
	case *Text:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "PreserveWhiteSpace":
			res.valueString = fmt.Sprintf("%t", inferedInstance.PreserveWhiteSpace)
			res.valueBool = inferedInstance.PreserveWhiteSpace
			res.GongFieldValueType = GongFieldValueTypeBool
		case "EnclosingRune":
			if inferedInstance.EnclosingRune != nil {
				res.valueString = inferedInstance.EnclosingRune.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Body:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Paragraphs":
			for idx, __instance__ := range inferedInstance.Paragraphs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Tables":
			for idx, __instance__ := range inferedInstance.Tables {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "LastParagraph":
			if inferedInstance.LastParagraph != nil {
				res.valueString = inferedInstance.LastParagraph.Name
			}
		}
	case Document:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "File":
			if inferedInstance.File != nil {
				res.valueString = inferedInstance.File.Name
			}
		case "Root":
			if inferedInstance.Root != nil {
				res.valueString = inferedInstance.Root.Name
			}
		case "Body":
			if inferedInstance.Body != nil {
				res.valueString = inferedInstance.Body.Name
			}
		}
	case Docx:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Files":
			for idx, __instance__ := range inferedInstance.Files {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Document":
			if inferedInstance.Document != nil {
				res.valueString = inferedInstance.Document.Name
			}
		}
	case File:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Node:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Nodes":
			for idx, __instance__ := range inferedInstance.Nodes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Paragraph:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "ParagraphProperties":
			if inferedInstance.ParagraphProperties != nil {
				res.valueString = inferedInstance.ParagraphProperties.Name
			}
		case "Runes":
			for idx, __instance__ := range inferedInstance.Runes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Text":
			res.valueString = inferedInstance.Text
		case "Next":
			if inferedInstance.Next != nil {
				res.valueString = inferedInstance.Next.Name
			}
		case "Previous":
			if inferedInstance.Previous != nil {
				res.valueString = inferedInstance.Previous.Name
			}
		case "EnclosingBody":
			if inferedInstance.EnclosingBody != nil {
				res.valueString = inferedInstance.EnclosingBody.Name
			}
		case "EnclosingTableColumn":
			if inferedInstance.EnclosingTableColumn != nil {
				res.valueString = inferedInstance.EnclosingTableColumn.Name
			}
		}
	case ParagraphProperties:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "ParagraphStyle":
			if inferedInstance.ParagraphStyle != nil {
				res.valueString = inferedInstance.ParagraphStyle.Name
			}
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		}
	case ParagraphStyle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Content":
			res.valueString = inferedInstance.Content
		case "ValAttr":
			res.valueString = inferedInstance.ValAttr
		}
	case Rune:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Text":
			if inferedInstance.Text != nil {
				res.valueString = inferedInstance.Text.Name
			}
		case "RuneProperties":
			if inferedInstance.RuneProperties != nil {
				res.valueString = inferedInstance.RuneProperties.Name
			}
		case "EnclosingParagraph":
			if inferedInstance.EnclosingParagraph != nil {
				res.valueString = inferedInstance.EnclosingParagraph.Name
			}
		}
	case RuneProperties:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "IsBold":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsBold)
			res.valueBool = inferedInstance.IsBold
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsStrike":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsStrike)
			res.valueBool = inferedInstance.IsStrike
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsItalic":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsItalic)
			res.valueBool = inferedInstance.IsItalic
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Content":
			res.valueString = inferedInstance.Content
		}
	case Table:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Content":
			res.valueString = inferedInstance.Content
		case "TableProperties":
			if inferedInstance.TableProperties != nil {
				res.valueString = inferedInstance.TableProperties.Name
			}
		case "TableRows":
			for idx, __instance__ := range inferedInstance.TableRows {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case TableColumn:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Paragraphs":
			for idx, __instance__ := range inferedInstance.Paragraphs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case TableProperties:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Content":
			res.valueString = inferedInstance.Content
		case "TableStyle":
			if inferedInstance.TableStyle != nil {
				res.valueString = inferedInstance.TableStyle.Name
			}
		}
	case TableRow:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "TableColumns":
			for idx, __instance__ := range inferedInstance.TableColumns {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case TableStyle:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "Content":
			res.valueString = inferedInstance.Content
		case "Val":
			res.valueString = inferedInstance.Val
		}
	case Text:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		case "Node":
			if inferedInstance.Node != nil {
				res.valueString = inferedInstance.Node.Name
			}
		case "PreserveWhiteSpace":
			res.valueString = fmt.Sprintf("%t", inferedInstance.PreserveWhiteSpace)
			res.valueBool = inferedInstance.PreserveWhiteSpace
			res.GongFieldValueType = GongFieldValueTypeBool
		case "EnclosingRune":
			if inferedInstance.EnclosingRune != nil {
				res.valueString = inferedInstance.EnclosingRune.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
