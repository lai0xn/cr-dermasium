package auth

import (
	"errors"
	"log"

	"github.com/lai0xn/cr-dermasuim/app/utils"
	"github.com/lai0xn/cr-dermasuim/models"
	"github.com/lai0xn/cr-dermasuim/storage"
	"github.com/satori/go.uuid"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Signup(user SignUpPayload) error {
	userModel := &models.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		IsActive:    user.IsActive,
	}
	db := storage.DB.Create(userModel)
	if db.Error != nil {
		return db.Error
	}
	s.SendOTP(user.PhoneNumber)
	return nil
}

func (s *Service) SendOTP(to string) error {
	params := &openapi.CreateVerificationParams{}
	log.Println("cc")
	params.SetTo(to)
	params.SetChannel("sms")

	_, err := client.VerifyV2.CreateVerification(verifyServiceSid, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Refresh(token string) (string, error) {
	token, err := utils.RefreshToken(token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Service) Signin(payload LoginPayload) (*models.User, error) {
	var user models.User
	// if the user already exists there is no need of creating a new one
	storage.DB.Where("phone_number = ?", payload.PhoneNumber).First(&user)
	if user.PhoneNumber != "" {
		err := s.SendOTP(user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	// create the user and send otp
	db := storage.DB.Create(&models.User{
		PhoneNumber: payload.PhoneNumber,
	})
	if db.Error != nil {
		return nil, db.Error
	}
	err := s.SendOTP(user.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) VerifyOTP(userID string, code string) (*models.User, error) {
	var user models.User
	id, err := uuid.FromString(userID)
	if err != nil {
		return nil, err
	}

	db := storage.DB.Where("id = ?", id).First(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(user.PhoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(verifyServiceSid, params)
	if err != nil {
		return nil, err
	} else if *resp.Status == "approved" {
		// if this is the first time the user verifies , activate his account
		if user.IsActive == false {
			log.Println("cc yakho")
			db := storage.DB.Model(&user).Updates(map[string]interface{}{"is_active": true})
			if db.Error != nil {
				log.Println(db.Error)
			}
		}
		return &user, nil
	} else {
		return nil, errors.New("invalid code")
	}
}
