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

export type CreateProjectRequest = {
  project?: OrgV1Org.Project
}

export type CreateProjectResponse = {
  project?: OrgV1Org.Project
}

export type GetProjectListRequest = {
  teamId?: string
}

export type GetProjectListResponse = {
  projects?: OrgV1Org.Project[]
}

export type GetProjectInfoRequest = {
  id?: string
}

export type GetProjectInfoResponse = {
  project?: OrgV1Org.Project
}

export type UpdateProjectInfoRequest = {
  project?: OrgV1Org.Project
}

export type UpdateProjectInfoResponse = {
  project?: OrgV1Org.Project
}

export class TeamService {
  static CreateOrg(req: CreateOrgRequest, initReq?: fm.InitReq): Promise<CreateOrgResponse> {
    return fm.fetchReq<CreateOrgRequest, CreateOrgResponse>(`/gapi/org/v1/create`, {...initReq, method: "POST"})
  }
  static GetOrgList(req: GetOrgListRequest, initReq?: fm.InitReq): Promise<GetOrgListResponse> {
    return fm.fetchReq<GetOrgListRequest, GetOrgListResponse>(`/gapi/org/v1/list?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
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
  static CreateProject(req: CreateProjectRequest, initReq?: fm.InitReq): Promise<CreateProjectResponse> {
    return fm.fetchReq<CreateProjectRequest, CreateProjectResponse>(`/gapi/team/v1/project/create`, {...initReq, method: "POST"})
  }
  static GetProjectList(req: GetProjectListRequest, initReq?: fm.InitReq): Promise<GetProjectListResponse> {
    return fm.fetchReq<GetProjectListRequest, GetProjectListResponse>(`/gapi/team/v1/project/list?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetProjectInfo(req: GetProjectInfoRequest, initReq?: fm.InitReq): Promise<GetProjectInfoResponse> {
    return fm.fetchReq<GetProjectInfoRequest, GetProjectInfoResponse>(`/gapi/team/v1/project/info?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static UpdateProjectInfo(req: UpdateProjectInfoRequest, initReq?: fm.InitReq): Promise<UpdateProjectInfoResponse> {
    return fm.fetchReq<UpdateProjectInfoRequest, UpdateProjectInfoResponse>(`/gapi/team/v1/project/info`, {...initReq, method: "PUT"})
  }
}