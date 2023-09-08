package controller

import (
	"context"
	"github/phat/product/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	Phone int   `json:"phone"`
	Contacted bool `json:"contacted"`
}

func (a *APIEnv) Register(c *gin.Context){
	_, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var customer model.Customers
	if err := c.BindJSON(&customer); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := a.DB.Model(&model.Customers{}).Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Failed",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Create Successfully",
		"data": customer,
	})
}

func (a *APIEnv) GetCustomers (c *gin.Context){

	var customers []model.Customers
	if data := a.DB.Find(&customers); data.Error != nil {
        c.AbortWithError(http.StatusNotFound, data.Error)
        return
    }

	log.Println(customers)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Customers Successfully",
		"data": customers,
	})
}


func (a *APIEnv) GetDetailCustomer(c *gin.Context){
	var customer model.Customers
	id := c.Param("id")

	if err := a.DB.Where("ID = ?", id).Find(&customer); err.Error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
		"Data": customer,
	})
}

func (a *APIEnv) DeleteMovie(c *gin.Context){                                                                       
	customer := []model.Customers{}
	id := c.Param("id")
	//check list 
	err :=  a.DB.Model(&model.Customers{}).Where("id = ?", id).Delete(&customer).Error
	if err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"Message": "Delete success",
		"Data": customer,
	})
}

func (a *APIEnv) UpdateMovie(c *gin.Context){
	var customer model.Customers
	var checkList []model.Customers
	id := c.Param("id")

	err := c.BindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	//check exists
	check := a.DB.Model(&model.Customers{}).Where("id = ?", id).Find(&checkList).Error
	if check != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Customer not exists!",
		})
		return
	}
	if len(checkList) < 1{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Customer not exists!",
		})
		return
	}
	//end check
	if err = a.DB.Model(&model.Customers{}).Where("id = ?", id).Updates(&customer).Error; err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Update Success",      
		"Data": customer,
	})
}