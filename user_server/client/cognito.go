package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoClient interface {
	LogIn(userName string, password string) (error, string, *cognito.InitiateAuthOutput)
	SignUp(userName, email, password string) (error, string)
	ConfirmSignUp(userName, code string) (error, string)
}

type awsCognitoClient struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientId   string
}

func NewCognitoClient(cognitoRegion string, cognitoAppClientID string) CognitoClient {
	conf := &aws.Config{Region: aws.String(cognitoRegion)}
	sess, err := session.NewSession(conf)
	client := cognito.New(sess)

	if err != nil {
		panic(err)
	}

	return &awsCognitoClient{
		cognitoClient: client,
		appClientId:   cognitoAppClientID,
	}
}

func (a *awsCognitoClient) LogIn(userName, password string) (error, string, *cognito.InitiateAuthOutput) {
	result, err := a.cognitoClient.InitiateAuth(&cognito.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": userName,
			"PASSWORD": password,
		}),
		ClientId: aws.String(a.appClientId),
	})
	if err != nil {
		return err, "", nil
	}
	return nil, result.String(), result
}

func (a *awsCognitoClient) SignUp(userName, email, password string) (error, string) {
	user := &cognito.SignUpInput{
		Username: aws.String(userName),
		Password: aws.String(password),
		ClientId: aws.String(a.appClientId),
		UserAttributes: []*cognito.AttributeType{
			{Name: aws.String("email"),
				Value: aws.String(email),
			},
		},
	}
	result, err := a.cognitoClient.SignUp(user)
	if err != nil {
		return err, ""
	}
	return nil, result.String()
}

func (a *awsCognitoClient) ConfirmSignUp(userName, code string) (error, string) {
	result, err := a.cognitoClient.ConfirmSignUp(&cognito.ConfirmSignUpInput{
		Username:         aws.String(userName),
		ConfirmationCode: aws.String(code),
		ClientId:         aws.String(a.appClientId),
	})
	if err != nil {
		return err, ""
	}
	return nil, result.String()
}
