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

import { RuneDB } from './rune-db';

// insertion point for imports
import { NodeDB } from './node-db'
import { TextDB } from './text-db'
import { RunePropertiesDB } from './runeproperties-db'
import { ParagraphDB } from './paragraph-db'

@Injectable({
  providedIn: 'root'
})
export class RuneService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  RuneServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private runesUrl: string

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
    this.runesUrl = origin + '/api/github.com/fullstack-lang/gongdocx/go/v1/runes';
  }

  /** GET runes from the server */
  getRunes(GONG__StackPath: string): Observable<RuneDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<RuneDB[]>(this.runesUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched runes')),
        catchError(this.handleError<RuneDB[]>('getRunes', []))
      );
  }

  /** GET rune by id. Will 404 if id not found */
  getRune(id: number, GONG__StackPath: string): Observable<RuneDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.runesUrl}/${id}`;
    return this.http.get<RuneDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched rune id=${id}`)),
      catchError(this.handleError<RuneDB>(`getRune id=${id}`))
    );
  }

  /** POST: add a new rune to the server */
  postRune(runedb: RuneDB, GONG__StackPath: string): Observable<RuneDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    runedb.Node = new NodeDB
    runedb.Text = new TextDB
    runedb.RuneProperties = new RunePropertiesDB
    runedb.EnclosingParagraph = new ParagraphDB
    let _Paragraph_Runes_reverse = runedb.Paragraph_Runes_reverse
    runedb.Paragraph_Runes_reverse = new ParagraphDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RuneDB>(this.runesUrl, runedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        runedb.Paragraph_Runes_reverse = _Paragraph_Runes_reverse
        // this.log(`posted runedb id=${runedb.ID}`)
      }),
      catchError(this.handleError<RuneDB>('postRune'))
    );
  }

  /** DELETE: delete the runedb from the server */
  deleteRune(runedb: RuneDB | number, GONG__StackPath: string): Observable<RuneDB> {
    const id = typeof runedb === 'number' ? runedb : runedb.ID;
    const url = `${this.runesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<RuneDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted runedb id=${id}`)),
      catchError(this.handleError<RuneDB>('deleteRune'))
    );
  }

  /** PUT: update the runedb on the server */
  updateRune(runedb: RuneDB, GONG__StackPath: string): Observable<RuneDB> {
    const id = typeof runedb === 'number' ? runedb : runedb.ID;
    const url = `${this.runesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    runedb.Node = new NodeDB
    runedb.Text = new TextDB
    runedb.RuneProperties = new RunePropertiesDB
    runedb.EnclosingParagraph = new ParagraphDB
    let _Paragraph_Runes_reverse = runedb.Paragraph_Runes_reverse
    runedb.Paragraph_Runes_reverse = new ParagraphDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<RuneDB>(url, runedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        runedb.Paragraph_Runes_reverse = _Paragraph_Runes_reverse
        // this.log(`updated runedb id=${runedb.ID}`)
      }),
      catchError(this.handleError<RuneDB>('updateRune'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in RuneService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("RuneService" + error); // log to console instead

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
