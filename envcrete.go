package main

import (
	"fmt"
	"os"
)

func main() {
	for key, value := range map[string]string{"DB_HOST": "192.168.31.99", "DB_PORT": "3306", "DB_USERNAME": "root", "DB_PASSWORD": "Qaqaqa00.0", "DB_CHARSET": "utf8mb4", "DB_NAME": "freelink"} {
		if err := os.Setenv(key, value); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
