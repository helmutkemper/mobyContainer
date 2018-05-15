package marketPlaceProcy

import (
  "golang.org/x/net/context"
  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types"
)

type ContainerExecAttachDataIn struct{
  Id          string
  Options     types.ExecStartCheck
}

func ContainerExecAttach(w ProxyResponseWriter, r *ProxyRequest) {
  output := JSonOutStt{}

  var inDataLStt ContainerExecAttachDataIn
  var err, id = Input(r, &inDataLStt)

  if err != nil {
    output.ToOutput( 0, err, []int{}, w )
    return
  }

  if id != "" {
    inDataLStt.Id = id
  }

  ctx := context.Background()
  cli, err := client.NewEnvClient()
  if err != nil {
    output.ToOutput( 0, err, []int{}, w )
    return
  }

  execResp, err := cli.ContainerExecAttach(ctx, inDataLStt.Id, inDataLStt.Options)
  if err != nil {
    output.ToOutput( 0, err, []int{}, w )
    return
  }

  output.ToOutput( 1, err, []types.HijackedResponse{ execResp }, w )
}
