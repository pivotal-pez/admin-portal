package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cloudfoundry-community/go-cfenv"
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
	cfapp, _ := cfenv.Current()
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
		creds := getAdminCreds(cfapp)
		heritageClient := getHeritageClient(creds)
		cfclient := cf.NewCloudFoundryClient(heritageClient, log)
		userSearch := new(users.UserSearch).Init(cfclient)
		userList, _ := userSearch.List("", "")
		userBlob := new(users.UserAggregate)
		userBlob.Compile(userList)
		r.JSON(200, userBlob)
	})

	m.Run()
}

const (
	AdminServiceName = "admin-user-information"
	AdminURI         = "cf-base-uri"
	AdminUser        = "cf-user"
	AdminPass        = "cf-pass"
)

type cfAdminCreds struct {
	AdminURI  string
	AdminUser string
	AdminPass string
	LoginURI  string
	APIURI    string
}

func getHeritageClient(creds *cfAdminCreds) (heritageClient *heritage) {
	heritageClient = &heritage{
		Client:   ccclient.New(creds.LoginURI, creds.AdminUser, creds.AdminPass, new(http.Client)),
		ccTarget: creds.APIURI,
	}
	heritageClient.Login()
	return
}

func getAdminCreds(cfapp *cfenv.App) (adminCreds *cfAdminCreds) {
	if cfAdminService, err := cfapp.Services.WithName(AdminServiceName); err == nil {
		creds := cfAdminService.Credentials
		adminCreds = &cfAdminCreds{
			AdminURI:  creds[AdminURI],
			AdminUser: creds[AdminUser],
			AdminPass: creds[AdminPass],
			LoginURI:  fmt.Sprintf("https://%s.%s", "login", creds[AdminURI]),
			APIURI:    fmt.Sprintf("https://%s.%s", "api", creds[AdminURI]),
		}
	} else {
		panic(fmt.Sprintf("There is a problem with your required service binding %s: %s", AdminServiceName, err.Error()))
	}
	return
}
