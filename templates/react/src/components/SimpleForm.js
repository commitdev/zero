import React from 'react'
import PropTypes from 'prop-types'
import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'
import InputLabel from '@material-ui/core/InputLabel'

/**
 * Stateless Bootstrap Form Manager
 * - inspired by Rails' SimpleForm
 */

export class SimpleForm extends React.PureComponent {
  constructor(props) {
    super(props)
    this.form = React.createRef()
  }

  render() {
    const { onSubmit, children = [], errors, ...props } = this.props
    return (
      <form {...props} onSubmit={this.onSubmit} ref={this.form}>
        {React.Children.map(children, this.wrapChildren)}
      </form>
    )
  }

  wrapChildren = childEl => {
    const { errors = {} } = this.props

    // For each children of Field type, match with the form error state and set the corresponding error messages
    // ref: https://mxstbr.blog/2017/02/react-children-deepdive/#manipulating-children
    if (childEl.props && (childEl.props.name || childEl.type === Field)) {
      return React.cloneElement(childEl, {
        error: Boolean(errors[childEl.props.name]),
        helperText: errors[childEl.props.name],
      })
    } else {
      return childEl
    }
  }

  onSubmit = e => {
    const { onSubmit, formTemplateId } = this.props

    e.preventDefault()
    const formData = getFormData(e.target)

    if (typeof onSubmit === 'function') {
      onSubmit(formData, formTemplateId)
    }
  }
}

export function Field({ name, label, errorText, helpText, ...fieldProps }) {
  return (
    <div>
      <InputLabel>{label || name}</InputLabel>
      <TextField
        id={name}
        name={name}
        helperText={errorText || helpText}
        {...fieldProps}
      />
    </div>
  )
}

export function SubmitBtn({ children = 'Submit', ...props }) {
  return (
    <Button color="primary" type="submit" {...props}>
      {children}
    </Button>
  )
}

export function getFormData(formEl) {
  const formData = new FormData(formEl)
  const formObj = {}
  formData.forEach((value, key) => {
    formObj[key] = value
  })
  return formObj
}

SimpleForm.propTypes = {
  onSubmit: PropTypes.func,
}
