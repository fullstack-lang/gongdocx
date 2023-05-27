import { TestBed } from '@angular/core/testing';

import { GongdocxspecificService } from './gongdocxspecific.service';

describe('GongdocxspecificService', () => {
  let service: GongdocxspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongdocxspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
