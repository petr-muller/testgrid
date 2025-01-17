// @generated by protobuf-ts 2.8.2 with parameter long_type_string
// @generated from protobuf file "pb/test_status/test_status.proto" (package "testgrid.test_status", syntax proto3)
// tslint:disable
/**
 * @generated from protobuf enum testgrid.test_status.TestStatus
 */
export enum TestStatus {
  /**
   * Proto versions of test_status.py's GathererStatus
   * Note that: NO_RESULT is used to signal that there should be no change.
   * This must be updated every time a new GathererStatus is added.
   *
   * @generated from protobuf enum value: NO_RESULT = 0;
   */
  NO_RESULT = 0,
  /**
   * @generated from protobuf enum value: PASS = 1;
   */
  PASS = 1,
  /**
   * @generated from protobuf enum value: PASS_WITH_ERRORS = 2;
   */
  PASS_WITH_ERRORS = 2,
  /**
   * @generated from protobuf enum value: PASS_WITH_SKIPS = 3;
   */
  PASS_WITH_SKIPS = 3,
  /**
   * @generated from protobuf enum value: RUNNING = 4;
   */
  RUNNING = 4,
  /**
   * @generated from protobuf enum value: CATEGORIZED_ABORT = 5;
   */
  CATEGORIZED_ABORT = 5,
  /**
   * @generated from protobuf enum value: UNKNOWN = 6;
   */
  UNKNOWN = 6,
  /**
   * @generated from protobuf enum value: CANCEL = 7;
   */
  CANCEL = 7,
  /**
   * @generated from protobuf enum value: BLOCKED = 8;
   */
  BLOCKED = 8,
  /**
   * @generated from protobuf enum value: TIMED_OUT = 9;
   */
  TIMED_OUT = 9,
  /**
   * @generated from protobuf enum value: CATEGORIZED_FAIL = 10;
   */
  CATEGORIZED_FAIL = 10,
  /**
   * @generated from protobuf enum value: BUILD_FAIL = 11;
   */
  BUILD_FAIL = 11,
  /**
   * @generated from protobuf enum value: FAIL = 12;
   */
  FAIL = 12,
  /**
   * @generated from protobuf enum value: FLAKY = 13;
   */
  FLAKY = 13,
  /**
   * @generated from protobuf enum value: TOOL_FAIL = 14;
   */
  TOOL_FAIL = 14,
  /**
   * @generated from protobuf enum value: BUILD_PASSED = 15;
   */
  BUILD_PASSED = 15,
}
