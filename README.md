# Jokenpo

This project is a simple game of rock, paper and scissors.

The main objective of this project is to practice integrating an external game
with a tournament system.

Inspiration comes from [Faceit Match Lifecycle](https://developers.faceit.com/start/games/best-practices) and
previous experience working at [GamersClub](https://gamersclub.gg).

From Faceit documentation:
```mermaid
sequenceDiagram
    participant API as Games API
    participant Game as Game Server

    API->>Game: Match Configuration
    API-->>Game: Match Cancel
    Game->>API: Match Ready
    Game-->>API: Match Aborted
    Game->>API: Match Started
    loop Match in progress
        Game->>API: Match Updated
    end
    Game->>API: Match Finished
    Game-->>API: Match Stats

```
