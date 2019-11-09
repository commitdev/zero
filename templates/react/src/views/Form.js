import React, { Component } from 'react'
import { compose, withApollo } from 'react-apollo'
import { makeStyles } from '@material-ui/core/styles'
import Grid from '@material-ui/core/Grid'
import _pickBy from 'lodash.pickby'
import cx from 'classnames'

import { SimpleForm, Field, SubmitBtn } from '../components/SimpleForm'
import FormProgress from '../components/progress/FormProgress'
import Query from '../components/graphql/Query'
import Queries from '../constants/queries'
import { Text, P } from '../components/Text'
import formFields from '../components/forms'
import SurveyNavBar from '../components/survey/NavBar'
import SurveyCompleted from '../components/survey/Completed'
import { trackException } from '../utils/error'
import { router, helpers } from '../utils'

// TODO get userId from session JWT paylod
const userId = 123
const MINIMUM_PROGRESS = 0.02

const useStyles = makeStyles(theme => ({
  form: {
    display: 'flex',
    minHeight: '100vh',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  formContent: {
    boxSizing: 'border-box',
    padding: '5vh 7vw',
    flex: '1 0 auto',
    maxWidth: '640px',
  },
  fullWidth: {
    width: '100%',
  },
  actionBar: {
    flex: 0,
  },
}))

function Form({ params, client, children }) {
  const { id } = params
  const classes = useStyles()

  return (
    <React.Fragment>
      <Query
        query={Queries.GET_FORM}
        variables={{ userId, templateId: id }}
        props={{ id }}
      >
        {renderForm}
      </Query>
    </React.Fragment>
  )

  function renderForm({ formTemplate, formResponse, refetch }) {
    if (!formTemplate) return null
    // TODO this is creating a unnecessary 2x fetch on first load
    // the proper way to solve the missing formResponseId is to make sure that the submit response's id is being added back to cache
    if (!formResponse) refetch()

    // TODO all this formatting logic should be moved into a selector, only pass in what we need
    const { title, description, fieldTypes } = formTemplate
    const respFields = formResponse ? formResponse.formFields : []
    const formResponseId = formResponse ? formResponse.id : undefined
    const fieldValues = respFields.reduce((curr, val) => {
      curr[val.name] = val.value
      return curr
    }, {})

    const fieldSetsMap = new Map()
    for (let i = 0; i < fieldTypes.length; i++) {
      let fieldSetName = fieldTypes[i].fieldset || fieldTypes[i].name
      let fields = fieldSetsMap.get(fieldSetName) || []
      fieldSetsMap.set(fieldSetName, [...fields, fieldTypes[i]])
    }
    const fieldSets = Array.from(fieldSetsMap)
    // if (!fieldSets[params.step]) {
    // return params.step >= fieldSets.length ? <FormCompleted/> : null
    // }
    const [fieldName, fields] = fieldSets[params.step] || []

    const progress =
      parseInt(params.step) / fieldSets.length || MINIMUM_PROGRESS

    return (
      <SimpleForm
        className={classes.form}
        style={{ width: '100%' }}
        onSubmit={handleSubmit(formResponseId, fieldValues, params, client)}
      >
        <SurveyNavBar title={title} backHref={getHref(-1, params)} />

        {params.step >= fieldSets.length ? (
          <SurveyCompleted />
        ) : (
          <React.Fragment>
            <div className={classes.fullWidth}>
              <FormProgress variant="determinate" value={progress * 100} />
            </div>

            <div className={cx(classes.formContent, classes.fullWidth)}>
              <Grid container spacing={3}>
                <Grid item xs={12}>
                  {description && <P>{description}</P>}
                </Grid>
                {fields.map(renderField(fieldValues))}
                {children}
              </Grid>
            </div>

            <div className={classes.fullWidth}>
              <SubmitBtn
                variant="contained"
                color="primary"
                size="large"
                fullWidth
              >
                <Text>next</Text>
              </SubmitBtn>
            </div>
          </React.Fragment>
        )}
      </SimpleForm>
    )
  }
}

function renderField(fieldValues) {
  return ({ fieldset, component, ...fieldProps }, i) => {
    const props = {
      ...fieldProps,
      defaultValue: fieldValues[fieldProps.name],
    }

    return (
      <Grid item xs={12} key={`${fieldProps.name}-${i}`}>
        {formFields[component] ? (
          React.createElement(formFields[component], props)
        ) : (
          <Field {...props} />
        )}
      </Grid>
    )
  }
}

function handleSubmit(formResponseId, fieldValues, params, client) {
  return async function(fieldsData) {
    const formData = { ...fieldValues, ...fieldsData }
    try {
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
            responseId: formResponseId,
            templateId: params.id,
            formFields,
          },
          val => val !== undefined
        ),
      })
      router.push(getHref(1, params))
    } catch (e) {
      trackException(e)
    }
  }
}

function getHref(stepChange, params) {
  let nextStep = parseInt(params.step) + stepChange
  if (nextStep < 0) return `/app`
  else return `/app/forms/${params.id}/step/${nextStep}`
}

export default compose(withApollo)(Form)
