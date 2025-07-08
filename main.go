package main

import (
	//"go_fiber/routes/api"
	//"log"
	//"github.com/gofiber/fiber/v2"

	"fmt"
	"go_fiber/crypt/util"
)

func main() {
	//app := fiber.New()

	//api.Routes(app)

	//log.Fatal(app.Listen(":3001"))

	plainText := "Hello, World!"
	fmt.Println("This is an original:", plainText)

	encrypted, err := util.GetAESEncrypted(plainText)

	if err != nil {
		fmt.Println("Error during encryption", err)
	}

	fmt.Println("This is an encrypted:", encrypted)

	decrypted, err := util.GetAESDecrypted(encrypted)

	if err != nil {
		fmt.Println("Error during decryption", err)
	}
	fmt.Println("This is a decrypted:", string(decrypted))

}
