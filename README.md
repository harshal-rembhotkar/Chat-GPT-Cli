# Chat-Gpt-Cli
1. **Install Required Libraries**:
   ```bash
   go get github.com/sashabaranov/go-openai
   go get github.com/spf13/cobra
   go get github.com/spf13/viper
   ```

2. **Create a `.env` File**:
   Add your OpenAI API key to a `.env` file in the same directory:
   ```env
   API_KEY=your-openai-api-key
   ```

3. **Run the Program**:
   ```bash
   go run .
   ```

4. **Interact with ChatGPT**:
   Type your questions into the console. Type `quit` to exit.
