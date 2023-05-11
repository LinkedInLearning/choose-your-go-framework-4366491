r := gin.Default()

r.GET("/api/users", func(c *gin.Context) {
    // Handle GET request for /api/users
})

r.POST("/api/users", func(c *gin.Context) {
    // Handle POST request for /api/users
})







r.LoadHTMLGlob("templates/*")

r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "Welcome to my website",
    })
})

r.Static("/static", "./static")










var name = flag.String("name", "world", "The name to greet")

func main() {
    flag.Parse()

    cmd := exec.Command("echo", "Hello, "+*name+"!")
    cmd.Stdout = os.Stdout
    cmd.Run()
}


