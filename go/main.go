package main

import (
	"context"
  "encoding/json"
  "io/ioutil"
  "os"
	"fmt"
  "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"time"
)

type App struct {
  Name  string `json:"name"`
  Url   string `json:"url"`
  Icon  string `json:"icon"`
}

type Apps struct {
  Apps []App `json:"apps"`
}

func main() {
  filename := "/config/apps.json"
  checkFileExists(filename)

	for {
    time.Sleep(10 * time.Second)
		fmt.Println("Starting run")
    go updateJson(filename)
		fmt.Println("Stopping run")
  }

}

func updateJson(filename string){
	  containers := getContainers()

	  apps_empty := []App{}
		apps := Apps{apps_empty}

	  for _, container := range containers {
	    app := App{}
	    for key, value := range container.Labels{
	      if key == "sui.app.name" {
	        app.Name = value
	      }
	      if key == "sui.app.url" {
	        app.Url = value
	      }
	      if key == "sui.app.icon" {
	        app.Icon = value
	      }
	    }

			if (App{}) != app  {
	    	  apps.AddItem(app)
			}
	  }

	  writeJson(filename, apps)
}

func (apps *Apps) AddItem(app App) []App {
	apps.Apps = append(apps.Apps, app)
	return apps.Apps
}

func writeJson(filename string, apps Apps){
  dat, err := json.MarshalIndent(apps, "", "    ")
  if err != nil {
    panic(err)
  }

  err = ioutil.WriteFile(filename, dat, 0644)

}

func getContainers() []types.Container {
  ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

  containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
  if err != nil {
    panic(err)
  }

  return containers

}

func checkFileExists(filename string) error {
    _, err := os.Stat(filename)
        if os.IsNotExist(err) {
            _, err := os.Create(filename)
                if err != nil {
                    return err
                }
        }
        return nil
}
