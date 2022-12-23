package contract

import "net/http"

const KernelKey = "htd:kernel"

type Kernel interface {
	HttpEngine() http.Handler
}
