import { Component } from '@angular/core';
import { TextFieldComponent } from '../../../../shared/components/text-field/text-field.component';
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { IconComponent } from "../../../../shared/components/icon/icon.component";
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { AuthService } from '../../services/auth.service';
import { DialogComponent } from "../../../../shared/components/dialog/dialog.component";
import { emailValidatorRegex } from '../../../../shared/utils/regex';

@Component({
  selector: 'app-login-form',
  imports: [TextFieldComponent, ButtonComponent, IconComponent, ReactiveFormsModule, DialogComponent],
  templateUrl: './login-form.component.html',
  styleUrl: './login-form.component.sass'
})
export class LoginFormComponent {
  constructor(
    private authService: AuthService
  ) {}

  form = new FormGroup({
    email: new FormControl('', [
      Validators.pattern(emailValidatorRegex)
    ]),
    password: new FormControl('', [
      Validators.minLength(6),
      Validators.maxLength(16)
    ])
  })

  loading = false
  error?: string
  showError = false

  get isDisabled() {
    const email = this.form.get('email'),
      password = this.form.get('password')

    return !email?.value
      || !password?.value
      || this.form.invalid
  }

  async handleSubmit() {
    if (this.loading || this.isDisabled) return
    this.loading = true

    try {
      await this.authService.login({
        email: this.form.get('email')?.value || '',
        password: this.form.get('password')?.value || ''
      })
    } catch (error: any) {
      if (error.status === 404) {
        this.error = `El correo electrónico no se encuentra registrado`
      } else if (error.status === 401) {
        this.error = 'La contraseña es incorrecta'
        this.form.get('password')?.reset()
      } else {
        this.error = 'Ha ocurrido un error desconocido'
      }
      this.showError = true
      this.loading = false
    }
  }
}
