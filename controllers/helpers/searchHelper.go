package helpers

import (
	"context"
	"net/http"
	"strings"

	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/repositories"
)

func Search(ctx context.Context, query string) (int, interface{}) {
	queryList := strings.Split(query, "+")
	listMentor, err := repositories.GetMentor(ctx, queryList)
	if err != nil {
		return 500, outputs.InternalServerErrorOutput{Code: 500, Message: "Internal Server Error"}
	}

	if len(listMentor) == 0 {
		return http.StatusNotFound, outputs.NotFoundOutput{Code: 404, Message: "Data not found"}
	}

	output := outputs.WebResponse{}
	output.BaseOutput = outputs.BaseOutput{Code: 200, Message: "Success get mentor data"}

	output.Data = listMentor
	return 200, output
}
