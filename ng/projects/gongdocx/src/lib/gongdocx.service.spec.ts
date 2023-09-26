import { TestBed } from '@angular/core/testing';

import { GongdocxService } from './gongdocx.service';

describe('GongdocxService', () => {
  let service: GongdocxService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongdocxService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
