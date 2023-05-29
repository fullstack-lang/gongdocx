import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { DocumentsTableComponent } from './documents-table/documents-table.component'
import { DocumentDetailComponent } from './document-detail/document-detail.component'

import { DocxsTableComponent } from './docxs-table/docxs-table.component'
import { DocxDetailComponent } from './docx-detail/docx-detail.component'

import { FilesTableComponent } from './files-table/files-table.component'
import { FileDetailComponent } from './file-detail/file-detail.component'

import { NodesTableComponent } from './nodes-table/nodes-table.component'
import { NodeDetailComponent } from './node-detail/node-detail.component'

import { ParagraphsTableComponent } from './paragraphs-table/paragraphs-table.component'
import { ParagraphDetailComponent } from './paragraph-detail/paragraph-detail.component'

import { ParagraphPropertiessTableComponent } from './paragraphpropertiess-table/paragraphpropertiess-table.component'
import { ParagraphPropertiesDetailComponent } from './paragraphproperties-detail/paragraphproperties-detail.component'

import { ParagraphStylesTableComponent } from './paragraphstyles-table/paragraphstyles-table.component'
import { ParagraphStyleDetailComponent } from './paragraphstyle-detail/paragraphstyle-detail.component'

import { RunesTableComponent } from './runes-table/runes-table.component'
import { RuneDetailComponent } from './rune-detail/rune-detail.component'

import { RunePropertiessTableComponent } from './runepropertiess-table/runepropertiess-table.component'
import { RunePropertiesDetailComponent } from './runeproperties-detail/runeproperties-detail.component'

import { TablesTableComponent } from './tables-table/tables-table.component'
import { TableDetailComponent } from './table-detail/table-detail.component'

import { TableColumnsTableComponent } from './tablecolumns-table/tablecolumns-table.component'
import { TableColumnDetailComponent } from './tablecolumn-detail/tablecolumn-detail.component'

import { TablePropertiessTableComponent } from './tablepropertiess-table/tablepropertiess-table.component'
import { TablePropertiesDetailComponent } from './tableproperties-detail/tableproperties-detail.component'

import { TableRowsTableComponent } from './tablerows-table/tablerows-table.component'
import { TableRowDetailComponent } from './tablerow-detail/tablerow-detail.component'

import { TableStylesTableComponent } from './tablestyles-table/tablestyles-table.component'
import { TableStyleDetailComponent } from './tablestyle-detail/tablestyle-detail.component'

import { TextsTableComponent } from './texts-table/texts-table.component'
import { TextDetailComponent } from './text-detail/text-detail.component'


@Injectable({
    providedIn: 'root'
})
export class RouteService {
    private routes: Routes = []

    constructor(private router: Router) { }

    public addRoutes(newRoutes: Routes): void {
        const existingRoutes = this.router.config
        this.routes = this.router.config

        for (let newRoute of newRoutes) {
            if (!existingRoutes.includes(newRoute)) {
                this.routes.push(newRoute)
            }
        }
        this.router.resetConfig(this.routes)
    }

