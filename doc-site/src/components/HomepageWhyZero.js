import React from 'react';
import styles from './HomepageWhyZero.module.scss';

const reasons = [
  {
    logo: require('../../static/img/icons/reason-technical-founder.svg').default,
    text: `As a technical founder or the first technical hire at a startup, 
    your sole focus is to build the logic for your application and get it into customers’ 
    hands as quickly and reliably as possible. Yet you immediately face multiple hurdles 
    before even writing the first line of code. You’re forced to make many tech trade-offs, 
    leading to decision fatigue. You waste countless hours building boilerplate SaaS 
    features not adding direct value to your customers. You spend precious time picking 
    up unfamiliar tech, make wrong choices that result in costly refactoring or rebuilding 
    in the future, and are unaware of tools and best practices that would speed up your 
    product iteration.`,
  },
  {
    logo: require('../../static/img/icons/reason-tools.svg').default,
    text: `Zero was built by a team of engineers with many years of experience in building
    and scaling startups. We have faced all the problems you will and want to provide a way 
    for new startups to avoid all those pitfalls. We also want to help you learn about the 
    tech choices we made so your team can become proficient in some of the great tools we 
    have included. The system you get starts small but allows you to scale well into the 
    future when you need to.`,
  }
];

const Reasons = ({data}) => (
  <div className={`${styles.reasons}`}>
    {
      data.map((i, idx) => (
        <div key={idx} className={styles.reason}>
          <i.logo className={styles.reason_logo} alt="logo" />
          <p className={`${styles.description} description`}>{i.text}</p>
        </div>
      ))
    }
  </div>
)

export default function FeatureWhyZero () {
const title = "Why is Zero good for startups ?"
return <div className={`${styles.reasons_container} featured-sections`}>
    <h3>{title}</h3>
    <Reasons data={reasons} />
  </div>
}
