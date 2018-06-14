package voice_it

import (
	"github.com/voiceittech/voiceit-go/voiceit"
	"github.com/rs/zerolog/log"
)

func main() {

	// Make Sure to add this at the top of your project

	myVoiceIt := voiceit.New("c882ab7aa83a4e6599c81a9fd930a318")

	// Now myVoiceIt is an instance of the VoiceIt class and can be used to make various different API Calls, as documented below.

	response := myVoiceIt.CreateUser("nEdAy", "abcd1234")

	log.Debug().Msg(response)

	response = myVoiceIt.CreateEnrollmentByWavURL("nEdAy", "abcd1234", "http://face-recognition-1253284991.file.myqcloud.com/store_8495613369280acb2277fb3546a42d33.mp3 ", "en-US")

	//log.Debug().Msg(response)

	response = myVoiceIt.AuthenticationByWavURL("nEdAy", "abcd1234", "http://face-recognition-1253284991.file.myqcloud.com/store_07cf9886e258d6ef63e6fe84a6121839.mp3  ", "en-US")

	log.Debug().Msg(response)
}
