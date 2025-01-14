import { Component } from '@angular/core';
import { CardComponent } from "../../../../shared/components/card/card.component";
import { RouterLink } from '@angular/router';
import { SignInFormComponent } from "../../ui/sign-in-form/sign-in-form.component";

@Component({
  selector: 'app-sign-in',
  imports: [CardComponent, RouterLink, SignInFormComponent],
  templateUrl: './sign-in.component.html',
  styleUrl: '../login/login.component.sass'
})
export class SignInComponent {

}
