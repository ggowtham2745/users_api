package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

type TtsUser struct {
	tableName  struct{} `pg:"TTS_USER,alias:u"`
	ID         string   `pg:"ID" json:"Id"`
	LOC_ID     string   `pg:"LOC_ID" json:"LocId,omitempty"`
	LOC_TYPE   string   `pg:"LOC_TYPE" json:"LocType,omitempty"`
	FIRST_NAME string   `pg:"FIRST_NAME" json:"FirstName,omitempty"`
	LAST_NAME  string   `pg:"LAST_NAME" json:"LastName,omitempty"`
	CREATED_ON string   `pg:"CREATED_ON" json:"CreatedOn,omitempty"`
	UPDATED_ON string   `pg:"UPDATED_ON" json:"UpdatedOn,omitempty"`
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")

}

func GetUserList(c *gin.Context) {
	var ttsusers []TtsUser

	err := dbConnect.Model(&ttsusers).Select()

	if err != nil {
		log.Printf("Error while getting all ttsUsers, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All ttsUsers",
		"data":    ttsusers,
	})
	return

}
