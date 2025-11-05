package handlers

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func UploadSingleFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"filename": file.Filename,
		"size":     file.Size,
		"header":   file.Header,
	})
}

func UploadMultipleFiles(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	files := form.File["files"]

	var fileInfo []fiber.Map
	for _, file := range files {
		fileInfo = append(fileInfo, fiber.Map{
			"filename": file.Filename,
			"size":     file.Size,
		})
	}
	return c.JSON(fiber.Map{
		"count": len(fileInfo),
		"files": fileInfo,
	})
}

func ProcessForm(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	age := c.FormValue("age", "0")

	return c.JSON(fiber.Map{
		"name":  name,
		"email": email,
		"age":   age,
	})
}

func SaveUploadedFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	filename := filepath.Base(file.Filename)
	dst := "./fiber_web/uploads/" + filename

	if err := c.SaveFile(file, dst); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":  "File uploaded successfully",
		"filename": filename,
		"path":     dst,
	})
}
