package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	I "github.com/mohdjishin/GoCart/interfaces"
	"github.com/mohdjishin/GoCart/model"
	utils "github.com/mohdjishin/GoCart/utils"
)

type Products struct{}

func NewProduct() I.IProduct {
	return &Products{}
}

func (*Products) AddProducts(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)

	product := new(model.Products)

	product.Product_Name = c.FormValue("pro_name")

	if price, err := strconv.ParseFloat(c.FormValue("price"), 64); err == nil {
		product.Price = price
	}
	product.Product_Category = strings.ToLower(c.FormValue("pro_category"))

	fileOne, err := c.FormFile("img_one")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "file upload error",
		})
	}

	fileTWo, err := c.FormFile("img_two")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "file upload error",
		})
	}
	fileThree, err := c.FormFile("img_three")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "file upload error",
		})
	}

	res := db.Save(&product)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed in creating products",
		})
	}

	productImage := new(model.ProductImage)

	productImage.ProductId = product.ID

	fileOne.Filename = uuid.New().String() + path.Ext(fileOne.Filename)
	fileTWo.Filename = uuid.New().String() + path.Ext(fileTWo.Filename)
	fileThree.Filename = uuid.New().String() + path.Ext(fileThree.Filename)

	url1, status1, _ := utils.UploadToBucket(fileOne)
	if !status1 {

		_ = utils.InternalServerError("img one upload failed", c)
	}

	url2, status2, _ := utils.UploadToBucket(fileTWo)
	if !status2 {

		_ = utils.InternalServerError("img two upload failed", c)
	}
	url3, status3, _ := utils.UploadToBucket(fileThree)
	if !status3 {

		_ = utils.InternalServerError("img three upload failed", c)
	}
	fileOne.Filename = url1
	fileTWo.Filename = url2
	fileThree.Filename = url3

	productImage.ImageOne = fileOne.Filename
	productImage.ImgTwo = fileTWo.Filename
	productImage.ImgThree = fileThree.Filename

	res = db.Save(&productImage)
	if res.Error != nil {

		_ = utils.InternalServerError("failed in creating products", c)

	}

	return c.Status(200).JSON(fiber.Map{
		"message": "product created",
	})
}

func (*Products) UpdatePro(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)

	id := c.Params("id")
	fmt.Println(id)

	pImages := new(model.ProductImage)
	e := db.First(&pImages, "product_id=?", id)
	if e.Error != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	pro := new(model.Products)
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	pro.ID = uint(u64)
	pro.Product_Name = c.FormValue("pro_name")
	if price, err := strconv.ParseFloat(c.FormValue("price"), 64); err == nil {
		pro.Price = price
	}
	pro.Product_Category = strings.ToLower(c.FormValue("pro_category"))

	fileOne, err := c.FormFile("img_one")

	if err != nil {

		return c.Status(400).JSON(fiber.Map{
			"message": pImages.ImageOne + " upload error",
		})
	}

	fileTwo, err := c.FormFile("img_two")

	if err != nil {

		return c.Status(400).JSON(fiber.Map{
			"message": pImages.ImgTwo + " upload error",
		})
	}

	fileThree, err := c.FormFile("img_three")
	if err != nil {

		return c.Status(400).JSON(fiber.Map{
			"message": pImages.ImgThree + " upload error",
		})
	}

	var wg sync.WaitGroup

	var url1 string
	var status1 bool

	var url2 string
	var status2 bool

	var url3 string
	var status3 bool

	wg.Add(1)
	go func(url *string, status *bool, fileOne *multipart.FileHeader, w *sync.WaitGroup) {
		*url, *status, _ = utils.UploadToBucket(fileOne)
		if !*status {

			_ = utils.InternalServerError("img one upload failed", c)
		}
		w.Done()
	}(&url1, &status1, fileOne, &wg)
	wg.Add(1)
	go func(url *string, status *bool, fileTwo *multipart.FileHeader, w *sync.WaitGroup) {
		*url, *status, _ = utils.UploadToBucket(fileTwo)
		if !*status {

			_ = utils.InternalServerError("img two upload failed", c)
		}
		w.Done()
	}(&url2, &status2, fileTwo, &wg)

	wg.Add(1)
	go func(url *string, status *bool, fileThree *multipart.FileHeader, w *sync.WaitGroup) {
		*url, *status, _ = utils.UploadToBucket(fileThree)
		if !*status {

			_ = utils.InternalServerError("img three upload failed", c)
		}
		w.Done()
	}(&url3, &status3, fileTwo, &wg)
	wg.Wait()

	fileOne.Filename = url1
	fileTwo.Filename = url2
	fileThree.Filename = url3

	pImages.ImageOne = fileOne.Filename
	pImages.ImgTwo = fileTwo.Filename
	pImages.ImgThree = fileThree.Filename

	db.Save(&pImages)
	db.Save(&pro)

	return c.Status(200).JSON(fiber.Map{"message": "updated"})

}

func (*Products) DelProduct(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
	id := c.Params("id")
	fmt.Println(id)

	proImg := new(model.ProductImage)

	db.First(&proImg, "product_id = ?", id)

	if proImg.ID == '0' {
		return c.Status(404).JSON(fiber.Map{
			"message": "no match found",
		})
	}

	err := db.Delete(&proImg, "product_id = ?", id)
	if err.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "unable to prossess img deletion",
		})
	}
	proImg = nil

	pro := new(model.Products)
	err = db.Delete(&pro, id)
	if err.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "unable to prossess deletion",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "product deleted",
	})
}
func (*Products) ViewProducts(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)

	var products []model.Products

	err := db.Where("id > ?", 0).Find(&products).Error

	if err != nil {
		fmt.Println(err)
	}

	var productsImages []model.ProductImage

	err = db.Where("id > ?", 0).Find(&productsImages).Error

	if err != nil {
		fmt.Println(err)
	}

	infoWithoutImageDetails := utils.ExtractProductInfo(products)
	infowithoutPersonal := utils.ExtractProImage(productsImages)

	comb := utils.CombinePRoductAndProductImage(infoWithoutImageDetails, infowithoutPersonal)

	return c.Status(200).JSON(comb)

}

func (*Products) GetbyCategory(c *fiber.Ctx) error {
	db := DB.OpenDb()
	defer DB.CloseDb(db)
	type category struct {
		Category string `json:"pro_category"`
	}
	ctgry := new(category)

	if err := c.BodyParser(ctgry); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var products []model.Products

	err := db.Where("product_category = ?", strings.ToLower(ctgry.Category)).Find(&products).Error

	if err != nil {
		return c.Status(200).JSON(fiber.Map{
			"message": "no product found with category : " + ctgry.Category,
		})
	}

	return c.Status(200).JSON(products)
}

func (*Products) SearchProduct(c *fiber.Ctx) error {
	db := DB.OpenDb()

	defer DB.CloseDb(db)
	id := c.Params("key")
	fmt.Println(id)

	var products []model.Products

	res := db.Select("id", "price", "product_category", "product_name").Find(&products, "product_name LIKE ?", id+"%")

	if res.Error != nil {
		return c.Status(200).JSON(fiber.Map{
			"message": "no product with search key +: " + id,
		})
	}
	if len(products) == 0 {
		return c.Status(200).JSON(fiber.Map{
			"message": "no product with search key +: " + id,
		})
	}

	return c.Status(200).JSON(products)
}
