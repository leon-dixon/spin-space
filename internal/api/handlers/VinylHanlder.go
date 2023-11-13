package handlers

import (
	"net/http"
	"spin-space/internal/store"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VinylHandler struct {
	Store *store.VinylRecordStore
}

func NewVinylHandler(store *store.VinylRecordStore) *VinylHandler {
	return &VinylHandler{
		Store: store,
	}
}

func (VinylHandler *VinylHandler) GetVinyls(context *gin.Context) {
	vinyls, err := VinylHandler.Store.GetVinyls()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching vinyl records"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"vinyls": vinyls})
}

func (vinylHandler *VinylHandler) GetVinylByID(context *gin.Context) {
	strID := context.Param("id")
	uint64ID, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error converting string to uint"})
		return
	}
	vinyl, err := vinylHandler.Store.GetByID(uint(uint64ID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching vinyl record"})
		return
	}
	context.JSON(http.StatusOK, vinyl)
}
