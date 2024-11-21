package productcontroller

import (
	"encoding/json"
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products" : products})
}

func Show(c *gin.Context){
	var product  models.Product

	id := c.Param("id")
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:	
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H {"Message" : "Data Tidak DiTemukan"}) 
				return
			default: 
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {"Message" : err.Error()}) 
				return
			}
		}	
		c.JSON(http.StatusOK, gin.H{"Product" : product})
}

func Create(c *gin.Context){

	var product models.Product

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {"Message" : err.Error()}) 
				return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"Product" : product})
}

func Update(c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {"Message" : err.Error()}) 
				return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {"Message" : "Tidak Dapat Mengupdate Data"}) 
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message" : "Data Berhasil Di Update"})

}

func Delete(c *gin.Context){

	var product models.Product

	// input := map[string] string{"id" : "0"} 
	var input struct{
		Id json.Number
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {"Message" : err.Error()}) 
				return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id ).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {"Message" : "Tidak Dapat Menghapus Data"}) 
		return	
	}
	c.JSON(http.StatusOK, gin.H{"Message" : "Data Berhasil Di Hapus"})
}