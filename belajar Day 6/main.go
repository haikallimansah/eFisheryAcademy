package main

import{
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"net/http"
}

type Product struct{
	ID int 'json:"id"'
	Name string 'json:"name"'
}

var (
	product = map[int]*Product{}
	nomor = 1
)

func CreateProduct(c echo.Context) error  {
	p := &Product{
		ID: nomor,
	}
	if err := c.Bind(p); err != nil{
		return err
	}
	product[p.ID] = p
	nomor++
	return c.JSON(http.StatusCreated, product)
}

func getProduct(c echo.Context) error{
	return c.JSON(http.StatusOK, product)
}

func main()  {
	// echo instance
	e := echo.New()

	e.Post("/",CreateProduct)
	e.GET("/products",getProduct)
	e.Logger.Fatal(e.Start(":1323"))

	// middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// routes
	// e.GET("/", hello)

	// start server
	// e.Logger.Fatal(e.Start(":1323"))	
}

// func hello(c echo.Context) error  {
// 	return c.String(http.StatusOK, "Hello World")

// }
