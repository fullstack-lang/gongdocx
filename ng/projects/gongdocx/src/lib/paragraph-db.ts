// insertion point for imports
import { NodeDB } from './node-db'
import { ParagraphPropertiesDB } from './paragraphproperties-db'
import { RuneDB } from './rune-db'
import { TableColumnDB } from './tablecolumn-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ParagraphDB {

	static GONGSTRUCT_NAME = "Paragraph"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other declarations
	Node?: NodeDB
	NodeID: NullInt64 = new NullInt64 // if pointer is null, Node.ID = 0

	ParagraphProperties?: ParagraphPropertiesDB
	ParagraphPropertiesID: NullInt64 = new NullInt64 // if pointer is null, ParagraphProperties.ID = 0

	Runes?: Array<RuneDB>
	TableColumn_ParagraphsDBID: NullInt64 = new NullInt64
	TableColumn_ParagraphsDBID_Index: NullInt64  = new NullInt64 // store the index of the paragraph instance in TableColumn.Paragraphs
	TableColumn_Paragraphs_reverse?: TableColumnDB 

}
