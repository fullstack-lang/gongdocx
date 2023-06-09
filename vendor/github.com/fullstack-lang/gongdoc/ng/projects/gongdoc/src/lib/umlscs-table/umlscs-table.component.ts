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
import { UmlscDB } from '../umlsc-db'
import { UmlscService } from '../umlsc.service'

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
  selector: 'app-umlscstable',
  templateUrl: './umlscs-table.component.html',
  styleUrls: ['./umlscs-table.component.css'],
})
export class UmlscsTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Umlsc instances
  selection: SelectionModel<UmlscDB> = new (SelectionModel)
  initialSelection = new Array<UmlscDB>()

  // the data source for the table
  umlscs: UmlscDB[] = []
  matTableDataSource: MatTableDataSource<UmlscDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.umlscs
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
    this.matTableDataSource.sortingDataAccessor = (umlscDB: UmlscDB, property: string) => {
      switch (property) {
        case 'ID':
          return umlscDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return umlscDB.Name;

        case 'Activestate':
          return umlscDB.Activestate;

        case 'IsInDrawMode':
          return umlscDB.IsInDrawMode ? "true" : "false";

        case 'DiagramPackage_Umlscs':
          if (this.frontRepo.DiagramPackages.get(umlscDB.DiagramPackage_UmlscsDBID.Int64) != undefined) {
            return this.frontRepo.DiagramPackages.get(umlscDB.DiagramPackage_UmlscsDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (umlscDB: UmlscDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the umlscDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += umlscDB.Name.toLowerCase()
      mergedContent += umlscDB.Activestate.toLowerCase()
      if (umlscDB.DiagramPackage_UmlscsDBID.Int64 != 0) {
        mergedContent += this.frontRepo.DiagramPackages.get(umlscDB.DiagramPackage_UmlscsDBID.Int64)!.Name.toLowerCase()
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
    private umlscService: UmlscService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of umlsc instances
    public dialogRef: MatDialogRef<UmlscsTableComponent>,
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
    this.umlscService.UmlscServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getUmlscs()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Activestate",
        "IsInDrawMode",
        "DiagramPackage_Umlscs",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Activestate",
        "IsInDrawMode",
        "DiagramPackage_Umlscs",
      ]
      this.selection = new SelectionModel<UmlscDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getUmlscs()

    this.matTableDataSource = new MatTableDataSource(this.umlscs)
  }

  getUmlscs(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.umlscs = this.frontRepo.Umlscs_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let umlsc of this.umlscs) {
            let ID = this.dialogData.ID
            let revPointer = umlsc[this.dialogData.ReversePointer as keyof UmlscDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(umlsc)
            }
            this.selection = new SelectionModel<UmlscDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, UmlscDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to UmlscDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as UmlscDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let umlsc = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as UmlscDB
              this.initialSelection.push(umlsc)
            }
          }

          this.selection = new SelectionModel<UmlscDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.umlscs
      }
    )
  }

  // newUmlsc initiate a new umlsc
  // create a new Umlsc objet
  newUmlsc() {
  }

  deleteUmlsc(umlscID: number, umlsc: UmlscDB) {
    // list of umlscs is truncated of umlsc before the delete
    this.umlscs = this.umlscs.filter(h => h !== umlsc);

    this.umlscService.deleteUmlsc(umlscID, this.GONG__StackPath).subscribe(
      umlsc => {
        this.umlscService.UmlscServiceChanged.next("delete")
      }
    );
  }

  editUmlsc(umlscID: number, umlsc: UmlscDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(umlscID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "umlsc" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, umlscID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.umlscs.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.umlscs.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<UmlscDB>()

      // reset all initial selection of umlsc that belong to umlsc
      for (let umlsc of this.initialSelection) {
        let index = umlsc[this.dialogData.ReversePointer as keyof UmlscDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(umlsc)

      }

      // from selection, set umlsc that belong to umlsc
      for (let umlsc of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = umlsc[this.dialogData.ReversePointer as keyof UmlscDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(umlsc)
      }


      // update all umlsc (only update selection & initial selection)
      for (let umlsc of toUpdate) {
        this.umlscService.updateUmlsc(umlsc, this.GONG__StackPath)
          .subscribe(umlsc => {
            this.umlscService.UmlscServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, UmlscDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedUmlsc = new Set<number>()
      for (let umlsc of this.initialSelection) {
        if (this.selection.selected.includes(umlsc)) {
          // console.log("umlsc " + umlsc.Name + " is still selected")
        } else {
          console.log("umlsc " + umlsc.Name + " has been unselected")
          unselectedUmlsc.add(umlsc.ID)
          console.log("is unselected " + unselectedUmlsc.has(umlsc.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let umlsc = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as UmlscDB
      if (unselectedUmlsc.has(umlsc.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<UmlscDB>) = new Array<UmlscDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          umlsc => {
            if (!this.initialSelection.includes(umlsc)) {
              // console.log("umlsc " + umlsc.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + umlsc.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = umlsc.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = umlsc.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("umlsc " + umlsc.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<UmlscDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
