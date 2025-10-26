package user

import (
	"api/ent"
	"api/ent/document"
	"api/ent/user"
	"api/utils"
	"context"
	"fmt"
	"mime/multipart"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type DocumentType string

const (
	DocumentTypeResume      DocumentType = "resume"
	DocumentTypePassport    DocumentType = "passport"
	DocumentTypeIDCard      DocumentType = "id_card"
	DocumentTypeCertificate DocumentType = "certificate"
	DocumentTypeOther       DocumentType = "other"
)

type UserHandler struct {
	client *ent.Client
}

func NewUserHandler(client *ent.Client) *UserHandler {
	return &UserHandler{
		client: client,
	}
}

func (uh *UserHandler) CheckUserAvailability(email string) bool {
	user, _ := uh.client.User.Query().Where(user.Email(email)).Only(context.Background())
	return user != nil
}

func (uh *UserHandler) CreateUser(email string, password string, username string) (*ent.User, error) {
	tx, err := uh.client.Tx(context.Background())
	if err != nil {
		return nil, err
	}
	newUser, err := tx.User.Create().
		SetEmail(email).
		SetPassword(password).SetUsername(username).Save(context.Background())

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return newUser, err
}

func (uh *UserHandler) UpdateUsers(id uint, u User) (*ent.User, error) {

	dob, err := utils.ConvertJsonDate(*u.DOB)

	if err != nil {
		return nil, err
	}
	user, err := uh.client.User.UpdateOneID(uint(id)).
		SetUsername(u.UserName).
		SetFirstName(*u.FirstName).
		SetLastName(*u.LastName).
		SetDob(dob).
		SetDescription(*u.Description).
		SetMobileNumber(int32(*u.MobileNumber)).
		SetAddressBlock(*u.AddresBlock).
		SetAddressStreet(*u.AddressStreet).
		SetRecidentialCountry(*u.ResidentialCountry).
		SetNationality(*u.Nationality).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (uh *UserHandler) FetchUserByID(id uint) (*ent.User, error) {
	return uh.client.User.Get(context.Background(), id)
}

func (uh *UserHandler) GenUploadToDrive(fileName string, file multipart.File, userId uint, documentType document.DocumentType) (string, error) {
	ctx := context.Background()
	tx, err := uh.client.Tx(context.Background())

	user, err := uh.FetchUserByID(userId)
	if err != nil {
		return "", fmt.Errorf("Failed to find a user: %v", err)
	}

	// Get authenticated HTTP client
	authSrv := utils.AuthService{}
	client, err := authSrv.GetGoogleClient(ctx, "/Users/theesh/dev/backends/Theesh/api/services/user/client_secret_181868859823-oj2chmhn3sjae4hu8fsouvtm49f3r3gf.apps.googleusercontent.com.json")
	if err != nil {
		return "", fmt.Errorf("failed to get Google client: %v", err)
	}

	// Create Google Drive service
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return "", fmt.Errorf("unable to create Drive service: %v", err)
	}

	// Create a Google Drive file metadata
	f := &drive.File{
		Name: fileName,
	}

	// Upload file
	// driveFile, err := srv.Files.Create(f).Media(file).Do()
	driveFile, err := srv.Files.
		Create(f).
		Media(file).
		Fields("*").
		Do()
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	fmt.Printf("File '%s' uploaded successfully with ID: %s\n", driveFile.Name, driveFile.Id)
	// driveFile
	_, err = uh.client.Document.Create().
		SetDocumentName(fileName).
		SetGoogleID(driveFile.Id).
		SetUser(user).
		SetDocumentType(documentType).
		SetDocumentWebLink(driveFile.WebViewLink).
		SetDocumentThumnailLink(driveFile.ThumbnailLink).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	if err := tx.Commit(); err != nil {
		return "", err
	}
	return driveFile.Id, nil
}
