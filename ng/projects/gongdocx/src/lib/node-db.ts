// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NodeDB {

	static GONGSTRUCT_NAME = "Node"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Nodes?: Array<NodeDB>
	Node_NodesDBID: NullInt64 = new NullInt64
	Node_NodesDBID_Index: NullInt64  = new NullInt64 // store the index of the node instance in Node.Nodes
	Node_Nodes_reverse?: NodeDB 

}
