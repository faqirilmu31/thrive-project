package controllers

import (
	"encoding/json"
	"golang-thrive/api/models"
	"golang-thrive/api/responses"
	"golang-thrive/api/utils/formaterror"
	"io/ioutil"
	"net/http"

	"github.com/gofrs/uuid"
)

type cartID struct {
	CartId int32 `json:"cart_id"`
}

type checkOut struct {
	Cart []cartID `json:"cart"`
}

type transactionID struct {
	TransactionID string `json:"transaction_id"`
}

type detailTransaction struct {
	ProductName   string  `json:"product_name"`
	PricePerPiece float32 `json:"price_per_piece"`
	Quantity      int32   `json:"quantity"`
	TotalPrice    float32 `json:"total_price"`
}

type responseTotalTransaction struct {
	TransactionID         string              `json:"transaction_id"`
	TotalPriceTransaction float32             `json:"total_price_trasaction"`
	DetailTransaction     []detailTransaction `json:"detail_transaction"`
}

func (server *Server) CheckOut(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	var cartID checkOut
	err = json.Unmarshal(body, &cartID)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	carts := models.Cart{}
	products := models.Product{}

	uuid := uuid.Must(uuid.NewV4())

	for _, v := range cartID.Cart {
		// Find Carts to be checkedout
		cart, err := carts.FindCartByID(server.DB, v.CartId)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		// Find Products data
		product, err := products.FindProductByID(server.DB, cart.ProductId)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}

		// Define item
		var harga float64
		harga = float64(product.Price) * float64(cart.Total)

		err = models.InsertTransaction(server.DB, models.Transaction{
			TransactionId: uuid.String(),
			ProductId:     cart.ProductId,
			UserId:        cart.UserId,
			CartId:        cart.CartId,
			Total:         cart.Total,
			Price:         float32(harga),
			PaymentStatus: "Done",
		})
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}

		// Update Quantity Product
		product.Quantity = product.Quantity - cart.Total
		err = product.UpdateUnit(server.DB, cart.ProductId)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}

		// Remove from cart table
		err = cart.DeleteCart(server.DB, cart.CartId)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}
	}

	response := map[string]interface{}{
		"status_code": "200",
		"status_desc": "Checkout Success",
	}
	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CheckTotalPrice(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	var transactionId transactionID
	err = json.Unmarshal(body, &transactionId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	transaction := models.Transaction{}
	//Find Transaction data
	transactions, err := transaction.FindTransactionByID(server.DB, transactionId.TransactionID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	var responseTotalPrice responseTotalTransaction
	responseTotalPrice.TransactionID = transactionId.TransactionID
	responseTotalPrice.TotalPriceTransaction = 100000

	var totalPriceTransaction float32

	products := models.Product{}
	for _, v := range *transactions {
		// Find Products data
		product, err := products.FindProductByID(server.DB, v.ProductId)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}

		totalPriceTransaction += v.Price

		newStruct := &detailTransaction{
			ProductName:   product.DisplayName,
			PricePerPiece: product.Price,
			Quantity:      v.Total,
			TotalPrice:    v.Price,
		}
		responseTotalPrice.DetailTransaction = append(responseTotalPrice.DetailTransaction, *newStruct)

	}
	responseTotalPrice.TotalPriceTransaction = totalPriceTransaction
	responses.JSON(w, http.StatusCreated, responseTotalPrice)
}
