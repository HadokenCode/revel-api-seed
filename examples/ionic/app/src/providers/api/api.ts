import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

interface LoginResponse {
  user_id: string;
  user_name: string;
  user_email: string;
  token: string;
}

/**
 * Api is a generic REST Api handler. Set your API url first.
 */
@Injectable()
export class Api {
  url: string = 'http://localhost:9000';
  token: string;
  headers: HttpHeaders = new HttpHeaders();

  constructor(public http: HttpClient) {
  }

  login(email: string, password: string) {
    return this.http.post<LoginResponse>(this.url + '/login', { email: email, password: password });
  }

  register(email: string, password: string, name: string) {
    return this.http.post<LoginResponse>(this.url + '/register', { email: email, password: password, name: name });
  }

  logout() {
    localStorage.removeItem('user_id');
    localStorage.removeItem('user_name');
    localStorage.removeItem('user_email');
    localStorage.removeItem('user_token');
    this.headers = new HttpHeaders();
  }

  get(endpoint: string, params?: any) {
    let reqOpts = {
      params: new HttpParams(),
      headers: this.headers
    };

    // Support easy query params for GET requests
    if (params) {
      reqOpts.params = new HttpParams();
      for (let k in params) {
        reqOpts.params.set(k, params[k]);
      }
    }

    return this.http.get(this.url + '/' + endpoint, reqOpts);
  }

  post(endpoint: string, body: any) {
    return this.http.post(this.url + '/' + endpoint, body, { headers: this.headers });
  }

  put(endpoint: string, body: any) {
    return this.http.put(this.url + '/' + endpoint, body, { headers: this.headers });
  }

  delete(endpoint: string) {
    return this.http.delete(this.url + '/' + endpoint, { headers: this.headers });
  }

  patch(endpoint: string, body: any) {
    return this.http.put(this.url + '/' + endpoint, body, { headers: this.headers });
  }
}
