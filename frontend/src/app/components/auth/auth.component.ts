import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html'
})
export class AuthComponent {
  username: string = '';
  password: string = '';
  email: string = '';
  isLoginMode = true;

  constructor(private authService: AuthService) {}

  toggleMode() {
    this.isLoginMode = !this.isLoginMode;
  }

  onSubmit() {
    if (this.isLoginMode) {
      this.authService.signIn(this.username, this.password);
    } else {
      this.authService.signUp(this.username, this.password, this.email);
    }
  }
}
