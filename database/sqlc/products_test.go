package database

import (
	"context"
	"database/sql"
	"ecom/database/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.NoError(t, err)
	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.MerchantID, product2.MerchantID)
	require.Equal(t, product1.Status, product2.Status)
}

func TestGetProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}
	arg := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}
	products, err := testQueries.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}

func TestDeleteProduct(t *testing.T) {
	product := createRandomProduct(t)
	err := testQueries.DeleteProduct(context.Background(), product.ID)
	require.NoError(t, err)
	product2, err := testQueries.GetProduct(context.Background(), product.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}

func TestUpdateProduct(t *testing.T) {
	arg := UpdateProductParams{
		Name:   util.RandomFullName(),
		Price:  int32(util.RandomInt(100, 5000)),
		Status: util.RandomString(4),
	}
	err := testQueries.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)
}

func createRandomProduct(t *testing.T) Product {
	merchant := createRandomMerchant(t)
	category := createRandomCategory(t)
	arg := CreateProductParams{
		Name:       util.RandomString(6),
		ID:         util.RandomInt(0, 999999999),
		MerchantID: int32(merchant.ID),
		Price:      int32(util.RandomInt(100, 8000)),
		Status:     util.RandomString(5),
		CategoryID: int32(category.ID),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.MerchantID, product.MerchantID)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.Status, product.Status)
	require.Equal(t, arg.CategoryID, product.CategoryID)
	require.Equal(t, arg.ID, product.ID)
	return product
}
