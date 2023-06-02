import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { BodyDB } from './body-db'
import { BodyService } from './body.service'

import { DocumentDB } from './document-db'
import { DocumentService } from './document.service'

import { DocxDB } from './docx-db'
import { DocxService } from './docx.service'

import { FileDB } from './file-db'
import { FileService } from './file.service'

import { NodeDB } from './node-db'
import { NodeService } from './node.service'

import { ParagraphDB } from './paragraph-db'
import { ParagraphService } from './paragraph.service'

import { ParagraphPropertiesDB } from './paragraphproperties-db'
import { ParagraphPropertiesService } from './paragraphproperties.service'

import { ParagraphStyleDB } from './paragraphstyle-db'
import { ParagraphStyleService } from './paragraphstyle.service'

import { RuneDB } from './rune-db'
import { RuneService } from './rune.service'

import { RunePropertiesDB } from './runeproperties-db'
import { RunePropertiesService } from './runeproperties.service'

import { TableDB } from './table-db'
import { TableService } from './table.service'

import { TableColumnDB } from './tablecolumn-db'
import { TableColumnService } from './tablecolumn.service'

import { TablePropertiesDB } from './tableproperties-db'
import { TablePropertiesService } from './tableproperties.service'

import { TableRowDB } from './tablerow-db'
import { TableRowService } from './tablerow.service'

import { TableStyleDB } from './tablestyle-db'
import { TableStyleService } from './tablestyle.service'

