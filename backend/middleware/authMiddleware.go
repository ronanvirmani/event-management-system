package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Missing Authorization header"))
            return
        }

        token := strings.Replace(authHeader, "Bearer ", "", 1)
        if token == "" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Invalid token"))
            return
        }

        // Validate token with Cognito
        region := aws.String(os.Getenv("AWS_REGION"))
        // poolID := aws.String(os.Getenv("AWS_COGNITO_USER_POOL_ID"))

        sess := session.Must(session.NewSession(&aws.Config{
            Region: region,
        }))

        svc := cognitoidentityprovider.New(sess)
        input := &cognitoidentityprovider.GetUserInput{
            AccessToken: aws.String(token),
        }

        _, err := svc.GetUser(input)
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Invalid token"))
            return
        }

        next.ServeHTTP(w, r)
    })
}
