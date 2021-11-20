package controller

import (
	"LoginRegister/dto"
	"LoginRegister/model"
	"LoginRegister/repository"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginController struct {

}

func (lc *LoginController) LoginByUsername(c *gin.Context){
	sess := sessions.Default(c)
	var loginRequest dto.LoginRequest
	var user model.User
	if err := c.ShouldBind(&loginRequest); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request !",
		})
		return
	}
	db := repository.Database
	if db == nil {
		log.Fatal("Database is null")
	}
	rows,_ := db.Query(fmt.Sprintf("SELECT ID, Username, isadmin FROM users WHERE Username='%s' and password='%s'", loginRequest.Username, loginRequest.Password))
	//log.Fatal(err)

	if rows != nil  && rows.Next(){

		rows.Scan(&user.ID, &user.Username, &user.Isadmin)
		sess.Set("ID", user.ID)
		sess.Set("isadmin", user.Isadmin)
		sess.Save()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Welcome %s !", user.Username),
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Username or Password is incorrect !",
		})
	}


}
