import { Component } from '@angular/core';
import { TextFieldComponent } from "../../../../shared/components/text-field/text-field.component";
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { IconComponent } from "../../../../shared/components/icon/icon.component";
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { emailValidatorRegex } from '../../../../shared/utils/regex';
import { DialogComponent } from "../../../../shared/components/dialog/dialog.component";
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-sign-in-form',
  imports: [TextFieldComponent, ButtonComponent, IconComponent, ReactiveFormsModule, DialogComponent],
  templateUrl: './sign-in-form.component.html',
  styleUrl: './sign-in-form.component.sass'
})
export class SignInFormComponent {
  constructor(
    private authService: AuthService
  ) {}

  form = new FormGroup({
    givenName: new FormControl('', [
      Validators.minLength(3),
      Validators.maxLength(24)
    ]),
    familyName: new FormControl('', [
      Validators.minLength(3),
      Validators.maxLength(24)
    ]),
    email: new FormControl('', [
      Validators.pattern(emailValidatorRegex)
    ]),
    password: new FormControl('', [
      Validators.minLength(6),
      Validators.maxLength(16)
    ]),
    repeatPassword: new FormControl('', [
      Validators.minLength(6),
      Validators.maxLength(16)
    ])
  })

  loading = false
  error?: string
  showError = false

  get isDisabled() {
    const email = this.form.get('email'),
      password = this.form.get('password'),
      repeatPassword = this.form.get('repeatPassword'),
      givenName = this.form.get('givenName'),
      familyName = this.form.get('familyName')

    return !email?.value
      || !password?.value
      || !repeatPassword?.value
      || !givenName?.value
      || !familyName?.value
      || this.form.invalid
  }

  async handleSubmit() {
    if (this.loading || this.isDisabled) return
    this.loading = true

    const password = this.form.get('password')?.value

    if (password !== this.form.get('repeatPassword')?.value) {
      this.error = 'Las contraseñas no coinciden'
      this.showError = true 
      this.loading = false
      return
    }
    
    try {
      await this.authService.signUp({
        email: this.form.get('email')?.value || '',
        password: password || '',
        givenName: this.form.get('givenName')?.value || '',
        familyName: this.form.get('familyName')?.value || ''
      })
    } catch (error: any) {
      if (error.status === 409) {
        this.error = 'Este correo electrónico ya está en uso'
      } else {
        this.error = 'Ha ocurrido un error inesperado'
      }
      this.showError = true
      this.loading = false
    }
  }
}
