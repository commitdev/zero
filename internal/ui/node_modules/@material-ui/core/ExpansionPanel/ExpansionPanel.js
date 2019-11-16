"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.default = exports.styles = void 0;

var _extends2 = _interopRequireDefault(require("@babel/runtime/helpers/extends"));

var _toArray2 = _interopRequireDefault(require("@babel/runtime/helpers/toArray"));

var _objectWithoutProperties2 = _interopRequireDefault(require("@babel/runtime/helpers/objectWithoutProperties"));

var _react = _interopRequireDefault(require("react"));

var _propTypes = _interopRequireDefault(require("prop-types"));

var _clsx = _interopRequireDefault(require("clsx"));

var _utils = require("@material-ui/utils");

var _Collapse = _interopRequireDefault(require("../Collapse"));

var _Paper = _interopRequireDefault(require("../Paper"));

var _withStyles = _interopRequireDefault(require("../styles/withStyles"));

var _ExpansionPanelContext = _interopRequireDefault(require("./ExpansionPanelContext"));

var styles = function styles(theme) {
  var transition = {
    duration: theme.transitions.duration.shortest
  };
  return {
    /* Styles applied to the root element. */
    root: {
      position: 'relative',
      transition: theme.transitions.create(['margin'], transition),
      '&:before': {
        position: 'absolute',
        left: 0,
        top: -1,
        right: 0,
        height: 1,
        content: '""',
        opacity: 1,
        backgroundColor: theme.palette.divider,
        transition: theme.transitions.create(['opacity', 'background-color'], transition)
      },
      '&:first-child': {
        '&:before': {
          display: 'none'
        }
      },
      '&$expanded': {
        margin: '16px 0',
        '&:first-child': {
          marginTop: 0
        },
        '&:last-child': {
          marginBottom: 0
        },
        '&:before': {
          opacity: 0
        }
      },
      '&$expanded + &': {
        '&:before': {
          display: 'none'
        }
      },
      '&$disabled': {
        backgroundColor: theme.palette.action.disabledBackground
      }
    },

    /* Styles applied to the root element if `square={false}`. */
    rounded: {
      borderRadius: 0,
      '&:first-child': {
        borderTopLeftRadius: theme.shape.borderRadius,
        borderTopRightRadius: theme.shape.borderRadius
      },
      '&:last-child': {
        borderBottomLeftRadius: theme.shape.borderRadius,
        borderBottomRightRadius: theme.shape.borderRadius,
        // Fix a rendering issue on Edge
        '@supports (-ms-ime-align: auto)': {
          borderBottomLeftRadius: 0,
          borderBottomRightRadius: 0
        }
      }
    },

    /* Styles applied to the root element if `expanded={true}`. */
    expanded: {},

    /* Styles applied to the root element if `disabled={true}`. */
    disabled: {}
  };
};

exports.styles = styles;

