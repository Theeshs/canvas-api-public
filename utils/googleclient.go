package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func tokenFromFile(file_ string) (*oauth2.Token, error) {
	file, err := os.Open(file_)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	return token, err
}

func getTokenFromWeb(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	authUrl := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authUrl)
	var authCode string
	fmt.Print("Enter the authorization code: ")
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Printf("Unable to read authorization code: %+v", err)
		return nil, err
	}

	tok, err := config.Exchange(ctx, authCode)
	if err != nil {
		log.Printf("Unable to retrieve token from web: %+v", err)
		return nil, err
	}

	return tok, nil
}

func saveToken(path string, token *oauth2.Token) error {
	log.Printf("Saving credential file to: %s\n", path)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Printf("Unable to cache oauth token: %+v", err)
		return err
	}
	defer file.Close()

	json.NewEncoder(file).Encode(token)
	return nil
}

func getClient(ctx context.Context, config *oauth2.Config) (*http.Client, error) {
	tokenFile := "token.json"
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		if tok, err = getTokenFromWeb(ctx, config); err != nil {
			return nil, err
		}
		if err = saveToken(tokenFile, tok); err != nil {
			return nil, err
		}
	}
	return config.Client(ctx, tok), nil
}

func (s *AuthService) GetGoogleClient(ctx context.Context, credentialsPath string) (*http.Client, error) {
	b, err := os.ReadFile(credentialsPath)
	if err != nil {
		log.Printf("Unable to read credentials file '%s': %+v", credentialsPath, err)
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Printf("Unable to parse credentials file to config: %+v", err)
		return nil, err
	}

	client, err := getClient(ctx, config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