    getPathRoot(): string {
        return 'github_com_fullstack_lang_gongdocx_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getDocumentTablePath(): string {
        return this.getPathRoot() + '-documents/:GONG__StackPath'
    }
    getDocumentTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDocumentTablePath(), component: DocumentsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getDocumentAdderPath(): string {
        return this.getPathRoot() + '-document-adder/:GONG__StackPath'
    }
    getDocumentAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDocumentAdderPath(), component: DocumentDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDocumentAdderForUsePath(): string {
        return this.getPathRoot() + '-document-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getDocumentAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDocumentAdderForUsePath(), component: DocumentDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDocumentDetailPath(): string {
        return this.getPathRoot() + '-document-detail/:id/:GONG__StackPath'
    }
    getDocumentDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDocumentDetailPath(), component: DocumentDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getDocxTablePath(): string {
        return this.getPathRoot() + '-docxs/:GONG__StackPath'
    }
    getDocxTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDocxTablePath(), component: DocxsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getDocxAdderPath(): string {
        return this.getPathRoot() + '-docx-adder/:GONG__StackPath'
    }
    getDocxAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDocxAdderPath(), component: DocxDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDocxAdderForUsePath(): string {
        return this.getPathRoot() + '-docx-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getDocxAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDocxAdderForUsePath(), component: DocxDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDocxDetailPath(): string {
        return this.getPathRoot() + '-docx-detail/:id/:GONG__StackPath'
    }
    getDocxDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDocxDetailPath(), component: DocxDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFileTablePath(): string {
        return this.getPathRoot() + '-files/:GONG__StackPath'
    }
    getFileTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFileTablePath(), component: FilesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFileAdderPath(): string {
        return this.getPathRoot() + '-file-adder/:GONG__StackPath'
    }
    getFileAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFileAdderPath(), component: FileDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFileAdderForUsePath(): string {
        return this.getPathRoot() + '-file-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFileAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFileAdderForUsePath(), component: FileDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFileDetailPath(): string {
        return this.getPathRoot() + '-file-detail/:id/:GONG__StackPath'
    }
    getFileDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFileDetailPath(), component: FileDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getNodeTablePath(): string {
        return this.getPathRoot() + '-nodes/:GONG__StackPath'
    }
    getNodeTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNodeTablePath(), component: NodesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getNodeAdderPath(): string {
        return this.getPathRoot() + '-node-adder/:GONG__StackPath'
    }
    getNodeAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNodeAdderPath(), component: NodeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getNodeAdderForUsePath(): string {
        return this.getPathRoot() + '-node-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getNodeAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNodeAdderForUsePath(), component: NodeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getNodeDetailPath(): string {
        return this.getPathRoot() + '-node-detail/:id/:GONG__StackPath'
    }
    getNodeDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getNodeDetailPath(), component: NodeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getParagraphTablePath(): string {
        return this.getPathRoot() + '-paragraphs/:GONG__StackPath'
    }
    getParagraphTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphTablePath(), component: ParagraphsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getParagraphAdderPath(): string {
        return this.getPathRoot() + '-paragraph-adder/:GONG__StackPath'
    }
    getParagraphAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphAdderPath(), component: ParagraphDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getParagraphAdderForUsePath(): string {
        return this.getPathRoot() + '-paragraph-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getParagraphAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphAdderForUsePath(), component: ParagraphDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getParagraphDetailPath(): string {
        return this.getPathRoot() + '-paragraph-detail/:id/:GONG__StackPath'
    }
    getParagraphDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphDetailPath(), component: ParagraphDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getParagraphPropertiesTablePath(): string {
        return this.getPathRoot() + '-paragraphpropertiess/:GONG__StackPath'
    }
    getParagraphPropertiesTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphPropertiesTablePath(), component: ParagraphPropertiessTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getParagraphPropertiesAdderPath(): string {
        return this.getPathRoot() + '-paragraphproperties-adder/:GONG__StackPath'
    }
    getParagraphPropertiesAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphPropertiesAdderPath(), component: ParagraphPropertiesDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getParagraphPropertiesAdderForUsePath(): string {
        return this.getPathRoot() + '-paragraphproperties-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getParagraphPropertiesAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphPropertiesAdderForUsePath(), component: ParagraphPropertiesDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getParagraphPropertiesDetailPath(): string {
        return this.getPathRoot() + '-paragraphproperties-detail/:id/:GONG__StackPath'
    }
    getParagraphPropertiesDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphPropertiesDetailPath(), component: ParagraphPropertiesDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getParagraphStyleTablePath(): string {
        return this.getPathRoot() + '-paragraphstyles/:GONG__StackPath'
    }
    getParagraphStyleTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphStyleTablePath(), component: ParagraphStylesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getParagraphStyleAdderPath(): string {
        return this.getPathRoot() + '-paragraphstyle-adder/:GONG__StackPath'
    }
    getParagraphStyleAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphStyleAdderPath(), component: ParagraphStyleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getParagraphStyleAdderForUsePath(): string {
        return this.getPathRoot() + '-paragraphstyle-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getParagraphStyleAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphStyleAdderForUsePath(), component: ParagraphStyleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getParagraphStyleDetailPath(): string {
        return this.getPathRoot() + '-paragraphstyle-detail/:id/:GONG__StackPath'
    }
    getParagraphStyleDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getParagraphStyleDetailPath(), component: ParagraphStyleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getRuneTablePath(): string {
        return this.getPathRoot() + '-runes/:GONG__StackPath'
    }
    getRuneTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRuneTablePath(), component: RunesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getRuneAdderPath(): string {
        return this.getPathRoot() + '-rune-adder/:GONG__StackPath'
    }
    getRuneAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRuneAdderPath(), component: RuneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRuneAdderForUsePath(): string {
        return this.getPathRoot() + '-rune-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getRuneAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRuneAdderForUsePath(), component: RuneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRuneDetailPath(): string {
        return this.getPathRoot() + '-rune-detail/:id/:GONG__StackPath'
    }
    getRuneDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRuneDetailPath(), component: RuneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getRunePropertiesTablePath(): string {
        return this.getPathRoot() + '-runepropertiess/:GONG__StackPath'
    }
    getRunePropertiesTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRunePropertiesTablePath(), component: RunePropertiessTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getRunePropertiesAdderPath(): string {
        return this.getPathRoot() + '-runeproperties-adder/:GONG__StackPath'
    }
    getRunePropertiesAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRunePropertiesAdderPath(), component: RunePropertiesDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRunePropertiesAdderForUsePath(): string {
        return this.getPathRoot() + '-runeproperties-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getRunePropertiesAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRunePropertiesAdderForUsePath(), component: RunePropertiesDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRunePropertiesDetailPath(): string {
        return this.getPathRoot() + '-runeproperties-detail/:id/:GONG__StackPath'
    }
    getRunePropertiesDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRunePropertiesDetailPath(), component: RunePropertiesDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getTableTablePath(): string {
        return this.getPathRoot() + '-tables/:GONG__StackPath'
    }
    getTableTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableTablePath(), component: TablesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTableAdderPath(): string {
        return this.getPathRoot() + '-table-adder/:GONG__StackPath'
    }
    getTableAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableAdderPath(), component: TableDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableAdderForUsePath(): string {
        return this.getPathRoot() + '-table-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTableAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableAdderForUsePath(), component: TableDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableDetailPath(): string {
        return this.getPathRoot() + '-table-detail/:id/:GONG__StackPath'
    }
    getTableDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableDetailPath(), component: TableDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getTableColumnTablePath(): string {
        return this.getPathRoot() + '-tablecolumns/:GONG__StackPath'
    }
    getTableColumnTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableColumnTablePath(), component: TableColumnsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTableColumnAdderPath(): string {
        return this.getPathRoot() + '-tablecolumn-adder/:GONG__StackPath'
    }
    getTableColumnAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableColumnAdderPath(), component: TableColumnDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableColumnAdderForUsePath(): string {
        return this.getPathRoot() + '-tablecolumn-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTableColumnAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableColumnAdderForUsePath(), component: TableColumnDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableColumnDetailPath(): string {
        return this.getPathRoot() + '-tablecolumn-detail/:id/:GONG__StackPath'
    }
    getTableColumnDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableColumnDetailPath(), component: TableColumnDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getTablePropertiesTablePath(): string {
        return this.getPathRoot() + '-tablepropertiess/:GONG__StackPath'
    }
    getTablePropertiesTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTablePropertiesTablePath(), component: TablePropertiessTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTablePropertiesAdderPath(): string {
        return this.getPathRoot() + '-tableproperties-adder/:GONG__StackPath'
    }
    getTablePropertiesAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTablePropertiesAdderPath(), component: TablePropertiesDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTablePropertiesAdderForUsePath(): string {
        return this.getPathRoot() + '-tableproperties-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTablePropertiesAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTablePropertiesAdderForUsePath(), component: TablePropertiesDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTablePropertiesDetailPath(): string {
        return this.getPathRoot() + '-tableproperties-detail/:id/:GONG__StackPath'
    }
    getTablePropertiesDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTablePropertiesDetailPath(), component: TablePropertiesDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getTableRowTablePath(): string {
        return this.getPathRoot() + '-tablerows/:GONG__StackPath'
    }
    getTableRowTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableRowTablePath(), component: TableRowsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTableRowAdderPath(): string {
        return this.getPathRoot() + '-tablerow-adder/:GONG__StackPath'
    }
    getTableRowAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableRowAdderPath(), component: TableRowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableRowAdderForUsePath(): string {
        return this.getPathRoot() + '-tablerow-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTableRowAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableRowAdderForUsePath(), component: TableRowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableRowDetailPath(): string {
        return this.getPathRoot() + '-tablerow-detail/:id/:GONG__StackPath'
    }
    getTableRowDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableRowDetailPath(), component: TableRowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getTableStyleTablePath(): string {
        return this.getPathRoot() + '-tablestyles/:GONG__StackPath'
    }
    getTableStyleTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableStyleTablePath(), component: TableStylesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTableStyleAdderPath(): string {
        return this.getPathRoot() + '-tablestyle-adder/:GONG__StackPath'
    }
    getTableStyleAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableStyleAdderPath(), component: TableStyleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableStyleAdderForUsePath(): string {
        return this.getPathRoot() + '-tablestyle-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTableStyleAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableStyleAdderForUsePath(), component: TableStyleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTableStyleDetailPath(): string {
        return this.getPathRoot() + '-tablestyle-detail/:id/:GONG__StackPath'
    }
    getTableStyleDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTableStyleDetailPath(), component: TableStyleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getTextTablePath(): string {
        return this.getPathRoot() + '-texts/:GONG__StackPath'
    }
    getTextTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTextTablePath(), component: TextsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTextAdderPath(): string {
        return this.getPathRoot() + '-text-adder/:GONG__StackPath'
    }
    getTextAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTextAdderPath(), component: TextDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTextAdderForUsePath(): string {
        return this.getPathRoot() + '-text-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTextAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTextAdderForUsePath(), component: TextDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTextDetailPath(): string {
        return this.getPathRoot() + '-text-detail/:id/:GONG__StackPath'
    }
    getTextDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTextDetailPath(), component: TextDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getDocumentTableRoute(stackPath),
            this.getDocumentAdderRoute(stackPath),
            this.getDocumentAdderForUseRoute(stackPath),
            this.getDocumentDetailRoute(stackPath),

