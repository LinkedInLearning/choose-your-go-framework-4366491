type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

app.Post("/", func(c *fiber.Ctx) error {
	p := new(Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}
})









func (c *Ctx) ParamsParser(out interface{}) error

func (c *Ctx) QueryParser(out interface{}) error

func (c *Ctx) ReqHeaderParser(out interface{}) error

func (c *Ctx) Redirect(location string, status ...int) error

func (c *Ctx) Request() *fasthttp.Request

func (c *Ctx) Response() *fasthttp.Response

func (c *Ctx) SaveFile(fh *multipart.FileHeader, path string) error

func (c *Ctx) SendFile(file string, compress ...bool) error








app.Use(csrf.New())

app.Use("/expose/envvars", envvar.New())

app.Use(logger.New())

app.Use(basicauth.New(basicauth.Config{
   Users: map[string]string{
       "john":  "doe",
       "admin": "123456",
   },
}))

app.Use(cors.New())








// Handlers define a function to create hooks for Fiber.
type OnRouteHandler = func(Route) error

type OnNameHandler = OnRouteHandler

type OnGroupHandler = func(Group) error

type OnGroupNameHandler = OnGroupHandler

type OnListenHandler = func() error

type OnForkHandler = func(int) error

type OnShutdownHandler = OnListenHandler

type OnMountHandler = func(*App) error
