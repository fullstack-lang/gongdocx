import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongdocxspecificComponent } from './gongdocxspecific.component';

describe('GongdocxspecificComponent', () => {
  let component: GongdocxspecificComponent;
  let fixture: ComponentFixture<GongdocxspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongdocxspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongdocxspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
