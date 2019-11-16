import _extends from "@babel/runtime/helpers/esm/extends";
import _objectWithoutProperties from "@babel/runtime/helpers/esm/objectWithoutProperties";
import React from 'react';
import clsx from 'clsx';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
export var styles = function styles(theme) {
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
var Skeleton = React.forwardRef(function Skeleton(props, ref) {
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
      other = _objectWithoutProperties(props, ["classes", "className", "component", "disableAnimate", "height", "variant", "width"]);

  return React.createElement(Component, _extends({
    ref: ref,
    className: clsx(classes.root, classes[variant], className, !disableAnimate && classes.animate)
  }, other, {
    style: _extends({
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
  classes: PropTypes.object.isRequired,

  /**
   * @ignore
   */
  className: PropTypes.string,

  /**
   * The component used for the root node.
   * Either a string to use a DOM element or a component.
   */
  component: PropTypes.elementType,

  /**
   * If `true` the animation effect is disabled.
   */
  disableAnimate: PropTypes.bool,

  /**
   * Height of the skeleton.
   * Useful when you don't want to adapt the skeleton to a text element but for instance a card.
   */
  height: PropTypes.oneOfType([PropTypes.number, PropTypes.string]),

  /**
   * The type of content that will be rendered.
   */
  variant: PropTypes.oneOf(['text', 'rect', 'circle']),

  /**
   * Width of the skeleton.
   * Useful when the skeleton is inside an inline element with no width of its own.
   */
  width: PropTypes.oneOfType([PropTypes.number, PropTypes.string])
} : void 0;
export default withStyles(styles, {
  name: 'MuiSkeleton'
})(Skeleton);