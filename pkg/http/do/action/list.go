package action

import (
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/repository"
	"github.com/boombaw/go-ws-sia/pkg/util"
	"github.com/gofiber/fiber/v2"
)

// ListDOService is the responsible for listing all the AKM
func ListDOService(kdProdi string, semester string) model.MessageResponse {
	var resp model.MessageResponse

	repo := repository.NewDORepository()

	arg := repository.DOParams{
		KdProdi:  kdProdi,
		Semester: semester,
	}

	do, err := repo.List(arg)

	if err != nil {
		resp.Event = "error"
		resp.Message = err.Error()
	}

	data := fiber.Map{
		"list": do,
	}

	resp.Event = "list_do"
	resp.Message = string(util.ToJson(data))

	return resp
}
