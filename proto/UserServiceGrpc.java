package proto;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.3.0)",
    comments = "Source: user.proto")
public final class UserServiceGrpc {

  private UserServiceGrpc() {}

  public static final String SERVICE_NAME = "proto.UserService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<proto.UserOuterClass.CreateRequest,
      proto.UserOuterClass.CreateResponse> METHOD_CREATE =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.UNARY,
          generateFullMethodName(
              "proto.UserService", "Create"),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.CreateRequest.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.CreateResponse.getDefaultInstance()));
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<proto.UserOuterClass.GetRequest,
      proto.UserOuterClass.GetResponse> METHOD_GET =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.UNARY,
          generateFullMethodName(
              "proto.UserService", "Get"),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.GetRequest.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.GetResponse.getDefaultInstance()));
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<proto.UserOuterClass.GetAllRequest,
      proto.UserOuterClass.GetAllResponse> METHOD_GET_ALL =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.UNARY,
          generateFullMethodName(
              "proto.UserService", "GetAll"),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.GetAllRequest.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.GetAllResponse.getDefaultInstance()));
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<proto.UserOuterClass.UpdateRequest,
      proto.UserOuterClass.UpdateResponse> METHOD_UPDATE =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.UNARY,
          generateFullMethodName(
              "proto.UserService", "Update"),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.UpdateRequest.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.UpdateResponse.getDefaultInstance()));
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<proto.UserOuterClass.DeleteRequest,
      proto.UserOuterClass.DeleteResponse> METHOD_DELETE =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.UNARY,
          generateFullMethodName(
              "proto.UserService", "Delete"),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.DeleteRequest.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(proto.UserOuterClass.DeleteResponse.getDefaultInstance()));

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UserServiceStub newStub(io.grpc.Channel channel) {
    return new UserServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UserServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new UserServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary and streaming output calls on the service
   */
  public static UserServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new UserServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class UserServiceImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * CREATE
     * </pre>
     */
    public void create(proto.UserOuterClass.CreateRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.CreateResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_CREATE, responseObserver);
    }

    /**
     * <pre>
     * READ
     * </pre>
     */
    public void get(proto.UserOuterClass.GetRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.GetResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET, responseObserver);
    }

    /**
     */
    public void getAll(proto.UserOuterClass.GetAllRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.GetAllResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_GET_ALL, responseObserver);
    }

    /**
     * <pre>
     * UPDATE
     * </pre>
     */
    public void update(proto.UserOuterClass.UpdateRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.UpdateResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_UPDATE, responseObserver);
    }

    /**
     * <pre>
     * DELETE
     * </pre>
     */
    public void delete(proto.UserOuterClass.DeleteRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.DeleteResponse> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_DELETE, responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            METHOD_CREATE,
            asyncUnaryCall(
              new MethodHandlers<
                proto.UserOuterClass.CreateRequest,
                proto.UserOuterClass.CreateResponse>(
                  this, METHODID_CREATE)))
          .addMethod(
            METHOD_GET,
            asyncUnaryCall(
              new MethodHandlers<
                proto.UserOuterClass.GetRequest,
                proto.UserOuterClass.GetResponse>(
                  this, METHODID_GET)))
          .addMethod(
            METHOD_GET_ALL,
            asyncUnaryCall(
              new MethodHandlers<
                proto.UserOuterClass.GetAllRequest,
                proto.UserOuterClass.GetAllResponse>(
                  this, METHODID_GET_ALL)))
          .addMethod(
            METHOD_UPDATE,
            asyncUnaryCall(
              new MethodHandlers<
                proto.UserOuterClass.UpdateRequest,
                proto.UserOuterClass.UpdateResponse>(
                  this, METHODID_UPDATE)))
          .addMethod(
            METHOD_DELETE,
            asyncUnaryCall(
              new MethodHandlers<
                proto.UserOuterClass.DeleteRequest,
                proto.UserOuterClass.DeleteResponse>(
                  this, METHODID_DELETE)))
          .build();
    }
  }

  /**
   */
  public static final class UserServiceStub extends io.grpc.stub.AbstractStub<UserServiceStub> {
    private UserServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private UserServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new UserServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * CREATE
     * </pre>
     */
    public void create(proto.UserOuterClass.CreateRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.CreateResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_CREATE, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * READ
     * </pre>
     */
    public void get(proto.UserOuterClass.GetRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.GetResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAll(proto.UserOuterClass.GetAllRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.GetAllResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_GET_ALL, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * UPDATE
     * </pre>
     */
    public void update(proto.UserOuterClass.UpdateRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.UpdateResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_UPDATE, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * DELETE
     * </pre>
     */
    public void delete(proto.UserOuterClass.DeleteRequest request,
        io.grpc.stub.StreamObserver<proto.UserOuterClass.DeleteResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_DELETE, getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class UserServiceBlockingStub extends io.grpc.stub.AbstractStub<UserServiceBlockingStub> {
    private UserServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private UserServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new UserServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * CREATE
     * </pre>
     */
    public proto.UserOuterClass.CreateResponse create(proto.UserOuterClass.CreateRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_CREATE, getCallOptions(), request);
    }

    /**
     * <pre>
     * READ
     * </pre>
     */
    public proto.UserOuterClass.GetResponse get(proto.UserOuterClass.GetRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET, getCallOptions(), request);
    }

    /**
     */
    public proto.UserOuterClass.GetAllResponse getAll(proto.UserOuterClass.GetAllRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_GET_ALL, getCallOptions(), request);
    }

    /**
     * <pre>
     * UPDATE
     * </pre>
     */
    public proto.UserOuterClass.UpdateResponse update(proto.UserOuterClass.UpdateRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_UPDATE, getCallOptions(), request);
    }

    /**
     * <pre>
     * DELETE
     * </pre>
     */
    public proto.UserOuterClass.DeleteResponse delete(proto.UserOuterClass.DeleteRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_DELETE, getCallOptions(), request);
    }
  }

  /**
   */
  public static final class UserServiceFutureStub extends io.grpc.stub.AbstractStub<UserServiceFutureStub> {
    private UserServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private UserServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new UserServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * CREATE
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.UserOuterClass.CreateResponse> create(
        proto.UserOuterClass.CreateRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_CREATE, getCallOptions()), request);
    }

    /**
     * <pre>
     * READ
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.UserOuterClass.GetResponse> get(
        proto.UserOuterClass.GetRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.UserOuterClass.GetAllResponse> getAll(
        proto.UserOuterClass.GetAllRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_GET_ALL, getCallOptions()), request);
    }

    /**
     * <pre>
     * UPDATE
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.UserOuterClass.UpdateResponse> update(
        proto.UserOuterClass.UpdateRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_UPDATE, getCallOptions()), request);
    }

    /**
     * <pre>
     * DELETE
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.UserOuterClass.DeleteResponse> delete(
        proto.UserOuterClass.DeleteRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_DELETE, getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE = 0;
  private static final int METHODID_GET = 1;
  private static final int METHODID_GET_ALL = 2;
  private static final int METHODID_UPDATE = 3;
  private static final int METHODID_DELETE = 4;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final UserServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(UserServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CREATE:
          serviceImpl.create((proto.UserOuterClass.CreateRequest) request,
              (io.grpc.stub.StreamObserver<proto.UserOuterClass.CreateResponse>) responseObserver);
          break;
        case METHODID_GET:
          serviceImpl.get((proto.UserOuterClass.GetRequest) request,
              (io.grpc.stub.StreamObserver<proto.UserOuterClass.GetResponse>) responseObserver);
          break;
        case METHODID_GET_ALL:
          serviceImpl.getAll((proto.UserOuterClass.GetAllRequest) request,
              (io.grpc.stub.StreamObserver<proto.UserOuterClass.GetAllResponse>) responseObserver);
          break;
        case METHODID_UPDATE:
          serviceImpl.update((proto.UserOuterClass.UpdateRequest) request,
              (io.grpc.stub.StreamObserver<proto.UserOuterClass.UpdateResponse>) responseObserver);
          break;
        case METHODID_DELETE:
          serviceImpl.delete((proto.UserOuterClass.DeleteRequest) request,
              (io.grpc.stub.StreamObserver<proto.UserOuterClass.DeleteResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static final class UserServiceDescriptorSupplier implements io.grpc.protobuf.ProtoFileDescriptorSupplier {
    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return proto.UserOuterClass.getDescriptor();
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (UserServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UserServiceDescriptorSupplier())
              .addMethod(METHOD_CREATE)
              .addMethod(METHOD_GET)
              .addMethod(METHOD_GET_ALL)
              .addMethod(METHOD_UPDATE)
              .addMethod(METHOD_DELETE)
              .build();
        }
      }
    }
    return result;
  }
}
