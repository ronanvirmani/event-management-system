import { Component } from '@angular/core';
import { AuthenticateService } from '../../services/auth.service';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html'
})
export class AuthComponent {
  username: string = '';
  password: string = '';
  email: string = '';
  isLoginMode = true;

  constructor(private authService: AuthenticateService) {}

  toggleMode() {
    this.isLoginMode = !this.isLoginMode;
  }

  onSubmit() {
    if (this.isLoginMode) {
      this.authService.login(this.username, this.password);
    }
  }
}
