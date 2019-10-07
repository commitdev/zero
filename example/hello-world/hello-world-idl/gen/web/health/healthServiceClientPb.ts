/**
 * @fileoverview gRPC-Web generated client stub for health
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as google_api_annotations_pb from '../../google/api/annotations_pb';
import * as protoc$gen$swagger_options_annotations_pb from '../../protoc-gen-swagger/options/annotations_pb';

import {
  HealthCheckRequest,
  HealthCheckResponse} from './health_pb';

export class HealthClient {
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
    HealthCheckResponse,
    (request: HealthCheckRequest) => {
      return request.serializeBinary();
    },
    HealthCheckResponse.deserializeBinary
  );

  check(
    request: HealthCheckRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: HealthCheckResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/health.Health/Check',
      request,
      metadata || {},
      this.methodInfoCheck,
      callback);
  }

  methodInfoWatch = new grpcWeb.AbstractClientBase.MethodInfo(
    HealthCheckResponse,
    (request: HealthCheckRequest) => {
      return request.serializeBinary();
    },
    HealthCheckResponse.deserializeBinary
  );

  watch(
    request: HealthCheckRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/health.Health/Watch',
      request,
      metadata || {},
      this.methodInfoWatch);
  }

}

