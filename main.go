package main

import (
	"context"
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type HiParams struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

func SayHi(ctx context.Context, session *mcp.ServerSession, req *mcp.CallToolParamsFor[HiParams]) (*mcp.CallToolResultFor[any], error) {
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "Hi " + req.Arguments.Name + ", Its nice to meet you"}},
	}, nil
}

func CodeReviewPrompt(ctx context.Context, session *mcp.ServerSession, req *mcp.GetPromptParams) (*mcp.GetPromptResult, error) {
	code, ok := req.Arguments["code"]
	if !ok {
		code = "// No code provided"
	}

	return &mcp.GetPromptResult{
		Description: "Code review request",
		Messages: []*mcp.PromptMessage{
			{
				Role: "user",
				Content: &mcp.TextContent{
					Text: "Please review this code and provide feedback on:\n" +
						"1. Code quality and style\n" +
						"2. Potential bugs or issues\n" +
						"3. Performance improvements\n" +
						"4. Best practices\n\n" +
						"Code to review:\n```\n" + code + "\n```",
				},
			},
		},
	}, nil
}

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "say hi"}, SayHi)

	server.AddPrompt(&mcp.Prompt{
		Name:        "code_review",
		Description: "Request code review with structured feedback",
		Arguments: []*mcp.PromptArgument{
			{
				Name:        "code",
				Description: "The code to review",
				Required:    true,
			},
		},
	}, CodeReviewPrompt)

	log.Printf("MCP servers serving at %s", "http://localhost:8080")
	handler := mcp.NewSSEHandler(func(request *http.Request) *mcp.Server {
		return server
	})
	log.Fatal(http.ListenAndServe(":8080", handler))
}
