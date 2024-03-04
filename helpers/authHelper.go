package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)



    func CheckUserType (c *gin.Context, role string)(err error){
        userType := c.GetString("user_type")
       
        err = nil
        if userType != role {
            err = errors.New("unauthorized to accesss this resource")
        }
        return nil
    }

func MatchUserTypeToUid(c *gin.Context, userId string) (err error){

    userType := c.GetString("user_type") 
    uid:= c.GetString("uid")
    err = nil

    if userType =="USER" && uid != userId {
            err = errors.New("unauthorized to access this resources")
            return err
    }

    err = CheckUserType(c, userType)

    return err
}