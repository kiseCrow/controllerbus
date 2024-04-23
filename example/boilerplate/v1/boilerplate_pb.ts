// @generated by protoc-gen-es v1.9.0 with parameter "target=ts,ts_nocheck=false"
// @generated from file github.com/aperturerobotics/controllerbus/example/boilerplate/v1/boilerplate.proto (package boilerplate.v1, syntax proto3)
/* eslint-disable */

import type {
  BinaryReadOptions,
  FieldList,
  JsonReadOptions,
  JsonValue,
  PartialMessage,
  PlainMessage,
} from '@bufbuild/protobuf'
import { Message, proto3 } from '@bufbuild/protobuf'

/**
 * Boilerplate implements the boilerplate directive.
 *
 * @generated from message boilerplate.v1.Boilerplate
 */
export class Boilerplate extends Message<Boilerplate> {
  /**
   * MessageText is the message to print with the boilerplate.
   * This is an example field.
   * The keyword "message" prevents us from using that as the field name.
   *
   * @generated from field: string message_text = 1;
   */
  messageText = ''

  constructor(data?: PartialMessage<Boilerplate>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'boilerplate.v1.Boilerplate'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'message_text',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): Boilerplate {
    return new Boilerplate().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): Boilerplate {
    return new Boilerplate().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): Boilerplate {
    return new Boilerplate().fromJsonString(jsonString, options)
  }

  static equals(
    a: Boilerplate | PlainMessage<Boilerplate> | undefined,
    b: Boilerplate | PlainMessage<Boilerplate> | undefined,
  ): boolean {
    return proto3.util.equals(Boilerplate, a, b)
  }
}

/**
 * BoilerplateResult implements the boilerplate directive result.
 *
 * @generated from message boilerplate.v1.BoilerplateResult
 */
export class BoilerplateResult extends Message<BoilerplateResult> {
  /**
   * PrintedLen is the final length of the printed message.
   *
   * @generated from field: uint32 printed_len = 1;
   */
  printedLen = 0

  constructor(data?: PartialMessage<BoilerplateResult>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'boilerplate.v1.BoilerplateResult'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'printed_len',
      kind: 'scalar',
      T: 13 /* ScalarType.UINT32 */,
    },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): BoilerplateResult {
    return new BoilerplateResult().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): BoilerplateResult {
    return new BoilerplateResult().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): BoilerplateResult {
    return new BoilerplateResult().fromJsonString(jsonString, options)
  }

  static equals(
    a: BoilerplateResult | PlainMessage<BoilerplateResult> | undefined,
    b: BoilerplateResult | PlainMessage<BoilerplateResult> | undefined,
  ): boolean {
    return proto3.util.equals(BoilerplateResult, a, b)
  }
}
