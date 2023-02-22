package slackclient_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/rneatherway/gh-slack/internal/slackclient"
	"github.com/tjgurwara99/mixtape"
	"github.com/tjgurwara99/mixtape/player"
)

func compareFuncWithBodyAndAuth(r *http.Request, recording *mixtape.Request) bool {
	if !mixtape.CompareFuncWithBody(r, recording) {
		return false
	}
	if r.Header.Get("Authorization") != recording.Header.Get("Authorization") {
		return false
	}
	return true
}

func TestWriteMessage(t *testing.T) {
	cassette, err := mixtape.Load("testdata/write_message")
	if err != nil {
		t.Fatal(err)
	}
	cassette.Comparer = compareFuncWithBodyAndAuth
	cassetteTranport := player.New(cassette, player.Replay, http.DefaultTransport)
	httpClient := &http.Client{
		Transport: cassetteTranport,
	}
	auth := &slackclient.SlackAuth{
		Token: "abc123",
	}
	client, err := slackclient.New("github", log.Default(), httpClient, auth)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.SendMessage("C04QJSB8G4D", "Sending a message from test to mock http interactions")
	if err != nil {
		t.Fatal(err)
	}
}
