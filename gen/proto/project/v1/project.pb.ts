/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

export type Project = {
  id?: string
  orgId?: string
  name?: string
}

export type Host = {
  id?: string
  orgId?: string
  projectId?: string
  createdAt?: Date
  updatedAt?: Date
  name?: string
  serverName?: string
  version?: string
  environment?: string
}

export type Profile = {
  id?: string
  hostId?: string
  createdAt?: Date
  updatedAt?: Date
  compatibleHash?: string
  startedAt?: Date
  endedAt?: Date
}

export type ProfileMeta = {
  id?: string
  profileId?: string
  createdAt?: Date
  updatedAt?: Date
  sampleType?: string
  sampleUnit?: string
  totalValue?: string
}

export type ProfileWithMeta = {
  profile?: Profile
  metas?: ProfileMeta[]
}