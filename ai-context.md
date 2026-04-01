
# Climbing-Go Project Overview

This document provides a summary of the `climbing-go` project, intended to be used as a context for AI-assisted development.

## Project Structure

The project is a 2D game written in Go, using the Ebiten game engine and the Donburi ECS library.

- **`main.go`**: The entry point of the application. It initializes the game and starts the game loop.
- **`go.mod`**: Defines the module and its dependencies.
  - `github.com/hajimehoshi/ebiten/v2`: The game engine.
  - `github.com/yohamta/donburi`: The ECS library.
- **`core/`**: Contains the core game logic.
  - **`ebiten/game.go`**: The main `Game` struct, which manages the game loop and delegates to the level manager.
  - **`level/`**: Manages game levels.
    - `manager.go`: Creates and manages levels, including setting up the initial level, systems, and entities.
    - `level.go`: Defines the `Level` struct, which contains the ECS world and the game map.
  - **`ecs/`**: The Entity-Component-System implementation.
    - **`archetypes/`**: Defines entity archetypes, such as the player.
      - `player.go`: A function to create a player entity with all its components.
    - **`components/`**: Defines the data components for entities.
      - `physics/`: Components related to physics, such as `Position`, `Velocity`, and `Transform`.
      - `visuals/`: Components related to visuals, such as `Sprite`.
      - `input/`: Components related to player input.
      - `stats/`: Components related to entity stats, such as `MoveSpeed`.
    - **`systems/`**: Defines the systems that operate on entities.
      - `physics.go`: Updates the position of entities based on their velocity.
      - `render.go`: Draws entities with sprites to the screen.
      - `movement/movement.go`: Updates the velocity of entities based on player input.
- **`assets/`**: Contains game assets, such as images.

## Game Loop

1.  **`main.go`**: The `main` function creates a new `Game` object and calls `ebiten.RunGame(game)`.
2.  **`core/ebiten/game.go`**: The `Game`'s `Update` and `Draw` methods are called on each frame.
3.  The `Game`'s `Update` and `Draw` methods delegate to the `levelManager.CurrentLevel.ECS()`.
4.  The ECS then calls the `Update` and `Draw` methods of all registered systems.

## Key Systems

- **`movement`**: Reads player input and updates the `Velocity` component of the player entity.
- **`physics`**: Updates the `Position` component of entities based on their `Velocity`.
- **`render`**: Draws entities with a `Sprite` and `Position` component to the screen.
