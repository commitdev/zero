"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;

var _react = _interopRequireDefault(require("react"));

var _createSvgIcon = _interopRequireDefault(require("./utils/createSvgIcon"));

var _default = (0, _createSvgIcon.default)(_react.default.createElement(_react.default.Fragment, null, _react.default.createElement("defs", null, _react.default.createElement("path", {
  id: "a",
  d: "M0 0h24v24H0z"
})), _react.default.createElement("path", {
  d: "M9 11.24V7.5C9 6.12 10.12 5 11.5 5S14 6.12 14 7.5v3.74c1.21-.81 2-2.18 2-3.74C16 5.01 13.99 3 11.5 3S7 5.01 7 7.5c0 1.56.79 2.93 2 3.74zm5.08 2.26H13v-6c0-.83-.67-1.5-1.5-1.5S10 6.67 10 7.5v10.74l-4.04-.85-1.21 1.23L10.13 24h8.67l1.07-7.62-5.79-2.88z"
})), 'TouchAppSharp');

exports.default = _default;