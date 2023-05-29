// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { ParagraphPropertiesDB } from '../paragraphproperties-db'
import { ParagraphPropertiesService } from '../paragraphproperties.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// ParagraphPropertiesDetailComponent is initilizaed from different routes
// ParagraphPropertiesDetailComponentState detail different cases 
enum ParagraphPropertiesDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
}

@Component({
	selector: 'app-paragraphproperties-detail',
	templateUrl: './paragraphproperties-detail.component.html',
	styleUrls: ['./paragraphproperties-detail.component.css'],
})
export class ParagraphPropertiesDetailComponent implements OnInit {

	// insertion point for declarations

	// the ParagraphPropertiesDB of interest
	paragraphproperties: ParagraphPropertiesDB = new ParagraphPropertiesDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: ParagraphPropertiesDetailComponentState = ParagraphPropertiesDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private paragraphpropertiesService: ParagraphPropertiesService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private activatedRoute: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		this.activatedRoute.params.subscribe(params => {
			this.onChangedActivatedRoute()
		});
	}
	onChangedActivatedRoute(): void {

		// compute state
		this.id = +this.activatedRoute.snapshot.paramMap.get('id')!;
		this.originStruct = this.activatedRoute.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.activatedRoute.snapshot.paramMap.get('originStructFieldName')!;

		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		const association = this.activatedRoute.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = ParagraphPropertiesDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = ParagraphPropertiesDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getParagraphProperties()

		// observable for changes in structs
		this.paragraphpropertiesService.ParagraphPropertiesServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getParagraphProperties()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getParagraphProperties(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case ParagraphPropertiesDetailComponentState.CREATE_INSTANCE:
						this.paragraphproperties = new (ParagraphPropertiesDB)
						break;
					case ParagraphPropertiesDetailComponentState.UPDATE_INSTANCE:
						let paragraphproperties = frontRepo.ParagraphPropertiess.get(this.id)
						console.assert(paragraphproperties != undefined, "missing paragraphproperties with id:" + this.id)
						this.paragraphproperties = paragraphproperties!
						break;
					// insertion point for init of association field
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.paragraphproperties.ParagraphStyleID == undefined) {
			this.paragraphproperties.ParagraphStyleID = new NullInt64
		}
		if (this.paragraphproperties.ParagraphStyle != undefined) {
			this.paragraphproperties.ParagraphStyleID.Int64 = this.paragraphproperties.ParagraphStyle.ID
			this.paragraphproperties.ParagraphStyleID.Valid = true
		} else {
			this.paragraphproperties.ParagraphStyleID.Int64 = 0
			this.paragraphproperties.ParagraphStyleID.Valid = true
		}
		if (this.paragraphproperties.NodeID == undefined) {
			this.paragraphproperties.NodeID = new NullInt64
		}
		if (this.paragraphproperties.Node != undefined) {
			this.paragraphproperties.NodeID.Int64 = this.paragraphproperties.Node.ID
			this.paragraphproperties.NodeID.Valid = true
		} else {
			this.paragraphproperties.NodeID.Int64 = 0
			this.paragraphproperties.NodeID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers

		switch (this.state) {
			case ParagraphPropertiesDetailComponentState.UPDATE_INSTANCE:
				this.paragraphpropertiesService.updateParagraphProperties(this.paragraphproperties, this.GONG__StackPath)
					.subscribe(paragraphproperties => {
						this.paragraphpropertiesService.ParagraphPropertiesServiceChanged.next("update")
					});
				break;
			default:
				this.paragraphpropertiesService.postParagraphProperties(this.paragraphproperties, this.GONG__StackPath).subscribe(paragraphproperties => {
					this.paragraphpropertiesService.ParagraphPropertiesServiceChanged.next("post")
					this.paragraphproperties = new (ParagraphPropertiesDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.paragraphproperties.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.paragraphproperties.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "ParagraphProperties"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.paragraphproperties.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
			GONG__StackPath: this.GONG__StackPath,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.paragraphproperties.Name == "") {
			this.paragraphproperties.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
