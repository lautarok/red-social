import { Component } from '@angular/core';
import { HeaderComponent } from "../header/header.component";
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-frame',
  imports: [HeaderComponent, RouterOutlet],
  templateUrl: './frame.component.html',
  styleUrl: './frame.component.sass'
})
export class FrameComponent {

}
