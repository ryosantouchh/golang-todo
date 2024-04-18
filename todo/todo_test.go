package todo

import (
	"fmt"
)

// create mirror stuffs

func TestCreateTodoNotAllowSleepTask() {
	handler := NewTodoHandler(&TestDB{})
	c := &TestContext{}

	handler.NewTask(c)

	want := "not allowed"

	// if want != c.v["error"] {
	// 	fmt.Errorf("want %s but get %s\n", want, c.v["error"])
	// }

	fmt.Println(want)
}

type TestDB struct{}

func (TestDB) New(*Todo) error {
	return nil
}

type TestContext struct {
	v map[string]interface{}
}

func (TestContext) Bind(v interface{}) error {
	// imaging that v is type Todo
	*v.(*Todo) = Todo{
		Title: "sleep",
	}
	// return c.Context.ShouldBindJSON(v)
	return nil
}

func (t *TestContext) JSON(c int, v interface{}) {
	// c.Context.JSON(statusCode, v)
	t.v = v.(map[string]interface{}) // assert
}

func (TestContext) TransactionID() string {
	return "TestTransactionID"
}

func (TestContext) Audience() string {
	// if aud, ok := c.Context.Get(("aud")); ok {
	// 	if s, ok := aud.(string); ok {
	// 		return s
	// 	}
	// }
	return ""
}
