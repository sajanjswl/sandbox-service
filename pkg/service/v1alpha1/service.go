package v1alpha1

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/sajanjswl/sandbox-service/config"
	v1 "github.com/sajanjswl/sandbox-service/gen/go/sandbox/v1alpha1"
	"github.com/sajanjswl/sandbox-service/models"

	"go.uber.org/zap"
)

type sandboxServiceServer struct {
	db     *gorm.DB
	logger *zap.Logger
	// v1.UnimplementedsandboxServiceServer
	config *config.Config
}

// register db wiht server
func NewSandboxServiceServer(db *gorm.DB, logger *zap.Logger, cfg *config.Config) v1.SandboxServiceServer {
	return &sandboxServiceServer{
		db:     db,
		logger: logger,
		config: cfg,
	}
}

func (s *sandboxServiceServer) LoginUser(ctx context.Context, req *v1.LoginUserRequest) (*v1.LoginUserResponse, error) {

	user := &models.User{}
	if err := models.GetUser(s.db, user, req.GetEmailId()); err != nil {
		s.logger.Warn("error reading user from datatbase", zap.String("email", user.Email), zap.Error(err))
		return nil, status.Error(codes.NotFound, "user not found")
	}

	//authenticating password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		s.logger.Info("password incorrect", zap.String("email", user.Email), zap.Error(err))
		return nil, errors.New("Password incorrect")
	}

	response := &v1.LoginUserResponse{
		Status:  "200",
		Message: "Hello  " + user.Name + "Logged in Successfully!!",
	}
	return response, nil

}

func (s *sandboxServiceServer) RegisterUser(ctx context.Context, req *v1.RegisterUserRequest) (*v1.RegisterUserResponse, error) {

	user := &models.User{}
	if err := models.GetUser(s.db, user, req.GetUser().EmailId); err == nil {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	//bycrpting the plaint text password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.GetUser().GetPassword()), bcrypt.MinCost)
	if err != nil {
		s.logger.Warn("Failed to hash password", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	user.Email = req.GetUser().EmailId
	user.Password = string(passwordHash)
	user.Name = req.GetUser().Name
	user.Mobile = req.GetUser().MobileNumber

	if err := models.CreateUser(s.db, user); err != nil {
		s.logger.Warn("Failed to register user", zap.String("email", user.Email), zap.Error(err))
	}

	s.logger.Info("register user", zap.String("email", user.Email))

	return &v1.RegisterUserResponse{Message: "successfully registerd  " + req.GetUser().GetName() + "  " + req.GetUser().EmailId}, nil
}
