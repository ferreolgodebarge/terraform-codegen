package main

import(
	"fmt"
	"context"
	"net/http"
	"openapi"
	//"bytes"
)

func init_api(base_path string) (API *openapi.APIClient) {
	conf := openapi.NewConfiguration();
	conf.BasePath = base_path;
	API = openapi.NewAPIClient(conf);
	return;
}

func create_server(api *openapi.APIClient, name string, description string) (server openapi.Server, res *http.Response, err error){
	server_request := openapi.ServerRequest{Name: name, Description: description};
	context := context.Background();
	server, res, err = api.ServersApi.CreateServer(context, server_request);
	return;
}

func list_servers(api *openapi.APIClient) (servers []openapi.Server, res *http.Response, err error) {
	context := context.Background();
	servers, res, err = api.ServersApi.ListServers(context);
	return;
}

func get_server(api *openapi.APIClient, id string) (server openapi.Server, res *http.Response, err error) {
	context := context.Background();
	server, res, err = api.ServersApi.GetServer(context, id);
	return;
}

func update_server(api *openapi.APIClient, id string, name string, description string) (server openapi.Server, res *http.Response, err error) {
	context := context.Background();
	server_request := openapi.ServerRequest{Name: name, Description: description};
	server, res, err = api.ServersApi.GetServer(context, id);
}

func delete_server(api *openapi.APIClient, id string) (res *http.Response, err error) {
	context := context.Background();
	res, err = api.ServersApi.DeleteServer(context, id);
	return;
}

func main() {
	api := init_api("http://127.0.0.1:5000/api/v1");
	name := "test"
	description := "test api"

	//server, c, err := create_server(api, name, description);

	//servers, c, err := list_servers(api)

	//server, c, err := get_server(api, "48987a4d-d76e-461d-a5ae-8a4d662448f1")

	//server, c, err := update_server(api, id, name + "_update", description + " updated")

	//c, err := delete_server(api, "48987a4d-d76e-461d-a5ae-8a4d662448f1")

	//fmt.Println(server);
	fmt.Println(c);
	fmt.Println(err);
	//fmt.Println(c.Status);
	//l := list_servers(api);
	//fmt.Println(l.Status)
	//fmt.Println("Here is the body response :");
	//fmt.Printf("%+v\n",l);
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(l.Body)
	//fmt.Println(buf)
	//str := buf.String()
	//fmt.Println(str)
}
