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
- After download completely, move to the go-template, open and edit the `init.sh` file and modify the `REPO_NAME` and `REPO_HOST` values bases on your git repository.
- Then use this command to create new project.
```
    ./init.sh create project-name path
```
This command will create new project with name project-name. Be careful this name is using by GO code.
Script saving data in path/project_name

How to generate a service?
----------------------
```
    ./init.sh gen ServiceName project_path/
```
This command will generate a service inside the created project, includes simple CURD for a model