import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';
import { environment } from '../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class WsService {
  private subject = new Subject<any>
  socket?: WebSocket

  constructor() {
    this.socket = new WebSocket(environment.apiUrl + '/ws?auth_token=' + localStorage.getItem('auth_token'))

    this.socket.onmessage = event => {
      this.subject.next(event)
    }

    this.socket.onclose = () => {
      window.location.reload()
    }

    this.socket.onerror = () => {
      window.location.reload()
    }
  }

  getObservable() {
    return this.subject.asObservable()
  }
}
