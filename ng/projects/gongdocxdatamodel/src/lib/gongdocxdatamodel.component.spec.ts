import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongdocxdatamodelComponent } from './gongdocxdatamodel.component';

describe('GongdocxdatamodelComponent', () => {
  let component: GongdocxdatamodelComponent;
  let fixture: ComponentFixture<GongdocxdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongdocxdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongdocxdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
