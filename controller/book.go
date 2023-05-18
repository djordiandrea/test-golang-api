package controller

import (
	"log"
	"net/http"
	"rest-api/database"
	"rest-api/model"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {

	var item []model.Mst_user
	database.DB.Table("mst_user").Select("*").Scan(&item)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": item,
	})

}

func GetUserDetail(c *gin.Context) {

	username := c.Param("username")
	var user = model.Mst_user{}
	database.DB.Table("mst_user").Select("*").Where("mus_username = ?", username).Scan(&user)

	if user.Mus_fullname == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errMsg": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": user,
		"errMsg": "",
	})

}

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Djordi Yoel Andrea",
		"bio":  "Saya tamvan",
	})
}

func RHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title": "Hello World",
		"Body":  "Biji Kuda Kamu",
	})
}

func RBook(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func RQuery(ctx *gin.Context) {
	query := ctx.Query("q") //yg di pake buat parameter, si q ini

	ctx.JSON(http.StatusOK, gin.H{
		"title": query,
	})
}

func DoubleParam(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"id":    id,
	})
}

func DoubleQuery(ctx *gin.Context) {
	item := ctx.Query("item") //yg di pake buat parameter, si q ini
	harga := ctx.Query("harga")

	ctx.JSON(http.StatusOK, gin.H{
		"item":  item,
		"harga": harga,
	})
}

func PostBook(ctx *gin.Context) {

	var ib model.InputBook

	err := ctx.ShouldBindJSON(&ib)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":    ib.Title,
		"price":    ib.Price,
		"subtitle": ib.SubTitle,
	})
}

func PostWithValidation(c *gin.Context) {
	var item model.ItemStruct

	err := c.ShouldBindJSON(&item)

	if err != nil {
		//
		// fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"descErr": err,
		})
		return //kayak break, biar codingan yg bawah ga dijalanin
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    item.Name,
		"price":   item.Price,
		"qty":     item.Quantity,
		"status":  http.StatusOK,
		"descErr": err,
	})

	// if err != nil {
	// 	//
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"descErr": err,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name":    item.Name,
	// 		"price":   item.Price,
	// 		"qty":     item.Quantity,
	// 		"status":  http.StatusOK,
	// 		"descErr": err,
	// 	})
	// }

}
