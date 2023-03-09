# ToDo

1. More complex password and username requirements.
2. Add email to users/email verification
3. ExerciseType stuff might need to be thought out more:
    * no custom exercise types for now
4. Logs
    * How to keep track of order of exercises/setgroups
        * Maybe keep workout id and delete/redo everything every time a workout is modified

5. SET UP FOREIGN KEY CONSTRAINTS 

6. CreateWorkoutTemplate:
    * Need to actually check that exercise type ID's for the exercises exist
    * Figure out how to cancel/delete everything already inserted if there is an error 

7. Modify data layer functions to update ID of the struct passed in
8. Graceful shutdown of database connection and server