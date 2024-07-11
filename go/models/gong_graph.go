// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Body:
		ok = stage.IsStagedBody(target)

	case *Document:
		ok = stage.IsStagedDocument(target)

	case *Docx:
		ok = stage.IsStagedDocx(target)

	case *File:
		ok = stage.IsStagedFile(target)

	case *Node:
		ok = stage.IsStagedNode(target)

	case *Paragraph:
		ok = stage.IsStagedParagraph(target)

	case *ParagraphProperties:
		ok = stage.IsStagedParagraphProperties(target)

	case *ParagraphStyle:
		ok = stage.IsStagedParagraphStyle(target)

	case *Rune:
		ok = stage.IsStagedRune(target)

	case *RuneProperties:
		ok = stage.IsStagedRuneProperties(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	case *TableColumn:
		ok = stage.IsStagedTableColumn(target)

	case *TableProperties:
		ok = stage.IsStagedTableProperties(target)

	case *TableRow:
		ok = stage.IsStagedTableRow(target)

	case *TableStyle:
		ok = stage.IsStagedTableStyle(target)

	case *Text:
		ok = stage.IsStagedText(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedBody(body *Body) (ok bool) {

	_, ok = stage.Bodys[body]

	return
}

func (stage *StageStruct) IsStagedDocument(document *Document) (ok bool) {

	_, ok = stage.Documents[document]

	return
}

func (stage *StageStruct) IsStagedDocx(docx *Docx) (ok bool) {

	_, ok = stage.Docxs[docx]

	return
}

func (stage *StageStruct) IsStagedFile(file *File) (ok bool) {

	_, ok = stage.Files[file]

	return
}

func (stage *StageStruct) IsStagedNode(node *Node) (ok bool) {

	_, ok = stage.Nodes[node]

	return
}

func (stage *StageStruct) IsStagedParagraph(paragraph *Paragraph) (ok bool) {

	_, ok = stage.Paragraphs[paragraph]

	return
}

func (stage *StageStruct) IsStagedParagraphProperties(paragraphproperties *ParagraphProperties) (ok bool) {

	_, ok = stage.ParagraphPropertiess[paragraphproperties]

	return
}

func (stage *StageStruct) IsStagedParagraphStyle(paragraphstyle *ParagraphStyle) (ok bool) {

	_, ok = stage.ParagraphStyles[paragraphstyle]

	return
}

func (stage *StageStruct) IsStagedRune(rune *Rune) (ok bool) {

	_, ok = stage.Runes[rune]

	return
}

func (stage *StageStruct) IsStagedRuneProperties(runeproperties *RuneProperties) (ok bool) {

	_, ok = stage.RunePropertiess[runeproperties]

	return
}

func (stage *StageStruct) IsStagedTable(table *Table) (ok bool) {

	_, ok = stage.Tables[table]

	return
}

func (stage *StageStruct) IsStagedTableColumn(tablecolumn *TableColumn) (ok bool) {

	_, ok = stage.TableColumns[tablecolumn]

	return
}

func (stage *StageStruct) IsStagedTableProperties(tableproperties *TableProperties) (ok bool) {

	_, ok = stage.TablePropertiess[tableproperties]

	return
}

func (stage *StageStruct) IsStagedTableRow(tablerow *TableRow) (ok bool) {

	_, ok = stage.TableRows[tablerow]

	return
}

func (stage *StageStruct) IsStagedTableStyle(tablestyle *TableStyle) (ok bool) {

	_, ok = stage.TableStyles[tablestyle]

	return
}

func (stage *StageStruct) IsStagedText(text *Text) (ok bool) {

	_, ok = stage.Texts[text]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Body:
		stage.StageBranchBody(target)

	case *Document:
		stage.StageBranchDocument(target)

	case *Docx:
		stage.StageBranchDocx(target)

	case *File:
		stage.StageBranchFile(target)

	case *Node:
		stage.StageBranchNode(target)

	case *Paragraph:
		stage.StageBranchParagraph(target)

	case *ParagraphProperties:
		stage.StageBranchParagraphProperties(target)

	case *ParagraphStyle:
		stage.StageBranchParagraphStyle(target)

	case *Rune:
		stage.StageBranchRune(target)

	case *RuneProperties:
		stage.StageBranchRuneProperties(target)

	case *Table:
		stage.StageBranchTable(target)

	case *TableColumn:
		stage.StageBranchTableColumn(target)

	case *TableProperties:
		stage.StageBranchTableProperties(target)

	case *TableRow:
		stage.StageBranchTableRow(target)

	case *TableStyle:
		stage.StageBranchTableStyle(target)

	case *Text:
		stage.StageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchBody(body *Body) {

	// check if instance is already staged
	if IsStaged(stage, body) {
		return
	}

	body.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if body.LastParagraph != nil {
		StageBranch(stage, body.LastParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range body.Paragraphs {
		StageBranch(stage, _paragraph)
	}
	for _, _table := range body.Tables {
		StageBranch(stage, _table)
	}

}

func (stage *StageStruct) StageBranchDocument(document *Document) {

	// check if instance is already staged
	if IsStaged(stage, document) {
		return
	}

	document.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if document.File != nil {
		StageBranch(stage, document.File)
	}
	if document.Root != nil {
		StageBranch(stage, document.Root)
	}
	if document.Body != nil {
		StageBranch(stage, document.Body)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDocx(docx *Docx) {

	// check if instance is already staged
	if IsStaged(stage, docx) {
		return
	}

	docx.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if docx.Document != nil {
		StageBranch(stage, docx.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _file := range docx.Files {
		StageBranch(stage, _file)
	}

}

func (stage *StageStruct) StageBranchFile(file *File) {

	// check if instance is already staged
	if IsStaged(stage, file) {
		return
	}

	file.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNode(node *Node) {

	// check if instance is already staged
	if IsStaged(stage, node) {
		return
	}

	node.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Nodes {
		StageBranch(stage, _node)
	}

}

func (stage *StageStruct) StageBranchParagraph(paragraph *Paragraph) {

	// check if instance is already staged
	if IsStaged(stage, paragraph) {
		return
	}

	paragraph.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraph.Node != nil {
		StageBranch(stage, paragraph.Node)
	}
	if paragraph.ParagraphProperties != nil {
		StageBranch(stage, paragraph.ParagraphProperties)
	}
	if paragraph.Next != nil {
		StageBranch(stage, paragraph.Next)
	}
	if paragraph.Previous != nil {
		StageBranch(stage, paragraph.Previous)
	}
	if paragraph.EnclosingBody != nil {
		StageBranch(stage, paragraph.EnclosingBody)
	}
	if paragraph.EnclosingTableColumn != nil {
		StageBranch(stage, paragraph.EnclosingTableColumn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rune := range paragraph.Runes {
		StageBranch(stage, _rune)
	}

}

func (stage *StageStruct) StageBranchParagraphProperties(paragraphproperties *ParagraphProperties) {

	// check if instance is already staged
	if IsStaged(stage, paragraphproperties) {
		return
	}

	paragraphproperties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphproperties.ParagraphStyle != nil {
		StageBranch(stage, paragraphproperties.ParagraphStyle)
	}
	if paragraphproperties.Node != nil {
		StageBranch(stage, paragraphproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchParagraphStyle(paragraphstyle *ParagraphStyle) {

	// check if instance is already staged
	if IsStaged(stage, paragraphstyle) {
		return
	}

	paragraphstyle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyle.Node != nil {
		StageBranch(stage, paragraphstyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRune(rune *Rune) {

	// check if instance is already staged
	if IsStaged(stage, rune) {
		return
	}

	rune.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rune.Node != nil {
		StageBranch(stage, rune.Node)
	}
	if rune.Text != nil {
		StageBranch(stage, rune.Text)
	}
	if rune.RuneProperties != nil {
		StageBranch(stage, rune.RuneProperties)
	}
	if rune.EnclosingParagraph != nil {
		StageBranch(stage, rune.EnclosingParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRuneProperties(runeproperties *RuneProperties) {

	// check if instance is already staged
	if IsStaged(stage, runeproperties) {
		return
	}

	runeproperties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if runeproperties.Node != nil {
		StageBranch(stage, runeproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTable(table *Table) {

	// check if instance is already staged
	if IsStaged(stage, table) {
		return
	}

	table.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if table.Node != nil {
		StageBranch(stage, table.Node)
	}
	if table.TableProperties != nil {
		StageBranch(stage, table.TableProperties)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablerow := range table.TableRows {
		StageBranch(stage, _tablerow)
	}

}

func (stage *StageStruct) StageBranchTableColumn(tablecolumn *TableColumn) {

	// check if instance is already staged
	if IsStaged(stage, tablecolumn) {
		return
	}

	tablecolumn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablecolumn.Node != nil {
		StageBranch(stage, tablecolumn.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range tablecolumn.Paragraphs {
		StageBranch(stage, _paragraph)
	}

}

func (stage *StageStruct) StageBranchTableProperties(tableproperties *TableProperties) {

	// check if instance is already staged
	if IsStaged(stage, tableproperties) {
		return
	}

	tableproperties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tableproperties.Node != nil {
		StageBranch(stage, tableproperties.Node)
	}
	if tableproperties.TableStyle != nil {
		StageBranch(stage, tableproperties.TableStyle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTableRow(tablerow *TableRow) {

	// check if instance is already staged
	if IsStaged(stage, tablerow) {
		return
	}

	tablerow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablerow.Node != nil {
		StageBranch(stage, tablerow.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablecolumn := range tablerow.TableColumns {
		StageBranch(stage, _tablecolumn)
	}

}

func (stage *StageStruct) StageBranchTableStyle(tablestyle *TableStyle) {

	// check if instance is already staged
	if IsStaged(stage, tablestyle) {
		return
	}

	tablestyle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablestyle.Node != nil {
		StageBranch(stage, tablestyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchText(text *Text) {

	// check if instance is already staged
	if IsStaged(stage, text) {
		return
	}

	text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if text.Node != nil {
		StageBranch(stage, text.Node)
	}
	if text.EnclosingRune != nil {
		StageBranch(stage, text.EnclosingRune)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Body:
		toT := CopyBranchBody(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Document:
		toT := CopyBranchDocument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Docx:
		toT := CopyBranchDocx(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *File:
		toT := CopyBranchFile(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Node:
		toT := CopyBranchNode(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Paragraph:
		toT := CopyBranchParagraph(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParagraphProperties:
		toT := CopyBranchParagraphProperties(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParagraphStyle:
		toT := CopyBranchParagraphStyle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rune:
		toT := CopyBranchRune(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RuneProperties:
		toT := CopyBranchRuneProperties(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Table:
		toT := CopyBranchTable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableColumn:
		toT := CopyBranchTableColumn(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableProperties:
		toT := CopyBranchTableProperties(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableRow:
		toT := CopyBranchTableRow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableStyle:
		toT := CopyBranchTableStyle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Text:
		toT := CopyBranchText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchBody(mapOrigCopy map[any]any, bodyFrom *Body) (bodyTo *Body) {

	// bodyFrom has already been copied
	if _bodyTo, ok := mapOrigCopy[bodyFrom]; ok {
		bodyTo = _bodyTo.(*Body)
		return
	}

	bodyTo = new(Body)
	mapOrigCopy[bodyFrom] = bodyTo
	bodyFrom.CopyBasicFields(bodyTo)

	//insertion point for the staging of instances referenced by pointers
	if bodyFrom.LastParagraph != nil {
		bodyTo.LastParagraph = CopyBranchParagraph(mapOrigCopy, bodyFrom.LastParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range bodyFrom.Paragraphs {
		bodyTo.Paragraphs = append(bodyTo.Paragraphs, CopyBranchParagraph(mapOrigCopy, _paragraph))
	}
	for _, _table := range bodyFrom.Tables {
		bodyTo.Tables = append(bodyTo.Tables, CopyBranchTable(mapOrigCopy, _table))
	}

	return
}

func CopyBranchDocument(mapOrigCopy map[any]any, documentFrom *Document) (documentTo *Document) {

	// documentFrom has already been copied
	if _documentTo, ok := mapOrigCopy[documentFrom]; ok {
		documentTo = _documentTo.(*Document)
		return
	}

	documentTo = new(Document)
	mapOrigCopy[documentFrom] = documentTo
	documentFrom.CopyBasicFields(documentTo)

	//insertion point for the staging of instances referenced by pointers
	if documentFrom.File != nil {
		documentTo.File = CopyBranchFile(mapOrigCopy, documentFrom.File)
	}
	if documentFrom.Root != nil {
		documentTo.Root = CopyBranchNode(mapOrigCopy, documentFrom.Root)
	}
	if documentFrom.Body != nil {
		documentTo.Body = CopyBranchBody(mapOrigCopy, documentFrom.Body)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDocx(mapOrigCopy map[any]any, docxFrom *Docx) (docxTo *Docx) {

	// docxFrom has already been copied
	if _docxTo, ok := mapOrigCopy[docxFrom]; ok {
		docxTo = _docxTo.(*Docx)
		return
	}

	docxTo = new(Docx)
	mapOrigCopy[docxFrom] = docxTo
	docxFrom.CopyBasicFields(docxTo)

	//insertion point for the staging of instances referenced by pointers
	if docxFrom.Document != nil {
		docxTo.Document = CopyBranchDocument(mapOrigCopy, docxFrom.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _file := range docxFrom.Files {
		docxTo.Files = append(docxTo.Files, CopyBranchFile(mapOrigCopy, _file))
	}

	return
}

func CopyBranchFile(mapOrigCopy map[any]any, fileFrom *File) (fileTo *File) {

	// fileFrom has already been copied
	if _fileTo, ok := mapOrigCopy[fileFrom]; ok {
		fileTo = _fileTo.(*File)
		return
	}

	fileTo = new(File)
	mapOrigCopy[fileFrom] = fileTo
	fileFrom.CopyBasicFields(fileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNode(mapOrigCopy map[any]any, nodeFrom *Node) (nodeTo *Node) {

	// nodeFrom has already been copied
	if _nodeTo, ok := mapOrigCopy[nodeFrom]; ok {
		nodeTo = _nodeTo.(*Node)
		return
	}

	nodeTo = new(Node)
	mapOrigCopy[nodeFrom] = nodeTo
	nodeFrom.CopyBasicFields(nodeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range nodeFrom.Nodes {
		nodeTo.Nodes = append(nodeTo.Nodes, CopyBranchNode(mapOrigCopy, _node))
	}

	return
}

func CopyBranchParagraph(mapOrigCopy map[any]any, paragraphFrom *Paragraph) (paragraphTo *Paragraph) {

	// paragraphFrom has already been copied
	if _paragraphTo, ok := mapOrigCopy[paragraphFrom]; ok {
		paragraphTo = _paragraphTo.(*Paragraph)
		return
	}

	paragraphTo = new(Paragraph)
	mapOrigCopy[paragraphFrom] = paragraphTo
	paragraphFrom.CopyBasicFields(paragraphTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphFrom.Node != nil {
		paragraphTo.Node = CopyBranchNode(mapOrigCopy, paragraphFrom.Node)
	}
	if paragraphFrom.ParagraphProperties != nil {
		paragraphTo.ParagraphProperties = CopyBranchParagraphProperties(mapOrigCopy, paragraphFrom.ParagraphProperties)
	}
	if paragraphFrom.Next != nil {
		paragraphTo.Next = CopyBranchParagraph(mapOrigCopy, paragraphFrom.Next)
	}
	if paragraphFrom.Previous != nil {
		paragraphTo.Previous = CopyBranchParagraph(mapOrigCopy, paragraphFrom.Previous)
	}
	if paragraphFrom.EnclosingBody != nil {
		paragraphTo.EnclosingBody = CopyBranchBody(mapOrigCopy, paragraphFrom.EnclosingBody)
	}
	if paragraphFrom.EnclosingTableColumn != nil {
		paragraphTo.EnclosingTableColumn = CopyBranchTableColumn(mapOrigCopy, paragraphFrom.EnclosingTableColumn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rune := range paragraphFrom.Runes {
		paragraphTo.Runes = append(paragraphTo.Runes, CopyBranchRune(mapOrigCopy, _rune))
	}

	return
}

func CopyBranchParagraphProperties(mapOrigCopy map[any]any, paragraphpropertiesFrom *ParagraphProperties) (paragraphpropertiesTo *ParagraphProperties) {

	// paragraphpropertiesFrom has already been copied
	if _paragraphpropertiesTo, ok := mapOrigCopy[paragraphpropertiesFrom]; ok {
		paragraphpropertiesTo = _paragraphpropertiesTo.(*ParagraphProperties)
		return
	}

	paragraphpropertiesTo = new(ParagraphProperties)
	mapOrigCopy[paragraphpropertiesFrom] = paragraphpropertiesTo
	paragraphpropertiesFrom.CopyBasicFields(paragraphpropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphpropertiesFrom.ParagraphStyle != nil {
		paragraphpropertiesTo.ParagraphStyle = CopyBranchParagraphStyle(mapOrigCopy, paragraphpropertiesFrom.ParagraphStyle)
	}
	if paragraphpropertiesFrom.Node != nil {
		paragraphpropertiesTo.Node = CopyBranchNode(mapOrigCopy, paragraphpropertiesFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchParagraphStyle(mapOrigCopy map[any]any, paragraphstyleFrom *ParagraphStyle) (paragraphstyleTo *ParagraphStyle) {

	// paragraphstyleFrom has already been copied
	if _paragraphstyleTo, ok := mapOrigCopy[paragraphstyleFrom]; ok {
		paragraphstyleTo = _paragraphstyleTo.(*ParagraphStyle)
		return
	}

	paragraphstyleTo = new(ParagraphStyle)
	mapOrigCopy[paragraphstyleFrom] = paragraphstyleTo
	paragraphstyleFrom.CopyBasicFields(paragraphstyleTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyleFrom.Node != nil {
		paragraphstyleTo.Node = CopyBranchNode(mapOrigCopy, paragraphstyleFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRune(mapOrigCopy map[any]any, runeFrom *Rune) (runeTo *Rune) {

	// runeFrom has already been copied
	if _runeTo, ok := mapOrigCopy[runeFrom]; ok {
		runeTo = _runeTo.(*Rune)
		return
	}

	runeTo = new(Rune)
	mapOrigCopy[runeFrom] = runeTo
	runeFrom.CopyBasicFields(runeTo)

	//insertion point for the staging of instances referenced by pointers
	if runeFrom.Node != nil {
		runeTo.Node = CopyBranchNode(mapOrigCopy, runeFrom.Node)
	}
	if runeFrom.Text != nil {
		runeTo.Text = CopyBranchText(mapOrigCopy, runeFrom.Text)
	}
	if runeFrom.RuneProperties != nil {
		runeTo.RuneProperties = CopyBranchRuneProperties(mapOrigCopy, runeFrom.RuneProperties)
	}
	if runeFrom.EnclosingParagraph != nil {
		runeTo.EnclosingParagraph = CopyBranchParagraph(mapOrigCopy, runeFrom.EnclosingParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRuneProperties(mapOrigCopy map[any]any, runepropertiesFrom *RuneProperties) (runepropertiesTo *RuneProperties) {

	// runepropertiesFrom has already been copied
	if _runepropertiesTo, ok := mapOrigCopy[runepropertiesFrom]; ok {
		runepropertiesTo = _runepropertiesTo.(*RuneProperties)
		return
	}

	runepropertiesTo = new(RuneProperties)
	mapOrigCopy[runepropertiesFrom] = runepropertiesTo
	runepropertiesFrom.CopyBasicFields(runepropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if runepropertiesFrom.Node != nil {
		runepropertiesTo.Node = CopyBranchNode(mapOrigCopy, runepropertiesFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTable(mapOrigCopy map[any]any, tableFrom *Table) (tableTo *Table) {

	// tableFrom has already been copied
	if _tableTo, ok := mapOrigCopy[tableFrom]; ok {
		tableTo = _tableTo.(*Table)
		return
	}

	tableTo = new(Table)
	mapOrigCopy[tableFrom] = tableTo
	tableFrom.CopyBasicFields(tableTo)

	//insertion point for the staging of instances referenced by pointers
	if tableFrom.Node != nil {
		tableTo.Node = CopyBranchNode(mapOrigCopy, tableFrom.Node)
	}
	if tableFrom.TableProperties != nil {
		tableTo.TableProperties = CopyBranchTableProperties(mapOrigCopy, tableFrom.TableProperties)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablerow := range tableFrom.TableRows {
		tableTo.TableRows = append(tableTo.TableRows, CopyBranchTableRow(mapOrigCopy, _tablerow))
	}

	return
}

func CopyBranchTableColumn(mapOrigCopy map[any]any, tablecolumnFrom *TableColumn) (tablecolumnTo *TableColumn) {

	// tablecolumnFrom has already been copied
	if _tablecolumnTo, ok := mapOrigCopy[tablecolumnFrom]; ok {
		tablecolumnTo = _tablecolumnTo.(*TableColumn)
		return
	}

	tablecolumnTo = new(TableColumn)
	mapOrigCopy[tablecolumnFrom] = tablecolumnTo
	tablecolumnFrom.CopyBasicFields(tablecolumnTo)

	//insertion point for the staging of instances referenced by pointers
	if tablecolumnFrom.Node != nil {
		tablecolumnTo.Node = CopyBranchNode(mapOrigCopy, tablecolumnFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range tablecolumnFrom.Paragraphs {
		tablecolumnTo.Paragraphs = append(tablecolumnTo.Paragraphs, CopyBranchParagraph(mapOrigCopy, _paragraph))
	}

	return
}

func CopyBranchTableProperties(mapOrigCopy map[any]any, tablepropertiesFrom *TableProperties) (tablepropertiesTo *TableProperties) {

	// tablepropertiesFrom has already been copied
	if _tablepropertiesTo, ok := mapOrigCopy[tablepropertiesFrom]; ok {
		tablepropertiesTo = _tablepropertiesTo.(*TableProperties)
		return
	}

	tablepropertiesTo = new(TableProperties)
	mapOrigCopy[tablepropertiesFrom] = tablepropertiesTo
	tablepropertiesFrom.CopyBasicFields(tablepropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if tablepropertiesFrom.Node != nil {
		tablepropertiesTo.Node = CopyBranchNode(mapOrigCopy, tablepropertiesFrom.Node)
	}
	if tablepropertiesFrom.TableStyle != nil {
		tablepropertiesTo.TableStyle = CopyBranchTableStyle(mapOrigCopy, tablepropertiesFrom.TableStyle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTableRow(mapOrigCopy map[any]any, tablerowFrom *TableRow) (tablerowTo *TableRow) {

	// tablerowFrom has already been copied
	if _tablerowTo, ok := mapOrigCopy[tablerowFrom]; ok {
		tablerowTo = _tablerowTo.(*TableRow)
		return
	}

	tablerowTo = new(TableRow)
	mapOrigCopy[tablerowFrom] = tablerowTo
	tablerowFrom.CopyBasicFields(tablerowTo)

	//insertion point for the staging of instances referenced by pointers
	if tablerowFrom.Node != nil {
		tablerowTo.Node = CopyBranchNode(mapOrigCopy, tablerowFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablecolumn := range tablerowFrom.TableColumns {
		tablerowTo.TableColumns = append(tablerowTo.TableColumns, CopyBranchTableColumn(mapOrigCopy, _tablecolumn))
	}

	return
}

func CopyBranchTableStyle(mapOrigCopy map[any]any, tablestyleFrom *TableStyle) (tablestyleTo *TableStyle) {

	// tablestyleFrom has already been copied
	if _tablestyleTo, ok := mapOrigCopy[tablestyleFrom]; ok {
		tablestyleTo = _tablestyleTo.(*TableStyle)
		return
	}

	tablestyleTo = new(TableStyle)
	mapOrigCopy[tablestyleFrom] = tablestyleTo
	tablestyleFrom.CopyBasicFields(tablestyleTo)

	//insertion point for the staging of instances referenced by pointers
	if tablestyleFrom.Node != nil {
		tablestyleTo.Node = CopyBranchNode(mapOrigCopy, tablestyleFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchText(mapOrigCopy map[any]any, textFrom *Text) (textTo *Text) {

	// textFrom has already been copied
	if _textTo, ok := mapOrigCopy[textFrom]; ok {
		textTo = _textTo.(*Text)
		return
	}

	textTo = new(Text)
	mapOrigCopy[textFrom] = textTo
	textFrom.CopyBasicFields(textTo)

	//insertion point for the staging of instances referenced by pointers
	if textFrom.Node != nil {
		textTo.Node = CopyBranchNode(mapOrigCopy, textFrom.Node)
	}
	if textFrom.EnclosingRune != nil {
		textTo.EnclosingRune = CopyBranchRune(mapOrigCopy, textFrom.EnclosingRune)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Body:
		stage.UnstageBranchBody(target)

	case *Document:
		stage.UnstageBranchDocument(target)

	case *Docx:
		stage.UnstageBranchDocx(target)

	case *File:
		stage.UnstageBranchFile(target)

	case *Node:
		stage.UnstageBranchNode(target)

	case *Paragraph:
		stage.UnstageBranchParagraph(target)

	case *ParagraphProperties:
		stage.UnstageBranchParagraphProperties(target)

	case *ParagraphStyle:
		stage.UnstageBranchParagraphStyle(target)

	case *Rune:
		stage.UnstageBranchRune(target)

	case *RuneProperties:
		stage.UnstageBranchRuneProperties(target)

	case *Table:
		stage.UnstageBranchTable(target)

	case *TableColumn:
		stage.UnstageBranchTableColumn(target)

	case *TableProperties:
		stage.UnstageBranchTableProperties(target)

	case *TableRow:
		stage.UnstageBranchTableRow(target)

	case *TableStyle:
		stage.UnstageBranchTableStyle(target)

	case *Text:
		stage.UnstageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchBody(body *Body) {

	// check if instance is already staged
	if !IsStaged(stage, body) {
		return
	}

	body.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if body.LastParagraph != nil {
		UnstageBranch(stage, body.LastParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range body.Paragraphs {
		UnstageBranch(stage, _paragraph)
	}
	for _, _table := range body.Tables {
		UnstageBranch(stage, _table)
	}

}

func (stage *StageStruct) UnstageBranchDocument(document *Document) {

	// check if instance is already staged
	if !IsStaged(stage, document) {
		return
	}

	document.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if document.File != nil {
		UnstageBranch(stage, document.File)
	}
	if document.Root != nil {
		UnstageBranch(stage, document.Root)
	}
	if document.Body != nil {
		UnstageBranch(stage, document.Body)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDocx(docx *Docx) {

	// check if instance is already staged
	if !IsStaged(stage, docx) {
		return
	}

	docx.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if docx.Document != nil {
		UnstageBranch(stage, docx.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _file := range docx.Files {
		UnstageBranch(stage, _file)
	}

}

func (stage *StageStruct) UnstageBranchFile(file *File) {

	// check if instance is already staged
	if !IsStaged(stage, file) {
		return
	}

	file.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNode(node *Node) {

	// check if instance is already staged
	if !IsStaged(stage, node) {
		return
	}

	node.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Nodes {
		UnstageBranch(stage, _node)
	}

}

func (stage *StageStruct) UnstageBranchParagraph(paragraph *Paragraph) {

	// check if instance is already staged
	if !IsStaged(stage, paragraph) {
		return
	}

	paragraph.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraph.Node != nil {
		UnstageBranch(stage, paragraph.Node)
	}
	if paragraph.ParagraphProperties != nil {
		UnstageBranch(stage, paragraph.ParagraphProperties)
	}
	if paragraph.Next != nil {
		UnstageBranch(stage, paragraph.Next)
	}
	if paragraph.Previous != nil {
		UnstageBranch(stage, paragraph.Previous)
	}
	if paragraph.EnclosingBody != nil {
		UnstageBranch(stage, paragraph.EnclosingBody)
	}
	if paragraph.EnclosingTableColumn != nil {
		UnstageBranch(stage, paragraph.EnclosingTableColumn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rune := range paragraph.Runes {
		UnstageBranch(stage, _rune)
	}

}

func (stage *StageStruct) UnstageBranchParagraphProperties(paragraphproperties *ParagraphProperties) {

	// check if instance is already staged
	if !IsStaged(stage, paragraphproperties) {
		return
	}

	paragraphproperties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphproperties.ParagraphStyle != nil {
		UnstageBranch(stage, paragraphproperties.ParagraphStyle)
	}
	if paragraphproperties.Node != nil {
		UnstageBranch(stage, paragraphproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchParagraphStyle(paragraphstyle *ParagraphStyle) {

	// check if instance is already staged
	if !IsStaged(stage, paragraphstyle) {
		return
	}

	paragraphstyle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyle.Node != nil {
		UnstageBranch(stage, paragraphstyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRune(rune *Rune) {

	// check if instance is already staged
	if !IsStaged(stage, rune) {
		return
	}

	rune.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rune.Node != nil {
		UnstageBranch(stage, rune.Node)
	}
	if rune.Text != nil {
		UnstageBranch(stage, rune.Text)
	}
	if rune.RuneProperties != nil {
		UnstageBranch(stage, rune.RuneProperties)
	}
	if rune.EnclosingParagraph != nil {
		UnstageBranch(stage, rune.EnclosingParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRuneProperties(runeproperties *RuneProperties) {

	// check if instance is already staged
	if !IsStaged(stage, runeproperties) {
		return
	}

	runeproperties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if runeproperties.Node != nil {
		UnstageBranch(stage, runeproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTable(table *Table) {

	// check if instance is already staged
	if !IsStaged(stage, table) {
		return
	}

	table.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if table.Node != nil {
		UnstageBranch(stage, table.Node)
	}
	if table.TableProperties != nil {
		UnstageBranch(stage, table.TableProperties)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablerow := range table.TableRows {
		UnstageBranch(stage, _tablerow)
	}

}

func (stage *StageStruct) UnstageBranchTableColumn(tablecolumn *TableColumn) {

	// check if instance is already staged
	if !IsStaged(stage, tablecolumn) {
		return
	}

	tablecolumn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablecolumn.Node != nil {
		UnstageBranch(stage, tablecolumn.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range tablecolumn.Paragraphs {
		UnstageBranch(stage, _paragraph)
	}

}

func (stage *StageStruct) UnstageBranchTableProperties(tableproperties *TableProperties) {

	// check if instance is already staged
	if !IsStaged(stage, tableproperties) {
		return
	}

	tableproperties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tableproperties.Node != nil {
		UnstageBranch(stage, tableproperties.Node)
	}
	if tableproperties.TableStyle != nil {
		UnstageBranch(stage, tableproperties.TableStyle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTableRow(tablerow *TableRow) {

	// check if instance is already staged
	if !IsStaged(stage, tablerow) {
		return
	}

	tablerow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablerow.Node != nil {
		UnstageBranch(stage, tablerow.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablecolumn := range tablerow.TableColumns {
		UnstageBranch(stage, _tablecolumn)
	}

}

func (stage *StageStruct) UnstageBranchTableStyle(tablestyle *TableStyle) {

	// check if instance is already staged
	if !IsStaged(stage, tablestyle) {
		return
	}

	tablestyle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablestyle.Node != nil {
		UnstageBranch(stage, tablestyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchText(text *Text) {

	// check if instance is already staged
	if !IsStaged(stage, text) {
		return
	}

	text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if text.Node != nil {
		UnstageBranch(stage, text.Node)
	}
	if text.EnclosingRune != nil {
		UnstageBranch(stage, text.EnclosingRune)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}
