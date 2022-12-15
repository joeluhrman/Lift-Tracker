# ToDo

1. Change testing db to localhost for development (this is much harder than anticipated)
2. More complex password and username requirements.
3. Figure out how to go about workouts/exercises/setgroups vs templates created by users vs default templates etc.
    a. use same structs for everything but just omit certain fields
    b. separate tables
        1. setgroups, exercises, workouts -> logged user data 
        2. default_exercises, default_workouts -> default templates shared across all users
        3. custom_exercises, custom_workouts -> user-made templates 
    c. Logged exercises and workouts must match either a default or custom template at the time of logging 
4. Make it more clear in code what deals with logged vs default vs custom exercises/workouts/whatever