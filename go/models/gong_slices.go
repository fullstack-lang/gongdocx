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
