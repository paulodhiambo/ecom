package database

import (
	"context"
	"database/sql"
	"ecom/database/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.Equal(t, category1.CatName, category2.CatName)
	require.Equal(t, category1.ID, category2.ID)
}

func TestDeleteCategory(t *testing.T) {
	category := createRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category.ID)
	require.NoError(t, err)
	category2, err := testQueries.GetCategory(context.Background(), category.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestGetCategories(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomCategory(t)
	}
	arg := ListCategoriesParams{
		Limit:  5,
		Offset: 5,
	}
	categories, err := testQueries.ListCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, 5)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}

func TestUpdateCategory(t *testing.T) {
	account1 := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:      account1.ID,
		CatName: util.RandomFullName(),
	}
	err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
}

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		CatName: util.RandomString(6),
		ID:      util.RandomInt(1, 99999999),
	}
	cat, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.CatName, cat.CatName)
	require.Equal(t, arg.ID, cat.ID)
	return cat
}
