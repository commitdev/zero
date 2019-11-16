/** @license Material-UI v4.6.0
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */
"use strict";

var _interopRequireWildcard = require("@babel/runtime/helpers/interopRequireWildcard");

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
Object.defineProperty(exports, "createGenerateClassName", {
  enumerable: true,
  get: function get() {
    return _createGenerateClassName.default;
  }
});
Object.defineProperty(exports, "createStyles", {
  enumerable: true,
  get: function get() {
    return _createStyles.default;
  }
});
Object.defineProperty(exports, "getThemeProps", {
  enumerable: true,
  get: function get() {
    return _getThemeProps.default;
  }
});
Object.defineProperty(exports, "jssPreset", {
  enumerable: true,
  get: function get() {
    return _jssPreset.default;
  }
});
Object.defineProperty(exports, "makeStyles", {
  enumerable: true,
  get: function get() {
    return _makeStyles.default;
  }
});
Object.defineProperty(exports, "mergeClasses", {
  enumerable: true,
  get: function get() {
    return _mergeClasses.default;
  }
});
Object.defineProperty(exports, "ServerStyleSheets", {
  enumerable: true,
  get: function get() {
    return _ServerStyleSheets.default;
  }
});
Object.defineProperty(exports, "styled", {
  enumerable: true,
  get: function get() {
    return _styled.default;
  }
});
Object.defineProperty(exports, "StylesProvider", {
  enumerable: true,
  get: function get() {
    return _StylesProvider.default;
  }
});
Object.defineProperty(exports, "ThemeProvider", {
  enumerable: true,
  get: function get() {
    return _ThemeProvider.default;
  }
});
Object.defineProperty(exports, "useTheme", {
  enumerable: true,
  get: function get() {
    return _useTheme.default;
  }
});
Object.defineProperty(exports, "withStyles", {
  enumerable: true,
  get: function get() {
    return _withStyles.default;
  }
});
Object.defineProperty(exports, "withTheme", {
  enumerable: true,
  get: function get() {
    return _withTheme.default;
  }
});
Object.defineProperty(exports, "withThemeCreator", {
  enumerable: true,
  get: function get() {
    return _withTheme.withThemeCreator;
  }
});

var _utils = require("@material-ui/utils");

var _createGenerateClassName = _interopRequireDefault(require("./createGenerateClassName"));

var _createStyles = _interopRequireDefault(require("./createStyles"));

var _getThemeProps = _interopRequireDefault(require("./getThemeProps"));

var _jssPreset = _interopRequireDefault(require("./jssPreset"));

var _makeStyles = _interopRequireDefault(require("./makeStyles"));

var _mergeClasses = _interopRequireDefault(require("./mergeClasses"));

var _ServerStyleSheets = _interopRequireDefault(require("./ServerStyleSheets"));

var _styled = _interopRequireDefault(require("./styled"));

var _StylesProvider = _interopRequireDefault(require("./StylesProvider"));

var _ThemeProvider = _interopRequireDefault(require("./ThemeProvider"));

var _useTheme = _interopRequireDefault(require("./useTheme"));

var _withStyles = _interopRequireDefault(require("./withStyles"));

var _withTheme = _interopRequireWildcard(require("./withTheme"));

/* Warning if there are several instances of @material-ui/styles */
if (process.env.NODE_ENV !== 'production' && process.env.NODE_ENV !== 'test' && typeof window !== 'undefined') {
  _utils.ponyfillGlobal['__@material-ui/styles-init__'] = _utils.ponyfillGlobal['__@material-ui/styles-init__'] || 0;

  if (_utils.ponyfillGlobal['__@material-ui/styles-init__'] === 1) {
    console.warn(['It looks like there are several instances of `@material-ui/styles` initialized in this application.', 'This may cause theme propagation issues, broken class names, ' + 'specificity issues, and makes your application bigger without a good reason.', '', 'See https://material-ui.com/r/styles-instance-warning for more info.'].join('\n'));
  }

  _utils.ponyfillGlobal['__@material-ui/styles-init__'] += 1;
}