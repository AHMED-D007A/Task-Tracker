# Task Tracker
- Task Tracker is a cli application that take arguments from the user and parse them to excute a certain operation depending on user input for these arguments.
- This project is from the project based road map, you can find it [here](https://roadmap.sh/backend/projects).
- You can read more about the project form [here](https://roadmap.sh/projects/task-tracker).
- I wanted to practice the basics of the language so, I work on this project using only file and string operation using the standard library packages (didn't use json marshal std).

### Usage
```
# Adding a new task
./bin/task_tracker add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
./bin/task_tracker update 1 "Buy groceries and cook dinner"
./bin/task_tracker delete 1

# Marking a task as in progress or done
./bin/task_tracker mark-in-progress 1
./bin/task_tracker mark-done 1

# Listing all tasks
./bin/task_tracker list

# Listing tasks by status
./bin/task_tracker list done
./bin/task_tracker list todo
./bin/task_tracker list in-progress
```