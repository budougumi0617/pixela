package pixela_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/budougumi0617/pixela"
)

func TestAcceptanceClient(t *testing.T) {
	name := os.Getenv("PIXELA_USER")
	if len(name) == 0 {
		t.Fatal("not find PIXELA_USER")
	}
	token := os.Getenv("PIXELA_TOKEN")

	if len(token) == 0 {
		t.Fatal("not find PIXELA_TOKEN")
	}
	cli := pixela.New(name, token)
	ctx := context.Background()
	gid := pixela.GraphID("acceptance-test")
	gd := &pixela.GraphDefinition{
		ID:             gid,
		Name:           "acc test",
		Unit:           "count",
		Type:           pixela.Int,
		Color:          pixela.Shibafu,
		TimeZone:       "Asia/Tokyo",
		SelfSufficient: pixela.SelfSufficientNone,
	}
	gds := []*pixela.GraphDefinition{gd}
	// cleanup
	_, _ = cli.DeleteGraph(ctx, gid)

	// create graph
	result, err := cli.CreateGraph(
		ctx,
		gd.ID, gd.Name,
		gd.Unit,
		gd.Type,
		gd.Color,
		pixela.TimeZone(gd.TimeZone),
		pixela.SelfSufficient(gd.SelfSufficient),
	)
	if err != nil {
		t.Fatalf("CreateGraph() failed: %v", err)
	}
	if !result.IsSuccess {
		t.Fatalf("CreateGraph() failed %#v", result)
	}

	// get graph
	gotGd, err := cli.GetGraph(ctx, gid)
	if err != nil {
		t.Fatalf("GetGraph() failed: %v", err)
	}
	if diff := cmp.Diff(gotGd, gd); diff != "" {
		t.Fatalf("GetGraph() diff: (-got +want)\n%s", diff)
	}

	// update graph
	gd.Color = pixela.Momiji
	gd.Unit = "time"
	result, err = cli.UpdateGraph(ctx, gd)
	if err != nil {
		t.Fatalf("UpdateGraph() failed: %v", err)
	}
	if !result.IsSuccess {
		t.Fatalf("UpdateGraph() failed %#v", result)
	}
	gotGd, err = cli.GetGraph(ctx, gid)
	if err != nil {
		t.Fatalf("GetGraph() failed: %v", err)
	}
	if diff := cmp.Diff(gotGd, gd); diff != "" {
		t.Fatalf("GetGraph() diff: (-got +want)\n%s", diff)
	}

	// get all graphs
	gotGds, err := cli.GetGraphs(ctx)
	if err != nil {
		t.Fatalf("GetGraphs() failed: %v", err)
	}
	if diff := cmp.Diff(gotGds, gds); diff != "" {
		t.Fatalf("GetGraphs() diff: (-got +want)\n%s", diff)
	}

	// delete graph
	result, err = cli.DeleteGraph(ctx, gid)
	if err != nil {
		t.Fatalf("DeleteGraph() failed: %v", err)
	}
	if !result.IsSuccess {
		t.Fatalf("DeleteGraph() failed %#v", result)
	}
}
