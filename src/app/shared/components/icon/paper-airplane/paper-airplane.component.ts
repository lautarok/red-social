import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-paper-airplane',
  imports: [],
  templateUrl: './paper-airplane.component.html',
  styleUrl: './paper-airplane.component.sass'
})
export class PaperAirplaneComponent {
  @Input() size?: string
}
