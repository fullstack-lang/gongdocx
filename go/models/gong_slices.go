// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Body:
		// insertion point per field
		if fieldName == "Paragraphs" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Body) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Body)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Paragraphs).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Paragraphs = _inferedTypeInstance.Paragraphs[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Paragraphs =
								append(_inferedTypeInstance.Paragraphs, any(fieldInstance).(*Paragraph))
						}
					}
				}
			}
		}
		if fieldName == "Tables" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Body) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Body)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Tables).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Tables = _inferedTypeInstance.Tables[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Tables =
								append(_inferedTypeInstance.Tables, any(fieldInstance).(*Table))
						}
					}
				}
			}
		}

	case *Document:
		// insertion point per field

	case *Docx:
		// insertion point per field
		if fieldName == "Files" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Docx) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Docx)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Files).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Files = _inferedTypeInstance.Files[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Files =
								append(_inferedTypeInstance.Files, any(fieldInstance).(*File))
						}
					}
				}
			}
		}

	case *File:
		// insertion point per field

	case *Node:
		// insertion point per field
		if fieldName == "Nodes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Node) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Node)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Nodes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Nodes = _inferedTypeInstance.Nodes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Nodes =
								append(_inferedTypeInstance.Nodes, any(fieldInstance).(*Node))
						}
					}
				}
			}
		}

	case *Paragraph:
		// insertion point per field
		if fieldName == "Runes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Paragraph) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Paragraph)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Runes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Runes = _inferedTypeInstance.Runes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Runes =
								append(_inferedTypeInstance.Runes, any(fieldInstance).(*Rune))
						}
					}
				}
			}
		}

	case *ParagraphProperties:
		// insertion point per field

	case *ParagraphStyle:
		// insertion point per field

	case *Rune:
		// insertion point per field

	case *RuneProperties:
		// insertion point per field

	case *Table:
		// insertion point per field
		if fieldName == "TableRows" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Table) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Table)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TableRows).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TableRows = _inferedTypeInstance.TableRows[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TableRows =
								append(_inferedTypeInstance.TableRows, any(fieldInstance).(*TableRow))
						}
					}
				}
			}
		}

	case *TableColumn:
		// insertion point per field
		if fieldName == "Paragraphs" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*TableColumn) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*TableColumn)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Paragraphs).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Paragraphs = _inferedTypeInstance.Paragraphs[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Paragraphs =
								append(_inferedTypeInstance.Paragraphs, any(fieldInstance).(*Paragraph))
						}
					}
				}
			}
		}

	case *TableProperties:
		// insertion point per field

	case *TableRow:
		// insertion point per field
		if fieldName == "TableColumns" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*TableRow) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*TableRow)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.TableColumns).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.TableColumns = _inferedTypeInstance.TableColumns[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.TableColumns =
								append(_inferedTypeInstance.TableColumns, any(fieldInstance).(*TableColumn))
						}
					}
				}
			}
		}

	case *TableStyle:
		// insertion point per field

	case *Text:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

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
