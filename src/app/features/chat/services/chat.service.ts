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
      email: peopleEmail
    })
  }

  getConversationList() {
    return this.apiService.get<Conversation[]>('conversation')
  }
}
