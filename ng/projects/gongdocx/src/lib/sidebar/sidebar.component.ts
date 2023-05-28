import { Component, Input, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { DocumentService } from '../document.service'
import { getDocumentUniqueID } from '../front-repo.service'
import { DocxService } from '../docx.service'
import { getDocxUniqueID } from '../front-repo.service'
import { FileService } from '../file.service'
import { getFileUniqueID } from '../front-repo.service'
import { NodeService } from '../node.service'
import { getNodeUniqueID } from '../front-repo.service'
import { TextService } from '../text.service'
import { getTextUniqueID } from '../front-repo.service'

import { RouteService } from '../route-service';

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongdocx-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private documentService: DocumentService,
    private docxService: DocxService,
    private fileService: FileService,
    private nodeService: NodeService,
    private textService: TextService,

    private routeService: RouteService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    console.log("Sidebar init: " + this.GONG__StackPath)

    // add the routes that will used by this side panel component and
    // by the component that are called from this component
    this.routeService.addDataPanelRoutes(this.GONG__StackPath)

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.documentService.DocumentServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.docxService.DocxServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.fileService.FileServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.nodeService.NodeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.textService.TextServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Document part of the mat tree
      */
      let documentGongNodeStruct: GongNode = {
        name: "Document",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Document",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(documentGongNodeStruct)

      this.frontRepo.Documents_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Documents_array.forEach(
        documentDB => {
          let documentGongNodeInstance: GongNode = {
            name: documentDB.Name,
            type: GongNodeType.INSTANCE,
            id: documentDB.ID,
            uniqueIdPerStack: getDocumentUniqueID(documentDB.ID),
            structName: "Document",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          documentGongNodeStruct.children!.push(documentGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association File
          */
          let FileGongNodeAssociation: GongNode = {
            name: "(File) File",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: documentDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Document",
            associationField: "File",
            associatedStructName: "File",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          documentGongNodeInstance.children!.push(FileGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation File
            */
          if (documentDB.File != undefined) {
            let documentGongNodeInstance_File: GongNode = {
              name: documentDB.File.Name,
              type: GongNodeType.INSTANCE,
              id: documentDB.File.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getDocumentUniqueID(documentDB.ID)
                + 5 * getFileUniqueID(documentDB.File.ID),
              structName: "File",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FileGongNodeAssociation.children.push(documentGongNodeInstance_File)
          }

          /**
          * let append a node for the association Root
          */
          let RootGongNodeAssociation: GongNode = {
            name: "(Node) Root",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: documentDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Document",
            associationField: "Root",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          documentGongNodeInstance.children!.push(RootGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Root
            */
          if (documentDB.Root != undefined) {
            let documentGongNodeInstance_Root: GongNode = {
              name: documentDB.Root.Name,
              type: GongNodeType.INSTANCE,
              id: documentDB.Root.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getDocumentUniqueID(documentDB.ID)
                + 5 * getNodeUniqueID(documentDB.Root.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RootGongNodeAssociation.children.push(documentGongNodeInstance_Root)
          }

        }
      )

      /**
      * fill up the Docx part of the mat tree
      */
      let docxGongNodeStruct: GongNode = {
        name: "Docx",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Docx",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(docxGongNodeStruct)

      this.frontRepo.Docxs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Docxs_array.forEach(
        docxDB => {
          let docxGongNodeInstance: GongNode = {
            name: docxDB.Name,
            type: GongNodeType.INSTANCE,
            id: docxDB.ID,
            uniqueIdPerStack: getDocxUniqueID(docxDB.ID),
            structName: "Docx",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          docxGongNodeStruct.children!.push(docxGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Files
          */
          let FilesGongNodeAssociation: GongNode = {
            name: "(File) Files",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: docxDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Docx",
            associationField: "Files",
            associatedStructName: "File",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          docxGongNodeInstance.children.push(FilesGongNodeAssociation)

          docxDB.Files?.forEach(fileDB => {
            let fileNode: GongNode = {
              name: fileDB.Name,
              type: GongNodeType.INSTANCE,
              id: fileDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getDocxUniqueID(docxDB.ID)
                + 11 * getFileUniqueID(fileDB.ID),
              structName: "File",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FilesGongNodeAssociation.children.push(fileNode)
          })

        }
      )

      /**
      * fill up the File part of the mat tree
      */
      let fileGongNodeStruct: GongNode = {
        name: "File",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "File",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(fileGongNodeStruct)

      this.frontRepo.Files_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Files_array.forEach(
        fileDB => {
          let fileGongNodeInstance: GongNode = {
            name: fileDB.Name,
            type: GongNodeType.INSTANCE,
            id: fileDB.ID,
            uniqueIdPerStack: getFileUniqueID(fileDB.ID),
            structName: "File",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          fileGongNodeStruct.children!.push(fileGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Node part of the mat tree
      */
      let nodeGongNodeStruct: GongNode = {
        name: "Node",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Node",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(nodeGongNodeStruct)

      this.frontRepo.Nodes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Nodes_array.forEach(
        nodeDB => {
          let nodeGongNodeInstance: GongNode = {
            name: nodeDB.Name,
            type: GongNodeType.INSTANCE,
            id: nodeDB.ID,
            uniqueIdPerStack: getNodeUniqueID(nodeDB.ID),
            structName: "Node",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nodeGongNodeStruct.children!.push(nodeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Nodes
          */
          let NodesGongNodeAssociation: GongNode = {
            name: "(Node) Nodes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: nodeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Node",
            associationField: "Nodes",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          nodeGongNodeInstance.children.push(NodesGongNodeAssociation)

          nodeDB.Nodes?.forEach(nodeDB => {
            let nodeNode: GongNode = {
              name: nodeDB.Name,
              type: GongNodeType.INSTANCE,
              id: nodeDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getNodeUniqueID(nodeDB.ID)
                + 11 * getNodeUniqueID(nodeDB.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodesGongNodeAssociation.children.push(nodeNode)
          })

        }
      )

      /**
      * fill up the Text part of the mat tree
      */
      let textGongNodeStruct: GongNode = {
        name: "Text",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Text",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(textGongNodeStruct)

      this.frontRepo.Texts_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Texts_array.forEach(
        textDB => {
          let textGongNodeInstance: GongNode = {
            name: textDB.Name,
            type: GongNodeType.INSTANCE,
            id: textDB.ID,
            uniqueIdPerStack: getTextUniqueID(textDB.ID),
            structName: "Text",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          textGongNodeStruct.children!.push(textGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Node
          */
          let NodeGongNodeAssociation: GongNode = {
            name: "(Node) Node",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: textDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Text",
            associationField: "Node",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          textGongNodeInstance.children!.push(NodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Node
            */
          if (textDB.Node != undefined) {
            let textGongNodeInstance_Node: GongNode = {
              name: textDB.Node.Name,
              type: GongNodeType.INSTANCE,
              id: textDB.Node.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getTextUniqueID(textDB.ID)
                + 5 * getNodeUniqueID(textDB.Node.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodeGongNodeAssociation.children.push(textGongNodeInstance_Node)
          }

        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    })
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path
    this.router.navigate([{
      outlets: {
        outletName: [fullPath]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
      let outletConf: any = {}
      outletConf[outletName] = [fullPath, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }

    if (type == GongNodeType.INSTANCE) {
      let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + structName.toLowerCase() + "-detail"

      let outletConf: any = {}
      outletConf[outletName] = [fullPath, id, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }
  }

  setEditorRouterOutlet(path: string) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
    let outletConf : any = {}
    outletConf[outletName] = [fullPath, this.GONG__StackPath]
    this.router.navigate([{ outlets: outletConf }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + node.associatedStructName.toLowerCase() + "-adder"
    this.router.navigate([{
      outlets: {
        outletName: [fullPath, node.id, node.structName, node.associationField, this.GONG__StackPath]
      }
    }]);
  }
}
