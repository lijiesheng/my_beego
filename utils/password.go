package utils

import (
	"strconv"
	"strings"
)

const (
	saltSize = 16
	delimiter  =  "|"
	stretchingPassword = 500
	saltLocalSecret = "wWnN&^bnmIIIEbW**WL"
)

func trimSaltHash(hash string) map[string]string {
	str := strings.Split(hash, delimiter)
	return map[string]string {
		"salt_secret" :     str[0],
		"interation_string" : str[1],
		"hash":              str[2],
		"salt":              str[3],
	}
}

// todo 不懂
func PasswordVerify(hasing string, pass string)(bool, error){
	data := trimSaltHash(hasing)

	interation, _ := strconv.ParseInt(data["interation_string"], 10, 64)

}


func hash(pass string, saltSecret string, salt, interation int64) (string, error) {
	var passSalt = saltSecret + pass + salt + saltSecret + pass + salt + pass + pass + salt

}
