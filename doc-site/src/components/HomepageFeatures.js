import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.scss';

const FeatureList = [
  {
    title: 'Reliable',
    Svg: require('../../static/img/icons/attr-reliable.svg').default,
    description: (
      <>
        Your infrastructure will be set up in multiple availability zones
        making it highly available and fault tolerant. All infrastructure is
        represented with code using HashiCorp Terraform so your
        environments are reproducible, auditable, and easy to configure.
      </>
    ),
  },
  {
    title: 'Scalable',
    Svg: require('../../static/img/icons/attr-scalable.svg').default,
    description: (
      <>
        Your services will be running in Kubernetes, with the EKS nodes 
        running in AWS Auto Scaling Group. The application workloads 
        and cluster size are ready to scale whenever the need arises. 
        Frontend assets will be served from AWS' Cloudfront CDN.
      </>
    ),
  },
  {
    title: 'Secure',
    Svg: require('../../static/img/icons/attr-secure.svg').default,
    description: (
      <>
        Properly configured access-control to resources/security groups. 
        Our practices are built on top of multiple security audits and 
        penetration tests. Automatic certificate management using 
        Let's Encrypt, database encryption, VPN support, and more means 
        your traffic will always be encrypted. 
      </>
    ),
  },
  {
    title: 'Modular',
    Svg: require('../../static/img/icons/attr-modular.svg').default,
    description: (
      <>
        Everything built by Zero is yours. After Zero generates your 
        infrastructure, backend, and frontend, all the code is checked 
        into your source control repositories and becomes the basis for 
        your new system. You can customize as much as you like. We 
        provide constant updates and new modules.
      </>
    ),
  },
];

const Feature = ({Svg, title, description}) => (
  <div className={clsx('col col--6') + " feature"}>
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
  return (
    <section className={`${styles.features} featured-sections`}>
      <div className="row">
        {FeatureList.map((props, idx) => (
          <Feature key={idx} {...props} />
        ))}
      </div>
    </section>
  )
}
