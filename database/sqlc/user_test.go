package database

import (
	"context"
	"database/sql"
	"ecom/database/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Gender, user2.Gender)
	require.Equal(t, user1.CountryCode, user2.CountryCode)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestGetUsers(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomUser(t)
	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, account := range users {
		require.NotEmpty(t, account)
	}
}

func TestDeleteUser(t *testing.T) {
	account1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetUser(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestUpdateUser(t *testing.T) {
	account1 := createRandomUser(t)

	arg := UpdateUsersParams{
		ID:       account1.ID,
		FullName: util.RandomFullName(),
	}
	err := testQueries.UpdateUsers(context.Background(), arg)
	require.NoError(t, err)
}

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		ID:          util.RandomInt(1, 2000000),
		Email:       util.RandomEmail(),
		Gender:      util.RandomGender(),
		DateOfBirth: time.Date(2000, 8, 15, 14, 30, 45, 100, time.Local),
		CreatedAt:   time.Now(),
		CountryCode: "KE",
		FullName:    util.RandomFullName(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.Gender)
	require.NotEmpty(t, user.DateOfBirth)
	require.NotEmpty(t, user.CountryCode)
	require.NotEmpty(t, user.FullName)
	return user
}
