/**
 * @fileoverview gRPC-Web generated client stub for helloworld
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as proto_health_health_pb from '../../proto/health/health_pb';
import * as google_api_annotations_pb from '../../google/api/annotations_pb';
import * as protoc$gen$swagger_options_annotations_pb from '../../protoc-gen-swagger/options/annotations_pb';

export class HelloworldClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoCheck = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_health_health_pb.HealthCheckResponse,
    (request: proto_health_health_pb.HealthCheckRequest) => {
      return request.serializeBinary();
    },
    proto_health_health_pb.HealthCheckResponse.deserializeBinary
  );

  check(
    request: proto_health_health_pb.HealthCheckRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_health_health_pb.HealthCheckResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/helloworld.Helloworld/Check',
      request,
      metadata || {},
      this.methodInfoCheck,
      callback);
  }

}

