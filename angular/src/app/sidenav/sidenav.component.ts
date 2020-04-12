import { Component, ViewChild, OnInit, ElementRef } from '@angular/core';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.scss'],
})
export class SidenavComponent {
  // TODO: get from API
  list = [{ name: 'Recent', icon: 'fiber_new', click: () => {} }];
  @ViewChild('drawer') drawer: ElementRef;
  

  constructor() {}
}
