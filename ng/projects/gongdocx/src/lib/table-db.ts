// insertion point for imports
import { NodeDB } from './node-db'
import { TablePropertiesDB } from './tableproperties-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TableDB {

	static GONGSTRUCT_NAME = "Table"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other declarations
	Node?: NodeDB
	NodeID: NullInt64 = new NullInt64 // if pointer is null, Node.ID = 0

	TableProperties?: TablePropertiesDB
	TablePropertiesID: NullInt64 = new NullInt64 // if pointer is null, TableProperties.ID = 0

}
