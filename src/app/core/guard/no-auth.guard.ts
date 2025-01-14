import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../../features/auth/services/auth.service';

export const noAuthGuard: CanActivateFn = (route, state) => {
  const authService = inject(AuthService),
    router = inject(Router)

  if (authService.isAuthenticated()) {
    router.navigate([''])
    return false
  }

  return true;
};
