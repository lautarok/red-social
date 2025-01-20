import { Component, Input } from '@angular/core';
import { LoaderComponent } from "../loader/loader.component";

@Component({
  selector: 'app-button',
  imports: [LoaderComponent],
  templateUrl: './button.component.html',
  styleUrl: './button.component.sass'
})
export class ButtonComponent {
  @Input() variant?: 'primary' | 'secondary' = 'secondary'
  @Input() type = 'button'
  @Input() loading = false
  @Input() disabled = false
  @Input() width?: string
  @Input() rounded = false
}
