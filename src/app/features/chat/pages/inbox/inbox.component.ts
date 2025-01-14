import { Component, Input, input } from '@angular/core';
import { CardComponent } from "../../../../shared/components/card/card.component";
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { AddConversationDialogComponent } from "../../ui/add-conversation-dialog/add-conversation-dialog.component";
import { ConversationListComponent } from "../../ui/conversation-list/conversation-list.component";

@Component({
  selector: 'app-inbox',
  imports: [ButtonComponent, AddConversationDialogComponent, ConversationListComponent],
  templateUrl: './inbox.component.html',
  styleUrl: './inbox.component.sass'
})
export class InboxComponent {
  @Input() conversation!: string
  showAddConversation = false
}
