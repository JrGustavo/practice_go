package users

import (
	"fmt"
	"github.com/JrGustavo/practice_go/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Junior", time.Now(), true)
	fmt.Println(u)

}
