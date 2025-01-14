import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddConversationDialogComponent } from './add-conversation-dialog.component';

describe('AddConversationDialogComponent', () => {
  let component: AddConversationDialogComponent;
  let fixture: ComponentFixture<AddConversationDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AddConversationDialogComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddConversationDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
