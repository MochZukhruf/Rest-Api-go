package productcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Fungsi untuk memuat data produk dari file JSON
func LoadProductsFromFile() ([]models.Product, error) {
	file, err := os.Open("product.json") // File JSON berisi data produk
	if err != nil {
		return nil, err      
	}
	defer file.Close()

	var products []models.Product
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}

// Fungsi untuk menyimpan data produk ke file JSON
func SaveProductsToFile(products []models.Product) error {
    file, err := os.OpenFile("product.json", os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    // Menggunakan encoder untuk menulis data JSON ke file
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")  // Agar format JSON lebih rapi (opsional)
    // Encode data produk ke dalam file JSON
    if err := encoder.Encode(products); err != nil {
        return err
    }

    return nil
}

// Fungsi untuk menampilkan semua produk
func Index(c *gin.Context) {
	// Mengambil produk dari file JSON
	products, err := LoadProductsFromFile()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// Fungsi untuk menampilkan produk berdasarkan ID
func Show(c *gin.Context) {
	id := c.Param("id") // id yang didapatkan dari parameter URL

	// Mengonversi id dari string ke int64
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID format"})
		return
	}

	// Memuat produk dari file JSON
	products, err := LoadProductsFromFile()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	// Mencari produk berdasarkan ID yang sudah dikonversi
	for _, product := range products {
		if product.Id == idInt { // Membandingkan dengan ID yang sudah dikonversi ke int64
			c.JSON(http.StatusOK, gin.H{"product": product})
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Product not found"})
}

// Fungsi untuk menambahkan produk baru
func Create(c *gin.Context) {
    var product models.Product

    // Bind data JSON ke struct Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
        return
    }

	 fmt.Printf("Received product: %+v\n", product)

    // Memuat produk yang ada
    products, err := LoadProductsFromFile()
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
        return
    }

    // Menambahkan produk baru ke dalam list produk
    products = append(products, product)

    // Menyimpan data ke file JSON
    if err := SaveProductsToFile(products); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"Product": product})
}


// Fungsi untuk memperbarui produk
func Update(c *gin.Context) {
	id := c.Param("id") // id yang didapatkan dari parameter URL

	// Mengonversi id dari string ke int64
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID format"})
		return
	}

	// Memuat produk dari file JSON
	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	products, err := LoadProductsFromFile()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	// Mencari produk berdasarkan ID dan memperbarui
	for i, product := range products {
		if product.Id == idInt { // Membandingkan dengan ID yang sudah dikonversi ke int64
			products[i] = updatedProduct
			break
		}
	}

	// Menyimpan perubahan ke file JSON
	if err := SaveProductsToFile(products); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Product updated successfully"})
}

// Fungsi untuk menghapus produk
func Delete(c *gin.Context) {
	var input struct {
		Id json.Number `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()

	// Memuat produk dari file JSON
	products, err := LoadProductsFromFile()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	// Mencari dan menghapus produk
	var found bool
	for i, product := range products {
		if product.Id == id {
			products = append(products[:i], products[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Product not found"})
		return
	}

	// Menyimpan perubahan setelah penghapusan
	if err := SaveProductsToFile(products); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Product deleted successfully"})
}
