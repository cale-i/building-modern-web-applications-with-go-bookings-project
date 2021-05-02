#!bin/bash

go get github.com/gobuffalo/pop/...

echo `development:
  dialect: postgres
  database: bookings
  user: postgres
  password: postgres
  host: 127.0.0.1
  pool: 5

test:
  url: {{envOr "TEST_DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_test"}}

production:
  url: {{envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/myapp_production"}}` >> ./database.yml


# Create DB
soda generate fizz CreateUsersTable
soda generate fizz CreateReservationsTable
soda generate fizz CreateRoomsTable
soda generate fizz CreateRestrictionsTable
soda generate fizz CreateRoomRestrictionsTable
soda migrate

# Setting up FK
soda generate fizz CreateFKFORReservationsTable
soda generate fizz CreateFKFORRoomRestrictionsTable
soda migrate

# UNQUE KEY
soda generate fizz CreateUniqueIndexForUsersTable
soda generate fizz CreateIndicesOnRoomRestrictions

# Add Indices
soda generate fizz CreateIndicesOnRoomRestrictions
soda generate fizz AddFKAndIndicesToReservationsTable

# recreate tables
soda reset
