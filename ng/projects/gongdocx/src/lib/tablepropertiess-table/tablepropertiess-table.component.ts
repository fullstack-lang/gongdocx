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
import { TablePropertiesDB } from '../tableproperties-db'
import { TablePropertiesService } from '../tableproperties.service'

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
  selector: 'app-tablepropertiesstable',
  templateUrl: './tablepropertiess-table.component.html',
  styleUrls: ['./tablepropertiess-table.component.css'],
})
export class TablePropertiessTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of TableProperties instances
  selection: SelectionModel<TablePropertiesDB> = new (SelectionModel)
  initialSelection = new Array<TablePropertiesDB>()

  // the data source for the table
  tablepropertiess: TablePropertiesDB[] = []
  matTableDataSource: MatTableDataSource<TablePropertiesDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.tablepropertiess
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
    this.matTableDataSource.sortingDataAccessor = (tablepropertiesDB: TablePropertiesDB, property: string) => {
      switch (property) {
        case 'ID':
          return tablepropertiesDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return tablepropertiesDB.Name;

        case 'Node':
          return (tablepropertiesDB.Node ? tablepropertiesDB.Node.Name : '');

        case 'Content':
          return tablepropertiesDB.Content;

        case 'TableStyle':
          return (tablepropertiesDB.TableStyle ? tablepropertiesDB.TableStyle.Name : '');

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (tablepropertiesDB: TablePropertiesDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the tablepropertiesDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += tablepropertiesDB.Name.toLowerCase()
      if (tablepropertiesDB.Node) {
        mergedContent += tablepropertiesDB.Node.Name.toLowerCase()
      }
      mergedContent += tablepropertiesDB.Content.toLowerCase()
      if (tablepropertiesDB.TableStyle) {
        mergedContent += tablepropertiesDB.TableStyle.Name.toLowerCase()
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
    private tablepropertiesService: TablePropertiesService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of tableproperties instances
    public dialogRef: MatDialogRef<TablePropertiessTableComponent>,
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
    this.tablepropertiesService.TablePropertiesServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getTablePropertiess()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Node",
        "Content",
        "TableStyle",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Node",
        "Content",
        "TableStyle",
      ]
      this.selection = new SelectionModel<TablePropertiesDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getTablePropertiess()

    this.matTableDataSource = new MatTableDataSource(this.tablepropertiess)
  }

  getTablePropertiess(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.tablepropertiess = this.frontRepo.TablePropertiess_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let tableproperties of this.tablepropertiess) {
            let ID = this.dialogData.ID
            let revPointer = tableproperties[this.dialogData.ReversePointer as keyof TablePropertiesDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(tableproperties)
            }
            this.selection = new SelectionModel<TablePropertiesDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, TablePropertiesDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to TablePropertiesDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as TablePropertiesDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let tableproperties = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as TablePropertiesDB
              this.initialSelection.push(tableproperties)
            }
          }

          this.selection = new SelectionModel<TablePropertiesDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.tablepropertiess
      }
    )
  }

  // newTableProperties initiate a new tableproperties
  // create a new TableProperties objet
  newTableProperties() {
  }

  deleteTableProperties(tablepropertiesID: number, tableproperties: TablePropertiesDB) {
    // list of tablepropertiess is truncated of tableproperties before the delete
    this.tablepropertiess = this.tablepropertiess.filter(h => h !== tableproperties);

    this.tablepropertiesService.deleteTableProperties(tablepropertiesID, this.GONG__StackPath).subscribe(
      tableproperties => {
        this.tablepropertiesService.TablePropertiesServiceChanged.next("delete")
      }
    );
  }

  editTableProperties(tablepropertiesID: number, tableproperties: TablePropertiesDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(tablepropertiesID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "tableproperties" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, tablepropertiesID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.tablepropertiess.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.tablepropertiess.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<TablePropertiesDB>()

      // reset all initial selection of tableproperties that belong to tableproperties
      for (let tableproperties of this.initialSelection) {
        let index = tableproperties[this.dialogData.ReversePointer as keyof TablePropertiesDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(tableproperties)

      }

      // from selection, set tableproperties that belong to tableproperties
      for (let tableproperties of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = tableproperties[this.dialogData.ReversePointer as keyof TablePropertiesDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(tableproperties)
      }


      // update all tableproperties (only update selection & initial selection)
      for (let tableproperties of toUpdate) {
        this.tablepropertiesService.updateTableProperties(tableproperties, this.GONG__StackPath)
          .subscribe(tableproperties => {
            this.tablepropertiesService.TablePropertiesServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, TablePropertiesDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedTableProperties = new Set<number>()
      for (let tableproperties of this.initialSelection) {
        if (this.selection.selected.includes(tableproperties)) {
          // console.log("tableproperties " + tableproperties.Name + " is still selected")
        } else {
          console.log("tableproperties " + tableproperties.Name + " has been unselected")
          unselectedTableProperties.add(tableproperties.ID)
          console.log("is unselected " + unselectedTableProperties.has(tableproperties.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let tableproperties = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as TablePropertiesDB
      if (unselectedTableProperties.has(tableproperties.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<TablePropertiesDB>) = new Array<TablePropertiesDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          tableproperties => {
            if (!this.initialSelection.includes(tableproperties)) {
              // console.log("tableproperties " + tableproperties.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + tableproperties.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = tableproperties.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = tableproperties.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("tableproperties " + tableproperties.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<TablePropertiesDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}