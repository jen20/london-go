package greeting

type GreeterRPCServer struct {
	// This is the real implementation
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(args *GreeterInputArgs, resp *string) error {
	*resp = s.Impl.Greet(args.Name)
	return nil
}
