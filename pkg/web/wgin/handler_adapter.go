package wgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pvillela/gfoa/pkg/web"
	log "github.com/sirupsen/logrus"
)

// TODO: Error handling is inconsistent between the two kinds of handler.

func PostHanderMaker(pInput Any, svc func(Any) Any) gin.HandlerFunc {

	return func(c *gin.Context) {
		filler := func(pInput Any) error {
			// Bind JSON content of request body to
			// struct created above
			err := c.BindJSON(pInput)
			if err != nil {
				// Gin automatically returns an error
				// response when the BindJSON operation
				// fails, we simply have to stop this
				// function from continuing to execute

				log.Info(err)

				c.JSON(400, gin.H{
					"msg":  err.Error(),
					"code": 999,
				})
				return err
			}
			return nil
		}

		pseudoHdlr := web.PostPseudoHandler(pInput, svc)

		pResp, err := pseudoHdlr(filler)

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, pResp)
	}

}

func GeneralGetHanderMaker(svc func(map[string]string) (Any, error)) gin.HandlerFunc {

	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		m := make(map[string]string, len(params))
		for k, vs := range params {
			m[k] = vs[0]
		}

		pResp, err := svc(m)

		if err != nil {
			c.JSON(400, gin.H{
				"msg":  err.Error(),
				"code": 888,
			})
			return
		}

		c.JSON(http.StatusOK, pResp)
	}
}
