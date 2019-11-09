import React from 'react'
import { Query } from 'react-apollo'
import LoadingIndicator from '../LoadingIndicator'

const GraphQLQuery = ({
  query,
  children,
  WrappedComponent,
  variables = {},
  props = {},
}) => (
  <Query query={query} variables={variables}>
    {({ loading, error, data, refetch }) => {
      if (loading) return <LoadingIndicator />
      if (error) {
        return <div>Error</div>
      }
      return WrappedComponent ? (
        <WrappedComponent {...data} {...props} refetch={refetch} />
      ) : (
        children({ ...data, ...props, refetch })
      )
    }}
  </Query>
)

export default GraphQLQuery
