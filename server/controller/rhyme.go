package controller

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ryomak/rhymer/server/util/api"
	"github.com/jinzhu/gorm"
	mg "github.com/ryomak/rhymer/server/my_graphqls"
)

type RhymeServer struct {
	DB *gorm.DB
}


func (rs RhymeServer)RhymeHandler(w http.ResponseWriter, r *http.Request) {
	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	query := bufBody.String()

	result, err := mg.ExecuteQuery(query, mg.GetWordSchema(),rs.DB)
	if err != nil {
		api.JsonErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w = api.WriteJsonHeader(w, http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
