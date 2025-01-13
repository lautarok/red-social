import { Component } from '@angular/core';
import { LoginFormComponent } from "../../ui/login-form/login-form.component";
import { CardComponent } from "../../../../shared/components/card/card.component";

@Component({
  selector: 'app-login',
  imports: [LoginFormComponent, CardComponent],
  templateUrl: './login.component.html',
  styleUrl: './login.component.sass'
})
export class LoginComponent {

}
