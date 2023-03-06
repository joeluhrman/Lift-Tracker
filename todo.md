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
5. Check exercises and workouts exist in either custom tables or default tables before creating (also get data for exercise musclegroups from there)
    - potential issue with users created exercises with same name as defaults, I guess dont allow that
6. add email to users

## the tables question
The issue is what to do with actual logged workouts vs exercise templates and workout templates.  
I guess it would be easiest to think of it in terms of the entities the user would be aware of:  
1. Logged data
    * this would be in the form of workouts (which include exercises and setgroups)

2. Exercise templates (i called it a "type" in the code instead of template)
    * a name, image, musclegroup / ppl type, etc. that would be selected by the user to add to their current (logged) workout (or saved one)

3. Workout templates
    * when starting a workout the user could (optionally) choose an existing workout template, which would have most of what a logged workout does.  
    * they should also be able to go off-template and edit things on the fly and have that work, so they wouldn't be locked into the exact template once they've chosen it.

Maybe it would be simplest to start with exercise templates for now and work up from there?

Exercise templates should be strict, as in a user cannot change anything about it on the fly (so we can do accurate stats on exercises).  
Workout templates/logged workouts are much looser because they don't actually provide any analytical value so who cares.