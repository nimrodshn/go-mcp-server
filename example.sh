#!/bin/bash

# Example script to run the agent
# Make sure to set your API key first!

echo "🚀 Starting MCP Agent Example"
echo ""

# Check if API key is set
if [ -z "$API_KEY" ]; then
    echo "❌ Please set your API_KEY environment variable first:"
    echo ""
    echo "For Claude:"
    echo "export API_KEY=your_claude_api_key"
    echo "export AI_PROVIDER=claude"
    echo ""
    echo "For OpenAI:"
    echo "export API_KEY=your_openai_api_key"
    echo "export AI_PROVIDER=openai"
    echo ""
    exit 1
fi

# Start the MCP server in the background
echo "📡 Starting MCP server..."
go run main.go &
MCP_PID=$!

# Wait a moment for the server to start
sleep 2

# Check if the server is running
if ! curl -s http://localhost:8080 > /dev/null; then
    echo "❌ MCP server failed to start"
    kill $MCP_PID 2>/dev/null
    exit 1
fi

echo "✅ MCP server is running at http://localhost:8080"

# Set default values if not provided
export MCP_SERVER_URL=${MCP_SERVER_URL:-"http://localhost:8080"}
export AI_PROVIDER=${AI_PROVIDER:-"claude"}

if [ "$AI_PROVIDER" = "claude" ]; then
    export MODEL=${MODEL:-"claude-3-sonnet-20240229"}
else
    export MODEL=${MODEL:-"gpt-3.5-turbo"}
fi

echo "🤖 Starting agent with:"
echo "   MCP Server: $MCP_SERVER_URL"
echo "   AI Provider: $AI_PROVIDER"
echo "   Model: $MODEL"
echo ""

# Start the agent
go run agent.go

# Clean up: kill the MCP server when the agent exits
echo "🧹 Cleaning up..."
kill $MCP_PID 2>/dev/null
echo "✅ Done!"