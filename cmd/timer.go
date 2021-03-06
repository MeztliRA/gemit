package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

const (
	timerStartedMessage      = "timer started!\n"
	timerDoneMessage         = "timer done!"
	notEnoughArgumentMessage = "timer: not enough argument"
	tooManyArgumentMessage   = "timer: to many argument"
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
}

func gemitTimer(args []string) {
	log.SetPrefix("gemit: ")
	log.SetFlags(0)

	if len(args) < 1 {
		log.Fatal(notEnoughArgumentMessage)
	} else if len(args) > 1 {
		log.Fatal(tooManyArgumentMessage)
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
	// no icon is just a placeholder so that beeep doesnt find a icon to use
	notificationErr := beeep.Notify("Gemit", timerDoneMessage, "no icon")
	if notificationErr != nil {
		log.Fatal(notificationErr)
	}
}
