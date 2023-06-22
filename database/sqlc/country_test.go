package database

import (
	"context"
	"database/sql"
	"ecom/database/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateCountry(t *testing.T) {
	createRandomCountry(t)
}

func TestGetCountry(t *testing.T) {
	country1 := createRandomCountry(t)
	country2, err := testQueries.GetCountry(context.Background(), country1.Code)
	require.NoError(t, err)
	require.Equal(t, country1.Name, country2.Name)
	require.Equal(t, country1.Code, country2.Code)
	require.Equal(t, country1.ContinentName, country2.ContinentName)
}

func TestGetCountries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCountry(t)
	}
	arg := ListCountriesParams{
		Limit:  5,
		Offset: 5,
	}
	countries, err := testQueries.ListCountries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, countries, 5)

	for _, country := range countries {
		require.NotEmpty(t, country)
	}
}

func TestDeleteCountry(t *testing.T) {
	country := createRandomCountry(t)
	err := testQueries.DeleteCountry(context.Background(), country.Code)
	require.NoError(t, err)
	country2, err := testQueries.GetCountry(context.Background(), country.Code)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, country2)
}

func TestUpdateCountry(t *testing.T) {
	country := createRandomCountry(t)

	arg := UpdateCountryParams{
		Code: country.Code,
		Name: util.RandomFullName(),
	}
	err := testQueries.UpdateCountry(context.Background(), arg)
	require.NoError(t, err)
}

func createRandomCountry(t *testing.T) Country {
	arg := CreateCountryParams{
		Code:          util.RandomString(6),
		Name:          util.RandomString(6),
		ContinentName: util.RandomString(5),
	}
	country, err := testQueries.CreateCountry(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Code, country.Code)
	require.Equal(t, arg.ContinentName, country.ContinentName)
	require.Equal(t, arg.Name, country.Name)
	return country
}
