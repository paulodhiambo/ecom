package database

import (
	"context"
	"database/sql"
	"ecom/database/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateOrderItem(t *testing.T) {
	createRandomOrderItem(t)
}

func TestGetOrderItem(t *testing.T) {
	orderItem1 := createRandomOrderItem(t)
	orderItem2, err := testQueries.GetOrderItem(context.Background(), orderItem1.OrderID)
	require.NoError(t, err)
	require.Equal(t, orderItem1.OrderID, orderItem2.OrderID)
	require.Equal(t, orderItem1.ProductID, orderItem2.ProductID)
	require.Equal(t, orderItem1.Quantity, orderItem2.Quantity)
}

func TestDeleteOrderItem(t *testing.T) {
	order := createRandomOrderItem(t)
	err := testQueries.DeleteOrderItem(context.Background(), order.OrderID)
	require.NoError(t, err)
	category2, err := testQueries.GetOrderItem(context.Background(), order.OrderID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestGetOrdersItem(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrderItem(t)
	}
	arg := ListOrderItemsParams{
		Limit:  5,
		Offset: 5,
	}
	ordersItems, err := testQueries.ListOrderItems(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, ordersItems, 5)

	for _, orderItem := range ordersItems {
		require.NotEmpty(t, orderItem)
	}
}

func TestUpdateOrderItem(t *testing.T) {
	order := createRandomOrderItem(t)

	arg := UpdateOrderItemParams{
		OrderID:  order.OrderID,
		Quantity: int32(util.RandomInt(19, 99999)),
	}
	err := testQueries.UpdateOrderItem(context.Background(), arg)
	require.NoError(t, err)
}

func createRandomOrderItem(t *testing.T) OrderItem {
	order := createRandomOrder(t)
	product := createRandomProduct(t)
	arg := CreateOrderItemParams{
		OrderID:   order.ID,
		ProductID: product.ID,
		Quantity:  int32(util.RandomInt(0, 999)),
	}
	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.OrderID, orderItem.OrderID)
	require.Equal(t, arg.Quantity, orderItem.Quantity)
	require.Equal(t, arg.ProductID, orderItem.ProductID)
	return orderItem
}
