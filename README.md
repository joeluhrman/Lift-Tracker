# Lift-Tracker
Webapp (and maybe mobile in the future) for tracking lifting progress.  
Right now I'm just working on the database and JSON API, going to start frontend soon.  
I'm focussing primarily on bare-bones functionality such as signup, login, logout, logging workouts/exercise/setgroups, and creating workout templates. Probably need to optimize data layer functions before expanding functionality past the basics.

# ToDo

1. Editing workout logs
2. Editing workout templates
3. Deleting workout logs
4. Deleting workout templates
5. Get user profile data layer/endpoint
7. Templates/Logs
    * How to keep track of order of exercises/setgroups
        * Maybe keep workout id and delete/redo everything every time a workout is modified
8. Figure out what to do with exercise type images
9. Make exercise types able to have multiple musclegroups and ppl types
10. FRONTEND

## Long term
1. Custom exercise types
2. Email verification
3. Password recovery
4. Stat tracking, graphs, etc.