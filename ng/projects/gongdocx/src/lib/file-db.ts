// insertion point for imports
import { DocxDB } from './docx-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class FileDB {

	static GONGSTRUCT_NAME = "File"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Docx_FilesDBID: NullInt64 = new NullInt64
	Docx_FilesDBID_Index: NullInt64  = new NullInt64 // store the index of the file instance in Docx.Files
	Docx_Files_reverse?: DocxDB 

}
