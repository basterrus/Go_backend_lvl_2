package response

import (
	"net/http"

	"github.com/lapitskyss/go_backend_2/internal/pkg/render"
)

func SendError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *BadRequestError:
		render.BadRequestError(w, err)
	case *NotFoundError:
		render.NotFoundError(w)
	default:
		render.InternalServerError(w)
	}
}
