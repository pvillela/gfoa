package wgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pvillela/gfoa/pkg/web"
)

func PostHanderMaker0(pseudoHdlr web.PseudoPostHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		filler := func(pReq interface{}) error {
			// Bind JSON content of request body to
			// struct created above
			err := c.BindJSON(pReq)
			if err != nil {
				// Gin automatically returns an error
				// response when the BindJSON operation
				// fails, we simply have to stop this
				// function from continuing to execute

				fmt.Println(err)

				c.JSON(400, gin.H{
					"msg":  err.Error(),
					"code": 999,
				})
				return err
			}
			return nil
		}
		pResp, err := pseudoHdlr(filler)

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, pResp)
	}
}

func pseudoPostHandler(pReq Any, svc func(Any) Any) func(web.Filler) (interface{}, error) {
	return func(filler web.Filler) (Any, error) {
		err := filler(pReq)

		if err != nil {
			return nil, err
		}

		resp := svc(pReq)

		return &resp, nil
	}
}

func PostHanderMaker(pData Any, svc func(Any) Any) gin.HandlerFunc {

	return func(c *gin.Context) {
		filler := func(pReq interface{}) error {
			// Bind JSON content of request body to
			// struct created above
			err := c.BindJSON(pReq)
			if err != nil {
				// Gin automatically returns an error
				// response when the BindJSON operation
				// fails, we simply have to stop this
				// function from continuing to execute

				fmt.Println(err)

				c.JSON(400, gin.H{
					"msg":  err.Error(),
					"code": 999,
				})
				return err
			}
			return nil
		}

		pseudoHdlr := pseudoPostHandler(pData, svc)

		pResp, err := pseudoHdlr(filler)

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, pResp)
	}

}