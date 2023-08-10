package logic_test

import (
	apiModels "RegressionHelperAPI/internal/regression/api/models"
	"RegressionHelperAPI/internal/regression/logic/models"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
)

func StartDockerContainer(t *testing.T) (*client.Client, string) {
	ctx := context.Background()
	cli, newClientErr := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if newClientErr != nil {
		t.Fatalf("Create docker client error: %v\n", newClientErr)
	}
	defer cli.Close()
	log.Println("Create client complete")

	buildCtx, openCtxErr := archive.TarWithOptions("./../", &archive.TarOptions{})
	if openCtxErr != nil {
		t.Fatalf("Archive docker context error: %v\n", openCtxErr)
	}

	buildOptions := types.ImageBuildOptions{
		Tags:       []string{"regression"},
		Dockerfile: "./production/services/regression/Dockerfile", // optional, is the default
	}
	buildResponse, buildErr := cli.ImageBuild(context.Background(), buildCtx, buildOptions)
	if buildErr != nil {
		t.Fatalf("Build docker image error: %v\n", buildErr)
	}
	defer buildResponse.Body.Close()
	log.Println("Image build complete")

	conf := &container.Config{
		Image: "regression",
		ExposedPorts: nat.PortSet{
			"8080/tcp": struct{}{},
		}}
	hostConf := &container.HostConfig{
		PortBindings: nat.PortMap{
			"8080/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8080",
				},
			},
		},
	}
	resp, startErr := cli.ContainerCreate(ctx, conf, hostConf, nil, nil, "regression")
	if startErr != nil {
		t.Fatalf("Create docker container error: %v\n", startErr)
	}
	log.Println("Container create complete")

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		t.Fatalf("Start docker container error: %v\n", err)
	}
	log.Println("Container start complete")

	out, logsErr := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if logsErr != nil {
		t.Fatalf("Container logs error: %v\n", logsErr)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return cli, resp.ID
}

func StopAndRemoveContainer(cli *client.Client, id string, t *testing.T) {
	if err := cli.ContainerStop(context.Background(), id, container.StopOptions{}); err != nil {
		t.Fatalf("Cannot stop container: %v\n", err)
	}
	if err := cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{}); err != nil {
		t.Fatalf("Cannot remove container: %s\n", err)
	}
}

func TestIntegrationLinearRegression(t *testing.T) {
	cli, id := StartDockerContainer(t)
	defer StopAndRemoveContainer(cli, id, t)

	var testCases = []struct {
		req apiModels.LinearRegressionRequest

		AExpected   float64
		BExpected   float64
		ContentType string
		ErrExpected string
	}{
		{
			req: apiModels.LinearRegressionRequest{
				Pairs: []models.Point{models.Point{1, 2}, models.Point{2, 4.1}, models.Point{3, 7}, models.Point{4, 17}, models.Point{5, 22}},
			},
			AExpected:   5.29,
			BExpected:   -5.45,
			ContentType: "application/json",
			ErrExpected: "",
		},
		{
			req: apiModels.LinearRegressionRequest{
				Pairs: []models.Point{models.Point{45.1, 68.8}, models.Point{59, 61.2}, models.Point{57.2, 59.9}, models.Point{61.8, 56.7}, models.Point{58.8, 55}, models.Point{47.2, 54.3}, models.Point{55.2, 49.3}},
			},
			AExpected:   0,
			BExpected:   0,
			ContentType: "plain/text",
			ErrExpected: "Unsupported Media Type",
		},
		{
			req: apiModels.LinearRegressionRequest{
				Pairs: []models.Point{models.Point{45.1, 68.8}, models.Point{59, 61.2}, models.Point{57.2, 59.9}, models.Point{61.8, 56.7}, models.Point{58.8, 55}, models.Point{47.2, 54.3}, models.Point{55.2, 49.3}},
			},
			AExpected:   -0.35,
			BExpected:   76.88,
			ContentType: "application/json",
			ErrExpected: "",
		},
	}

	var httpClient = http.DefaultClient

	for i := 0; i < len(testCases); i++ {
		data, _ := json.Marshal(testCases[i].req)
		reader := bytes.NewReader(data)
		req, reqErr := http.NewRequest("GET", "http://localhost:8080/linear", reader)
		req.Header.Add("Content-Type", testCases[i].ContentType)
		if reqErr != nil {
			t.Fatalf("creating http request error: %v\n", reqErr)
		}
		resp, doErr := httpClient.Do(req)
		if doErr != nil {
			t.Fatalf("doing http request error: %v,\n", doErr)
		}
		respData, bodyReadErr := io.ReadAll(resp.Body)
		if bodyReadErr != nil {
			t.Fatalf("Body read error: %v\n", bodyReadErr)
		}
		var respModel apiModels.LinearRegressionResponce

		unmarshErr := json.Unmarshal(respData, &respModel)
		if unmarshErr != nil {
			t.Fatalf("Responce unmarshal error: %v\n", unmarshErr)
		}
		assert.Equal(t, testCases[i].ErrExpected, respModel.Err)
		assert.LessOrEqual(t, math.Abs(respModel.A-testCases[i].AExpected), 0.01)
		assert.LessOrEqual(t, math.Abs(respModel.B-testCases[i].BExpected), 0.01)
	}

}

func TestIntegrationNonlinear2Regression(t *testing.T) {
	cli, id := StartDockerContainer(t)

	defer StopAndRemoveContainer(cli, id, t)
}
