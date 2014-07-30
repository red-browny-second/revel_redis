package controllers

import(
  "revel_redis/app/models"
  "github.com/revel/revel"
  "fmt"
) 

type App struct {
	ModelController
}

func (c App) New() revel.Result {
  greeting := "Hello, Revel Redis Sample Appp!!"
	return c.Render(greeting)
}

func (c App) Index(myName string) revel.Result {
  fmt.Println("In instert data.")
  fmt.Println("Values--->", myName)
	vals := []string{myName}
	for _, v := range vals {
	    err	:= client.Rpush("username", []byte(v))
		fmt.Println("Redis Error:", err)

	}
  
  
  fmt.Println("In get data.")
	usernames,_ := client.Lrange("username", 0, 10000)
	fmt.Println("User--->", usernames)
  
  users := []models.User{}
  for i, v := range usernames {
	  println("Name", i,":",string(v))

	  user := models.User{}
	  user.Username = string(v)
	  users = append(users, user)
  }
	
  fmt.Println("All users>>>", users)
	return c.Render(users)
}

