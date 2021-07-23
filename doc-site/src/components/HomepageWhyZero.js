import React, { useState } from 'react';
import clsx from 'clsx';
import styles from './HomepageWhyZero.module.scss';

const reasons = [
  {
    logo: require('../../static/img/icons/reason-diamond.svg').default,
    title: 'Quality',
    text: `Like the best DevOps engineer you’ve ever met - except open source and free.`,
    details: [
      `The devops skill gap is real. Why spend precious time picking up unfamiliar tech, 
      making wrong choices that result in costly refactoring or rebuilding in the future, 
      and missing tools and best practices that would speed up your product iteration?`,
      `Get real-time support for all your questions from Zero’s community.`
    ]
  },
  {
    logo: require('../../static/img/icons/reason-clockwise.svg').default,
    title: 'Speed',
    text: `Just as fast as other tools like Heroku to get up and running.`,
    details: [
      `Building foundational infrastructure the right way doesn’t have to take a long time. Our team has years of experience building and scaling startups and have poured that knowledge into Zero. What used to take us weeks of DevOps work can now take you 30 minutes.`,
      `We provide constant updates and new modules that you can pull in on an ongoing basis.`,
    ]
  },
  {
    logo: require('../../static/img/icons/reason-key.svg').default,
    title: 'Ownership',
    text: `You own 100% of the code. No lock-in!`,
    details: [
      `Everything built by Zero is yours. It’s your code to change or migrate off at any point.`,
      `Cloud application hosting tools are built to lock you in and don’t scale. `,
      `Infrastructure is created in your cloud provider account. You can customize as much 
      as you like with no strings attached. You control how much you spend.`
    ]
  }
];

const Reasons = ({ data, expanded, setExpanded }) => (
  <div className={`${styles.reasons} row`}>
    {
      data.map((i, idx) => (
        <div key={idx} className={`${styles.reason} ${clsx('col col--3') }`}>
          <i.logo className={styles.reason_logo} alt="logo" />
          <h4 className={styles.title}>{i.title}</h4>

          <p className={`${styles.description} description`}>{i.text}</p>
          {expanded && <ul className={`${styles.description} description`}>{i.details.map(content=> <li>{content}</li>)}</ul>}
        </div>
      ))
    }
  </div>
)

export default function FeatureWhyZero () {
  const [expanded, setExpanded] = useState(false)
const title = "Why is Zero good for startups ?"
return <div className={`${styles.reasons_container} featured-sections`}>
    <h2 className={styles.title}>
      {title}
      <h5 className={styles.subtitle}>
        As engineer #1, your sole priority is to build the logic for your application and get it into customers’ hands as quickly and reliably as possible.
      </h5>
    </h2>
    <Reasons data={reasons} expanded={expanded} setExpanded={setExpanded} />
    <div className={`${styles.expand} ${expanded && styles.expanded}`}>
      <a href="javascript:void(0);" onClick={()=>{setExpanded(!expanded)}}>{expanded ? "Less" : "More" } Details</a>
    </div>
  </div>
}
