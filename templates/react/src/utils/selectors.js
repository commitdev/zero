import React from 'react'
import hoistNonReactStatic from 'hoist-non-react-statics'
import { getProp, keyBy } from './helpers'
import Interface from './Interface'
import PropTypes from 'prop-types'

// ref: https://reactjs.org/docs/higher-order-components.html
function withSelector(selector) {
  return WrappedComponent => {
    class Enhance extends React.PureComponent {
      render() {
        const { data, ...passThrough } = this.props
        const processedData = selector(this.props)
        return (
          <WrappedComponent
            {...passThrough}
            mutate={data.mutate}
            refetch={data.refetch}
            data={processedData}
          />
        )
      }
    }
    Enhance.displayName = `withSelector(${getDisplayName(WrappedComponent)})`
    hoistNonReactStatic(Enhance, WrappedComponent)

    return Enhance
  }
}

function getDisplayName(WrappedComponent) {
  return WrappedComponent.displayName || WrappedComponent.name || 'Component'
}

function createProgressSelector(forms, formTabs) {
  return function getFormProgress({ data }) {
    const formResponses = getProp(data, ['myProfile', 'formResponses']) || []

    return {
      orgId: getProp(data, ['myProfile', 'profile', 'orgId']),
      formProgress: new FormProgress({ forms, formResponses, formTabs }),
    }
  }
}

const FormProgressPropDefs = {
  forms: PropTypes.arrayOf(
    PropTypes.shape({
      formTemplateId: PropTypes.string,
    })
  ).isRequired,
  formResponses: PropTypes.arrayOf(
    PropTypes.shape({
      formTemplateId: PropTypes.string.isRequired,
    })
  ).isRequired,
  formTabs: PropTypes.arrayOf(
    PropTypes.shape({
      index: PropTypes.number.isRequired,
    })
  ).isRequired,
}

class FormProgress extends Interface {
  constructor(params) {
    super(params, FormProgressPropDefs)
    Object.assign(this, params)
    // TODO rename all step => tab, too confusing
    this.formResponseMap = keyBy(params.formResponses, 'formTemplateId')
    this.tabIndices = this.formTabs.map(s => s.index)
  }

  // returns the index of the first form in formTabs that's incomplete (memoized)
  getNextIndex() {
    return (
      this.nextIndex ||
      (this.nextIndex = this.forms.findIndex(
        form => !this.formResponseMap[form.formTemplateId]
      ))
    )
  }

  getNextPath() {
    const lastFormIndex = this.forms.length - 1
    return this.isComplete() ? lastFormIndex : this.getNextIndex()
  }

  // form is complete when no more forms to fill
  isComplete() {
    return this.getNextIndex() < 0
  }

  // returns the index of the active step tab
  getActiveTab() {
    if (this.isComplete()) return this.tabIndices.length - 1

    // find the latest step that's the form has
    let activeForm = this.tabIndices
      .slice()
      .reverse()
      .find(step => this.getNextIndex() >= step)
    return this.tabIndices.findIndex(step => step === activeForm)
  }

  // calculate the step progress based on the number of forms in the current step
  getTabProgress() {
    const nextIndex = this.getNextIndex()
    if (nextIndex < 0) {
      return 0
    }
    const maxIndex = this.tabIndices.findIndex(tabIndex => tabIndex > nextIndex)
    const max = this.tabIndices[maxIndex]
    const min = this.tabIndices[maxIndex - 1]
    if (maxIndex <= 0) {
      return 0
    }
    return Math.floor(((nextIndex - min) / (max - min)) * 100)
  }
}

export default {
  withSelector,
  getDisplayName,
  // getFormProgress,
}
