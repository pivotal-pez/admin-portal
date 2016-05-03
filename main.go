package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/go-martini/martini"
	"github.com/jasonlvhit/gocron"
	"github.com/martini-contrib/render"
	"github.com/pivotal-pez/admin-portal/applications"
	"github.com/pivotal-pez/admin-portal/events"
	"github.com/pivotal-pez/admin-portal/users"
	cf "github.com/pivotal-pez/pezdispenser/cloudfoundryclient"
	"github.com/xchapter7x/cloudcontroller-client"
)

const (
	//SuccessStatus - success status of a call to the rest API
	SuccessStatus = 200
)

const (
	//AdminServiceName - userprovided service name for admin user of foundation
	AdminServiceName = "admin-user-information"
	//AdminURI - foundatio uri cups fieldname
	AdminURI = "cf-base-uri"
	//AdminUser - admin users cups fieldname
	AdminUser = "cf-user"
	//AdminPass - admin pass cups fieldname
	AdminPass = "cf-pass"
)

var (
	localCache ghettoCache
)

type ghettoCache struct {
	UserBlob   *users.UserAggregate
	AppsBlob   *applications.AppAggregate
	EventsBlob []cf.APIResponse
}

type cfAdminCreds struct {
	AdminURI  string
	AdminUser string
	AdminPass string
	LoginURI  string
	APIURI    string
}

type heritage struct {
	*ccclient.Client
	ccTarget string
}

func (s *heritage) CCTarget() string {
	return s.ccTarget
}

func main() {
	scheduler := gocron.NewScheduler()
	localLogger := log.New(os.Stdout, "[martini] ", 0)
	cfapp, _ := cfenv.Current()
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(martini.Static("public"))
	m.Get("/", func(params martini.Params, log *log.Logger, r render.Render) {
		r.HTML(SuccessStatus, "index", nil)
	})

	m.Get("/v1/info/apps", func(log *log.Logger, r render.Render) {
		r.JSON(200, localCache.AppsBlob)
	})

	m.Get("/v1/info/users", func(log *log.Logger, r render.Render) {
		r.JSON(200, localCache.UserBlob)
	})

	m.Get("/v1/info/events", func(log *log.Logger, r render.Render) {
		r.JSON(200, localCache.EventsBlob)
	})

	scheduler.Every(10).Seconds().Do(func() {
		FetchEventsInfo(cfapp, localLogger)
	})
	scheduler.Every(1).Minute().Do(func() {
		FetchUserInfo(cfapp, localLogger)
	})
	scheduler.Every(1).Minute().Do(func() {
		FetchAppsInfo(cfapp, localLogger)
	})
	scheduler.RunAll()
	go func() {
		scheduler.Start()
	}()
	m.Run()
}

//FetchEventsInfo - function to poll for events information in the CCDB via REST API
func FetchEventsInfo(cfapp *cfenv.App, log *log.Logger) {
	log.Println("running FetchEventsInfo cron")
	heritageClient := getHeritageClient(cfapp)
	cfclient := cf.NewCloudFoundryClient(heritageClient, log)
	eventSearch := new(events.EventSearch).Init(cfclient)
	eventSearch.CompileRecentEvents()
	localCache.EventsBlob = eventSearch.EventsBlob
}

//FetchAppsInfo - function to poll for app information in the CCDB via REST API
func FetchAppsInfo(cfapp *cfenv.App, log *log.Logger) {
	log.Println("running FetchAppsInfo cron")
	heritageClient := getHeritageClient(cfapp)
	cfclient := cf.NewCloudFoundryClient(heritageClient, log)
	appSearch := new(applications.AppSearch).Init(cfclient)
	appSearch.CompileAllApps()
	localCache.AppsBlob = appSearch.AppStats
}

//FetchUserInfo - function to poll for user information in the UAA via REST API
func FetchUserInfo(cfapp *cfenv.App, log *log.Logger) {
	log.Printf("running FetchUserInfo cron")
	heritageClient := getHeritageClient(cfapp)
	cfclient := cf.NewCloudFoundryClient(heritageClient, log)
	userSearch := new(users.UserSearch).Init(cfclient)
	userList, _ := userSearch.List("", "")
	localCache.UserBlob = new(users.UserAggregate)
	localCache.UserBlob.Compile(userList)
}

func getHeritageClient(cfapp *cfenv.App) (heritageClient *heritage) {
	creds := getAdminCreds(cfapp)
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
			AdminURI:  creds[AdminURI].(string),
			AdminUser: creds[AdminUser].(string),
			AdminPass: creds[AdminPass].(string),
			LoginURI:  fmt.Sprintf("https://%s.%s", "login", creds[AdminURI]),
			APIURI:    fmt.Sprintf("https://%s.%s", "api", creds[AdminURI]),
		}
	} else {
		panic(fmt.Sprintf("There is a problem with your required service binding %s: %s", AdminServiceName, err.Error()))
	}
	return
}
