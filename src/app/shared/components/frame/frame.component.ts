import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { AsideMenuComponent } from "../aside-menu/aside-menu.component";

@Component({
  selector: 'app-frame',
  imports: [RouterOutlet, AsideMenuComponent],
  templateUrl: './frame.component.html',
  styleUrl: './frame.component.sass'
})
export class FrameComponent {

}
