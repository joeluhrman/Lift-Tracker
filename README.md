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
6. Better "enums" for PPLTypes and MuscleGroups (custom enum implementation?)
    - Potential Issues
        1. Being able to convert any string into a PPLType or mscgrp / use ones other than consts defined in types
            a. fixed for pplType and mscgrp
        2. Creation of eTypes allowing any string to be decoded from JSON into a pplType or mscgrp
            a. only an issue if allow custom exercise type creation
7. Finish musclegroups

## Frontend
1. Make a select exercise component that can be put in a modal to show
a nice list of each exericse type and return the chosen exercise type
to the parent component.
2. Refactor workout form for adding logs

### Notes
1. Workout Form (1 form component, 1 configurable route component, 4 actual routes) configurable for:
    a. Add/Edit Workout Template
    b. Add/Edit Workout Log
    - Log has:
        - setgroups have weight
        - Notes on each exercise
        - Notes on log
        - Date on log
    - Also need to add buttons to remove and exercise and setgroup inputs from the form
2. Workout Templates and History are totally separate components / routes

## Long term
1. Custom exercise types
2. Email verification
3. Password recovery
4. Stat tracking, graphs, etc.