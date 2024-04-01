/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as UserV1User from "./user.pb"
export type GetUserInfoRequest = {
  id?: string
}

export type GetUserInfoResponse = {
  info?: UserV1User.User
}

export type EditUserInfoRequest = {
  info?: UserV1User.User
}

export type EditUserInfoResponse = {
  info?: UserV1User.User
}

export class UserService {
  static GetUserInfo(req: GetUserInfoRequest, initReq?: fm.InitReq): Promise<GetUserInfoResponse> {
    return fm.fetchReq<GetUserInfoRequest, GetUserInfoResponse>(`/gapi/user/v1/info?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static EditUserInfo(req: EditUserInfoRequest, initReq?: fm.InitReq): Promise<EditUserInfoResponse> {
    return fm.fetchReq<EditUserInfoRequest, EditUserInfoResponse>(`/gapi/user/v1/edit`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}