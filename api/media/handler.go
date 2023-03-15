package media

import (
	db "developer-network/database"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var allmedia []db.Media

func GetAllMedia(c *gin.Context) {
	err := db.DB.Model(&allmedia).Select()
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, allmedia)

}

func GetMediaById(c *gin.Context) {
	id := c.Param("id")
	media, err := mediaById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "media with id doesn't exists"})
		return
	}

	c.IndentedJSON(http.StatusOK, media)

}

func mediaById(id string) (*db.Media, error) {
	ids, _ := strconv.ParseUint(id, 10, 64)
	media := &db.Media{ID: ids}
	err := db.DB.Model(media).WherePK().Select()
	if err != nil {
		return nil, errors.New("media not found")
	}
	return media, nil
}

func CreateNewMedia(c *gin.Context) {
	var newMedia db.Media

	err := c.BindJSON(&newMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	_, err = db.DB.Model(&newMedia).Insert()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//reading the media files stored in /tmp/media directory
	files, fileErr := os.ReadDir("/tmp/media")
	if fileErr != nil {
		panic(fileErr)
	}

	for _, file := range files {
		log.Println(file.Name(), file.Type().IsRegular())
		if !file.IsDir() {
			//checking whether directory /media/user exists or not
			_, dirErr := os.Stat("./mediaData/user/")
			if os.IsNotExist(dirErr) {
				//creating directory for storing media data
				dirErr := os.MkdirAll("./mediaData/user/", 0755)
				if dirErr != nil {
					log.Fatal(dirErr)
				}
			}
			//checking whether directory /media/user/USERID exists or not
			_, err := os.Stat("./mediaData/user/" + strconv.FormatUint(newMedia.ID, 10))
			if os.IsNotExist(err) {
				//creating directory for storing media data
				errDir := os.Mkdir(strconv.FormatUint(newMedia.ID, 10), 0755)
				if errDir != nil {
					log.Fatal(err)
				}
				//moving directory /USERID to /mediaData/user/
				errMov := os.Rename((strconv.FormatUint(newMedia.ID, 10)), "./mediaData/user/"+(strconv.FormatUint(newMedia.ID, 10)))
				if errMov != nil {
					log.Fatal(errMov)
				}
			}
			//moving media files from /tmp/media to /mediaData/user/USERID
			errMov := os.Rename("/tmp/media/"+file.Name(), "./mediaData/user/"+strconv.FormatUint(newMedia.ID, 10)+"/"+file.Name())
			if errMov != nil {
				log.Fatal(errMov)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Media created successfully"})
}

func UpdateMedia(c *gin.Context) {
	id := c.Param("id")

	mediaById, err := mediaById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "media with id doesn't exists"})
		return
	}

	// Get the media from the database
	updateMedia := &db.Media{ID: mediaById.ID}
	err = db.DB.Model(updateMedia).WherePK().Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the media fields
	err = c.BindJSON(updateMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated media back to the database
	_, err = db.DB.Model(updateMedia).WherePK().Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//working with files
	//reading the media files stored in /tmp/media directory
	files, fileErr := os.ReadDir("/tmp/media")
	if fileErr != nil {
		panic(fileErr)
	}

	for _, file := range files {
		log.Println(file.Name(), file.Type().IsRegular())
		if !file.IsDir() {
			//checking whether directory /media/user exists or not
			_, dirErr := os.Stat("./mediaData/user/")
			if os.IsNotExist(dirErr) {
				//creating directory for storing media data
				dirErr := os.MkdirAll("./mediaData/user/", 0755)
				if dirErr != nil {
					log.Fatal(dirErr)
				}
			}

				//checking whether directory /media/user/USERID exists or not
				_, err := os.Stat("./mediaData/user/" + strconv.FormatUint(updateMedia.ID, 10))
				if os.IsNotExist(err) {
					//creating directory for storing media data
					errDir := os.Mkdir(strconv.FormatUint(updateMedia.ID, 10), 0755)
					if errDir != nil {
						log.Fatal(err)
					}

					//moving directory /USERID to /mediaData/user/
					errMov := os.Rename((strconv.FormatUint(updateMedia.ID, 10)), "./mediaData/user/"+(strconv.FormatUint(updateMedia.ID, 10)))
					if errMov != nil {
						log.Fatal(errMov)
					}
				}
				//moving media files from /tmp/media to /mediaData/user/USERID
				errMov := os.Rename("/tmp/media/"+file.Name(), "./mediaData/user/"+strconv.FormatUint(updateMedia.ID, 10)+"/"+file.Name())
				if errMov != nil {
					log.Fatal(errMov)
				}
			
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Media updated successfully"})

}

func DeleteMediaById(c *gin.Context) {
	id := c.Param("id")

	mediaById, err := mediaById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "media with id doesn't exists"})
		return
	}

	delMedia := &db.Media{ID: mediaById.ID}

	_, delerror := db.DB.Model(delMedia).WherePK().Delete()
	if delerror != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Media deleted successfully"})
	// os.RemoveAll("./mediaData/user/" + id)
}
