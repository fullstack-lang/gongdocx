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
import { ParagraphService } from '../paragraph.service'
import { getParagraphUniqueID } from '../front-repo.service'
import { ParagraphPropertiesService } from '../paragraphproperties.service'
import { getParagraphPropertiesUniqueID } from '../front-repo.service'
import { ParagraphStyleService } from '../paragraphstyle.service'
import { getParagraphStyleUniqueID } from '../front-repo.service'
import { RuneService } from '../rune.service'
import { getRuneUniqueID } from '../front-repo.service'
import { RunePropertiesService } from '../runeproperties.service'
import { getRunePropertiesUniqueID } from '../front-repo.service'
import { TableService } from '../table.service'
import { getTableUniqueID } from '../front-repo.service'
import { TablePropertiesService } from '../tableproperties.service'
import { getTablePropertiesUniqueID } from '../front-repo.service'
import { TableStyleService } from '../tablestyle.service'
import { getTableStyleUniqueID } from '../front-repo.service'
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
    private paragraphService: ParagraphService,
    private paragraphpropertiesService: ParagraphPropertiesService,
    private paragraphstyleService: ParagraphStyleService,
    private runeService: RuneService,
    private runepropertiesService: RunePropertiesService,
    private tableService: TableService,
    private tablepropertiesService: TablePropertiesService,
    private tablestyleService: TableStyleService,
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
    this.paragraphService.ParagraphServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.paragraphpropertiesService.ParagraphPropertiesServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.paragraphstyleService.ParagraphStyleServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.runeService.RuneServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.runepropertiesService.RunePropertiesServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.tableService.TableServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.tablepropertiesService.TablePropertiesServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.tablestyleService.TableStyleServiceChanged.subscribe(
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
      * fill up the Paragraph part of the mat tree
      */
      let paragraphGongNodeStruct: GongNode = {
        name: "Paragraph",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Paragraph",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(paragraphGongNodeStruct)

      this.frontRepo.Paragraphs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Paragraphs_array.forEach(
        paragraphDB => {
          let paragraphGongNodeInstance: GongNode = {
            name: paragraphDB.Name,
            type: GongNodeType.INSTANCE,
            id: paragraphDB.ID,
            uniqueIdPerStack: getParagraphUniqueID(paragraphDB.ID),
            structName: "Paragraph",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          paragraphGongNodeStruct.children!.push(paragraphGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Node
          */
          let NodeGongNodeAssociation: GongNode = {
            name: "(Node) Node",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: paragraphDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Paragraph",
            associationField: "Node",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          paragraphGongNodeInstance.children!.push(NodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Node
            */
          if (paragraphDB.Node != undefined) {
            let paragraphGongNodeInstance_Node: GongNode = {
              name: paragraphDB.Node.Name,
              type: GongNodeType.INSTANCE,
              id: paragraphDB.Node.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getParagraphUniqueID(paragraphDB.ID)
                + 5 * getNodeUniqueID(paragraphDB.Node.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodeGongNodeAssociation.children.push(paragraphGongNodeInstance_Node)
          }

          /**
          * let append a node for the association ParagraphProperties
          */
          let ParagraphPropertiesGongNodeAssociation: GongNode = {
            name: "(ParagraphProperties) ParagraphProperties",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: paragraphDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Paragraph",
            associationField: "ParagraphProperties",
            associatedStructName: "ParagraphProperties",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          paragraphGongNodeInstance.children!.push(ParagraphPropertiesGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation ParagraphProperties
            */
          if (paragraphDB.ParagraphProperties != undefined) {
            let paragraphGongNodeInstance_ParagraphProperties: GongNode = {
              name: paragraphDB.ParagraphProperties.Name,
              type: GongNodeType.INSTANCE,
              id: paragraphDB.ParagraphProperties.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getParagraphUniqueID(paragraphDB.ID)
                + 5 * getParagraphPropertiesUniqueID(paragraphDB.ParagraphProperties.ID),
              structName: "ParagraphProperties",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ParagraphPropertiesGongNodeAssociation.children.push(paragraphGongNodeInstance_ParagraphProperties)
          }

          /**
          * let append a node for the slide of pointer Runes
          */
          let RunesGongNodeAssociation: GongNode = {
            name: "(Rune) Runes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: paragraphDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Paragraph",
            associationField: "Runes",
            associatedStructName: "Rune",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          paragraphGongNodeInstance.children.push(RunesGongNodeAssociation)

          paragraphDB.Runes?.forEach(runeDB => {
            let runeNode: GongNode = {
              name: runeDB.Name,
              type: GongNodeType.INSTANCE,
              id: runeDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getParagraphUniqueID(paragraphDB.ID)
                + 11 * getRuneUniqueID(runeDB.ID),
              structName: "Rune",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RunesGongNodeAssociation.children.push(runeNode)
          })

        }
      )

      /**
      * fill up the ParagraphProperties part of the mat tree
      */
      let paragraphpropertiesGongNodeStruct: GongNode = {
        name: "ParagraphProperties",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "ParagraphProperties",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(paragraphpropertiesGongNodeStruct)

      this.frontRepo.ParagraphPropertiess_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.ParagraphPropertiess_array.forEach(
        paragraphpropertiesDB => {
          let paragraphpropertiesGongNodeInstance: GongNode = {
            name: paragraphpropertiesDB.Name,
            type: GongNodeType.INSTANCE,
            id: paragraphpropertiesDB.ID,
            uniqueIdPerStack: getParagraphPropertiesUniqueID(paragraphpropertiesDB.ID),
            structName: "ParagraphProperties",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          paragraphpropertiesGongNodeStruct.children!.push(paragraphpropertiesGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association ParagraphStyle
          */
          let ParagraphStyleGongNodeAssociation: GongNode = {
            name: "(ParagraphStyle) ParagraphStyle",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: paragraphpropertiesDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "ParagraphProperties",
            associationField: "ParagraphStyle",
            associatedStructName: "ParagraphStyle",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          paragraphpropertiesGongNodeInstance.children!.push(ParagraphStyleGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation ParagraphStyle
            */
          if (paragraphpropertiesDB.ParagraphStyle != undefined) {
            let paragraphpropertiesGongNodeInstance_ParagraphStyle: GongNode = {
              name: paragraphpropertiesDB.ParagraphStyle.Name,
              type: GongNodeType.INSTANCE,
              id: paragraphpropertiesDB.ParagraphStyle.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getParagraphPropertiesUniqueID(paragraphpropertiesDB.ID)
                + 5 * getParagraphStyleUniqueID(paragraphpropertiesDB.ParagraphStyle.ID),
              structName: "ParagraphStyle",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ParagraphStyleGongNodeAssociation.children.push(paragraphpropertiesGongNodeInstance_ParagraphStyle)
          }

          /**
          * let append a node for the association Node
          */
          let NodeGongNodeAssociation: GongNode = {
            name: "(Node) Node",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: paragraphpropertiesDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "ParagraphProperties",
            associationField: "Node",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          paragraphpropertiesGongNodeInstance.children!.push(NodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Node
            */
          if (paragraphpropertiesDB.Node != undefined) {
            let paragraphpropertiesGongNodeInstance_Node: GongNode = {
              name: paragraphpropertiesDB.Node.Name,
              type: GongNodeType.INSTANCE,
              id: paragraphpropertiesDB.Node.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getParagraphPropertiesUniqueID(paragraphpropertiesDB.ID)
                + 5 * getNodeUniqueID(paragraphpropertiesDB.Node.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodeGongNodeAssociation.children.push(paragraphpropertiesGongNodeInstance_Node)
          }

        }
      )

      /**
      * fill up the ParagraphStyle part of the mat tree
      */
      let paragraphstyleGongNodeStruct: GongNode = {
        name: "ParagraphStyle",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "ParagraphStyle",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(paragraphstyleGongNodeStruct)

      this.frontRepo.ParagraphStyles_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.ParagraphStyles_array.forEach(
        paragraphstyleDB => {
          let paragraphstyleGongNodeInstance: GongNode = {
            name: paragraphstyleDB.Name,
            type: GongNodeType.INSTANCE,
            id: paragraphstyleDB.ID,
            uniqueIdPerStack: getParagraphStyleUniqueID(paragraphstyleDB.ID),
            structName: "ParagraphStyle",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          paragraphstyleGongNodeStruct.children!.push(paragraphstyleGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Node
          */
          let NodeGongNodeAssociation: GongNode = {
            name: "(Node) Node",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: paragraphstyleDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "ParagraphStyle",
            associationField: "Node",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          paragraphstyleGongNodeInstance.children!.push(NodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Node
            */
          if (paragraphstyleDB.Node != undefined) {
            let paragraphstyleGongNodeInstance_Node: GongNode = {
              name: paragraphstyleDB.Node.Name,
              type: GongNodeType.INSTANCE,
              id: paragraphstyleDB.Node.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getParagraphStyleUniqueID(paragraphstyleDB.ID)
                + 5 * getNodeUniqueID(paragraphstyleDB.Node.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodeGongNodeAssociation.children.push(paragraphstyleGongNodeInstance_Node)
          }

        }
      )

      /**
      * fill up the Rune part of the mat tree
      */
      let runeGongNodeStruct: GongNode = {
        name: "Rune",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Rune",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(runeGongNodeStruct)

      this.frontRepo.Runes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Runes_array.forEach(
        runeDB => {
          let runeGongNodeInstance: GongNode = {
            name: runeDB.Name,
            type: GongNodeType.INSTANCE,
            id: runeDB.ID,
            uniqueIdPerStack: getRuneUniqueID(runeDB.ID),
            structName: "Rune",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          runeGongNodeStruct.children!.push(runeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Node
          */
          let NodeGongNodeAssociation: GongNode = {
            name: "(Node) Node",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: runeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Rune",
            associationField: "Node",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          runeGongNodeInstance.children!.push(NodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Node
            */
          if (runeDB.Node != undefined) {
            let runeGongNodeInstance_Node: GongNode = {
              name: runeDB.Node.Name,
              type: GongNodeType.INSTANCE,
              id: runeDB.Node.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getRuneUniqueID(runeDB.ID)
                + 5 * getNodeUniqueID(runeDB.Node.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodeGongNodeAssociation.children.push(runeGongNodeInstance_Node)
          }

          /**
          * let append a node for the association Text
          */
          let TextGongNodeAssociation: GongNode = {
            name: "(Text) Text",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: runeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Rune",
            associationField: "Text",
            associatedStructName: "Text",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          runeGongNodeInstance.children!.push(TextGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Text
            */
          if (runeDB.Text != undefined) {
            let runeGongNodeInstance_Text: GongNode = {
              name: runeDB.Text.Name,
              type: GongNodeType.INSTANCE,
              id: runeDB.Text.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getRuneUniqueID(runeDB.ID)
                + 5 * getTextUniqueID(runeDB.Text.ID),
              structName: "Text",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            TextGongNodeAssociation.children.push(runeGongNodeInstance_Text)
          }

          /**
          * let append a node for the association RuneProperties
          */
          let RunePropertiesGongNodeAssociation: GongNode = {
            name: "(RuneProperties) RuneProperties",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: runeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Rune",
            associationField: "RuneProperties",
            associatedStructName: "RuneProperties",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          runeGongNodeInstance.children!.push(RunePropertiesGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation RuneProperties
            */
          if (runeDB.RuneProperties != undefined) {
            let runeGongNodeInstance_RuneProperties: GongNode = {
              name: runeDB.RuneProperties.Name,
              type: GongNodeType.INSTANCE,
              id: runeDB.RuneProperties.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getRuneUniqueID(runeDB.ID)
                + 5 * getRunePropertiesUniqueID(runeDB.RuneProperties.ID),
              structName: "RuneProperties",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RunePropertiesGongNodeAssociation.children.push(runeGongNodeInstance_RuneProperties)
          }

        }
      )

      /**
      * fill up the RuneProperties part of the mat tree
      */
      let runepropertiesGongNodeStruct: GongNode = {
        name: "RuneProperties",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "RuneProperties",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(runepropertiesGongNodeStruct)

      this.frontRepo.RunePropertiess_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.RunePropertiess_array.forEach(
        runepropertiesDB => {
          let runepropertiesGongNodeInstance: GongNode = {
            name: runepropertiesDB.Name,
            type: GongNodeType.INSTANCE,
            id: runepropertiesDB.ID,
            uniqueIdPerStack: getRunePropertiesUniqueID(runepropertiesDB.ID),
            structName: "RuneProperties",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          runepropertiesGongNodeStruct.children!.push(runepropertiesGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Node
          */
          let NodeGongNodeAssociation: GongNode = {
            name: "(Node) Node",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: runepropertiesDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "RuneProperties",
            associationField: "Node",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          runepropertiesGongNodeInstance.children!.push(NodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Node
            */
          if (runepropertiesDB.Node != undefined) {
            let runepropertiesGongNodeInstance_Node: GongNode = {
              name: runepropertiesDB.Node.Name,
              type: GongNodeType.INSTANCE,
              id: runepropertiesDB.Node.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getRunePropertiesUniqueID(runepropertiesDB.ID)
                + 5 * getNodeUniqueID(runepropertiesDB.Node.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodeGongNodeAssociation.children.push(runepropertiesGongNodeInstance_Node)
          }

        }
      )

      /**
      * fill up the Table part of the mat tree
      */
      let tableGongNodeStruct: GongNode = {
        name: "Table",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Table",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(tableGongNodeStruct)

      this.frontRepo.Tables_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Tables_array.forEach(
        tableDB => {
          let tableGongNodeInstance: GongNode = {
            name: tableDB.Name,
            type: GongNodeType.INSTANCE,
            id: tableDB.ID,
            uniqueIdPerStack: getTableUniqueID(tableDB.ID),
            structName: "Table",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          tableGongNodeStruct.children!.push(tableGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Node
          */
          let NodeGongNodeAssociation: GongNode = {
            name: "(Node) Node",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: tableDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Table",
            associationField: "Node",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          tableGongNodeInstance.children!.push(NodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Node
            */
          if (tableDB.Node != undefined) {
            let tableGongNodeInstance_Node: GongNode = {
              name: tableDB.Node.Name,
              type: GongNodeType.INSTANCE,
              id: tableDB.Node.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getTableUniqueID(tableDB.ID)
                + 5 * getNodeUniqueID(tableDB.Node.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodeGongNodeAssociation.children.push(tableGongNodeInstance_Node)
          }

          /**
          * let append a node for the association TableProperties
          */
          let TablePropertiesGongNodeAssociation: GongNode = {
            name: "(TableProperties) TableProperties",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: tableDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Table",
            associationField: "TableProperties",
            associatedStructName: "TableProperties",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          tableGongNodeInstance.children!.push(TablePropertiesGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation TableProperties
            */
          if (tableDB.TableProperties != undefined) {
            let tableGongNodeInstance_TableProperties: GongNode = {
              name: tableDB.TableProperties.Name,
              type: GongNodeType.INSTANCE,
              id: tableDB.TableProperties.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getTableUniqueID(tableDB.ID)
                + 5 * getTablePropertiesUniqueID(tableDB.TableProperties.ID),
              structName: "TableProperties",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            TablePropertiesGongNodeAssociation.children.push(tableGongNodeInstance_TableProperties)
          }

        }
      )

      /**
      * fill up the TableProperties part of the mat tree
      */
      let tablepropertiesGongNodeStruct: GongNode = {
        name: "TableProperties",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "TableProperties",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(tablepropertiesGongNodeStruct)

      this.frontRepo.TablePropertiess_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.TablePropertiess_array.forEach(
        tablepropertiesDB => {
          let tablepropertiesGongNodeInstance: GongNode = {
            name: tablepropertiesDB.Name,
            type: GongNodeType.INSTANCE,
            id: tablepropertiesDB.ID,
            uniqueIdPerStack: getTablePropertiesUniqueID(tablepropertiesDB.ID),
            structName: "TableProperties",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          tablepropertiesGongNodeStruct.children!.push(tablepropertiesGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Node
          */
          let NodeGongNodeAssociation: GongNode = {
            name: "(Node) Node",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: tablepropertiesDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "TableProperties",
            associationField: "Node",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          tablepropertiesGongNodeInstance.children!.push(NodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Node
            */
          if (tablepropertiesDB.Node != undefined) {
            let tablepropertiesGongNodeInstance_Node: GongNode = {
              name: tablepropertiesDB.Node.Name,
              type: GongNodeType.INSTANCE,
              id: tablepropertiesDB.Node.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getTablePropertiesUniqueID(tablepropertiesDB.ID)
                + 5 * getNodeUniqueID(tablepropertiesDB.Node.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodeGongNodeAssociation.children.push(tablepropertiesGongNodeInstance_Node)
          }

          /**
          * let append a node for the association TableStyle
          */
          let TableStyleGongNodeAssociation: GongNode = {
            name: "(TableStyle) TableStyle",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: tablepropertiesDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "TableProperties",
            associationField: "TableStyle",
            associatedStructName: "TableStyle",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          tablepropertiesGongNodeInstance.children!.push(TableStyleGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation TableStyle
            */
          if (tablepropertiesDB.TableStyle != undefined) {
            let tablepropertiesGongNodeInstance_TableStyle: GongNode = {
              name: tablepropertiesDB.TableStyle.Name,
              type: GongNodeType.INSTANCE,
              id: tablepropertiesDB.TableStyle.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getTablePropertiesUniqueID(tablepropertiesDB.ID)
                + 5 * getTableStyleUniqueID(tablepropertiesDB.TableStyle.ID),
              structName: "TableStyle",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            TableStyleGongNodeAssociation.children.push(tablepropertiesGongNodeInstance_TableStyle)
          }

        }
      )

      /**
      * fill up the TableStyle part of the mat tree
      */
      let tablestyleGongNodeStruct: GongNode = {
        name: "TableStyle",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "TableStyle",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(tablestyleGongNodeStruct)

      this.frontRepo.TableStyles_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.TableStyles_array.forEach(
        tablestyleDB => {
          let tablestyleGongNodeInstance: GongNode = {
            name: tablestyleDB.Name,
            type: GongNodeType.INSTANCE,
            id: tablestyleDB.ID,
            uniqueIdPerStack: getTableStyleUniqueID(tablestyleDB.ID),
            structName: "TableStyle",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          tablestyleGongNodeStruct.children!.push(tablestyleGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Node
          */
          let NodeGongNodeAssociation: GongNode = {
            name: "(Node) Node",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: tablestyleDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "TableStyle",
            associationField: "Node",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          tablestyleGongNodeInstance.children!.push(NodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Node
            */
          if (tablestyleDB.Node != undefined) {
            let tablestyleGongNodeInstance_Node: GongNode = {
              name: tablestyleDB.Node.Name,
              type: GongNodeType.INSTANCE,
              id: tablestyleDB.Node.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getTableStyleUniqueID(tablestyleDB.ID)
                + 5 * getNodeUniqueID(tablestyleDB.Node.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            NodeGongNodeAssociation.children.push(tablestyleGongNodeInstance_Node)
          }

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
