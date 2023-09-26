import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongdocxComponent } from './gongdocx.component';

describe('GongdocxComponent', () => {
  let component: GongdocxComponent;
  let fixture: ComponentFixture<GongdocxComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GongdocxComponent]
    });
    fixture = TestBed.createComponent(GongdocxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
