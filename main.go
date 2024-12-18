package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
   
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
   )
   
   func GetResponse(client *openai.Client, ctx context.Context, question string) {
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
	 Model: "gpt-3.5-turbo", // Use "gpt-4" for more advanced responses
	 Messages: []openai.ChatCompletionMessage{
	  {Role: "user", Content: question},
	 },
	})
	if err != nil {
	 fmt.Println("Error:", err)
	 os.Exit(13)
	}
   
	// Print the response from ChatGPT
	fmt.Println(resp.Choices[0].Message.Content)
   }
   
   type NullWriter int
   
   func (NullWriter) Write([]byte) (int, error) { return 0, nil }
   
   func main() {
	log.SetOutput(new(NullWriter)) // Disable logging
   
	// Load API key from environment variables or .env file
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
	 fmt.Println("Error reading config file:", err)
	 os.Exit(1)
	}
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
	 fmt.Println("Missing API KEY in .env file")
	 os.Exit(1)
	}
   
	ctx := context.Background()
	client := openai.NewClient(apiKey)
   
	// Set up the CLI application
	rootCmd := &cobra.Command{
	 Use:   "chatgpt",
	 Short: "Chat with ChatGPT in console.",
	 Run: func(cmd *cobra.Command, args []string) {
	  scanner := bufio.NewScanner(os.Stdin)
	  quit := false
   
	  for !quit {
	   fmt.Print("Say something ('quit' to end): ")
   
	   if !scanner.Scan() {
		break
	   }
   
	   question := scanner.Text()
	   switch question {
	   case "quit":
		quit = true
   
	   default:
		GetResponse(client, ctx, question)
	   }
	  }
	 },
	}
   
	// Execute the CLI command
	if err := rootCmd.Execute(); err != nil {
	 fmt.Println("Error:", err)
	 os.Exit(1)
	}
   }