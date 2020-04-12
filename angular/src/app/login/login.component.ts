import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AppConfiguration } from 'src/configuration';
import { ApiService } from '../api.service';

export interface DialogData {
  username: string;
  password: string;
}

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent {
  loader = false;
  error = '';

  constructor(
    private api: ApiService,
    public dialogRef: MatDialogRef<LoginComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
  ) {}

  onLogin() {
    this.loader = true;

    this.api.authorize(this.data.username, this.data.password).catch(e => this.error = e.message).then(res => {
      this.error = res;
      this.loader = false;
    });
  }

  onSignup() {
    this.loader = true;

    this.api.signup(this.data.username, this.data.password).catch(e => this.error = e.message).then(res => {
      this.error = res;
      this.loader = false;
    });
  }
}
