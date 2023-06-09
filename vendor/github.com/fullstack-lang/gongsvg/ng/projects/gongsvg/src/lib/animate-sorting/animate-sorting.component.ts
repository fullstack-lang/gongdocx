// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { AnimateDB } from '../animate-db'
import { AnimateService } from '../animate.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-animate-sorting',
  templateUrl: './animate-sorting.component.html',
  styleUrls: ['./animate-sorting.component.css']
})
export class AnimateSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Animate instances that are in the association
  associatedAnimates = new Array<AnimateDB>();

  constructor(
    private animateService: AnimateService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of animate instances
    public dialogRef: MatDialogRef<AnimateSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getAnimates()
  }

  getAnimates(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let animate of this.frontRepo.Animates_array) {
          let ID = this.dialogData.ID
          let revPointerID = animate[this.dialogData.ReversePointer as keyof AnimateDB] as unknown as NullInt64
          let revPointerID_Index = animate[this.dialogData.ReversePointer + "_Index" as keyof AnimateDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedAnimates.push(animate)
          }
        }

        // sort associated animate according to order
        this.associatedAnimates.sort((t1, t2) => {
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
    moveItemInArray(this.associatedAnimates, event.previousIndex, event.currentIndex);

    // set the order of Animate instances
    let index = 0

    for (let animate of this.associatedAnimates) {
      let revPointerID_Index = animate[this.dialogData.ReversePointer + "_Index" as keyof AnimateDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedAnimates.forEach(
      animate => {
        this.animateService.updateAnimate(animate, this.dialogData.GONG__StackPath)
          .subscribe(animate => {
            this.animateService.AnimateServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}
