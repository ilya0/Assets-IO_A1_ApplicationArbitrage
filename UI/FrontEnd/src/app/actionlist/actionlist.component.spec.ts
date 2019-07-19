import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ActionlistComponent } from './actionlist.component';

describe('ActionlistComponent', () => {
  let component: ActionlistComponent;
  let fixture: ComponentFixture<ActionlistComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ActionlistComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ActionlistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
