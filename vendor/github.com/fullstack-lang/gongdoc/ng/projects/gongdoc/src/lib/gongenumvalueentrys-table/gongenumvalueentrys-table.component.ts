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
import { GongEnumValueEntryDB } from '../gongenumvalueentry-db'
import { GongEnumValueEntryService } from '../gongenumvalueentry.service'

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
  selector: 'app-gongenumvalueentrystable',
  templateUrl: './gongenumvalueentrys-table.component.html',
  styleUrls: ['./gongenumvalueentrys-table.component.css'],
})
export class GongEnumValueEntrysTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of GongEnumValueEntry instances
  selection: SelectionModel<GongEnumValueEntryDB> = new (SelectionModel)
  initialSelection = new Array<GongEnumValueEntryDB>()

  // the data source for the table
  gongenumvalueentrys: GongEnumValueEntryDB[] = []
  matTableDataSource: MatTableDataSource<GongEnumValueEntryDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.gongenumvalueentrys
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
    this.matTableDataSource.sortingDataAccessor = (gongenumvalueentryDB: GongEnumValueEntryDB, property: string) => {
      switch (property) {
        case 'ID':
          return gongenumvalueentryDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return gongenumvalueentryDB.Name;

        case 'Identifier':
          return gongenumvalueentryDB.Identifier;

        case 'GongEnumShape_GongEnumValueEntrys':
          if (this.frontRepo.GongEnumShapes.get(gongenumvalueentryDB.GongEnumShape_GongEnumValueEntrysDBID.Int64) != undefined) {
            return this.frontRepo.GongEnumShapes.get(gongenumvalueentryDB.GongEnumShape_GongEnumValueEntrysDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (gongenumvalueentryDB: GongEnumValueEntryDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the gongenumvalueentryDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += gongenumvalueentryDB.Name.toLowerCase()
      mergedContent += gongenumvalueentryDB.Identifier.toLowerCase()
      if (gongenumvalueentryDB.GongEnumShape_GongEnumValueEntrysDBID.Int64 != 0) {
        mergedContent += this.frontRepo.GongEnumShapes.get(gongenumvalueentryDB.GongEnumShape_GongEnumValueEntrysDBID.Int64)!.Name.toLowerCase()
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
    private gongenumvalueentryService: GongEnumValueEntryService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongenumvalueentry instances
    public dialogRef: MatDialogRef<GongEnumValueEntrysTableComponent>,
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
    this.gongenumvalueentryService.GongEnumValueEntryServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongEnumValueEntrys()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Identifier",
        "GongEnumShape_GongEnumValueEntrys",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Identifier",
        "GongEnumShape_GongEnumValueEntrys",
      ]
      this.selection = new SelectionModel<GongEnumValueEntryDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getGongEnumValueEntrys()

    this.matTableDataSource = new MatTableDataSource(this.gongenumvalueentrys)
  }

  getGongEnumValueEntrys(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.gongenumvalueentrys = this.frontRepo.GongEnumValueEntrys_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let gongenumvalueentry of this.gongenumvalueentrys) {
            let ID = this.dialogData.ID
            let revPointer = gongenumvalueentry[this.dialogData.ReversePointer as keyof GongEnumValueEntryDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(gongenumvalueentry)
            }
            this.selection = new SelectionModel<GongEnumValueEntryDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongEnumValueEntryDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to GongEnumValueEntryDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as GongEnumValueEntryDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let gongenumvalueentry = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongEnumValueEntryDB
              this.initialSelection.push(gongenumvalueentry)
            }
          }

          this.selection = new SelectionModel<GongEnumValueEntryDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.gongenumvalueentrys
      }
    )
  }

  // newGongEnumValueEntry initiate a new gongenumvalueentry
  // create a new GongEnumValueEntry objet
  newGongEnumValueEntry() {
  }

  deleteGongEnumValueEntry(gongenumvalueentryID: number, gongenumvalueentry: GongEnumValueEntryDB) {
    // list of gongenumvalueentrys is truncated of gongenumvalueentry before the delete
    this.gongenumvalueentrys = this.gongenumvalueentrys.filter(h => h !== gongenumvalueentry);

    this.gongenumvalueentryService.deleteGongEnumValueEntry(gongenumvalueentryID, this.GONG__StackPath).subscribe(
      gongenumvalueentry => {
        this.gongenumvalueentryService.GongEnumValueEntryServiceChanged.next("delete")
      }
    );
  }

  editGongEnumValueEntry(gongenumvalueentryID: number, gongenumvalueentry: GongEnumValueEntryDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(gongenumvalueentryID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "gongenumvalueentry" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, gongenumvalueentryID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongenumvalueentrys.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongenumvalueentrys.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<GongEnumValueEntryDB>()

      // reset all initial selection of gongenumvalueentry that belong to gongenumvalueentry
      for (let gongenumvalueentry of this.initialSelection) {
        let index = gongenumvalueentry[this.dialogData.ReversePointer as keyof GongEnumValueEntryDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(gongenumvalueentry)

      }

      // from selection, set gongenumvalueentry that belong to gongenumvalueentry
      for (let gongenumvalueentry of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = gongenumvalueentry[this.dialogData.ReversePointer as keyof GongEnumValueEntryDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(gongenumvalueentry)
      }


      // update all gongenumvalueentry (only update selection & initial selection)
      for (let gongenumvalueentry of toUpdate) {
        this.gongenumvalueentryService.updateGongEnumValueEntry(gongenumvalueentry, this.GONG__StackPath)
          .subscribe(gongenumvalueentry => {
            this.gongenumvalueentryService.GongEnumValueEntryServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongEnumValueEntryDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedGongEnumValueEntry = new Set<number>()
      for (let gongenumvalueentry of this.initialSelection) {
        if (this.selection.selected.includes(gongenumvalueentry)) {
          // console.log("gongenumvalueentry " + gongenumvalueentry.Name + " is still selected")
        } else {
          console.log("gongenumvalueentry " + gongenumvalueentry.Name + " has been unselected")
          unselectedGongEnumValueEntry.add(gongenumvalueentry.ID)
          console.log("is unselected " + unselectedGongEnumValueEntry.has(gongenumvalueentry.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let gongenumvalueentry = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongEnumValueEntryDB
      if (unselectedGongEnumValueEntry.has(gongenumvalueentry.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<GongEnumValueEntryDB>) = new Array<GongEnumValueEntryDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          gongenumvalueentry => {
            if (!this.initialSelection.includes(gongenumvalueentry)) {
              // console.log("gongenumvalueentry " + gongenumvalueentry.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + gongenumvalueentry.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = gongenumvalueentry.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = gongenumvalueentry.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("gongenumvalueentry " + gongenumvalueentry.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<GongEnumValueEntryDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
