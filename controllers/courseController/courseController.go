package coursecontroller

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/Kazukite12/e-learning/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var course []models.Course

	models.DB.Find(&course)
	c.JSON(http.StatusOK, gin.H{"Course": course})
}

func Show(c *gin.Context) {
	var course models.Course

	id := c.Param("id")

	if err := models.DB.First(&course, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"massage": "course tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"massage": err.Error()})

			return

		}
	}

	c.JSON(http.StatusOK, gin.H{"course": course})

}

func Create(c *gin.Context) {
	var Course models.Course
	if err := c.ShouldBindJSON(&Course); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var existingKeys = make(map[string]bool) // map to store existing keys
	randKey := func() string {
		b := make([]byte, 3) // 3 bytes = 6 hex characters
		rand.Read(b)
		return hex.EncodeToString(b)
	}
	generateKey := func() string {
		key := randKey()
		for existingKeys[key] { // check if the key already exists
			key = randKey()
		}
		existingKeys[key] = true
		return key

	}
	Course.Key = generateKey()
	models.DB.Create(&Course)
	c.JSON(http.StatusOK, gin.H{"Course": Course})
}

func Update(c *gin.Context) {
	var Course models.Course

	id := c.Param("id")
	if err := c.ShouldBindJSON(&Course); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&Course).Where("id = ?", id).Updates(&Course).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"massage": "data tidak dapat di update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"massage": "Data Berhasil di Update"})
}

func Delete(c *gin.Context) {
	var Course models.Course

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&Course, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"massage": "Gagal Menghapus Course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"massage": "Course Berhasil di Hapus"})
}
