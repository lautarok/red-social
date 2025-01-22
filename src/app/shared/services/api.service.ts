import { isPlatformBrowser } from '@angular/common';
import { HttpClient, HttpContext, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Inject, Injectable, PLATFORM_ID } from '@angular/core';
import { catchError, throwError } from 'rxjs';
import { environment } from '../../../environments/environment';
import { CACHING_ENABLED } from '../../core/interceptors/cache.interceptor';

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
      body?: Record<string, unknown>,
      invalidateCache?: boolean
    }
  ): Promise<T> | undefined {
    return new Promise((resolve, reject) => {
      this.httpClient.request(
        options?.method || 'GET',
        environment.apiUrl + '/' + path,
        {
          ...options,
          context: new HttpContext().set(CACHING_ENABLED, !!options?.invalidateCache)
        }
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

  get<T>(path: string, options?: {
    invalidateCache?: boolean
  }) {
    return this.api<T>(path, options)
  }
}
