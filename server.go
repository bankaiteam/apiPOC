package main


import (
  "apiPoc/model"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "time"
  "encoding/json"
)
import "fmt"

func main() {
  m := martini.Classic()
  
  m.Use(render.Renderer()) //HTML & JSON renderer middleware

  m.Get("/", func() string {
    return "Hello world!"
  })
  
  m.Get("/error", func() (int, string) {
    return 400, "i'm a teapot" // HTTP 418 : "i'm a teapot"
  })

  m.Get("/user", func(r render.Render) {
    user := ormTest()
    userJson, err := json.Marshal(user)
    fmt.Printf("%s\n", userJson)
    if err == nil {
      r.JSON(201, user)
    }
  })

  m.Run()
}


func migrateDb(db gorm.DB){
  // Create table
  db.CreateTable(&model.User{})
  db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})

  // Drop table
  db.DropTable(&model.User{})

  // Automating Migration
  db.AutoMigrate(&model.User{})
  db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
  // Feel free to change your struct, AutoMigrate will keep your database up-to-date.
  // AutoMigrate will ONLY add *new columns* and *new indexes*,
  // WON'T update current column's type or delete unused columns, to protect your data.
  // If the table is not existing, AutoMigrate will create the table automatically.
}

func ormTest() model.User{
    db, err := gorm.Open("sqlite3", "./gorm.db")
    migrateDb(db)
    if err == nil {
      user := model.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
      db.NewRecord(user) // => returns `true` if primary key is blank
      db.Create(&user)
      db.NewRecord(user) // => return `false` after `user` created
      user2 := model.User{}
      db.First(&user2)
      return user2
    }
    panic("Not implemented")
}