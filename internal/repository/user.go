package repository

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/saleh-ghazimoradi/GopherInn/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserById(ctx context.Context, id string) (*domain.User, error)
	GetUsers(ctx context.Context, offset, limit int) ([]*domain.User, error)
	CountUsers(ctx context.Context) (int64, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userRepository struct {
	collection *mongo.Collection
	logger     *zerolog.Logger
}

func (u *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	if user.Id.IsZero() {
		user.Id = bson.NewObjectID()
	}

	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		u.logger.Error().Err(err).Msg("failed to create user")
		return err
	}

	return nil
}

func (u *userRepository) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user domain.User
	err = u.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetUsers(ctx context.Context, offset, limit int) ([]*domain.User, error) {
	opts := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit))

	cursor, err := u.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*domain.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) CountUsers(ctx context.Context) (int64, error) {
	return u.collection.CountDocuments(ctx, bson.M{})
}

func (u *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	update := bson.M{
		"$set": bson.M{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		},
	}

	result, err := u.collection.UpdateOne(
		ctx,
		bson.M{"_id": user.Id},
		update,
	)
	if err != nil {
		u.logger.Err(err).Msg("failed to update user")
		return err
	}

	if result.MatchedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id string) error {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := u.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		u.logger.Err(err).Msg("failed to delete user")
		return err
	}

	if result.DeletedCount == 0 {
		return ErrNotFound
	}
	return nil
}

func NewUserRepository(database *mongo.Database, collectionName string, logger *zerolog.Logger) UserRepository {
	return &userRepository{
		collection: database.Collection(collectionName),
		logger:     logger,
	}
}
