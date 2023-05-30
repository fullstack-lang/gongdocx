// insertion point for imports
import { ParagraphDB } from './paragraph-db'
import { TableDB } from './table-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BodyDB {

	static GONGSTRUCT_NAME = "Body"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Paragraphs?: Array<ParagraphDB>
	Tables?: Array<TableDB>
	LastParagraph?: ParagraphDB
	LastParagraphID: NullInt64 = new NullInt64 // if pointer is null, LastParagraph.ID = 0

}
