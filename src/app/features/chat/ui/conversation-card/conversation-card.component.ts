import { Component, Input } from '@angular/core';
import { RouterLink } from '@angular/router';
import parseDate from '../../../../shared/utils/parseDate';

@Component({
  selector: 'app-conversation-card',
  imports: [RouterLink],
  templateUrl: './conversation-card.component.html',
  styleUrl: './conversation-card.component.sass'
})
export class ConversationCardComponent {
  @Input() conversation!: Conversation
  @Input() selected = false

  parseDate(dateStr: string) {
    return parseDate(dateStr)
  }
}
