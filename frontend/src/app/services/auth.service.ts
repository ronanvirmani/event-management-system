import { Injectable } from "@angular/core";
import {
  AuthenticationDetails,
  CognitoUser,
  CognitoUserPool,
} from "amazon-cognito-identity-js";
import { environment } from "../../environments/environment";
import { Router } from "@angular/router";

@Injectable({
  providedIn: "root",
})
export class AuthenticateService {
  userPool: any;
  cognitoUser: any;
  username: string = "";

  constructor(private router: Router) {}

  // Login
  login(emailaddress: any, password: any) {
    let authenticationDetails = new AuthenticationDetails({
      Username: emailaddress,
      Password: password,
    });

    let poolData = {
      UserPoolId: environment.awsAmplifyConfig.Auth.userPoolId,
      ClientId: environment.awsAmplifyConfig.Auth.userPoolWebClientId,
    };

    this.username = emailaddress;
    this.userPool = new CognitoUserPool(poolData);
    let userData = { Username: emailaddress, Pool: this.userPool };
    this.cognitoUser = new CognitoUser(userData);

    this.cognitoUser.authenticateUser(authenticationDetails, {
      onSuccess: (result: any) => {
        this.router.navigate(["/home"]);
        console.log("Success Results : ", result);
      },
      // First time login attempt
      newPasswordRequired: () => {
        this.router.navigate(["/newPasswordRequire"]);
      },
      onFailure: (error: any) => {
        console.log("error", error);
      },
    });
  }
  
}