            this.getDocxTableRoute(stackPath),
            this.getDocxAdderRoute(stackPath),
            this.getDocxAdderForUseRoute(stackPath),
            this.getDocxDetailRoute(stackPath),

            this.getFileTableRoute(stackPath),
            this.getFileAdderRoute(stackPath),
            this.getFileAdderForUseRoute(stackPath),
            this.getFileDetailRoute(stackPath),

            this.getNodeTableRoute(stackPath),
            this.getNodeAdderRoute(stackPath),
            this.getNodeAdderForUseRoute(stackPath),
            this.getNodeDetailRoute(stackPath),

            this.getParagraphTableRoute(stackPath),
            this.getParagraphAdderRoute(stackPath),
            this.getParagraphAdderForUseRoute(stackPath),
            this.getParagraphDetailRoute(stackPath),

            this.getParagraphPropertiesTableRoute(stackPath),
            this.getParagraphPropertiesAdderRoute(stackPath),
            this.getParagraphPropertiesAdderForUseRoute(stackPath),
            this.getParagraphPropertiesDetailRoute(stackPath),

            this.getParagraphStyleTableRoute(stackPath),
            this.getParagraphStyleAdderRoute(stackPath),
            this.getParagraphStyleAdderForUseRoute(stackPath),
            this.getParagraphStyleDetailRoute(stackPath),

