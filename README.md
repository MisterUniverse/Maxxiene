# Maxxiene

A personalized cli for day to day tasks.

**COMMANDS:**
- `maxx backup`: "Backup a file or directory"
- `maxx convert`: "A way to convert files"
- `maxx crypto`: "Encrypt/decrypt"
- `maxx firewall` "Edit windows firewall rules"
- `maxx notes`: "Stores notes, files, photos, and dumps in the database"
- `maxx proc`: "Running process utilities"
- `maxx shred`: "Shreds a file or directory"
- `maxx sites`: "A simple bookmark manager"
- `maxx task`: "Task/todo manager"
- `maxx uwu`: "..."
- `maxx version`: "Prints the version number"


## backup cmd
- `maxx backup {FileOrDirectory}`: Backup a file or directory. Files get saved as a .mxbkup file
- `maxx backup --config or -c`: Backs up the app configuration and supporting files

## convert cmd
- `maxx convert md2html` "Converts a markdown file into an html file"

## crypto cmd
- `maxx crypto encrypt`: "encrypts a line of text"
- `maxx crypto decrypt`: "decrypts a line of text that was encrypted by you"

## firewall cmd
- `maxx firewall add` "Add windows firewall rules"
- `maxx firewall list` "List windows firewall rules"
- `maxx firewall del` "Delete windows firewall rules"
- `maxx firewall enable` "Enable windows firewall for specific profile"
- `maxx firewall disable` "Disable windows firewall for specific profile"
- `maxx firewall export` "Export windows firewall rules as binary and text file"
- `maxx firewall import` "Imports windows firewall rules binary file"

## notes cmd
- `maxx notes save`

## proc cmd
- `maxx proc dump {process name}`: "Creates a memory dump"
- `maxx proc kill {process name}`: "Kills a running process"
- `maxx proc watch {process name}`: "Monitors a running process memory and cpu usage"
- `maxx proc inject {process name} {payload}`: "Injects a dll into a running process"

## setup cmd
- `maxx setup` Sets up the initial configuration. ex:(maxxdb.db, cofig/.env, data, etc.).

## shred cmd
- `maxx shred {fileOrDirectory}`: Delete a file or directory

## sites cmd
- `maxx sites add {name} {url} {category}` Adds a bookmark to your bookmarks.json file.
- `maxx sites list` List all bookmarks sorted by category.
- `maxx sites delete {name}` Deletes a bookmark by name from bookmarks.json 

## tasks cmd
- `maxx task add {task}` Adds a task to the todo list.
- `maxx task list` List all tasks on todo list.
- `maxx task delete {index}` deletes a task at a given index.

## uwu cmd
- `maxx uwu`: "For those down bad days..."
