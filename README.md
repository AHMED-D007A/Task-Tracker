# Task Tracker
- Task Tracker is a cli application that take arguments from the user and parse them to excute a certain operation depending on user input for these arguments.
- This project is from the project based road map, you can find it [here](https://roadmap.sh/backend/projects)

- I wanted to practice the basics of the language so, I work on this project using only file and string operation using the standard library packages.

### Usage
```
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```