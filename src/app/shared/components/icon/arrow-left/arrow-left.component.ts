import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-arrow-left',
  imports: [],
  templateUrl: './arrow-left.component.html',
  styleUrl: './arrow-left.component.sass'
})
export class ArrowLeftComponent {
  @Input() size?: string
}
