import { Injectable } from '@angular/core';
import Auth from 'aws-amplify';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  async signUp(username: string, password: string, email: string) {
    try {
      const user = await Auth.signUp({
        username,
        password,
        attributes: {
          email
        }
      });
      console.log({ user });
    } catch (error) {
      console.log('error signing up:', error);
    }
  }

  async signIn(username: string, password: string) {
    try {
      const user = await Auth.signIn(username, password);
      console.log({ user });
    } catch (error) {
      console.log('error signing in', error);
    }
  }

  async signOut() {
    try {
      await Auth.signOut();
    } catch (error) {
      console.log('error signing out:', error);
    }
  }
}
