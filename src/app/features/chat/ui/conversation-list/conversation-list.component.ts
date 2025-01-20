import { Component, Input, output } from '@angular/core';
import { ChatService } from '../../services/chat.service';
import { ConversationCardComponent } from "../conversation-card/conversation-card.component";
import { LoaderComponent } from '../../../../shared/components/loader/loader.component';
import { AuthService } from '../../../auth/services/auth.service';
import { ButtonComponent } from "../../../../shared/components/button/button.component";

@Component({
  selector: 'app-conversation-list',
  imports: [ConversationCardComponent, LoaderComponent, ButtonComponent],
  templateUrl: './conversation-list.component.html',
  styleUrl: './conversation-list.component.sass'
})
export class ConversationListComponent {
  constructor(
    private chatService: ChatService,
    private authService: AuthService
  ) {}

  @Input() selectedId?: string
  addConversation = output()

  myUser?: User
  conversations?: Conversation[]

  async ngOnInit() {
    this.myUser = await this.authService.getMyUser()

    const conversations = await this.chatService.getConversationList()
    if (conversations) {
      this.conversations = conversations?.map(conversation => {
        return {
          ...conversation,
          users: conversation.users.filter(user => user.id !== this.myUser?.id)
        }
      })
    } else {
      this.conversations = []
    }

    this.chatService.observeConversations().subscribe((data: {type: string, conversation: Conversation}) => {
      if (!this.conversations) return
      this.conversations = [
        ...this.conversations,
        {
          ...data.conversation,
          users: data.conversation.users.filter(user => user.id !== this.myUser?.id)
        }
      ]
    })

    this.chatService.observeMessages().subscribe((data: {type: string, message: Message}) => {
      if (!this.conversations) return

      const conversationIndex = this.conversations.map(conversation => conversation.id)
        .indexOf(data.message.conversationId)

      this.conversations[conversationIndex].lastMessage = data.message
      this.conversations = this.conversations.sort((a, b) => (b.lastMessage?.id || 0) - (a.lastMessage?.id || 0))
    })
  }
}
