package bank

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"stockexchange.com/entity"
)

type Status struct {
	Money  float64 `json:"money"`
	Status string  `json:"status"`
}

func CreateBankAccount(db *gorm.DB, Account *entity.Bankaccount) (err error) {
	err = db.Create(Account).Error
	if err != nil {
		return err
	}
	return nil
}

//get
func GetAllAccounts(db *gorm.DB, Account *[]entity.Bankaccount) (err error) {
	err = db.Find(Account).Error
	if err != nil {
		return err
	}
	return nil
}

func GetBankAccount(db *gorm.DB, Account *entity.Bankaccount, account string, bank string) (err error) {
	err = db.Where("account = ? and bank = ?", account, bank).First(Account).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBankAccountMoney(db *gorm.DB, status *Status, account string,bank string,userid string) (err error) {
   fmt.Println(status)
   var op string 
	if status.Status == "1" {
		op = "-"
	}
	if status.Status == "2" {
		op = "+"
	}
	err =  db.Model(&entity.Bankaccount{ID : userid ,Account: account, Bank: bank}).Update("money", gorm.Expr(fmt.Sprintf("money %s ?", op), status.Money)).Error
	if err != nil {
		return err
	}
	return nil
}

var DB *gorm.DB

//create
func CreateAccount(c *gin.Context) {
	var account entity.Bankaccount
	c.BindJSON(&account)
	err := CreateBankAccount(DB, &account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, account)
}

func GetAccounts(c *gin.Context) {
	var account []entity.Bankaccount
	err := GetAllAccounts(DB, &account)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, account)
}
func GetAccount(c *gin.Context) {
	account, _ := c.Params.Get("account")
	bank, _ := c.Params.Get("bank")
	var accounts entity.Bankaccount
	err := GetBankAccount(DB, &accounts, account, bank)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func GetAllBankAccount(db *gorm.DB,Account *[]entity.Bankaccount, bank string) (err error) {
	err = db.Where("bank = ?", bank).Find(Account).Error
	if err != nil {
		return err
	}
	return nil
}
func GetAccountBank(c *gin.Context) {
	bank, _ := c.Params.Get("bank")
	var accounts []entity.Bankaccount
	err := GetAllBankAccount(DB, &accounts, bank)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func UpdateAccountMoney(c *gin.Context) {
   var accounts entity.Bankaccount
   var account string = c.Param("account")
   var bank string = c.Param("bank")
   var userid string = c.Param("userid")
   var status Status
   c.BindJSON(&status)
   err := UpdateBankAccountMoney(DB, &status, account,bank,userid)
   if err != nil {
      if errors.Is(err, gorm.ErrRecordNotFound) {
         c.AbortWithStatus(http.StatusNotFound)
         return
      }
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
      return
   }
   err = GetBankAccount(DB, &accounts, account,bank)
   if err != nil {
      if errors.Is(err, gorm.ErrRecordNotFound) {
         c.AbortWithStatus(http.StatusNotFound)
         return
      }
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
      return
   }
   c.JSON(http.StatusOK, accounts)
}
func GroupApi(router *gin.RouterGroup) {

	api := router.Group("/api/v1")
	{
		api.POST("/bank-account/create", CreateAccount)
		api.GET("/bank-account/:account/:bank", GetAccount)
		api.GET("/bank-account/bank/:bank", GetAccountBank)
		api.GET("/bank-account/accounts", GetAccounts)
		api.PUT("/bank-account/:account/:bank/:userid/update/money", UpdateAccountMoney)
	}
}
