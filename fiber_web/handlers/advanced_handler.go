package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func PathParams(c *fiber.Ctx) error {
	id := c.Params("id")
	name := c.Params("name")
	return c.JSON(fiber.Map{
		"id":   id,
		"name": name,
	})
}

func WildcardParams(c *fiber.Ctx) error {
	filepath := c.Params("*")
	return c.JSON(fiber.Map{
		"filepath": filepath,
	})
}

func QueryArray(c *fiber.Ctx) error {
	tags := c.Query("tags")
	return c.JSON(fiber.Map{
		"tags": tags,
	})
}

func QueryMap(c *fiber.Ctx) error {
	query := c.Queries()
	return c.JSON(query)
}

func PostArray(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	values := form.Value["items"]
	return c.JSON(fiber.Map{
		"items": values,
	})
}

func PostMap(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(form.Value)
}

func ClientIP(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"ip": c.IP(),
	})
}

func RequestInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"method":      c.Method(),
		"path":        c.Path(),
		"protocol":    c.Protocol(),
		"host":        c.Hostname(),
		"ip":          c.IP(),
		"ips":         c.IPs(),
		"originalURL": c.OriginalURL(),
		"baseURL":     c.BaseURL(),
	})
}

func BindQuery(c *fiber.Ctx) error {
	type QueryParams struct {
		Page  int    `query:"page"`
		Limit int    `query:"limit"`
		Sort  string `query:"sort"`
	}
	var params QueryParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(params)
}

func BindForm(c *fiber.Ctx) error {
	type FormData struct {
		Name  string `form:"name"`
		Email string `form:"email"`
		Age   int    `form:"age"`
	}
	var data FormData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(data)
}

func BindURI(c *fiber.Ctx) error {
	type URIParams struct {
		ID   int    `params:"id"`
		Name string `params:"name"`
	}
	var params URIParams
	if err := c.ParamsParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(params)
}

func BindHeader(c *fiber.Ctx) error {
	type HeaderData struct {
		Authorization string `reqHeader:"authorization"`
		ContentType   string `reqHeader:"content-type"`
		UserAgent     string `reqHeader:"user-agent"`
	}
	var headers HeaderData
	if err := c.ReqHeaderParser(&headers); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(headers)
}
