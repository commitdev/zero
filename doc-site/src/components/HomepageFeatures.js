import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.scss';

const FeatureList = [
  {
    title: 'Reliable',
    Svg: require('../../static/img/icons/attr-reliable.svg').default,
    description: (
      <>
        Fault-tolerant infrastructure. Production workloads will be 
        highly available, with traffic load balanced to multiple 
        instances of your application. All the infrastructure is 
        represented with code to be reproducible and easy to configure.
      </>
    ),
  },
  {
    title: 'Scalable',
    Svg: require('../../static/img/icons/attr-scalable.svg').default,
    description: (
      <>
        Your system will scale automatically based on your application’s 
        needs. For frontend assets, using a CDN will ensure availability 
        at global scale.
      </>
    ),
  },
  {
    title: 'Secure',
    Svg: require('../../static/img/icons/attr-secure.svg').default,
    description: (
      <>
        All your systems will follow security best practices backed up 
        by multiple security audits and penetration tests, and will be 
        properly configured to limit access to private networks, 
        secrets, and data. Bullet-proof your application by default 
        using existing, tested tools.
      </>
    ),
  },
];

const Feature = ({Svg, title, description}) => (
  <div className={clsx('col col--5') + " feature"}>
    <div className="text--center">
      <Svg className={styles.featureSvg} alt={title} />
    </div>
    <div className="text--center padding-horiz--md">
      <h3>{title}</h3>
      <p className="description">{description}</p>
    </div>
  </div>
)

export default function HomepageFeatures() {
  return (<>
    <section className={`${styles.features} featured-sections`}>

      <h2 className={styles.title}>Building something <strong>fast</strong> doesn't mean you can't also build it <strong>right</strong></h2>
      <div className={`${styles.row} row`}>
        {FeatureList.map((props, idx) => (
          <Feature key={idx} {...props} />
        ))}
      </div>
    </section>
  </>)
}
