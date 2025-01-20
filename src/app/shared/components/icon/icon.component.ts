import { Component, Input } from '@angular/core';
import { MentionComponent } from "./mention/mention.component";
import { ArrowRightComponent } from './arrow-right/arrow-right.component';
import { KeyComponent } from './key/key.component';
import { ArrowLeftComponent } from "./arrow-left/arrow-left.component";
import { HomeComponent } from "./home/home.component";
import { PaperAirplaneComponent } from "./paper-airplane/paper-airplane.component";

@Component({
  selector: 'app-icon',
  imports: [MentionComponent, ArrowRightComponent, KeyComponent, ArrowLeftComponent, HomeComponent, PaperAirplaneComponent],
  templateUrl: './icon.component.html',
  styleUrl: './icon.component.sass'
})
export class IconComponent {
  @Input() set!: string
  @Input() size?: string
}
