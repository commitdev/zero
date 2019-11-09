import React from 'react'
import PropTypes from 'prop-types'
import FormHelperText from '@material-ui/core/FormHelperText'
import FormControl from '@material-ui/core/FormControl'
import Input from '@material-ui/core/Input'
import NativeSelect from '@material-ui/core/NativeSelect'
import FormLabel from '@material-ui/core/FormLabel'

// TODO temporary solution, the best way is to probably strip the annoying label styles from MUI all together
const formLabelStyle = {
  fontSize: '90%',
  marginBottom: -18,
  paddingBottom: 6,
  marginTop: -2,
}

class NativeSelectField extends React.Component {
  render() {
    const {
      name,
      label,
      allowEmpty,
      options = [],
      helperText,
      defaultValue,
      ...props
    } = this.props
    const fieldId = `${name || Date.now()}-field`
    return (
      <FormControl>
        <FormLabel htmlFor={fieldId} style={formLabelStyle}>
          {label || name}
        </FormLabel>
        <NativeSelect
          defaultValue={defaultValue}
          input={<Input name={name} id={fieldId} />}
          {...props}
        >
          {allowEmpty && <option value="" />}
          {options.map((option, i) => (
            <option value={option.value} key={i}>
              {option.label === undefined || option.label === null
                ? option.value
                : option.label}
            </option>
          ))}
        </NativeSelect>
        {helperText && <FormHelperText>{helperText}</FormHelperText>}
      </FormControl>
    )
  }
}

NativeSelectField.propTypes = {
  defaultValue: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
  allowEmpty: PropTypes.bool,
  name: PropTypes.string,
  label: PropTypes.node,
  helperText: PropTypes.string,
  options: PropTypes.arrayOf(
    PropTypes.shape({
      value: PropTypes.string.isRequired,
      label: PropTypes.string,
    })
  ),
}

export default NativeSelectField
