package action

import (
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/repository"
	"github.com/boombaw/go-ws-sia/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func ListLulusan(kdProdi string, semester string) model.MessageResponse {
	var resp model.MessageResponse

	repo := repository.NewSyncLulusanRepository()

	arg := repository.SyncLulusanParams{
		KdProdi:  kdProdi,
		Semester: semester,
	}

	lulusan, err := repo.ListLulusan(arg)

	if err != nil {
		resp.Event = "error"
		resp.Message = err.Error()
	}

	data := fiber.Map{
		"list": lulusan,
	}

	resp.Event = "list_lulusan"
	resp.Message = string(util.ToJson(data))

	return resp
}
