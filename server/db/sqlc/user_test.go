package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	util "github.com/3iOj/OnlineJudge/utils"
	"github.com/stretchr/testify/require"
)
func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(8));
	require.NoError(t, err);
	arg := CreateUserParams{
		Username:       util.RandomUser(),
		Password: 		hashedPassword,
		Name:       	util.RandomUser(),
		Email:       	util.RandomEmail(),
		Dob:           	time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC),
	
	}
	
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Dob, user.Dob)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}