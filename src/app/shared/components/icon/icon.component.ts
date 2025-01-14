import { Component, Input } from '@angular/core';
import { MentionComponent } from "./mention/mention.component";
import { ArrowRightComponent } from './arrow-right/arrow-right.component';
import { KeyComponent } from './key/key.component';

@Component({
  selector: 'app-icon',
  imports: [MentionComponent, ArrowRightComponent, KeyComponent],
  templateUrl: './icon.component.html',
  styleUrl: './icon.component.sass'
})
export class IconComponent {
  @Input() set!: string
  @Input() size?: string
}
