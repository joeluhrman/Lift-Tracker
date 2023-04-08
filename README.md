# Lift-Tracker
Webapp (and maybe mobile in the future) for tracking lifting progress.  
Right now I'm just working on the database and JSON API, going to start frontend soon.  
I'm focussing primarily on bare-bones functionality such as signup, login, logout, logging workouts/exercise/setgroups, and creating workout templates. Probably need to optimize data layer functions before expanding functionality past the basics.

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