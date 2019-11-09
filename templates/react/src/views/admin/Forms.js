import React from 'react'
import Button from '@material-ui/core/Button'
import AddIcon from '@material-ui/icons/Add'

import { router } from '../../utils'
import Centered from '../../components/Centered'
import AdminNavbar from '../../components/admin/Navbar'
import GraphQLQuery from '../../components/graphql/Query'
import Queries from '../../constants/queries'
import FormList from '../../components/admin/FormList'
import Breadcrumbs from '../../components/Breadcrumbs'

class Forms extends React.PureComponent {
  render() {
    return (
      <React.Fragment>
        <AdminNavbar location={this.props.location} />
        <Centered horizontal container>
          <Breadcrumbs location={this.props.location} />
          <GraphQLQuery
            query={Queries.GET_ALL_FORM_TEMPLATES}
            WrappedComponent={FormList}
          />
          <br />
          <Button
            variant="outlined"
            onClick={router.to('/admin/forms/new')}
            fullWidth
          >
            <AddIcon />
          </Button>
        </Centered>
      </React.Fragment>
    )
  }
}

export default Forms
