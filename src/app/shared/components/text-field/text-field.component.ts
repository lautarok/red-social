import { Component, ElementRef, forwardRef, input, viewChild } from '@angular/core';
import { IconComponent } from '../icon/icon.component';
import { ControlValueAccessor, FormsModule, NG_VALUE_ACCESSOR } from '@angular/forms';

@Component({
  selector: 'app-text-field',
  imports: [IconComponent, FormsModule],
  templateUrl: './text-field.component.html',
  styleUrl: './text-field.component.sass',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      multi: true,
      useExisting: forwardRef(() => TextFieldComponent)
    }
  ]
})
export class TextFieldComponent implements ControlValueAccessor {
  label = input.required<string>()
  icon = input<string>()
  type = input<string>('text')

  inputElement = viewChild.required<ElementRef<HTMLInputElement>>('input')

  focusInput() {
    this.inputElement().nativeElement.focus()
  }

  value = ''

  private onChange = (value: string) => {}
  private onTouched = () => {}

  writeValue(value: string): void {
    this.value = value
  }

  registerOnChange(fn: any): void {
    this.onChange = fn
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn
  }

  onInput(event: Event) {
    const target = event.target as HTMLInputElement
    this.value = target.value
    this.onChange(this.value)
  }
}
