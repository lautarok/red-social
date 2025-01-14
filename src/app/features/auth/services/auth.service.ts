import { Injectable } from '@angular/core';
import { ApiService } from '../../../shared/services/api.service';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(
    private apiService: ApiService,
    private router: Router
  ) { }

  private sessionStart(response: any) {
    if (response.token) {
      localStorage.setItem('auth_token', response.token)
      this.router.navigate([''])
    }
  }

  isAuthenticated() {
    return !!localStorage.getItem('auth_token')
  }

  async login(body: {email: string, password: string}) {
    try {
      const response = await this.apiService.post<{token: string}>('auth', body)      
      this.sessionStart(response)
      return response
    } catch (error) {
      throw error
    }
  }

  async signUp(body: {
    email: string,
    password: string,
    givenName: string,
    familyName: string
  }) {
    try {
      const response = await this.apiService.put<{token: string}>('auth/sign-up', body)
      this.sessionStart(response)
      return response
    } catch (error) {
      throw error
    }
  }

  getMyUser() {
    return this.apiService.get<User>('user/me')
  }
}
