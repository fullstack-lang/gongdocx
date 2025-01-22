// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Body
	// insertion point per field
	clear(stage.Body_Paragraphs_reverseMap)
	stage.Body_Paragraphs_reverseMap = make(map[*Paragraph]*Body)
	for body := range stage.Bodys {
		_ = body
		for _, _paragraph := range body.Paragraphs {
			stage.Body_Paragraphs_reverseMap[_paragraph] = body
		}
	}
	clear(stage.Body_Tables_reverseMap)
	stage.Body_Tables_reverseMap = make(map[*Table]*Body)
	for body := range stage.Bodys {
		_ = body
		for _, _table := range body.Tables {
			stage.Body_Tables_reverseMap[_table] = body
		}
	}

	// Compute reverse map for named struct Document
	// insertion point per field

	// Compute reverse map for named struct Docx
	// insertion point per field
	clear(stage.Docx_Files_reverseMap)
	stage.Docx_Files_reverseMap = make(map[*File]*Docx)
	for docx := range stage.Docxs {
		_ = docx
		for _, _file := range docx.Files {
			stage.Docx_Files_reverseMap[_file] = docx
		}
	}

	// Compute reverse map for named struct File
	// insertion point per field

	// Compute reverse map for named struct Node
	// insertion point per field
	clear(stage.Node_Nodes_reverseMap)
	stage.Node_Nodes_reverseMap = make(map[*Node]*Node)
	for node := range stage.Nodes {
		_ = node
		for _, _node := range node.Nodes {
			stage.Node_Nodes_reverseMap[_node] = node
		}
	}

	// Compute reverse map for named struct Paragraph
	// insertion point per field
	clear(stage.Paragraph_Runes_reverseMap)
	stage.Paragraph_Runes_reverseMap = make(map[*Rune]*Paragraph)
	for paragraph := range stage.Paragraphs {
		_ = paragraph
		for _, _rune := range paragraph.Runes {
			stage.Paragraph_Runes_reverseMap[_rune] = paragraph
		}
	}

	// Compute reverse map for named struct ParagraphProperties
	// insertion point per field

	// Compute reverse map for named struct ParagraphStyle
	// insertion point per field

	// Compute reverse map for named struct Rune
	// insertion point per field

	// Compute reverse map for named struct RuneProperties
	// insertion point per field

	// Compute reverse map for named struct Table
	// insertion point per field
	clear(stage.Table_TableRows_reverseMap)
	stage.Table_TableRows_reverseMap = make(map[*TableRow]*Table)
	for table := range stage.Tables {
		_ = table
		for _, _tablerow := range table.TableRows {
			stage.Table_TableRows_reverseMap[_tablerow] = table
		}
	}

	// Compute reverse map for named struct TableColumn
	// insertion point per field
	clear(stage.TableColumn_Paragraphs_reverseMap)
	stage.TableColumn_Paragraphs_reverseMap = make(map[*Paragraph]*TableColumn)
	for tablecolumn := range stage.TableColumns {
		_ = tablecolumn
		for _, _paragraph := range tablecolumn.Paragraphs {
			stage.TableColumn_Paragraphs_reverseMap[_paragraph] = tablecolumn
		}
	}

	// Compute reverse map for named struct TableProperties
	// insertion point per field

	// Compute reverse map for named struct TableRow
	// insertion point per field
	clear(stage.TableRow_TableColumns_reverseMap)
	stage.TableRow_TableColumns_reverseMap = make(map[*TableColumn]*TableRow)
	for tablerow := range stage.TableRows {
		_ = tablerow
		for _, _tablecolumn := range tablerow.TableColumns {
			stage.TableRow_TableColumns_reverseMap[_tablecolumn] = tablerow
		}
	}

	// Compute reverse map for named struct TableStyle
	// insertion point per field

	// Compute reverse map for named struct Text
	// insertion point per field

}
