package marketPlaceProcy

import (
  "golang.org/x/net/context"
  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types"
)

func ImageWebList(w ProxyResponseWriter, r *ProxyRequest){
  output := JSonOutStt{}

  cli, err := client.NewEnvClient()
  if err != nil {
    output.ToOutput(0, err, []int{}, w)
  }

  images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
  if err != nil {
    output.ToOutput(0, err, []int{}, w)
  }

  output.ToOutput( len(images), err, images, w )
}