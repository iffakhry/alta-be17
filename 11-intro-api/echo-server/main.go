package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type article struct {
	ID      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// create ne echo instance
	e := echo.New()
	// membuat route / endpoint
	e.GET("/articles", GetAllArticles)
	// e.POST("/articles", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]any{
	// 		"message": "data berhasil ditambahkan",
	// 		"status":  true,
	// 	})
	// })

	// endpoint with param
	e.GET("/users/:id", GetUserByIdController)
	// sample endpoint with query param
	e.GET("/users", GetAllUsersController)
	// sample endpoint with form value
	e.POST("/users", AddUserController)
	// sample endpoint with bind json
	e.POST("/articles", AddArticleController)
	e.PUT("/articles/:id", UpdateArticleController)

	e.DELETE("/users/:id", GetUserByIdController)
	// start server di port
	e.Logger.Fatal(e.Start(":8080"))

}

func GetAllArticles(c echo.Context) error {
	// return c.JSON(http.StatusOK, map[string]any{
	// 	"message": "data ditemukan",
	// 	"status":  true,
	// })
	// dummy data
	var data = []article{
		article{1, "lorem", "lorem"},
		article{2, "ipsum", "ipsum"},
		article{3, "data3", "content3"},
		article{4, "data4", "content4"},
	}
	// return c.JSON(http.StatusOK, data)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "data ditemukan",
		"status":  true,
		"data":    data,
	})
}

// func with param
func GetUserByIdController(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "id harus angka",
		})
	}

	user1 := User{
		Id:    idConv,
		Name:  "Budi",
		Email: "budi@mail.com",
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "berhasil",
		"id":      id,
		"data":    user1,
	})
}

// query param
func GetAllUsersController(c echo.Context) error {
	pageQuery := c.QueryParam("page")
	nameQuery := c.QueryParam("name")
	domisiliQuery := c.QueryParam("domisili")
	if domisiliQuery != "" {
		// ...
	}

	// select * from users where name like "%budi%" and umur = 30

	return c.JSON(http.StatusOK, map[string]any{
		"message":  "berhasil",
		"page":     pageQuery,
		"name":     nameQuery,
		"domisili": domisiliQuery,
	})
}

// form value
func AddUserController(c echo.Context) error {
	//get data from form value
	nameForm := c.FormValue("name")
	emailForm := c.FormValue("email")

	var userNew = User{
		Name:  nameForm,
		Email: emailForm,
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "berhasil",
		"data":    userNew,
	})
}

// bind
func AddArticleController(c echo.Context) error {
	articleInput := article{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&articleInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}
	// c.Request().FormFile("gambar")

	log.Println("article title", articleInput.Title)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "berhasil",
		"data":    articleInput,
	})
}

func UpdateArticleController(c echo.Context) error {
	id := c.Param("id")

	articleInput := article{}
	errBind := c.Bind(&articleInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}
	// articleInput.ID = id
	return c.JSON(http.StatusOK, map[string]any{
		"message": "berhasil",
		"id":      id,
		"data":    articleInput,
	})
}
