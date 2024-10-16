import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from '@angular/common/http';
import { Observable, from, throwError } from 'rxjs';
import { switchMap, catchError } from 'rxjs/operators';
import { fetchAuthSession } from 'aws-amplify/auth';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return from(fetchAuthSession()).pipe(
      switchMap(session => {
        // Check if the session has valid tokens
        
        if (session.tokens && session.tokens.accessToken) {
          const token = session.tokens.accessToken;

          // Check if the token is valid based on expiration
          if (this.isTokenValid(token.toString())) {
            const authReq = req.clone({
              setHeaders: {
                Authorization: `Bearer ${token.toString()}` // Convert token to string
              }
            });
            return next.handle(authReq);
          } else {
            // Token is expired or invalid, handle accordingly
            console.warn('Access token is invalid or expired.');
            return next.handle(req); // Proceed without attaching the token
          }

        }
        // If no tokens, proceed without attaching the header
        return next.handle(req);
      }),
      catchError(error => {
        console.error('Error fetching auth session:', error);
        // Optionally, redirect to login or handle the error as needed
        return throwError(error);
      })
    );
  }

  /**
   * Helper function to check if the JWT token is valid based on its expiration time.
   * @param token - The JWT access token.
   * @returns True if the token is still valid, false otherwise.
   */
  private isTokenValid(token: string): boolean {
    try {
      const payload = this.decodeJwt(token);
      const currentTime = Math.floor(Date.now() / 1000);
      return payload.exp && payload.exp > currentTime;
    } catch (error) {
      console.error('Error decoding token:', error);
      return false;
    }
  }

  /**
   * Decodes a JWT token and returns its payload.
   * @param token - The JWT token to decode.
   * @returns The payload of the JWT token.
   */
  private decodeJwt(token: string): any {
    const payload = token.split('.')[1];
    const decoded = atob(payload);
    return JSON.parse(decoded);
  }
}
