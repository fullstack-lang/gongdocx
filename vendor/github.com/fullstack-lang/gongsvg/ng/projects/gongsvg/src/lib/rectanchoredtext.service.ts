// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { RectAnchoredTextDB } from './rectanchoredtext-db';

// insertion point for imports
import { RectDB } from './rect-db'

@Injectable({
  providedIn: 'root'
})
export class RectAnchoredTextService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  RectAnchoredTextServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private rectanchoredtextsUrl: string

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
    this.rectanchoredtextsUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/rectanchoredtexts';
  }

  /** GET rectanchoredtexts from the server */
  getRectAnchoredTexts(GONG__StackPath: string): Observable<RectAnchoredTextDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<RectAnchoredTextDB[]>(this.rectanchoredtextsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched rectanchoredtexts')),
        catchError(this.handleError<RectAnchoredTextDB[]>('getRectAnchoredTexts', []))
      );
  }

  /** GET rectanchoredtext by id. Will 404 if id not found */
  getRectAnchoredText(id: number, GONG__StackPath: string): Observable<RectAnchoredTextDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.rectanchoredtextsUrl}/${id}`;
    return this.http.get<RectAnchoredTextDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched rectanchoredtext id=${id}`)),
      catchError(this.handleError<RectAnchoredTextDB>(`getRectAnchoredText id=${id}`))
    );
  }

  /** POST: add a new rectanchoredtext to the server */
  postRectAnchoredText(rectanchoredtextdb: RectAnchoredTextDB, GONG__StackPath: string): Observable<RectAnchoredTextDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    rectanchoredtextdb.Animates = []
    let _Rect_RectAnchoredTexts_reverse = rectanchoredtextdb.Rect_RectAnchoredTexts_reverse
    rectanchoredtextdb.Rect_RectAnchoredTexts_reverse = new RectDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RectAnchoredTextDB>(this.rectanchoredtextsUrl, rectanchoredtextdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        rectanchoredtextdb.Rect_RectAnchoredTexts_reverse = _Rect_RectAnchoredTexts_reverse
        // this.log(`posted rectanchoredtextdb id=${rectanchoredtextdb.ID}`)
      }),
      catchError(this.handleError<RectAnchoredTextDB>('postRectAnchoredText'))
    );
  }

  /** DELETE: delete the rectanchoredtextdb from the server */
  deleteRectAnchoredText(rectanchoredtextdb: RectAnchoredTextDB | number, GONG__StackPath: string): Observable<RectAnchoredTextDB> {
    const id = typeof rectanchoredtextdb === 'number' ? rectanchoredtextdb : rectanchoredtextdb.ID;
    const url = `${this.rectanchoredtextsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<RectAnchoredTextDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted rectanchoredtextdb id=${id}`)),
      catchError(this.handleError<RectAnchoredTextDB>('deleteRectAnchoredText'))
    );
  }

  /** PUT: update the rectanchoredtextdb on the server */
  updateRectAnchoredText(rectanchoredtextdb: RectAnchoredTextDB, GONG__StackPath: string): Observable<RectAnchoredTextDB> {
    const id = typeof rectanchoredtextdb === 'number' ? rectanchoredtextdb : rectanchoredtextdb.ID;
    const url = `${this.rectanchoredtextsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    rectanchoredtextdb.Animates = []
    let _Rect_RectAnchoredTexts_reverse = rectanchoredtextdb.Rect_RectAnchoredTexts_reverse
    rectanchoredtextdb.Rect_RectAnchoredTexts_reverse = new RectDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<RectAnchoredTextDB>(url, rectanchoredtextdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        rectanchoredtextdb.Rect_RectAnchoredTexts_reverse = _Rect_RectAnchoredTexts_reverse
        // this.log(`updated rectanchoredtextdb id=${rectanchoredtextdb.ID}`)
      }),
      catchError(this.handleError<RectAnchoredTextDB>('updateRectAnchoredText'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in RectAnchoredTextService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("RectAnchoredTextService" + error); // log to console instead

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
