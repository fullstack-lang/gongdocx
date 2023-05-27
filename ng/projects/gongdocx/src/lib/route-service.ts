import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { DocxsTableComponent } from './docxs-table/docxs-table.component'
import { DocxDetailComponent } from './docx-detail/docx-detail.component'


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



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getDocxTableRoute(stackPath),
            this.getDocxAdderRoute(stackPath),
            this.getDocxAdderForUseRoute(stackPath),
            this.getDocxDetailRoute(stackPath),

        ])
    }
}
