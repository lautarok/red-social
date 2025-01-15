import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-arrow-right',
  imports: [],
  templateUrl: './arrow-right.component.html',
  styleUrl: './arrow-right.component.sass'
})
export class ArrowRightComponent {
  @Input() size?: string
}
