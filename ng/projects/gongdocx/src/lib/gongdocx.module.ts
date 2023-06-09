import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { BodysTableComponent } from './bodys-table/bodys-table.component'
import { BodySortingComponent } from './body-sorting/body-sorting.component'
import { BodyDetailComponent } from './body-detail/body-detail.component'

import { DocumentsTableComponent } from './documents-table/documents-table.component'
import { DocumentSortingComponent } from './document-sorting/document-sorting.component'
import { DocumentDetailComponent } from './document-detail/document-detail.component'

import { DocxsTableComponent } from './docxs-table/docxs-table.component'
import { DocxSortingComponent } from './docx-sorting/docx-sorting.component'
import { DocxDetailComponent } from './docx-detail/docx-detail.component'

import { FilesTableComponent } from './files-table/files-table.component'
import { FileSortingComponent } from './file-sorting/file-sorting.component'
import { FileDetailComponent } from './file-detail/file-detail.component'

import { NodesTableComponent } from './nodes-table/nodes-table.component'
import { NodeSortingComponent } from './node-sorting/node-sorting.component'
import { NodeDetailComponent } from './node-detail/node-detail.component'

import { ParagraphsTableComponent } from './paragraphs-table/paragraphs-table.component'
import { ParagraphSortingComponent } from './paragraph-sorting/paragraph-sorting.component'
import { ParagraphDetailComponent } from './paragraph-detail/paragraph-detail.component'

import { ParagraphPropertiessTableComponent } from './paragraphpropertiess-table/paragraphpropertiess-table.component'
import { ParagraphPropertiesSortingComponent } from './paragraphproperties-sorting/paragraphproperties-sorting.component'
import { ParagraphPropertiesDetailComponent } from './paragraphproperties-detail/paragraphproperties-detail.component'

import { ParagraphStylesTableComponent } from './paragraphstyles-table/paragraphstyles-table.component'
import { ParagraphStyleSortingComponent } from './paragraphstyle-sorting/paragraphstyle-sorting.component'
import { ParagraphStyleDetailComponent } from './paragraphstyle-detail/paragraphstyle-detail.component'

import { RunesTableComponent } from './runes-table/runes-table.component'
import { RuneSortingComponent } from './rune-sorting/rune-sorting.component'
import { RuneDetailComponent } from './rune-detail/rune-detail.component'

import { RunePropertiessTableComponent } from './runepropertiess-table/runepropertiess-table.component'
import { RunePropertiesSortingComponent } from './runeproperties-sorting/runeproperties-sorting.component'
import { RunePropertiesDetailComponent } from './runeproperties-detail/runeproperties-detail.component'

import { TablesTableComponent } from './tables-table/tables-table.component'
import { TableSortingComponent } from './table-sorting/table-sorting.component'
import { TableDetailComponent } from './table-detail/table-detail.component'

import { TableColumnsTableComponent } from './tablecolumns-table/tablecolumns-table.component'
import { TableColumnSortingComponent } from './tablecolumn-sorting/tablecolumn-sorting.component'
import { TableColumnDetailComponent } from './tablecolumn-detail/tablecolumn-detail.component'

import { TablePropertiessTableComponent } from './tablepropertiess-table/tablepropertiess-table.component'
import { TablePropertiesSortingComponent } from './tableproperties-sorting/tableproperties-sorting.component'
import { TablePropertiesDetailComponent } from './tableproperties-detail/tableproperties-detail.component'

import { TableRowsTableComponent } from './tablerows-table/tablerows-table.component'
import { TableRowSortingComponent } from './tablerow-sorting/tablerow-sorting.component'
import { TableRowDetailComponent } from './tablerow-detail/tablerow-detail.component'

import { TableStylesTableComponent } from './tablestyles-table/tablestyles-table.component'
import { TableStyleSortingComponent } from './tablestyle-sorting/tablestyle-sorting.component'
import { TableStyleDetailComponent } from './tablestyle-detail/tablestyle-detail.component'

import { TextsTableComponent } from './texts-table/texts-table.component'
import { TextSortingComponent } from './text-sorting/text-sorting.component'
import { TextDetailComponent } from './text-detail/text-detail.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		BodysTableComponent,
		BodySortingComponent,
		BodyDetailComponent,

		DocumentsTableComponent,
		DocumentSortingComponent,
		DocumentDetailComponent,

		DocxsTableComponent,
		DocxSortingComponent,
		DocxDetailComponent,

		FilesTableComponent,
		FileSortingComponent,
		FileDetailComponent,

		NodesTableComponent,
		NodeSortingComponent,
		NodeDetailComponent,

		ParagraphsTableComponent,
		ParagraphSortingComponent,
		ParagraphDetailComponent,

		ParagraphPropertiessTableComponent,
		ParagraphPropertiesSortingComponent,
		ParagraphPropertiesDetailComponent,

		ParagraphStylesTableComponent,
		ParagraphStyleSortingComponent,
		ParagraphStyleDetailComponent,

		RunesTableComponent,
		RuneSortingComponent,
		RuneDetailComponent,

		RunePropertiessTableComponent,
		RunePropertiesSortingComponent,
		RunePropertiesDetailComponent,

		TablesTableComponent,
		TableSortingComponent,
		TableDetailComponent,

		TableColumnsTableComponent,
		TableColumnSortingComponent,
		TableColumnDetailComponent,

		TablePropertiessTableComponent,
		TablePropertiesSortingComponent,
		TablePropertiesDetailComponent,

		TableRowsTableComponent,
		TableRowSortingComponent,
		TableRowDetailComponent,

		TableStylesTableComponent,
		TableStyleSortingComponent,
		TableStyleDetailComponent,

		TextsTableComponent,
		TextSortingComponent,
		TextDetailComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		BodysTableComponent,
		BodySortingComponent,
		BodyDetailComponent,

		DocumentsTableComponent,
		DocumentSortingComponent,
		DocumentDetailComponent,

		DocxsTableComponent,
		DocxSortingComponent,
		DocxDetailComponent,

		FilesTableComponent,
		FileSortingComponent,
		FileDetailComponent,

		NodesTableComponent,
		NodeSortingComponent,
		NodeDetailComponent,

		ParagraphsTableComponent,
		ParagraphSortingComponent,
		ParagraphDetailComponent,

		ParagraphPropertiessTableComponent,
		ParagraphPropertiesSortingComponent,
		ParagraphPropertiesDetailComponent,

		ParagraphStylesTableComponent,
		ParagraphStyleSortingComponent,
		ParagraphStyleDetailComponent,

		RunesTableComponent,
		RuneSortingComponent,
		RuneDetailComponent,

		RunePropertiessTableComponent,
		RunePropertiesSortingComponent,
		RunePropertiesDetailComponent,

		TablesTableComponent,
		TableSortingComponent,
		TableDetailComponent,

		TableColumnsTableComponent,
		TableColumnSortingComponent,
		TableColumnDetailComponent,

		TablePropertiessTableComponent,
		TablePropertiesSortingComponent,
		TablePropertiesDetailComponent,

		TableRowsTableComponent,
		TableRowSortingComponent,
		TableRowDetailComponent,

		TableStylesTableComponent,
		TableStyleSortingComponent,
		TableStyleDetailComponent,

		TextsTableComponent,
		TextSortingComponent,
		TextDetailComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongdocxModule { }
