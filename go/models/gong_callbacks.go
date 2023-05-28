package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Document:
		if stage.OnAfterDocumentCreateCallback != nil {
			stage.OnAfterDocumentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Docx:
		if stage.OnAfterDocxCreateCallback != nil {
			stage.OnAfterDocxCreateCallback.OnAfterCreate(stage, target)
		}
	case *File:
		if stage.OnAfterFileCreateCallback != nil {
			stage.OnAfterFileCreateCallback.OnAfterCreate(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeCreateCallback != nil {
			stage.OnAfterNodeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Text:
		if stage.OnAfterTextCreateCallback != nil {
			stage.OnAfterTextCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Document:
		newTarget := any(new).(*Document)
		if stage.OnAfterDocumentUpdateCallback != nil {
			stage.OnAfterDocumentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Docx:
		newTarget := any(new).(*Docx)
		if stage.OnAfterDocxUpdateCallback != nil {
			stage.OnAfterDocxUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *File:
		newTarget := any(new).(*File)
		if stage.OnAfterFileUpdateCallback != nil {
			stage.OnAfterFileUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Node:
		newTarget := any(new).(*Node)
		if stage.OnAfterNodeUpdateCallback != nil {
			stage.OnAfterNodeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Text:
		newTarget := any(new).(*Text)
		if stage.OnAfterTextUpdateCallback != nil {
			stage.OnAfterTextUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Document:
		if stage.OnAfterDocumentDeleteCallback != nil {
			staged := any(staged).(*Document)
			stage.OnAfterDocumentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Docx:
		if stage.OnAfterDocxDeleteCallback != nil {
			staged := any(staged).(*Docx)
			stage.OnAfterDocxDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *File:
		if stage.OnAfterFileDeleteCallback != nil {
			staged := any(staged).(*File)
			stage.OnAfterFileDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Node:
		if stage.OnAfterNodeDeleteCallback != nil {
			staged := any(staged).(*Node)
			stage.OnAfterNodeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Text:
		if stage.OnAfterTextDeleteCallback != nil {
			staged := any(staged).(*Text)
			stage.OnAfterTextDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Document:
		if stage.OnAfterDocumentReadCallback != nil {
			stage.OnAfterDocumentReadCallback.OnAfterRead(stage, target)
		}
	case *Docx:
		if stage.OnAfterDocxReadCallback != nil {
			stage.OnAfterDocxReadCallback.OnAfterRead(stage, target)
		}
	case *File:
		if stage.OnAfterFileReadCallback != nil {
			stage.OnAfterFileReadCallback.OnAfterRead(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeReadCallback != nil {
			stage.OnAfterNodeReadCallback.OnAfterRead(stage, target)
		}
	case *Text:
		if stage.OnAfterTextReadCallback != nil {
			stage.OnAfterTextReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Document:
		stage.OnAfterDocumentUpdateCallback = any(callback).(OnAfterUpdateInterface[Document])
	
	case *Docx:
		stage.OnAfterDocxUpdateCallback = any(callback).(OnAfterUpdateInterface[Docx])
	
	case *File:
		stage.OnAfterFileUpdateCallback = any(callback).(OnAfterUpdateInterface[File])
	
	case *Node:
		stage.OnAfterNodeUpdateCallback = any(callback).(OnAfterUpdateInterface[Node])
	
	case *Text:
		stage.OnAfterTextUpdateCallback = any(callback).(OnAfterUpdateInterface[Text])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Document:
		stage.OnAfterDocumentCreateCallback = any(callback).(OnAfterCreateInterface[Document])
	
	case *Docx:
		stage.OnAfterDocxCreateCallback = any(callback).(OnAfterCreateInterface[Docx])
	
	case *File:
		stage.OnAfterFileCreateCallback = any(callback).(OnAfterCreateInterface[File])
	
	case *Node:
		stage.OnAfterNodeCreateCallback = any(callback).(OnAfterCreateInterface[Node])
	
	case *Text:
		stage.OnAfterTextCreateCallback = any(callback).(OnAfterCreateInterface[Text])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Document:
		stage.OnAfterDocumentDeleteCallback = any(callback).(OnAfterDeleteInterface[Document])
	
	case *Docx:
		stage.OnAfterDocxDeleteCallback = any(callback).(OnAfterDeleteInterface[Docx])
	
	case *File:
		stage.OnAfterFileDeleteCallback = any(callback).(OnAfterDeleteInterface[File])
	
	case *Node:
		stage.OnAfterNodeDeleteCallback = any(callback).(OnAfterDeleteInterface[Node])
	
	case *Text:
		stage.OnAfterTextDeleteCallback = any(callback).(OnAfterDeleteInterface[Text])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Document:
		stage.OnAfterDocumentReadCallback = any(callback).(OnAfterReadInterface[Document])
	
	case *Docx:
		stage.OnAfterDocxReadCallback = any(callback).(OnAfterReadInterface[Docx])
	
	case *File:
		stage.OnAfterFileReadCallback = any(callback).(OnAfterReadInterface[File])
	
	case *Node:
		stage.OnAfterNodeReadCallback = any(callback).(OnAfterReadInterface[Node])
	
	case *Text:
		stage.OnAfterTextReadCallback = any(callback).(OnAfterReadInterface[Text])
	
	}
}
