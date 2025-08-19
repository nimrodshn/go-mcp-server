# MCP Greeter Server

A simple Model Context Protocol (MCP) server written in Go that provides a greeting tool.

## Features

- **Greet Tool**: Say hello to someone by name

## Prerequisites

- Go 1.24.5 or later

## Installation

```bash
go mod download
```

## Usage

1. Start the server:
```bash
./mcp-server
```

2. The server will be available at `http://localhost:8080`

3. Connect to Claude Code:
```bash
claude mcp add --transport sse greeter http://localhost:8080
```

## API

### Tools

#### `greet`
- **Description**: Say hi to a person
- **Parameters**:
  - `name` (string): The name of the person to greet
- **Returns**: A greeting message

## Development

To modify the server, edit `main.go` and restart the application.