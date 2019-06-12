package security

import (
  "fmt"
  "errors"
  "github.com/gin-gonic/gin"
  data "github.com/estebanborai/songs-share-server/data"
  models "github.com/estebanborai/songs-share-server/models"
)

func ValidateUser(c *gin.Context, userName string, plainPassword string) (models.User, error) {
  var user models.User
  var userSecret models.UserSecret

  db := data.Connection(c)
  
  if userResult := db.Where(&models.User{ UserName: userName }).First(&user); userResult.Error == nil {
    if credsResult := db.Where(&models.UserSecret{ UserId: user.Id }).First(&userSecret); credsResult.Error == nil {
      if ValidatePassword(userSecret.Hash, plainPassword) {
        return user, nil
      } else {
        return user, errors.New("Invalid Password")
      }
    }

    return user, errors.New("Cannot validate User")
  } else {
    return user, errors.New(fmt.Sprintf("%s doesn't exists", userName))
  }
}