import { TextDB } from './text-db'
import { TextService } from './text.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Bodys_array = new Array<BodyDB>(); // array of repo instances
  Bodys = new Map<number, BodyDB>(); // map of repo instances
  Bodys_batch = new Map<number, BodyDB>(); // same but only in last GET (for finding repo instances to delete)
  Documents_array = new Array<DocumentDB>(); // array of repo instances
  Documents = new Map<number, DocumentDB>(); // map of repo instances
  Documents_batch = new Map<number, DocumentDB>(); // same but only in last GET (for finding repo instances to delete)
  Docxs_array = new Array<DocxDB>(); // array of repo instances
  Docxs = new Map<number, DocxDB>(); // map of repo instances
  Docxs_batch = new Map<number, DocxDB>(); // same but only in last GET (for finding repo instances to delete)
  Files_array = new Array<FileDB>(); // array of repo instances
  Files = new Map<number, FileDB>(); // map of repo instances
  Files_batch = new Map<number, FileDB>(); // same but only in last GET (for finding repo instances to delete)
  Nodes_array = new Array<NodeDB>(); // array of repo instances
  Nodes = new Map<number, NodeDB>(); // map of repo instances
  Nodes_batch = new Map<number, NodeDB>(); // same but only in last GET (for finding repo instances to delete)
  Paragraphs_array = new Array<ParagraphDB>(); // array of repo instances
  Paragraphs = new Map<number, ParagraphDB>(); // map of repo instances
  Paragraphs_batch = new Map<number, ParagraphDB>(); // same but only in last GET (for finding repo instances to delete)
  ParagraphPropertiess_array = new Array<ParagraphPropertiesDB>(); // array of repo instances
  ParagraphPropertiess = new Map<number, ParagraphPropertiesDB>(); // map of repo instances
  ParagraphPropertiess_batch = new Map<number, ParagraphPropertiesDB>(); // same but only in last GET (for finding repo instances to delete)
  ParagraphStyles_array = new Array<ParagraphStyleDB>(); // array of repo instances
  ParagraphStyles = new Map<number, ParagraphStyleDB>(); // map of repo instances
  ParagraphStyles_batch = new Map<number, ParagraphStyleDB>(); // same but only in last GET (for finding repo instances to delete)
  Runes_array = new Array<RuneDB>(); // array of repo instances
  Runes = new Map<number, RuneDB>(); // map of repo instances
  Runes_batch = new Map<number, RuneDB>(); // same but only in last GET (for finding repo instances to delete)
  RunePropertiess_array = new Array<RunePropertiesDB>(); // array of repo instances
  RunePropertiess = new Map<number, RunePropertiesDB>(); // map of repo instances
  RunePropertiess_batch = new Map<number, RunePropertiesDB>(); // same but only in last GET (for finding repo instances to delete)
  Tables_array = new Array<TableDB>(); // array of repo instances
  Tables = new Map<number, TableDB>(); // map of repo instances
  Tables_batch = new Map<number, TableDB>(); // same but only in last GET (for finding repo instances to delete)
  TableColumns_array = new Array<TableColumnDB>(); // array of repo instances
  TableColumns = new Map<number, TableColumnDB>(); // map of repo instances
  TableColumns_batch = new Map<number, TableColumnDB>(); // same but only in last GET (for finding repo instances to delete)
  TablePropertiess_array = new Array<TablePropertiesDB>(); // array of repo instances
  TablePropertiess = new Map<number, TablePropertiesDB>(); // map of repo instances
  TablePropertiess_batch = new Map<number, TablePropertiesDB>(); // same but only in last GET (for finding repo instances to delete)
  TableRows_array = new Array<TableRowDB>(); // array of repo instances
  TableRows = new Map<number, TableRowDB>(); // map of repo instances
  TableRows_batch = new Map<number, TableRowDB>(); // same but only in last GET (for finding repo instances to delete)
  TableStyles_array = new Array<TableStyleDB>(); // array of repo instances
  TableStyles = new Map<number, TableStyleDB>(); // map of repo instances
  TableStyles_batch = new Map<number, TableStyleDB>(); // same but only in last GET (for finding repo instances to delete)
  Texts_array = new Array<TextDB>(); // array of repo instances
  Texts = new Map<number, TextDB>(); // map of repo instances
  Texts_batch = new Map<number, TextDB>(); // same but only in last GET (for finding repo instances to delete)
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"

  GONG__StackPath: string = ""
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template 
    private bodyService: BodyService,
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
    private tablecolumnService: TableColumnService,
    private tablepropertiesService: TablePropertiesService,
    private tablerowService: TableRowService,
    private tablestyleService: TableStyleService,
    private textService: TextService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<BodyDB[]>,
    Observable<DocumentDB[]>,
    Observable<DocxDB[]>,
    Observable<FileDB[]>,
    Observable<NodeDB[]>,
    Observable<ParagraphDB[]>,
    Observable<ParagraphPropertiesDB[]>,
    Observable<ParagraphStyleDB[]>,
    Observable<RuneDB[]>,
    Observable<RunePropertiesDB[]>,
    Observable<TableDB[]>,
    Observable<TableColumnDB[]>,
    Observable<TablePropertiesDB[]>,
    Observable<TableRowDB[]>,
    Observable<TableStyleDB[]>,
    Observable<TextDB[]>,
  ] = [ // insertion point sub template
      this.bodyService.getBodys(this.GONG__StackPath),
      this.documentService.getDocuments(this.GONG__StackPath),
      this.docxService.getDocxs(this.GONG__StackPath),
      this.fileService.getFiles(this.GONG__StackPath),
      this.nodeService.getNodes(this.GONG__StackPath),
      this.paragraphService.getParagraphs(this.GONG__StackPath),
      this.paragraphpropertiesService.getParagraphPropertiess(this.GONG__StackPath),
      this.paragraphstyleService.getParagraphStyles(this.GONG__StackPath),
      this.runeService.getRunes(this.GONG__StackPath),
      this.runepropertiesService.getRunePropertiess(this.GONG__StackPath),
      this.tableService.getTables(this.GONG__StackPath),
      this.tablecolumnService.getTableColumns(this.GONG__StackPath),
      this.tablepropertiesService.getTablePropertiess(this.GONG__StackPath),
      this.tablerowService.getTableRows(this.GONG__StackPath),
      this.tablestyleService.getTableStyles(this.GONG__StackPath),
      this.textService.getTexts(this.GONG__StackPath),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [ // insertion point sub template
      this.bodyService.getBodys(this.GONG__StackPath),
      this.documentService.getDocuments(this.GONG__StackPath),
      this.docxService.getDocxs(this.GONG__StackPath),
      this.fileService.getFiles(this.GONG__StackPath),
      this.nodeService.getNodes(this.GONG__StackPath),
      this.paragraphService.getParagraphs(this.GONG__StackPath),
      this.paragraphpropertiesService.getParagraphPropertiess(this.GONG__StackPath),
      this.paragraphstyleService.getParagraphStyles(this.GONG__StackPath),
      this.runeService.getRunes(this.GONG__StackPath),
      this.runepropertiesService.getRunePropertiess(this.GONG__StackPath),
      this.tableService.getTables(this.GONG__StackPath),
      this.tablecolumnService.getTableColumns(this.GONG__StackPath),
      this.tablepropertiesService.getTablePropertiess(this.GONG__StackPath),
      this.tablerowService.getTableRows(this.GONG__StackPath),
      this.tablestyleService.getTableStyles(this.GONG__StackPath),
      this.textService.getTexts(this.GONG__StackPath),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            bodys_,
            documents_,
            docxs_,
            files_,
            nodes_,
            paragraphs_,
            paragraphpropertiess_,
            paragraphstyles_,
            runes_,
            runepropertiess_,
            tables_,
            tablecolumns_,
            tablepropertiess_,
            tablerows_,
            tablestyles_,
            texts_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var bodys: BodyDB[]
            bodys = bodys_ as BodyDB[]
            var documents: DocumentDB[]
            documents = documents_ as DocumentDB[]
            var docxs: DocxDB[]
            docxs = docxs_ as DocxDB[]
            var files: FileDB[]
            files = files_ as FileDB[]
            var nodes: NodeDB[]
            nodes = nodes_ as NodeDB[]
            var paragraphs: ParagraphDB[]
            paragraphs = paragraphs_ as ParagraphDB[]
            var paragraphpropertiess: ParagraphPropertiesDB[]
            paragraphpropertiess = paragraphpropertiess_ as ParagraphPropertiesDB[]
            var paragraphstyles: ParagraphStyleDB[]
            paragraphstyles = paragraphstyles_ as ParagraphStyleDB[]
            var runes: RuneDB[]
            runes = runes_ as RuneDB[]
            var runepropertiess: RunePropertiesDB[]
            runepropertiess = runepropertiess_ as RunePropertiesDB[]
            var tables: TableDB[]
            tables = tables_ as TableDB[]
            var tablecolumns: TableColumnDB[]
            tablecolumns = tablecolumns_ as TableColumnDB[]
            var tablepropertiess: TablePropertiesDB[]
            tablepropertiess = tablepropertiess_ as TablePropertiesDB[]
            var tablerows: TableRowDB[]
            tablerows = tablerows_ as TableRowDB[]
            var tablestyles: TableStyleDB[]
            tablestyles = tablestyles_ as TableStyleDB[]
            var texts: TextDB[]
            texts = texts_ as TextDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.Bodys_array = bodys

            // clear the map that counts Body in the GET
            this.frontRepo.Bodys_batch.clear()

            bodys.forEach(
              body => {
                this.frontRepo.Bodys.set(body.ID, body)
                this.frontRepo.Bodys_batch.set(body.ID, body)
              }
            )

            // clear bodys that are absent from the batch
            this.frontRepo.Bodys.forEach(
              body => {
                if (this.frontRepo.Bodys_batch.get(body.ID) == undefined) {
                  this.frontRepo.Bodys.delete(body.ID)
                }
              }
            )

            // sort Bodys_array array
            this.frontRepo.Bodys_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Documents_array = documents

            // clear the map that counts Document in the GET
            this.frontRepo.Documents_batch.clear()

            documents.forEach(
              document => {
                this.frontRepo.Documents.set(document.ID, document)
                this.frontRepo.Documents_batch.set(document.ID, document)
              }
            )

            // clear documents that are absent from the batch
            this.frontRepo.Documents.forEach(
              document => {
                if (this.frontRepo.Documents_batch.get(document.ID) == undefined) {
                  this.frontRepo.Documents.delete(document.ID)
                }
              }
            )

            // sort Documents_array array
            this.frontRepo.Documents_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Docxs_array = docxs

            // clear the map that counts Docx in the GET
            this.frontRepo.Docxs_batch.clear()

            docxs.forEach(
              docx => {
                this.frontRepo.Docxs.set(docx.ID, docx)
                this.frontRepo.Docxs_batch.set(docx.ID, docx)
              }
            )

            // clear docxs that are absent from the batch
            this.frontRepo.Docxs.forEach(
              docx => {
                if (this.frontRepo.Docxs_batch.get(docx.ID) == undefined) {
                  this.frontRepo.Docxs.delete(docx.ID)
                }
              }
            )

            // sort Docxs_array array
            this.frontRepo.Docxs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Files_array = files

            // clear the map that counts File in the GET
            this.frontRepo.Files_batch.clear()

            files.forEach(
              file => {
                this.frontRepo.Files.set(file.ID, file)
                this.frontRepo.Files_batch.set(file.ID, file)
              }
            )

            // clear files that are absent from the batch
            this.frontRepo.Files.forEach(
              file => {
                if (this.frontRepo.Files_batch.get(file.ID) == undefined) {
                  this.frontRepo.Files.delete(file.ID)
                }
              }
            )

            // sort Files_array array
            this.frontRepo.Files_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Nodes_array = nodes

            // clear the map that counts Node in the GET
            this.frontRepo.Nodes_batch.clear()

            nodes.forEach(
              node => {
                this.frontRepo.Nodes.set(node.ID, node)
                this.frontRepo.Nodes_batch.set(node.ID, node)
              }
            )

            // clear nodes that are absent from the batch
            this.frontRepo.Nodes.forEach(
              node => {
                if (this.frontRepo.Nodes_batch.get(node.ID) == undefined) {
                  this.frontRepo.Nodes.delete(node.ID)
                }
              }
            )

            // sort Nodes_array array
            this.frontRepo.Nodes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Paragraphs_array = paragraphs

            // clear the map that counts Paragraph in the GET
            this.frontRepo.Paragraphs_batch.clear()

            paragraphs.forEach(
              paragraph => {
                this.frontRepo.Paragraphs.set(paragraph.ID, paragraph)
                this.frontRepo.Paragraphs_batch.set(paragraph.ID, paragraph)
              }
            )

            // clear paragraphs that are absent from the batch
            this.frontRepo.Paragraphs.forEach(
              paragraph => {
                if (this.frontRepo.Paragraphs_batch.get(paragraph.ID) == undefined) {
                  this.frontRepo.Paragraphs.delete(paragraph.ID)
                }
              }
            )

            // sort Paragraphs_array array
            this.frontRepo.Paragraphs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.ParagraphPropertiess_array = paragraphpropertiess

            // clear the map that counts ParagraphProperties in the GET
            this.frontRepo.ParagraphPropertiess_batch.clear()

            paragraphpropertiess.forEach(
              paragraphproperties => {
                this.frontRepo.ParagraphPropertiess.set(paragraphproperties.ID, paragraphproperties)
                this.frontRepo.ParagraphPropertiess_batch.set(paragraphproperties.ID, paragraphproperties)
              }
            )

            // clear paragraphpropertiess that are absent from the batch
            this.frontRepo.ParagraphPropertiess.forEach(
              paragraphproperties => {
                if (this.frontRepo.ParagraphPropertiess_batch.get(paragraphproperties.ID) == undefined) {
                  this.frontRepo.ParagraphPropertiess.delete(paragraphproperties.ID)
                }
              }
            )

            // sort ParagraphPropertiess_array array
            this.frontRepo.ParagraphPropertiess_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.ParagraphStyles_array = paragraphstyles

            // clear the map that counts ParagraphStyle in the GET
            this.frontRepo.ParagraphStyles_batch.clear()

            paragraphstyles.forEach(
              paragraphstyle => {
                this.frontRepo.ParagraphStyles.set(paragraphstyle.ID, paragraphstyle)
                this.frontRepo.ParagraphStyles_batch.set(paragraphstyle.ID, paragraphstyle)
              }
            )

            // clear paragraphstyles that are absent from the batch
            this.frontRepo.ParagraphStyles.forEach(
              paragraphstyle => {
                if (this.frontRepo.ParagraphStyles_batch.get(paragraphstyle.ID) == undefined) {
                  this.frontRepo.ParagraphStyles.delete(paragraphstyle.ID)
                }
              }
            )

            // sort ParagraphStyles_array array
            this.frontRepo.ParagraphStyles_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Runes_array = runes

            // clear the map that counts Rune in the GET
            this.frontRepo.Runes_batch.clear()

            runes.forEach(
              rune => {
                this.frontRepo.Runes.set(rune.ID, rune)
                this.frontRepo.Runes_batch.set(rune.ID, rune)
              }
            )

            // clear runes that are absent from the batch
            this.frontRepo.Runes.forEach(
              rune => {
                if (this.frontRepo.Runes_batch.get(rune.ID) == undefined) {
                  this.frontRepo.Runes.delete(rune.ID)
                }
              }
            )

            // sort Runes_array array
            this.frontRepo.Runes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.RunePropertiess_array = runepropertiess

            // clear the map that counts RuneProperties in the GET
            this.frontRepo.RunePropertiess_batch.clear()

            runepropertiess.forEach(
              runeproperties => {
                this.frontRepo.RunePropertiess.set(runeproperties.ID, runeproperties)
                this.frontRepo.RunePropertiess_batch.set(runeproperties.ID, runeproperties)
              }
            )

            // clear runepropertiess that are absent from the batch
            this.frontRepo.RunePropertiess.forEach(
              runeproperties => {
                if (this.frontRepo.RunePropertiess_batch.get(runeproperties.ID) == undefined) {
                  this.frontRepo.RunePropertiess.delete(runeproperties.ID)
                }
              }
            )

            // sort RunePropertiess_array array
            this.frontRepo.RunePropertiess_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Tables_array = tables

            // clear the map that counts Table in the GET
            this.frontRepo.Tables_batch.clear()

            tables.forEach(
              table => {
                this.frontRepo.Tables.set(table.ID, table)
                this.frontRepo.Tables_batch.set(table.ID, table)
              }
            )

            // clear tables that are absent from the batch
            this.frontRepo.Tables.forEach(
              table => {
                if (this.frontRepo.Tables_batch.get(table.ID) == undefined) {
                  this.frontRepo.Tables.delete(table.ID)
                }
              }
            )

            // sort Tables_array array
            this.frontRepo.Tables_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.TableColumns_array = tablecolumns

            // clear the map that counts TableColumn in the GET
            this.frontRepo.TableColumns_batch.clear()

            tablecolumns.forEach(
              tablecolumn => {
                this.frontRepo.TableColumns.set(tablecolumn.ID, tablecolumn)
                this.frontRepo.TableColumns_batch.set(tablecolumn.ID, tablecolumn)
              }
            )

            // clear tablecolumns that are absent from the batch
            this.frontRepo.TableColumns.forEach(
              tablecolumn => {
                if (this.frontRepo.TableColumns_batch.get(tablecolumn.ID) == undefined) {
                  this.frontRepo.TableColumns.delete(tablecolumn.ID)
                }
              }
            )

            // sort TableColumns_array array
            this.frontRepo.TableColumns_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.TablePropertiess_array = tablepropertiess

            // clear the map that counts TableProperties in the GET
            this.frontRepo.TablePropertiess_batch.clear()

            tablepropertiess.forEach(
              tableproperties => {
                this.frontRepo.TablePropertiess.set(tableproperties.ID, tableproperties)
                this.frontRepo.TablePropertiess_batch.set(tableproperties.ID, tableproperties)
              }
            )

            // clear tablepropertiess that are absent from the batch
            this.frontRepo.TablePropertiess.forEach(
              tableproperties => {
                if (this.frontRepo.TablePropertiess_batch.get(tableproperties.ID) == undefined) {
                  this.frontRepo.TablePropertiess.delete(tableproperties.ID)
                }
              }
            )

            // sort TablePropertiess_array array
            this.frontRepo.TablePropertiess_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.TableRows_array = tablerows

            // clear the map that counts TableRow in the GET
            this.frontRepo.TableRows_batch.clear()

            tablerows.forEach(
              tablerow => {
                this.frontRepo.TableRows.set(tablerow.ID, tablerow)
                this.frontRepo.TableRows_batch.set(tablerow.ID, tablerow)
              }
            )

            // clear tablerows that are absent from the batch
            this.frontRepo.TableRows.forEach(
              tablerow => {
                if (this.frontRepo.TableRows_batch.get(tablerow.ID) == undefined) {
                  this.frontRepo.TableRows.delete(tablerow.ID)
                }
              }
            )

            // sort TableRows_array array
            this.frontRepo.TableRows_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.TableStyles_array = tablestyles

            // clear the map that counts TableStyle in the GET
            this.frontRepo.TableStyles_batch.clear()

            tablestyles.forEach(
              tablestyle => {
                this.frontRepo.TableStyles.set(tablestyle.ID, tablestyle)
                this.frontRepo.TableStyles_batch.set(tablestyle.ID, tablestyle)
              }
            )

            // clear tablestyles that are absent from the batch
            this.frontRepo.TableStyles.forEach(
              tablestyle => {
                if (this.frontRepo.TableStyles_batch.get(tablestyle.ID) == undefined) {
                  this.frontRepo.TableStyles.delete(tablestyle.ID)
                }
              }
            )

            // sort TableStyles_array array
            this.frontRepo.TableStyles_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Texts_array = texts

            // clear the map that counts Text in the GET
            this.frontRepo.Texts_batch.clear()

            texts.forEach(
              text => {
                this.frontRepo.Texts.set(text.ID, text)
                this.frontRepo.Texts_batch.set(text.ID, text)
              }
            )

            // clear texts that are absent from the batch
            this.frontRepo.Texts.forEach(
              text => {
                if (this.frontRepo.Texts_batch.get(text.ID) == undefined) {
                  this.frontRepo.Texts.delete(text.ID)
                }
              }
            )

            // sort Texts_array array
            this.frontRepo.Texts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            bodys.forEach(
              body => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field LastParagraph redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(body.LastParagraphID.Int64)
                  if (_paragraph) {
                    body.LastParagraph = _paragraph
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            documents.forEach(
              document => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field File redeeming
                {
                  let _file = this.frontRepo.Files.get(document.FileID.Int64)
                  if (_file) {
                    document.File = _file
                  }
                }
                // insertion point for pointer field Root redeeming
                {
                  let _node = this.frontRepo.Nodes.get(document.RootID.Int64)
                  if (_node) {
                    document.Root = _node
                  }
                }
                // insertion point for pointer field Body redeeming
                {
                  let _body = this.frontRepo.Bodys.get(document.BodyID.Int64)
                  if (_body) {
                    document.Body = _body
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            docxs.forEach(
              docx => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Document redeeming
                {
                  let _document = this.frontRepo.Documents.get(docx.DocumentID.Int64)
                  if (_document) {
                    docx.Document = _document
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            files.forEach(
              file => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Docx.Files redeeming
                {
                  let _docx = this.frontRepo.Docxs.get(file.Docx_FilesDBID.Int64)
                  if (_docx) {
                    if (_docx.Files == undefined) {
                      _docx.Files = new Array<FileDB>()
                    }
                    _docx.Files.push(file)
                    if (file.Docx_Files_reverse == undefined) {
                      file.Docx_Files_reverse = _docx
                    }
                  }
                }
              }
            )
            nodes.forEach(
              node => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Node.Nodes redeeming
                {
                  let _node = this.frontRepo.Nodes.get(node.Node_NodesDBID.Int64)
                  if (_node) {
                    if (_node.Nodes == undefined) {
                      _node.Nodes = new Array<NodeDB>()
                    }
                    _node.Nodes.push(node)
                    if (node.Node_Nodes_reverse == undefined) {
                      node.Node_Nodes_reverse = _node
                    }
                  }
                }
              }
            )
            paragraphs.forEach(
              paragraph => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(paragraph.NodeID.Int64)
                  if (_node) {
                    paragraph.Node = _node
                  }
                }
                // insertion point for pointer field ParagraphProperties redeeming
                {
                  let _paragraphproperties = this.frontRepo.ParagraphPropertiess.get(paragraph.ParagraphPropertiesID.Int64)
                  if (_paragraphproperties) {
                    paragraph.ParagraphProperties = _paragraphproperties
                  }
                }
                // insertion point for pointer field Next redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(paragraph.NextID.Int64)
                  if (_paragraph) {
                    paragraph.Next = _paragraph
                  }
                }
                // insertion point for pointer field Previous redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(paragraph.PreviousID.Int64)
                  if (_paragraph) {
                    paragraph.Previous = _paragraph
                  }
                }
                // insertion point for pointer field EnclosingBody redeeming
                {
                  let _body = this.frontRepo.Bodys.get(paragraph.EnclosingBodyID.Int64)
                  if (_body) {
                    paragraph.EnclosingBody = _body
                  }
                }
                // insertion point for pointer field EnclosingTableColumn redeeming
                {
                  let _tablecolumn = this.frontRepo.TableColumns.get(paragraph.EnclosingTableColumnID.Int64)
                  if (_tablecolumn) {
                    paragraph.EnclosingTableColumn = _tablecolumn
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Body.Paragraphs redeeming
                {
                  let _body = this.frontRepo.Bodys.get(paragraph.Body_ParagraphsDBID.Int64)
                  if (_body) {
                    if (_body.Paragraphs == undefined) {
                      _body.Paragraphs = new Array<ParagraphDB>()
                    }
                    _body.Paragraphs.push(paragraph)
                    if (paragraph.Body_Paragraphs_reverse == undefined) {
                      paragraph.Body_Paragraphs_reverse = _body
                    }
                  }
                }
                // insertion point for slice of pointer field TableColumn.Paragraphs redeeming
                {
                  let _tablecolumn = this.frontRepo.TableColumns.get(paragraph.TableColumn_ParagraphsDBID.Int64)
                  if (_tablecolumn) {
                    if (_tablecolumn.Paragraphs == undefined) {
                      _tablecolumn.Paragraphs = new Array<ParagraphDB>()
                    }
                    _tablecolumn.Paragraphs.push(paragraph)
                    if (paragraph.TableColumn_Paragraphs_reverse == undefined) {
                      paragraph.TableColumn_Paragraphs_reverse = _tablecolumn
                    }
                  }
                }
              }
            )
            paragraphpropertiess.forEach(
              paragraphproperties => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field ParagraphStyle redeeming
                {
                  let _paragraphstyle = this.frontRepo.ParagraphStyles.get(paragraphproperties.ParagraphStyleID.Int64)
                  if (_paragraphstyle) {
                    paragraphproperties.ParagraphStyle = _paragraphstyle
                  }
                }
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(paragraphproperties.NodeID.Int64)
                  if (_node) {
                    paragraphproperties.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            paragraphstyles.forEach(
              paragraphstyle => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(paragraphstyle.NodeID.Int64)
                  if (_node) {
                    paragraphstyle.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            runes.forEach(
              rune => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(rune.NodeID.Int64)
                  if (_node) {
                    rune.Node = _node
                  }
                }
                // insertion point for pointer field Text redeeming
                {
                  let _text = this.frontRepo.Texts.get(rune.TextID.Int64)
                  if (_text) {
                    rune.Text = _text
                  }
                }
                // insertion point for pointer field RuneProperties redeeming
                {
                  let _runeproperties = this.frontRepo.RunePropertiess.get(rune.RunePropertiesID.Int64)
                  if (_runeproperties) {
                    rune.RuneProperties = _runeproperties
                  }
                }
                // insertion point for pointer field EnclosingParagraph redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(rune.EnclosingParagraphID.Int64)
                  if (_paragraph) {
                    rune.EnclosingParagraph = _paragraph
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Paragraph.Runes redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(rune.Paragraph_RunesDBID.Int64)
                  if (_paragraph) {
                    if (_paragraph.Runes == undefined) {
                      _paragraph.Runes = new Array<RuneDB>()
                    }
                    _paragraph.Runes.push(rune)
                    if (rune.Paragraph_Runes_reverse == undefined) {
                      rune.Paragraph_Runes_reverse = _paragraph
                    }
                  }
                }
              }
            )
            runepropertiess.forEach(
              runeproperties => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(runeproperties.NodeID.Int64)
                  if (_node) {
                    runeproperties.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            tables.forEach(
              table => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(table.NodeID.Int64)
                  if (_node) {
                    table.Node = _node
                  }
                }
                // insertion point for pointer field TableProperties redeeming
                {
                  let _tableproperties = this.frontRepo.TablePropertiess.get(table.TablePropertiesID.Int64)
                  if (_tableproperties) {
                    table.TableProperties = _tableproperties
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Body.Tables redeeming
                {
                  let _body = this.frontRepo.Bodys.get(table.Body_TablesDBID.Int64)
                  if (_body) {
                    if (_body.Tables == undefined) {
                      _body.Tables = new Array<TableDB>()
                    }
                    _body.Tables.push(table)
                    if (table.Body_Tables_reverse == undefined) {
                      table.Body_Tables_reverse = _body
                    }
                  }
                }
              }
            )
            tablecolumns.forEach(
              tablecolumn => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(tablecolumn.NodeID.Int64)
                  if (_node) {
                    tablecolumn.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field TableRow.TableColumns redeeming
                {
                  let _tablerow = this.frontRepo.TableRows.get(tablecolumn.TableRow_TableColumnsDBID.Int64)
                  if (_tablerow) {
                    if (_tablerow.TableColumns == undefined) {
                      _tablerow.TableColumns = new Array<TableColumnDB>()
                    }
                    _tablerow.TableColumns.push(tablecolumn)
                    if (tablecolumn.TableRow_TableColumns_reverse == undefined) {
                      tablecolumn.TableRow_TableColumns_reverse = _tablerow
                    }
                  }
                }
              }
            )
            tablepropertiess.forEach(
              tableproperties => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(tableproperties.NodeID.Int64)
                  if (_node) {
                    tableproperties.Node = _node
                  }
                }
                // insertion point for pointer field TableStyle redeeming
                {
                  let _tablestyle = this.frontRepo.TableStyles.get(tableproperties.TableStyleID.Int64)
                  if (_tablestyle) {
                    tableproperties.TableStyle = _tablestyle
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            tablerows.forEach(
              tablerow => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(tablerow.NodeID.Int64)
                  if (_node) {
                    tablerow.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Table.TableRows redeeming
                {
                  let _table = this.frontRepo.Tables.get(tablerow.Table_TableRowsDBID.Int64)
                  if (_table) {
                    if (_table.TableRows == undefined) {
                      _table.TableRows = new Array<TableRowDB>()
                    }
                    _table.TableRows.push(tablerow)
                    if (tablerow.Table_TableRows_reverse == undefined) {
                      tablerow.Table_TableRows_reverse = _table
                    }
                  }
                }
              }
            )
            tablestyles.forEach(
              tablestyle => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(tablestyle.NodeID.Int64)
                  if (_node) {
                    tablestyle.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            texts.forEach(
              text => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(text.NodeID.Int64)
                  if (_node) {
                    text.Node = _node
                  }
                }
                // insertion point for pointer field EnclosingRune redeeming
                {
                  let _rune = this.frontRepo.Runes.get(text.EnclosingRuneID.Int64)
                  if (_rune) {
                    text.EnclosingRune = _rune
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // BodyPull performs a GET on Body of the stack and redeem association pointers 
  BodyPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.bodyService.getBodys(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            bodys,
          ]) => {
            // init the array
            this.frontRepo.Bodys_array = bodys

            // clear the map that counts Body in the GET
            this.frontRepo.Bodys_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bodys.forEach(
              body => {
                this.frontRepo.Bodys.set(body.ID, body)
                this.frontRepo.Bodys_batch.set(body.ID, body)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field LastParagraph redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(body.LastParagraphID.Int64)
                  if (_paragraph) {
                    body.LastParagraph = _paragraph
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear bodys that are absent from the GET
            this.frontRepo.Bodys.forEach(
              body => {
                if (this.frontRepo.Bodys_batch.get(body.ID) == undefined) {
                  this.frontRepo.Bodys.delete(body.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // DocumentPull performs a GET on Document of the stack and redeem association pointers 
  DocumentPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.documentService.getDocuments(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            documents,
          ]) => {
            // init the array
            this.frontRepo.Documents_array = documents

            // clear the map that counts Document in the GET
            this.frontRepo.Documents_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            documents.forEach(
              document => {
                this.frontRepo.Documents.set(document.ID, document)
                this.frontRepo.Documents_batch.set(document.ID, document)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field File redeeming
                {
                  let _file = this.frontRepo.Files.get(document.FileID.Int64)
                  if (_file) {
                    document.File = _file
                  }
                }
                // insertion point for pointer field Root redeeming
                {
                  let _node = this.frontRepo.Nodes.get(document.RootID.Int64)
                  if (_node) {
                    document.Root = _node
                  }
                }
                // insertion point for pointer field Body redeeming
                {
                  let _body = this.frontRepo.Bodys.get(document.BodyID.Int64)
                  if (_body) {
                    document.Body = _body
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear documents that are absent from the GET
            this.frontRepo.Documents.forEach(
              document => {
                if (this.frontRepo.Documents_batch.get(document.ID) == undefined) {
                  this.frontRepo.Documents.delete(document.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // DocxPull performs a GET on Docx of the stack and redeem association pointers 
  DocxPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.docxService.getDocxs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            docxs,
          ]) => {
            // init the array
            this.frontRepo.Docxs_array = docxs

            // clear the map that counts Docx in the GET
            this.frontRepo.Docxs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            docxs.forEach(
              docx => {
                this.frontRepo.Docxs.set(docx.ID, docx)
                this.frontRepo.Docxs_batch.set(docx.ID, docx)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Document redeeming
                {
                  let _document = this.frontRepo.Documents.get(docx.DocumentID.Int64)
                  if (_document) {
                    docx.Document = _document
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear docxs that are absent from the GET
            this.frontRepo.Docxs.forEach(
              docx => {
                if (this.frontRepo.Docxs_batch.get(docx.ID) == undefined) {
                  this.frontRepo.Docxs.delete(docx.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FilePull performs a GET on File of the stack and redeem association pointers 
  FilePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.fileService.getFiles(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            files,
          ]) => {
            // init the array
            this.frontRepo.Files_array = files

            // clear the map that counts File in the GET
            this.frontRepo.Files_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            files.forEach(
              file => {
                this.frontRepo.Files.set(file.ID, file)
                this.frontRepo.Files_batch.set(file.ID, file)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Docx.Files redeeming
                {
                  let _docx = this.frontRepo.Docxs.get(file.Docx_FilesDBID.Int64)
                  if (_docx) {
                    if (_docx.Files == undefined) {
                      _docx.Files = new Array<FileDB>()
                    }
                    _docx.Files.push(file)
                    if (file.Docx_Files_reverse == undefined) {
                      file.Docx_Files_reverse = _docx
                    }
                  }
                }
              }
            )

            // clear files that are absent from the GET
            this.frontRepo.Files.forEach(
              file => {
                if (this.frontRepo.Files_batch.get(file.ID) == undefined) {
                  this.frontRepo.Files.delete(file.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // NodePull performs a GET on Node of the stack and redeem association pointers 
  NodePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.nodeService.getNodes(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            nodes,
          ]) => {
            // init the array
            this.frontRepo.Nodes_array = nodes

            // clear the map that counts Node in the GET
            this.frontRepo.Nodes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            nodes.forEach(
              node => {
                this.frontRepo.Nodes.set(node.ID, node)
                this.frontRepo.Nodes_batch.set(node.ID, node)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Node.Nodes redeeming
                {
                  let _node = this.frontRepo.Nodes.get(node.Node_NodesDBID.Int64)
                  if (_node) {
                    if (_node.Nodes == undefined) {
                      _node.Nodes = new Array<NodeDB>()
                    }
                    _node.Nodes.push(node)
                    if (node.Node_Nodes_reverse == undefined) {
                      node.Node_Nodes_reverse = _node
                    }
                  }
                }
              }
            )

            // clear nodes that are absent from the GET
            this.frontRepo.Nodes.forEach(
              node => {
                if (this.frontRepo.Nodes_batch.get(node.ID) == undefined) {
                  this.frontRepo.Nodes.delete(node.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // ParagraphPull performs a GET on Paragraph of the stack and redeem association pointers 
  ParagraphPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.paragraphService.getParagraphs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            paragraphs,
          ]) => {
            // init the array
            this.frontRepo.Paragraphs_array = paragraphs

            // clear the map that counts Paragraph in the GET
            this.frontRepo.Paragraphs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            paragraphs.forEach(
              paragraph => {
                this.frontRepo.Paragraphs.set(paragraph.ID, paragraph)
                this.frontRepo.Paragraphs_batch.set(paragraph.ID, paragraph)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(paragraph.NodeID.Int64)
                  if (_node) {
                    paragraph.Node = _node
                  }
                }
                // insertion point for pointer field ParagraphProperties redeeming
                {
                  let _paragraphproperties = this.frontRepo.ParagraphPropertiess.get(paragraph.ParagraphPropertiesID.Int64)
                  if (_paragraphproperties) {
                    paragraph.ParagraphProperties = _paragraphproperties
                  }
                }
                // insertion point for pointer field Next redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(paragraph.NextID.Int64)
                  if (_paragraph) {
                    paragraph.Next = _paragraph
                  }
                }
                // insertion point for pointer field Previous redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(paragraph.PreviousID.Int64)
                  if (_paragraph) {
                    paragraph.Previous = _paragraph
                  }
                }
                // insertion point for pointer field EnclosingBody redeeming
                {
                  let _body = this.frontRepo.Bodys.get(paragraph.EnclosingBodyID.Int64)
                  if (_body) {
                    paragraph.EnclosingBody = _body
                  }
                }
                // insertion point for pointer field EnclosingTableColumn redeeming
                {
                  let _tablecolumn = this.frontRepo.TableColumns.get(paragraph.EnclosingTableColumnID.Int64)
                  if (_tablecolumn) {
                    paragraph.EnclosingTableColumn = _tablecolumn
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Body.Paragraphs redeeming
                {
                  let _body = this.frontRepo.Bodys.get(paragraph.Body_ParagraphsDBID.Int64)
                  if (_body) {
                    if (_body.Paragraphs == undefined) {
                      _body.Paragraphs = new Array<ParagraphDB>()
                    }
                    _body.Paragraphs.push(paragraph)
                    if (paragraph.Body_Paragraphs_reverse == undefined) {
                      paragraph.Body_Paragraphs_reverse = _body
                    }
                  }
                }
                // insertion point for slice of pointer field TableColumn.Paragraphs redeeming
                {
                  let _tablecolumn = this.frontRepo.TableColumns.get(paragraph.TableColumn_ParagraphsDBID.Int64)
                  if (_tablecolumn) {
                    if (_tablecolumn.Paragraphs == undefined) {
                      _tablecolumn.Paragraphs = new Array<ParagraphDB>()
                    }
                    _tablecolumn.Paragraphs.push(paragraph)
                    if (paragraph.TableColumn_Paragraphs_reverse == undefined) {
                      paragraph.TableColumn_Paragraphs_reverse = _tablecolumn
                    }
                  }
                }
              }
            )

            // clear paragraphs that are absent from the GET
            this.frontRepo.Paragraphs.forEach(
              paragraph => {
                if (this.frontRepo.Paragraphs_batch.get(paragraph.ID) == undefined) {
                  this.frontRepo.Paragraphs.delete(paragraph.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // ParagraphPropertiesPull performs a GET on ParagraphProperties of the stack and redeem association pointers 
  ParagraphPropertiesPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.paragraphpropertiesService.getParagraphPropertiess(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            paragraphpropertiess,
          ]) => {
            // init the array
            this.frontRepo.ParagraphPropertiess_array = paragraphpropertiess

            // clear the map that counts ParagraphProperties in the GET
            this.frontRepo.ParagraphPropertiess_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            paragraphpropertiess.forEach(
              paragraphproperties => {
                this.frontRepo.ParagraphPropertiess.set(paragraphproperties.ID, paragraphproperties)
                this.frontRepo.ParagraphPropertiess_batch.set(paragraphproperties.ID, paragraphproperties)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field ParagraphStyle redeeming
                {
                  let _paragraphstyle = this.frontRepo.ParagraphStyles.get(paragraphproperties.ParagraphStyleID.Int64)
                  if (_paragraphstyle) {
                    paragraphproperties.ParagraphStyle = _paragraphstyle
                  }
                }
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(paragraphproperties.NodeID.Int64)
                  if (_node) {
                    paragraphproperties.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear paragraphpropertiess that are absent from the GET
            this.frontRepo.ParagraphPropertiess.forEach(
              paragraphproperties => {
                if (this.frontRepo.ParagraphPropertiess_batch.get(paragraphproperties.ID) == undefined) {
                  this.frontRepo.ParagraphPropertiess.delete(paragraphproperties.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // ParagraphStylePull performs a GET on ParagraphStyle of the stack and redeem association pointers 
  ParagraphStylePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.paragraphstyleService.getParagraphStyles(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            paragraphstyles,
          ]) => {
            // init the array
            this.frontRepo.ParagraphStyles_array = paragraphstyles

            // clear the map that counts ParagraphStyle in the GET
            this.frontRepo.ParagraphStyles_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            paragraphstyles.forEach(
              paragraphstyle => {
                this.frontRepo.ParagraphStyles.set(paragraphstyle.ID, paragraphstyle)
                this.frontRepo.ParagraphStyles_batch.set(paragraphstyle.ID, paragraphstyle)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(paragraphstyle.NodeID.Int64)
                  if (_node) {
                    paragraphstyle.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear paragraphstyles that are absent from the GET
            this.frontRepo.ParagraphStyles.forEach(
              paragraphstyle => {
                if (this.frontRepo.ParagraphStyles_batch.get(paragraphstyle.ID) == undefined) {
                  this.frontRepo.ParagraphStyles.delete(paragraphstyle.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RunePull performs a GET on Rune of the stack and redeem association pointers 
  RunePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.runeService.getRunes(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            runes,
          ]) => {
            // init the array
            this.frontRepo.Runes_array = runes

            // clear the map that counts Rune in the GET
            this.frontRepo.Runes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            runes.forEach(
              rune => {
                this.frontRepo.Runes.set(rune.ID, rune)
                this.frontRepo.Runes_batch.set(rune.ID, rune)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(rune.NodeID.Int64)
                  if (_node) {
                    rune.Node = _node
                  }
                }
                // insertion point for pointer field Text redeeming
                {
                  let _text = this.frontRepo.Texts.get(rune.TextID.Int64)
                  if (_text) {
                    rune.Text = _text
                  }
                }
                // insertion point for pointer field RuneProperties redeeming
                {
                  let _runeproperties = this.frontRepo.RunePropertiess.get(rune.RunePropertiesID.Int64)
                  if (_runeproperties) {
                    rune.RuneProperties = _runeproperties
                  }
                }
                // insertion point for pointer field EnclosingParagraph redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(rune.EnclosingParagraphID.Int64)
                  if (_paragraph) {
                    rune.EnclosingParagraph = _paragraph
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Paragraph.Runes redeeming
                {
                  let _paragraph = this.frontRepo.Paragraphs.get(rune.Paragraph_RunesDBID.Int64)
                  if (_paragraph) {
                    if (_paragraph.Runes == undefined) {
                      _paragraph.Runes = new Array<RuneDB>()
                    }
                    _paragraph.Runes.push(rune)
                    if (rune.Paragraph_Runes_reverse == undefined) {
                      rune.Paragraph_Runes_reverse = _paragraph
                    }
                  }
                }
              }
            )

            // clear runes that are absent from the GET
            this.frontRepo.Runes.forEach(
              rune => {
                if (this.frontRepo.Runes_batch.get(rune.ID) == undefined) {
                  this.frontRepo.Runes.delete(rune.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RunePropertiesPull performs a GET on RuneProperties of the stack and redeem association pointers 
  RunePropertiesPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.runepropertiesService.getRunePropertiess(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            runepropertiess,
          ]) => {
            // init the array
            this.frontRepo.RunePropertiess_array = runepropertiess

            // clear the map that counts RuneProperties in the GET
            this.frontRepo.RunePropertiess_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            runepropertiess.forEach(
              runeproperties => {
                this.frontRepo.RunePropertiess.set(runeproperties.ID, runeproperties)
                this.frontRepo.RunePropertiess_batch.set(runeproperties.ID, runeproperties)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(runeproperties.NodeID.Int64)
                  if (_node) {
                    runeproperties.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear runepropertiess that are absent from the GET
            this.frontRepo.RunePropertiess.forEach(
              runeproperties => {
                if (this.frontRepo.RunePropertiess_batch.get(runeproperties.ID) == undefined) {
                  this.frontRepo.RunePropertiess.delete(runeproperties.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TablePull performs a GET on Table of the stack and redeem association pointers 
  TablePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.tableService.getTables(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            tables,
          ]) => {
            // init the array
            this.frontRepo.Tables_array = tables

            // clear the map that counts Table in the GET
            this.frontRepo.Tables_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            tables.forEach(
              table => {
                this.frontRepo.Tables.set(table.ID, table)
                this.frontRepo.Tables_batch.set(table.ID, table)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(table.NodeID.Int64)
                  if (_node) {
                    table.Node = _node
                  }
                }
                // insertion point for pointer field TableProperties redeeming
                {
                  let _tableproperties = this.frontRepo.TablePropertiess.get(table.TablePropertiesID.Int64)
                  if (_tableproperties) {
                    table.TableProperties = _tableproperties
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Body.Tables redeeming
                {
                  let _body = this.frontRepo.Bodys.get(table.Body_TablesDBID.Int64)
                  if (_body) {
                    if (_body.Tables == undefined) {
                      _body.Tables = new Array<TableDB>()
                    }
                    _body.Tables.push(table)
                    if (table.Body_Tables_reverse == undefined) {
                      table.Body_Tables_reverse = _body
                    }
                  }
                }
              }
            )

            // clear tables that are absent from the GET
            this.frontRepo.Tables.forEach(
              table => {
                if (this.frontRepo.Tables_batch.get(table.ID) == undefined) {
                  this.frontRepo.Tables.delete(table.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TableColumnPull performs a GET on TableColumn of the stack and redeem association pointers 
  TableColumnPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.tablecolumnService.getTableColumns(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            tablecolumns,
          ]) => {
            // init the array
            this.frontRepo.TableColumns_array = tablecolumns

            // clear the map that counts TableColumn in the GET
            this.frontRepo.TableColumns_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            tablecolumns.forEach(
              tablecolumn => {
                this.frontRepo.TableColumns.set(tablecolumn.ID, tablecolumn)
                this.frontRepo.TableColumns_batch.set(tablecolumn.ID, tablecolumn)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(tablecolumn.NodeID.Int64)
                  if (_node) {
                    tablecolumn.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field TableRow.TableColumns redeeming
                {
                  let _tablerow = this.frontRepo.TableRows.get(tablecolumn.TableRow_TableColumnsDBID.Int64)
                  if (_tablerow) {
                    if (_tablerow.TableColumns == undefined) {
                      _tablerow.TableColumns = new Array<TableColumnDB>()
                    }
                    _tablerow.TableColumns.push(tablecolumn)
                    if (tablecolumn.TableRow_TableColumns_reverse == undefined) {
                      tablecolumn.TableRow_TableColumns_reverse = _tablerow
                    }
                  }
                }
              }
            )

            // clear tablecolumns that are absent from the GET
            this.frontRepo.TableColumns.forEach(
              tablecolumn => {
                if (this.frontRepo.TableColumns_batch.get(tablecolumn.ID) == undefined) {
                  this.frontRepo.TableColumns.delete(tablecolumn.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TablePropertiesPull performs a GET on TableProperties of the stack and redeem association pointers 
  TablePropertiesPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.tablepropertiesService.getTablePropertiess(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            tablepropertiess,
          ]) => {
            // init the array
            this.frontRepo.TablePropertiess_array = tablepropertiess

            // clear the map that counts TableProperties in the GET
            this.frontRepo.TablePropertiess_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            tablepropertiess.forEach(
              tableproperties => {
                this.frontRepo.TablePropertiess.set(tableproperties.ID, tableproperties)
                this.frontRepo.TablePropertiess_batch.set(tableproperties.ID, tableproperties)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(tableproperties.NodeID.Int64)
                  if (_node) {
                    tableproperties.Node = _node
                  }
                }
                // insertion point for pointer field TableStyle redeeming
                {
                  let _tablestyle = this.frontRepo.TableStyles.get(tableproperties.TableStyleID.Int64)
                  if (_tablestyle) {
                    tableproperties.TableStyle = _tablestyle
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear tablepropertiess that are absent from the GET
            this.frontRepo.TablePropertiess.forEach(
              tableproperties => {
                if (this.frontRepo.TablePropertiess_batch.get(tableproperties.ID) == undefined) {
                  this.frontRepo.TablePropertiess.delete(tableproperties.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TableRowPull performs a GET on TableRow of the stack and redeem association pointers 
  TableRowPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.tablerowService.getTableRows(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            tablerows,
          ]) => {
            // init the array
            this.frontRepo.TableRows_array = tablerows

            // clear the map that counts TableRow in the GET
            this.frontRepo.TableRows_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            tablerows.forEach(
              tablerow => {
                this.frontRepo.TableRows.set(tablerow.ID, tablerow)
                this.frontRepo.TableRows_batch.set(tablerow.ID, tablerow)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(tablerow.NodeID.Int64)
                  if (_node) {
                    tablerow.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Table.TableRows redeeming
                {
                  let _table = this.frontRepo.Tables.get(tablerow.Table_TableRowsDBID.Int64)
                  if (_table) {
                    if (_table.TableRows == undefined) {
                      _table.TableRows = new Array<TableRowDB>()
                    }
                    _table.TableRows.push(tablerow)
                    if (tablerow.Table_TableRows_reverse == undefined) {
                      tablerow.Table_TableRows_reverse = _table
                    }
                  }
                }
              }
            )

            // clear tablerows that are absent from the GET
            this.frontRepo.TableRows.forEach(
              tablerow => {
                if (this.frontRepo.TableRows_batch.get(tablerow.ID) == undefined) {
                  this.frontRepo.TableRows.delete(tablerow.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TableStylePull performs a GET on TableStyle of the stack and redeem association pointers 
  TableStylePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.tablestyleService.getTableStyles(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            tablestyles,
          ]) => {
            // init the array
            this.frontRepo.TableStyles_array = tablestyles

            // clear the map that counts TableStyle in the GET
            this.frontRepo.TableStyles_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            tablestyles.forEach(
              tablestyle => {
                this.frontRepo.TableStyles.set(tablestyle.ID, tablestyle)
                this.frontRepo.TableStyles_batch.set(tablestyle.ID, tablestyle)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(tablestyle.NodeID.Int64)
                  if (_node) {
                    tablestyle.Node = _node
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear tablestyles that are absent from the GET
            this.frontRepo.TableStyles.forEach(
              tablestyle => {
                if (this.frontRepo.TableStyles_batch.get(tablestyle.ID) == undefined) {
                  this.frontRepo.TableStyles.delete(tablestyle.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // TextPull performs a GET on Text of the stack and redeem association pointers 
  TextPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.textService.getTexts(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            texts,
          ]) => {
            // init the array
            this.frontRepo.Texts_array = texts

            // clear the map that counts Text in the GET
            this.frontRepo.Texts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            texts.forEach(
              text => {
                this.frontRepo.Texts.set(text.ID, text)
                this.frontRepo.Texts_batch.set(text.ID, text)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Node redeeming
                {
                  let _node = this.frontRepo.Nodes.get(text.NodeID.Int64)
                  if (_node) {
                    text.Node = _node
                  }
                }
                // insertion point for pointer field EnclosingRune redeeming
                {
                  let _rune = this.frontRepo.Runes.get(text.EnclosingRuneID.Int64)
                  if (_rune) {
                    text.EnclosingRune = _rune
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear texts that are absent from the GET
            this.frontRepo.Texts.forEach(
              text => {
                if (this.frontRepo.Texts_batch.get(text.ID) == undefined) {
                  this.frontRepo.Texts.delete(text.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getBodyUniqueID(id: number): number {
  return 31 * id
}
export function getDocumentUniqueID(id: number): number {
  return 37 * id
}
export function getDocxUniqueID(id: number): number {
  return 41 * id
}
export function getFileUniqueID(id: number): number {
  return 43 * id
}
export function getNodeUniqueID(id: number): number {
  return 47 * id
}
export function getParagraphUniqueID(id: number): number {
  return 53 * id
}
export function getParagraphPropertiesUniqueID(id: number): number {
  return 59 * id
}
export function getParagraphStyleUniqueID(id: number): number {
  return 61 * id
}
export function getRuneUniqueID(id: number): number {
  return 67 * id
}
export function getRunePropertiesUniqueID(id: number): number {
  return 71 * id
}
export function getTableUniqueID(id: number): number {
  return 73 * id
}
export function getTableColumnUniqueID(id: number): number {
  return 79 * id
}
export function getTablePropertiesUniqueID(id: number): number {
  return 83 * id
}
export function getTableRowUniqueID(id: number): number {
  return 89 * id
}
export function getTableStyleUniqueID(id: number): number {
  return 97 * id
}
export function getTextUniqueID(id: number): number {
  return 101 * id
}
