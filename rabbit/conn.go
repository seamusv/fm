package rabbit

import "github.com/streadway/amqp"

type Conn struct {
	Channel *amqp.Channel
}

func GetConn(rabbitURL string) (Conn, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return Conn{}, err
	}

	ch, err := conn.Channel()
	return Conn{
		Channel: ch,
	}, err
}
