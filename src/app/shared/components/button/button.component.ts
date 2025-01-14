import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-button',
  imports: [],
  templateUrl: './button.component.html',
  styleUrl: './button.component.sass'
})
export class ButtonComponent {
  @Input() variant?: 'primary'
  @Input() type = 'button'
  @Input() loading = false
  @Input() disabled = false
  @Input() width?: string
  @Input() rounded = false
}
