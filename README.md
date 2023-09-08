# Maxxiene

A personalized cli for day to day tasks.

**COMMANDS:**
- `maxx backup`: "Backup a file or directory"
- `maxx convert`: "A way to convert files"
- `maxx proc`: "Running process utilities"
- `maxx setup`: "Set up initial configuration"
- `maxx sites`: "A simple bookmark manager"
- `maxx task`: Task/todo manager


## backup cmd
- `maxx backup {FileOrDirectory}`: Backup a file or directory. Files get saved as a .mxbkup file
- `maxx backup --config or -c`: Backs up the app configuration and supporting files

## convert cmd
- `maxx convert md2html` "Converts a markdown file into an html file"

## proc cmd
- `maxx proc dump {process name}`: "Creates a memory dump"
- `maxx proc kill {process name}`: "Kills a running process"
- `maxx proc watch {process name}`: "Monitors a running process memory and cpu usage"

## setup cmd
- `maxx setup` Sets up the initial configuration. ex:(todo.md, cofig/.env, data, etc.).

## sites cmd
- `maxx sites add {name} {url} {category}` Adds a bookmark to your bookmarks.json file.
- `maxx sites list` List all bookmarks sorted by category.
- `maxx sites delete {name}` Deletes a bookmark by name from bookmarks.json 

## tasks cmd
- `maxx task add {task}` Adds a task to the todo list (./todo.md).
- `maxx task list` List all tasks on todo list.
- `maxx task delete {index}` deletes a task at a given index.
