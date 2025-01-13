import { Component, input } from '@angular/core';

@Component({
  selector: 'app-button',
  imports: [],
  templateUrl: './button.component.html',
  styleUrl: './button.component.sass'
})
export class ButtonComponent {
  variant = input<'primary'>()
  type = input<string>('button')
  loading = input(false)
  disabled = input(false)
}
