/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as ProjectV1Project from "./project.pb"
export type CreateProjectRequest = {
  project?: ProjectV1Project.Project
  teamId?: string
}

export type CreateProjectResponse = {
  project?: ProjectV1Project.Project
}

export type GetProjectListRequest = {
  teamId?: string
}

export type GetProjectListResponse = {
  projects?: ProjectV1Project.Project[]
}

export type GetProjectInfoRequest = {
  id?: string
}

export type GetProjectInfoResponse = {
  project?: ProjectV1Project.Project
}

export type UpdateProjectInfoRequest = {
  project?: ProjectV1Project.Project
}

export type UpdateProjectInfoResponse = {
  project?: ProjectV1Project.Project
}

export class ProjectService {
  static CreateProject(req: CreateProjectRequest, initReq?: fm.InitReq): Promise<CreateProjectResponse> {
    return fm.fetchReq<CreateProjectRequest, CreateProjectResponse>(`/gapi/team/v1/project/create`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetProjectList(req: GetProjectListRequest, initReq?: fm.InitReq): Promise<GetProjectListResponse> {
    return fm.fetchReq<GetProjectListRequest, GetProjectListResponse>(`/gapi/team/v1/project/list?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetProjectInfo(req: GetProjectInfoRequest, initReq?: fm.InitReq): Promise<GetProjectInfoResponse> {
    return fm.fetchReq<GetProjectInfoRequest, GetProjectInfoResponse>(`/gapi/team/v1/project/info?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static UpdateProjectInfo(req: UpdateProjectInfoRequest, initReq?: fm.InitReq): Promise<UpdateProjectInfoResponse> {
    return fm.fetchReq<UpdateProjectInfoRequest, UpdateProjectInfoResponse>(`/gapi/team/v1/project/info`, {...initReq, method: "PUT", body: JSON.stringify(req, fm.replacer)})
  }
}