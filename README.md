# Multiplayer Game Server

This repository contains a basic implementation of a multiplayer game server written in Go, comprising a server component and a player module. The server facilitates communication between multiple players in a game session, allowing them to exchange messages and update their positions within the game world.

## Server Component

The `main` package includes the server implementation. Here's a brief overview of its components:

### Server Struct

- **Message**: Defines the structure of messages exchanged between clients.
- **Server**: Manages the server instance, handling connections, message passing, and session management.

### Server Functions

- **NewServer**: Initializes a new server instance.
- **Start**: Starts the server, accepting incoming connections and handling messages.
- **acceptLoop**: Handles accepting new client connections.
- **readLoop**: Handles reading messages from clients and updating player positions accordingly.

## Player Module

The `playermodule` package includes functionalities related to managing player sessions and positions within the game.

### PlayerModule Functions

- **GenAllSessions**: Initializes the session manager.
- **AddToSession**: Adds a player to a game session.
- **UpdatePlayerPos**: Updates a player's position within a session.
- **GetSession**: Retrieves session information based on session ID.
- **GetSessionIdFromBuf**: Extracts session ID from a byte buffer.
- **GetPlayerPosFromBuf**: Extracts player position from a byte buffer.

## Usage

To use this server implementation, follow these steps:

1. Clone the repository.
2. Build the server using Go.
3. Run the server binary, specifying the desired listening address.
