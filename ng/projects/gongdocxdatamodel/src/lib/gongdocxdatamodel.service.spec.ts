import { TestBed } from '@angular/core/testing';

import { GongdocxdatamodelService } from './gongdocxdatamodel.service';

describe('GongdocxdatamodelService', () => {
  let service: GongdocxdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongdocxdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
