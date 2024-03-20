# Folder Structure Documentation

This document outlines the directory structure of the `frontend` of the `Meja Belajar` web application.

## Folder Tree
```
frontend
|
├───docs
├───public
└───src
    ├───apis
    │   └───services
    ├───assets
    │   └───images
    │       ├───icon
    │       └───social
    ├───components
    ├───contexts
    ├───data
    ├───helpers
    │   ├───configs
    │   ├───hooks
    │   └───utils
    ├───models
    │   ├───database
    │   └───zod
    └───pages
```


## apis
This directory contains modules responsible for handling API requests and interactions. File name will use `camelCase`.
- **services**: Houses service files that encapsulate API endpoints and related functionality.

## assets
This directory stores static assets such as `images`, `fonts`, and `animation component` used in the application.
- **images**: Contains image assets used throughout the application.
  - **icon**: Stores icons used within the application.
  - **social**: Holds social media icons or images.

## components
Contains reusable React components that can be used across different pages or sections of the application. File name will use `PascalCase`.

## contexts
Houses React context providers used for managing `global state` within the application. File name will use `PascalCase`.

## data
Contains data files or modules used by the application, such as JSON files, mock data, or data manipulation utilities. File name will use `camelCase`.

## helpers
This directory houses helper functions and utilities used throughout the application. File name will use `camelCase`.
- **configs**: Stores configuration files or constants used across the application.
- **hooks**: Contains custom React hooks used for state management, side effects, or other common patterns.
- **utils**: Houses miscellaneous utility functions or components that do not fit into other specific categories.

## models
Contains data `models` or `schemas` used within the application. File name will use `camelCase`.
- **database**: Holds database-related models or schemas.
- **zod**: Stores data validation schemas using the Zod library.

## Pages
The `pages` directory houses top-level React components representing individual pages or views of the application. Pages naming will use `PascalCase`. Every page that is built with multiple components will be encapsulated into one folder using the parent page's name. For example:
- The `Landing` folder may contain one parent component (`Landing.tsx`) and three subcomponents (`LandingGuest.tsx`, `LandingUser.tsx`, `LandingMentor.tsx`).
