import React, { Component } from 'react'

import FormBuilder from '../../components/admin/FormBuilder'
import GraphQLQuery from '../../components/graphql/Query'
import Breadcrumbs from '../../components/Breadcrumbs'
import Queries from '../../constants/queries'
import AdminNavbar from '../../components/admin/Navbar'
import Centered from '../../components/Centered'

class FormEdit extends Component {
  render() {
    const { params, location } = this.props
    const id = params.id
    return (
      <React.Fragment>
        <AdminNavbar location={location} />
        <Centered container>
          <Breadcrumbs location={location} />
          <GraphQLQuery
            query={Queries.GET_FORM_TEMPLATE}
            variables={{ id }}
            WrappedComponent={FormBuilder}
            props={{ id }}
          />
        </Centered>
      </React.Fragment>
    )
  }
}

export default FormEdit
