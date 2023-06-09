// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { ParagraphDB } from '../paragraph-db'
import { ParagraphService } from '../paragraph.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { BodyDB } from '../body-db'
import { TableColumnDB } from '../tablecolumn-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// ParagraphDetailComponent is initilizaed from different routes
// ParagraphDetailComponentState detail different cases 
enum ParagraphDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Body_Paragraphs_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_TableColumn_Paragraphs_SET,
}

@Component({
	selector: 'app-paragraph-detail',
	templateUrl: './paragraph-detail.component.html',
	styleUrls: ['./paragraph-detail.component.css'],
})
export class ParagraphDetailComponent implements OnInit {

	// insertion point for declarations

	// the ParagraphDB of interest
	paragraph: ParagraphDB = new ParagraphDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: ParagraphDetailComponentState = ParagraphDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private paragraphService: ParagraphService,
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
			this.state = ParagraphDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = ParagraphDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Paragraphs":
						// console.log("Paragraph" + " is instanciated with back pointer to instance " + this.id + " Body association Paragraphs")
						this.state = ParagraphDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Body_Paragraphs_SET
						break;
					case "Paragraphs":
						// console.log("Paragraph" + " is instanciated with back pointer to instance " + this.id + " TableColumn association Paragraphs")
						this.state = ParagraphDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_TableColumn_Paragraphs_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getParagraph()

		// observable for changes in structs
		this.paragraphService.ParagraphServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getParagraph()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getParagraph(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case ParagraphDetailComponentState.CREATE_INSTANCE:
						this.paragraph = new (ParagraphDB)
						break;
					case ParagraphDetailComponentState.UPDATE_INSTANCE:
						let paragraph = frontRepo.Paragraphs.get(this.id)
						console.assert(paragraph != undefined, "missing paragraph with id:" + this.id)
						this.paragraph = paragraph!
						break;
					// insertion point for init of association field
					case ParagraphDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Body_Paragraphs_SET:
						this.paragraph = new (ParagraphDB)
						this.paragraph.Body_Paragraphs_reverse = frontRepo.Bodys.get(this.id)!
						break;
					case ParagraphDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_TableColumn_Paragraphs_SET:
						this.paragraph = new (ParagraphDB)
						this.paragraph.TableColumn_Paragraphs_reverse = frontRepo.TableColumns.get(this.id)!
						break;
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
		if (this.paragraph.NodeID == undefined) {
			this.paragraph.NodeID = new NullInt64
		}
		if (this.paragraph.Node != undefined) {
			this.paragraph.NodeID.Int64 = this.paragraph.Node.ID
			this.paragraph.NodeID.Valid = true
		} else {
			this.paragraph.NodeID.Int64 = 0
			this.paragraph.NodeID.Valid = true
		}
		if (this.paragraph.ParagraphPropertiesID == undefined) {
			this.paragraph.ParagraphPropertiesID = new NullInt64
		}
		if (this.paragraph.ParagraphProperties != undefined) {
			this.paragraph.ParagraphPropertiesID.Int64 = this.paragraph.ParagraphProperties.ID
			this.paragraph.ParagraphPropertiesID.Valid = true
		} else {
			this.paragraph.ParagraphPropertiesID.Int64 = 0
			this.paragraph.ParagraphPropertiesID.Valid = true
		}
		if (this.paragraph.NextID == undefined) {
			this.paragraph.NextID = new NullInt64
		}
		if (this.paragraph.Next != undefined) {
			this.paragraph.NextID.Int64 = this.paragraph.Next.ID
			this.paragraph.NextID.Valid = true
		} else {
			this.paragraph.NextID.Int64 = 0
			this.paragraph.NextID.Valid = true
		}
		if (this.paragraph.PreviousID == undefined) {
			this.paragraph.PreviousID = new NullInt64
		}
		if (this.paragraph.Previous != undefined) {
			this.paragraph.PreviousID.Int64 = this.paragraph.Previous.ID
			this.paragraph.PreviousID.Valid = true
		} else {
			this.paragraph.PreviousID.Int64 = 0
			this.paragraph.PreviousID.Valid = true
		}
		if (this.paragraph.EnclosingBodyID == undefined) {
			this.paragraph.EnclosingBodyID = new NullInt64
		}
		if (this.paragraph.EnclosingBody != undefined) {
			this.paragraph.EnclosingBodyID.Int64 = this.paragraph.EnclosingBody.ID
			this.paragraph.EnclosingBodyID.Valid = true
		} else {
			this.paragraph.EnclosingBodyID.Int64 = 0
			this.paragraph.EnclosingBodyID.Valid = true
		}
		if (this.paragraph.EnclosingTableColumnID == undefined) {
			this.paragraph.EnclosingTableColumnID = new NullInt64
		}
		if (this.paragraph.EnclosingTableColumn != undefined) {
			this.paragraph.EnclosingTableColumnID.Int64 = this.paragraph.EnclosingTableColumn.ID
			this.paragraph.EnclosingTableColumnID.Valid = true
		} else {
			this.paragraph.EnclosingTableColumnID.Int64 = 0
			this.paragraph.EnclosingTableColumnID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.paragraph.Body_Paragraphs_reverse != undefined) {
			if (this.paragraph.Body_ParagraphsDBID == undefined) {
				this.paragraph.Body_ParagraphsDBID = new NullInt64
			}
			this.paragraph.Body_ParagraphsDBID.Int64 = this.paragraph.Body_Paragraphs_reverse.ID
			this.paragraph.Body_ParagraphsDBID.Valid = true
			if (this.paragraph.Body_ParagraphsDBID_Index == undefined) {
				this.paragraph.Body_ParagraphsDBID_Index = new NullInt64
			}
			this.paragraph.Body_ParagraphsDBID_Index.Valid = true
			this.paragraph.Body_Paragraphs_reverse = new BodyDB // very important, otherwise, circular JSON
		}
		if (this.paragraph.TableColumn_Paragraphs_reverse != undefined) {
			if (this.paragraph.TableColumn_ParagraphsDBID == undefined) {
				this.paragraph.TableColumn_ParagraphsDBID = new NullInt64
			}
			this.paragraph.TableColumn_ParagraphsDBID.Int64 = this.paragraph.TableColumn_Paragraphs_reverse.ID
			this.paragraph.TableColumn_ParagraphsDBID.Valid = true
			if (this.paragraph.TableColumn_ParagraphsDBID_Index == undefined) {
				this.paragraph.TableColumn_ParagraphsDBID_Index = new NullInt64
			}
			this.paragraph.TableColumn_ParagraphsDBID_Index.Valid = true
			this.paragraph.TableColumn_Paragraphs_reverse = new TableColumnDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case ParagraphDetailComponentState.UPDATE_INSTANCE:
				this.paragraphService.updateParagraph(this.paragraph, this.GONG__StackPath)
					.subscribe(paragraph => {
						this.paragraphService.ParagraphServiceChanged.next("update")
					});
				break;
			default:
				this.paragraphService.postParagraph(this.paragraph, this.GONG__StackPath).subscribe(paragraph => {
					this.paragraphService.ParagraphServiceChanged.next("post")
					this.paragraph = new (ParagraphDB) // reset fields
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

			dialogData.ID = this.paragraph.ID!
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
			dialogData.ID = this.paragraph.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "Paragraph"
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
			ID: this.paragraph.ID,
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
		if (this.paragraph.Name == "") {
			this.paragraph.Name = event.value.Name
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
