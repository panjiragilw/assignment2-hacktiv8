package router

import (
	"assignment2-v4/db"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	db.InitializeDB()
}

type Item struct {
	ItemId      int    `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"orderId"`
}

type Order struct {
	OrderId      int       `json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Item    `json:"item"`
}

func StartRouter() {
	route := gin.Default()

	orderRoute := route.Group("/orders")
	{
		// Membuat data order dan item
		// Contoh tujuan: 127.0.0.1:8080/
		orderRoute.POST("/", CreateOrderHandler)

		// Membaca seluruh data order dan item
		// Contoh tujuan: 127.0.0.1:8080/
		orderRoute.GET("/", ReadAllOrderHandler)

		// Membaca data order dan item berdasarkan order ID
		// Contoh tujuan: 127.0.0.1:8080/1
		orderRoute.GET("/:orderId", ReadOrderHandler)

		// Memperbarui data order dan data item terpilih
		// Output response di postman hanya menampilkan data item yang terupdate
		// Contoh tujuan: 127.0.0.1:8080/1
		orderRoute.PUT("/:orderId", UpdateOrderHandler)

		// Menghapus data order dan item berdasarkan order ID
		// Contoh tujuan: 127.0.0.1:8080/1
		orderRoute.DELETE("/:orderId", DeleteOrderHandler)
	}

	route.Run(":8080")
}

// Fungsi handler untuk alamat tujuan 127.0.0.1:8080/ dengan METHOD POST.
// Fungsi yang mengatur proses pembuatan data
func CreateOrderHandler(c *gin.Context) {
	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"errMsg": err.Error(),
		})

		return
	}

	var orderedAtStr string
	if val, ok := data["orderedAt"].(string); ok {
		orderedAtStr = val
	}

	y, err := time.Parse("2006-01-02", orderedAtStr)

	if err != nil {
		log.Fatal(err)
	}

	items := data["item"].([]interface{})

	var itemReq []Item
	for _, v := range items {
		item := v.(map[string]interface{})
		itemSend := Item{
			ItemCode:    item["itemCode"].(string),
			Description: item["description"].(string),
			Quantity:    int(item["quantity"].(float64)),
		}
		itemReq = append(itemReq, itemSend)
	}

	req := Order{
		CustomerName: data["customerName"].(string),
		OrderedAt:    y,
		Item:         itemReq,
	}

	res := CreateOrderDB(req)

	c.JSON(201, res)
}

// Fungsi handler untuk alamat tujuan 127.0.0.1:8080/ dengan METHOD GET.
// Fungsi yang mengatur proses pembacaan data
func ReadAllOrderHandler(c *gin.Context) {

	res := ReadAllOrderDB()

	c.JSON(201, res)
}

// Fungsi handler untuk alamat tujuan 127.0.0.1:8080/<orderId> dengan METHOD GET.
// Fungsi yang mengatur proses pembacaan data berdasarkan order ID
func ReadOrderHandler(c *gin.Context) {

	paramId, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		log.Fatal(err)
	}
	res := ReadOrderDB(paramId)

	c.JSON(201, res)
}

// Fungsi handler untuk alamat tujuan 127.0.0.1:8080/<orderId> dengan METHOD PUT.
// Fungsi yang mengatur proses pembaruan data
func UpdateOrderHandler(c *gin.Context) {
	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"errMsg": err.Error(),
		})

		return
	}

	var orderedAtStr string
	if val, ok := data["orderedAt"].(string); ok {
		orderedAtStr = val
	}

	y, err := time.Parse(time.RFC3339, orderedAtStr)

	if err != nil {
		log.Fatal(err)
	}

	items := data["item"].([]interface{})

	var itemReq []Item
	for _, v := range items {
		item := v.(map[string]interface{})
		itemSend := Item{
			ItemId:      int(item["itemId"].(float64)),
			ItemCode:    item["itemCode"].(string),
			Description: item["description"].(string),
			Quantity:    int(item["quantity"].(float64)),
		}
		itemReq = append(itemReq, itemSend)
	}

	req := Order{
		CustomerName: data["customerName"].(string),
		OrderedAt:    y,
		Item:         itemReq,
	}

	paramOrderId, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		log.Fatal(err)
	}

	res := UpdateOrderDB(req, paramOrderId)

	c.JSON(201, res)
}

// Fungsi handler untuk alamat tujuan 127.0.0.1:8080/<orderId> dengan METHOD DELETE.
// Fungsi yang mengatur proses penghapusan data
func DeleteOrderHandler(c *gin.Context) {

	paramId, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		log.Fatal(err)
	}
	res := DeleteOrderDB(paramId)

	c.JSON(201, res)
}
