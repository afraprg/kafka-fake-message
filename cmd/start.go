package cmd

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/spf13/cobra"
	"io/ioutil"
	"kafka-fake-message/kafka"
	"strings"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start produce",
	Long:  `example description`,
	Run: func(cmd *cobra.Command, args []string) {
		brokers, _ := cmd.Flags().GetString("brokers")
		brokersList := strings.Split(brokers, ",")
		topic, _ := cmd.Flags().GetString("topic")
		interval, _ := cmd.Flags().GetInt64("interval")
		count, _ := cmd.Flags().GetInt64("count")
		generate, _ := cmd.Flags().GetBool("generate")
		file, _ := cmd.Flags().GetString("file")
		//init kafka topic
		kafkaProducer := kafka.Kafka{}
		kafkaProducer.Topic = topic
		kafkaProducer.Brokers = brokersList
		//send message on interval
		var msgCount int64 = 0
		ticker := time.NewTicker(time.Duration(interval) * time.Second)

		var message string
		var messages []string
		if file != "" {
			//read file parse and to messages array
			fileContent, err := ioutil.ReadFile(file) // just pass the file name
			if err != nil {
				panic("Cannot read file")
			}
			messages = strings.Split(string(fileContent), "\n")
		}

		messageFileReadIndex := 0

		for {
			select {
			case <-ticker.C:
				if generate {
					message = faker.Sentence()
				} else if file != "" {
					message = messages[messageFileReadIndex]
					messageFileReadIndex++
				}
				kafkaProducer.SendMessage(message)
				fmt.Println(message)
				msgCount++
				if msgCount == count {
					return
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().Int64P("count", "c", 10, "Count message send")
	startCmd.Flags().Int64P("interval", "i", 10, "Interval in second")
	startCmd.Flags().StringP("brokers", "b", "", "Brokers list")
	startCmd.Flags().StringP("topic", "t", "", "Topic name")
	startCmd.Flags().String("file", "", "file path")
	startCmd.Flags().Bool("generate", false, "Generate random string as message")
}
