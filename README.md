# Lift-Tracker
Webapp (and maybe mobile in the future) for tracking lifting progress.  
It uses a Postgresql database, Golang web-server, and React frontend scaffolded with Vite.  

# ToDo

## Backend
1. Setup config file for db urls, server port, etc
2. Editing logs/templates
3. Deleting logs/templates
4. Templates/Logs
    * How to keep track of order of exercises/setgroups
        * Maybe keep workout id and delete/redo everything every time a workout is modified
5. Remove exType images for now
6. Better "enums" for PPLTypes and MuscleGroups
    - Problems 
        - Don't want invalid values for either making through creation endpoint
        and into the db
        - Don't want devs to be able to create them on the fly / accidentally 
        convert an invalid string into one or something
7. Finish musclegroups

## Frontend
1. Add workout template form
    - select exercise modal

## Long term
1. Custom exercise types
2. Email verification
3. Password recovery
4. Stat tracking, graphs, etc.