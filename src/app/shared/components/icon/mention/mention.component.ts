import { Component, Input, input } from '@angular/core';

@Component({
  selector: 'app-mention',
  imports: [],
  templateUrl: './mention.component.html',
  styleUrl: './mention.component.sass'
})
export class MentionComponent {
  @Input() size?: string
}
