import { Component } from '@angular/core';
import { ChatService } from '../../services/chat.service';
import { ConversationCardComponent } from "../conversation-card/conversation-card.component";
import { LoaderComponent } from '../../../../shared/components/loader/loader.component';
import { AuthService } from '../../../auth/services/auth.service';

@Component({
  selector: 'app-conversation-list',
  imports: [ConversationCardComponent, LoaderComponent],
  templateUrl: './conversation-list.component.html',
  styleUrl: './conversation-list.component.sass'
})
export class ConversationListComponent {
  constructor(
    private chatService: ChatService,
    private authService: AuthService
  ) {}

  myUser?: User
  conversations?: Conversation[]

  async ngOnInit() {
    this.myUser = await this.authService.getMyUser()

    const conversations = await this.chatService.getConversationList()
    this.conversations = conversations?.map(conversation => {
      return {
        ...conversation,
        users: conversation.users.filter(user => user.id !== this.myUser?.id)
      }
    })
  }
}
