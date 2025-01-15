import { Component } from '@angular/core';
import { DialogComponent } from "../dialog/dialog.component";
import { ButtonComponent } from "../button/button.component";
import { AuthService } from '../../../features/auth/services/auth.service';

@Component({
  selector: 'app-aside-menu',
  imports: [DialogComponent, ButtonComponent],
  templateUrl: './aside-menu.component.html',
  styleUrl: './aside-menu.component.sass'
})
export class AsideMenuComponent {
  constructor(
    private authService: AuthService
  ) {}

  myUser?: User
  showMyUserDialog = false

  logout() {
    this.authService.logout()
  }

  async ngOnInit() {
    this.myUser = await this.authService.getMyUser()
  }
}
