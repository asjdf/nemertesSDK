/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/
export type Org = {
  id?: string
  name?: string
  ownerId?: string
}

export type OrgMember = {
  id?: string
  orgId?: string
  userId?: string
  role?: string
}

export type Team = {
  id?: string
  name?: string
}

export type TeamMember = {
  id?: string
  teamId?: string
  userId?: string
  role?: string
}

export type TeamProject = {
  id?: string
  teamId?: string
  projectId?: string
}