package routes

import (
	"freelink/DB"
	"freelink/encryption"
	"freelink/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Login(engine *gin.Engine, db *gorm.DB) {
	engine.POST("/refresh", func(ctx *gin.Context) {

		var databaseuser = new(DB.User)

		forms := new(struct {
			Username string `json:"username"`
			Password string `json:"password"`
		})

		if err := ctx.ShouldBindJSON(forms); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Where("username = ?", forms.Username).First(&databaseuser).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Username does not exist",
			})
			return
		}

		if encryption.DecryptPassword(forms.Password, databaseuser.Salt) != databaseuser.Password {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid password",
			})
			return
		}

		if usertoken, err := token.Generatetoken(databaseuser.Username); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": usertoken})
		}
	})
}

func Register(engine *gin.Engine, db *gorm.DB) {}
