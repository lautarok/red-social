import { Component, Input, output, SimpleChanges } from '@angular/core';
import { DialogComponent } from "../../../../shared/components/dialog/dialog.component";
import { TextFieldComponent } from "../../../../shared/components/text-field/text-field.component";
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { IconComponent } from "../../../../shared/components/icon/icon.component";
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { emailValidatorRegex } from '../../../../shared/utils/regex';
import { ChatService } from '../../services/chat.service';

@Component({
  selector: 'app-add-conversation-dialog',
  imports: [DialogComponent, TextFieldComponent, ButtonComponent, IconComponent, ReactiveFormsModule],
  templateUrl: './add-conversation-dialog.component.html',
  styleUrl: './add-conversation-dialog.component.sass'
})
export class AddConversationDialogComponent {
  constructor(
    private chatService: ChatService
  ) {}

  @Input() show!: boolean
  showChange = output<boolean>()
  onConversation = output<Conversation>()

  form = new FormGroup({
    email: new FormControl('', [
      Validators.pattern(emailValidatorRegex)
    ])
  })

  loading = false
  error?: string
  showError = false

  get isDisabled() {
    return !this.form.get('email')?.value
      || this.form.invalid
  }

  async handleSubmit() {
    if (this.loading || this.isDisabled) return
    this.loading = true
    
    try {
      const response = await this.chatService.createConversation(this.form.get('email')?.value || '')
      if (response) {
        this.onConversation.emit(response)
      }
    } catch (error: any) {
      if (error.status === 404) {
        this.error = 'No se ha encontrado el usuario'
      } else if (error.status === 409) {
        this.error = 'Ya existe una conversación con este usuario'
      } else {
        this.error = 'Ha ocurrido un error desconocido'
      }
      this.showError = true
      this.loading = false
    }
  }

  ngOnChanges(changes: SimpleChanges) {
    if (changes['show'] && this.show === true) {
      this.form.reset()
    }
  }
}
