# MaX Pondus

Ever dreamed to have maximum performance with maximum weights?

Worry no more, MaX Pondus is here!

## Structure

This monorepo is split into backend and frontend projects.

Both projects and their usage instructions can be found under the corresponding folders.

## Features

This section contains an abstract list of the service features.

**NOTE**: Project is currently under biiiiiig brainstorming so all features are currently TODO.

### User Management

- Authentication uses a simple HTTP Basic authentication.
- User roles are split into a normal and administrator roles.
- Administrators may add new users and modify or delete existing users.
- Administrators may also act as normal users.

### Exercises

- Mechanics can be classified into isolated and compound exercises.
- Exercises contain at least name and description.
- Exercises may contain additional links e.g. links to videos or additional information.
- Exercises may contain a list of the target and synergist muscles.
- Muscle and muscle group catalog is provided by the service.
- Administrators may add new exercises and modify or delete existing exercises.

### Workouts

- Users may add, modify and delete own workouts.
- Users may add additional notes per workout.
- Workouts can be individual or they can be part of an existing workout program.
- Workouts may contain 1..N sets.
- Workout set may contain 1..N exercises (e.g. supersets).
- Workout set exercises may contain differing weights (e.g. drop sets).
- A clock is shown at all times during the workout.
- A simple buttons to start and stop the workout and the workout set timer are shown in the UI.
- Workout and workout set durations are saved.

### Programs

- Users may add, modify and delete own programs.
- Programs may contain additional notes.
- Program workouts may contain additional notes.
- Program workouts contain rest times between sets.
- Program workouts can be individually customized (e.g. light and intense weeks etc.).
- User may see their own program progress and history.

## Entities

This section contains information about the application entities and their relationships.

```mermaid
erDiagram
  Role         ||--o{ User     : ""
  User         ||--o{ Program  : ""
  User         ||--o{ Workout  : ""
  Program      |o--o{ Workout  : ""
  Workout      ||--o{ Set      : ""
  Set          ||--o{ Subset   : ""
  Muscle-Group ||--o{ Muscle   : ""
  Muscle       }o--o{ Exercise : "" 
  Mechanics    ||--o{ Exercise : ""
  Exercise     ||--o{ Link     : ""
  Exercise     ||--o{ Subset   : ""
  User         ||--o{ Workout  : ""
```

| Entity       | Description                                                                            |
| ------------ | -------------------------------------------------------------------------------------- |
| Exercise     | An activity which requires physical effort.                                            |
| Link         | A link to a external source providing additional info (e.g. videos & additional info). |
| Mechanics    | A category how an exercise handles motion and force.                                   |
| Muscle       | A part of the body that handles the physical effort.                                   |
| Muscle Group | A logical group of muscles in a body.                                                  |
| Program      | A set of scheduled physical activities for a time period.                              |
| Role         | A group of application use permissions.                                                |
| Set          | A logical group of exercise repetition sets.                                           |
| Subset       | A set of repetitions of an exercise.                                                   |
| User         | An identifiable user of the application.                                               |
| Workout      | A single session of physical activity.                                                 |

## Schema

This section contains information about the application database schema.

TODO
