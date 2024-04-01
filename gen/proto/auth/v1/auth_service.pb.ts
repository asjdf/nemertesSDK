/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type RefreshTokenRequest = {
  refresh_token?: string
}

export type RefreshTokenResponse = {
  access_token?: string
  refresh_token?: string
}

export type OAuthRedirectRequest = {
  platform?: string
  from?: string
}

export type OAuthRedirectResponse = {
  url?: string
}

export type OAuthCallbackRequest = {
  code?: string
  state?: string
}

export type OAuthCallbackResponse = {
  access_token?: string
  refresh_token?: string
}

export class AuthService {
  static RefreshToken(req: RefreshTokenRequest, initReq?: fm.InitReq): Promise<RefreshTokenResponse> {
    return fm.fetchReq<RefreshTokenRequest, RefreshTokenResponse>(`/gapi/auth/v1/refreshToken`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static OAuthRedirect(req: OAuthRedirectRequest, initReq?: fm.InitReq): Promise<OAuthRedirectResponse> {
    return fm.fetchReq<OAuthRedirectRequest, OAuthRedirectResponse>(`/gapi/auth/v1/oauth/redirect?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static OAuthCallback(req: OAuthCallbackRequest, initReq?: fm.InitReq): Promise<OAuthCallbackResponse> {
    return fm.fetchReq<OAuthCallbackRequest, OAuthCallbackResponse>(`/gapi/auth/v1/oauth/callback`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}