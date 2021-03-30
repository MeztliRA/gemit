package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

const (
	timerStartedMessage      = "timer started!"
	timerDoneMessage         = "timer done!"
	notEnoughArgumentMessage = "timer: to many or not enough argument"
)

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Start a new timer",
	Long: `Start a new gemit timer of a specified duration.
	the duration used for the timer is from the argument passed.`,
	Run: func(cmd *cobra.Command, args []string) {
		gemitTimer(args)
	},
}

func init() {
	rootCmd.AddCommand(timerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func gemitTimer(args []string) {
	if len(args) != 1 {
		log.Fatal(notEnoughArgumentMessage)
	}
	inputtedDuration := args[0]
	duration, err := time.ParseDuration(inputtedDuration)
	if err != nil {
		log.Fatal(err)
	}
	timer := time.NewTimer(duration)
	fmt.Println(timerStartedMessage)
	<-timer.C
	fmt.Println(timerDoneMessage)
}
