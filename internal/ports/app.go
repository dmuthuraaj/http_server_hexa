package ports

import "net/http"

type HandlerAppPort interface {
	Handler() *http.ServeMux
}
