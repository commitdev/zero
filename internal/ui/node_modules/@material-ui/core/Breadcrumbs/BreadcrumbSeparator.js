"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = void 0;

var _extends2 = _interopRequireDefault(require("@babel/runtime/helpers/extends"));

var _objectWithoutProperties2 = _interopRequireDefault(require("@babel/runtime/helpers/objectWithoutProperties"));

var _react = _interopRequireDefault(require("react"));

var _propTypes = _interopRequireDefault(require("prop-types"));

var _clsx = _interopRequireDefault(require("clsx"));

var _withStyles = _interopRequireDefault(require("../styles/withStyles"));

var styles = {
  root: {
    display: 'flex',
    userSelect: 'none',
    marginLeft: 8,
    marginRight: 8
  }
};
/**
 * @ignore - internal component.
 */

function BreadcrumbSeparator(props) {
  var classes = props.classes,
      className = props.className,
      other = (0, _objectWithoutProperties2.default)(props, ["classes", "className"]);
  return _react.default.createElement("li", (0, _extends2.default)({
    "aria-hidden": true,
    className: (0, _clsx.default)(classes.root, className)
  }, other));
}

process.env.NODE_ENV !== "production" ? BreadcrumbSeparator.propTypes = {
  children: _propTypes.default.node.isRequired,
  classes: _propTypes.default.object.isRequired,
  className: _propTypes.default.string
} : void 0;

var _default = (0, _withStyles.default)(styles, {
  name: 'PrivateBreadcrumbSeparator'
})(BreadcrumbSeparator);

exports.default = _default;