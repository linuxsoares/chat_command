package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	configprompt "github.com/linuxsoares/chat_command/configPrompt"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a phrase: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	res, err := generateCommitMessageWithOpenAI(phrase)
	if err != nil {
		fmt.Println("Error to generate command line:", err)
		return
	}
	fmt.Println("Generated command line:", res)
	err = executeCommand(res)
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}

}

func generateCommitMessageWithOpenAI(changeText string) (string, error) {
	client := openai.NewClient(CHAT_COMMAND_OPEN_AI_TOKEN)
	ctx := context.Background()

	prompt := configprompt.UserPrompt + changeText
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: configprompt.SystemPrompt,
			},
		},
		MaxTokens: 500,
	}
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

func executeCommand(command string) error {
	var userResponse string
	fmt.Println("Can ? (yes/no)")
	fmt.Scanln(&userResponse)

	if strings.ToLower(userResponse) == "yes" {
		// Commit the changes
		cmd := exec.Command("bash", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Command execution` aborted by user.")
	}
	return nil
}
