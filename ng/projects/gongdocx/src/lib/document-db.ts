// insertion point for imports
import { FileDB } from './file-db'
import { NodeDB } from './node-db'
import { BodyDB } from './body-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DocumentDB {

	static GONGSTRUCT_NAME = "Document"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	File?: FileDB
	FileID: NullInt64 = new NullInt64 // if pointer is null, File.ID = 0

	Root?: NodeDB
	RootID: NullInt64 = new NullInt64 // if pointer is null, Root.ID = 0

	Body?: BodyDB
	BodyID: NullInt64 = new NullInt64 // if pointer is null, Body.ID = 0

}
