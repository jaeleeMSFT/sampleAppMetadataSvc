package main

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "net/mail"  
  "github.com/hashicorp/go-version"
    
  "github.com/emicklei/go-restful/v3"

)

type App struct {
  title string
  ver version.Version                   // import from https://golangrepo.com/repo/hashicorp-go-version-go-utilities
  maintainers []mail.Address    // https://pkg.go.dev/net/mail#Address
  company string
  website url.URL               // https://pkg.go.dev/net/url
  source url.URL                // https://pkg.go.dev/net/url
  license string
  description string
}

// ordered list of creation TimeBuckets
// Each TimeBucket contains an expandable array of App objects(ptr)
//  [2020-01-23T08:23:59] --> { "myApp v1.0.0", "Foo v0.0.1", "Apple v0.0.3" }
//  [2020-01-23T08:24:20] --> { "myApp v2.0.0", "Foo v0.0.2" }
//  [2020-12-04T17:01:34] --> { "myApp v3.0.0" }
//  [2021-04-12T23:56:55] --> { "myApp v3.5.0", "Orange v0.0.1" }
//  [2021-06-06T02:11:01] --> { "myApp v4.0.2", "Apple v2.0.0" }
//  [2020-12-04T17:01:34] --> { "myApp v5.0.0" }
//  ......


func returnApps(req *restful.Request, resp *restful.Response) {
  fmt.Println("request handled:  returnApps")

}

func returnSingleVersion(req *restful.Request, resp *restful.Response) {
  title := req.PathParameter("title")
  version := req.PathParameter("version")
  
  fmt.Println("request handled:  returnSingleVersion", title, version)
  
}

func returnVersionsOfApp(req *restful.Request, resp *restful.Response) {
  title := req.PathParameter("title")
  
  fmt.Println("request handled:  returnVersionsOfApp", title)
  
}

func createOrUpdateAppVersion(req *restful.Request, resp *restful.Response) {
  title := req.PathParameter("title")
  version := req.PathParameter("version")
  
  fmt.Println("request handled:  createOrUpdateAppVersion", title, version)
  
}

func deleteSingleVersion(req *restful.Request, resp *restful.Response) {
  title := req.PathParameter("title")
  version := req.PathParameter("version")
  
  fmt.Println("request handled:  deleteSingleVersion", title, version)
  
}

func deleteAllVersionsOfApp(req *restful.Request, resp *restful.Response) {
  title := req.PathParameter("title")

  fmt.Println("request handled:  deleteAllVersionsOfApp", title)
  
}



func main() { 

  fmt.Println("starting REST server...")
  
  ws := new(restful.WebService)
  
  ws.Route(ws.GET("/apps").To(returnApps))
  
  ws.Route(ws.GET("/apps/{title}").To(returnVersionsOfApp))
  ws.Route(ws.DELETE("/apps/{title}").To(deleteAllVersionsOfApp))

  ws.Route(ws.PUT("/apps/{title}/{version}").To(createOrUpdateAppVersion))
  ws.Route(ws.GET("/apps/{title}/{version}").To(returnSingleVersion))
  ws.Route(ws.DELETE("/apps/{title}/{version}").To(deleteSingleVersion))
   
  restful.DefaultResponseContentType(restful.MIME_JSON)
  restful.RegisterEntityAccessor(MediaTypeApplicationYaml, NewYamlReaderWriter(MediaTypeApplicationYaml))
  
  restful.Add(ws)  
  
  log.Fatal(http.ListenAndServe(":8080", nil))
 
}
