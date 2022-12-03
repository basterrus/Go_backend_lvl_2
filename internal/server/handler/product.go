package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/lapitskyss/go_backend_2/internal/pkg/render"
	"github.com/lapitskyss/go_backend_2/internal/pkg/response"
	"github.com/lapitskyss/go_backend_2/internal/srv/producsrv"
)

type Handler struct {
	productSrv *producsrv.ProductSrv
	log        *zap.Logger
}

func InitHandler(productSrv *producsrv.ProductSrv, log *zap.Logger) *Handler {
	return &Handler{
		productSrv: productSrv,
		log:        log,
	}
}

type createProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
	Quantity    uint64 `json:"quantity"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req = &createProductRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		render.BadRequestError(w, errors.New("incorrect request params"))
		return
	}

	product, err := h.productSrv.Create(r.Context(), producsrv.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	})
	if err != nil {
		response.SendError(w, err)
		return
	}

	render.Success(w, product)
}

func (h *Handler) SearchProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	products, err := h.productSrv.Search(r.Context(), query)
	if err != nil {
		response.SendError(w, err)
		return
	}

	render.Success(w, products)
}
