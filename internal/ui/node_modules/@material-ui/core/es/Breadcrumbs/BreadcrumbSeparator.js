import _extends from "@babel/runtime/helpers/esm/extends";
import _objectWithoutPropertiesLoose from "@babel/runtime/helpers/esm/objectWithoutPropertiesLoose";
import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import withStyles from '../styles/withStyles';
const styles = {
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
  const {
    classes,
    className
  } = props,
        other = _objectWithoutPropertiesLoose(props, ["classes", "className"]);

  return React.createElement("li", _extends({
    "aria-hidden": true,
    className: clsx(classes.root, className)
  }, other));
}

process.env.NODE_ENV !== "production" ? BreadcrumbSeparator.propTypes = {
  children: PropTypes.node.isRequired,
  classes: PropTypes.object.isRequired,
  className: PropTypes.string
} : void 0;
export default withStyles(styles, {
  name: 'PrivateBreadcrumbSeparator'
})(BreadcrumbSeparator);