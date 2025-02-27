import { Component } from '@angular/core';
import { TextFieldComponent } from '../../../../shared/components/text-field/text-field.component';
import { ButtonComponent } from "../../../../shared/components/button/button.component";

@Component({
  selector: 'app-sell',
  imports: [TextFieldComponent, ButtonComponent],
  templateUrl: './sell.component.html',
  styleUrl: './sell.component.sass'
})
export class SellComponent {

}
