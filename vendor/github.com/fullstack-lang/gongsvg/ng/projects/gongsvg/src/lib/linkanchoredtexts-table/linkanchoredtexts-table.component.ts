// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { ActivatedRoute, Router, RouterState } from '@angular/router';
import { LinkAnchoredTextDB } from '../linkanchoredtext-db'
import { LinkAnchoredTextService } from '../linkanchoredtext.service'

// insertion point for additional imports

import { RouteService } from '../route-service';

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-linkanchoredtextstable',
  templateUrl: './linkanchoredtexts-table.component.html',
  styleUrls: ['./linkanchoredtexts-table.component.css'],
})
export class LinkAnchoredTextsTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of LinkAnchoredText instances
  selection: SelectionModel<LinkAnchoredTextDB> = new (SelectionModel)
  initialSelection = new Array<LinkAnchoredTextDB>()

  // the data source for the table
  linkanchoredtexts: LinkAnchoredTextDB[] = []
  matTableDataSource: MatTableDataSource<LinkAnchoredTextDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.linkanchoredtexts
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (linkanchoredtextDB: LinkAnchoredTextDB, property: string) => {
      switch (property) {
        case 'ID':
          return linkanchoredtextDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return linkanchoredtextDB.Name;

        case 'Content':
          return linkanchoredtextDB.Content;

        case 'X_Offset':
          return linkanchoredtextDB.X_Offset;

        case 'Y_Offset':
          return linkanchoredtextDB.Y_Offset;

        case 'FontWeight':
          return linkanchoredtextDB.FontWeight;

        case 'Color':
          return linkanchoredtextDB.Color;

        case 'FillOpacity':
          return linkanchoredtextDB.FillOpacity;

        case 'Stroke':
          return linkanchoredtextDB.Stroke;

        case 'StrokeWidth':
          return linkanchoredtextDB.StrokeWidth;

        case 'StrokeDashArray':
          return linkanchoredtextDB.StrokeDashArray;

        case 'StrokeDashArrayWhenSelected':
          return linkanchoredtextDB.StrokeDashArrayWhenSelected;

        case 'Transform':
          return linkanchoredtextDB.Transform;

        case 'Link_TextAtArrowEnd':
          if (this.frontRepo.Links.get(linkanchoredtextDB.Link_TextAtArrowEndDBID.Int64) != undefined) {
            return this.frontRepo.Links.get(linkanchoredtextDB.Link_TextAtArrowEndDBID.Int64)!.Name
          } else {
            return ""
          }

        case 'Link_TextAtArrowStart':
          if (this.frontRepo.Links.get(linkanchoredtextDB.Link_TextAtArrowStartDBID.Int64) != undefined) {
            return this.frontRepo.Links.get(linkanchoredtextDB.Link_TextAtArrowStartDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (linkanchoredtextDB: LinkAnchoredTextDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the linkanchoredtextDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += linkanchoredtextDB.Name.toLowerCase()
      mergedContent += linkanchoredtextDB.Content.toLowerCase()
      mergedContent += linkanchoredtextDB.X_Offset.toString()
      mergedContent += linkanchoredtextDB.Y_Offset.toString()
      mergedContent += linkanchoredtextDB.FontWeight.toLowerCase()
      mergedContent += linkanchoredtextDB.Color.toLowerCase()
      mergedContent += linkanchoredtextDB.FillOpacity.toString()
      mergedContent += linkanchoredtextDB.Stroke.toLowerCase()
      mergedContent += linkanchoredtextDB.StrokeWidth.toString()
      mergedContent += linkanchoredtextDB.StrokeDashArray.toLowerCase()
      mergedContent += linkanchoredtextDB.StrokeDashArrayWhenSelected.toLowerCase()
      mergedContent += linkanchoredtextDB.Transform.toLowerCase()
      if (linkanchoredtextDB.Link_TextAtArrowEndDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Links.get(linkanchoredtextDB.Link_TextAtArrowEndDBID.Int64)!.Name.toLowerCase()
      }

      if (linkanchoredtextDB.Link_TextAtArrowStartDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Links.get(linkanchoredtextDB.Link_TextAtArrowStartDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private linkanchoredtextService: LinkAnchoredTextService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of linkanchoredtext instances
    public dialogRef: MatDialogRef<LinkAnchoredTextsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
    private activatedRoute: ActivatedRoute,

    private routeService: RouteService,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      this.GONG__StackPath = dialogData.GONG__StackPath
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.linkanchoredtextService.LinkAnchoredTextServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getLinkAnchoredTexts()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Content",
        "X_Offset",
        "Y_Offset",
        "FontWeight",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "StrokeDashArrayWhenSelected",
        "Transform",
        "Link_TextAtArrowEnd",
        "Link_TextAtArrowStart",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Content",
        "X_Offset",
        "Y_Offset",
        "FontWeight",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "StrokeDashArrayWhenSelected",
        "Transform",
        "Link_TextAtArrowEnd",
        "Link_TextAtArrowStart",
      ]
      this.selection = new SelectionModel<LinkAnchoredTextDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getLinkAnchoredTexts()

    this.matTableDataSource = new MatTableDataSource(this.linkanchoredtexts)
  }

  getLinkAnchoredTexts(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.linkanchoredtexts = this.frontRepo.LinkAnchoredTexts_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let linkanchoredtext of this.linkanchoredtexts) {
            let ID = this.dialogData.ID
            let revPointer = linkanchoredtext[this.dialogData.ReversePointer as keyof LinkAnchoredTextDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(linkanchoredtext)
            }
            this.selection = new SelectionModel<LinkAnchoredTextDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, LinkAnchoredTextDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to LinkAnchoredTextDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as LinkAnchoredTextDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let linkanchoredtext = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as LinkAnchoredTextDB
              this.initialSelection.push(linkanchoredtext)
            }
          }

          this.selection = new SelectionModel<LinkAnchoredTextDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.linkanchoredtexts
      }
    )
  }

  // newLinkAnchoredText initiate a new linkanchoredtext
  // create a new LinkAnchoredText objet
  newLinkAnchoredText() {
  }

  deleteLinkAnchoredText(linkanchoredtextID: number, linkanchoredtext: LinkAnchoredTextDB) {
    // list of linkanchoredtexts is truncated of linkanchoredtext before the delete
    this.linkanchoredtexts = this.linkanchoredtexts.filter(h => h !== linkanchoredtext);

    this.linkanchoredtextService.deleteLinkAnchoredText(linkanchoredtextID, this.GONG__StackPath).subscribe(
      linkanchoredtext => {
        this.linkanchoredtextService.LinkAnchoredTextServiceChanged.next("delete")
      }
    );
  }

  editLinkAnchoredText(linkanchoredtextID: number, linkanchoredtext: LinkAnchoredTextDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(linkanchoredtextID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "linkanchoredtext" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, linkanchoredtextID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.linkanchoredtexts.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.linkanchoredtexts.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<LinkAnchoredTextDB>()

      // reset all initial selection of linkanchoredtext that belong to linkanchoredtext
      for (let linkanchoredtext of this.initialSelection) {
        let index = linkanchoredtext[this.dialogData.ReversePointer as keyof LinkAnchoredTextDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(linkanchoredtext)

      }

      // from selection, set linkanchoredtext that belong to linkanchoredtext
      for (let linkanchoredtext of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = linkanchoredtext[this.dialogData.ReversePointer as keyof LinkAnchoredTextDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(linkanchoredtext)
      }


      // update all linkanchoredtext (only update selection & initial selection)
      for (let linkanchoredtext of toUpdate) {
        this.linkanchoredtextService.updateLinkAnchoredText(linkanchoredtext, this.GONG__StackPath)
          .subscribe(linkanchoredtext => {
            this.linkanchoredtextService.LinkAnchoredTextServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, LinkAnchoredTextDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedLinkAnchoredText = new Set<number>()
      for (let linkanchoredtext of this.initialSelection) {
        if (this.selection.selected.includes(linkanchoredtext)) {
          // console.log("linkanchoredtext " + linkanchoredtext.Name + " is still selected")
        } else {
          console.log("linkanchoredtext " + linkanchoredtext.Name + " has been unselected")
          unselectedLinkAnchoredText.add(linkanchoredtext.ID)
          console.log("is unselected " + unselectedLinkAnchoredText.has(linkanchoredtext.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let linkanchoredtext = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as LinkAnchoredTextDB
      if (unselectedLinkAnchoredText.has(linkanchoredtext.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<LinkAnchoredTextDB>) = new Array<LinkAnchoredTextDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          linkanchoredtext => {
            if (!this.initialSelection.includes(linkanchoredtext)) {
              // console.log("linkanchoredtext " + linkanchoredtext.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + linkanchoredtext.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = linkanchoredtext.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = linkanchoredtext.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("linkanchoredtext " + linkanchoredtext.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<LinkAnchoredTextDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
