// @generated by protoc-gen-es v1.9.0 with parameter "target=ts,ts_nocheck=false"
// @generated from file github.com/aperturerobotics/controllerbus/example/plugin-demo/demo-controller/config.proto (package hot.compile.demo.controller, syntax proto3)
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
 * Config is the boilerplate configuration.
 *
 * @generated from message hot.compile.demo.controller.Config
 */
export class Config extends Message<Config> {
  /**
   * ExitAfterDur requests an os.Exit(0) after a duration.
   *
   * @generated from field: string exit_after_dur = 1;
   */
  exitAfterDur = ''

  constructor(data?: PartialMessage<Config>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'hot.compile.demo.controller.Config'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    {
      no: 1,
      name: 'exit_after_dur',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): Config {
    return new Config().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): Config {
    return new Config().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): Config {
    return new Config().fromJsonString(jsonString, options)
  }

  static equals(
    a: Config | PlainMessage<Config> | undefined,
    b: Config | PlainMessage<Config> | undefined,
  ): boolean {
    return proto3.util.equals(Config, a, b)
  }
}
