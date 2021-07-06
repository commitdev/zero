/** @type {import('@docusaurus/types').DocusaurusConfig} */
const { stylesheets, misc } = require('@commitdev/zero-doc-site-common-elements');

const siteUrl = process.env.BUILD_DOMAIN ? `https://${process.env.BUILD_DOMAIN}` : 'https://staging.getzero.dev';
const baseUrl = '/';
const repositoryName = 'zero';

module.exports = {
  title: 'Zero',
  tagline: 'Opinionated infrastructure to take you from idea to production on day one',
  url: siteUrl,
  baseUrl,
  ...misc(),
  projectName: repositoryName,
  themeConfig: {

    colorMode: {
      defaultMode: 'dark',
    },
    navbar: {
      logo: {
        alt: 'Zero Logo',
        src: 'img/zero.svg',
      },
      items: [
        {
          to: '/docs/zero/about/overview',
          label: 'Docs',
          className: 'header-docs-link header-logo-24',
          position: 'right'
        },
        {
          href: 'https://slack.getzero.dev',
          label: 'Slack',
          className: 'header-slack-link header-logo-24',
          position: 'right',
        },
        {
          href: 'https://github.com/commitdev/zero',
          label: 'Github',
          className: 'header-github-link header-logo-24',
          position: 'right',
        },
      ],
    },
    footer: {
      links: [
        {
          items: [
            {
              to: '/docs/zero/about/overview',
              label: 'Docs',
              className: 'header-docs-link header-logo-24',
              position: 'right'
            },
            {
              href: 'https://slack.getzero.dev',
              label: 'Slack',
              className: 'header-slack-link header-logo-24',
              position: 'right',
            },
            {
              href: 'https://github.com/commitdev/zero',
              label: 'Github',
              className: 'header-github-link header-logo-24',
              position: 'right',
            },
          ],
        },
      ],
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          path: 'docs',
          routeBasePath: 'docs/zero/',
          include: ['**/*.md', '**/*.mdx'],
          editUrl: 'https://github.com/commitdev/zero/blob/main/doc-site/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
        debug: true,
      },
    ],
  ],
  plugins: [
    'docusaurus-plugin-sass'
  ],
  stylesheets: stylesheets(),
};
