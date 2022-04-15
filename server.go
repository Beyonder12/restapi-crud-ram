package main

import (
	"fmt"
	"net/http"
	"restapi/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

// OOP we can say replaced by struct
type Users struct {
	ID      int
	Email   string `json:"email"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

//receiver function
func (Users) TableName() string {
	return "users"
}

func main() {
	route := echo.New()
	db := db.Init()

	route.POST("user/create_user", func(c echo.Context) error {
		user := new(Users) //skeleton
		c.Bind(user)       //skeleton binds body

		err := db.Create(user).Error
		if err != nil {
			fmt.Println("error created")
			return c.JSON(http.StatusBadRequest, err)
		}

		response := struct {
			Message string
			Data    Users
		}{
			Message: "New user has been created succesfully",
			Data:    *user, //pointer concept
		}

		return c.JSON(http.StatusOK, response)
	})

	route.GET("user/get_data", func(c echo.Context) error {
		// user := new(Users)
		// c.Bind(user)

		var result []Users
		db.Find(&result)

		response := struct {
			Message string
			Data    []Users
		}{
			Message: "Successfully seen user's data",
			Data:    result,
		}

		return c.JSON(http.StatusOK, response)
	})

	route.PUT("user/update_user/:id", func(c echo.Context) error {
		user := new(Users)
		user.ID, _ = strconv.Atoi(c.Param("id"))
		c.Bind(user)

		var aUser Users

		err := db.
			Where("id = ?", user.ID).
			Find(&aUser).Error

		if err != nil {
			fmt.Println("error created")
			return c.JSON(http.StatusBadRequest, err)
		}

		aUser.Name = user.Name
		aUser.Address = user.Address
		aUser.Age = user.Age

		err = db.Save(aUser).Error
		if err != nil {
			fmt.Println("error created")
			return c.JSON(http.StatusBadRequest, err)
		}

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

	route.DELETE("user/delete_user/:id", func(c echo.Context) error {
		user := new(Users)
		user.ID, _ = strconv.Atoi(c.Param("id"))

		err := db.
			Delete(&user).Error

		if err != nil {
			fmt.Println("error delete")
			return c.JSON(http.StatusBadRequest, err)
		}
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

	route.GET("user/get_data/:id", func(c echo.Context) error {
		user := new(Users)
		user.ID, _ = strconv.Atoi(c.Param("id"))

		err := db.Where("id=?", user.ID).First(&user).Error

		if err != nil {
			fmt.Println("not found the data")
			return c.JSON(http.StatusBadRequest, err)
		}

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
