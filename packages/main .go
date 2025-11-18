package packages

import (
	"fmt"

	"github.com/md-sanim-mia/packages/auth"
)

func main() {

	auth.LoginWithCredientials("coder-sanim", "sanim1234")

	session := auth.GetSession()

	fmt.Println("session", session)
}
