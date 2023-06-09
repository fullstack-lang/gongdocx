import { NgModule } from '@angular/core';

import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { GongdocModule } from 'gongdoc'
import { GongdocspecificModule } from 'gongdocspecific'

import { GongdocxModule } from 'gongdocx'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    DataModelPanelComponent
  ],
  imports: [
    GongdocModule,
    GongdocspecificModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongdocxModule,
  ],
  exports: [
    DataModelPanelComponent
  ]
})
export class GongdocxdatamodelModule { }
