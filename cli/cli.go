package cli

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
	appdata "translate/appData"
	"translate/config"
	"translate/query"
	"translate/query/sdo"
	"translate/utils"
)

func initLogin() (user *string, pass *string, loginCmd *flag.FlagSet) {
	cmd := flag.NewFlagSet("login", flag.ExitOnError)

	loginUser := cmd.String("u", "", "User name")
	loginPass := cmd.String("p", "", "Password")

	return loginUser, loginPass, cmd
}

func initMe() (meCmd *flag.FlagSet) {
	return flag.NewFlagSet("me", flag.ExitOnError)
}

func initCreate() (w *bool, url *string, language *string, title *string, createCmd *flag.FlagSet) {
	cmd := flag.NewFlagSet("create", flag.ExitOnError)

	wait := cmd.Bool("w", false, "wait for job to finish")
	textUrl := cmd.String("url", "", "text file url to be translated")
	targetLanguage := cmd.String("lang", "", "desired translation")
	jobTitle := cmd.String("title", "", "Job title")

	return wait, textUrl, targetLanguage, jobTitle, cmd
}

func initStatus() (*flag.FlagSet, *string, *bool) {
	cmd := flag.NewFlagSet("status", flag.ExitOnError)
	jId := cmd.String("id", "", "status id")
	statusWait := cmd.Bool("w", false, "wait for job to finish")

	return cmd, jId, statusWait
}

func initSearch() (*flag.FlagSet, *string) {
	cmd := flag.NewFlagSet("search", flag.ExitOnError)
	value := cmd.String("v", "", "Search value")

	return cmd, value
}

func checkStatus(id string, c chan string) {
	res := query.JobStatus(id);
	c <- res.Data.Job.Status
}

func WatchStatus(id string) {
	c := make(chan string)	
	go checkStatus(id, c)
	for status := range c {
		if (status == "complete") {
			fmt.Println("Job Status: ", status)
			os.Exit(0)
		}
		fmt.Printf("Job Status: %s\r", status)
		time.Sleep(3*time.Second)
		go checkStatus(id, c)
	}
}

