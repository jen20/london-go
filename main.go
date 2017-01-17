package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/jen20/london-go/greeting"
)

// Run ./greeter to use the default English plugin.
// Run ./greeter french to use the French plugin.
// Run ./greeter anythingelse to see an impressive crash.
func main() {
	language := "english"
	if len(os.Args) > 1 {
		language = os.Args[1]
	}

	// Note this is BAD, but demonstrates the concept.
	pluginCmd := fmt.Sprintf("./greeter-%s", language)

	log.SetOutput(ioutil.Discard)

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "LONDON_GO_DEMO",
			MagicCookieValue: "Hi",
		},
		Plugins: map[string]plugin.Plugin{
			"greeter": new(greeting.GreeterPlugin),
		},
		Cmd: exec.Command(pluginCmd),
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(greeting.Greeter)
	fmt.Println(greeter.Greet())
}
