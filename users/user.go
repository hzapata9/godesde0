package users

import (
	"fmt"
	"time"

	"github.com/hzapata9/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)

	fmt.Println("Usuario nuevo: ", u)
}
