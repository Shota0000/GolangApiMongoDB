package data

import (
	"context"
	"projectName/pkg/config"
	"projectName/pkg/domain"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserProvider interface {
	CreateAccount(user *domain.User) error
	UsernameExists(username string) (bool, error)
	FindByUsername(username string) (*domain.User, error)
}

type UserProvider struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserProvider(cfg *config.Settings, mongo *mongo.Client) IUserProvider {
	userCollection := mongo.Database(cfg.DbName).Collection("users")
	return &UserProvider{
		userCollection: userCollection,
		//ctx作らなきゃダメだからとりま使うときに使うのが, context.TODO()
		ctx: context.TODO(),
	}
}

func (u UserProvider) CreateAccount(user *domain.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	if err != nil {
		return errors.Wrap(err, "Error inserting user")
	}
	return nil
}

func (u UserProvider) UsernameExists(username string) (bool, error) {
	var userFound *domain.User
	//filterはスライスなので複数の条件を入れれる
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	//Findoneでfilterの条件に合致するドキュメントをとってきて,Decodeで構造体に変換しいれる
	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrap(err, "Error finding by username")
	}
	return true, nil
}

func (u UserProvider) FindByUsername(username string) (*domain.User, error) {
	var userFound domain.User
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error finding by username")
	}
	return &userFound, nil
}
