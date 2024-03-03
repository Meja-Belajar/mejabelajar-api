# Meja Belajar

## About
This repository is collectively owned by all contributors and serves as a university project developed as part of the Software Engineering [COMP6100001] course. The project's objective is to equip students with the essential knowledge and skills required for success in software development, covering various aspects such as the software development lifecycle, process models, project management, architecture, and quality assurance.

Each contributor in this repository plays a vital role in the Software Engineering project, with every group member contributing based on their assigned roles and responsibilities

## Idea 
Meja Belajar is a web application designed to facilitate university students in becoming mentors or mentees. Mentors are students with excellent academic performance who assist other students in enhancing their academic excellence. This application will offer features to connect mentors with mentees, fostering a supportive learning environment within the university community.

## Tech and Stack
Description of the technologies and stack utilized in this project:

### React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

### Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type aware lint rules:

- Configure the top-level `parserOptions` property like this:

```js
export default {
  // other rules...
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
    project: ['./tsconfig.json', './tsconfig.node.json'],
    tsconfigRootDir: __dirname,
  },
}
```

- Replace `plugin:@typescript-eslint/recommended` to `plugin:@typescript-eslint/recommended-type-checked` or `plugin:@typescript-eslint/strict-type-checked`
- Optionally add `plugin:@typescript-eslint/stylistic-type-checked`
- Install [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) and add `plugin:react/recommended` & `plugin:react/jsx-runtime` to the `extends` list
