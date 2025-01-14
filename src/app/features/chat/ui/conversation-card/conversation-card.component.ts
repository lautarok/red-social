import { Component, Input } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-conversation-card',
  imports: [RouterLink],
  templateUrl: './conversation-card.component.html',
  styleUrl: './conversation-card.component.sass'
})
export class ConversationCardComponent {
  @Input() conversation!: Conversation
}
