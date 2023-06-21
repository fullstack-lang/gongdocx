import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongdocx from 'gongdocx'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongdocx Data/Model'
  view = this.default

  views: string[] = [this.default];

  DataStack = "gongdocx"
  ModelStacks = "github.com/fullstack-lang/gongdocx/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
