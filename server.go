package main


import (
  "apiPoc/model"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "time"
  "fmt"
)

func main() {
  m := martini.Classic()
  db, err := gorm.Open("sqlite3", "./gorm.db")
  if err != nil{
    panic("Could not open db")
  }
  migrateDb(db)
  
  m.Use(render.Renderer()) //HTML & JSON renderer middleware

  m.Get("/", func() string {
    return "Hello world!"
  })
     
  m.Get("/error", func() (int, string) {
    return 400, "i'm a teapot" // HTTP 418 : "i'm a teapot"
  })

  m.Get("/error500", func() (int, string) {
    return 500, "i'm a teapot" // HTTP 418 : "i'm a teapot"
  })

  m.Get("/api/user", func(r render.Render) {
    user := ormTest(db)
    r.JSON(201, user)   
  })

  m.Get("/user", func(r render.Render) {
    users := ormTest(db)
    r.HTML(200, "user", users)
  })

  m.Get("/users", func(r render.Render) {
    users := GetAllUsers(db)
    r.HTML(200, "createUser", users)
  })

  m.Post("/users", binding.Form(model.User{}), func(user model.User, r render.Render) {
    db.Create(&user)
    fmt.Printf("%v\n", user)
    users := GetAllUsers(db)
    r.HTML(200, "createUser", users)
  })

  m.Run()
}

func GetAllUsers(db gorm.DB)[]model.User{
  users := []model.User{}
  db.Find(&users)
  return users
}

func migrateDb(db gorm.DB){
  // Create table
  db.CreateTable(&model.User{})
  db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})

  // Automating Migration
  db.AutoMigrate(&model.User{})
  db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
  // Feel free to change your struct, AutoMigrate will keep your database up-to-date.
  // AutoMigrate will ONLY add *new columns* and *new indexes*,
  // WON'T update current column's type or delete unused columns, to protect your data.
  // If the table is not existing, AutoMigrate will create the table automatically.
}
   
func ormTest(db gorm.DB) []model.User{
      user := model.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
      db.NewRecord(user) // => returns `true` if primary key is blank
      db.Create(&user)
      db.NewRecord(user) // => return `false` after `user` created
      users := []model.User{}
      db.First(&users)
      return users
}