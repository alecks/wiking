import { Component, Input } from '@angular/core';
import { AppConfiguration } from 'src/configuration';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss']
})
export class ToolbarComponent {

  @Input() title: string;
  config = AppConfiguration

  // TODO: get from API
  buttons = [{ name: "GitHub", icon: "code", href: "https://github.com/fjah/wiking" }]

}
