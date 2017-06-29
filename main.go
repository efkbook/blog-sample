package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/efkbook/blog-sample/controller"
	"github.com/efkbook/blog-sample/db"

	csrf "github.com/utrack/gin-csrf"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	elastic "gopkg.in/olivere/elastic.v5"
)

// Server is whole server implementation for this blog app.
// This holds database connection and router settings based on gin.
type Server struct {
	db     *sql.DB
	Engine *gin.Engine
	es     *elastic.Client
	fluent *fluent.Fluent
}

// Close makes the database connection to close.
func (s *Server) Close() error {
	return s.db.Close()
}

// Init initialize server state. Connecting to database, compiling templates,
// and settings router.
func (s *Server) Init(dbconf, env, esurl, fluentHost string) {
	cs, err := db.NewConfigsFromFile(dbconf)
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	db, err := cs.Open(env)
	if err != nil {
		log.Fatalf("db initialization failed: %s", err)
	}
	s.db = db

	client, err := elastic.NewClient(
		elastic.SetURL(esurl),
		elastic.SetMaxRetries(5),
		elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("initialize Elasticsearch client failed: %s", err)
	}
	s.es = client

	for {
		logger, err := fluent.New(fluent.Config{
			FluentHost: fluentHost,
		})
		if err != nil {
			log.Printf("initialize fluentd client failed: %s", err)
			time.Sleep(time.Second)
			continue
		}
		s.fluent = logger
		break
	}

	// NOTE: define helper func to use from templates here.
	t := template.Must(template.New("").Funcs(template.FuncMap{
		"LoggedIn":    controller.LoggedIn,
		"CurrentName": controller.CurrentName,
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
	}).ParseGlob("templates/*"))
	s.Engine.SetHTMLTemplate(t)

	store := sessions.NewCookieStore([]byte("secretkey"))
	s.Engine.Use(sessions.Sessions("blogsession", store))
	s.Engine.Use(csrf.Middleware(csrf.Options{
		Secret: "secretkey",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismach")
			c.Abort()
		},
	}))

	s.Route()
}

// New returns server object.
func New() *Server {
	r := gin.Default()
	return &Server{Engine: r}
}

// Run starts running http server.
func (s *Server) Run(addr ...string) {
	s.Engine.Run(addr...)
}

// Route setting router for this blog.
func (s *Server) Route() {
	article := &controller.Article{DB: s.db, ES: s.es, Fluent: s.fluent}
	user := &controller.User{DB: s.db}

	auth := s.Engine.Group("/")
	auth.Use(controller.AuthRequired())
	{
		auth.GET("/authtest", func(c *gin.Context) {
			c.String(200, "you're authed")
		})
		auth.GET("/new", func(c *gin.Context) {
			c.HTML(200, "new.tmpl", gin.H{
				"title":   "New: go-blog",
				"csrf":    csrf.GetToken(c),
				"context": c,
			})
		})
		auth.GET("/article/:id/edit", article.Edit)
		auth.POST("/save", article.Save)
		auth.POST("/delete", article.Delete)
		auth.GET("/logout", func(c *gin.Context) {
			c.HTML(http.StatusOK, "logout.tmpl", gin.H{
				"csrf":    csrf.GetToken(c),
				"context": c,
			})
		})
		auth.POST("/logout", user.Logout)
	}

	s.Engine.GET("/", article.Root)
	s.Engine.GET("/search", article.Search)
	s.Engine.GET("/article/:id", article.Get)
	s.Engine.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{
			"csrf": csrf.GetToken(c),
		})
	})
	s.Engine.POST("/signup", user.SignUp)
	s.Engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"csrf": csrf.GetToken(c),
		})
	})
	s.Engine.POST("/login", user.Login)

	s.Engine.Static("/static", "static")
}

func main() {
	var (
		dbconf     = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env        = flag.String("env", "development", "application envirionment (production, development etc.)")
		esurl      = flag.String("esurl", "http://elasticsearch:9200", "Elasticsearch node URL")
		fluentHost = flag.String("fluenthost", "fluentd", "Elasticsearch node URL")
	)
	flag.Parse()
	b := New()
	b.Init(*dbconf, *env, *esurl, *fluentHost)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	b.Run(":" + port)
}
