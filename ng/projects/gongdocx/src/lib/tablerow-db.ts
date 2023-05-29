// insertion point for imports
import { NodeDB } from './node-db'
import { TableColumnDB } from './tablecolumn-db'
import { TableDB } from './table-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TableRowDB {

	static GONGSTRUCT_NAME = "TableRow"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other declarations
	Node?: NodeDB
	NodeID: NullInt64 = new NullInt64 // if pointer is null, Node.ID = 0

	TableColumns?: Array<TableColumnDB>
	Table_TableRowsDBID: NullInt64 = new NullInt64
	Table_TableRowsDBID_Index: NullInt64  = new NullInt64 // store the index of the tablerow instance in Table.TableRows
	Table_TableRows_reverse?: TableDB 

}
