// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { TableRowDB } from '../tablerow-db'
import { TableRowService } from '../tablerow.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-tablerow-sorting',
  templateUrl: './tablerow-sorting.component.html',
  styleUrls: ['./tablerow-sorting.component.css']
})
export class TableRowSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of TableRow instances that are in the association
  associatedTableRows = new Array<TableRowDB>();

  constructor(
    private tablerowService: TableRowService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of tablerow instances
    public dialogRef: MatDialogRef<TableRowSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getTableRows()
  }

  getTableRows(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let tablerow of this.frontRepo.TableRows_array) {
          let ID = this.dialogData.ID
          let revPointerID = tablerow[this.dialogData.ReversePointer as keyof TableRowDB] as unknown as NullInt64
          let revPointerID_Index = tablerow[this.dialogData.ReversePointer + "_Index" as keyof TableRowDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedTableRows.push(tablerow)
          }
        }

        // sort associated tablerow according to order
        this.associatedTableRows.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedTableRows, event.previousIndex, event.currentIndex);

    // set the order of TableRow instances
    let index = 0

    for (let tablerow of this.associatedTableRows) {
      let revPointerID_Index = tablerow[this.dialogData.ReversePointer + "_Index" as keyof TableRowDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedTableRows.forEach(
      tablerow => {
        this.tablerowService.updateTableRow(tablerow, this.dialogData.GONG__StackPath)
          .subscribe(tablerow => {
            this.tablerowService.TableRowServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}