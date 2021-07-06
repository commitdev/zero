import React from 'react';
import styles from './HomepageVideo.module.scss';

export default function FeatureVideo () {
return <div className={styles.video}>
    <iframe width="850" height="450" src="https://www.youtube.com/embed/IZ45Wm3wN7s" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
  </div>
}
