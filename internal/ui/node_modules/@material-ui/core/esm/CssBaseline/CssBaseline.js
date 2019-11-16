import _extends from "@babel/runtime/helpers/esm/extends";
import React from 'react';
import PropTypes from 'prop-types';
import makeStyles from '../styles/makeStyles';
import { exactProp } from '@material-ui/utils';
var useStyles = makeStyles(function (theme) {
  return {
    '@global': {
      html: {
        WebkitFontSmoothing: 'antialiased',
        // Antialiasing.
        MozOsxFontSmoothing: 'grayscale',
        // Antialiasing.
        // Change from `box-sizing: content-box` so that `width`
        // is not affected by `padding` or `border`.
        boxSizing: 'border-box'
      },
      '*, *::before, *::after': {
        boxSizing: 'inherit'
      },
      'strong, b': {
        fontWeight: 'bolder'
      },
      body: _extends({
        margin: 0,
        // Remove the margin in all browsers.
        color: theme.palette.text.primary
      }, theme.typography.body2, {
        backgroundColor: theme.palette.background.default,
        '@media print': {
          // Save printer ink.
          backgroundColor: theme.palette.common.white
        },
        // Add support for document.body.requestFullScreen().
        // Other elements, if background transparent, are not supported.
        '&::backdrop': {
          backgroundColor: theme.palette.background.default
        }
      })
    }
  };
}, {
  name: 'MuiCssBaseline'
});
/**
 * Kickstart an elegant, consistent, and simple baseline to build upon.
 */

function CssBaseline(props) {
  var _props$children = props.children,
      children = _props$children === void 0 ? null : _props$children;
  useStyles();
  return React.createElement(React.Fragment, null, children);
}

process.env.NODE_ENV !== "production" ? CssBaseline.propTypes = {
  /**
   * You can wrap a node.
   */
  children: PropTypes.node
} : void 0;

if (process.env.NODE_ENV !== 'production') {
  // eslint-disable-next-line
  CssBaseline['propTypes' + ''] = exactProp(CssBaseline.propTypes);
}

export default CssBaseline;