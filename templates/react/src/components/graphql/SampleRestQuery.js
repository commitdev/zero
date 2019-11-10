import React, { Component } from 'react'
import gql from 'graphql-tag'
import { Query } from 'react-apollo'

const PERSON_QUERY = gql`
  query luke {
    person @rest(type: "Person", path: "people/1/") {
      name
    }
  }
`

class SampleRestQuery extends Component {
  render() {
    return (
      <Query query={PERSON_QUERY}>
        {({ loading, error, data }) => {
          if (loading) return <div>Fetching</div>
          if (error) return <div>Error</div>

          const person = data.person

          return <div>{person.name}</div>
        }}
      </Query>
    )
  }
}

export default SampleRestQuery
