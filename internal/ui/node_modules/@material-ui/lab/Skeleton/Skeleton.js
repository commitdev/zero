"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = exports.styles = void 0;

var _extends2 = _interopRequireDefault(require("@babel/runtime/helpers/extends"));

var _objectWithoutProperties2 = _interopRequireDefault(require("@babel/runtime/helpers/objectWithoutProperties"));

var _react = _interopRequireDefault(require("react"));

var _clsx = _interopRequireDefault(require("clsx"));

var _propTypes = _interopRequireDefault(require("prop-types"));

var _styles = require("@material-ui/core/styles");

var styles = function styles(theme) {
  return {
    /* Styles applied to the root element. */
    root: {
      display: 'block',
      backgroundColor: theme.palette.action.hover,
      height: '1.2em'
    },

    /* Styles applied to the root element if `variant="text"`. */
    text: {
      marginTop: '0.8em',
      marginBottom: '0.8em',
      borderRadius: theme.shape.borderRadius
    },

    /* Styles applied to the root element if `variant="rect"`. */
    rect: {},

    /* Styles applied to the root element if `variant="circle"`. */
    circle: {
      borderRadius: '50%'
    },

    /* Styles applied to the root element if `disabledAnimate={false}`. */
    animate: {
      animation: '$animate 1.5s ease-in-out infinite'
    },
    '@keyframes animate': {
      '0%': {
        opacity: 1
      },
      '50%': {
        opacity: 0.4
      },
      '100%': {
        opacity: 1
      }
    }
  };
};

exports.styles = styles;

var Skeleton = _react.default.forwardRef(function Skeleton(props, ref) {
  var classes = props.classes,
      className = props.className,
      _props$component = props.component,
      Component = _props$component === void 0 ? 'div' : _props$component,
      _props$disableAnimate = props.disableAnimate,
      disableAnimate = _props$disableAnimate === void 0 ? false : _props$disableAnimate,
      height = props.height,
      _props$variant = props.variant,
      variant = _props$variant === void 0 ? 'text' : _props$variant,
      width = props.width,
      other = (0, _objectWithoutProperties2.default)(props, ["classes", "className", "component", "disableAnimate", "height", "variant", "width"]);
  return _react.default.createElement(Component, (0, _extends2.default)({
    ref: ref,
    className: (0, _clsx.default)(classes.root, classes[variant], className, !disableAnimate && classes.animate)
  }, other, {
    style: (0, _extends2.default)({
      width: width,
      height: height
    }, other.style)
  }));
});

process.env.NODE_ENV !== "production" ? Skeleton.propTypes = {
  /**
   * Override or extend the styles applied to the component.
   * See [CSS API](#css) below for more details.
   */
  classes: _propTypes.default.object.isRequired,

  /**
   * @ignore
   */
  className: _propTypes.default.string,

  /**
   * The component used for the root node.
   * Either a string to use a DOM element or a component.
   */
  component: _propTypes.default.elementType,

  /**
   * If `true` the animation effect is disabled.
   */
  disableAnimate: _propTypes.default.bool,

  /**
   * Height of the skeleton.
   * Useful when you don't want to adapt the skeleton to a text element but for instance a card.
   */
  height: _propTypes.default.oneOfType([_propTypes.default.number, _propTypes.default.string]),

  /**
   * The type of content that will be rendered.
   */
  variant: _propTypes.default.oneOf(['text', 'rect', 'circle']),

  /**
   * Width of the skeleton.
   * Useful when the skeleton is inside an inline element with no width of its own.
   */
  width: _propTypes.default.oneOfType([_propTypes.default.number, _propTypes.default.string])
} : void 0;

var _default = (0, _styles.withStyles)(styles, {
  name: 'MuiSkeleton'
})(Skeleton);

exports.default = _default;