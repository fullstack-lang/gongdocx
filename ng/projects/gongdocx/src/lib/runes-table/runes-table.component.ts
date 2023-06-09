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
import { RuneDB } from '../rune-db'
import { RuneService } from '../rune.service'

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
  selector: 'app-runestable',
  templateUrl: './runes-table.component.html',
  styleUrls: ['./runes-table.component.css'],
})
export class RunesTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Rune instances
  selection: SelectionModel<RuneDB> = new (SelectionModel)
  initialSelection = new Array<RuneDB>()

  // the data source for the table
  runes: RuneDB[] = []
  matTableDataSource: MatTableDataSource<RuneDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.runes
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
    this.matTableDataSource.sortingDataAccessor = (runeDB: RuneDB, property: string) => {
      switch (property) {
        case 'ID':
          return runeDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return runeDB.Name;

        case 'Content':
          return runeDB.Content;

        case 'Node':
          return (runeDB.Node ? runeDB.Node.Name : '');

        case 'Text':
          return (runeDB.Text ? runeDB.Text.Name : '');

        case 'RuneProperties':
          return (runeDB.RuneProperties ? runeDB.RuneProperties.Name : '');

        case 'EnclosingParagraph':
          return (runeDB.EnclosingParagraph ? runeDB.EnclosingParagraph.Name : '');

        case 'Paragraph_Runes':
          if (this.frontRepo.Paragraphs.get(runeDB.Paragraph_RunesDBID.Int64) != undefined) {
            return this.frontRepo.Paragraphs.get(runeDB.Paragraph_RunesDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (runeDB: RuneDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the runeDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += runeDB.Name.toLowerCase()
      mergedContent += runeDB.Content.toLowerCase()
      if (runeDB.Node) {
        mergedContent += runeDB.Node.Name.toLowerCase()
      }
      if (runeDB.Text) {
        mergedContent += runeDB.Text.Name.toLowerCase()
      }
      if (runeDB.RuneProperties) {
        mergedContent += runeDB.RuneProperties.Name.toLowerCase()
      }
      if (runeDB.EnclosingParagraph) {
        mergedContent += runeDB.EnclosingParagraph.Name.toLowerCase()
      }
      if (runeDB.Paragraph_RunesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Paragraphs.get(runeDB.Paragraph_RunesDBID.Int64)!.Name.toLowerCase()
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
    private runeService: RuneService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of rune instances
    public dialogRef: MatDialogRef<RunesTableComponent>,
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
    this.runeService.RuneServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getRunes()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Content",
        "Node",
        "Text",
        "RuneProperties",
        "EnclosingParagraph",
        "Paragraph_Runes",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Content",
        "Node",
        "Text",
        "RuneProperties",
        "EnclosingParagraph",
        "Paragraph_Runes",
      ]
      this.selection = new SelectionModel<RuneDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getRunes()

    this.matTableDataSource = new MatTableDataSource(this.runes)
  }

  getRunes(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.runes = this.frontRepo.Runes_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let rune of this.runes) {
            let ID = this.dialogData.ID
            let revPointer = rune[this.dialogData.ReversePointer as keyof RuneDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(rune)
            }
            this.selection = new SelectionModel<RuneDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, RuneDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to RuneDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as RuneDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let rune = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as RuneDB
              this.initialSelection.push(rune)
            }
          }

          this.selection = new SelectionModel<RuneDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.runes
      }
    )
  }

  // newRune initiate a new rune
  // create a new Rune objet
  newRune() {
  }

  deleteRune(runeID: number, rune: RuneDB) {
    // list of runes is truncated of rune before the delete
    this.runes = this.runes.filter(h => h !== rune);

    this.runeService.deleteRune(runeID, this.GONG__StackPath).subscribe(
      rune => {
        this.runeService.RuneServiceChanged.next("delete")
      }
    );
  }

  editRune(runeID: number, rune: RuneDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(runeID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "rune" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, runeID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.runes.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.runes.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<RuneDB>()

      // reset all initial selection of rune that belong to rune
      for (let rune of this.initialSelection) {
        let index = rune[this.dialogData.ReversePointer as keyof RuneDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(rune)

      }

      // from selection, set rune that belong to rune
      for (let rune of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = rune[this.dialogData.ReversePointer as keyof RuneDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(rune)
      }


      // update all rune (only update selection & initial selection)
      for (let rune of toUpdate) {
        this.runeService.updateRune(rune, this.GONG__StackPath)
          .subscribe(rune => {
            this.runeService.RuneServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, RuneDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedRune = new Set<number>()
      for (let rune of this.initialSelection) {
        if (this.selection.selected.includes(rune)) {
          // console.log("rune " + rune.Name + " is still selected")
        } else {
          console.log("rune " + rune.Name + " has been unselected")
          unselectedRune.add(rune.ID)
          console.log("is unselected " + unselectedRune.has(rune.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let rune = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as RuneDB
      if (unselectedRune.has(rune.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<RuneDB>) = new Array<RuneDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          rune => {
            if (!this.initialSelection.includes(rune)) {
              // console.log("rune " + rune.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + rune.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = rune.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = rune.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("rune " + rune.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<RuneDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
