import React from 'react'
import { withSnackbar } from 'notistack'
import PropTypes from 'prop-types'

import Grid from '@material-ui/core/Grid'
import Button from '@material-ui/core/Button'
import Table from '@material-ui/core/Table/Table'
import TableHead from '@material-ui/core/TableHead/TableHead'
import TableRow from '@material-ui/core/TableRow/TableRow'
import TableCell from '@material-ui/core/TableCell/TableCell'
import TableBody from '@material-ui/core/TableBody/TableBody'
import Paper from '@material-ui/core/Paper/Paper'
import AddIcon from '@material-ui/icons/Add'
import RemoveIcon from '@material-ui/icons/Remove'
import TextField from '@material-ui/core/TextField'
import Checkbox from '@material-ui/core/Checkbox'
import Select from '@material-ui/core/Select'
import MenuItem from '@material-ui/core/MenuItem'
import Queries from '../../constants/queries'
import Mutation from '../graphql/Mutation'

const DEFAULT_FIELD = {
  name: 'nameOfField',
  label: 'Label of field',
  helpText: 'Additional info',
  component: undefined,
  required: true,
  width: 0,
  fieldset: '',
}

class FormBuilder extends React.Component {
  state = {
    title: 'Form Name',
    fields: null,
    submitBtnStr: 'CREATE FORM',
  }

  componentDidMount() {
    if (this.props.formTemplate) {
      this.setState({
        id: this.props.formTemplate.id,
        title: this.props.formTemplate.title,
        fields: this.props.formTemplate.fieldTypes,
        submitBtnStr: 'UPDATE FORM',
      })
    } else {
      this.setState({
        fields: [DEFAULT_FIELD],
      })
    }
  }

  render() {
    return (
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <TextField
            required
            label="Form Title"
            value={this.state.title}
            margin="normal"
            variant="outlined"
            fullWidth
            onChange={this._onChangeFormName}
          />
        </Grid>
        <Grid item xs={12}>
          <Paper>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Field Name</TableCell>
                  <TableCell>Label</TableCell>
                  <TableCell>Value Type</TableCell>
                  <TableCell>Width</TableCell>
                  <TableCell>Required</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>{this._renderTableRow()}</TableBody>
            </Table>
          </Paper>
        </Grid>
        <Grid item xs={6}>
          <Button variant="outlined" onClick={this._addField} fullWidth>
            <AddIcon />
          </Button>
        </Grid>
        <Grid item xs={6}>
          <Button variant="outlined" onClick={this._removeField} fullWidth>
            <RemoveIcon />
          </Button>
        </Grid>
        <Grid item xs={12}>
          {this._renderSubmitButton()}
        </Grid>
      </Grid>
    )
  }

  _renderTableRow() {
    if (this.state.fields) {
      return this.state.fields.map((field, i) => {
        return (
          <TableRow key={i}>
            <TableCell>
              <TextField
                name="name"
                defaultValue={field.name}
                onChange={e => this._onChange(i, e)}
              />
            </TableCell>
            <TableCell>
              <TextField
                name="label"
                defaultValue={field.label}
                onChange={e => this._onChange(i, e)}
              />
            </TableCell>
            <TableCell>
              <Select
                name="component"
                value={field.component}
                onChange={e => this._onChange(i, e)}
              >
                <MenuItem value="">Text Field</MenuItem>
                <MenuItem value="TextArea">Text Area</MenuItem>
                <MenuItem value="Radio">Radio</MenuItem>
                <MenuItem value="CheckboxSelect">CheckboxSelect</MenuItem>
                <MenuItem value="FileUpload">File Upload</MenuItem>
                <MenuItem value="Currency">Currency</MenuItem>
                <MenuItem value="LabeledCheckbox">Labeled Checkbox</MenuItem>
              </Select>
            </TableCell>
            <TableCell>
              <TextField
                name="width"
                type="number"
                inputProps={{ min: 1, max: 12 }}
                defaultValue={field.width}
                onChange={e => this._onChangeWidth(i, e)}
              />
            </TableCell>
            <TableCell>
              <Checkbox
                name="required"
                checked={!!field.required}
                onChange={() => this._toggleRequired(i)}
              />
            </TableCell>
          </TableRow>
        )
      })
    }
    return null
  }

  _renderSubmitButton() {
    const { title, fields } = this.state
    const { id } = this.props
    return (
      <Mutation
        mutation={Queries.UPSERT_TEMPLATE_MUTATION}
        variables={{ id: id || null, title, fieldTypes: fields }}
      >
        {mutation => (
          <Button
            variant="contained"
            fullWidth
            color="primary"
            onClick={mutation}
          >
            {this.state.submitBtnStr}
          </Button>
        )}
      </Mutation>
    )
  }

  _onChangeFormName = event => {
    const value = event.target.value
    this.setState({ title: value })
  }

  _onChange = (index, event) => {
    let currentField = this.state.fields[index]
    currentField[event.target.name] = event.target.value
    const fields = this.state.fields
    fields[index] = currentField
    this.setState({ fields: fields })
  }

  _toggleRequired = index => {
    const fields = this.state.fields
    fields[index].required = !fields[index].required
    this.setState({ fields: fields })
  }

  _onChangeWidth = (index, event) => {
    let value = parseInt(event.target.value)
    if (isNaN(value)) value = null
    const fields = this.state.fields
    fields[index].width = value
    this.setState({ fields: fields })
  }

  _addField = () => {
    const fields = this.state.fields
    fields.push({ ...DEFAULT_FIELD })
    this.setState({ fields: fields })
  }

  _removeField = () => {
    const fields = this.state.fields
    fields.pop()
    this.setState({ fields: fields })
  }
}

FormBuilder.propTypes = {
  id: PropTypes.string,
  title: PropTypes.string,
  fields: PropTypes.shape({
    name: PropTypes.string,
    label: PropTypes.string,
    helpText: PropTypes.string,
    component: PropTypes.string,
    required: PropTypes.bool,
    width: PropTypes.number,
    fieldset: PropTypes.string,
  }),
  submitBtnStr: PropTypes.string,
}

export default withSnackbar(FormBuilder)
