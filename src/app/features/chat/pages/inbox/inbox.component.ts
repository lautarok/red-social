import { Component, Input, input } from '@angular/core';
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { AddConversationDialogComponent } from "../../ui/add-conversation-dialog/add-conversation-dialog.component";
import { ConversationListComponent } from "../../ui/conversation-list/conversation-list.component";
import { ConversationComponent } from '../../ui/conversation/conversation.component';

@Component({
  selector: 'app-inbox',
  imports: [ButtonComponent, AddConversationDialogComponent, ConversationListComponent, ConversationComponent],
  templateUrl: './inbox.component.html',
  styleUrl: './inbox.component.sass'
})
export class InboxComponent {
  @Input() conversation!: string
  showAddConversation = false
}
