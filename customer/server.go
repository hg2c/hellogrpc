package customer

func NewServer(hostPort string, tracer opentracing.Tracer, metricsFactory metrics.Factory, logger log.Factory) *Server {
	return &Server{
		hostPort: hostPort,
		tracer:   tracer,
		logger:   logger,
		database: newDatabase(
			tracing.Init("mysql", metricsFactory.Namespace("mysql", nil), logger),
			logger.With(zap.String("component", "mysql")),
		),
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Nihao " + in.Name}, nil
}

func main() {
	logger0, _ := zap.NewDevelopment()
	logger := log.NewFactory(logger0.With(zap.String("service", "greeter")))

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Bg().Fatal("Unable to start server", zap.Error(err))
	}

	metricsFactory := xkit.Wrap("", expvar.NewFactory(10)) // 10 buckets for histograms
	logger.Bg().Info("Using expvar as metrics backend")

	tracer := tracing.Init("greeter", metricsFactory.Namespace("greeter", nil), logger)

	th := otgrpc.NewTraceHandler(tracer)
	s := grpc.NewServer(grpc.StatsHandler(th))
	// s := grpc.NewServer(
	// 	grpc.UnaryInterceptor(
	// 		otgrpc.OpenTracingServerInterceptor(tracer)))

	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logger.Bg().Fatal("failed to serve: %v", zap.Error(err))
	}
}
