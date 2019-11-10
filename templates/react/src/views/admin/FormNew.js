import React, { Component } from 'react'

import FormBuilder from '../../components/admin/FormBuilder'
import Breadcrumbs from '../../components/Breadcrumbs'
import AdminNavbar from '../../components/admin/Navbar'
import Centered from '../../components/Centered'

class FormNew extends Component {
  render() {
    const { location } = this.props
    return (
      <React.Fragment>
        <AdminNavbar location={location} />
        <Centered horizontal container>
          <Breadcrumbs location={location} />
          <FormBuilder />
        </Centered>
      </React.Fragment>
    )
  }
}

export default FormNew
