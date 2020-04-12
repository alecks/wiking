import { Component, Inject, Output } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { ApiService } from '../api.service';
import { EventEmitter } from '@angular/core';

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
  public snackbarEvent = new EventEmitter<string>();

  constructor(
    private api: ApiService,
    public dialogRef: MatDialogRef<LoginComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData
  ) {}

  onLogin() {
    this.loader = true;

    this.api
      .authorize(this.data.username, this.data.password)
      .then((res) => this.snackbarEvent.emit(res))
      .catch((e) => this.snackbarEvent.emit(e.message));
  }

  onSignup() {
    this.loader = true;

    this.api
      .signup(this.data.username, this.data.password)
      .then((res) => this.snackbarEvent.emit(res))
      .catch((e) => this.snackbarEvent.emit(e.message));
  }
}
