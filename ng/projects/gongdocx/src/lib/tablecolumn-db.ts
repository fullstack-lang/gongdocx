// insertion point for imports
import { NodeDB } from './node-db'
import { ParagraphDB } from './paragraph-db'
import { TableRowDB } from './tablerow-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TableColumnDB {

	static GONGSTRUCT_NAME = "TableColumn"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other declarations
	Node?: NodeDB
	NodeID: NullInt64 = new NullInt64 // if pointer is null, Node.ID = 0

	Paragraphs?: Array<ParagraphDB>
	TableRow_TableColumnsDBID: NullInt64 = new NullInt64
	TableRow_TableColumnsDBID_Index: NullInt64  = new NullInt64 // store the index of the tablecolumn instance in TableRow.TableColumns
	TableRow_TableColumns_reverse?: TableRowDB 

}
