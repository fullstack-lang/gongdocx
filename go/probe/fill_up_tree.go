package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongdocx/go/models"
)

func fillUpTree(
	playground *Playground,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range playground.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	playground.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(playground.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](playground.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", playground.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(playground.treeStage)


		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}
		
		switch gongStruct.Name {
		// insertion point
		case "Body":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Body](playground.stageOfInterest)
			for _body := range set {
				nodeInstance := (&tree.Node{Name: _body.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_body, "Body", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Document":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Document](playground.stageOfInterest)
			for _document := range set {
				nodeInstance := (&tree.Node{Name: _document.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_document, "Document", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Docx":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Docx](playground.stageOfInterest)
			for _docx := range set {
				nodeInstance := (&tree.Node{Name: _docx.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_docx, "Docx", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "File":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.File](playground.stageOfInterest)
			for _file := range set {
				nodeInstance := (&tree.Node{Name: _file.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_file, "File", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Node":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Node](playground.stageOfInterest)
			for _node := range set {
				nodeInstance := (&tree.Node{Name: _node.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_node, "Node", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Paragraph":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Paragraph](playground.stageOfInterest)
			for _paragraph := range set {
				nodeInstance := (&tree.Node{Name: _paragraph.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_paragraph, "Paragraph", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ParagraphProperties":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ParagraphProperties](playground.stageOfInterest)
			for _paragraphproperties := range set {
				nodeInstance := (&tree.Node{Name: _paragraphproperties.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_paragraphproperties, "ParagraphProperties", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ParagraphStyle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ParagraphStyle](playground.stageOfInterest)
			for _paragraphstyle := range set {
				nodeInstance := (&tree.Node{Name: _paragraphstyle.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_paragraphstyle, "ParagraphStyle", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Rune":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Rune](playground.stageOfInterest)
			for _rune := range set {
				nodeInstance := (&tree.Node{Name: _rune.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_rune, "Rune", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RuneProperties":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RuneProperties](playground.stageOfInterest)
			for _runeproperties := range set {
				nodeInstance := (&tree.Node{Name: _runeproperties.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_runeproperties, "RuneProperties", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Table":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Table](playground.stageOfInterest)
			for _table := range set {
				nodeInstance := (&tree.Node{Name: _table.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_table, "Table", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TableColumn":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TableColumn](playground.stageOfInterest)
			for _tablecolumn := range set {
				nodeInstance := (&tree.Node{Name: _tablecolumn.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tablecolumn, "TableColumn", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TableProperties":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TableProperties](playground.stageOfInterest)
			for _tableproperties := range set {
				nodeInstance := (&tree.Node{Name: _tableproperties.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tableproperties, "TableProperties", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TableRow":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TableRow](playground.stageOfInterest)
			for _tablerow := range set {
				nodeInstance := (&tree.Node{Name: _tablerow.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tablerow, "TableRow", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TableStyle":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TableStyle](playground.stageOfInterest)
			for _tablestyle := range set {
				nodeInstance := (&tree.Node{Name: _tablestyle.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_tablestyle, "TableStyle", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Text":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Text](playground.stageOfInterest)
			for _text := range set {
				nodeInstance := (&tree.Node{Name: _text.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_text, "Text", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}	
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(playground.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			playground,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}
	playground.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	playground     *Playground
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	playground *Playground) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.playground = playground
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.playground,
	)
}
