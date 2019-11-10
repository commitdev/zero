import React from 'react'
import { compose } from 'react-apollo'
import { withSnackbar } from 'notistack'
import PropTypes from 'prop-types'
import Dropzone from 'react-dropzone'
import classNames from 'classnames'

import Grid from '@material-ui/core/Grid'
import { withStyles } from '@material-ui/core/styles'
import CloudUploadIcon from '@material-ui/icons/CloudUpload'
import Button from '@material-ui/core/Button'
import DeleteIcon from '@material-ui/icons/Delete'
import { Typography } from '@material-ui/core'
import InputLabel from '@material-ui/core/InputLabel'

import LoadingIndicator from '../LoadingIndicator'
import { auth } from '../../utils'

const styles = theme => ({
  dropZone: {
    position: 'relative',
    width: '100%',
    minHeight: '100px',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignContent: 'center',
    padding: theme.spacing(3),
    margin: `${theme.spacing(1)}px 0`,
    borderWidth: 1,
    border: 'dashed',
    borderColor: theme.palette.grey[300],
    boxSizing: 'border-box',
    cursor: 'grab',
  },
  stripes: {
    backgroundImage: `repeating-linear-gradient(-45deg, #F0F0F0, #F0F0F0 25px, #C8C8C8 25px, #C8C8C8 50px)`,
    animation: 'progress 2s linear infinite !important',
    backgroundSize: '150% 100%',
  },
  rejectStripes: {
    backgroundImage: `repeating-linear-gradient(-45deg, #fc8785, #fc8785 25px, #f4231f 25px, #f4231f 50px)`,
    animation: 'progress 2s linear infinite !important',
    backgroundSize: '150% 100%',
  },
  removeBtn: {
    backgroundColor: 'black',
    color: 'white',
    '&:hover': {
      backgroundColor: 'white',
      color: 'black',
      border: 'solid',
      borderWidth: 1,
      borderColor: 'black',
    },
  },
  smallPreviewImg: {
    height: 120,
    width: 'initial',
    maxWidth: '100%',
    color: 'rgba(0, 0, 0, 0.87)',
    transition: 'all 450ms cubic-bezier(0.23, 1, 0.32, 1) 0ms',
    boxSizing: 'border-box',
    boxShadow: 'rgba(0, 0, 0, 0.12) 0 1px 6px, rgba(0, 0, 0, 0.12) 0 1px 4px',
    borderRadius: theme.shape.borderRadius,
    zIndex: 5,
    opacity: 1,
  },
  imageContainer: {
    position: 'relative',
    zIndex: 10,
    marginBottom: theme.spacing(3),
    '&:hover $smallPreviewImg': {
      opacity: 0.3,
    },
    '&:hover $viewBtn': {
      opacity: 1,
    },
  },
  viewBtn: {
    transition: '.5s ease',
    position: 'absolute',
    opacity: 0,
    top: 50,
    left: '30%',
    backgroundColor: 'white',
    color: 'black',
    '&:hover': {
      backgroundColor: 'white',
      color: 'black',
    },
  },
  paper: {
    position: 'absolute',
    backgroundColor: theme.palette.background.paper,
    boxShadow: theme.shadows[5],
    padding: theme.spacing(4),
    outline: 'none',
  },
  inputLabel: {
    fontSize: '0.9rem',
  },
})

