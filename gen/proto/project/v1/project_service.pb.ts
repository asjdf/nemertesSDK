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

export type DeleteProjectRequest = {
  id?: string
}

export type DeleteProjectResponse = {
}

export type GetHostListRequest = {
  projectId?: string
}

export type GetHostListResponse = {
  hosts?: ProjectV1Project.Host[]
}

export type GetHostInfoRequest = {
  projectId?: string
  hostId?: string
}

export type GetHostInfoResponse = {
  host?: ProjectV1Project.Host
}

export type GetProfileListRequest = {
  hostId?: string
  startTime?: Date
  endTime?: Date
  sampleType?: string
  sampleUnit?: string
}

export type GetProfileListResponse = {
  profiles?: ProjectV1Project.ProfileWithMeta[]
}

export type GetProfileListMetaListRequest = {
  hostId?: string
  startTime?: Date
  endTime?: Date
}

export type GetProfileListMetaListResponse = {
  list?: ProjectV1Project.ProfileMeta[]
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
  static DeleteProject(req: DeleteProjectRequest, initReq?: fm.InitReq): Promise<DeleteProjectResponse> {
    return fm.fetchReq<DeleteProjectRequest, DeleteProjectResponse>(`/gapi/team/v1/project/delete`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetHostList(req: GetHostListRequest, initReq?: fm.InitReq): Promise<GetHostListResponse> {
    return fm.fetchReq<GetHostListRequest, GetHostListResponse>(`/gapi/team/v1/project/host/list?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetHostInfo(req: GetHostInfoRequest, initReq?: fm.InitReq): Promise<GetHostInfoResponse> {
    return fm.fetchReq<GetHostInfoRequest, GetHostInfoResponse>(`/gapi/team/v1/project/host/info?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetProfileList(req: GetProfileListRequest, initReq?: fm.InitReq): Promise<GetProfileListResponse> {
    return fm.fetchReq<GetProfileListRequest, GetProfileListResponse>(`/gapi/team/v1/project/profile/list?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetProfileListMetaList(req: GetProfileListMetaListRequest, initReq?: fm.InitReq): Promise<GetProfileListMetaListResponse> {
    return fm.fetchReq<GetProfileListMetaListRequest, GetProfileListMetaListResponse>(`/gapi/team/v1/project/profile/list/meta/list?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}