func Init() {
	user, pass, loginCmd := initLogin()
	meCmd := initMe()
	wait, textUrl, lang, jobTitle, createCmd := initCreate()
	statusCmd, jobId, statusWait := initStatus()
	searchCmd, searchValue := initSearch()

	if len(os.Args) < 2 {
		fmt.Println("Command is required")
		fmt.Println("login for authorized actions")
		fmt.Println("Me to see current user")
		fmt.Println("create -url -lang to create text translation")
		fmt.Println("status to get all job statuses or -id to get specific job status")
		fmt.Println("search to find TDO's based on title (SDO field)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "login":
		loginCmd.Parse(os.Args[2:])
	case "me":
		meCmd.Parse(os.Args[2:])
	case "create":
		createCmd.Parse(os.Args[2:])
	case "status":
		statusCmd.Parse(os.Args[2:])
	case "search":
		searchCmd.Parse(os.Args[2:])
		
	}

	if meCmd.Parsed() {
		res := query.Me()
		utils.HandleGraphQLError(res.Errors)
		fmt.Println("Logged in as: ", res.Data.Me.Name)
	}

	if loginCmd.Parsed() {
		res := query.Login(*user, *pass)
		utils.HandleGraphQLError(res.Errors)
		err := utils.SaveAuthInfo(res.Data.UserLogin.Token, res.Data.UserLogin.Organization.Name, res.Data.UserLogin.Organization.ID)
		utils.HandleErr(err)
		regRes := query.GetRegistry()
		// need to save registry ID
		appdata.GetAppData()
		if (len(regRes.Data.DataRegistries.Records) == 0) {
			// create registry, schema draft, and publish schema draft
			createdReg := query.CreateRegistry();
			utils.HandleGraphQLError(res.Errors)
			appdata.SetRegistryId(createdReg.Data.CreateDataRegistry.ID);
			err := appdata.Commit()
			utils.HandleErr(err)

			// create schema and publish
			schemaDraft := query.CreateSchemaDraft(appdata.GetAppData().RegistryId)
			utils.HandleGraphQLError(schemaDraft.Errors)
			appdata.SetSchemaId(schemaDraft.Data.UpsertSchemaDraft.ID)
			err = appdata.Commit()
			utils.HandleErr(err)

			schemaPub := query.PublishSchemaDraft(appdata.GetAppData().SchemaId)
			utils.HandleGraphQLError(schemaPub.Errors)
			appdata.SetSchemaStatus(schemaPub.Data.UpdateSchemaState.Status)
			err = appdata.Commit()
			utils.HandleErr(err)

			// create a root folder
			rootFolder := query.GetRootFolder()
			utils.HandleGraphQLError(rootFolder.Errors)
			rootFolderId := rootFolder.Data.RootFolders[0].ID
			createFolderRes := query.CreateCLIFolder(config.CLI_FOLDER_NAME, config.CLI_FOLDER_DESCRIPTION, rootFolderId)
			appdata.SetCLIFolderId(createFolderRes.Data.CreateFolder.ID)
			err = appdata.Commit()
			utils.HandleErr(err)
		}
		
	}

	if createCmd.Parsed() {
		if (*textUrl != "" || *lang != "") {
			aData := appdata.GetAppData()
			authData := utils.GetAuthInfo()
			res := query.CreateTranslateJob("", *textUrl, *lang)
			utils.HandleGraphQLError(res.Errors)
			// Add job to appData
			appdata.AddJob(appdata.Job{
				Id: res.Data.CreateJob.ID,
				TargetId: res.Data.CreateJob.TargetID,
				Status: "pending",
				JobTitle: *jobTitle,
			})
			appdata.Commit()

			// Update target tdo with template content
			fmt.Printf("Updating TDO..\r")
			updateTdoRes  := query.UpdateTdoWithContent(sdo.UpdateTdoWithContentInput{
				TdoId: res.Data.CreateJob.TargetID,
				SchemaId: aData.SchemaId,
				Email: authData.Email,
				JobTitle: *jobTitle,
				TranslatedTo: *lang,
				Url: *textUrl,
			})
			utils.HandleGraphQLError(updateTdoRes.Errors)

			// Move TDO to cli folder
			fmt.Printf("Moving TDO..\r")
			moveTdoRes := query.MoveTDOToFolder(res.Data.CreateJob.TargetID, aData.FolderId)
			utils.HandleGraphQLError(moveTdoRes.Errors)
			fmt.Printf("moved tdo: %+v", moveTdoRes)

			if (*wait) {
				WatchStatus(res.Data.CreateJob.ID)
			}
		}
	}

	if statusCmd.Parsed() {
		w := new(tabwriter.Writer)
	
		w.Init(os.Stdout, 8, 8, 8, '\t', 0)
		
		defer w.Flush()
		
		fmt.Fprintf(w, "\n %s\t%s\t%s\t", "ID ", "STATUS", "TASKS DONE")
		fmt.Fprintf(w, "\n %s\t%s\t%s\t", "--", "------", "----------")

		if (*jobId != "") {
			if (*statusWait) {
				WatchStatus(*jobId)
			} else {
				res := query.JobStatus(*jobId);
				utils.HandleGraphQLError(res.Errors)
				
				tasksCompleted := utils.GetCompletedTasks(res.Data.Job.Tasks.Records)
				fmt.Fprintf(w, "\n %s\t%s\t%s\t", *jobId, res.Data.Job.Status, tasksCompleted)
			}
		} else {
			for _, j := range appdata.GetJobs() {
				res := query.JobStatus(j.Id);
				utils.HandleGraphQLError(res.Errors)
				tasksCompleted := utils.GetCompletedTasks(res.Data.Job.Tasks.Records)

				fmt.Fprintf(w, "\n %s\t%s\t%s\t", j.Id, res.Data.Job.Status, tasksCompleted)
			}
		}	
	}

	if searchCmd.Parsed() {
		if (*searchValue != "") {
			data := query.SearchByTitle(*searchValue)
			fmt.Println("====== RESULTS ======")
			for _, res := range data.Data.SearchMedia.Jsondata.Results {
				fmt.Println("TDO ID: ", res.Recording.RecordingID)		
			}
		}
	}
}