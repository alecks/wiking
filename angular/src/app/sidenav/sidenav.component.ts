import { Component } from '@angular/core';
import { MatSnackBarModule, MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.scss']
})
export class SidenavComponent {

  // TODO: get from API
  list = [{ name: "Most Recent", icon: "fiber_new", href: "" }];

  constructor(private snackbar: MatSnackBar) { }

  onClick() {
    this.snackbar.open("This is WIP.");
  }

}
