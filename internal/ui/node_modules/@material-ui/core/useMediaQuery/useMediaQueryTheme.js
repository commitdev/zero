"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;

var _useMediaQuery = _interopRequireDefault(require("./useMediaQuery"));

// TODO to deprecate in v4.x and remove in v5
function useMediaQueryTheme() {
  return _useMediaQuery.default.apply(void 0, arguments);
}

var _default = useMediaQueryTheme;
exports.default = _default;