package echo

import (
	"log"
	"os"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

func Setup(s *grpc.Server) {
	RegisterServerServer(s, &EchoServer{
		logger: log.New(os.Stderr, "[echo] ", log.Ldate|log.Lmicroseconds),
	})
}

var _ ServerServer = (*EchoServer)(nil)

type EchoServer struct {
	logger *log.Logger
}

func (s *EchoServer) Echo(_ context.Context, dto *EchoDTO) (*EchoDTO, error) {
	s.logger.Println(">> str=%s, raw_length=%d", dto.Str, len(dto.Raw))
	return dto, nil
}
