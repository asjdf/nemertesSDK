/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as TeamV1Team from "./team.pb"
export type CreateTeamRequest = {
  team?: TeamV1Team.Team
}

export type CreateTeamResponse = {
  team?: TeamV1Team.Team
}

export type GetTeamListRequest = {
  userId?: string
}

export type GetTeamListResponse = {
  teams?: TeamV1Team.Team[]
}

export type GetTeamInfoRequest = {
  id?: string
}

export type GetTeamInfoResponse = {
  team?: TeamV1Team.Team
}

export type UpdateTeamInfoRequest = {
  team?: TeamV1Team.Team
}

export type UpdateTeamInfoResponse = {
  team?: TeamV1Team.Team
}

export type CreateProjectRequest = {
  project?: TeamV1Team.Project
}

export type CreateProjectResponse = {
  project?: TeamV1Team.Project
}

export type GetProjectListRequest = {
  teamId?: string
}

export type GetProjectListResponse = {
  projects?: TeamV1Team.Project[]
}

export type GetProjectInfoRequest = {
  id?: string
}

export type GetProjectInfoResponse = {
  project?: TeamV1Team.Project
}

export type UpdateProjectInfoRequest = {
  project?: TeamV1Team.Project
}

export type UpdateProjectInfoResponse = {
  project?: TeamV1Team.Project
}

export class TeamService {
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