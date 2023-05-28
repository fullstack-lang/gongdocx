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

            this.getTextTableRoute(stackPath),
            this.getTextAdderRoute(stackPath),
            this.getTextAdderForUseRoute(stackPath),
            this.getTextDetailRoute(stackPath),

        ])
    }
}
