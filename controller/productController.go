package controller

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mohdjishin/GoCart/database"
	"github.com/mohdjishin/GoCart/model"
	utils "github.com/mohdjishin/GoCart/utils"
)

func AddProducts(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)

	product := new(model.Products)

	product.Product_Name = c.FormValue("pro_name")
	fmt.Println(product.Product_Name)
	if price, err := strconv.ParseFloat(c.FormValue("price"), 64); err == nil {
		product.Price = price
	}
	product.Product_Category = c.FormValue("pro_category")

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

	fmt.Println(product.ID)
	productImage := new(model.ProductImage)

	productImage.ProductId = product.ID

	fileOne.Filename = uuid.New().String() + path.Ext(fileOne.Filename)
	fileTWo.Filename = uuid.New().String() + path.Ext(fileTWo.Filename)
	fileThree.Filename = uuid.New().String() + path.Ext(fileThree.Filename)

	// err = c.SaveFile(fileOne, "public/upload/"+fileOne.Filename)
	// if err != nil {
	// 	return c.Status(501).JSON(fiber.Map{"message": fileOne.Filename + " upload not completed successfull"})
	// }

	// err = c.SaveFile(fileTWo, "public/upload/"+fileTWo.Filename)
	// if err != nil {
	// 	return c.Status(501).JSON(fiber.Map{"message": fileTWo.Filename + " upload not completed successfull"})
	// }

	// err = c.SaveFile(fileThree, "public/upload/"+fileThree.Filename)
	// if err != nil {
	// 	return c.Status(501).JSON(fiber.Map{"message": fileThree.Filename + " upload not completed successfull"})
	// }

	url1, status1, _ := utils.UploadToBucket(fileOne)
	if !status1 {
		return c.Status(201).JSON(fiber.Map{
			"message": "img one upload failed",
		})
	}

	url2, status2, _ := utils.UploadToBucket(fileOne)
	if !status2 {
		return c.Status(201).JSON(fiber.Map{
			"message": "img two upload failed",
		})
	}
	url3, status3, _ := utils.UploadToBucket(fileOne)
	if !status3 {
		return c.Status(201).JSON(fiber.Map{
			"message": "img three upload failed",
		})
	}
	fileOne.Filename = url1
	fileTWo.Filename = url2
	fileTWo.Filename = url3

	productImage.ImageOne = fileOne.Filename
	productImage.ImgTwo = fileTWo.Filename
	productImage.ImgThree = fileThree.Filename

	res = db.Save(&productImage)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed in creating products",
		})

	}

	return c.Status(200).JSON(fiber.Map{
		"message": "product created",
	})
}

func UpdatePro(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)

	id := c.Params("id")
	fmt.Println(id)

	pImages := new(model.ProductImage)
	e := db.First(&pImages, "product_id=?", id)
	if e.Error != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	fmt.Println(pImages.ImageOne)

	fmt.Println(pImages.ImgTwo)
	fmt.Println(pImages.ImgThree)

	err := os.Remove("public/upload/" + pImages.ImageOne)
	if err != nil {
		fmt.Println("FIle not fount")
	}

	err = os.Remove("public/upload/" + pImages.ImgTwo)
	if err != nil {
		fmt.Println("FIle not fount")
	}

	err = os.Remove("public/upload/" + pImages.ImgThree)
	if err != nil {
		fmt.Println("FIle not fount")
	}
	// update... just deleted files

	pro := new(model.Products)
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	pro.ID = uint(u64)
	pro.Product_Name = c.FormValue("product_name")
	if price, err := strconv.ParseFloat(c.FormValue("price"), 64); err == nil {
		pro.Price = price
	}
	pro.Product_Category = c.FormValue("product_category")

	fileOne, err := c.FormFile("img_one")

	if err != nil {

		return c.Status(400).JSON(fiber.Map{
			"message": pImages.ImageOne + " upload error",
		})
	}

	err = c.SaveFile(fileOne, "public/upload/"+fileOne.Filename)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{"message": "file upload not completed successfull"})
	}

	fileTwo, err := c.FormFile("img_two")

	if err != nil {

		return c.Status(400).JSON(fiber.Map{
			"message": pImages.ImgTwo + " upload error",
		})
	}

	err = c.SaveFile(fileTwo, "public/upload/"+fileTwo.Filename)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{"message": "file upload not completed successfull"})
	}

	fileThree, err := c.FormFile("img_three")

	if err != nil {

		return c.Status(400).JSON(fiber.Map{
			"message": pImages.ImgThree + " upload error",
		})
	}

	err = c.SaveFile(fileThree, "public/upload/"+fileThree.Filename)
	if err != nil {
		return c.Status(501).JSON(fiber.Map{"message": "file upload not completed successfull"})
	}

	pImages.ImageOne = fileOne.Filename
	pImages.ImgTwo = fileTwo.Filename
	pImages.ImgThree = fileThree.Filename
	fmt.Println("---------------")
	fmt.Println(pImages.ImageOne)
	fmt.Println(pImages.ImgTwo)
	fmt.Println(pImages.ImgThree)

	db.Save(&pImages)
	db.Save(&pro)

	return c.Status(200).JSON(fiber.Map{"message": "updated"})

}

func DelProduct(c *fiber.Ctx) error {
	db := database.OpenDb()
	defer database.CloseDb(db)
	id := c.Params("id")
	fmt.Println(id)

	proImg := new(model.ProductImage)

	db.First(&proImg, "product_id = ?", id)

	if proImg.ID == '0' {
		return c.Status(404).JSON(fiber.Map{
			"message": "no match found",
		})
	}

	e := os.Remove("public/upload/" + proImg.ImageOne)
	if e != nil {
		fmt.Println("FIle not fount")
	}
	e = os.Remove("public/upload/" + proImg.ImgTwo)
	if e != nil {
		fmt.Println("FIle not fount")
	}
	e = os.Remove("public/upload/" + proImg.ImgThree)
	if e != nil {
		fmt.Println("FIle not fount")
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

func ViewUsers(c *fiber.Ctx) error {

	db := database.OpenDb()
	defer database.CloseDb(db)
	var user []model.Users
	err := db.Where("id > ?", 0).Find(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	var address []model.Address
	err = db.Where("id > ?", 0).Find(&address).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(address)

	infoWithoutAddress := utils.ExtractPersonalInfo(user)
	infowithPersonal := utils.ExtractAdresses(address)

	comb := utils.Combined(infoWithoutAddress, infowithPersonal)

	return c.Status(200).JSON(comb)

}
