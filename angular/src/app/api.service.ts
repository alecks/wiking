import { Injectable } from '@angular/core';
import { AppConfiguration } from 'src/configuration';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private path: string = AppConfiguration.siteURL + AppConfiguration.apiPath;

  constructor(private http: HttpClient) {}

  async authorize(username: string, password: string): Promise<any> {
    return this.http
      .post(
        this.path + 'auth/authorize',
        {
          username: username.trim(),
          password,
        },
        {
          responseType: 'text',
        }
      )
      .toPromise();
  }

  async signup(
    username: string,
    password: string,
    permitUsername?: string,
    permitPassword?: string,
    superuser?: string
  ): Promise<any> {
    return this.http
      .post(
        this.path + 'auth/user',
        {
          username: username.trim(),
          password,
          permitUsername,
          permitPassword,
          superuser,
        },
        {
          responseType: 'text',
        }
      )
      .toPromise();
  }
}
