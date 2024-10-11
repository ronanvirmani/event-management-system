package services

import (
    "errors"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
    "github.com/ronanvirmani/event-management-system/backend/models"
)

var svc *cognitoidentityprovider.CognitoIdentityProvider

func init() {
    region := os.Getenv("AWS_REGION")
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String(region),
    }))
    svc = cognitoidentityprovider.New(sess)
}

func SignUpUser(user models.User) error {
    input := &cognitoidentityprovider.SignUpInput{
        ClientId: aws.String(os.Getenv("AWS_COGNITO_APP_CLIENT_ID")),
        Username: aws.String(user.Username),
        Password: aws.String(user.Password),
        UserAttributes: []*cognitoidentityprovider.AttributeType{
            {
                Name:  aws.String("email"),
                Value: aws.String(user.Email),
            },
        },
    }

    _, err := svc.SignUp(input)
    if err != nil {
        return err
    }

    return nil
}

func LoginUser(user models.User) (string, error) {
    input := &cognitoidentityprovider.InitiateAuthInput{
        AuthFlow: aws.String("USER_PASSWORD_AUTH"),
        ClientId: aws.String(os.Getenv("AWS_COGNITO_APP_CLIENT_ID")),
        AuthParameters: map[string]*string{
            "USERNAME": aws.String(user.Username),
            "PASSWORD": aws.String(user.Password),
        },
    }

    result, err := svc.InitiateAuth(input)
    if err != nil {
        return "", err
    }

    if result.AuthenticationResult == nil {
        return "", errors.New("authentication failed")
    }

    return *result.AuthenticationResult.AccessToken, nil
}
