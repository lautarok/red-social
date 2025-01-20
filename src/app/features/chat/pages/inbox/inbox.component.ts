import { Component, Input } from '@angular/core';
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { AddConversationDialogComponent } from "../../ui/add-conversation-dialog/add-conversation-dialog.component";
import { ConversationListComponent } from "../../ui/conversation-list/conversation-list.component";
import { ConversationComponent } from '../../ui/conversation/conversation.component';
import { IconComponent } from "../../../../shared/components/icon/icon.component";

@Component({
  selector: 'app-inbox',
  imports: [ButtonComponent, AddConversationDialogComponent, ConversationListComponent, ConversationComponent, IconComponent],
  templateUrl: './inbox.component.html',
  styleUrl: './inbox.component.sass'
})
export class InboxComponent {
  @Input() conversation!: string
  showAddConversation = false
}
