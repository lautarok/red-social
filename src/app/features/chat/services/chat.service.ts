import { Injectable } from '@angular/core';
import { ApiService } from '../../../shared/services/api.service';

@Injectable({
  providedIn: 'root'
})
export class ChatService {
  constructor(
    private apiService: ApiService
  ) {}

  createConversation(peopleEmail: string) {
    return this.apiService.post('conversation', {
      id: peopleEmail
    })
  }

  getConversationList() {
    return this.apiService.get<Conversation[]>('conversation')
  }

  getConversation(id: number) {
    return this.apiService.get<Conversation>('conversation/' + id)
  }

  sendMessage(conversationId: number, message: string) {
    return this.apiService.post('message', {
      conversationId, message
    })
  }
}
