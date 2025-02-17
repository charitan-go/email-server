package rotation

import (
	"log"
	"time"

	"github.com/charitan-go/key-server/internal/key/service"
	"github.com/charitan-go/key-server/pkg/env"
)

type RotationServer struct {
	initKeyDuration   time.Duration
	rotateKeyDuration time.Duration

	keySvc service.KeyService
}

func NewRotationServer(keySvc service.KeyService) *RotationServer {
	return &RotationServer{keySvc: keySvc}
}

func (s *RotationServer) readConfig() {
	initKeyDurationStr, _ := env.ReadEnv("INIT_KEY_DURATION")
	rotateKeyDurationStr, _ := env.ReadEnv("ROTATE_KEY_DURATION")

	if initKeyDuration, err := time.ParseDuration(initKeyDurationStr); err != nil {
		log.Fatalln("Error in parsing init key duration")
	} else {
		s.initKeyDuration = initKeyDuration
	}

	if rotateKeyDuration, err := time.ParseDuration(rotateKeyDurationStr); err != nil {
		log.Fatalln("Error in parsing rotate key duration")
	} else {
		s.rotateKeyDuration = rotateKeyDuration
	}
}

func (s *RotationServer) Run() {
	s.readConfig()

	time.Sleep(s.initKeyDuration)
	for true {
		s.keySvc.GenerateKeyPairs()
		time.Sleep(s.rotateKeyDuration)
	}
}
