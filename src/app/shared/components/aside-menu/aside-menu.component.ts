import { Component } from '@angular/core';
import { DialogComponent } from "../dialog/dialog.component";
import { ButtonComponent } from "../button/button.component";
import { AuthService } from '../../../features/auth/services/auth.service';
import { RouterLink } from '@angular/router';
import { IconComponent } from "../icon/icon.component";

@Component({
  selector: 'app-aside-menu',
  imports: [DialogComponent, ButtonComponent, RouterLink, IconComponent],
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
