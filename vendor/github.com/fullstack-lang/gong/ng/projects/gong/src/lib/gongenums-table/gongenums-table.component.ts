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
import { GongEnumDB } from '../gongenum-db'
import { GongEnumService } from '../gongenum.service'

// insertion point for additional imports
import { GongEnumTypeList } from '../GongEnumType'

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
  selector: 'app-gongenumstable',
  templateUrl: './gongenums-table.component.html',
  styleUrls: ['./gongenums-table.component.css'],
})
export class GongEnumsTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of GongEnum instances
  selection: SelectionModel<GongEnumDB> = new (SelectionModel)
  initialSelection = new Array<GongEnumDB>()

  // the data source for the table
  gongenums: GongEnumDB[] = []
  matTableDataSource: MatTableDataSource<GongEnumDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.gongenums
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
    this.matTableDataSource.sortingDataAccessor = (gongenumDB: GongEnumDB, property: string) => {
      switch (property) {
        case 'ID':
          return gongenumDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return gongenumDB.Name;

        case 'Type':
          return gongenumDB.Type_string!;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (gongenumDB: GongEnumDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the gongenumDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += gongenumDB.Name.toLowerCase()
      mergedContent += gongenumDB.Type_string!

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
    private gongenumService: GongEnumService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongenum instances
    public dialogRef: MatDialogRef<GongEnumsTableComponent>,
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
    this.gongenumService.GongEnumServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongEnums()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Type",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Type",
      ]
      this.selection = new SelectionModel<GongEnumDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getGongEnums()

    this.matTableDataSource = new MatTableDataSource(this.gongenums)
  }

  getGongEnums(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.gongenums = this.frontRepo.GongEnums_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        for (let gongenum of this.gongenums) {
          gongenum.Type_string = GongEnumTypeList[gongenum.Type].viewValue
        }

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let gongenum of this.gongenums) {
            let ID = this.dialogData.ID
            let revPointer = gongenum[this.dialogData.ReversePointer as keyof GongEnumDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(gongenum)
            }
            this.selection = new SelectionModel<GongEnumDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongEnumDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to GongEnumDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as GongEnumDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let gongenum = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongEnumDB
              this.initialSelection.push(gongenum)
            }
          }

          this.selection = new SelectionModel<GongEnumDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.gongenums
      }
    )
  }

  // newGongEnum initiate a new gongenum
  // create a new GongEnum objet
  newGongEnum() {
  }

  deleteGongEnum(gongenumID: number, gongenum: GongEnumDB) {
    // list of gongenums is truncated of gongenum before the delete
    this.gongenums = this.gongenums.filter(h => h !== gongenum);

    this.gongenumService.deleteGongEnum(gongenumID, this.GONG__StackPath).subscribe(
      gongenum => {
        this.gongenumService.GongEnumServiceChanged.next("delete")
      }
    );
  }

  editGongEnum(gongenumID: number, gongenum: GongEnumDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(gongenumID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "gongenum" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, gongenumID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongenums.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongenums.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<GongEnumDB>()

      // reset all initial selection of gongenum that belong to gongenum
      for (let gongenum of this.initialSelection) {
        let index = gongenum[this.dialogData.ReversePointer as keyof GongEnumDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(gongenum)

      }

      // from selection, set gongenum that belong to gongenum
      for (let gongenum of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = gongenum[this.dialogData.ReversePointer as keyof GongEnumDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(gongenum)
      }


      // update all gongenum (only update selection & initial selection)
      for (let gongenum of toUpdate) {
        this.gongenumService.updateGongEnum(gongenum, this.GONG__StackPath)
          .subscribe(gongenum => {
            this.gongenumService.GongEnumServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongEnumDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedGongEnum = new Set<number>()
      for (let gongenum of this.initialSelection) {
        if (this.selection.selected.includes(gongenum)) {
          // console.log("gongenum " + gongenum.Name + " is still selected")
        } else {
          console.log("gongenum " + gongenum.Name + " has been unselected")
          unselectedGongEnum.add(gongenum.ID)
          console.log("is unselected " + unselectedGongEnum.has(gongenum.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let gongenum = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongEnumDB
      if (unselectedGongEnum.has(gongenum.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<GongEnumDB>) = new Array<GongEnumDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          gongenum => {
            if (!this.initialSelection.includes(gongenum)) {
              // console.log("gongenum " + gongenum.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + gongenum.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = gongenum.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = gongenum.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("gongenum " + gongenum.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<GongEnumDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
