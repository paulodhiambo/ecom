package database

import (
	"context"
	"database/sql"
	"ecom/database/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateMerchant(t *testing.T) {
	createRandomMerchant(t)
}

func TestGetMerchant(t *testing.T) {
	merchant1 := createRandomMerchant(t)
	merchant2, err := testQueries.GetMerchant(context.Background(), merchant1.ID)
	require.NoError(t, err)
	require.Equal(t, merchant1.MerchantName, merchant2.MerchantName)
	require.Equal(t, merchant1.CountryCode, merchant2.CountryCode)
	require.Equal(t, merchant1.UpdatedAt, merchant2.UpdatedAt)
	require.Equal(t, merchant1.CreatedAt, merchant2.CreatedAt)
}

func TestGetMerchants(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMerchant(t)
	}
	arg := ListMerchantsParams{
		Limit:  5,
		Offset: 5,
	}
	merchants, err := testQueries.ListMerchants(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, merchants, 5)

	for _, merchant := range merchants {
		require.NotEmpty(t, merchant)
	}
}

func TestDeleteMerchant(t *testing.T) {
	merchant := createRandomMerchant(t)
	err := testQueries.DeleteMerchant(context.Background(), merchant.ID)
	require.NoError(t, err)
	country2, err := testQueries.GetMerchant(context.Background(), merchant.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, country2)
}

func TestUpdateMerchant(t *testing.T) {
	merchant := createRandomMerchant(t)

	arg := UpdateMerchantParams{
		ID:           merchant.ID,
		MerchantName: util.RandomFullName(),
	}
	err := testQueries.UpdateMerchant(context.Background(), arg)
	require.NoError(t, err)
}

func createRandomMerchant(t *testing.T) Merchant {
	country := createRandomCountry(t)
	arg := CreateMerchantParams{
		ID:           util.RandomInt(1, 999999),
		MerchantName: util.RandomString(10),
		CountryCode:  country.Code,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	merchant, err := testQueries.CreateMerchant(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.ID, merchant.ID)
	require.Equal(t, arg.MerchantName, merchant.MerchantName)
	require.Equal(t, arg.CountryCode, merchant.CountryCode)
	return merchant
}
