import React from 'react';
import styles from './HomepageTrustedBy.module.scss';

const trustedByData = [
  {
    img: "img/partners/planworth.png",
    src: "https://www.planworth.co/",
  },
  {
    img: "img/partners/patch.png",
    src: "https://www.patch.io/",
  },
  {
    img: "img/partners/atlasone.png",
    src: "https://atlasone.ca/",
  },
  {
    img: "img/partners/placeholder.png",
    src: "https://placeholder.co/",
  },
]

const Carousel = ({data}) => (
  <ul className={styles.trusted}>
    {
      data.map((item, idx) => (
        <li key={idx}>
          <a href={item.src} target="_blank">
            <img src={item.img} />
          </a>
        </li>
      ))
    }
  </ul>
)

export default function TrustedByCarousel() {
  return <div className="featured-sections">
    <h3 className={styles.title}>Trusted By</h3>
    <Carousel data={trustedByData} />
  </div>
}