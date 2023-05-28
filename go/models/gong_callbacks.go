package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Docx:
		if stage.OnAfterDocxCreateCallback != nil {
			stage.OnAfterDocxCreateCallback.OnAfterCreate(stage, target)
		}
	case *File:
		if stage.OnAfterFileCreateCallback != nil {
			stage.OnAfterFileCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
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
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
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
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Docx:
		if stage.OnAfterDocxReadCallback != nil {
			stage.OnAfterDocxReadCallback.OnAfterRead(stage, target)
		}
	case *File:
		if stage.OnAfterFileReadCallback != nil {
			stage.OnAfterFileReadCallback.OnAfterRead(stage, target)
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
	case *Docx:
		stage.OnAfterDocxUpdateCallback = any(callback).(OnAfterUpdateInterface[Docx])
	
	case *File:
		stage.OnAfterFileUpdateCallback = any(callback).(OnAfterUpdateInterface[File])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Docx:
		stage.OnAfterDocxCreateCallback = any(callback).(OnAfterCreateInterface[Docx])
	
	case *File:
		stage.OnAfterFileCreateCallback = any(callback).(OnAfterCreateInterface[File])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Docx:
		stage.OnAfterDocxDeleteCallback = any(callback).(OnAfterDeleteInterface[Docx])
	
	case *File:
		stage.OnAfterFileDeleteCallback = any(callback).(OnAfterDeleteInterface[File])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Docx:
		stage.OnAfterDocxReadCallback = any(callback).(OnAfterReadInterface[Docx])
	
	case *File:
		stage.OnAfterFileReadCallback = any(callback).(OnAfterReadInterface[File])
	
	}
}
