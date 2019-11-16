import _objectWithoutProperties from "@babel/runtime/helpers/esm/objectWithoutProperties";
import _defineProperty from "@babel/runtime/helpers/esm/defineProperty";
import React from 'react';
import PropTypes from 'prop-types';
import capitalize from '../utils/capitalize';
import withStyles from '../styles/withStyles';
import useTheme from '../styles/useTheme';

var styles = function styles(theme) {
  var hidden = {
    display: 'none'
  };
  return theme.breakpoints.keys.reduce(function (acc, key) {
    acc["only".concat(capitalize(key))] = _defineProperty({}, theme.breakpoints.only(key), hidden);
    acc["".concat(key, "Up")] = _defineProperty({}, theme.breakpoints.up(key), hidden);
    acc["".concat(key, "Down")] = _defineProperty({}, theme.breakpoints.down(key), hidden);
    return acc;
  }, {});
};
/**
 * @ignore - internal component.
 */


function HiddenCss(props) {
  var children = props.children,
      classes = props.classes,
      className = props.className,
      lgDown = props.lgDown,
      lgUp = props.lgUp,
      mdDown = props.mdDown,
      mdUp = props.mdUp,
      only = props.only,
      smDown = props.smDown,
      smUp = props.smUp,
      xlDown = props.xlDown,
      xlUp = props.xlUp,
      xsDown = props.xsDown,
      xsUp = props.xsUp,
      other = _objectWithoutProperties(props, ["children", "classes", "className", "lgDown", "lgUp", "mdDown", "mdUp", "only", "smDown", "smUp", "xlDown", "xlUp", "xsDown", "xsUp"]);

  var theme = useTheme();

  if (process.env.NODE_ENV !== 'production') {
    if (Object.keys(other).length !== 0 && !(Object.keys(other).length === 1 && other.hasOwnProperty('ref'))) {
      console.error("Material-UI: unsupported props received ".concat(Object.keys(other).join(', '), " by `<Hidden />`."));
    }
  }

  var clsx = [];

  if (className) {
    clsx.push(className);
  }

  for (var i = 0; i < theme.breakpoints.keys.length; i += 1) {
    var breakpoint = theme.breakpoints.keys[i];
    var breakpointUp = props["".concat(breakpoint, "Up")];
    var breakpointDown = props["".concat(breakpoint, "Down")];

    if (breakpointUp) {
      clsx.push(classes["".concat(breakpoint, "Up")]);
    }

    if (breakpointDown) {
      clsx.push(classes["".concat(breakpoint, "Down")]);
    }
  }

  if (only) {
    var onlyBreakpoints = Array.isArray(only) ? only : [only];
    onlyBreakpoints.forEach(function (breakpoint) {
      clsx.push(classes["only".concat(capitalize(breakpoint))]);
    });
  }

  return React.createElement("div", {
    className: clsx.join(' ')
  }, children);
}

process.env.NODE_ENV !== "production" ? HiddenCss.propTypes = {
  /**
   * The content of the component.
   */
  children: PropTypes.node,

  /**
   * Override or extend the styles applied to the component.
   * See [CSS API](#css) below for more details.
   */
  classes: PropTypes.object.isRequired,

  /**
   * @ignore
   */
  className: PropTypes.string,

  /**
   * Specify which implementation to use.  'js' is the default, 'css' works better for
   * server-side rendering.
   */
  implementation: PropTypes.oneOf(['js', 'css']),

  /**
   * If true, screens this size and down will be hidden.
   */
  lgDown: PropTypes.bool,

  /**
   * If true, screens this size and up will be hidden.
   */
  lgUp: PropTypes.bool,

  /**
   * If true, screens this size and down will be hidden.
   */
  mdDown: PropTypes.bool,

  /**
   * If true, screens this size and up will be hidden.
   */
  mdUp: PropTypes.bool,

  /**
   * Hide the given breakpoint(s).
   */
  only: PropTypes.oneOfType([PropTypes.oneOf(['xs', 'sm', 'md', 'lg', 'xl']), PropTypes.arrayOf(PropTypes.oneOf(['xs', 'sm', 'md', 'lg', 'xl']))]),

  /**
   * If true, screens this size and down will be hidden.
   */
  smDown: PropTypes.bool,

  /**
   * If true, screens this size and up will be hidden.
   */
  smUp: PropTypes.bool,

  /**
   * If true, screens this size and down will be hidden.
   */
  xlDown: PropTypes.bool,

  /**
   * If true, screens this size and up will be hidden.
   */
  xlUp: PropTypes.bool,

  /**
   * If true, screens this size and down will be hidden.
   */
  xsDown: PropTypes.bool,

  /**
   * If true, screens this size and up will be hidden.
   */
  xsUp: PropTypes.bool
} : void 0;
export default withStyles(styles, {
  name: 'PrivateHiddenCss'
})(HiddenCss);