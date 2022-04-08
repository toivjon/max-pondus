# MaX Pondus - Backend

The backend for the MaX Pondus weight training journal.

## Build
The build scripts can be found from the scripts folder.

Windows users should note that the build script uses goimports which needs the diff to be present in the $PATH.

### Windows
```console
./scripts/build.bat
```

### Unix/Linux
```console
./scripts/build.sh
```

## REST API
This application introduces a REST API to access the desired features.

Furthermore the REST API is split into Admin API and Personal API.

### Admin API
The Admin API is used to perform administrative operations like adding new users and exercises.

This API requires the user to have the ADMIN role.

Here is an overview of the available end points.

| URL                          | Method | Description                            |
| ---------------------------- | ------ | -------------------------------------- |
| /admin/api/v1/users          | GET    | Get the list of users.                 |
| /admin/api/v1/users          | POST   | Add a user.                            |
| /admin/api/v1/users/{id}     | GET    | Get info about an individual user.     |
| /admin/api/v1/users/{id}     | PUT    | Update a user.                         |
| /admin/api/v1/users/{id}     | DELETE | Delete a user.                         |
| /admin/api/v1/exercises      | GET    | Get the list of exercises.             |
| /admin/api/v1/exercises      | POST   | Add an exercise.                       |
| /admin/api/v1/exercises/{id} | GET    | Get info about an individual exercise. |
| /admin/api/v1/exercises/{id} | PUT    | Update an exercise.                    |
| /admin/api/v1/exercises/{id} | DELETE | Delete an exercise.                    |

### Personal API
The Personal API is used to perform personal workout operations like writing down workouts etc.

Here is an overview of the available end points.

TODO A list of available end points.