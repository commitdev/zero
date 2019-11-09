import React from 'react'
import { compose, withApollo } from 'react-apollo'
import PropTypes from 'prop-types'
import _pickBy from 'lodash.pickby'
import Grid from '@material-ui/core/Grid'

import { SimpleForm, Field } from './SimpleForm'
import Queries from '../constants/queries'
import Query from './graphql/Query'
import { trackException } from '../utils/error'
import formFields from './forms'
import { H5, P } from './Text'

class ExampleForm extends React.Component {
  render() {
    const { templateId, orgId } = this.props
    return (
      <Query query={Queries.GET_FORM} variables={{ templateId, orgId }}>
        {this._renderForm}
      </Query>
    )
  }

  _renderForm = ({ formTemplate, myProfile }) => {
    const { templateId } = this.props
    const formResponse =
      myProfile.formResponses.find(r => r.formTemplateId === templateId) || {}
    const { formFields = [] } = formResponse
    const fieldMap = formFields.reduce((curr, val) => {
      curr[val.name] = val.value
      return curr
    }, {})

    if (formTemplate) {
      return (
        <SimpleForm onSubmit={this._handleSubmit(formResponse)}>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <H5 gutterBottom>{formTemplate.title}</H5>
              {formTemplate.description && <P>{formTemplate.description}</P>}
            </Grid>
            {formTemplate.fieldTypes.map(this._renderField(fieldMap))}
            {this.props.children}
          </Grid>
        </SimpleForm>
      )
    }
    return null
  }

  _handleSubmit = formResponse => async formData => {
    try {
      const { orgId, templateId, onCompleted, client } = this.props
      const formFields = Object.keys(formData).reduce((arr, key) => {
        const value = formData[key]
        const subObj = {
          name: key,
          value: value instanceof File ? value.name : value,
        }
        return arr.concat(subObj)
      }, [])
      await client.mutate({
        mutation: Queries.SUBMIT_FORM_MUTATION,
        variables: _pickBy(
          {
            responseId: formResponse.id,
            templateId,
            formFields,
            orgId,
          },
          val => val !== undefined
        ),
      })
      if (onCompleted) {
        onCompleted(formData)
      }
    } catch (e) {
      trackException(e)
    }
  }

  _renderField = fieldMap => ({
    width = 4,
    fieldset,
    component,
    ...fieldProps
  }) => {
    const props = {
      ...fieldProps,
      defaultValue: fieldMap[fieldProps.name],
    }
    return (
      <Grid item xs={width}>
        {formFields[component] ? (
          React.createElement(formFields[component], props)
        ) : (
          <Field {...props} />
        )}
      </Grid>
    )
  }
}

ExampleForm.propTypes = {
  orgId: PropTypes.string.isRequired,
  templateId: PropTypes.string.isRequired,
  onCompleted: PropTypes.func,
  children: PropTypes.node,
}

export default compose(withApollo)(ExampleForm)
