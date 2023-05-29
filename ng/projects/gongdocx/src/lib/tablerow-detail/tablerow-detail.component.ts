// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { TableRowDB } from '../tablerow-db'
import { TableRowService } from '../tablerow.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { TableDB } from '../table-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// TableRowDetailComponent is initilizaed from different routes
// TableRowDetailComponentState detail different cases 
enum TableRowDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Table_TableRows_SET,
}

@Component({
	selector: 'app-tablerow-detail',
	templateUrl: './tablerow-detail.component.html',
	styleUrls: ['./tablerow-detail.component.css'],
})
export class TableRowDetailComponent implements OnInit {

	// insertion point for declarations

	// the TableRowDB of interest
	tablerow: TableRowDB = new TableRowDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: TableRowDetailComponentState = TableRowDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private tablerowService: TableRowService,
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
			this.state = TableRowDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = TableRowDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "TableRows":
						// console.log("TableRow" + " is instanciated with back pointer to instance " + this.id + " Table association TableRows")
						this.state = TableRowDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Table_TableRows_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getTableRow()

		// observable for changes in structs
		this.tablerowService.TableRowServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getTableRow()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getTableRow(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case TableRowDetailComponentState.CREATE_INSTANCE:
						this.tablerow = new (TableRowDB)
						break;
					case TableRowDetailComponentState.UPDATE_INSTANCE:
						let tablerow = frontRepo.TableRows.get(this.id)
						console.assert(tablerow != undefined, "missing tablerow with id:" + this.id)
						this.tablerow = tablerow!
						break;
					// insertion point for init of association field
					case TableRowDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Table_TableRows_SET:
						this.tablerow = new (TableRowDB)
						this.tablerow.Table_TableRows_reverse = frontRepo.Tables.get(this.id)!
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
		if (this.tablerow.NodeID == undefined) {
			this.tablerow.NodeID = new NullInt64
		}
		if (this.tablerow.Node != undefined) {
			this.tablerow.NodeID.Int64 = this.tablerow.Node.ID
			this.tablerow.NodeID.Valid = true
		} else {
			this.tablerow.NodeID.Int64 = 0
			this.tablerow.NodeID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.tablerow.Table_TableRows_reverse != undefined) {
			if (this.tablerow.Table_TableRowsDBID == undefined) {
				this.tablerow.Table_TableRowsDBID = new NullInt64
			}
			this.tablerow.Table_TableRowsDBID.Int64 = this.tablerow.Table_TableRows_reverse.ID
			this.tablerow.Table_TableRowsDBID.Valid = true
			if (this.tablerow.Table_TableRowsDBID_Index == undefined) {
				this.tablerow.Table_TableRowsDBID_Index = new NullInt64
			}
			this.tablerow.Table_TableRowsDBID_Index.Valid = true
			this.tablerow.Table_TableRows_reverse = new TableDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case TableRowDetailComponentState.UPDATE_INSTANCE:
				this.tablerowService.updateTableRow(this.tablerow, this.GONG__StackPath)
					.subscribe(tablerow => {
						this.tablerowService.TableRowServiceChanged.next("update")
					});
				break;
			default:
				this.tablerowService.postTableRow(this.tablerow, this.GONG__StackPath).subscribe(tablerow => {
					this.tablerowService.TableRowServiceChanged.next("post")
					this.tablerow = new (TableRowDB) // reset fields
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

			dialogData.ID = this.tablerow.ID!
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
			dialogData.ID = this.tablerow.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "TableRow"
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
			ID: this.tablerow.ID,
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
		if (this.tablerow.Name == "") {
			this.tablerow.Name = event.value.Name
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
