// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { TableColumnDB } from '../tablecolumn-db'
import { TableColumnService } from '../tablecolumn.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { TableRowDB } from '../tablerow-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// TableColumnDetailComponent is initilizaed from different routes
// TableColumnDetailComponentState detail different cases 
enum TableColumnDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_TableRow_TableColumns_SET,
}

@Component({
	selector: 'app-tablecolumn-detail',
	templateUrl: './tablecolumn-detail.component.html',
	styleUrls: ['./tablecolumn-detail.component.css'],
})
export class TableColumnDetailComponent implements OnInit {

	// insertion point for declarations

	// the TableColumnDB of interest
	tablecolumn: TableColumnDB = new TableColumnDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: TableColumnDetailComponentState = TableColumnDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private tablecolumnService: TableColumnService,
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
			this.state = TableColumnDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = TableColumnDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "TableColumns":
						// console.log("TableColumn" + " is instanciated with back pointer to instance " + this.id + " TableRow association TableColumns")
						this.state = TableColumnDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_TableRow_TableColumns_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getTableColumn()

		// observable for changes in structs
		this.tablecolumnService.TableColumnServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getTableColumn()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getTableColumn(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case TableColumnDetailComponentState.CREATE_INSTANCE:
						this.tablecolumn = new (TableColumnDB)
						break;
					case TableColumnDetailComponentState.UPDATE_INSTANCE:
						let tablecolumn = frontRepo.TableColumns.get(this.id)
						console.assert(tablecolumn != undefined, "missing tablecolumn with id:" + this.id)
						this.tablecolumn = tablecolumn!
						break;
					// insertion point for init of association field
					case TableColumnDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_TableRow_TableColumns_SET:
						this.tablecolumn = new (TableColumnDB)
						this.tablecolumn.TableRow_TableColumns_reverse = frontRepo.TableRows.get(this.id)!
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
		if (this.tablecolumn.NodeID == undefined) {
			this.tablecolumn.NodeID = new NullInt64
		}
		if (this.tablecolumn.Node != undefined) {
			this.tablecolumn.NodeID.Int64 = this.tablecolumn.Node.ID
			this.tablecolumn.NodeID.Valid = true
		} else {
			this.tablecolumn.NodeID.Int64 = 0
			this.tablecolumn.NodeID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.tablecolumn.TableRow_TableColumns_reverse != undefined) {
			if (this.tablecolumn.TableRow_TableColumnsDBID == undefined) {
				this.tablecolumn.TableRow_TableColumnsDBID = new NullInt64
			}
			this.tablecolumn.TableRow_TableColumnsDBID.Int64 = this.tablecolumn.TableRow_TableColumns_reverse.ID
			this.tablecolumn.TableRow_TableColumnsDBID.Valid = true
			if (this.tablecolumn.TableRow_TableColumnsDBID_Index == undefined) {
				this.tablecolumn.TableRow_TableColumnsDBID_Index = new NullInt64
			}
			this.tablecolumn.TableRow_TableColumnsDBID_Index.Valid = true
			this.tablecolumn.TableRow_TableColumns_reverse = new TableRowDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case TableColumnDetailComponentState.UPDATE_INSTANCE:
				this.tablecolumnService.updateTableColumn(this.tablecolumn, this.GONG__StackPath)
					.subscribe(tablecolumn => {
						this.tablecolumnService.TableColumnServiceChanged.next("update")
					});
				break;
			default:
				this.tablecolumnService.postTableColumn(this.tablecolumn, this.GONG__StackPath).subscribe(tablecolumn => {
					this.tablecolumnService.TableColumnServiceChanged.next("post")
					this.tablecolumn = new (TableColumnDB) // reset fields
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

			dialogData.ID = this.tablecolumn.ID!
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
			dialogData.ID = this.tablecolumn.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "TableColumn"
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
			ID: this.tablecolumn.ID,
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
		if (this.tablecolumn.Name == "") {
			this.tablecolumn.Name = event.value.Name
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
