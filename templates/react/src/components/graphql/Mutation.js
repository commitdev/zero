import React from 'react'
import { Mutation } from 'react-apollo'
import { withSnackbar } from 'notistack'
import LoadingIndicator from '../LoadingIndicator'

class GraphQLMutation extends React.Component {
  render() {
    const {
      onCompleted = this._onCompleted,
      onError = this._onError,
      ...props
    } = this.props
    return (
      <Mutation
        {...props}
        optimisticResponse={LoadingIndicator}
        onCompleted={onCompleted}
        onError={onError}
      >
        {this.props.children}
      </Mutation>
    )
  }

  _onError = (err = {}) => {
    this.props.enqueueSnackbar(err.message || 'Unexpected Error', {
      variant: 'error',
    })
  }

  _onCompleted = () => {
    this.props.enqueueSnackbar(this.props.successMsg || 'success', {
      variant: 'success',
    })
  }
}

export default withSnackbar(GraphQLMutation)
