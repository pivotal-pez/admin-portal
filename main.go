package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/pivotal-pez/admin-portal/users"
	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
	"github.com/xchapter7x/cloudcontroller-client"
)

const (
	SuccessStatus = 200
)

type heritage struct {
	*ccclient.Client
	ccTarget string
}

func (s *heritage) CCTarget() string {
	return s.ccTarget
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(martini.Static("public"))
	m.Get("/", func(params martini.Params, log *log.Logger, r render.Render) {
		r.HTML(SuccessStatus, "index", nil)
	})

	m.Get("/v1/info/apps", func(params martini.Params, log *log.Logger, r render.Render) {
		//http://apidocs.cloudfoundry.org/213/apps/list_all_apps.html
		//grab instance count, app count, buildpack and state aggregates
		r.HTML(SuccessStatus, "index", nil)
	})

	m.Get("/v1/info/users", func(log *log.Logger, r render.Render) {
		//grab total user count, okta users, uaa users, orphaned users
		baseURI := os.Getenv("CF_BASE_URI")
		user := os.Getenv("CF_USER")
		pass := os.Getenv("CF_PASS")
		loginURI := fmt.Sprintf("https://%s.%s", "login", baseURI)
		apiURI := fmt.Sprintf("https://%s.%s", "api", baseURI)
		heritageClient := &heritage{
			Client:   ccclient.New(loginURI, user, pass, new(http.Client)),
			ccTarget: apiURI,
		}
		heritageClient.Login()
		cfclient := cf.NewCloudFoundryClient(heritageClient, log)
		userSearch := new(users.UserSearch).Init(cfclient)
		userList, _ := userSearch.List("", "")
		userBlob := new(users.UserAggregate)
		userBlob.Compile(userList)
		r.JSON(200, userBlob)
	})

	m.Run()
}
