package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/pvillela/gfoa/pkg/web"
)

func HandlerAdapter(helper web.HandlerHelper) gin.HandlerFunc {
	return nil
}

func HandlerAdapterW(helper web.WithHandlerHelper) gin.HandlerFunc {
	return nil
}
