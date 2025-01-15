import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-key',
  imports: [],
  templateUrl: './key.component.html',
  styleUrl: './key.component.sass'
})
export class KeyComponent {
  @Input() size?: string
}
