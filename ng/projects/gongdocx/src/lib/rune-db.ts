// insertion point for imports
import { NodeDB } from './node-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RuneDB {

	static GONGSTRUCT_NAME = "Rune"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	RuneStyle: string = ""

	// insertion point for other declarations
	Node?: NodeDB
	NodeID: NullInt64 = new NullInt64 // if pointer is null, Node.ID = 0

}
