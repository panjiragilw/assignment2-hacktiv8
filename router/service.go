package router

import (
	"assignment2-v4/db"
	"log"
)

// Fungsi untuk membuat data order dengan memasukkan data order dan item ke dalam database.
// Data item bisa banyak (berbentuk []).
// Contoh request body:
//
// {
//     "orderedAt": "2020-10-23",
//     "customerName": "Tom Jerry",
//     "item": [
//         {
//             "itemCode": "125",
//             "description": "iPhone 11",
//             "quantity": 2
//         },
//         {
//             "itemCode": "126",
//             "description": "iPhone 11 Pro",
//             "quantity": 1
//         }
//     ]
// }
func CreateOrderDB(orderReq Order) Order {
	db := db.GetDB()

	queryOrder := `
	INSERT INTO orders (customer_name, ordered_at) 
	VALUES ($1, $2) 
	RETURNING order_id, customer_name, ordered_at`

	row := db.QueryRow(queryOrder, orderReq.CustomerName, orderReq.OrderedAt)

	var orderRes Order

	err := row.Scan(&orderRes.OrderId, &orderRes.CustomerName, &orderRes.OrderedAt)

	if err != nil {
		log.Fatal(err)
	}

	queryItem := `
	INSERT INTO items (item_code, description, quantity, order_id)
	VALUES ($1, $2, $3, $4)
	RETURNING item_id, item_code, description, quantity, order_id
	`
	//dibuat untuk bisa insert banyak data jika diperlukan
	itemReq := orderReq.Item

	var itemRes []Item
	for _, v := range itemReq {
		row = db.QueryRow(queryItem, v.ItemCode, v.Description, v.Quantity, orderRes.OrderId)

		var itemOrder Item
		err = row.Scan(&itemOrder.ItemId, &itemOrder.ItemCode, &itemOrder.Description, &itemOrder.Quantity, &itemOrder.OrderId)

		if err != nil {
			log.Fatal(err)
		}

		itemRes = append(itemRes, itemOrder)
	}

	orderReq.Item = itemRes
	orderReq.OrderId = orderRes.OrderId

	return orderReq
}

// Fungsi untuk membaca semua data order dari database.
func ReadAllOrderDB() []Order {
	db := db.GetDB()

	query := `
	SELECT o.order_id, o.customer_name, o.ordered_at, i.item_id, i.item_code, i.description, i.quantity
	FROM orders o
	JOIN items i
	ON o.order_id = i.order_id`

	row, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	var orderRes []Order
	var itemRes []Item

	var tempOrder Order
	_ = tempOrder

	for row.Next() {
		var order Order
		var item Item

		if err := row.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &item.ItemId, &item.ItemCode, &item.Description, &item.Quantity); err != nil {
			log.Fatal(err)
		}
		item.OrderId = order.OrderId
		itemRes = append(itemRes, item)

		if orderRes == nil {
			orderRes = append(orderRes, order)
			tempOrder = order
		} else {

			if order.OrderId != tempOrder.OrderId {
				orderRes = append(orderRes, order)
				tempOrder = order
			}
		}
	}

	for i, v := range orderRes {
		var itemTemp []Item
		for _, w := range itemRes {
			if w.OrderId == v.OrderId {
				itemTemp = append(itemTemp, w)
			}
		}
		orderRes[i].Item = itemTemp
	}

	return orderRes
}

// Fungsi untuk membaca data berdasarkan orderID yang dimasukkan ke parameter
func ReadOrderDB(orderId int) Order {
	db := db.GetDB()

	query := `
	SELECT o.order_id, o.customer_name, o.ordered_at, i.item_id, i.item_code, i.description, i.quantity
	FROM orders o
	JOIN items i
	ON o.order_id = i.order_id
	WHERE o.order_id = $1`

	row, err := db.Query(query, orderId)
	if err != nil {
		log.Fatal(err)
	}

	var orderRes Order
	var itemRes []Item

	for row.Next() {
		var item Item
		if err := row.Scan(&orderRes.OrderId, &orderRes.CustomerName, &orderRes.OrderedAt, &item.ItemId, &item.ItemCode, &item.Description, &item.Quantity); err != nil {
			log.Fatal(err)
		}
		item.OrderId = orderRes.OrderId
		itemRes = append(itemRes, item)
	}

	orderRes.Item = itemRes

	return orderRes
}

// Fungsi untuk memperbarui data order dan item berdasarkan orderID dan berdasarkan itemID yang dimasukkan.
// Beberapa data item dapat diperbarui menyesuaikan request Body.
// Crashed jika request update Item lebih dari data Item yang ada di database (jika dibuat solusi, akan terlalu panjang).
// Contoh request Body:
//
// {
//     "orderedAt": "2019-11-09T21:21:46+07:00",
//     "customerName": "Spike Tyke",
//     "item": [
//         {
//             "itemId": 1,
//             "itemCode": "123",
//             "description": "iPhone 10X",
//             "quantity": 5
//         },
//         {
//             "itemId": 2,
//             "itemCode": "124",
//             "description": "iPhone 10Xs",
//             "quantity": 10
//         }
//     ]
// }
func UpdateOrderDB(orderReq Order, orderId int) Order {
	db := db.GetDB()

	queryOrder := `
	UPDATE orders
	SET customer_name = $2, ordered_at = $3
	WHERE order_id = $1
	RETURNING order_id, customer_name, ordered_at
	`

	row := db.QueryRow(queryOrder, orderId, orderReq.CustomerName, orderReq.OrderedAt)

	var orderRes Order

	err := row.Scan(&orderRes.OrderId, &orderRes.CustomerName, &orderRes.OrderedAt)

	if err != nil {
		log.Fatal(err)
	}

	queryItem := `
	UPDATE items 
	SET item_code = $3, description = $4, quantity = $5
	WHERE order_id = $1 AND item_id = $2
	RETURNING item_id, item_code, description, quantity, order_id
	`
	//dibuat untuk bisa update banyak data item jika diperlukan
	itemReq := orderReq.Item

	var itemRes []Item
	for _, v := range itemReq {
		row = db.QueryRow(queryItem, orderRes.OrderId, v.ItemId, v.ItemCode, v.Description, v.Quantity)

		var itemOrder Item
		err = row.Scan(&itemOrder.ItemId, &itemOrder.ItemCode, &itemOrder.Description, &itemOrder.Quantity, &itemOrder.OrderId)

		if err != nil {
			log.Fatal("Error itemRes: ", err)
		}

		itemRes = append(itemRes, itemOrder)
	}

	orderReq.Item = itemRes
	orderReq.OrderId = orderRes.OrderId

	return orderReq
}

// Fungsi untuk menghapus data order dan item sekaligus berdasarkan orderId yang dimasukkan ke parameter
func DeleteOrderDB(orderId int) int64 {
	db := db.GetDB()

	query := `
	DELETE
	FROM orders o
	USING items i
	WHERE o.order_id = i.order_id
	AND o.order_id = $1
	`

	res, err := db.Exec(query, orderId)
	if err != nil {
		log.Fatal(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return count
}
