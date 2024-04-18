package todo

import "github.com/gin-gonic/gin"

// Create Implementation Context Interface

type MyGinContext struct {
	*gin.Context // we can use the thing we compose as a key name
}

func NewGinContext(c *gin.Context) *MyGinContext {
	return &MyGinContext{Context: c}
}

func (c *MyGinContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}

func (c *MyGinContext) JSON(statusCode int, v interface{}) {
	c.Context.JSON(statusCode, v)
}

func (c *MyGinContext) TransactionID() string {
	return c.Context.Request.Header.Get("TransactionID")
}

func (c *MyGinContext) Audience() string {
	if aud, ok := c.Context.Get(("aud")); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}

func ConvertGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(NewGinContext(ctx))
		// handler(&MyGinContext{Context: ctx})
	}
}
