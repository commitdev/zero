import React from 'react'
import PropTypes from 'prop-types'

import Table from '@material-ui/core/Table/Table'
import TableHead from '@material-ui/core/TableHead/TableHead'
import TableRow from '@material-ui/core/TableRow/TableRow'
import TableCell from '@material-ui/core/TableCell/TableCell'
import TableBody from '@material-ui/core/TableBody/TableBody'
import Paper from '@material-ui/core/Paper/Paper'
import { withStyles } from '@material-ui/core/styles'
import IconButton from '@material-ui/core/IconButton'
import DeleteIcon from '@material-ui/icons/Delete'
import EditIcon from '@material-ui/icons/Edit'

import { router } from '../../utils'
import { queries } from '../../constants'
import Mutation from '../graphql/Mutation'

const styles = theme => ({
  container: {
    width: '100%',
  },
  tableCell: {
    paddingRight: theme.spacing(1),
    minWidth: 100,
  },
  textField: {
    marginBottom: theme.spacing(1),
    width: '50%',
  },
})

class FormList extends React.PureComponent {
  render() {
    const { classes } = this.props
    return (
      <Paper className={classes.container}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell className={classes.tableCell}>Form ID</TableCell>
              <TableCell className={classes.tableCell}>Form Title</TableCell>
              <TableCell align="right" className={classes.tableCell}>
                Actions
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>{this._renderTableRow()}</TableBody>
        </Table>
      </Paper>
    )
  }

  _renderTableRow() {
    const { classes, formTemplates } = this.props
    if (formTemplates) {
      return formTemplates.map((template, i) => {
        return (
          <TableRow key={i}>
            <TableCell name="id" className={classes.tableCell}>
              {template.id}
            </TableCell>
            <TableCell name="title" className={classes.tableCell}>
              {template.title}
            </TableCell>
            <TableCell
              align="right"
              padding="checkbox"
              className={classes.tableCell}
            >
              <IconButton
                className={classes.button}
                aria-label="Edit"
                onClick={router.to(`/admin/forms/${template.id}`)}
              >
                <EditIcon />
              </IconButton>
              {this._renderDeleteButton(template.id)}
            </TableCell>
          </TableRow>
        )
      })
    }
    return
  }

  _renderDeleteButton(id) {
    return (
      <Mutation
        mutation={queries.DESTROY_TEMPLATE_MUTATION}
        variables={{ id }}
        refetchQueries={[{ query: queries.GET_ALL_FORM_TEMPLATES }]}
        successMsg="Deleted"
      >
        {mutation => (
          <IconButton
            className={this.props.classes.button}
            aria-label="Delete"
            onClick={mutation}
          >
            <DeleteIcon />
          </IconButton>
        )}
      </Mutation>
    )
  }
}

const formTemplatesType = PropTypes.shape({
  id: PropTypes.string,
  title: PropTypes.string,
})

FormList.propTypes = {
  classes: PropTypes.object.isRequired,
  formTemplates: PropTypes.arrayOf(formTemplatesType),
}

export default withStyles(styles)(FormList)
