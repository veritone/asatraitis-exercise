package cli

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
	"translate/query"
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

func initCreate() (w *bool, url *string, language *string, createCmd *flag.FlagSet) {
	cmd := flag.NewFlagSet("create", flag.ExitOnError)

	wait := cmd.Bool("w", false, "wait for job to finish")
	textUrl := cmd.String("url", "", "text file url to be translated")
	targetLanguage := cmd.String("lang", "", "desired translation")	

	return wait, textUrl, targetLanguage, cmd
}

func initStatus() (*flag.FlagSet, *string, *bool) {
	cmd := flag.NewFlagSet("status", flag.ExitOnError)
	jId := cmd.String("id", "", "status id")
	statusWait := cmd.Bool("w", false, "wait for job to finish")

	return cmd, jId, statusWait
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
	wait, textUrl, lang, createCmd := initCreate()
	statusCmd, jobId, statusWait := initStatus()

	if len(os.Args) < 2 {
		fmt.Println("Command is required")
		fmt.Println("login for authorized actions")
		fmt.Println("Me to see current user")
		fmt.Println("create -url -lang to create text translation")
		fmt.Println("status to get all job statuses or -id to get specific job status")
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
	}

	if meCmd.Parsed() {
		res := query.Me()
		utils.HandleGraphQLError(res.Errors)
		fmt.Println("Logged in as: ", res.Data.Me.Name)
	}

	if loginCmd.Parsed() {
		res := query.Login(*user, *pass)
		utils.HandleGraphQLError(res.Errors)
		fmt.Print(res.Data.UserLogin.Token)
		err := utils.SaveToken(res.Data.UserLogin.Token)
		utils.HandleErr(err)
	}

	if createCmd.Parsed() {
		fmt.Println("Need to create go routine to poll status?", *wait)
		if (*textUrl != "" || *lang != "") {
			res := query.CreateTranslateJob("", *textUrl, *lang)
			utils.HandleGraphQLError(res.Errors)
			utils.SaveJob(res.Data.CreateJob.ID)
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
			for _, jId := range utils.GetJobs() {
				res := query.JobStatus(jId);
				utils.HandleGraphQLError(res.Errors)
				tasksCompleted := utils.GetCompletedTasks(res.Data.Job.Tasks.Records)

				fmt.Fprintf(w, "\n %s\t%s\t%s\t", jId, res.Data.Job.Status, tasksCompleted)
			}
		}	
	}
}