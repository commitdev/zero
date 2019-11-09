This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Key Features

- Dynamic loading of application
- Session handling
- Simple Routing using React Enroute
- Redux state management
- Simple server side rendering (React-snap) and SEO tag management (React-helmet)
- Redux Saga for side effect management
- Minimalist theme for MUI
- Simple language management, designed to be compatible and easily swappable with i18next
- Logger / analytics service
- In memory Localstorage fallback
- ScourJs as a immutability helper (optional)

https://auth0.com/docs/universal-login
https://trellis.auth0.com/authorize?response_type=token&client_id=e8Sk7ToG3y7fJMpNN006iVhEU1W2yONN&redirect_uri=http://localhost:3000/auth/callback&state=STATE
&connection=CONNECTION

## Code Formatting

- following [AirBnB styleguide](https://github.com/airbnb/javascript)
- linting is built into react-scripts (CRA config source)[https://github.com/facebook/create-react-app/blob/master/packages/eslint-config-react-app/index.js]
- use prettier to automatically format code

## Styles

We are using a hybrid approach to styling. The preferred system is CSS modules for its simplicity and statically compiled performance. However MUI forces us to use their style and theming system.

- use MUI css-in-js styling for theme and theme dependent styling
- use CSS modules for components that are independent and reusable. Can support using css variable overrides to customize it.
- share variables using css variables
- using some global styles for reset css and libraries like animate.css

## Documentation

- use http://documentation.js.org/ it supports Flow syntax

## Formatting

## Tools:

We should strive to use minimalistic tools or write your own

- https://bundlephobia.com/result?p=dayjs@1.7.8
- https://github.com/zalmoxisus/redux-devtools-extension

## Bundling / Codesplitting

https://facebook.github.io/create-react-app/docs/code-splitting
https://reactjs.org/docs/code-splitting.html

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.<br>
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br>
You will also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.<br>
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.<br>
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br>
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).
