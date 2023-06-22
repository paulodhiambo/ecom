package database

import (
	"context"
	"database/sql"
	"ecom/database/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestGetOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	order2, err := testQueries.GetOrder(context.Background(), order1.ID)
	require.NoError(t, err)
	require.Equal(t, order1.Status, order2.Status)
	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.UserID, order2.UserID)
}

func TestDeleteOrder(t *testing.T) {
	order := createRandomOrder(t)
	err := testQueries.DeleteOrder(context.Background(), order.ID)
	require.NoError(t, err)
	category2, err := testQueries.GetOrder(context.Background(), order.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestGetOrders(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}
	arg := ListOrdersParams{
		Limit:  5,
		Offset: 5,
	}
	orders, err := testQueries.ListOrders(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orders, 5)

	for _, order := range orders {
		require.NotEmpty(t, order)
	}
}

func TestUpdateOrder(t *testing.T) {
	order := createRandomOrder(t)

	arg := UpdateOrderParams{
		ID:     order.ID,
		Status: util.RandomFullName(),
	}
	err := testQueries.UpdateOrder(context.Background(), arg)
	require.NoError(t, err)
}

func createRandomOrder(t *testing.T) Order {
	user := createRandomUser(t)
	arg := CreateOrderParams{
		ID:        util.RandomInt(1, 99999999),
		UserID:    int32(user.ID),
		Status:    util.RandomString(4),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Status, order.Status)
	require.Equal(t, arg.ID, order.ID)
	require.Equal(t, arg.UserID, order.UserID)
	return order
}
