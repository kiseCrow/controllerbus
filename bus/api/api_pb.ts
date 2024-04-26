// @generated by protoc-gen-es-lite unknown with parameter "target=ts,ts_nocheck=false"
// @generated from file github.com/aperturerobotics/controllerbus/bus/api/api.proto (package bus.api, syntax proto3)
/* eslint-disable */

import {
  createMessageType,
  Message,
  MessageType,
  PartialFieldInfo,
} from '@aptre/protobuf-es-lite'
import { Info } from '../../controller/controller_pb.js'
import { DirectiveState } from '../../directive/directive_pb.js'

export const protobufPackage = 'bus.api'

/**
 * Config are configuration arguments.
 *
 * @generated from message bus.api.Config
 */
export interface Config extends Message<Config> {
  /**
   * EnableExecController enables the exec controller API.
   *
   * @generated from field: bool enable_exec_controller = 1;
   */
  enableExecController?: boolean
}

export const Config: MessageType<Config> = createMessageType({
  typeName: 'bus.api.Config',
  fields: [
    {
      no: 1,
      name: 'enable_exec_controller',
      kind: 'scalar',
      T: 8 /* ScalarType.BOOL */,
    },
  ] as readonly PartialFieldInfo[],
  packedByDefault: true,
})

/**
 * GetBusInfoRequest is the request type for GetBusInfo.
 *
 * @generated from message bus.api.GetBusInfoRequest
 */
export interface GetBusInfoRequest extends Message<GetBusInfoRequest> {}

export const GetBusInfoRequest: MessageType<GetBusInfoRequest> =
  createMessageType({
    typeName: 'bus.api.GetBusInfoRequest',
    fields: [] as readonly PartialFieldInfo[],
    packedByDefault: true,
  })

/**
 * GetBusInfoResponse is the response type for GetBusInfo.
 *
 * @generated from message bus.api.GetBusInfoResponse
 */
export interface GetBusInfoResponse extends Message<GetBusInfoResponse> {
  /**
   * RunningControllers is the list of running controllers.
   *
   * @generated from field: repeated controller.Info running_controllers = 1;
   */
  runningControllers?: Info[]
  /**
   * RunningDirectives is the list of running directives.
   *
   * @generated from field: repeated directive.DirectiveState running_directives = 2;
   */
  runningDirectives?: DirectiveState[]
}

export const GetBusInfoResponse: MessageType<GetBusInfoResponse> =
  createMessageType({
    typeName: 'bus.api.GetBusInfoResponse',
    fields: [
      {
        no: 1,
        name: 'running_controllers',
        kind: 'message',
        T: () => Info,
        repeated: true,
      },
      {
        no: 2,
        name: 'running_directives',
        kind: 'message',
        T: () => DirectiveState,
        repeated: true,
      },
    ] as readonly PartialFieldInfo[],
    packedByDefault: true,
  })
