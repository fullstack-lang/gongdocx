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
import { ParagraphPropertiesDB } from '../paragraphproperties-db'
import { ParagraphPropertiesService } from '../paragraphproperties.service'

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
  selector: 'app-paragraphpropertiesstable',
  templateUrl: './paragraphpropertiess-table.component.html',
  styleUrls: ['./paragraphpropertiess-table.component.css'],
})
export class ParagraphPropertiessTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of ParagraphProperties instances
  selection: SelectionModel<ParagraphPropertiesDB> = new (SelectionModel)
  initialSelection = new Array<ParagraphPropertiesDB>()

  // the data source for the table
  paragraphpropertiess: ParagraphPropertiesDB[] = []
  matTableDataSource: MatTableDataSource<ParagraphPropertiesDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.paragraphpropertiess
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
    this.matTableDataSource.sortingDataAccessor = (paragraphpropertiesDB: ParagraphPropertiesDB, property: string) => {
      switch (property) {
        case 'ID':
          return paragraphpropertiesDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return paragraphpropertiesDB.Name;

        case 'Node':
          return (paragraphpropertiesDB.Node ? paragraphpropertiesDB.Node.Name : '');

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (paragraphpropertiesDB: ParagraphPropertiesDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the paragraphpropertiesDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += paragraphpropertiesDB.Name.toLowerCase()
      if (paragraphpropertiesDB.Node) {
        mergedContent += paragraphpropertiesDB.Node.Name.toLowerCase()
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
    private paragraphpropertiesService: ParagraphPropertiesService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of paragraphproperties instances
    public dialogRef: MatDialogRef<ParagraphPropertiessTableComponent>,
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
    this.paragraphpropertiesService.ParagraphPropertiesServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getParagraphPropertiess()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Node",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Node",
      ]
      this.selection = new SelectionModel<ParagraphPropertiesDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getParagraphPropertiess()

    this.matTableDataSource = new MatTableDataSource(this.paragraphpropertiess)
  }

  getParagraphPropertiess(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.paragraphpropertiess = this.frontRepo.ParagraphPropertiess_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let paragraphproperties of this.paragraphpropertiess) {
            let ID = this.dialogData.ID
            let revPointer = paragraphproperties[this.dialogData.ReversePointer as keyof ParagraphPropertiesDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(paragraphproperties)
            }
            this.selection = new SelectionModel<ParagraphPropertiesDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ParagraphPropertiesDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to ParagraphPropertiesDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as ParagraphPropertiesDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let paragraphproperties = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ParagraphPropertiesDB
              this.initialSelection.push(paragraphproperties)
            }
          }

          this.selection = new SelectionModel<ParagraphPropertiesDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.paragraphpropertiess
      }
    )
  }

  // newParagraphProperties initiate a new paragraphproperties
  // create a new ParagraphProperties objet
  newParagraphProperties() {
  }

  deleteParagraphProperties(paragraphpropertiesID: number, paragraphproperties: ParagraphPropertiesDB) {
    // list of paragraphpropertiess is truncated of paragraphproperties before the delete
    this.paragraphpropertiess = this.paragraphpropertiess.filter(h => h !== paragraphproperties);

    this.paragraphpropertiesService.deleteParagraphProperties(paragraphpropertiesID, this.GONG__StackPath).subscribe(
      paragraphproperties => {
        this.paragraphpropertiesService.ParagraphPropertiesServiceChanged.next("delete")
      }
    );
  }

  editParagraphProperties(paragraphpropertiesID: number, paragraphproperties: ParagraphPropertiesDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(paragraphpropertiesID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "paragraphproperties" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, paragraphpropertiesID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.paragraphpropertiess.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.paragraphpropertiess.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<ParagraphPropertiesDB>()

      // reset all initial selection of paragraphproperties that belong to paragraphproperties
      for (let paragraphproperties of this.initialSelection) {
        let index = paragraphproperties[this.dialogData.ReversePointer as keyof ParagraphPropertiesDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(paragraphproperties)

      }

      // from selection, set paragraphproperties that belong to paragraphproperties
      for (let paragraphproperties of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = paragraphproperties[this.dialogData.ReversePointer as keyof ParagraphPropertiesDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(paragraphproperties)
      }


      // update all paragraphproperties (only update selection & initial selection)
      for (let paragraphproperties of toUpdate) {
        this.paragraphpropertiesService.updateParagraphProperties(paragraphproperties, this.GONG__StackPath)
          .subscribe(paragraphproperties => {
            this.paragraphpropertiesService.ParagraphPropertiesServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ParagraphPropertiesDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedParagraphProperties = new Set<number>()
      for (let paragraphproperties of this.initialSelection) {
        if (this.selection.selected.includes(paragraphproperties)) {
          // console.log("paragraphproperties " + paragraphproperties.Name + " is still selected")
        } else {
          console.log("paragraphproperties " + paragraphproperties.Name + " has been unselected")
          unselectedParagraphProperties.add(paragraphproperties.ID)
          console.log("is unselected " + unselectedParagraphProperties.has(paragraphproperties.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let paragraphproperties = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ParagraphPropertiesDB
      if (unselectedParagraphProperties.has(paragraphproperties.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<ParagraphPropertiesDB>) = new Array<ParagraphPropertiesDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          paragraphproperties => {
            if (!this.initialSelection.includes(paragraphproperties)) {
              // console.log("paragraphproperties " + paragraphproperties.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + paragraphproperties.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = paragraphproperties.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = paragraphproperties.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("paragraphproperties " + paragraphproperties.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<ParagraphPropertiesDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}