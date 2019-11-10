import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import LoadingIndicator from './LoadingIndicator'

export default function AuthorizedImage(props) {
  const [img, setImg] = useState('')

  async function fetchImage() {
    await fetch(props.url, props.requestOptions).then(res => {
      return res.blob()
    }).then(imageBlob => {
      let imgUrl = URL.createObjectURL(imageBlob)
      setImg(imgUrl)
    })
    .catch(err => {
      // console.log(err)
    })
  }

  useEffect(() => {
    setImg('')
    fetchImage()
  }, [props.url])

  function _handleImageClick() {
    window.open(img)
  }

  function _renderImage() {
    if (img === '') {
      return <LoadingIndicator />
    }

    return <img src={img} alt="" onClick={_handleImageClick} />
  }

  return _renderImage()
}

AuthorizedImage.propTypes = {
  url: PropTypes.string.isRequired,
  requestOptions: PropTypes.object.isRequired,
}
