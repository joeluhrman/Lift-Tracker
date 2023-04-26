# Lift-Tracker
Webapp (and maybe mobile in the future) for tracking lifting progress.  
It uses a Postgresql database, Golang web-server, and React frontend scaffolded with Vite.  

### Structure
#### Golang web-server (root of repository)
Packages:
* `main`: just contains `main.go` which just contains main function
* `server`: contains functionality for handling HTTP requests/JSON API
* `storage`: contains functionality for data CRUDs
* `types`: contains types used across `storage` and `server`

#### React frontend (`web` directory)
Important folders:
* root: contains `index.html` entry point
* `/src`: contains `main.jsx` main file called by `index.html`
* `/src/routes`: contains components w/ client-side routes (full pages)
* `/src/components`: contains components like forms, etc. used to makeup full pages
* `/src/handlers`: contains classes for handling HTTP requests in a uniform way 

# ToDo

## Backend
1. Setup config file for db urls, server port, etc
1. Editing workout logs
2. Editing workout templates
3. Deleting workout logs
4. Deleting workout templates
5. Templates/Logs
    * How to keep track of order of exercises/setgroups
        * Maybe keep workout id and delete/redo everything every time a workout is modified
6. Figure out what to do with exercise type images
7. Make exercise types able to have multiple musclegroups and ppl types

## Frontend
1. Figure out how to change items on navbar based on logged in or not
2. Figure out how to abstract forms so all have same appearance

## Long term
1. Custom exercise types
2. Email verification
3. Password recovery
4. Stat tracking, graphs, etc.