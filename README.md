# aiWARE Translate
CLI tool written in go to create txt file translation jobs. 

# Commands
## login
Running translate w/ login command will allow you to authenticate with dev environment and save token locally in `.token`
### Flow
```mermaid
flowchart TD
  A[Login] --> B{CLI has data registry already?};
  B -- Yes --> C[Set auth info into auth.json]
  B -- No --> W[Set auth info into auth.json]
  W --> D[Create Data Registry and save id in app_data.json]
  C --> Q[Done]  
  D --> E[Create Schema and save id in app_data.json]
  E --> F[Publish Schema and save id in app_data.json]
  F --> G[Create Folder and save id in app_data.json]
  G --> Q[Done]
```
### flags
- `-u` Username
- `-p` Password
### Example
`./translate login -u <username> -p <password>` or `go run . login -u <username> -p <password>`
## me
Me command will print the authenticate user if `.token` is valid
### Example
`./translate me` or `go run . me`
## create
Create command will create a txt file translate job and append job id to `jobs` file for reference.
### flags
- `-url` URL to a txt file to be translated
- `-lang` Language the txt file content to be translated to
- `-w` Wait for the job to complete before exiting
### Example
`./translate create -url https://example-files.online-convert.com/document/txt/example.txt -lang de -w`
or
`go run . create -url https://example-files.online-convert.com/document/txt/example.txt -lang de -w`
## status
Status command will display all job statuses or single job status if id is provided
### flags
- `-id` Job Id
- `-w` Wait for the job to be complete before exiting
### Example
- `./translate status` or `go run . status` Displays all jobs with their statuses and tasks in a table format
- `./translate status -id <jobId>` or `go run . status -d <jobId>` Displays a job and task status
- `./translate status -id <jobId> -w` or `go run . status -d <jobId> -w` Wait for a job to be done before exiting
## search
Search for TDO by SDO title field
### flags
- `-v` Value used for fidning TDO by title field
### Example
- `./translate search -v testing` or `go run . search -v testing` Displays all jobs with their statuses and tasks in a table format
