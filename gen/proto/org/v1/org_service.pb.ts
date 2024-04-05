/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as OrgV1Org from "./org.pb"
export type CreateOrgRequest = {
  org?: OrgV1Org.Org
}

export type CreateOrgResponse = {
  org?: OrgV1Org.Org
}

export type GetOrgListRequest = {
  userId?: string
}

export type GetOrgListResponse = {
  orgs?: OrgV1Org.Org[]
}

export type GetOrgMembersRequest = {
  orgId?: string
}

export type GetOrgMembersResponse = {
  member?: OrgV1Org.OrgMember[]
}

export type CreateTeamRequest = {
  team?: OrgV1Org.Team
}

export type CreateTeamResponse = {
  team?: OrgV1Org.Team
}

export type GetTeamListRequest = {
  userId?: string
}

export type GetTeamListResponse = {
  teams?: OrgV1Org.Team[]
}

export type GetTeamInfoRequest = {
  id?: string
}

export type GetTeamInfoResponse = {
  team?: OrgV1Org.Team
}

export type UpdateTeamInfoRequest = {
  team?: OrgV1Org.Team
}

export type UpdateTeamInfoResponse = {
  team?: OrgV1Org.Team
}

export type GetTeamMembersRequest = {
  teamId?: string
}

export type GetTeamMembersResponse = {
  member?: OrgV1Org.TeamMember[]
}

export class TeamService {
  static CreateOrg(req: CreateOrgRequest, initReq?: fm.InitReq): Promise<CreateOrgResponse> {
    return fm.fetchReq<CreateOrgRequest, CreateOrgResponse>(`/gapi/org/v1/create`, {...initReq, method: "POST"})
  }
  static GetOrgList(req: GetOrgListRequest, initReq?: fm.InitReq): Promise<GetOrgListResponse> {
    return fm.fetchReq<GetOrgListRequest, GetOrgListResponse>(`/gapi/org/v1/list?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetOrgMembers(req: GetOrgMembersRequest, initReq?: fm.InitReq): Promise<GetOrgMembersResponse> {
    return fm.fetchReq<GetOrgMembersRequest, GetOrgMembersResponse>(`/gapi/org/v1/members?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static CreateTeam(req: CreateTeamRequest, initReq?: fm.InitReq): Promise<CreateTeamResponse> {
    return fm.fetchReq<CreateTeamRequest, CreateTeamResponse>(`/gapi/team/v1/create`, {...initReq, method: "POST"})
  }
  static GetTeamList(req: GetTeamListRequest, initReq?: fm.InitReq): Promise<GetTeamListResponse> {
    return fm.fetchReq<GetTeamListRequest, GetTeamListResponse>(`/gapi/team/v1/list?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetTeamInfo(req: GetTeamInfoRequest, initReq?: fm.InitReq): Promise<GetTeamInfoResponse> {
    return fm.fetchReq<GetTeamInfoRequest, GetTeamInfoResponse>(`/gapi/team/v1/info?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static UpdateTeamInfo(req: UpdateTeamInfoRequest, initReq?: fm.InitReq): Promise<UpdateTeamInfoResponse> {
    return fm.fetchReq<UpdateTeamInfoRequest, UpdateTeamInfoResponse>(`/gapi/team/v1/info`, {...initReq, method: "PUT"})
  }
  static GetTeamMembers(req: GetTeamMembersRequest, initReq?: fm.InitReq): Promise<GetTeamMembersResponse> {
    return fm.fetchReq<GetTeamMembersRequest, GetTeamMembersResponse>(`/gapi/team/v1/members?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}