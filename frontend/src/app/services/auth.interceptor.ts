import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from '@angular/common/http';
import Auth from '@aws-amplify/auth';
import { CognitoUserSession } from 'amazon-cognito-identity-js';
import { Observable, from } from 'rxjs';
import { switchMap, catchError } from 'rxjs/operators';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return from(Auth.currentSession()).pipe(
      switchMap((session: CognitoUserSession) => {
        const token = session.getAccessToken().getJwtToken();
        const authReq = req.clone({
          headers: req.headers.set('Authorization', `Bearer ${token}`)
        });
        return next.handle(authReq);
      }),
      catchError(() => {
        return next.handle(req);
      })
    );
  }
}
