package models

// GONGNOTE(NodeOnText) :
// [models.Text]
// -
// t: This stands for 'text'. This element is used to represent a run of text within
// a paragraph. The t element will contain the actual string of text as its content.
// -
// When parsing these nodes, your code should handle each type of node appropriately
// based on its name. For example, when you encounter a t node, you might simply
// extract and store the text content.
const NodeOnText = ""

type Text struct {
	Name    string
	Content string
	Node    *Node
}
