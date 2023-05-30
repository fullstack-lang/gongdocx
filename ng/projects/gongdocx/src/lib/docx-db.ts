// insertion point for imports
import { FileDB } from './file-db'
import { DocumentDB } from './document-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DocxDB {

	static GONGSTRUCT_NAME = "Docx"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Files?: Array<FileDB>
	Document?: DocumentDB
	DocumentID: NullInt64 = new NullInt64 // if pointer is null, Document.ID = 0

}
