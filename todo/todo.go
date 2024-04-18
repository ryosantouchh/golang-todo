package todo

import (
	"log"
	"net/http"
	"time"
)

// 1.) define model
// 2.) define handler
// 3.) bind method to handler

type Todo struct {
	Title string `json:"text" binding:"required"` // attach the tag -- to refer the frontend service need to send the {text: ...}
	// gorm.Model        // embedded gorm model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt DeletedAt `gorm:"index"`
}

type storer interface {
	New(*Todo) error
}

func (Todo) TableName() string {
	return "todos"
}

type H map[string]interface{}

type TodoHandler struct {
	store storer
}

// construct NewHandler to accept the db as a dependency to return TodoHandler that bind with methods
// Once using it ==> todo := NewTodoHandler(&db) -- todo.NewTask(&context)
func NewTodoHandler(store storer) *TodoHandler {
	return &TodoHandler{store: store}
}

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
	TransactionID() string
	Audience() string
}

func (t *TodoHandler) NewTask(c Context) {
	// extract token
	// s := c.Request.Header.Get("Authorization")
	// tokenString := strings.TrimPrefix(s, "Bearer ")

	// if err := auth.Protect(tokenString); err != nil {
	// c.AbortWithStatus(http.StatusUnauthorized) // Abort will not sent the Request through the next middleware
	// return
	// }

	var todo Todo // -- as Modifiable data --> so using pointer
	// we have BindJson and ShouldBindJson -- second one you can handle error msg , first one, sent 400 code if error
	// if err := c.ShouldBindJSON(&todo); err != nil {
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, H{
			"error": err.Error(),
		})
		return
	}

	if todo.Title == "sleep" {
		// transactionID := c.Request.Header.Get("TransactionID")
		transactionID := c.TransactionID()
		// aud, _ := c.Get("aud")
		aud := c.Audience()
		log.Println(transactionID, aud, "not allowed")
		c.JSON(http.StatusBadRequest, H{
			"error": "not allowed",
		})
		return
	}

	// r := t.db.Create(&todo)
	// if err := r.Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	err := t.store.New(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, H{
		"ID": todo.ID,
	})
}

// func (t *TodoHandler) GetTodoList(c *gin.Context) {
// 	var todos []Todo
// 	r := t.db.Find(&todos)
// 	if err := r.Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, todos)
// }

// func (t *TodoHandler) DeleteTodo(c *gin.Context) {
// 	idParam := c.Param("id")
// 	log.Println(idParam)

// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	response := t.db.Delete(&Todo{}, id)
// 	if err := response.Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "success",
// 	})
// }
