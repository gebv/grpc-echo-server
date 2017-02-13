package echo

import (
	"io"
	"log"
	"os"

	context "golang.org/x/net/context"

	"fmt"

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
	s.logger.Printf("\t+\t[Echo] str=%s, raw_length=%d\n", dto.Str, len(dto.Raw))
	return dto, nil
}

func (s *EchoServer) Track(stream Server_TrackServer) error {
	s.logger.Println("\t+\t[Track] start")
	for {
		track, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.logger.Printf("[Track] err %v\n", err)
			break
		}
		s.logger.Printf("\t+\t[Track] %v\n", track)
	}
	s.logger.Println("\t+\t[Track] end")
	return nil
}

func (s *EchoServer) ChatStream(req *StreamChatRequest, stream Server_ChatStreamServer) error {
	arr := make([]string, 10)
	s.logger.Println("\t+\t[ChatStream] start...\n")
	for index, str := range arr {
		if err := stream.Send(&ChatMessage{
			Text: fmt.Sprintf("msg: %s#%d", str, index),
		}); err != nil {
			return err
		}
	}

	s.logger.Printf("\t+\t[ChatStream] sent messages %d\n", len(arr))
	return nil
}
