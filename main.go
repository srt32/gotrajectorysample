package main

import(
  trajectory "github.com/srt32/gotrajectory"
  "fmt"
)

func main(){
  accountName := "simon"
  projectName := "test"

  stories, err := trajectory.GetStories(accountName, projectName)

  if err != nil {
    fmt.Println(err)
  } else {
    for _, story := range stories {
      fmt.Println(story)
    }
  }
}