var ExpansionPanel = _react.default.forwardRef(function ExpansionPanel(props, ref) {
  var childrenProp = props.children,
      classes = props.classes,
      className = props.className,
      _props$defaultExpande = props.defaultExpanded,
      defaultExpanded = _props$defaultExpande === void 0 ? false : _props$defaultExpande,
      _props$disabled = props.disabled,
      disabled = _props$disabled === void 0 ? false : _props$disabled,
      expandedProp = props.expanded,
      onChange = props.onChange,
      _props$square = props.square,
      square = _props$square === void 0 ? false : _props$square,
      _props$TransitionComp = props.TransitionComponent,
      TransitionComponent = _props$TransitionComp === void 0 ? _Collapse.default : _props$TransitionComp,
      TransitionProps = props.TransitionProps,
      other = (0, _objectWithoutProperties2.default)(props, ["children", "classes", "className", "defaultExpanded", "disabled", "expanded", "onChange", "square", "TransitionComponent", "TransitionProps"]);

  var _React$useRef = _react.default.useRef(expandedProp != null),
      isControlled = _React$useRef.current;

  var _React$useState = _react.default.useState(defaultExpanded),
      expandedState = _React$useState[0],
      setExpandedState = _React$useState[1];

  var expanded = isControlled ? expandedProp : expandedState;

  if (process.env.NODE_ENV !== 'production') {
    // eslint-disable-next-line react-hooks/rules-of-hooks
    _react.default.useEffect(function () {
      if (isControlled !== (expandedProp != null)) {
        console.error(["Material-UI: A component is changing ".concat(isControlled ? 'a ' : 'an un', "controlled ExpansionPanel to be ").concat(isControlled ? 'un' : '', "controlled."), 'Elements should not switch from uncontrolled to controlled (or vice versa).', 'Decide between using a controlled or uncontrolled ExpansionPanel ' + 'element for the lifetime of the component.', 'More info: https://fb.me/react-controlled-components'].join('\n'));
      }
    }, [expandedProp, isControlled]);
  }

  var handleChange = _react.default.useCallback(function (event) {
    if (!isControlled) {
      setExpandedState(!expanded);
    }

    if (onChange) {
      onChange(event, !expanded);
    }
  }, [expanded, isControlled, onChange]);

  var _React$Children$toArr = _react.default.Children.toArray(childrenProp),
      _React$Children$toArr2 = (0, _toArray2.default)(_React$Children$toArr),
      summary = _React$Children$toArr2[0],
      children = _React$Children$toArr2.slice(1);

  var contextValue = _react.default.useMemo(function () {
    return {
      expanded: expanded,
      disabled: disabled,
      toggle: handleChange
    };
  }, [expanded, disabled, handleChange]);

  return _react.default.createElement(_Paper.default, (0, _extends2.default)({
    className: (0, _clsx.default)(classes.root, className, expanded && classes.expanded, disabled && classes.disabled, !square && classes.rounded),
    ref: ref,
    square: square
  }, other), _react.default.createElement(_ExpansionPanelContext.default.Provider, {
    value: contextValue
  }, summary), _react.default.createElement(TransitionComponent, (0, _extends2.default)({
    in: expanded,
    timeout: "auto"
  }, TransitionProps), _react.default.createElement("div", {
    "aria-labelledby": summary.props.id,
    id: summary.props['aria-controls'],
    role: "region"
  }, children)));
});

process.env.NODE_ENV !== "production" ? ExpansionPanel.propTypes = {
  /**
   * The content of the expansion panel.
   */
  children: (0, _utils.chainPropTypes)(_propTypes.default.node.isRequired, function (props) {
    var summary = _react.default.Children.toArray(props.children)[0];

    if (summary.type === _react.default.Fragment) {
      return new Error("Material-UI: the ExpansionPanel doesn't accept a Fragment as a child. " + 'Consider providing an array instead.');
    }

    if (!_react.default.isValidElement(summary)) {
      return new Error('Material-UI: expected the first child of ExpansionPanel to be a valid element.');
    }

    return null;
  }),

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
   * If `true`, expands the panel by default.
   */
  defaultExpanded: _propTypes.default.bool,

  /**
   * If `true`, the panel will be displayed in a disabled state.
   */
  disabled: _propTypes.default.bool,

  /**
   * If `true`, expands the panel, otherwise collapse it.
   * Setting this prop enables control over the panel.
   */
  expanded: _propTypes.default.bool,

  /**
   * Callback fired when the expand/collapse state is changed.
   *
   * @param {object} event The event source of the callback.
   * @param {boolean} expanded The `expanded` state of the panel.
   */
  onChange: _propTypes.default.func,

  /**
   * @ignore
   */
  square: _propTypes.default.bool,

  /**
   * The component used for the collapse effect.
   */
  TransitionComponent: _propTypes.default.elementType,

  /**
   * Props applied to the `Transition` element.
   */
  TransitionProps: _propTypes.default.object
} : void 0;

var _default = (0, _withStyles.default)(styles, {
  name: 'MuiExpansionPanel'
})(ExpansionPanel);

exports.default = _default;