package action

import (
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/repository"
	"github.com/boombaw/go-ws-sia/pkg/util"
	"github.com/gofiber/fiber/v2"
)

// ListAKMService is the responsible for listing all the AKM
func ListAKMService(kdProdi string, semester string) model.MessageResponse {
	var resp model.MessageResponse

	repo := repository.NewAKMRepository()

	arg := repository.AkmParams{
		KdProdi:  kdProdi,
		Semester: semester,
	}

	akm, err := repo.List(arg)

	if err != nil {
		resp.Event = "error"
		resp.Message = err.Error()
	}

	data := fiber.Map{
		"list": akm,
	}

	resp.Event = "list_akm"
	resp.Message = string(util.ToJson(data))

	return resp
}

func ListAKMNAService(kdProdi string, semester string) model.MessageResponse {
	var resp model.MessageResponse

	repo := repository.NewAKMRepository()

	arg := repository.AkmParams{
		KdProdi:  kdProdi,
		Semester: semester,
	}

	akm, err := repo.ListNA(arg)

	if err != nil {
		resp.Event = "error"
		resp.Message = err.Error()
	}

	data := fiber.Map{
		"list": akm,
	}

	resp.Event = "list_akm"
	resp.Message = string(util.ToJson(data))

	return resp
}
