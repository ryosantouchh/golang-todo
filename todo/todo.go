package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 1.) define model
// 2.) define handler
// 3.) bind method to handler

type Todo struct {
	gorm.Model // embedded gorm model
	Title string `json:"text"` // attach the tag -- to refer the frontend service need to send the {text: ...}
}

func (Todo) TableName() string {
	return "todos"
}

type TodoHandler struct {
	db *gorm.DB
}

// construct NewHandler to accept the db as a dependency to return TodoHandler that bind with methods
// Once using it ==> todo := NewTodoHandler(&db) -- todo.NewTask(&context)
func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

func (t *TodoHandler) NewTask(c *gin.Context) {
	// extract token
	// s := c.Request.Header.Get("Authorization")
	// tokenString := strings.TrimPrefix(s, "Bearer ")

	// if err := auth.Protect(tokenString); err != nil {
		// c.AbortWithStatus(http.StatusUnauthorized) // Abort will not sent the Request through the next middleware
		// return
	// }

	var todo Todo // -- as Modifiable data --> so using pointer
	// we have BindJson and ShouldBindJson -- second one you can handle error msg , first one, sent 400 code if error
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	r := t.db.Create(&todo)
	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID": todo.ID,
	})
}
