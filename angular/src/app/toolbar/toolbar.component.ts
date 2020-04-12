import { Component, Input } from '@angular/core';
import { AppConfiguration } from 'src/configuration';
import { MatDialog } from '@angular/material/dialog';
import { LoginComponent } from '../login/login.component';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss'],
})
export class ToolbarComponent {

  @Input() title: string;
  config = AppConfiguration;
  username = '';
  password = '';
  error = '';

  // This is static; no need to get from api. Might implement customisation.
  buttons = [
    { name: 'Login', icon: 'account_circle', click: 'openLogin()' },
    { name: 'GitHub', icon: 'code', href: AppConfiguration.githubLink },
  ];

  constructor(private dialog: MatDialog, private snackbar: MatSnackBar) {}

  openLogin(): void {
    const dialogRef = this.dialog.open(LoginComponent, {
      width: '25rem',
      data: { username: this.username, password: this.password, error: this.error },
    });
    dialogRef.componentInstance.snackbarEvent.subscribe(res => {
      dialogRef.componentInstance.loader = false;
      this.snackbar.open(res, "Close");
    });
  }

}
