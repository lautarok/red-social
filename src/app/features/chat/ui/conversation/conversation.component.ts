import { Component, ElementRef, Input, SimpleChanges, viewChild } from '@angular/core';
import { RouterLink } from '@angular/router';
import { IconComponent } from "../../../../shared/components/icon/icon.component";
import { ChatService } from '../../services/chat.service';
import { AuthService } from '../../../auth/services/auth.service';
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { LoaderComponent } from "../../../../shared/components/loader/loader.component";

@Component({
  selector: 'app-conversation',
  imports: [RouterLink, IconComponent, ButtonComponent, ReactiveFormsModule, LoaderComponent],
  templateUrl: './conversation.component.html',
  styleUrl: './conversation.component.sass'
})
export class ConversationComponent {
  constructor(
    private chatService: ChatService,
    private authService: AuthService
  ) {}

  @Input() id!: string

  messagesScrollElement = viewChild<ElementRef<HTMLDivElement>>('messagesScroll')
  messagesWrapperElement = viewChild<ElementRef<HTMLDivElement>>('messagesWrapper')

  conversation?: Conversation
  myUser?: User

  scrollReady = false

  form = new FormGroup({
    message: new FormControl('')
  })

  async handleSubmit() {
    const message = this.form.get('message')?.value
    if (!this.conversation?.id || !message) return
    await this.chatService.sendMessage(this.conversation.id, message)
    this.form.reset()
  }

  async fetchData() {
    this.conversation = undefined
    this.myUser = await this.authService.getMyUser()
    const conversation = await this.chatService.getConversation(parseInt(this.id))!
    this.conversation = {
      ...conversation,
      users: conversation.users.filter(user => user.id !== this.myUser?.id)
    }
  }

  async ngOnChanges(changes: SimpleChanges) {
    if (changes['id']) {
      await this.fetchData()
      this.scrollReady = false
      this.scrollToBottom()
    }
  }

  private scrollToBottom() {
    if (this.scrollReady) return

    const messagesScroll = this.messagesScrollElement()?.nativeElement,
      messagesWrapper = this.messagesWrapperElement()?.nativeElement

    if (messagesScroll && messagesWrapper) {
      messagesScroll.scrollTop = messagesWrapper.clientHeight
      setTimeout(() => {
        messagesScroll.scrollTop = messagesWrapper.clientHeight
        this.scrollReady = true
      }, 10)
    }
  }

  ngAfterViewInit() {
    this.scrollToBottom()
  }

  ngAfterViewChecked() {
    this.scrollToBottom()
  }
}
