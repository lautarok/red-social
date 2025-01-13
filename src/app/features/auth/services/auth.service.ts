import { Injectable } from '@angular/core';
import { ApiService } from '../../../shared/services/api.service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(
    private apiService: ApiService
  ) { }

  login(body: {email: string, password: string}) {
    return this.apiService.post<{token: string}>('auth', body)
  }
}
