"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = setRef;

// TODO: Make it private only in v5
function setRef(ref, value) {
  if (typeof ref === 'function') {
    ref(value);
  } else if (ref) {
    ref.current = value;
  }
}