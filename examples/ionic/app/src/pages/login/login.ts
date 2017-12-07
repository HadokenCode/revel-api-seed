import { Component } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { IonicPage, NavController, ToastController } from 'ionic-angular';
import { HttpHeaders } from '@angular/common/http';



import { Api } from '../../providers/providers';
import { MainPage } from '../pages';

@IonicPage()
@Component({
  selector: 'page-login',
  templateUrl: 'login.html'
})
export class LoginPage {
  // The account fields for the login form.
  // If you're using the username field with or without email, make
  // sure to add it to the type
  account: { email: string, password: string } = {
    email: 'demo@demo.com',
    password: 'demo'
  };

  // Our translated text strings
  private loginErrorString: string;

  constructor(public navCtrl: NavController,
    public toastCtrl: ToastController,
    public translateService: TranslateService,
    private api: Api) {

    this.translateService.get('LOGIN_ERROR').subscribe((value) => {
      this.loginErrorString = value;
    })
  }

  // Attempt to login
  doLogin() {
    this.api.login(this.account.email, this.account.password).subscribe(data => {
      localStorage.setItem('user_id', data.user_id);
      localStorage.setItem('user_name', data.user_name);
      localStorage.setItem('user_email', this.account.email);
      localStorage.setItem('user_token', data.token);
      this.api.headers = new HttpHeaders().set('Authorization', data.token);
      this.navCtrl.push(MainPage);
    }, err => {
      let toast = this.toastCtrl.create({
        message: this.loginErrorString,
        duration: 3000,
        position: 'top'
      });
      toast.present();
    })
  }

}
