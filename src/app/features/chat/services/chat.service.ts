import { Injectable } from '@angular/core';
import { ApiService } from '../../../shared/services/api.service';
import { WsService } from '../../../shared/services/ws.service';
import { filter, map, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ChatService {
  constructor(
    private apiService: ApiService,
    private WsService: WsService
  ) {}

  createConversation(peopleEmail: string) {
    return this.apiService.post<Conversation>('conversation', {
      email: peopleEmail
    })
  }

  getConversationList() {
    return this.apiService.get<Conversation[]>('conversation')
  }

  getConversation(id: number) {
    return this.apiService.get<Conversation>('conversation/' + id)
  }

  sendMessage(conversationId: number, message: string) {
    return this.apiService.post<{insertedId: number}>('message', {
      conversationId,
      message
    })
  }

  observeMessages(): Observable<any> {
    return this.WsService.getObservable().pipe(
      map((event: MessageEvent) => {
        try {
          return JSON.parse(event.data)
        } catch (error) {
          console.error(error)
          return null
        }
      }),
      filter((parsedData) => parsedData !== null && parsedData.type === 'message')
    )
  }

  observeConversations(): Observable<any> {
    return this.WsService.getObservable().pipe(
      map((event: MessageEvent) => {
        try {
          return JSON.parse(event.data)
        } catch (error) {
          console.error(error)
          return null
        }
      }),
      filter((parsedData) => parsedData !== null && parsedData.type === 'conversation')
    )
  }
}
