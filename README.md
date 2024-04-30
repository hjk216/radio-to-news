### Summary
This project was originally going to be an automated radio channel (think police, fire, EMS, park ranger) to news story (with a nice UI) pipeline and website. Since speech-to-text is relatively expensive this wouldn't really be feasible to create and host unless with some other buisness model. 

This is partially finished and you can see the next TODOs below. The general process is to retrieve the original audio, chunk it up, send to speech-to-text API (in this case Whisper), combine the text, and then use OpenAI to analyze the text in order to retrieve stories.

### TODO
- Take date as user input
- Hit Broadcastify API to get audio archives for a given radio channel and date
- Make speech-to-text and text-to-stories async
- If hosting this, retrieve and store content from s3 or the like.
    - Also, set up cron job to execute once a day for the given radio channels.
- In analyzing the text, play around with the prompt.
    - Sometimes an LLM can only do one thing well at a time, potentially have another step to remove personally identifiable information.

### OpenAI

- Models: https://platform.openai.com/docs/models/overview

### Broadcastify

- https://www.broadcastify.com/archives/feed/11446

### Go

- Effective Go: https://go.dev/doc/effective_go#if
- godub: https://pkg.go.dev/github.com/Vernacular-ai/godub
- audio: https://pkg.go.dev/github.com/auroraapi/aurora-go/audio
- OpenAI: https://pkg.go.dev/github.com/sashabaranov/go-openai
- AWS S3: https://pkg.go.dev/github.com/aws/aws-sdk-go/service/s3
