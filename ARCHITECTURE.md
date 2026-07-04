# Go Mine Game - Architecture Overview

This document provides a high-level overview of the go-mine-game architecture.

## System Architecture

```mermaid
graph TB
    subgraph "User Interface"
        CLI["Command Line Interface"]
    end
    
    subgraph "Game Core"
        GameEngine["Game Engine"]
        GameState["Game State Manager"]
        Board["Board Logic"]
    end
    
    subgraph "Game Logic"
        MineLogic["Mine Logic"]
        RevealLogic["Reveal Logic"]
        WinLoseLogic["Win/Lose Detection"]
        FlagLogic["Flag Management"]
    end
    
    subgraph "Utilities"
        Renderer["Board Renderer"]
        InputHandler["Input Handler"]
        Validator["Move Validator"]
    end
    
    CLI -->|sends commands| InputHandler
    InputHandler -->|parses input| GameEngine
    GameEngine -->|manages flow| GameState
    GameState -->|queries/updates| Board
    
    Board -->|uses| MineLogic
    Board -->|uses| RevealLogic
    Board -->|uses| WinLoseLogic
    Board -->|uses| FlagLogic
    
    GameEngine -->|validates moves| Validator
    GameEngine -->|displays board| Renderer
    Renderer -->|reads| GameState
    
    style CLI fill:#e1f5ff
    style GameEngine fill:#fff3e0
    style GameState fill:#fff3e0
    style Board fill:#fff3e0
    style MineLogic fill:#f3e5f5
    style RevealLogic fill:#f3e5f5
    style WinLoseLogic fill:#f3e5f5
    style FlagLogic fill:#f3e5f5
    style Renderer fill:#e8f5e9
    style InputHandler fill:#e8f5e9
    style Validator fill:#e8f5e9
```

## Component Descriptions

### User Interface Layer
- **Command Line Interface**: Entry point for user interactions with the game

### Game Core Layer
- **Game Engine**: Main orchestrator that coordinates game flow and operations
- **Game State Manager**: Maintains current game state (board configuration, flags, revealed cells)
- **Board Logic**: Core board data structure and operations

### Game Logic Layer
- **Mine Logic**: Handles mine placement and mine-related operations
- **Reveal Logic**: Manages cell revelation and cascade effects
- **Win/Lose Detection**: Determines game end conditions
- **Flag Management**: Handles flagging and unflagging cells

### Utilities Layer
- **Board Renderer**: Displays the game board to the player
- **Input Handler**: Processes and parses player commands
- **Move Validator**: Validates user moves before execution

## Data Flow

1. Player enters command via CLI
2. Input Handler parses the command
3. Game Engine receives the validated command
4. Game Engine uses Move Validator to check if the move is legal
5. Game Engine updates Game State and Board Logic
6. Game Logic components execute the move (reveal, flag, etc.)
7. Win/Lose Detection checks for end conditions
8. Board Renderer displays updated board
9. Game continues or ends based on game state
