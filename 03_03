



app.Get("/", func(c *fiber.Ctx) error {
   return fiber.NewError(782, "Custom error message")
})













app := fiber.New(fiber.Config{
   Prefork: true,
})

if !fiber.IsChild() {
   fmt.Println("I'm the parent process")
} else {
   fmt.Println("I'm a child process")
}






func main() {
   app := fiber.New()
   micro := fiber.New()
   app.Mount("/john", micro) 

   micro.Get("/doe", func(c *fiber.Ctx) error {
       return c.SendStatus(fiber.StatusOK)
   })

   log.Fatal(app.Listen(":3000"))
}





func main() {
   app := fiber.New()
  
   api := app.Group("/api", handler) 
  
   v1 := api.Group("/v1", handler)  
   v1.Get("/list", handler)          
   v1.Get("/user", handler)
  
   log.Fatal(app.Listen(":3000"))
  } 
  