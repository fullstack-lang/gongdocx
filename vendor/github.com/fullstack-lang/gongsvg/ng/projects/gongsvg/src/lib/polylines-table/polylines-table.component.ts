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
import { PolylineDB } from '../polyline-db'
import { PolylineService } from '../polyline.service'

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
  selector: 'app-polylinestable',
  templateUrl: './polylines-table.component.html',
  styleUrls: ['./polylines-table.component.css'],
})
export class PolylinesTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Polyline instances
  selection: SelectionModel<PolylineDB> = new (SelectionModel)
  initialSelection = new Array<PolylineDB>()

  // the data source for the table
  polylines: PolylineDB[] = []
  matTableDataSource: MatTableDataSource<PolylineDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.polylines
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
    this.matTableDataSource.sortingDataAccessor = (polylineDB: PolylineDB, property: string) => {
      switch (property) {
        case 'ID':
          return polylineDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return polylineDB.Name;

        case 'Points':
          return polylineDB.Points;

        case 'Color':
          return polylineDB.Color;

        case 'FillOpacity':
          return polylineDB.FillOpacity;

        case 'Stroke':
          return polylineDB.Stroke;

        case 'StrokeWidth':
          return polylineDB.StrokeWidth;

        case 'StrokeDashArray':
          return polylineDB.StrokeDashArray;

        case 'StrokeDashArrayWhenSelected':
          return polylineDB.StrokeDashArrayWhenSelected;

        case 'Transform':
          return polylineDB.Transform;

        case 'Layer_Polylines':
          if (this.frontRepo.Layers.get(polylineDB.Layer_PolylinesDBID.Int64) != undefined) {
            return this.frontRepo.Layers.get(polylineDB.Layer_PolylinesDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (polylineDB: PolylineDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the polylineDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += polylineDB.Name.toLowerCase()
      mergedContent += polylineDB.Points.toLowerCase()
      mergedContent += polylineDB.Color.toLowerCase()
      mergedContent += polylineDB.FillOpacity.toString()
      mergedContent += polylineDB.Stroke.toLowerCase()
      mergedContent += polylineDB.StrokeWidth.toString()
      mergedContent += polylineDB.StrokeDashArray.toLowerCase()
      mergedContent += polylineDB.StrokeDashArrayWhenSelected.toLowerCase()
      mergedContent += polylineDB.Transform.toLowerCase()
      if (polylineDB.Layer_PolylinesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Layers.get(polylineDB.Layer_PolylinesDBID.Int64)!.Name.toLowerCase()
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
    private polylineService: PolylineService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of polyline instances
    public dialogRef: MatDialogRef<PolylinesTableComponent>,
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
    this.polylineService.PolylineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getPolylines()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Points",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "StrokeDashArrayWhenSelected",
        "Transform",
        "Layer_Polylines",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Points",
        "Color",
        "FillOpacity",
        "Stroke",
        "StrokeWidth",
        "StrokeDashArray",
        "StrokeDashArrayWhenSelected",
        "Transform",
        "Layer_Polylines",
      ]
      this.selection = new SelectionModel<PolylineDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getPolylines()

    this.matTableDataSource = new MatTableDataSource(this.polylines)
  }

  getPolylines(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.polylines = this.frontRepo.Polylines_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let polyline of this.polylines) {
            let ID = this.dialogData.ID
            let revPointer = polyline[this.dialogData.ReversePointer as keyof PolylineDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(polyline)
            }
            this.selection = new SelectionModel<PolylineDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, PolylineDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to PolylineDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as PolylineDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let polyline = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as PolylineDB
              this.initialSelection.push(polyline)
            }
          }

          this.selection = new SelectionModel<PolylineDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.polylines
      }
    )
  }

  // newPolyline initiate a new polyline
  // create a new Polyline objet
  newPolyline() {
  }

  deletePolyline(polylineID: number, polyline: PolylineDB) {
    // list of polylines is truncated of polyline before the delete
    this.polylines = this.polylines.filter(h => h !== polyline);

    this.polylineService.deletePolyline(polylineID, this.GONG__StackPath).subscribe(
      polyline => {
        this.polylineService.PolylineServiceChanged.next("delete")
      }
    );
  }

  editPolyline(polylineID: number, polyline: PolylineDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(polylineID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "polyline" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, polylineID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.polylines.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.polylines.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<PolylineDB>()

      // reset all initial selection of polyline that belong to polyline
      for (let polyline of this.initialSelection) {
        let index = polyline[this.dialogData.ReversePointer as keyof PolylineDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(polyline)

      }

      // from selection, set polyline that belong to polyline
      for (let polyline of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = polyline[this.dialogData.ReversePointer as keyof PolylineDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(polyline)
      }


      // update all polyline (only update selection & initial selection)
      for (let polyline of toUpdate) {
        this.polylineService.updatePolyline(polyline, this.GONG__StackPath)
          .subscribe(polyline => {
            this.polylineService.PolylineServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, PolylineDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedPolyline = new Set<number>()
      for (let polyline of this.initialSelection) {
        if (this.selection.selected.includes(polyline)) {
          // console.log("polyline " + polyline.Name + " is still selected")
        } else {
          console.log("polyline " + polyline.Name + " has been unselected")
          unselectedPolyline.add(polyline.ID)
          console.log("is unselected " + unselectedPolyline.has(polyline.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let polyline = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as PolylineDB
      if (unselectedPolyline.has(polyline.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<PolylineDB>) = new Array<PolylineDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          polyline => {
            if (!this.initialSelection.includes(polyline)) {
              // console.log("polyline " + polyline.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + polyline.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = polyline.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = polyline.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("polyline " + polyline.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<PolylineDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
