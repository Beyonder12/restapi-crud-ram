package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Users struct {
	Email   string
	Name    string
	Age     int
	Address string
}

func main() {
	route := echo.New()

	route.POST("user/create_user", func(c echo.Context) error {
		user := new(Users)
		c.Bind(user)

		response := struct {
			Message string
			Data    Users
		}{
			Message: "New user has been created succesfully",
			Data:    *user,
		}

		return c.JSON(http.StatusOK, response)
	})

	route.PUT("user/update_user/:email", func(c echo.Context) error {
		user := new(Users)
		user.Email = c.Param("email")
		//update to the database
		response := struct {
			Message string
			Data    Users
		}{
			Message: "User data has been updated successfully!",
			Data:    *user,
		}

		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("user/delete_user/:email", func(c echo.Context) error {
		user := new(Users)
		user.Email = c.Param("email")
		//delete data
		response := struct {
			Message string
			ID      string
		}{
			Message: "Successfully deleted data",
			ID:      user.Email,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("user/get_data", func(c echo.Context) error {
		user := new(Users)
		user.Email = "fajri.illahi1211@gmail.com"
		user.Name = "Fajri Illahi"
		user.Age = 27
		user.Address = "Flamboyan"

		response := struct {
			Message string
			Data    Users
		}{
			Message: "Successfully seen user's data",
			Data:    *user,
		}

		return c.JSON(http.StatusOK, response)
	})

	route.Start(":8081")

}
