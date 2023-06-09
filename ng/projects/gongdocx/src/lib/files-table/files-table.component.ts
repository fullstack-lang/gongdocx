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
import { FileDB } from '../file-db'
import { FileService } from '../file.service'

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
  selector: 'app-filestable',
  templateUrl: './files-table.component.html',
  styleUrls: ['./files-table.component.css'],
})
export class FilesTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of File instances
  selection: SelectionModel<FileDB> = new (SelectionModel)
  initialSelection = new Array<FileDB>()

  // the data source for the table
  files: FileDB[] = []
  matTableDataSource: MatTableDataSource<FileDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.files
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
    this.matTableDataSource.sortingDataAccessor = (fileDB: FileDB, property: string) => {
      switch (property) {
        case 'ID':
          return fileDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return fileDB.Name;

        case 'Docx_Files':
          if (this.frontRepo.Docxs.get(fileDB.Docx_FilesDBID.Int64) != undefined) {
            return this.frontRepo.Docxs.get(fileDB.Docx_FilesDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (fileDB: FileDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the fileDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += fileDB.Name.toLowerCase()
      if (fileDB.Docx_FilesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Docxs.get(fileDB.Docx_FilesDBID.Int64)!.Name.toLowerCase()
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
    private fileService: FileService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of file instances
    public dialogRef: MatDialogRef<FilesTableComponent>,
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
    this.fileService.FileServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getFiles()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Docx_Files",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Docx_Files",
      ]
      this.selection = new SelectionModel<FileDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getFiles()

    this.matTableDataSource = new MatTableDataSource(this.files)
  }

  getFiles(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.files = this.frontRepo.Files_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let file of this.files) {
            let ID = this.dialogData.ID
            let revPointer = file[this.dialogData.ReversePointer as keyof FileDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(file)
            }
            this.selection = new SelectionModel<FileDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, FileDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to FileDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as FileDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let file = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as FileDB
              this.initialSelection.push(file)
            }
          }

          this.selection = new SelectionModel<FileDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.files
      }
    )
  }

  // newFile initiate a new file
  // create a new File objet
  newFile() {
  }

  deleteFile(fileID: number, file: FileDB) {
    // list of files is truncated of file before the delete
    this.files = this.files.filter(h => h !== file);

    this.fileService.deleteFile(fileID, this.GONG__StackPath).subscribe(
      file => {
        this.fileService.FileServiceChanged.next("delete")
      }
    );
  }

  editFile(fileID: number, file: FileDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(fileID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "file" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, fileID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.files.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.files.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<FileDB>()

      // reset all initial selection of file that belong to file
      for (let file of this.initialSelection) {
        let index = file[this.dialogData.ReversePointer as keyof FileDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(file)

      }

      // from selection, set file that belong to file
      for (let file of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = file[this.dialogData.ReversePointer as keyof FileDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(file)
      }


      // update all file (only update selection & initial selection)
      for (let file of toUpdate) {
        this.fileService.updateFile(file, this.GONG__StackPath)
          .subscribe(file => {
            this.fileService.FileServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, FileDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedFile = new Set<number>()
      for (let file of this.initialSelection) {
        if (this.selection.selected.includes(file)) {
          // console.log("file " + file.Name + " is still selected")
        } else {
          console.log("file " + file.Name + " has been unselected")
          unselectedFile.add(file.ID)
          console.log("is unselected " + unselectedFile.has(file.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let file = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as FileDB
      if (unselectedFile.has(file.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<FileDB>) = new Array<FileDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          file => {
            if (!this.initialSelection.includes(file)) {
              // console.log("file " + file.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + file.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = file.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = file.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("file " + file.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<FileDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
