import { HttpContextToken, HttpInterceptorFn, HttpResponse } from '@angular/common/http';
import { of, tap } from 'rxjs';

export const CACHING_ENABLED = new HttpContextToken<boolean>(() => true);

const cacheMap = new Map<string, HttpResponse<any>>()

export const cacheInterceptor: HttpInterceptorFn = (req, next) => {
  if (req.method !== 'GET') {
    return next(req)
  }

  if (req.context.get(CACHING_ENABLED)) {
    const cachedResponse = cacheMap.get(req.urlWithParams)
    if (cachedResponse) {
      console.log(cachedResponse)
      return of(cachedResponse)
    }
  }

  return next(req).pipe(
    tap(event => {
      if (event instanceof HttpResponse) {
        cacheMap.set(req.urlWithParams, event)
      }
    })
  );
};
