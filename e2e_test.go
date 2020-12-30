package pixela_test

import (
	"context"
	"os"
	"testing"

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
	result, err := cli.CreateGraph(
		ctx,
		gid, "acc test",
		"count",
		pixela.Int,
		pixela.Shibafu,
		pixela.TimeZone("Asia/Tokyo"),
	)
	if err != nil {
		t.Fatalf("CreateGraph() failed: %v", err)
	}
	if !result.IsSuccess {
		t.Fatalf("CreateGraph() failed %#v", result)
	}

	result, err = cli.DeleteGraph(ctx, gid)
	if err != nil {
		t.Fatalf("DeleteGraph() failed: %v", err)
	}
	if !result.IsSuccess {
		t.Fatalf("DeleteGraph() failed %#v", result)
	}
}