class FileUploadField extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      fileObjects: [],
      loading: false,
    }
  }

  isLoading = () => {
    this.setState({ loading: true })
  }

  bytesToSize = bytes => {
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
    if (bytes === 0) return 'n/a'
    const i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)), 10)
    if (i === 0) return `${bytes} ${sizes[i]})`
    return `${(bytes / 1024 ** i).toFixed(1)} ${sizes[i]}`
  }

  async onDrop(files) {
    if (this.state.fileObjects.length + files.length > this.props.filesLimit) {
      this.props.enqueueSnackbar(
        `Maximum allowed number of files exceeded. Only ${this.props.filesLimit} allowed`,
        { variant: 'error' }
      )
    } else {
      const file = files[0]

      this.isLoading()

      let uploadedFile = file.name
      const formData = new FormData()
      formData.append('upfile', file)
      const { jwt } = await auth.getToken()

      uploadedFile = await fetch(
        `${process.env.REACT_APP_FILE_HOST}/v1/file/upload`,
        {
          method: 'POST',
          headers: new Headers({
            Authorization: `Bearer ${jwt}`,
          }),
          body: formData,
        }
      )
        .then(r => r.json())
        .then(namespacedFilename => {
          let filename = namespacedFilename.split('/')[1]
          this.setState({ value: filename })
        })
        .catch(err => {
          const errorMsg = err.message
          this.setState({ loading: false })
          this.props.enqueueSnackbar(errorMsg, { variant: 'error' })
        })

      this.setState({ loading: false })
      const reader = new FileReader()
      reader.onload = event => {
        this.setState({
          fileObjects: this.state.fileObjects.concat({
            file,
            data: event.target.result,
            uploadedFile,
          }),
        })
        this.props.enqueueSnackbar(
          `File ${file.name} successfully uploaded. `,
          { variant: 'success' }
        )
      }
      reader.readAsDataURL(file)
    }
  }

  handleRemove = (fileIndex, event) => {
    const { fileObjects } = this.state
    const file = fileObjects.filter((fileObject, i) => {
      return i === fileIndex
    })[0].file
    fileObjects.splice(fileIndex, 1)
    this.setState(fileObjects, () => {
      this.props.enqueueSnackbar(`File ${file.name} removed`)
    })
  }

  handleDropRejected(rejectedFiles, event) {
    var message = ''
    rejectedFiles.forEach(rejectedFile => {
      message = `File ${rejectedFile.name} was rejected. `
      if (!this.props.acceptedFiles.includes(rejectedFile.type)) {
        message += 'File type not supported. '
      }
      if (rejectedFile.size > this.props.maxFileSize) {
        message +=
          'File is too big. Size limit is ' +
          this.bytesToSize(this.props.maxFileSize) +
          '. '
      }
    })
    if (this.props.onDropRejected) {
      this.props.onDropRejected(rejectedFiles, event)
    }

    this.props.enqueueSnackbar(message, { variant: 'error' })
  }

  render() {
    const { classes, name, label, helperText, dropzoneText, value } = this.props

    return (
      <React.Fragment>
        <InputLabel className={classes.inputLabel}>{label || name}</InputLabel>
        <Dropzone
          accept={this.props.acceptedFiles.join(',')}
          onDrop={files => this.onDrop(files)}
          onDropRejected={(rejectedFiles, event) =>
            this.handleDropRejected(rejectedFiles, event)
          }
          maxSize={this.props.maxFileSize}
          disableClick
        >
          {({
            getRootProps,
            getInputProps,
            isDragActive,
            isDragReject,
            open,
          }) => {
            let styles = classes.dropZone
            styles = isDragActive
              ? classNames(classes.dropZone, classes.stripes)
              : styles
            styles = isDragReject
              ? classNames(classes.dropZone, classes.rejectStripes)
              : styles
            return (
              <div {...getRootProps()} className={styles}>
                <input
                  type="hidden"
                  name={name}
                  value={this.state.value || value}
                />
                <input {...getInputProps()} />

                {this.state.loading && <LoadingIndicator />}
                {this.state.fileObjects.map((fileObject, index) => {
                  return (
                    <Grid
                      container
                      spacing={8}
                      direction="row"
                      justify="space-between"
                    >
                      <Grid
                        item
                        xs={6}
                        key={index}
                        className={classes.imageContainer}
                      >
                        <img
                          className={classes.smallPreviewImg}
                          alt="uploaded file"
                          src={fileObject.data || value}
                        />
                        <Typography variant="subtitle2">
                          {fileObject.file.name}
                        </Typography>
                      </Grid>
                      <Grid item xs={4} key={index}>
                        <Button
                          className={classes.removeBtn}
                          size="small"
                          onClick={event => this.handleRemove(index, event)}
                        >
                          <DeleteIcon /> Remove
                        </Button>
                      </Grid>
                    </Grid>
                  )
                })}

                <Grid container>
                  <Grid item xs={8}>
                    <Typography variant="body2" gutterBottom={true}>
                      {helperText || dropzoneText}
                    </Typography>
                    <Typography variant="caption">
                      Accepted formats: JPEG, PNG
                    </Typography>
                  </Grid>
                  <Grid item xs={4}>
                    <Button onClick={() => open()} variant="outlined">
                      Upload&nbsp;
                      <CloudUploadIcon />
                    </Button>
                  </Grid>
                </Grid>
              </div>
            )
          }}
        </Dropzone>
      </React.Fragment>
    )
  }
}

FileUploadField.defaultProps = {
  acceptedFiles: ['image/jpeg'],
  filesLimit: 2,
  maxFileSize: 3000000, // 3 MB
  dropzoneText: 'Upload a copy of your document',
}

FileUploadField.propTypes = {
  name: PropTypes.string.isRequired,
  label: PropTypes.node,
  value: PropTypes.string,
  helperText: PropTypes.string,
  defaultValue: PropTypes.string,
  acceptedFiles: PropTypes.array,
  filesLimit: PropTypes.number,
  maxFileSize: PropTypes.number,
  dropzoneText: PropTypes.string,
}

export default compose(
  withStyles(styles),
  withSnackbar
)(FileUploadField)
