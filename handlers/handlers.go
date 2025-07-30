package handlers

import (
	"github.com/gofiber/fiber/v2"
	"main/models"
	"main/services"
	"main/utils"
	"net/http"
)

func GetAllItems(c *fiber.Ctx) error {
	var items []models.Item = services.GetAllItems()
	return c.JSON(models.Response[[]models.Item]{
		Success: true,
		Message: "All items Data",
		Data:    items,
	})
}

func GetItemByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	item, err := services.GetItemByID(idParam)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "Item found",
		Data:    item,
	})
}

func CreateItem(c *fiber.Ctx) error {
	isValid, err := utils.CheckToken(c)
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	var itemInput *models.ItemRequest = new(models.ItemRequest)

	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	errors := itemInput.ValidateStruct()
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}
	var createdItem models.Item = services.CreateItem(*itemInput)
	return c.Status(http.StatusCreated).JSON(models.Response[models.Item]{
		Success: true,
		Message: "item created",
		Data:    createdItem,
	})

}

func UpdateItem(c *fiber.Ctx) error {
	isValid, err := utils.CheckToken(c)
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	var inputItem *models.ItemRequest = new(models.ItemRequest)

	if err := c.BodyParser(inputItem); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	errors := inputItem.ValidateStruct()
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}
	idPar := c.Params("id")
	updatedItem, err := services.UpdateItem(*inputItem, idPar)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "item updated",
		Data:    updatedItem,
	})
}

func DeleteItem(c *fiber.Ctx) error {
	isValid, err := utils.CheckToken(c)
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	idPar := c.Params("id")
	result := services.DeleteItem(idPar)
	if result {
		return c.JSON(models.Response[any]{
			Success: true,
			Message: "item deleted",
		})
	}
	return c.Status(http.StatusNotFound).JSON(models.Response[any]{
		Success: false,
		Message: "failed to delete item",
	})
}
