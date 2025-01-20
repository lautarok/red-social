import { isPlatformBrowser } from '@angular/common';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Inject, Injectable, PLATFORM_ID } from '@angular/core';
import { catchError, throwError } from 'rxjs';
import { environment } from '../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(
    private httpClient: HttpClient
  ) { }

  private api<T>(
    path: string,
    options?: {
      method?: 'GET' | 'POST' | 'PUT' | 'DELETE',
      headers?: HttpHeaders,
      body?: Record<string, unknown>
    }
  ): Promise<T> | undefined {
    return new Promise((resolve, reject) => {
      this.httpClient.request(
        options?.method || 'GET', environment.apiUrl + '/' + path, options
      ).pipe(
        catchError((error: HttpErrorResponse) => {
          reject(error)
          return throwError(() => new Error(error.message))
        })
      ).subscribe((data: any) => {
        resolve(data as T)
      })
    })
  }

  post<T>(path: string, body: Record<string, unknown>) {
    return this.api<T>(path, {
      body,
      method: 'POST'
    })
  }

  put<T>(path: string, body: Record<string, unknown>) {
    return this.api<T>(path, {
      body,
      method: 'PUT'
    })
  }

  get<T>(path: string) {
    return this.api<T>(path)
  }
}
