// @generated by protoc-gen-connect-es v0.8.6 with parameter "target=ts"
// @generated from file proto/user.proto (package proto.user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserRequest, CreateUserResponse, GetUserByEmailRequest, GetUserByEmailResponse, GetUserRequest, GetUserResponse, LoginRequest, LoginResponse, WhoAmIReponse, WhoAmIRequest } from "./user_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service proto.user.v1.UserService
 */
export const UserService = {
  typeName: "proto.user.v1.UserService",
  methods: {
    /**
     * @generated from rpc proto.user.v1.UserService.GetUser
     */
    getUser: {
      name: "GetUser",
      I: GetUserRequest,
      O: GetUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.user.v1.UserService.GetUserByEmail
     */
    getUserByEmail: {
      name: "GetUserByEmail",
      I: GetUserByEmailRequest,
      O: GetUserByEmailResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.user.v1.UserService.CreateUser
     */
    createUser: {
      name: "CreateUser",
      I: CreateUserRequest,
      O: CreateUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.user.v1.UserService.Login
     */
    login: {
      name: "Login",
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.user.v1.UserService.WhoAmI
     */
    whoAmI: {
      name: "WhoAmI",
      I: WhoAmIRequest,
      O: WhoAmIReponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

