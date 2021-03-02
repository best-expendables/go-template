Go service template
===================
# Folder structure
- `internal/handler/` - defines your app routes and their logic
- `internal/service/` - defines your business logic
- `internal/util/` - code and functionality to be shared by different parts of the project
- `internal/model/` - defines model data
- `internal/repository/` - implements business logic and handles storage
- `internal/config/` - this folder contains configuration files, such as application settings, constants, ..., etc.
- `internal/constant/` - this folder contains constants value.
- `internal/cmd` - defines function to excute app, such as http, migration, seeder, ..., etc.

How to create project?
----------------------
```
    ./init.sh create project-name path
```
This command will create new project with name project-name. Be careful this name is using by GO code.
Script saving data in path/project_name