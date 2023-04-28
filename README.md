# Lift-Tracker
Webapp (and maybe mobile in the future) for tracking lifting progress.  
It uses a Postgresql database, Golang web-server, and React frontend scaffolded with Vite.  

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
1. Figure out how to abstract forms so all have same appearance
2. Add workout template form
    - select exercise modal

## Long term
1. Custom exercise types
2. Email verification
3. Password recovery
4. Stat tracking, graphs, etc.