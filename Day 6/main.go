package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Article struct {
	Title	string	`form:"title" json:"title"`
	Description  string `form:"description" json:"description"`
	Content	string    `form:"content" json:"content"`
}

var articles []Article

func createArticle(c echo.Context) error {
	var article Article
	if err :=c.Bind(&article); err != nil {
		return err
	}
	

	var article Article{
		Title := c.FormValue("title"),
		Description := c.FormValue("description"),
		Content := c.FormValue("content"),
	}
	articles = append(articles, article)
	printProduct(products)
	return c.NoContent(http.StatusCreated)
}

func printArticles(articles []Article) {
	for i, article := range articles {
		fmt.Printf("%d. %s", i + 1, article.title)
	}
	fmt.Printf("Total produk : %d\n", len(articles))
}

func showArticle(c echo.Context) error  {
	articleId, err := strconv.Atoi(c.Param("id"))

	if len(articles) < articleId {
		return c.NoContent(http.StatusNotFound)
	}

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	
	article := articles[articleId]
	if article.Deleted {
		return c.NoContent(http.StatusNotFound)
	}

	
	if err = c.Bind(&article); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	articles[articleId] = article
	
	return c.NoContent(http.StatusOK)
}

func deleteArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("id"))

	if len(articles) < articleId {
		return c.NoContent(http.StatusNotFound)
	}

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	
	article := articles[articleId]
	article.Deleted = true

	articles[articleId].Deleted = true
	return c.NoContent(http.StatusGone)
}

func main() {
	articles = make([]Article, 0)
	e := echo.New()
	e.POST("/articles", createArticle)
	e.GET("/articles", func(c echo.Context) error {
		return c.JSON(http.StatusOK, articles)
	})
	e.PUT("/articles/:id", updateArticle)
	e.DELETE("/articles/:id", deleteArticle)
	e.Logger.Fatal(e.Start(":1323"))
	fmt.Println("Server started at localhost:1323")
}
