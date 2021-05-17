package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func PostRotateMatrix(c *fiber.Ctx) error {
	var info Square
		response := Response {
			Status: false,
			Message: "",
			Data: nil,
		}

		err := c.BodyParser(&info)
		if err != nil {
			response.Message = err.Error()
			return c.JSON(response)
		}

		if (len(info.Data) == 0) {
			response.Message = "Array without data"
			c.JSON(response)
			return nil
		}

		if len(info.Data) != len(info.Data[0]) {
			response.Message = "Array is not NxN"
			c.JSON(response)
			return nil
		}

		size := len(info.Data)
		data := make([][]int, 0)

		for i := 0; i < size; i++ {
			tmp := make([]int, size)
			for j := 0; j < size; j++ {
				tmp[size - j - 1] = info.Data[size - j - 1][size - i - 1]
			}
			data = append(data, tmp)
		}

		response.Status = true
		response.Message = "The array was rotated successfully"
		response.Data = data

		return c.JSON(response)
}

type Square struct {
	Data [][]int `json:"data"`
}

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data [][]int `json:"data"`
}