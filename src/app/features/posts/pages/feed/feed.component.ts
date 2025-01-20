import { Component } from '@angular/core';
import { ButtonComponent } from "../../../../shared/components/button/button.component";
import { IconComponent } from "../../../../shared/components/icon/icon.component";
import { Router } from '@angular/router';

@Component({
  selector: 'app-feed',
  imports: [ButtonComponent, IconComponent],
  templateUrl: './feed.component.html',
  styleUrl: './feed.component.sass'
})
export class FeedComponent {
  constructor(
    private router: Router
  ) {}

  goToChat() {
    this.router.navigate(['chat'])
  }
}
