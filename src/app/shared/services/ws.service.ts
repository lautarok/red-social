import { Injectable } from '@angular/core';
import { ReplaySubject } from 'rxjs';
import { environment } from '../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class WsService {
  private subject = new ReplaySubject<any>()
  socket?: WebSocket

  constructor() {
    const authToken = localStorage.getItem('auth_token')
    if (authToken) {
      this.connect(authToken)
    }
  }

  private reconnect() {
    window.location.reload()
  }

  disconnect() {
    this.socket?.close()
  }

  connect(authToken: string) {
    this.socket = new WebSocket(environment.apiUrl + '/ws?auth_token=' + authToken)

    this.socket.onmessage = event => {
      this.subject.next(event)
    }

    this.socket.onclose = event => {
      if (!event.wasClean) {
        this.reconnect()
      }
    }

    this.socket.onerror = () => {
      this.reconnect()
    }
  }

  getObservable() {
    return this.subject.asObservable()
  }
}
