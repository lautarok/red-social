import { HttpContextToken, HttpInterceptorFn, HttpResponse } from '@angular/common/http';
import { catchError, filter, of, shareReplay, tap } from 'rxjs';

export const CACHING_ENABLED = new HttpContextToken<boolean>(() => true);

export const cacheMap = new Map<string, HttpResponse<any>>(),
  pendingRequests = new Map<string, any>()

export const cacheInterceptor: HttpInterceptorFn = (req, next) => {
  if (req.method !== 'GET') {
    return next(req)
  }

  if (req.context.get(CACHING_ENABLED)) {
    const cachedResponse = cacheMap.get(req.urlWithParams)
    if (cachedResponse) {
      return of(cachedResponse)
    } else if (pendingRequests.has(req.urlWithParams)) {
      return pendingRequests.get(req.urlWithParams)
    }

    const request = next(req).pipe(
      filter(event => event instanceof HttpResponse),
      tap(event => {
        cacheMap.set(req.urlWithParams, event)
      }),
      catchError((error) => {
        pendingRequests.delete(req.urlWithParams)
        throw error
      }),
      shareReplay(1)
    )

    pendingRequests.set(req.urlWithParams, request)

    request.subscribe({
      complete: () => pendingRequests.delete(req.urlWithParams)
    })

    return request
  }

  return next(req)
};
