
```bash
set APP_NAME="your_app_name"

npx create-react-app %APP_NAME% --template typescript

cd %APP_NAME%
npm install --save typescript @types/node @types/react @types/react-dom @types/jest

npm install --save react-router-dom @types/react-router-dom
npm install --save react-bootstrap@next bootstrap@5.1.1 @types/react-bootstrap
npm install --save react-router-bootstrap @types/react-router-bootstrap # 这个不好用
npm install --save axios
npm install --save react-redux @types/react-redux 
npm install --save @reduxjs/toolkit 
npm install --save redux-thunk @types/redux-thunk

npm install --save-dev node-sass
```

Storybook

```bash
npx sb init
```