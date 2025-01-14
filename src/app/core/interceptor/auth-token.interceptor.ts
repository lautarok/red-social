import { HttpInterceptorFn } from '@angular/common/http';

export const authTokenInterceptor: HttpInterceptorFn = (req, next) => {
  const token = localStorage.getItem('auth_token')

  if (token) {
    req = req.clone({
      headers: req.headers.set('Authorization', 'Bearer ' + token)
    })
  }

  return next(req);
};
