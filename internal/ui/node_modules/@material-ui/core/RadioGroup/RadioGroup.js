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

var _FormGroup = _interopRequireDefault(require("../FormGroup"));

var _useForkRef = _interopRequireDefault(require("../utils/useForkRef"));

var _RadioGroupContext = _interopRequireDefault(require("./RadioGroupContext"));

var RadioGroup = _react.default.forwardRef(function RadioGroup(props, ref) {
  var actions = props.actions,
      children = props.children,
      name = props.name,
      valueProp = props.value,
      onChange = props.onChange,
      other = (0, _objectWithoutProperties2.default)(props, ["actions", "children", "name", "value", "onChange"]);

  var rootRef = _react.default.useRef(null);

  var _React$useRef = _react.default.useRef(valueProp != null),
      isControlled = _React$useRef.current;

  var _React$useState = _react.default.useState(props.defaultValue),
      valueState = _React$useState[0],
      setValue = _React$useState[1];

  var value = isControlled ? valueProp : valueState;

  if (process.env.NODE_ENV !== 'production') {
    // eslint-disable-next-line react-hooks/rules-of-hooks
    _react.default.useEffect(function () {
      if (isControlled !== (valueProp != null)) {
        console.error(["Material-UI: A component is changing ".concat(isControlled ? 'a ' : 'an un', "controlled RadioGroup to be ").concat(isControlled ? 'un' : '', "controlled."), 'Elements should not switch from uncontrolled to controlled (or vice versa).', 'Decide between using a controlled or uncontrolled RadioGroup ' + 'element for the lifetime of the component.', 'More info: https://fb.me/react-controlled-components'].join('\n'));
      }
    }, [valueProp, isControlled]);
  }

  _react.default.useImperativeHandle(actions, function () {
    return {
      focus: function focus() {
        var input = rootRef.current.querySelector('input:not(:disabled):checked');

        if (!input) {
          input = rootRef.current.querySelector('input:not(:disabled)');
        }

        if (input) {
          input.focus();
        }
      }
    };
  }, []);

  var handleRef = (0, _useForkRef.default)(ref, rootRef);

  var handleChange = function handleChange(event) {
    if (!isControlled) {
      setValue(event.target.value);
    }

    if (onChange) {
      onChange(event, event.target.value);
    }
  };

  return _react.default.createElement(_RadioGroupContext.default.Provider, {
    value: {
      name: name,
      onChange: handleChange,
      value: value
    }
  }, _react.default.createElement(_FormGroup.default, (0, _extends2.default)({
    role: "radiogroup",
    ref: handleRef
  }, other), children));
});

process.env.NODE_ENV !== "production" ? RadioGroup.propTypes = {
  /**
   * @ignore
   */
  actions: _propTypes.default.shape({
    current: _propTypes.default.object
  }),

  /**
   * The content of the component.
   */
  children: _propTypes.default.node,

  /**
   * The default `input` element value. Use when the component is not controlled.
   */
  defaultValue: _propTypes.default.any,

  /**
   * The name used to reference the value of the control.
   */
  name: _propTypes.default.string,

  /**
   * @ignore
   */
  onBlur: _propTypes.default.func,

  /**
   * Callback fired when a radio button is selected.
   *
   * @param {object} event The event source of the callback.
   * You can pull out the new value by accessing `event.target.value` (string).
   */
  onChange: _propTypes.default.func,

  /**
   * @ignore
   */
  onKeyDown: _propTypes.default.func,

  /**
   * Value of the selected radio button. The DOM API casts this to a string.
   */
  value: _propTypes.default.any
} : void 0;
var _default = RadioGroup;
exports.default = _default;