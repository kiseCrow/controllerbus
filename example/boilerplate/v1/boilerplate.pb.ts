// @generated by protoc-gen-es-lite unknown with parameter "target=ts,ts_nocheck=false"
// @generated from file github.com/aperturerobotics/controllerbus/example/boilerplate/v1/boilerplate.proto (package boilerplate.v1, syntax proto3)
/* eslint-disable */

import type { MessageType, PartialFieldInfo } from '@aptre/protobuf-es-lite'
import { createMessageType, Message } from '@aptre/protobuf-es-lite'

export const protobufPackage = 'boilerplate.v1'

/**
 * Boilerplate implements the boilerplate directive.
 *
 * @generated from message boilerplate.v1.Boilerplate
 */
export type Boilerplate = Message<{
  /**
   * MessageText is the message to print with the boilerplate.
   * This is an example field.
   * The keyword "message" prevents us from using that as the field name.
   *
   * @generated from field: string message_text = 1;
   */
  messageText?: string
}>

export const Boilerplate: MessageType<Boilerplate> = createMessageType({
  typeName: 'boilerplate.v1.Boilerplate',
  fields: [
    {
      no: 1,
      name: 'message_text',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
  ] as readonly PartialFieldInfo[],
  packedByDefault: true,
})

/**
 * BoilerplateResult implements the boilerplate directive result.
 *
 * @generated from message boilerplate.v1.BoilerplateResult
 */
export type BoilerplateResult = Message<{
  /**
   * PrintedLen is the final length of the printed message.
   *
   * @generated from field: uint32 printed_len = 1;
   */
  printedLen?: number
}>

export const BoilerplateResult: MessageType<BoilerplateResult> =
  createMessageType({
    typeName: 'boilerplate.v1.BoilerplateResult',
    fields: [
      {
        no: 1,
        name: 'printed_len',
        kind: 'scalar',
        T: 13 /* ScalarType.UINT32 */,
      },
    ] as readonly PartialFieldInfo[],
    packedByDefault: true,
  })
