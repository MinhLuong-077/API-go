package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"stockexchange.com/entity"
)

var DB *gorm.DB
func CreateNewTransaction(db *gorm.DB, Transaction *entity.Transaction) (err error) {
	err = db.Create(Transaction).Error
	if err != nil {
		return err
	}
	return nil
}


//create
func CreateTransaction(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	 c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	 c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var transaction entity.Transaction
	c.BindJSON(&transaction)
	err := CreateNewTransaction(DB, &transaction)
	if err != nil {
	   c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	   return
	}
	c.JSON(http.StatusOK, transaction)
	
 }
func GroupApi(router *gin.RouterGroup) {
	api := router.Group("/api/v1")
	{
		api.POST("/users/transaction/create", CreateTransaction)
	}
}