            this.getRuneTableRoute(stackPath),
            this.getRuneAdderRoute(stackPath),
            this.getRuneAdderForUseRoute(stackPath),
            this.getRuneDetailRoute(stackPath),

            this.getRunePropertiesTableRoute(stackPath),
            this.getRunePropertiesAdderRoute(stackPath),
            this.getRunePropertiesAdderForUseRoute(stackPath),
            this.getRunePropertiesDetailRoute(stackPath),

            this.getTableTableRoute(stackPath),
            this.getTableAdderRoute(stackPath),
            this.getTableAdderForUseRoute(stackPath),
            this.getTableDetailRoute(stackPath),

            this.getTableColumnTableRoute(stackPath),
            this.getTableColumnAdderRoute(stackPath),
            this.getTableColumnAdderForUseRoute(stackPath),
            this.getTableColumnDetailRoute(stackPath),

            this.getTablePropertiesTableRoute(stackPath),
            this.getTablePropertiesAdderRoute(stackPath),
            this.getTablePropertiesAdderForUseRoute(stackPath),
            this.getTablePropertiesDetailRoute(stackPath),

            this.getTableRowTableRoute(stackPath),
            this.getTableRowAdderRoute(stackPath),
            this.getTableRowAdderForUseRoute(stackPath),
            this.getTableRowDetailRoute(stackPath),

            this.getTableStyleTableRoute(stackPath),
            this.getTableStyleAdderRoute(stackPath),
            this.getTableStyleAdderForUseRoute(stackPath),
            this.getTableStyleDetailRoute(stackPath),

            this.getTextTableRoute(stackPath),
            this.getTextAdderRoute(stackPath),
            this.getTextAdderForUseRoute(stackPath),
            this.getTextDetailRoute(stackPath),

        ])
    }
}
