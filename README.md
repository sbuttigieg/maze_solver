# maze_solver

## DATABASE
IMPORTANT: It is assumed that Postgres is installed on the machine running this code.</br>
When the code is run, it checks that the database 'mazesolver' is created and that it has a table called 'levels'.</br>
If they do not already exist, they are automatically created.</br>

## Coding process
The process to write this app was as follows:
1. Sketch a plan    [30 mins]
2. Create the REST API [30 mins]
3. Create the level validation [45 mins]
4. Add database functionality [2 hours]
5. Add the minimum path functionality [5 hours]
6. Final Testing and debugging [1 hour]

## Configuration
The app configuration parameters are all located in constants/constants.go</br>
The database table structure is configured in database/tables.go

## REST API
The REST API has these endpoints:
- GET /levels ==> gets all levels
- GET /levels/[id] ==> gets a level by ID
- POST /level ==> adds a level ==> returns the ID of the level added
- DELETE /[id] ==> deletes a level by ID
- PATCH /[id] ==> updates a level by ID

Example web calls are stored in the postman collection (folder 'postman').
API endpoints are also tested in 'api_test.go'

## Validation
For both new levels and updated levels, before storing a level in the database the following validation is performed:
- that maps are rectangular by checking if all rows in the level have the same length as the first row.
- that maps are not larger than 100 in both X and Y directions. (Here it is assuming that the map is already validated as rectangular as only the length of row 0 is validated)
- that all tiles are valid (values 0 to 2).
- that a level has one starting point (a count of 0 or >1 triggers an error)
- that a level has one exit point, meaning a tile of value 0 in the top row (a count of 0 or >1 triggers an error)
- NOTE: the minimum path is calculated during validation but even if there is no solution, a minimum path of 0 is still accepted.

## Minimum Path
To find the minimum path, there are 2 steps:
1. Use a recursive method to find all the winning paths in the maze
2. Find the shortest path out of all winning paths

The recursive method works like this:
1. Set the start tile as the initial path
2. Make a copy of the initial path to use it as a starting point for the new path
3. Try to move up (unless the incoming direction is from front)
4. Check if the new location is the same as the previous location.
    - If it is then it means there was a wall.
        - Set path as dead and jump to step 5.
    - If it is not then it means that the tile was open.
        - Add the new location to the path
        - Check if the new location is the exit point.
            - If it is
                - Add the path to the list of winning paths.
                - Jump to step 5.
            - If it is not start recursion of method from step 2.
5. Try to move left (unless the incoming direction is from the left)
6. Check if the new location is the same as the previous location.
    - If it is then it means there was a wall.
        - Set path as dead and jump to step 7.
    - If it is not then it means that the tile was open.
        - Add the new location to the path
        - Check if the new location is the exit point.
            - If it is
                - Add the path to the list of winning paths.
                - Jump to step 7.
            - If it is not start recursion of method from step 2.
7. Try to move right (unless the incoming direction is from the right)
8. Check if the new location is the same as the previous location.
    - If it is then it means there was a wall.
        - Set path as dead and jump to step 9.
    - If it is not then it means that the tile was open.
        - Add the new location to the path
        - Check if the new location is the exit point.
            - If it is
                - Add the path to the list of winning paths.
                - Jump to step 9.
            - If it is not start recursion of method from step 2.
9. Try to move down (unless the incoming direction is from behind)
10. Check if the new location is the same as the previous location.
    - If it is then it means there was a wall.
        - Set path as dead and jump to step 11.
    - If it is not then it means that the tile was open.
        - Add the new location to the path
        - Check if the new location is the exit point.
            - If it is
                - Add the path to the list of winning paths.
                - Jump to step 11.
            - If it is not start recursion of method from step 2.
11. Return to the point that called the method, be it the initial call or within the recursion.

## Testing
- All major functions have a test code.
- While the database functions do not have direct test code, they are indirectly tested in api_test.go.
- These tests write data into the mazesolver database. This was done for simplicity. If this code was for a production app, I would have mocked a database for testing.
- Similarly fastestWinningPath() and findPath() do not have direct test code, and are also indirectly tested in api_test.go.
- The Postman collection also includes different level examples including a 100x100 maze and a maze with over a million possible paths and 770 different winning paths. These where used to test for speed of the minimum path calculation on the largest maze. With more time, I would have tried to optimise the code to make it faster.
