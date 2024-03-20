package mail

import (
	"testing"

	"github.com/dpurbosakti/booknest-grpc/internal/config"

	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	config, err := config.LoadConfig("../..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from moko</p>
	`

	to := []string{"dwiatmokop@gmail.com", "sirensym@gmail.com"}
	attachFiles := []string{"../../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
