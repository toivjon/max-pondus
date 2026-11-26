# Domain Model

## Overview

This document defines the core domain concepts, relationships and business rules of Max Pondus.

It serves as a reference for frontend and backend development, architectural decisions and AI-assisted code generation.

The domain model is technology-independent. It describes business concepts rather than database tables, API payloads or implementation-specific structures.

## Entity Relationship Diagram

```mermaid
erDiagram
  USER         ||--o{ WORKOUT  : "owns"
  WORKOUT      ||--o{ SET      : "contains"
  SET          ||--o{ SUBSET   : "contains"
  MUSCLE-GROUP ||--o{ MUSCLE   : "contains"
  MUSCLE       }o--o{ EXERCISE : "associated with" 
  EXERCISE     ||--o{ SUBSET   : "referenced by"

  USER {
    string id
  }

  WORKOUT {
    uuid     id
    datetime startedAt
    datetime endedAt
  }

  SET {
    uuid id
  }

  SUBSET {
    uuid       id
    datetime   startedAt
    datetime   endedAt
    uint       repetitions
    uint       weight
    WeightUnit weightUnit
  }
```

## Entities

| Entity       | Description                                                                     |
| ------------ | ------------------------------------------------------------------------------- |
| Exercise     | A defined physical movement performed during training (e.g. Dumbbell Row).      |
| Muscle       | A specific muscle involved in performing exercises.                             |
| Muscle Group | A category of related muscles (e.g. Chest).                                     |
| Set          | A logical training unit that can represent a normal set, superset, or drop set. |
| Subset       | A recorded set of repetitions of a single exercise.                             |
| User         | An identifiable user of the application.                                        |
| Workout      | An individual training session.                                                 |

## Reference Data

Exercises, Muscles and Muscle Groups are predefined domain concepts identified by stable, language-independent codes.

Their allowed values and relationships are defined in [shared/exercises.json](../shared/exercises.json).

The catalog serves as the single source of truth for these definitions.

Display names and descriptions are localized by the frontend.

## Business Rules

### Workout

- A Workout must have a start time, but its end time is optional.
- When provided, endedAt must not precede startedAt.
- Workouts can be saved without being completed.
- Recorded Workouts can be modified afterwards.
- Sets within a Workout are ordered by the earliest available Subset start time, when available.

### Set

- Subsets within a Set are ordered by their start times, when available.
- Subsets within the same Set may reference the same or different Exercises.
- Subsets within the same Set may have different weights and repetition counts.

### Subset

- Weight and repetition count cannot be negative.
- Exercises may be performed without additional weight.
- Subset timestamps are optional. When both are provided, endedAt must not precede startedAt.

### Identifier Stability

- Published identifiers must not be changed or reused, as they may be referenced by recorded Workouts.
