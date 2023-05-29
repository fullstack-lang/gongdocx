// insertion point for imports
import { ParagraphStyleDB } from './paragraphstyle-db'
import { NodeDB } from './node-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ParagraphPropertiesDB {

	static GONGSTRUCT_NAME = "ParagraphProperties"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other declarations
	ParagraphStyle?: ParagraphStyleDB
	ParagraphStyleID: NullInt64 = new NullInt64 // if pointer is null, ParagraphStyle.ID = 0

	Node?: NodeDB
	NodeID: NullInt64 = new NullInt64 // if pointer is null, Node.ID = 0

}
