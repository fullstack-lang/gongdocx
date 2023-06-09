// insertion point for imports
import { NodeDB } from './node-db'
import { TextDB } from './text-db'
import { RunePropertiesDB } from './runeproperties-db'
import { ParagraphDB } from './paragraph-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RuneDB {

	static GONGSTRUCT_NAME = "Rune"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other declarations
	Node?: NodeDB
	NodeID: NullInt64 = new NullInt64 // if pointer is null, Node.ID = 0

	Text?: TextDB
	TextID: NullInt64 = new NullInt64 // if pointer is null, Text.ID = 0

	RuneProperties?: RunePropertiesDB
	RunePropertiesID: NullInt64 = new NullInt64 // if pointer is null, RuneProperties.ID = 0

	EnclosingParagraph?: ParagraphDB
	EnclosingParagraphID: NullInt64 = new NullInt64 // if pointer is null, EnclosingParagraph.ID = 0

	Paragraph_RunesDBID: NullInt64 = new NullInt64
	Paragraph_RunesDBID_Index: NullInt64  = new NullInt64 // store the index of the rune instance in Paragraph.Runes
	Paragraph_Runes_reverse?: ParagraphDB 

}
