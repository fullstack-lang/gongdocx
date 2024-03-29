// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { SVGIconDB } from './svgicon-db'
import { SVGIcon, CopySVGIconToSVGIconDB } from './svgicon'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class SVGIconService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  SVGIconServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private svgiconsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.svgiconsUrl = origin + '/api/github.com/fullstack-lang/gongtree/go/v1/svgicons';
  }

  /** GET svgicons from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGIconDB[]> {
    return this.getSVGIcons(GONG__StackPath, frontRepo)
  }
  getSVGIcons(GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGIconDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<SVGIconDB[]>(this.svgiconsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<SVGIconDB[]>('getSVGIcons', []))
      );
  }

  /** GET svgicon by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGIconDB> {
    return this.getSVGIcon(id, GONG__StackPath, frontRepo)
  }
  getSVGIcon(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGIconDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.svgiconsUrl}/${id}`;
    return this.http.get<SVGIconDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched svgicon id=${id}`)),
      catchError(this.handleError<SVGIconDB>(`getSVGIcon id=${id}`))
    );
  }

  // postFront copy svgicon to a version with encoded pointers and post to the back
  postFront(svgicon: SVGIcon, GONG__StackPath: string): Observable<SVGIconDB> {
    let svgiconDB = new SVGIconDB
    CopySVGIconToSVGIconDB(svgicon, svgiconDB)
    const id = typeof svgiconDB === 'number' ? svgiconDB : svgiconDB.ID
    const url = `${this.svgiconsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<SVGIconDB>(url, svgiconDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<SVGIconDB>('postSVGIcon'))
    );
  }
  
  /** POST: add a new svgicon to the server */
  post(svgicondb: SVGIconDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGIconDB> {
    return this.postSVGIcon(svgicondb, GONG__StackPath, frontRepo)
  }
  postSVGIcon(svgicondb: SVGIconDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGIconDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<SVGIconDB>(this.svgiconsUrl, svgicondb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted svgicondb id=${svgicondb.ID}`)
      }),
      catchError(this.handleError<SVGIconDB>('postSVGIcon'))
    );
  }

  /** DELETE: delete the svgicondb from the server */
  delete(svgicondb: SVGIconDB | number, GONG__StackPath: string): Observable<SVGIconDB> {
    return this.deleteSVGIcon(svgicondb, GONG__StackPath)
  }
  deleteSVGIcon(svgicondb: SVGIconDB | number, GONG__StackPath: string): Observable<SVGIconDB> {
    const id = typeof svgicondb === 'number' ? svgicondb : svgicondb.ID;
    const url = `${this.svgiconsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<SVGIconDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted svgicondb id=${id}`)),
      catchError(this.handleError<SVGIconDB>('deleteSVGIcon'))
    );
  }

  // updateFront copy svgicon to a version with encoded pointers and update to the back
  updateFront(svgicon: SVGIcon, GONG__StackPath: string): Observable<SVGIconDB> {
    let svgiconDB = new SVGIconDB
    CopySVGIconToSVGIconDB(svgicon, svgiconDB)
    const id = typeof svgiconDB === 'number' ? svgiconDB : svgiconDB.ID
    const url = `${this.svgiconsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<SVGIconDB>(url, svgiconDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<SVGIconDB>('updateSVGIcon'))
    );
  }

  /** PUT: update the svgicondb on the server */
  update(svgicondb: SVGIconDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGIconDB> {
    return this.updateSVGIcon(svgicondb, GONG__StackPath, frontRepo)
  }
  updateSVGIcon(svgicondb: SVGIconDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SVGIconDB> {
    const id = typeof svgicondb === 'number' ? svgicondb : svgicondb.ID;
    const url = `${this.svgiconsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<SVGIconDB>(url, svgicondb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated svgicondb id=${svgicondb.ID}`)
      }),
      catchError(this.handleError<SVGIconDB>('updateSVGIcon'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in SVGIconService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("SVGIconService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
