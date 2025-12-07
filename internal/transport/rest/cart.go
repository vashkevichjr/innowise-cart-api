package rest

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vashkevichjr/innowise-cart-api/internal/entity"
	"github.com/vashkevichjr/innowise-cart-api/internal/service"
)

type Handler struct {
	service *service.Cart
}

func NewCartHandler(service *service.Cart) *Handler {
	return &Handler{
		service: service,
	}
}

type CreateCartResponse struct {
	ID        int32             `json:"id"`
	Items     []entity.CartItem `json:"items"`
	CreatedAt time.Time         `json:"created_at"`
}

func (h *Handler) CreateCart(c *gin.Context) {
	row, err := h.service.CreateCart(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	cart := CreateCartResponse{
		ID:        row.Id,
		Items:     row.Items,
		CreatedAt: row.CreatedAt,
	}
	c.JSON(http.StatusCreated, cart)
}

type ItemRequest struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type ItemResponse struct {
	Id        int32     `json:"item_id"`
	Product   string    `json:"product_id"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *Handler) CreateItem(c *gin.Context) {
	var req ItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if req.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price is required"})
		return
	}

	row, err := h.service.CreateItem(c, req.Name, req.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	item := ItemResponse{
		Id:        row.ID,
		Product:   row.Name,
		Price:     row.Price,
		CreatedAt: row.CreatedAt,
	}

	c.JSON(http.StatusCreated, item)
}

type AddItemRequest struct {
	Quantity int32 `json:"quantity"`
}

type AddItemResponse struct {
	CartID    int32     `json:"cart_id"`
	ItemID    int32     `json:"item_id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h *Handler) AddItemToCart(c *gin.Context) {
	cartIdStr := c.Param("cart_id")
	cartId, err := strconv.Atoi(cartIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart id"})
		return
	}

	itemIdStr := c.Param("item_id")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	var req AddItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid AddItem request"})
		return
	}

	cartItemRow, err := h.service.AddItemToCart(c.Request.Context(), int32(cartId), int32(itemId), req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	cartItem := AddItemResponse{
		CartID:    cartItemRow.CartID,
		ItemID:    cartItemRow.ItemID,
		Name:      cartItemRow.Name,
		Price:     cartItemRow.Price,
		Quantity:  cartItemRow.Quantity,
		CreatedAt: cartItemRow.CreatedAt,
		UpdatedAt: cartItemRow.UpdatedAt,
	}

	c.JSON(http.StatusOK, cartItem)
}

func (h *Handler) RemoveItemFromCart(c *gin.Context) {
	cartIdStr := c.Param("cart_id")
	cartId, err := strconv.Atoi(cartIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart id"})
		return
	}

	itemIdStr := c.Param("item_id")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	err = h.service.RemoveFromCart(c.Request.Context(), int32(cartId), int32(itemId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type ViewCartResponse struct {
	CartID    int32             `json:"cart_id"`
	ItemId    int32             `json:"item_id"`
	Items     []entity.CartItem `json:"items"`
	CreatedAt time.Time         `json:"created_at"`
}

func (h *Handler) ViewCart(c *gin.Context) {
	cartIdStr := c.Param("id")
	cartId, err := strconv.Atoi(cartIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart id"})
		return
	}

	cart, err := h.service.ViewCart(c.Request.Context(), int32(cartId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	viewCart := ViewCartResponse{
		CartID:    cart.Id,
		Items:     cart.Items,
		CreatedAt: cart.CreatedAt,
	}

	c.JSON(http.StatusOK, viewCart)
}

type CalculateCartResponse struct {
	CartID          int32   `json:"cart_id"`
	TotalPrice      float32 `json:"total_price"`
	DiscountPercent int32   `json:"discount_percent"`
	FinalPrice      float32 `json:"final_price"`
}

func (h *Handler) CalculateCart(c *gin.Context) {
	cartIdStr := c.Param("id")
	cartId, err := strconv.Atoi(cartIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart id"})
		return
	}

	row, err := h.service.CalculatePrice(c.Request.Context(), int32(cartId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	calculator := CalculateCartResponse{
		CartID:          row.CartID,
		TotalPrice:      row.TotalPrice,
		DiscountPercent: row.DiscountPercent,
		FinalPrice:      row.FinalPrice,
	}

	c.JSON(http.StatusOK, calculator)
}
