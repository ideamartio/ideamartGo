package ideago;

import "github.com/joomtriggers/ideago/sms"

//  wh

// config = sms.NewConfiguration().SetPassword().SetServer().SetApplication()
// sms.NewMessage("Message").SendMessage(config);

//proper code to handle a SMS call should be done

func SMS() (*sms.Sender) {
	sender := &sms.Sender{};

	return sender
}
